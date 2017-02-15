// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package crud

import (
	"time"

	"strings"

	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	//	"fmt"
	"fmt"
)

const FiveMinutes time.Duration = 5 * time.Minute

type BaseCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
}

func (s *BaseCrud) VoidState() {
	s.D.SetId("")
}

func handleMissingResourceError(sync ResourceVoider, err *error) {
	if err != nil && strings.Contains((*err).Error(), "does not exist") {
		fmt.Println("Object does not exist, voiding and nullifying error")
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

	stateful, ok := sync.(StatefullyDeletedResource)
	if ok {
		pending := stateful.DeletedPending()
		target := stateful.DeletedTarget()
		e = waitForStateRefresh(stateful, pending, target)
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
		return sync, sync.State(), e
	}
}

func waitForStateRefresh(sync StatefulResource, pending, target []string) (e error) {
	stateConf := &resource.StateChangeConf{
		Pending: pending,
		Target:  target,
		Refresh: stateRefreshFunc(sync),
		Timeout: FiveMinutes,
	}

	if _, e = stateConf.WaitForState(); e != nil {
		handleMissingResourceError(sync, &e)
		return
	}
	sync.SetData()

	return
}
