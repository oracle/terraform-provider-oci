// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package crud

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
	FiveMinutes time.Duration = 5 * time.Minute
	TwoHours    time.Duration = 120 * time.Minute
	ZeroTime    time.Duration = 0

	DefaultTimeout = &schema.ResourceTimeout{
		Create: &FiveMinutes,
		Update: &FiveMinutes,
		Delete: &FiveMinutes,
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
					log.Printf("[DEBUG] crud.BaseCrud.setState: state: %#v", currentState)
					return s.D.Set("state", currentState)
				} else if stateValue := resourceValue.FieldByName("State"); stateValue.IsValid() {
					currentState := stateValue.String()
					log.Printf("[DEBUG] crud.BaseCrud.setState: state: %#v", currentState)
					return s.D.Set("state", currentState)
				}
			}
		}
	}

	panic("Could not set resource state, sync did not have a valid .Res.State, .Resource.State, or .WorkRequest.State")
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
	id = d.Id()
	// NOTE: if the id is for a work request, refresh its state and loadBalancerID.
	if strings.HasPrefix(id, "ocid1.loadbalancerworkrequest.") {
		getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
		getWorkRequestRequest.WorkRequestId = &id
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
	var e error
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

	if _, e = stateConf.WaitForState(); e != nil {
		return e
	}
	if wr.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateFailed {
		return fmt.Errorf("WorkRequest FAILED: %+v", wr.ErrorDetails)
	}
	return nil
}

func CreateDBSystemResource(d *schema.ResourceData, sync ResourceCreator) (e error) {
	if e = sync.Create(); e != nil {
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
		e = waitForStateRefresh(stateful, timeout, "creation", stateful.CreatedPending(), stateful.CreatedTarget())
	}

	d.SetId(sync.ID())
	sync.SetData()

	if ew, waitOK := sync.(ExtraWaitPostCreateDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostCreateDelete())
	}

	return
}

func CreateResource(d *schema.ResourceData, sync ResourceCreator) (e error) {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	if e = sync.Create(); e != nil {
		return e
	}

	// ID is required for state refresh
	d.SetId(sync.ID())

	if stateful, ok := sync.(StatefullyCreatedResource); ok {
		e = waitForStateRefresh(stateful, d.Timeout(schema.TimeoutCreate), "creation", stateful.CreatedPending(), stateful.CreatedTarget())
		if stateful.State() == string(oci_load_balancer.WorkRequestLifecycleStateFailed) {
			// Remove resource from state if asynchronous work request has failed so that it is recreated on next apply
			// TODO: automatic retry on WorkRequestFailed
			sync.VoidState()
			return
		}
	}

	d.SetId(sync.ID())
	sync.SetData()

	if ew, waitOK := sync.(ExtraWaitPostCreateDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostCreateDelete())
	}

	return
}

func ReadResource(sync ResourceReader) (e error) {
	if e = sync.Get(); e != nil {
		log.Printf("ERROR IN GET: %v\n", e.Error())
		handleMissingResourceError(sync, &e)
		return
	}

	sync.SetData()

	// Remove resource from state if it has been terminated so that it is recreated on next apply
	if dr, ok := sync.(StatefullyDeletedResource); ok {
		for _, target := range dr.DeletedTarget() {
			if dr.State() == target && dr.State() != string(oci_load_balancer.WorkRequestLifecycleStateSucceeded) {
				dr.VoidState()
				return
			}
		}
	}

	return
}

func UpdateResource(d *schema.ResourceData, sync ResourceUpdater) (e error) {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	d.Partial(true)
	if e = sync.Update(); e != nil {
		return
	}
	d.Partial(false)

	if stateful, ok := sync.(StatefullyUpdatedResource); ok {
		e = waitForStateRefresh(stateful, d.Timeout(schema.TimeoutUpdate), "update", stateful.UpdatedPending(), stateful.UpdatedTarget())
	}

	sync.SetData()

	return
}

// DeleteResource requests a Delete(). If the resource deletes
// statefully (not immediately), poll State to ensure:
// () -> Pending -> Deleted.
// Finally, sets the ResourceData state to empty.
func DeleteResource(d *schema.ResourceData, sync ResourceDeleter) (e error) {
	if synchronizedResource, ok := sync.(SynchronizedResource); ok {
		if mutex := synchronizedResource.GetMutex(); mutex != nil {
			mutex.Lock()
			defer mutex.Unlock()
		}
	}

	if e = sync.Delete(); e != nil {
		handleMissingResourceError(sync, &e)
		return
	}

	//d.SetId(sync.ID())
	if stateful, ok := sync.(StatefullyDeletedResource); ok {
		e = waitForStateRefresh(stateful, d.Timeout(schema.TimeoutDelete), "deletion", stateful.DeletedPending(), stateful.DeletedTarget())
	}

	if ew, waitOK := sync.(ExtraWaitPostCreateDelete); waitOK {
		time.Sleep(ew.ExtraWaitPostCreateDelete())
	}

	if e == nil {
		sync.VoidState()
	} else {
		handleMissingResourceError(sync, &e)
	}

	return
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

// waitForStateRefresh takes a StatefulResource, a timeout duration, a list of states to treat as Pending, and a list of states to treat as Target. It uses those to wrap resource.StateChangeConf.WaitForState(). If the resource returns a missing status, it will not be treated as an error.
//
// sync.D.Id must be set.
// It does not set state from that refreshed state.
func waitForStateRefresh(sync StatefulResource, timeout time.Duration, operationName string, pending, target []string) (e error) {
	// TODO: try to move this onto sync
	stateConf := &resource.StateChangeConf{
		Pending: pending,
		Target:  target,
		Refresh: stateRefreshFunc(sync),
		Timeout: timeout,
	}

	if _, e = stateConf.WaitForState(); e != nil {
		handleMissingResourceError(sync, &e)
		return
	}
	if sync.State() == FAILED {
		return fmt.Errorf("Resource %s failed, state FAILED", operationName)
	}

	return
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

func FieldDeprecated(deprecatedFieldName string) string {
	return fmt.Sprintf("The '%s' field has been deprecated. It is no longer supported.", deprecatedFieldName)
}

func FieldDeprecatedForAnother(deprecatedFieldName string, newFieldName string) string {
	return fmt.Sprintf("The '%s' field has been deprecated. Please use '%s' instead.", deprecatedFieldName, newFieldName)
}

func FieldDeprecatedAndOverridenByAnother(deprecatedFieldName string, newFieldName string) string {
	return fmt.Sprintf("%s If both fields are specified, then '%s' will be used.",
		FieldDeprecatedForAnother(deprecatedFieldName, newFieldName), newFieldName)
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
