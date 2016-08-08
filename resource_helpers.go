package main

import (
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

var resourceSchema = map[string]*schema.Schema{
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

func createResource(d *schema.ResourceData, sync ResourceSync) (e error) {
	if e = sync.Create(); e != nil {
		return
	}
	d.SetId(sync.Id())
	sync.SetData()

	if sync.State() != baremtlsdk.ResourceCreated {
		e = waitForStateRefresh(sync)
	}

	return
}

func readResource(sync ResourceSync) (e error) {
	if e = sync.Get(); e != nil {
		return
	}
	sync.SetData()

	return
}

func updateResource(d *schema.ResourceData, sync ResourceSync) (e error) {
	d.Partial(true)
	if e = sync.Update(); e != nil {
		return
	}
	d.Partial(false)
	sync.SetData()

	return
}

func stateRefreshFunc(sync ResourceSync) resource.StateRefreshFunc {
	return func() (res interface{}, s string, e error) {
		if e = sync.Get(); e != nil {
			return nil, "", e
		}
		s = sync.State()
		return
	}
}

func waitForStateRefresh(sync ResourceSync) (e error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{baremtlsdk.ResourceCreating},
		Target:  []string{baremtlsdk.ResourceCreated},
		Refresh: stateRefreshFunc(sync),
		Timeout: 5 * time.Minute,
	}

	if _, e = stateConf.WaitForState(); e != nil {
		return
	}

	sync.SetData()

	return
}
