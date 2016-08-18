package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

const fiveMinutes time.Duration = 5 * time.Minute

var identitySchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"description": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},
	"compartment_id": &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	},
	"state": &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	},
	"time_created": &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	},
	"time_modified": &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	},
}

func createResource(d *schema.ResourceData, sync ResourceCreator) (e error) {
	if e = sync.Create(); e != nil {
		return
	}
	d.SetId(sync.Id())
	sync.SetData()

	stateful, ok := sync.(StatefullyCreatedResource)
	if ok {
		pending := stateful.CreatedPending()
		target := stateful.CreatedTarget()
		e = waitForStateRefresh(stateful, pending, target)
	}

	return
}

func readResource(sync ResourceReader) (e error) {
	if e = sync.Get(); e != nil {
		return
	}
	sync.SetData()

	return
}

func updateResource(d *schema.ResourceData, sync ResourceUpdater) (e error) {
	d.Partial(true)
	if e = sync.Update(); e != nil {
		return
	}
	d.Partial(false)
	sync.SetData()

	return
}

func deleteResource(sync ResourceDeleter) (e error) {
	if e = sync.Delete(); e != nil {
		return
	}

	stateful, ok := sync.(StatefullyDeletedResource)
	if ok {
		pending := stateful.DeletedPending()
		target := stateful.DeletedTarget()
		e = waitForStateRefresh(stateful, pending, target)
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
		Timeout: fiveMinutes,
	}

	if _, e = stateConf.WaitForState(); e != nil {
		return
	}

	sync.SetData()

	return
}
