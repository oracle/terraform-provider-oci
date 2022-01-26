// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package acctest

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
	"text/template"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"

	"github.com/hashicorp/go-multierror"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"

	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	tf_provider "github.com/terraform-providers/terraform-provider-oci/internal/provider"

	//tf_resource_discovery "github.com/terraform-providers/terraform-provider-oci/oci/resourcediscovery"
	tf_resource "github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var tmpl template.Template = *template.New("tmpl")
var lineSeparator = "\n"

// Applies values from a map to a string template
func apply(template string, values map[string]string) string {
	b := new(bytes.Buffer)
	t, _ := tmpl.Parse(template)
	if err := t.Execute(b, values); err != nil {
		log.Printf("[WARN] Unable to apply values to template: '%s'", err)
	}
	return b.String()
}

type TokenFn func(string, map[string]string) string

// Creates a form of "apply" above that will always supply the same value for {{.token}} use hard code value for HTTP replay
func TokenizeWithHttpReplay(defaultString string) (string, TokenFn) {
	var ts string
	if httpreplay.ModeRecordReplay() {
		ts = defaultString
	} else {
		ts = utils.Timestamp()
	}
	return ts, func(template string, values map[string]string) string {
		if values == nil {
			values = map[string]string{}
		}
		values["token"] = ts
		return apply(template, values)
	}
}

// custom TestCheckFunc helper, returns a value associated with a key from an instance in the current state
func FromInstanceState(s *terraform.State, name, key string) (string, error) {
	ms := s.RootModule()
	rs, ok := ms.Resources[name]
	if !ok {
		return "", fmt.Errorf("Not found: %s", name)
	}

	is := rs.Primary
	if is == nil {
		return "", fmt.Errorf("No primary instance: %s", name)
	}

	v, ok := is.Attributes[key]

	if ok {
		return v, nil
	} else {
		return "", fmt.Errorf("%s: Attribute '%s' not found", name, key)
	}
}

// TestCheckResourceAttributesEqual is a TestCheckFunc which ensures that the values of two
// attributes in two different resources are equal.
func TestCheckResourceAttributesEqual(name1, key1, name2, key2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		val1, err := FromInstanceState(s, name1, key1)
		if err != nil {
			return err
		}

		val2, err := FromInstanceState(s, name2, key2)
		if err != nil {
			return err
		}

		if val1 != val2 {
			return fmt.Errorf(
				"%s: attribute '%s' value %#v not equal to '%s' attribute '%s' value %#v",
				name1, key1, val1, name2, key2, val2)
		}
		return nil
	}
}

func TestCheckAttributeBase64Encoded(name, key string, expectBase64Encoded bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		content, err := FromInstanceState(s, name, key)
		if err != nil {
			return err
		}

		isBase64Encoded := true
		if _, err := base64.StdEncoding.DecodeString(content); err != nil {
			isBase64Encoded = false
		}

		if isBase64Encoded != expectBase64Encoded {
			if expectBase64Encoded {
				return fmt.Errorf("expected '%s' to be base64 encoded, but it is not", key)
			}
			return fmt.Errorf("expected '%s' to not be base64 encoded, but it is", key)
		}

		return nil
	}
}

// Function to be implemented by resources that wish to wait on a certain condition and this function is responsible for evaluating the specific condition for that resource
type ShouldWaitFunc func(response oci_common.OCIOperationResponse) bool

// Function to be implemented by resources that wish to wait on a certain condition and this function is responsible for fetching the latest state using the resourceId
type FetchOperationFunc func(client *tf_client.OracleClients, resourceId *string, retryPolicy *oci_common.RetryPolicy) error

// This function waits for the given time and retries the ShouldWaitFunc and periodically invokes the FetchOperationFunc to fetch the latest response
func WaitTillCondition(testAccProvider *schema.Provider, resourceId *string, shouldWait ShouldWaitFunc, timeout time.Duration,
	fetchOperationFunc FetchOperationFunc, service string, disableNotFoundRetries bool) func() {
	return func() {
		client := testAccProvider.Meta().(*tf_client.OracleClients)
		log.Printf("[INFO] start of WaitTillCondition for resource %s ", *resourceId)
		retryPolicy := tf_resource.GetRetryPolicy(disableNotFoundRetries, service)
		retryPolicy.ShouldRetryOperation = ConditionShouldRetry(timeout, shouldWait, service, disableNotFoundRetries)

		err := fetchOperationFunc(client, resourceId, retryPolicy)
		if err != nil {
			log.Printf("[WARN] WaitTillCondition failed for %s resource with error %v", *resourceId, err)
		} else {
			log.Printf("[INFO] end of WaitTillCondition for resource %s ", *resourceId)
		}
	}
}

// This function is responsible for the actual check for ShouldWaitFunc and the aborting
func ConditionShouldRetry(timeout time.Duration, shouldWait ShouldWaitFunc, service string, disableNotFoundRetries bool, optionals ...interface{}) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		//Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		//Make sure we stop on default rules
		if tf_resource.ShouldRetry(response, disableNotFoundRetries, service, startTime, optionals...) {
			return true
		}

		return shouldWait(response)
	}
}

type RepresentationType int

const (
	Required RepresentationType = iota + 1
	Optional
)

type RepresentationMode int

const (
	Create RepresentationMode = iota
	Update
)

type Representation struct {
	RepType RepresentationType
	Create  interface{}
	Update  interface{}
}

type RepresentationGroup struct {
	RepType RepresentationType
	Group   map[string]interface{}
}

func cloneRepresentation(representations map[string]interface{}) map[string]interface{} {
	copyMap := map[string]interface{}{}

	for key, value := range representations {
		representation, ok := value.(Representation)
		if ok {
			copyMap[key] = Representation{representation.RepType, representation.Create, representation.Update}
		}
		representationGroup, ok := value.(RepresentationGroup)
		if ok {
			copyMap[key] = RepresentationGroup{representationGroup.RepType, cloneRepresentation(representationGroup.Group)}
		}
		representationGroupArr, ok := value.([]RepresentationGroup)
		if ok {
			representationGroupArrClone := make([]RepresentationGroup, len(representationGroupArr))
			for index, representationGroupItem := range representationGroupArr {
				representationGroupArrClone[index] = RepresentationGroup{representationGroup.RepType, cloneRepresentation(representationGroupItem.Group)}
			}
			copyMap[key] = representationGroupArrClone
		}
	}

	return copyMap
}

func RepresentationCopyWithRemovedProperties(representations map[string]interface{}, removedProperties []string) map[string]interface{} {
	representationsCopy := cloneRepresentation(representations)
	for _, propName := range removedProperties {
		delete(representationsCopy, propName)
	}
	return representationsCopy
}

func RepresentationCopyWithNewProperties(representations map[string]interface{}, newProperties map[string]interface{}) map[string]interface{} {
	representationsCopy := cloneRepresentation(representations)
	for propName, value := range newProperties {
		representationsCopy[propName] = value
	}
	return representationsCopy
}

func GetUpdatedRepresentationCopy(propertyNameStr string, newValue interface{}, representations map[string]interface{}) map[string]interface{} {
	propertyNames := strings.Split(propertyNameStr, ".")
	return updateNestedRepresentation(0, propertyNames, newValue, cloneRepresentation(representations))
}

func GetMultipleUpdatedRepresenationCopy(propertyNames []string, newValues []interface{}, representations map[string]interface{}) map[string]interface{} {
	for i := 0; i < len(propertyNames); i++ {
		representations = GetUpdatedRepresentationCopy(propertyNames[i], newValues[i], representations)
	}
	return representations
}

func updateNestedRepresentation(currIndex int, propertyNames []string, newValue interface{}, representations map[string]interface{}) map[string]interface{} {
	//recursively search the property to Update
	for prop := range representations {
		if prop == propertyNames[currIndex] {
			representationGroup, ok := representations[prop].(RepresentationGroup)
			if ok && currIndex+1 < len(propertyNames) {
				updateNestedRepresentation(currIndex+1, propertyNames, newValue, representationGroup.Group)
			} else {
				representations[prop] = newValue
			}
			return representations
		}
	}

	return nil
}

// removes the list of properties at nested level(given the full qualified name) from the representation map
// example for fully qualified name of a nested level property: "specification.request_policies.authentication.audiences"
func GetRepresentationCopyWithMultipleRemovedProperties(propertyNames []string, representation map[string]interface{}) map[string]interface{} {
	for i := 0; i < len(propertyNames); i++ {
		representation = RepresentationCopyWithRemovedNestedProperties(propertyNames[i], representation)
	}
	return representation
}

func RepresentationCopyWithRemovedNestedProperties(propertyNameStr string, representation map[string]interface{}) map[string]interface{} {
	propertyNames := strings.Split(propertyNameStr, ".")
	return updateNestedRepresentationRemoveProperty(0, propertyNames, cloneRepresentation(representation))
}

func updateNestedRepresentationRemoveProperty(currIndex int, propertyNames []string, representation map[string]interface{}) map[string]interface{} {
	//recursively search the property to remove
	for prop := range representation {
		if prop == propertyNames[currIndex] {
			representationGroup, ok := representation[prop].(RepresentationGroup)
			if ok && currIndex+1 < len(propertyNames) {
				updateNestedRepresentationRemoveProperty(currIndex+1, propertyNames, representationGroup.Group)
			} else {
				delete(representation, prop)
			}
			return representation
		}
	}
	return nil
}

func GenerateDataSourceFromRepresentationMap(resourceType string, resourceName string, representationType RepresentationType, representationMode RepresentationMode, representations map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(`%sdata "%s" "%s" %s`, lineSeparator, resourceType, resourceName, GenerateResourceFromMap(representationType, representationMode, representations)))
	return buffer.String()
}

func GenerateResourceFromRepresentationMap(resourceType string, resourceName string, representationType RepresentationType, representationMode RepresentationMode, representations map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(`%sresource "%s" "%s" %s`, lineSeparator, resourceType, resourceName, GenerateResourceFromMap(representationType, representationMode, representations)))
	return buffer.String()
}

func GenerateResourceFromMap(representationType RepresentationType, representationMode RepresentationMode, representations map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString("{" + lineSeparator)

	sortedRepresentations := make([]string, 0, len(representations))
	for key := range representations {
		sortedRepresentations = append(sortedRepresentations, key)
	}
	sort.Strings(sortedRepresentations)

	for _, prop := range sortedRepresentations {
		representation, ok := representations[prop].(Representation)
		if ok && representation.RepType <= representationType {

			representationValue := representation.Create
			if representationMode == Update && representation.Update != nil {
				representationValue = representation.Update
			}

			repStrValue, strRep := representationValue.(string)
			if strRep {
				buffer.WriteString(fmt.Sprintf(`%s = "%s"%s`, prop, repStrValue, lineSeparator))
			}

			repArrayStrValue, arrayRep := representationValue.([]string)
			if arrayRep {
				var repArrayStrEscValue []string
				for _, arrayValue := range repArrayStrValue {
					repArrayStrEscValue = append(repArrayStrEscValue, fmt.Sprintf(`"%s"`, arrayValue))
				}
				buffer.WriteString(fmt.Sprintf(`%s = [%s]%s`, prop, strings.Join(repArrayStrEscValue, ", "), lineSeparator))
			}

			repMapStrValue, mapRep := representationValue.(map[string]string)
			if mapRep {
				sortedKeys := make([]string, 0, len(repMapStrValue))
				for key := range repMapStrValue {
					sortedKeys = append(sortedKeys, key)
				}
				sort.Strings(sortedKeys)

				var repMapStrEscValue []string
				for _, key := range sortedKeys {
					repMapStrEscValue = append(repMapStrEscValue, fmt.Sprintf(`"%s" = "%s"`, key, repMapStrValue[key]))
				}
				buffer.WriteString(fmt.Sprintf("%s = {\n%s\n}%s", prop, strings.Join(repMapStrEscValue, lineSeparator), lineSeparator))

			}

		}
		representationGroup, ok := representations[prop].(RepresentationGroup)
		if ok && representationGroup.RepType <= representationType {
			buffer.WriteString(fmt.Sprintf("%s %s", prop, GenerateResourceFromMap(representationType, representationMode, representationGroup.Group)))
		}
		representationGroupArray, ok := representations[prop].([]RepresentationGroup)
		if ok {
			for _, representationGroupInArray := range representationGroupArray {
				if representationGroupInArray.RepType <= representationType {
					buffer.WriteString(fmt.Sprintf("%s %s", prop, GenerateResourceFromMap(representationType, representationMode, representationGroupInArray.Group)))
				}
			}
		}
	}
	buffer.WriteString(fmt.Sprintf("}%s", lineSeparator))
	return buffer.String()
}

func setEnvSetting(s, v string) error {
	error := os.Setenv(s, v)
	if error != nil {
		return fmt.Errorf("Failed to set env setting '%s', encountered error: %v", s, error)
	}
	return nil
}

/*

 */
func CheckJsonStringsEqual(expectedJsonString string, actualJsonString string) error {
	if expectedJsonString == actualJsonString {
		return nil
	}

	var expected, actual interface{}
	if err := json.Unmarshal([]byte(expectedJsonString), &expected); err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(actualJsonString), &actual); err != nil {
		return err
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON '%s', but got JSON '%s'", expectedJsonString, actualJsonString)
	}
	return nil
}

// Compares an attribute against a JSON string's unmarshalled value
func TestCheckJsonResourceAttr(name, key, expectedJson string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actualJsonFromState, err := FromInstanceState(s, name, key)
		if err != nil {
			return err
		}

		if err := CheckJsonStringsEqual(expectedJson, actualJsonFromState); err != nil {
			return fmt.Errorf("%s: Attribute '%s' %s", name, key, err)
		}
		return nil
	}
}

func SaveConfigContent(content string, service string, resource string, t *testing.T) {
	if strings.ToLower(utils.GetEnvSettingWithBlankDefault("save_configs")) == "true" {
		if len(content) > 0 {
			if err := WriteToFile(content, service, resource); err != nil {
				log.Printf("Failed to write TF content to file with error: %q", err)
			}
		}

		t.Skip()
	}
}

func WriteToFile(content string, service string, resource string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	path = path + "/output/" + service + "/"
	if err := os.MkdirAll(path, 0770); err != nil {
		return err
	}
	f, err := os.OpenFile(path+"/"+resource+".tf", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	if _, err = f.WriteString(content); err != nil {
		return err
	}
	return nil
}

func GenericTestStepPreConfiguration(stepNumber int) func() {
	return func() {

		// add logs for notifying execution
		log.Println()
		log.Printf("====================== Executing Test Step %d ===================", stepNumber)
		log.Println()
	}
}

/*
	This struct extends the HashiCorp plugin framework testing.T
	It adds a slice to store all error messages encountered during test execution
*/
type OciTestT struct {
	T             *testing.T
	ErrorMessages []string
}

func (t *OciTestT) Error(args ...interface{}) {
	t.T.Error(args...)
	str := fmt.Sprintf("%v", args)
	t.ErrorMessages = append(t.ErrorMessages, str)
}

func (t *OciTestT) Fatal(args ...interface{}) {
	t.T.Fatal(args...)
	str := fmt.Sprintf("%v", args)
	t.ErrorMessages = append(t.ErrorMessages, str)
}

func (t *OciTestT) Skip(args ...interface{}) {
	t.T.Skip(args...)
}

func (t *OciTestT) Name() string {
	return t.T.Name()
}

func (t *OciTestT) Parallel() {
	t.T.Parallel()
}

// Method to execute tests
func ResourceTest(t *testing.T, checkDestroyFunc resource.TestCheckFunc, steps []resource.TestStep) {
	// set Generic preconfiguration method if not explicitly set
	for index, _ := range steps {
		if steps[index].PreConfig == nil {
			steps[index].PreConfig = GenericTestStepPreConfiguration(index)
		}
	}

	ociTest := OciTestT{t, make([]string, 0)}
	resource.Test(&ociTest, resource.TestCase{
		PreCheck: func() { PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": TestAccProvider,
		},
		CheckDestroy: checkDestroyFunc,
		Steps:        steps,
	})

	// check if any error was logged
	if len(ociTest.ErrorMessages) <= 0 {
		return
	}

	fmt.Println("================ Error Summary ================")
	// print out the errors in an error summary
	for _, error := range ociTest.ErrorMessages {
		fmt.Println(error)
	}
}

func PreCheck(t *testing.T) {
	envVarChecklist := []string{}
	copy(envVarChecklist, requiredTestEnvVars)
	if utils.GetEnvSettingWithDefault("use_obo_token", "false") != "false" {
		envVarChecklist = append(envVarChecklist, requiredOboTokenAuthEnvVars...)
	} else {
		envVarChecklist = append(envVarChecklist, requiredKeyAuthEnvVars...)
	}

	for _, envVar := range envVarChecklist {
		if v := utils.GetEnvSettingWithBlankDefault(envVar); v == "" {
			t.Fatal("TF_VAR_" + envVar + " must be set for acceptance tests")
		}
	}

}

var requiredTestEnvVars = []string{"compartment_ocid", "compartment_id_for_create", "compartment_id_for_update", "tags_import_if_exists"}
var requiredKeyAuthEnvVars = []string{"tenancy_ocid", "user_ocid", "fingerprint"}
var requiredOboTokenAuthEnvVars = []string{"tenancy_ocid", "obo_token"}
var TestAccProvider *schema.Provider
var TestAccProviders map[string]terraform.ResourceProvider

const (
	requestQueryOpcTimeMaintenanceRebootDue = "opc-time-maintenance-reboot-due"
)

// Provider is the adapter for terraform, that gives access to all the resources
func ProviderTestCopy(configfn schema.ConfigureFunc) terraform.ResourceProvider {
	result := &schema.Provider{
		DataSourcesMap: tf_provider.DataSourcesMap(),
		Schema:         tf_provider.SchemaMap(),
		ResourcesMap:   tf_provider.ResourcesMap(),
		ConfigureFunc:  configfn,
	}

	// Additions for test parameters
	result.Schema["test_time_maintenance_reboot_due"] = &schema.Schema{Type: schema.TypeString, Optional: true}

	return result
}

func ProviderTestConfig() string {
	return `
	# Need to have this block even though it's empty; for import testing
	provider "oci" {
	}
	` + CommonTestVariables()
}

func CommonTestVariables() string {
	return `
	variable "tenancy_ocid" {
		default = "` + utils.GetEnvSettingWithBlankDefault("tenancy_ocid") + `"
	}

	variable "ssh_public_key" {
		default = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"
	}

	variable "region" {
		default = "` + utils.GetEnvSettingWithBlankDefault("region") + `"
	}

	`
}

func GetTestClients(data *schema.ResourceData) *tf_client.OracleClients {
	r := &schema.Resource{
		Schema: tf_provider.SchemaMap(),
	}
	d := r.Data(nil)
	d.SetId(utils.GetEnvSettingWithBlankDefault("tenancy_ocid"))
	d.Set("tenancy_ocid", utils.GetEnvSettingWithBlankDefault("tenancy_ocid"))
	d.Set("region", utils.GetEnvSettingWithDefault("region", "us-phoenix-1"))

	if auth := utils.GetEnvSettingWithDefault("auth", globalvar.AuthAPIKeySetting); auth == globalvar.AuthAPIKeySetting {
		d.Set("auth", utils.GetEnvSettingWithDefault("auth", globalvar.AuthAPIKeySetting))
		d.Set("user_ocid", utils.GetEnvSettingWithBlankDefault("user_ocid"))
		d.Set("fingerprint", utils.GetEnvSettingWithBlankDefault("fingerprint"))
		d.Set("private_key_path", utils.GetEnvSettingWithBlankDefault("private_key_path"))
		d.Set("private_key_password", utils.GetEnvSettingWithBlankDefault("private_key_password"))
		d.Set("private_key", utils.GetEnvSettingWithBlankDefault("private_key"))
	} else {
		d.Set("auth", utils.GetEnvSettingWithDefault("auth", auth))
	}

	tf_provider.TerraformCLIVersion = globalvar.TestTerraformCLIVersion
	client, err := tf_provider.ProviderConfig(d)
	if err != nil {
		panic(err)
	}

	// This is a test hook to support creating instances that have a maintenance reboot time set
	// The test hook allows 'time_maintenance_reboot_due' field to be tested for instance datasources/resources
	// This is controlled by a provider option rather than environment variable: so that the tests can run in parallel
	// without affecting one another and also allow individual test steps to alter this
	//
	// If we have additional test hooks that need to be supported in this manner, then the following logic should be
	// compartmentalized and registered with the test provider in a scalable manner.
	/*maintenanceRebootTime, ok := data.GetOkExists("test_time_maintenance_reboot_due")
	if ok {
		computeClient := client.(*tf_client.OracleClients).computeClient()
		baseInterceptor := computeClient.Interceptor
		computeClient.Interceptor = func(r *http.Request) error {
			if err := baseInterceptor(r); err != nil {
				return err
			}

			if r.Method == http.MethodPost && (strings.Contains(r.URL.Path, "/instances")) {
				query := r.URL.Query()
				query.Set(requestQueryOpcTimeMaintenanceRebootDue, maintenanceRebootTime.(string))
				r.URL.RawQuery = query.Encode()
			}
			return nil
		}
	}*/

	return client.(*tf_client.OracleClients)
}

// wrapper over resource.ComposeAggregateTestCheckFunc to use customErrorFormat for multierror
func ComposeAggregateTestCheckFuncWrapper(fs ...resource.TestCheckFunc) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var result *multierror.Error

		for i, f := range fs {
			if err := f(s); err != nil {
				result = multierror.Append(result, fmt.Errorf("Check %d/%d error: %s", i+1, len(fs), err))
			}
		}

		err := result.ErrorOrNil()
		if err != nil {
			result.ErrorFormat = utils.CustomErrorFormat
		}

		return err
	}
}

func ProviderConfigTest(t *testing.T, disableRetries bool, skipRequiredField bool, auth string, configFileProfile string, configFunc ConfigFunc) {
	r := &schema.Resource{
		Schema: tf_provider.SchemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")
	d.Set("auth", auth)
	if !skipRequiredField {
		d.Set("tenancy_ocid", testTenancyOCID)
	}
	if configFileProfile == "" || configFileProfile == "DEFAULT" {
		d.Set("user_ocid", testUserOCID)
		d.Set("fingerprint", testKeyFingerPrint)
		d.Set("private_key", testPrivateKey)
		//d.Set("private_key_path", "")
		d.Set("region", "us-phoenix-1")
		d.Set("private_key_password", "password")
	}
	if configFileProfile == "PROFILE3" {
		d.Set("fingerprint", testKeyFingerPrint)
	}
	if disableRetries {
		d.Set("disable_auto_retries", disableRetries)
	}
	if configFileProfile != "" {
		d.Set("config_file_profile", configFileProfile)
	}

	// Use config func for export (resource discovery)
	configureProviderFn := configFunc
	//userAgent := fmt.Sprintf(globalvar.ExportUserAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, globalvar.Version)

	// If no ConfigFunc use ProviderConfig
	if configureProviderFn == nil {
		configureProviderFn = tf_provider.ProviderConfig
		//userAgent = fmt.Sprintf(globalvar.UserAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, sdkMeta.SDKVersionString(), provider.TerraformCLIVersion, globalvar.DefaultUserAgentProviderName, globalvar.Version)

	}
	client, err := configureProviderFn(d)

	if configFileProfile == "wrongProfile" {
		assert.Equal(t, "configuration file did not contain profile: wrongProfile", err.Error())
		return
	}
	if configFileProfile == "PROFILE2" {
		assert.Equal(t, "can not Create client, bad configuration: did not find a proper configuration for private key", err.Error())
		return
	}
	switch auth {
	case globalvar.AuthAPIKeySetting, "":
		if skipRequiredField {
			assert.Equal(t, err, nil)
			return
		}
	default:
		assert.Error(t, err, fmt.Sprintf("auth must be one of '%s' or '%s' or '%s'", globalvar.AuthAPIKeySetting, globalvar.AuthInstancePrincipalSetting, globalvar.AuthInstancePrincipalWithCertsSetting))
		return
	}
	assert.Nil(t, err)
	assert.NotNil(t, client)

	oracleClient, ok := client.(*tf_client.OracleClients)
	assert.True(t, ok)

	testClient := func(c *oci_common.BaseClient) {
		assert.NotNil(t, c)
		assert.NotNil(t, c.HTTPClient)
		assert.Exactly(t, c.UserAgent, globalvar.UserAgentFormatter)
		assert.NotNil(t, c.Interceptor)
		assert.NotNil(t, c.Signer)
	}

	testClient(&oracleClient.IdentityClient().BaseClient)
}

/* This function is used in the test asserts to verify that an element in a set contains certain properties
 * properties is a map of nameOfProperty -> expectedValueOfProperty
 * presentProperties is an array of property names that are expected to be set in the set element but we don't care about matching the value
 * will return nil (the positive response) if there is an element in the set that matches all properties in properties and presentProperties
 */
func CheckResourceSetContainsElementWithProperties(name, setKey string, properties map[string]string, presentProperties []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rm := s.RootModule()
		rs, ok := rm.Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		is := rs.Primary
		if is == nil {
			return fmt.Errorf("No primary instance: %s", name)
		}

		orderedKeys := []string{}
		for key, _ := range is.Attributes {
			orderedKeys = append(orderedKeys, key)
		}
		sort.Strings(orderedKeys)
		var currSetElementId string
		currMatchedAttributes := []string{}
		currMatchedPresentProperties := []string{}
		setElementMatch := func() bool {
			return len(currMatchedAttributes) == len(properties) && (presentProperties == nil || len(currMatchedPresentProperties) == len(presentProperties))
		}
		for _, key := range orderedKeys {
			prefix := fmt.Sprintf("%s.", setKey)
			if !strings.HasPrefix(key, prefix) {
				continue
			}
			attrWithSetIdRaw := strings.TrimPrefix(key, prefix)

			attrWithSetIdRawArr := strings.Split(attrWithSetIdRaw, ".")
			if len(attrWithSetIdRawArr) < 2 {
				continue
			}
			if currSetElementId == "" {
				currSetElementId = attrWithSetIdRawArr[0]
			}
			if attrWithSetIdRawArr[0] != currSetElementId {
				if setElementMatch() {
					return nil
				}
				currMatchedPresentProperties = []string{}
				currMatchedAttributes = []string{}
				currSetElementId = attrWithSetIdRawArr[0]
			}
			attributeName := strings.Join(attrWithSetIdRawArr[1:], ".")
			for propName, value := range properties {
				if propName == attributeName && value == is.Attributes[key] {
					currMatchedAttributes = append(currMatchedAttributes, propName)
				}
			}
			if presentProperties != nil {
				for _, propName := range presentProperties {
					if propName == attributeName {
						currMatchedPresentProperties = append(currMatchedPresentProperties, propName)
					}
				}
			}
		}
		if setElementMatch() {
			return nil
		}

		return fmt.Errorf("%s: Set Attribute '%s' does not contain an element with attributes %v %v\nAttributesInStatefile: %v", name, setKey, properties, presentProperties, is.Attributes)
	}
}

func CheckResourceSetContainsElementWithPropertiesContainingNestedSets(name, setKey string, properties map[string]interface{}, presentProperties []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rm := s.RootModule()
		rs, ok := rm.Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		is := rs.Primary
		if is == nil {
			return fmt.Errorf("No primary instance: %s", name)
		}

		orderedKeys := []string{}
		for key, _ := range is.Attributes {
			orderedKeys = append(orderedKeys, key)
		}
		sort.Strings(orderedKeys)
		var currSetElementId string
		currMatchedAttributes := []string{}
		currMatchedPresentProperties := []string{}
		leafProperties := 0
		for _, value := range properties {
			if _, ok := value.(string); ok {
				leafProperties++
			}
		}
		setElementMatch := func() bool {
			return len(currMatchedAttributes) == leafProperties && (presentProperties == nil || len(currMatchedPresentProperties) == len(presentProperties))
		}
		for _, key := range orderedKeys {
			prefix := fmt.Sprintf("%s.", setKey)
			if !strings.HasPrefix(key, prefix) {
				continue
			}
			attrWithSetIdRaw := strings.TrimPrefix(key, prefix)

			attrWithSetIdRawArr := strings.Split(attrWithSetIdRaw, ".")
			if len(attrWithSetIdRawArr) < 2 {
				continue
			}
			if attrWithSetIdRawArr[0] != currSetElementId {
				if setElementMatch() {
					return nil
				}
				currMatchedPresentProperties = []string{}
				currMatchedAttributes = []string{}
				currSetElementId = attrWithSetIdRawArr[0]

				//check nested set properties, we do it in this if statement to avoid repeating the same checks for each key in the loop. We only need to check once per set element id
				for propName, value := range properties {
					if valueSet, ok := value.([]map[string]interface{}); ok {
						for _, nestedSetElement := range valueSet {
							nestedSetCheck := CheckResourceSetContainsElementWithPropertiesContainingNestedSets(name, fmt.Sprintf("%s.%s.%s", setKey, currSetElementId, propName), nestedSetElement, nil)
							if err := nestedSetCheck(s); err != nil {
								return err
							}
						}
					}
				}
			}
			attributeName := strings.Join(attrWithSetIdRawArr[1:], ".")
			for propName, value := range properties {
				if valueStr, ok := value.(string); ok {
					if propName == attributeName && valueStr == is.Attributes[key] {
						currMatchedAttributes = append(currMatchedAttributes, propName)
					}
				}
			}
			if presentProperties != nil {
				for _, propName := range presentProperties {
					if propName == attributeName {
						currMatchedPresentProperties = append(currMatchedPresentProperties, propName)
					}
				}
			}
		}
		if setElementMatch() {
			return nil
		}

		return fmt.Errorf("%s: Set Attribute '%s' does not contain an element with attributes %v %v\nAttributesInStatefile: %v", name, setKey, properties, presentProperties, is.Attributes)
	}
}

func TestAccPreCheck(t *testing.T) {
	envVarChecklist := []string{}
	copy(envVarChecklist, requiredTestEnvVars)
	if utils.GetEnvSettingWithDefault("use_obo_token", "false") != "false" {
		envVarChecklist = append(envVarChecklist, requiredOboTokenAuthEnvVars...)
	} else {
		envVarChecklist = append(envVarChecklist, requiredKeyAuthEnvVars...)
	}

	for _, envVar := range envVarChecklist {
		assertEnvAvailable(envVar, t)
	}
}
