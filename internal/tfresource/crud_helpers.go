// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"

	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v59/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v59/loadbalancer"
	oci_work_requests "github.com/oracle/oci-go-sdk/v59/workrequests"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	FifteenMinutes                = 15 * time.Minute
	TwentyMinutes                 = 20 * time.Minute
	ThirtyMinutes                 = 30 * time.Minute
	OneHour                       = 60 * time.Minute
	TwoHours                      = 120 * time.Minute
	TwoAndHalfHours               = 150 * time.Minute
	ThreeHours                    = 180 * time.Minute
	TwelveHours                   = 12 * time.Hour
	ZeroTime        time.Duration = 0

	DefaultTimeout = &schema.ResourceTimeout{
		Create: &TwentyMinutes,
		Update: &TwentyMinutes,
		Delete: &TwentyMinutes,
	}
	convertResFieldsToDSFields             = convertResourceFieldsToDatasourceFields
	jsonMarshal                            = json.Marshal
	waitForStateRefreshVar                 = WaitForStateRefresh
	WaitForWorkRequestVar                  = WaitForWorkRequest
	getWorkRequestErrorsVar                = getWorkRequestErrors
	waitForStateRefreshForHybridPollingVar = waitForStateRefreshForHybridPolling
	stateRefreshFuncVar                    = stateRefreshFunc
	HandleErrorVar                         = HandleError
	ShouldRetryVar                         = ShouldRetry
	reflectValueOf                         = reflect.ValueOf
)

const (
	FAILED    = "FAILED"
	SUCCEEDED = "SUCCEEDED"
)

const (
	OpcNextPageHeader = "Opc-Next-Page"
)

type BaseCrud struct {
	D     *schema.ResourceData
	Mutex *sync.Mutex
}

func (s *BaseCrud) VoidState() {
	s.D.SetId("")
}

// Default implementation, used in conjunction with State()
func (s *BaseCrud) setState(sync StatefulResource) error {
	// Pseudo code:
	//   currentState := sync.Res.State || sync.Resource.State || sync.WorkRequest.State
	//   s.D.Set("state", currentState)
	v := reflectValueOf(sync).Elem()
	for _, key := range []string{"Res", "Resource", "WorkRequest"} {
		// Yes, this "valid"ation is terrible
		if resourceReferenceValue := v.FieldByName(key); resourceReferenceValue.IsValid() {
			if resourceValue := resourceReferenceValue.Elem(); resourceValue.IsValid() {
				// In rare cases, the kind for "Res" is an interface (e.g. if the resource itself is
				// a polymorphic type, opposed to a field on the resource). Use Elem() to get the value
				// the interface contains, otherwise the ".FieldByName()" method will throw.
				if resourceValue.Kind() == reflect.Interface {
					resourceValue = resourceValue.Elem()
				}

				if stateValue := resourceValue.FieldByName("LifecycleState"); stateValue.IsValid() {
					currentState := stateValue.String()
					log.Printf("[DEBUG] BaseCrud.setState: state: %#v", currentState)
					return s.D.Set("state", currentState)
				} else if stateValue := resourceValue.FieldByName("State"); stateValue.IsValid() {
					currentState := stateValue.String()
					log.Printf("[DEBUG] BaseCrud.setState: state: %#v", currentState)
					return s.D.Set("state", currentState)
				}
			}
		}
	}

	return fmt.Errorf("Could not set resource state, sync did not have a valid .Res.State, .Resource.State, or .WorkRequest.State")
}

// Default implementation pulls state off of the schema
func (s *BaseCrud) State() string {
	str, ok := s.D.Get("state").(string)
	if ok {
		return str
	}
	return ""
}

type schemaResourceData interface {
	GetOkExists(string) (interface{}, bool)
	SetId(string)
	Timeout(string) time.Duration
	Partial(bool)
	HasChange(string) bool
	GetChange(string) (interface{}, interface{})
}

type workReqClient interface {
	GetWorkRequest(context.Context, oci_work_requests.GetWorkRequestRequest) (oci_work_requests.GetWorkRequestResponse, error)
	ListWorkRequestErrors(context.Context, oci_work_requests.ListWorkRequestErrorsRequest) (oci_work_requests.ListWorkRequestErrorsResponse, error)
}

func waitForStateRefreshForHybridPolling(workRequestClient workReqClient, workRequestIds *string, entityType string, action oci_work_requests.WorkRequestResourceActionTypeEnum,
	disableFoundRetries bool, sync StatefulResource, timeout time.Duration, operationName string, pending, target []string) error {
	// TODO: try to move this onto sync
	stateConf := &resource.StateChangeConf{
		Pending: pending,
		Target:  target,
		Refresh: stateRefreshFuncVar(sync),
		Timeout: timeout,
	}

	// Should not wait when in replay mode
	if httpreplay.ShouldRetryImmediately() {
		stateConf.PollInterval = 1
	}

	if _, e := stateConf.WaitForState(); e != nil {
		handleMissingResourceError(sync, &e)
		if _, ok := e.(*resource.UnexpectedStateError); ok {
			retryPolicy := GetRetryPolicy(disableFoundRetries, "work_request")
			retryPolicy.ShouldRetryOperation = workRequestShouldRetryFunc(timeout)
			e = getWorkRequestErrorsVar(workRequestClient, workRequestIds, retryPolicy, entityType, action)
			return e
		}

		if _, ok := e.(*resource.TimeoutError); ok {
			e = fmt.Errorf("%s, you may need to increase the Terraform Operation timeouts for your resource to continue polling for longer", e)
		}
		return e
	}

	if sync.State() == FAILED {
		return fmt.Errorf("Resource %s failed, state FAILED", operationName)
	}

	return nil
}

func ResourceRefreshForHybridPolling(workRequestClient workReqClient, workRequestIds *string, entityType string, action oci_work_requests.WorkRequestResourceActionTypeEnum,
	disableFoundRetries bool, d schemaResourceData, sync ResourceCreator) error {

	// ID is required for state refresh
	d.SetId(sync.ID())

	if stateful, ok := sync.(StatefullyCreatedResource); ok {
		if e := waitForStateRefreshForHybridPollingVar(workRequestClient, workRequestIds, entityType, action, disableFoundRetries, stateful, d.Timeout(schema.TimeoutCreate), "creation", stateful.CreatedPending(), stateful.CreatedTarget()); e != nil {
			if stateful.State() == FAILED {
				// Remove resource from state if asynchronous work request has failed so that it is recreated on next apply
				// TODO: automatic retry on WorkRequestFailed
				sync.VoidState()
			}

			//We need to SetData() here because if there is an error or timeout in the wait for state after the Create() was successful we want to store the resource in the statefile to avoid dangling resources
			if setDataErr := sync.SetData(); setDataErr != nil {
				log.Printf("[ERROR] error setting data after WaitForStateRefresh() error: %v", setDataErr)
			}

			return e
		}
	}

	d.SetId(sync.ID())
	if e := sync.SetData(); e != nil {
		return e
	}

	if ew, waitOK := sync.(ExtraWaitPostCreateDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostCreateDelete())
	}

	return nil
}

func CreateResourceUsingHybridPolling(sync ResourceCreator) error {
	if e := sync.Create(); e != nil {
		return HandleErrorVar(sync, e)
	}

	return nil
}

func CreateResource(d schemaResourceData, sync ResourceCreator) error {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	if e := sync.Create(); e != nil {
		return HandleError(sync, e)
	}

	// ID is required for state refresh
	d.SetId(sync.ID())

	if stateful, ok := sync.(StatefullyCreatedResource); ok {
		if e := waitForStateRefreshVar(stateful, d.Timeout(schema.TimeoutCreate), "creation", stateful.CreatedPending(), stateful.CreatedTarget()); e != nil {
			if stateful.State() == FAILED {
				// Remove resource from state if asynchronous work request has failed so that it is recreated on next apply
				// TODO: automatic retry on WorkRequestFailed
				sync.VoidState()
			}

			//We need to SetData() here because if there is an error or timeout in the wait for state after the Create() was successful we want to store the resource in the statefile to avoid dangling resources
			if setDataErr := sync.SetData(); setDataErr != nil {
				log.Printf("[ERROR] error setting data after WaitForStateRefresh() error: %v", setDataErr)
			}

			return e
		}
	}

	d.SetId(sync.ID())
	if e := sync.SetData(); e != nil {
		return e
	}

	if ew, waitOK := sync.(ExtraWaitPostCreateDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostCreateDelete())
	}
	return nil
}

func ReadResource(sync ResourceReader) error {
	if e := sync.Get(); e != nil {
		log.Printf("ERROR IN GET: %v\n", e.Error())
		handleMissingResourceError(sync, &e)
		return HandleError(sync, e)
	}

	if e := sync.SetData(); e != nil {
		return e
	}

	// Remove resource from state if it has been terminated so that it is recreated on next apply
	if dr, ok := sync.(StatefullyDeletedResource); ok {
		for _, target := range dr.DeletedTarget() {
			if dr.State() == target && dr.State() != string(oci_load_balancer.WorkRequestLifecycleStateSucceeded) {
				dr.VoidState()
				return nil
			}
		}
	}

	return nil
}

func UpdateResource(d schemaResourceData, sync ResourceUpdater) error {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	d.Partial(true)
	if e := sync.Update(); e != nil {

		return HandleError(sync, e)
	}
	d.Partial(false)

	if stateful, ok := sync.(StatefullyUpdatedResource); ok {
		if e := waitForStateRefreshVar(stateful, d.Timeout(schema.TimeoutUpdate), "update", stateful.UpdatedPending(), stateful.UpdatedTarget()); e != nil {

			return e
		}
	}

	if e := sync.SetData(); e != nil {
		return e
	}

	return nil
}

// DeleteResource requests a Delete(). If the resource deletes
// statefully (not immediately), poll State to ensure:
// () -> Pending -> Deleted.
// Finally, sets the ResourceData state to empty.
func DeleteResource(d schemaResourceData, sync ResourceDeleter) error {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	if e := sync.Delete(); e != nil {
		handleMissingResourceError(sync, &e)
		return HandleError(sync, e)
	}

	if stateful, ok := sync.(StatefullyDeletedResource); ok {
		if e := waitForStateRefreshVar(stateful, d.Timeout(schema.TimeoutDelete), "deletion", stateful.DeletedPending(), stateful.DeletedTarget()); e != nil {
			handleMissingResourceError(sync, &e)
			return e
		}
	}

	if ew, waitOK := sync.(ExtraWaitPostCreateDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostCreateDelete())
	}

	if ew, waitOK := sync.(ExtraWaitPostDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostDelete())
	}

	sync.VoidState()

	return nil
}

func stateRefreshFunc(sync StatefulResource) resource.StateRefreshFunc {
	return func() (res interface{}, s string, e error) {
		if e = sync.Get(); e != nil {
			return nil, "", e
		}
		// We don't set all the state here, because not found errors are handled elsewhere.
		// But we do need the new state for the default State() function
		if e = sync.setState(sync); e != nil {
			return nil, "", e
		}
		return sync, sync.State(), e
	}
}

// Helper function to wait for Update to reach terminal state before doing another Update
// Useful in situations where more than one Update is needed and prior Update needs to complete
func WaitForUpdatedState(d schemaResourceData, sync ResourceUpdater) error {
	if stateful, ok := sync.(StatefullyUpdatedResource); ok {
		if e := waitForStateRefreshVar(stateful, d.Timeout(schema.TimeoutUpdate), "update", stateful.UpdatedPending(), stateful.UpdatedTarget()); e != nil {
			return e
		}
	}

	return nil
}

// Helper function to wait for Create to reach terminal state before doing another operation
// Useful in situations where another operation is done right after Create
func WaitForCreatedState(d schemaResourceData, sync ResourceCreator) error {
	d.SetId(sync.ID())
	if stateful, ok := sync.(StatefullyCreatedResource); ok {
		if e := waitForStateRefreshVar(stateful, d.Timeout(schema.TimeoutCreate), "creation", stateful.CreatedPending(), stateful.CreatedTarget()); e != nil {
			return e
		}
	}

	return nil
}

// WaitForStateRefresh takes a StatefulResource, a timeout duration, a list of states to treat as Pending, and a list of states to treat as Target. It uses those to wrap resource.StateChangeConf.WaitForState(). If the resource returns a missing status, it will not be treated as an error.
//
// sync.D.Id must be set.
// It does not set state from that refreshed state.
func WaitForStateRefresh(sync StatefulResource, timeout time.Duration, operationName string, pending, target []string) error {
	// TODO: try to move this onto sync
	stateConf := &resource.StateChangeConf{
		Pending: pending,
		Target:  target,
		Refresh: stateRefreshFuncVar(sync),
		Timeout: timeout,
	}

	// Should not wait when in replay mode
	if httpreplay.ShouldRetryImmediately() {
		stateConf.PollInterval = 1
	}

	if _, e := stateConf.WaitForState(); e != nil {
		handleMissingResourceError(sync, &e)
		if _, ok := e.(*resource.UnexpectedStateError); ok {
			if len(target) > 0 {
				e = fmt.Errorf("During %s, Terraform expected the resource to reach state(s): %s, but the service reported unexpected state: %s.", operationName, strings.Join(target, ","), sync.State())
			} else {
				e = fmt.Errorf("During %s, service reported unexpected state: %s.", operationName, sync.State())
			}
			return e
		}

		if _, ok := e.(*resource.TimeoutError); ok {
			e = fmt.Errorf("%s, you may need to increase the Terraform Operation timeouts for your resource to continue polling for longer", e)
		}
		return e
	}

	if sync.State() == FAILED {
		return fmt.Errorf("Resource %s failed, state FAILED", operationName)
	}

	return nil
}

func FilterMissingResourceError(sync ResourceVoider, err *error) {
	if err != nil && strings.Contains((*err).Error(), "does not exist") {
		log.Println("[DEBUG] Object does not exist, voiding resource and nullifying error")
		sync.VoidState()
		*err = nil
	}
}

// no unit test
// In the Exadata case the service return the hostname provided by the service with a suffix
func DbSystemHostnameDiffSuppress(key string, old string, new string, d *schema.ResourceData) bool {
	return EqualIgnoreCaseSuppressDiff(key, old, new, d) || NewIsPrefixOfOldDiffSuppress(key, old, new, d)
}

func NewIsPrefixOfOldDiffSuppress(key string, old string, new string, d *schema.ResourceData) bool {
	return strings.HasPrefix(strings.ToLower(old), strings.ToLower(new))
}

func DbVersionDiffSuppress(key string, old string, new string, d *schema.ResourceData) bool {
	if old == "" || new == "" {
		return false
	}
	if new == "18.0.0.0" || new == "19.0.0.0" {
		oldVersion := strings.Split(old, ".")
		newVersion := strings.Split(new, ".")
		oldVersionNumber, err := strconv.Atoi(oldVersion[0])
		if err != nil {
			return false
		}
		newVersionNumber, err := strconv.Atoi(newVersion[0])
		if err != nil {
			return false
		}

		return oldVersionNumber == newVersionNumber
	}
	return strings.HasPrefix(strings.ToLower(old), strings.ToLower(new))
}

func AdDiffSuppress(key string, old string, new string, d *schema.ResourceData) bool {
	const float64EqualityThreshold = 1e-6
	oldf, err := strconv.ParseFloat(old, 64)
	if err != nil {
		return false
	}
	newf, err := strconv.ParseFloat(new, 64)
	if err != nil {
		return false
	}

	return math.Abs(oldf-newf) <= float64EqualityThreshold
}

func GiVersionDiffSuppress(key string, old string, new string, d *schema.ResourceData) bool {
	if old == "" || new == "" {
		return false
	}
	oldVersion := strings.Split(old, ".")
	newVersion := strings.Split(new, ".")

	if oldVersion[0] == newVersion[0] {
		return true
	}
	return false
}

func MySqlVersionDiffSuppress(key string, old string, new string, d *schema.ResourceData) bool {
	if old == "" || new == "" {
		return false
	}
	oldVersion := strings.Split(old, ".")
	newVersion := strings.Split(new, ".")

	oldMajorVersion := oldVersion[0] + "." + oldVersion[1]
	newMajorVersion := newVersion[0] + "." + newVersion[1]

	if oldMajorVersion == newMajorVersion {
		return true
	}

	return false
}

func LoadBalancersSuppressDiff(key string, old string, new string, d *schema.ResourceData) bool {
	return loadBalancersSuppressDiff(d)
}
func loadBalancersSuppressDiff(d schemaResourceData) bool {
	if !d.HasChange("load_balancers") {
		return true
	}
	oldRaw, newRaw := d.GetChange("load_balancers")
	interfaces := oldRaw.([]interface{})
	oldBalancers := make(map[string]bool, len(interfaces))
	for _, item := range interfaces {
		oldBalancers[item.(map[string]interface{})["load_balancer_id"].(string)] = true
	}

	interfaces = newRaw.([]interface{})
	if len(interfaces) != len(oldBalancers) {
		return false
	}

	for _, item := range interfaces {
		converted := item.(map[string]interface{})["load_balancer_id"].(string)
		if _, ok := oldBalancers[converted]; !ok {
			return false
		}
	}
	return true
}

func EqualIgnoreCaseSuppressDiff(key string, old string, new string, d *schema.ResourceData) bool {
	return strings.EqualFold(old, new)
}

func FieldDeprecatedAndAvoidReferences(deprecatedFieldName string) string {
	return fmt.Sprintf("The '%s' field has been deprecated and may be removed in a future version. Do not use this field.", deprecatedFieldName)
}

func FieldDeprecated(deprecatedFieldName string) string {
	return fmt.Sprintf("The '%s' field has been deprecated. It is no longer supported.", deprecatedFieldName)
}

func FieldDeprecatedForAnother(deprecatedFieldName string, newFieldName string) string {
	return fmt.Sprintf("The '%s' field has been deprecated. Please use '%s' instead.", deprecatedFieldName, newFieldName)
}

func FieldDeprecatedButSupportedThroughAnotherResource(deprecatedFieldName string, newResourceName string) string {
	return fmt.Sprintf("The '%s' field has been deprecated. Please use the '%s' resource instead.", deprecatedFieldName, newResourceName)
}

func FieldDeprecatedButSupportedThroughAnotherDataSource(deprecatedFieldName string, newDataSourceName string) string {
	return fmt.Sprintf("The '%s' field has been deprecated. Please use the '%s' data source instead.", deprecatedFieldName, newDataSourceName)
}

func FieldDeprecatedAndOverridenByAnother(deprecatedFieldName string, newFieldName string) string {
	return fmt.Sprintf("%s If both fields are specified, then '%s' will be used.",
		FieldDeprecatedForAnother(deprecatedFieldName, newFieldName), newFieldName)
}

func ResourceDeprecatedForAnother(deprecatedResourceName string, newResourceName string) string {
	return fmt.Sprintf("The '%s' resource has been deprecated. Please use '%s' instead.", deprecatedResourceName, newResourceName)
}

func ResourceNotFoundErrorMessage(resourceName string, reason string) error {
	// Use this function to generate an error message for any resource that is not found.  The message is specially
	// formatted so that it is detected by the handleMissingResourceError function correctly.  Do not change the message format.
	return fmt.Errorf("%s not found. %s \n", resourceName, reason)
}

// GenerateDataSourceID generates an ID for the data source based on the current time stamp.
func GenerateDataSourceID() string {
	// Important, if you don't have an ID, make one up for your datasource
	// or things will end in tears.

	// Consider prefixing with resource name or useful identifier beyond just a timestamp.
	return time.Now().UTC().String()
}

func GenerateDataSourceHashID(idPrefix string, resourceSchema *schema.Resource, resourceData schemaResourceData) string {
	// Important, if you don't have an ID, make one up for your datasource
	// or things will end in tears.

	if resourceSchema == nil || resourceData == nil {
		return ""
	}

	var buf bytes.Buffer
	// sort keys of the map
	keys := make([]string, 0, len(resourceSchema.Schema))
	for key := range resourceSchema.Schema {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		// parse schema field is user input
		value := resourceSchema.Schema[key]
		if !(value.Required || value.Optional) {
			continue
		}

		// Ignoring TypeList, TypeSet and TypeMap
		if value.Type == schema.TypeList || value.Type == schema.TypeSet || value.Type == schema.TypeMap {
			continue
		}

		if element, ok := resourceData.GetOkExists(key); ok {
			buf.WriteString(fmt.Sprintf("%v-", element))
		}
	}
	return fmt.Sprintf("%s%d", idPrefix, hashcode.String(buf.String()))
}

// stringsToSet encodes an []string into a
// *schema.Set in the appropriate structure for the schema
func StringsToSet(ss []string) *schema.Set {
	st := &schema.Set{F: schema.HashString}
	for _, s := range ss {
		st.Add(s)
	}
	return st
}

// SetToString encodes an *schema.Set into an []string honoring the structure for the schema
func SetToStrings(volumeIdsSet *schema.Set) []string {
	interfaces := volumeIdsSet.List()
	tmp := make([]string, len(interfaces))
	for i, toBeConverted := range interfaces {
		tmp[i] = toBeConverted.(string)
	}
	return tmp
}

// NormalizeBoolString parses a string value into a bool value, and if successful, formats it back
// into a string & throws an error otherwise. This allows for normalizing the different formats of
// valid bool strings (e.g. "1", "false", "TRUE", "F", etc.) to a uniform string representation of
// a boolean value ("true" & "false").
func NormalizeBoolString(v string) (string, error) {
	boolVal, err := strconv.ParseBool(v)
	if err != nil {
		return "", err
	}
	return strconv.FormatBool(boolVal), nil
}

// WaitForResourceCondition polls on a resource, waiting for it to reach a specified condition. This is done with exponential
// backoff. The terminating condition is specified as a boolean function; and this will return a timeout error if the
// specified condition isn't reached within the specified timeout period.
func WaitForResourceCondition(s ResourceFetcher, resourceChangedFunc func() bool, timeout time.Duration) error {
	backoffTime := time.Second
	startTime := time.Now()
	endTime := startTime.Add(timeout)
	lastAttempt := false
	for {
		if err := s.Get(); err != nil {
			return err
		}

		if resourceChangedFunc() {
			break
		}

		if lastAttempt || time.Now().After(endTime) {
			return fmt.Errorf("Timed out waiting for configuration to reach specified condition.")
		}

		backoffTime = backoffTime * 2

		// If next attempt occurs after timeout, then retry earlier
		nextAttemptTime := time.Now().Add(backoffTime)
		if nextAttemptTime.After(endTime) {
			backoffTime = endTime.Sub(time.Now())
			lastAttempt = true
		}

		if httpreplay.ShouldRetryImmediately() {
			backoffTime = 10 * time.Millisecond
		}

		time.Sleep(backoffTime)
	}

	return nil
}

// Get the schema for a nested DataSourceSchema generated from the ResourceSchema
func GetDataSourceItemSchema(resourceSchema *schema.Resource) *schema.Resource {
	if _, idExists := resourceSchema.Schema["id"]; !idExists {
		resourceSchema.Schema["id"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
	}

	// Ensure Create/Read are not set for nested sub-resource schemas. Otherwise, terraform will validate them
	// as though they were resources.
	resourceSchema.Create = nil
	resourceSchema.Read = nil

	return convertResFieldsToDSFields(resourceSchema)
}

// Get the Singular DataSource Schema from Resource Schema with additional fields and Read Function
func GetSingularDataSourceItemSchema(resourceSchema *schema.Resource, addFieldMap map[string]*schema.Schema, readFunc schema.ReadFunc) *schema.Resource {
	if _, idExists := resourceSchema.Schema["id"]; !idExists {
		resourceSchema.Schema["id"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
	}

	// Ensure Create,Read, Update and Delete are not set for data source schemas. Otherwise, terraform will validate them
	// as though they were resources.
	resourceSchema.Create = nil
	resourceSchema.Update = nil
	resourceSchema.Delete = nil
	resourceSchema.Read = readFunc
	resourceSchema.Importer = nil
	resourceSchema.Timeouts = nil
	resourceSchema.CustomizeDiff = nil

	var dataSourceSchema *schema.Resource = convertResFieldsToDSFields(resourceSchema)

	for key, value := range addFieldMap {
		dataSourceSchema.Schema[key] = value
	}

	return dataSourceSchema
}

// This is mainly used to ensure that fields of a datasource item are compliant with Terraform schema validation
// All datasource return items should have computed-only fields; and not require Diff, Validation, or Default settings.
func convertResourceFieldsToDatasourceFields(resourceSchema *schema.Resource) *schema.Resource {
	for _, fieldSchema := range resourceSchema.Schema {
		fieldSchema.Computed = true
		fieldSchema.Required = false
		fieldSchema.Optional = false
		fieldSchema.DiffSuppressFunc = nil
		fieldSchema.ValidateFunc = nil
		fieldSchema.ConflictsWith = nil
		fieldSchema.Default = nil
		fieldSchema.DefaultFunc = nil
		if fieldSchema.Type == schema.TypeSet {
			fieldSchema.Type = schema.TypeList
			fieldSchema.Set = nil
		}

		if fieldSchema.Elem != nil {
			if resource, ok := fieldSchema.Elem.(*schema.Resource); ok {
				fieldSchema.Elem = convertResourceFieldsToDatasourceFields(resource)
			}
		}
	}

	return resourceSchema
}

func GetRetryPolicyWithAdditionalRetryCondition(timeout time.Duration, retryConditionFunction func(oci_common.OCIOperationResponse) bool, service string) *oci_common.RetryPolicy {
	startTime := time.Now()
	return &oci_common.RetryPolicy{
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			if ShouldRetryVar(response, false, service, startTime) {
				return true
			}
			if retryConditionFunction(response) {
				timeWaited := GetElapsedRetryDuration(startTime)
				return timeWaited < timeout
			}
			return false
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return GetRetryBackoffDuration(response, false, service, startTime)
		},
		MaximumNumberAttempts: 0,
	}
}

func elaspedInMillisecond(start time.Time) int64 {
	return time.Since(start).Nanoseconds() / int64(time.Millisecond)
}

func WaitForWorkRequestWithErrorHandling(workRequestClient workReqClient, workRequestIds *string, entityType string, action oci_work_requests.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool) (*string, error) {
	var identifier *string
	workRequestIdsSet := map[string]bool{}

	for _, wId := range strings.Split(strings.TrimSpace(*workRequestIds), ",") {
		if wId != "" {
			workRequestIdsSet[strings.TrimSpace(wId)] = true
		}
	}

	for wId := range workRequestIdsSet {
		id, err := WaitForWorkRequestVar(workRequestClient, &wId, entityType, action, timeout, disableFoundRetries, true)
		if err != nil {
			return id, err
		}
		identifier = id
	}
	return identifier, nil

}

func WaitForWorkRequest(workRequestClient workReqClient, workRequestId *string, entityType string, action oci_work_requests.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, expectIdentifier bool) (*string, error) {
	retryPolicy := GetRetryPolicy(disableFoundRetries, "work_request")
	retryPolicy.ShouldRetryOperation = workRequestShouldRetryFunc(timeout)

	response := oci_work_requests.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_work_requests.WorkRequestStatusInProgress),
			string(oci_work_requests.WorkRequestStatusAccepted),
			string(oci_work_requests.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_work_requests.WorkRequestStatusSucceeded),
			string(oci_work_requests.WorkRequestStatusFailed),
			string(oci_work_requests.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = workRequestClient.GetWorkRequest(context.Background(),
				oci_work_requests.GetWorkRequestRequest{
					WorkRequestId: workRequestId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}

	var identifier *string

	if _, e := stateConf.WaitForState(); e != nil {
		for _, res := range response.Resources {
			if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
				if res.Identifier != nil {
					identifier = res.Identifier
					break
				}
			}
		}

		return identifier, e
	}

	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	if expectIdentifier && identifier == nil {
		if response.Status == oci_work_requests.WorkRequestStatusSucceeded {
			return nil, fmt.Errorf("work request succeeded but no identifier was found, workId: %s, entity: %s, action: %s",
				*workRequestId, entityType, action)
		}
		return nil, getWorkRequestErrorsVar(workRequestClient, workRequestId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func workRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if ShouldRetryVar(response, false, "work_request", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_work_requests.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func getWorkRequestErrors(workRequestClient workReqClient, workRequestId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_work_requests.WorkRequestResourceActionTypeEnum) error {
	response, err := workRequestClient.ListWorkRequestErrors(context.Background(), oci_work_requests.ListWorkRequestErrorsRequest{
		WorkRequestId: workRequestId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workRequestId, entityType, action, errorMessage)

	return workRequestErr
}

// Helper to marshal JSON objects from service into strings that can be stored in state.
// This limitation exists because Terraform doesn't support maps of nested objects and so we use JSON strings representation
// as a workaround.
func GenericMapToJsonMap(genericMap map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	for key, value := range genericMap {
		switch v := value.(type) {
		case string:
			result[key] = v
		default:
			bytes, err := jsonMarshal(v)
			if err != nil {
				continue
			}
			result[key] = string(bytes)
		}
	}

	return result
}

func GetTimeoutDuration(timeout string) *time.Duration {
	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		// Return the OCI Provider's default timeout if there is an error
		return &TwentyMinutes
	}
	return &timeoutDuration
}

func ConvertObjectToJsonString(object interface{}) (string, error) {
	bytes, err := jsonMarshal(object)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
