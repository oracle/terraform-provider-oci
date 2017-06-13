// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package crud

import (
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/terraform-provider-baremetal/client"
)

var (
	FiveMinutes    time.Duration = 5 * time.Minute
	TwoHours       time.Duration = 120 * time.Minute
	DefaultTimeout               = &schema.ResourceTimeout{
		Create: &FiveMinutes,
		Update: &FiveMinutes,
		Delete: &FiveMinutes,
	}
)

type BaseCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
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
				if stateValue := resourceValue.FieldByName("State"); stateValue.IsValid() {
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

func LoadBalancerResourceID(res interface{}, workReq *baremetal.WorkRequest) (id *string, workReqSucceeded bool) {
	v := reflect.ValueOf(res).Elem()
	if v.IsValid() {
		// This is super fugly. It's this way because the LB API has no convention for ID formats.

		// Load balancer
		id := v.FieldByName("ID")
		if id.IsValid() {
			s := id.String()
			return &s, false
		}
		// backendset, listener
		name := v.FieldByName("Name")
		if name.IsValid() {
			s := name.String()
			return &s, false
		}
		// certificate
		certName := v.FieldByName("CertificateName")
		if certName.IsValid() {
			s := certName.String()
			return &s, false
		}
		// backend
		ip := v.FieldByName("ip_address")
		port := v.FieldByName("port")
		if ip.IsValid() && port.IsValid() {
			s := ip.String() + ":" + strconv.Itoa(int(int(port.Int())))
			return &s, false
		}
	}
	if workReq != nil {
		if workReq.State == baremetal.WorkRequestSucceeded {
			return nil, true
		} else {
			return &workReq.ID, false
		}
	}
	return nil, false
}

func LoadBalancerResourceGet(s BaseCrud, workReq *baremetal.WorkRequest) (id string, stillWorking bool, err error) {
	id = s.D.Id()
	// NOTE: if the id is for a work request, refresh its state and loadBalancerID.
	if strings.HasPrefix(id, "ocid1.loadbalancerworkrequest.") {
		updatedWorkReq, err := s.Client.GetWorkRequest(id, nil)
		if err != nil {
			return "", false, err
		}
		if workReq != nil {
			*workReq = *updatedWorkReq
			s.D.Set("state", workReq.State)
			if workReq.State == baremetal.WorkRequestSucceeded {
				return "", false, nil
			}
		}
		return "", true, nil
	}
	return id, false, nil
}

func LoadBalancerWaitForWorkRequest(client client.BareMetalClient, d *schema.ResourceData, wr *baremetal.WorkRequest) error {
	var e error
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			baremetal.ResourceWaitingForWorkRequest,
			baremetal.WorkRequestInProgress,
			baremetal.WorkRequestAccepted,
		},
		Target:  []string{
			baremetal.ResourceSucceededWorkRequest,
			baremetal.WorkRequestSucceeded,
			baremetal.ResourceFailed,
		},
		Refresh: func() (interface{}, string, error) {
			wr, e = client.GetWorkRequest(wr.ID, nil)
			return wr, wr.State, e
		},
		Timeout: d.Timeout(schema.TimeoutCreate),
	}

	if _, e = stateConf.WaitForState(); e != nil {
		return e
	}
	if wr.State == baremetal.ResourceFailed {
		return errors.New("Resource creation failed, state FAILED")
	}
	return nil
}


func CreateResource(d *schema.ResourceData, sync ResourceCreator) (e error) {
	if e = sync.Create(); e != nil {
		// Check for conflicts and retry
		// This happens with concurrent volume attachments, etc
		if strings.Contains(strings.ToLower(e.Error()), "try again later") {
			log.Println("[DEBUG] Resource creation conflicts with other resources. Waiting 10 seconds and trying again...")
			time.Sleep(10 * time.Second)
			e = CreateResource(d, sync)
		}
		return e
	}

	// ID is required for state refresh
	d.SetId(sync.ID())

	if stateful, ok := sync.(StatefullyCreatedResource); ok {
		e = waitForStateRefresh(stateful, d.Timeout(schema.TimeoutCreate), stateful.CreatedPending(), stateful.CreatedTarget())
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

	/* Attempt at #113, but this breaks everything. Probably because this is used internally by other state checking mechanisms.

	if dr, ok := sync.(StatefullyDeletedResource); ok {
		for _, target := range dr.DeletedTarget() {
			if dr.State() == target {
				dr.VoidState()
				return
			}
		}
	}*/

	return
}

func UpdateResource(d *schema.ResourceData, sync ResourceUpdater) (e error) {
	d.Partial(true)
	if e = sync.Update(); e != nil {
		return
	}
	d.Partial(false)
	sync.SetData()

	return
}

// DeleteResource requests a Delete(). If the resource deletes
// statefully (not immediately), poll State to ensure:
// () -> Pending -> Deleted.
// Finally, sets the ResourceData state to empty.
func DeleteResource(d *schema.ResourceData, sync ResourceDeleter) (e error) {
	if e = sync.Delete(); e != nil {
		handleMissingResourceError(sync, &e)
		if e != nil {
			return
		}
	}

	//d.SetId(sync.ID())
	if stateful, ok := sync.(StatefullyDeletedResource); ok {
		e = waitForStateRefresh(stateful, d.Timeout(schema.TimeoutDelete), stateful.DeletedPending(), stateful.DeletedTarget())
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
func waitForStateRefresh(sync StatefulResource, timeout time.Duration, pending, target []string) (e error) {
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
	if sync.State() == baremetal.ResourceFailed {
		return errors.New("Resource creation failed, state FAILED")
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
