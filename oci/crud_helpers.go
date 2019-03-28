// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"sync"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

var (
	FifteenMinutes                = 15 * time.Minute
	TwentyMinutes                 = 20 * time.Minute
	OneHour                       = 60 * time.Minute
	TwoHours                      = 120 * time.Minute
	TwoAndHalfHours               = 150 * time.Minute
	ZeroTime        time.Duration = 0

	DefaultTimeout = &schema.ResourceTimeout{
		Create: &FifteenMinutes,
		Update: &FifteenMinutes,
		Delete: &FifteenMinutes,
	}
)

const (
	FAILED = "FAILED"
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
	v := reflect.ValueOf(sync).Elem()
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

func handleMissingResourceError(sync ResourceVoider, err *error) {

	if err != nil {
		if strings.Contains((*err).Error(), "does not exist") ||
			strings.Contains((*err).Error(), " not present in ") ||
			strings.Contains((*err).Error(), "not found") ||
			(strings.Contains((*err).Error(), "Load balancer") && strings.Contains((*err).Error(), " has no ")) {

			log.Println("[DEBUG] Object does not exist, voiding resource and nullifying error")
			sync.VoidState()
			*err = nil
		}
	}
}

func LoadBalancerResourceID(res interface{}, workReq *oci_load_balancer.WorkRequest) (id *string, workReqSucceeded bool) {
	v := reflect.ValueOf(res).Elem()
	if v.IsValid() {
		// This is super fugly. It's this way because the LB API has no convention for ID formats.

		// Load balancer
		id := v.FieldByName("Id")
		if id.IsValid() && !id.IsNil() {
			s := id.Elem().String()
			return &s, false
		}
		// backendset, listener
		name := v.FieldByName("Name")
		if name.IsValid() && !name.IsNil() {
			s := name.Elem().String()
			return &s, false
		}
		// certificate
		certName := v.FieldByName("CertificateName")
		if certName.IsValid() && !certName.IsNil() {
			s := certName.Elem().String()
			return &s, false
		}
		// backend TODO The following can probably be removed because the Backend object has a Name parameter)
		ip := v.FieldByName("IpAddress")
		port := v.FieldByName("Port")
		if ip.IsValid() && !ip.IsNil() && port.IsValid() && !port.IsNil() {
			s := ip.Elem().String() + ":" + strconv.Itoa(int(int(port.Elem().Int())))
			return &s, false
		}
	}
	if workReq != nil {
		if workReq.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return nil, true
		} else {
			return workReq.Id, false
		}
	}
	return nil, false
}

func LoadBalancerResourceGet(client *oci_load_balancer.LoadBalancerClient, d *schema.ResourceData, wr *oci_load_balancer.WorkRequest, retryPolicy *oci_common.RetryPolicy) (id string, stillWorking bool, err error) {
	// NOTE: if the id is for a work request, refresh its state and loadBalancerID.
	if wr != nil && wr.Id != nil {
		getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
		getWorkRequestRequest.WorkRequestId = wr.Id
		getWorkRequestRequest.RequestMetadata.RetryPolicy = retryPolicy
		updatedWorkRes, err := client.GetWorkRequest(context.Background(), getWorkRequestRequest)
		if err != nil {
			return "", false, err
		}
		if wr != nil {
			*wr = updatedWorkRes.WorkRequest
			d.Set("state", wr.LifecycleState)
			if wr.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
				return "", false, nil
			}
			if wr.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateFailed {
				return "", false, fmt.Errorf("WorkRequest FAILED: %+v", wr.ErrorDetails)
			}
		}
		return "", true, nil
	}
	return id, false, nil
}

func LoadBalancerWaitForWorkRequest(client *oci_load_balancer.LoadBalancerClient, d *schema.ResourceData, wr *oci_load_balancer.WorkRequest, retryPolicy *oci_common.RetryPolicy) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
			string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
		},
		Target: []string{
			string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
			string(oci_load_balancer.WorkRequestLifecycleStateFailed),
		},
		Refresh: func() (interface{}, string, error) {
			getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
			getWorkRequestRequest.WorkRequestId = wr.Id
			getWorkRequestRequest.RequestMetadata.RetryPolicy = retryPolicy
			workRequestResponse, err := client.GetWorkRequest(context.Background(), getWorkRequestRequest)
			wr = &workRequestResponse.WorkRequest
			return wr, string(wr.LifecycleState), err
		},
		Timeout: d.Timeout(schema.TimeoutCreate),
	}

	if _, e := stateConf.WaitForState(); e != nil {
		return e
	}

	if wr.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateFailed {
		return fmt.Errorf("WorkRequest FAILED: %+v", wr.ErrorDetails)
	}
	return nil
}

func CreateDBSystemResource(d *schema.ResourceData, sync ResourceCreator) error {
	if e := sync.Create(); e != nil {
		return e
	}

	// ID is required for state refresh
	d.SetId(sync.ID())

	var timeout time.Duration
	shape := d.Get("shape")
	timeout = d.Timeout(schema.TimeoutCreate)
	if timeout == 0 {
		if strings.HasPrefix(shape.(string), "Exadata") {
			timeout = time.Duration(12) * time.Hour
		} else {
			timeout = time.Duration(2) * time.Hour
		}
	}
	if stateful, ok := sync.(StatefullyCreatedResource); ok {
		if e := waitForStateRefresh(stateful, timeout, "creation", stateful.CreatedPending(), stateful.CreatedTarget()); e != nil {
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

func CreateResource(d *schema.ResourceData, sync ResourceCreator) error {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	if e := sync.Create(); e != nil {
		return e
	}

	// ID is required for state refresh
	d.SetId(sync.ID())

	if stateful, ok := sync.(StatefullyCreatedResource); ok {
		if e := waitForStateRefresh(stateful, d.Timeout(schema.TimeoutCreate), "creation", stateful.CreatedPending(), stateful.CreatedTarget()); e != nil {
			if stateful.State() == FAILED {
				// Remove resource from state if asynchronous work request has failed so that it is recreated on next apply
				// TODO: automatic retry on WorkRequestFailed
				sync.VoidState()
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
		return e
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

func UpdateResource(d *schema.ResourceData, sync ResourceUpdater) error {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	d.Partial(true)
	if e := sync.Update(); e != nil {
		return e
	}
	d.Partial(false)

	if stateful, ok := sync.(StatefullyUpdatedResource); ok {
		if e := waitForStateRefresh(stateful, d.Timeout(schema.TimeoutUpdate), "update", stateful.UpdatedPending(), stateful.UpdatedTarget()); e != nil {
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
func DeleteResource(d *schema.ResourceData, sync ResourceDeleter) error {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	if e := sync.Delete(); e != nil {
		handleMissingResourceError(sync, &e)
		return e
	}

	if stateful, ok := sync.(StatefullyDeletedResource); ok {
		if e := waitForStateRefresh(stateful, d.Timeout(schema.TimeoutDelete), "deletion", stateful.DeletedPending(), stateful.DeletedTarget()); e != nil {
			handleMissingResourceError(sync, &e)
			return e
		}
	}

	if ew, waitOK := sync.(ExtraWaitPostCreateDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostCreateDelete())
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

// Helper function to wait for update to reach terminal state before doing another update
// Useful in situations where more than one update is needed and prior update needs to complete
func waitForUpdatedState(d *schema.ResourceData, sync ResourceUpdater) error {
	if stateful, ok := sync.(StatefullyUpdatedResource); ok {
		if e := waitForStateRefresh(stateful, d.Timeout(schema.TimeoutUpdate), "update", stateful.UpdatedPending(), stateful.UpdatedTarget()); e != nil {
			return e
		}
	}

	return nil
}

// waitForStateRefresh takes a StatefulResource, a timeout duration, a list of states to treat as Pending, and a list of states to treat as Target. It uses those to wrap resource.StateChangeConf.WaitForState(). If the resource returns a missing status, it will not be treated as an error.
//
// sync.D.Id must be set.
// It does not set state from that refreshed state.
func waitForStateRefresh(sync StatefulResource, timeout time.Duration, operationName string, pending, target []string) error {
	// TODO: try to move this onto sync
	stateConf := &resource.StateChangeConf{
		Pending: pending,
		Target:  target,
		Refresh: stateRefreshFunc(sync),
		Timeout: timeout,
	}

	if _, e := stateConf.WaitForState(); e != nil {
		handleMissingResourceError(sync, &e)
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

// GenerateDataSourceID generates an ID for the data source based on the current time stamp.
func GenerateDataSourceID() string {
	// Important, if you don't have an ID, make one up for your datasource
	// or things will end in tears.

	// Consider prefixing with resource name or useful identifier beyond just a timestamp.
	return time.Now().UTC().String()
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

	return convertResourceFieldsToDatasourceFields(resourceSchema)
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

	var dataSourceSchema *schema.Resource = convertResourceFieldsToDatasourceFields(resourceSchema)

	for key, value := range addFieldMap {
		if _, fieldExists := resourceSchema.Schema[key]; !fieldExists {
			dataSourceSchema.Schema[key] = value
		}
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
