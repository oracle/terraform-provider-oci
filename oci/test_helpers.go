// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

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
	"text/template"
	"time"

	"github.com/hashicorp/terraform/command"
	"github.com/mitchellh/cli"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	oci_common "github.com/oracle/oci-go-sdk/common"
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

// Returns date-time formatted as a string, ex: 2017-10-12-000934-119299083"
func timestamp() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d-%02d%02d%02d-%d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
}

type TokenFn func(string, map[string]string) string

// Creates a form of "apply" above that will always supply the same value for {{.token}} use hard code value for HTTP replay
func tokenizeWithHttpReplay(defaultString string) (string, TokenFn) {
	var ts string
	if httpreplay.ModeRecordReplay() {
		ts = defaultString
	} else {
		ts = timestamp()
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
func fromInstanceState(s *terraform.State, name, key string) (string, error) {
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
		val1, err := fromInstanceState(s, name1, key1)
		if err != nil {
			return err
		}

		val2, err := fromInstanceState(s, name2, key2)
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

func testCheckAttributeBase64Encoded(name, key string, expectBase64Encoded bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		content, err := fromInstanceState(s, name, key)
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
type FetchOperationFunc func(client *OracleClients, resourceId *string, retryPolicy *oci_common.RetryPolicy) error

// This function waits for the given time and retries the ShouldWaitFunc and periodically invokes the FetchOperationFunc to fetch the latest response
func waitTillCondition(testAccProvider *schema.Provider, resourceId *string, shouldWait ShouldWaitFunc, timeout time.Duration,
	fetchOperationFunc FetchOperationFunc, service string, disableNotFoundRetries bool) func() {
	return func() {
		client := testAccProvider.Meta().(*OracleClients)
		log.Printf("[INFO] start of waitTillCondition for resource %s ", *resourceId)
		retryPolicy := getRetryPolicy(disableNotFoundRetries, service)
		retryPolicy.ShouldRetryOperation = conditionShouldRetry(timeout, shouldWait, service, disableNotFoundRetries)

		err := fetchOperationFunc(client, resourceId, retryPolicy)
		if err != nil {
			log.Printf("[WARN] waitTillCondition failed for %s resource with error %v", *resourceId, err)
		} else {
			log.Printf("[INFO] end of waitTillCondition for resource %s ", *resourceId)
		}
	}
}

// This function is responsible for the actual check for ShouldWaitFunc and the aborting
func conditionShouldRetry(timeout time.Duration, shouldWait ShouldWaitFunc, service string, disableNotFoundRetries bool, optionals ...interface{}) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		//Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		//Make sure we stop on default rules
		if shouldRetry(response, disableNotFoundRetries, service, startTime, optionals...) {
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
	repType RepresentationType
	create  interface{}
	update  interface{}
}

type RepresentationGroup struct {
	repType RepresentationType
	group   map[string]interface{}
}

func cloneRepresentation(representations map[string]interface{}) map[string]interface{} {
	copyMap := map[string]interface{}{}

	for key, value := range representations {
		representation, ok := value.(Representation)
		if ok {
			copyMap[key] = Representation{representation.repType, representation.create, representation.update}
		}
		representationGroup, ok := value.(RepresentationGroup)
		if ok {
			copyMap[key] = RepresentationGroup{representationGroup.repType, cloneRepresentation(representationGroup.group)}
		}
		representationGroupArr, ok := value.([]RepresentationGroup)
		if ok {
			representationGroupArrClone := make([]RepresentationGroup, len(representationGroupArr))
			for index, representationGroupItem := range representationGroupArr {
				representationGroupArrClone[index] = RepresentationGroup{representationGroup.repType, cloneRepresentation(representationGroupItem.group)}
			}
			copyMap[key] = representationGroupArrClone
		}
	}

	return copyMap
}

func representationCopyWithRemovedProperties(representations map[string]interface{}, removedProperties []string) map[string]interface{} {
	representationsCopy := cloneRepresentation(representations)
	for _, propName := range removedProperties {
		delete(representationsCopy, propName)
	}
	return representationsCopy
}

func representationCopyWithNewProperties(representations map[string]interface{}, newProperties map[string]interface{}) map[string]interface{} {
	representationsCopy := cloneRepresentation(representations)
	for propName, value := range newProperties {
		representationsCopy[propName] = value
	}
	return representationsCopy
}

func getUpdatedRepresentationCopy(propertyNameStr string, newValue interface{}, representations map[string]interface{}) map[string]interface{} {
	propertyNames := strings.Split(propertyNameStr, ".")
	return updateNestedRepresentation(0, propertyNames, newValue, cloneRepresentation(representations))
}

func getMultipleUpdatedRepresenationCopy(propertyNames []string, newValues []interface{}, representations map[string]interface{}) map[string]interface{} {
	for i := 0; i < len(propertyNames); i++ {
		representations = getUpdatedRepresentationCopy(propertyNames[i], newValues[i], representations)
	}
	return representations
}

func updateNestedRepresentation(currIndex int, propertyNames []string, newValue interface{}, representations map[string]interface{}) map[string]interface{} {
	//recursively search the property to update
	for prop := range representations {
		if prop == propertyNames[currIndex] {
			representationGroup, ok := representations[prop].(RepresentationGroup)
			if ok && currIndex+1 < len(propertyNames) {
				updateNestedRepresentation(currIndex+1, propertyNames, newValue, representationGroup.group)
			} else {
				representations[prop] = newValue
			}
			return representations
		}
	}

	return nil
}

func generateDataSourceFromRepresentationMap(resourceType string, resourceName string, representationType RepresentationType, representationMode RepresentationMode, representations map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(`%sdata "%s" "%s" %s`, lineSeparator, resourceType, resourceName, generateResourceFromMap(representationType, representationMode, representations)))
	return buffer.String()
}

func generateResourceFromRepresentationMap(resourceType string, resourceName string, representationType RepresentationType, representationMode RepresentationMode, representations map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(`%sresource "%s" "%s" %s`, lineSeparator, resourceType, resourceName, generateResourceFromMap(representationType, representationMode, representations)))
	return buffer.String()
}

func generateResourceFromMap(representationType RepresentationType, representationMode RepresentationMode, representations map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString("{" + lineSeparator)

	sortedRepresentations := make([]string, 0, len(representations))
	for key := range representations {
		sortedRepresentations = append(sortedRepresentations, key)
	}
	sort.Strings(sortedRepresentations)

	for _, prop := range sortedRepresentations {
		representation, ok := representations[prop].(Representation)
		if ok && representation.repType <= representationType {

			representationValue := representation.create
			if representationMode == Update && representation.update != nil {
				representationValue = representation.update
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
		if ok && representationGroup.repType <= representationType {
			buffer.WriteString(fmt.Sprintf("%s %s", prop, generateResourceFromMap(representationType, representationMode, representationGroup.group)))
		}
		representationGroupArray, ok := representations[prop].([]RepresentationGroup)
		if ok {
			for _, representationGroupInArray := range representationGroupArray {
				if representationGroupInArray.repType <= representationType {
					buffer.WriteString(fmt.Sprintf("%s %s", prop, generateResourceFromMap(representationType, representationMode, representationGroupInArray.group)))
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

// Temporary fix for identity resource export tests
func isIdentityOcid(ocid *string) bool {
	identityOcidPrefixes := []string{
		"authenticationPolicies/",
		"ocid1.compartment.",
		"ocid1.dynamicgroup.",
		"ocid1.group.",
		"ocid1.user.",
	}

	for _, prefix := range identityOcidPrefixes {
		if strings.HasPrefix(*ocid, prefix) {
			return true
		}
	}
	return false
}

// Temporarily skip export tests for following resources for now, because services
// don't return full information for running terraform plan to succeed
func skipExportForOcid(ocid *string) bool {
	skipExportOcidPrefixes := []string{
		"ocid1.saml2idp.",
		"ocid1.tagdefinition.",
		"ocid1.policy.",
	}

	for _, prefix := range skipExportOcidPrefixes {
		if strings.HasPrefix(*ocid, prefix) {
			return true
		}
	}
	return false
}

func testExportCompartment(OCID *string, compartmentId *string) error {
	var arg ExportCommandArgs
	dir, _ := os.Getwd()
	outputDir := fmt.Sprintf(dir + "/exportCompartment")
	if err := os.RemoveAll(outputDir); err != nil {
		log.Printf("unable to remove existing '%s' due to error '%v'", outputDir, err)
		return err
	}
	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		log.Printf("unable to create '%s' due to error '%v'", outputDir, err)
		return err
	}
	defer func() {
		if err := os.RemoveAll(outputDir); err != nil {
			log.Printf("unable to cleanup '%s' due to error '%v'", outputDir, err)
		}
	}()
	arg.CompartmentId = compartmentId
	arg.GenerateState = true
	arg.OutputDir = &outputDir
	arg.IDs = []string{*OCID}

	// Temporary fix for handling identity test requirements
	if isIdentityOcid(OCID) {
		arg.Services = []string{"identity"}
	} else if skipExportForOcid(OCID) {
		log.Printf("Skipping export test for OCID: %s", *OCID)
		return nil
	}

	if errExport := RunExportCommand(&arg); errExport != nil {
		return fmt.Errorf("[ERROR] RunExportCommand failed: %s", errExport)
	}
	meta := command.Meta{
		//Color:            true,
		//GlobalPluginDirs: globalPluginDirs(),
		//PluginOverrides:  &PluginOverrides,
		Ui: &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},

		//Services: services,

		RunningInAutomation: true,
		//PluginCacheDir:      config.PluginCacheDir,
		//OverrideDataDir:     dataDir,

		//ShutdownCh: makeShutdownCh(),
	}
	initCmd := command.InitCommand{Meta: meta}
	var initArgs []string
	if pluginDir := getEnvSettingWithBlankDefault("provider_bin_path"); pluginDir != "" {
		log.Printf("[INFO] plugin dir: '%s'", pluginDir)
		initArgs = append(initArgs, fmt.Sprintf("-plugin-dir=%v", pluginDir))
	}
	initArgs = append(initArgs, *arg.OutputDir)
	if errCode := initCmd.Run(initArgs); errCode != 0 {
		return nil
	}

	// Need to set the compartment OCID environment variable for plan step
	compartmentOcidVarName := "TF_VAR_compartment_ocid"
	storeCompartmentId := os.Getenv(compartmentOcidVarName)
	if err := os.Setenv(compartmentOcidVarName, *compartmentId); err != nil {
		return fmt.Errorf("could not set %s environment in export test", compartmentOcidVarName)
	}

	defer func() {
		if storeCompartmentId != "" {
			if err := os.Setenv(compartmentOcidVarName, storeCompartmentId); err != nil {
				log.Printf("[WARN] unable to restore %s to %s", compartmentOcidVarName, storeCompartmentId)
			}
		}
	}()

	planCmd := command.PlanCommand{Meta: meta}
	statefile := fmt.Sprintf(*arg.OutputDir + "/terraform.tfstate")
	if errCode := planCmd.Run([]string{"-detailed-exitcode", fmt.Sprintf("-state=%v", statefile), *arg.OutputDir}); errCode != 0 {
		if errCode == 1 {
			return fmt.Errorf("[ERROR] terraform plan command failed")
		} else {
			return fmt.Errorf("[ERROR] terraform plan command return non-empty diff")
		}
	}
	return nil
}

func checkJsonStringsEqual(expectedJsonString string, actualJsonString string) error {
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
func testCheckJsonResourceAttr(name, key, expectedJson string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actualJsonFromState, err := fromInstanceState(s, name, key)
		if err != nil {
			return err
		}

		if err := checkJsonStringsEqual(expectedJson, actualJsonFromState); err != nil {
			return fmt.Errorf("%s: Attribute '%s' %s", name, key, err)
		}
		return nil
	}
}
