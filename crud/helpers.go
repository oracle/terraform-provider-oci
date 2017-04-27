// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package crud

import (
	"log"
	"os"
	"reflect"
	"strings"
	"time"

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
	if err != nil && strings.Contains((*err).Error(), "does not exist") {
		log.Println("[DEBUG] Object does not exist, voiding resource and nullifying error")
		sync.VoidState()
		*err = nil
	}
}

func CreateResource(d *schema.ResourceData, sync ResourceCreator) (e error) {
	if e = sync.Create(); e != nil {
		return e
	}

	// ID is required for state refresh
	d.SetId(sync.ID())

	if stateful, ok := sync.(StatefullyCreatedResource); ok {
		e = waitForStateRefresh(stateful, d.Timeout(schema.TimeoutCreate), stateful.CreatedPending(), stateful.CreatedTarget())
	}

	d.SetId(sync.ID())
	sync.SetData()

	return
}

func ReadResource(sync ResourceReader) (e error) {
	if e = sync.Get(); e != nil {
		handleMissingResourceError(sync, &e)
		return
	}
	sync.SetData()

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
		return
	}

	if stateful, ok := sync.(StatefullyDeletedResource); ok {
		e = waitForStateRefresh(stateful, d.Timeout(schema.TimeoutDelete), stateful.DeletedPending(), stateful.DeletedTarget())

	}

	if ew, waitOK := sync.(ExtraWaitPostDelete); waitOK {
		if os.Getenv("TF_ORACLE_ENV") != "test" {
			time.Sleep(ew.ExtraWaitPostDelete())
		}
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
// used by:
// - CreateResource()
// - DeleteResource()
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

	return
}
