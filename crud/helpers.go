// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package crud

import (
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"errors"
	"strconv"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
)

const FiveMinutes time.Duration = 5 * time.Minute

type BaseCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
}

func (s *BaseCrud) VoidState() {
	s.D.SetId("")
}

// Default implementation, used in conjunction with State()
func (s *BaseCrud) setState(sync StatefulResource) error {
	v := reflect.ValueOf(sync).Elem()
	for _, resVal := range []reflect.Value{v.FieldByName("Res"), v.FieldByName("Resource")} {
		if resVal.IsValid() {
			err := s.D.Set("state", resVal.Elem().FieldByName("State").String())
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
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
		log.Println("Object does not exist, voiding and nullifying error")
		sync.VoidState()
		*err = nil
	}
}

func CreateResource(d *schema.ResourceData, sync ResourceCreator) (e error) {
	if e = sync.Create(); e != nil {
		return
	}
	d.SetId(sync.ID())
	sync.SetData()

	stateful, ok := sync.(StatefullyCreatedResource)
	if ok {
		pending := stateful.CreatedPending()
		target := stateful.CreatedTarget()
		e = waitForStateRefresh(stateful, pending, target)
	}

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

func DeleteResource(sync ResourceDeleter) (e error) {
	if e = sync.Delete(); e != nil {
		return
	}

	ew, waitOK := sync.(ExtraWaitPostDelete)

	stateful, ok := sync.(StatefullyDeletedResource)
	if ok {
		pending := stateful.DeletedPending()
		target := stateful.DeletedTarget()
		e = waitForStateRefresh(stateful, pending, target)
	}

	if waitOK {
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

func waitForStateRefresh(sync StatefulResource, pending, target []string) (e error) {
	timeoutStr := os.Getenv("TF_VAR_timeout_minutes")
	t, err := strconv.Atoi(timeoutStr)
	if err != nil {
		return errors.New("timeout_minutes: " + err.Error())
	}
	timeout := time.Duration(t) * time.Minute
	if customTimeouter, ok := sync.(CustomTimeouter); ok {
		if customTimeouter.CustomTimeout() > timeout {
			timeout = customTimeouter.CustomTimeout()
		}
	}
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
	sync.SetData()

	return
}
