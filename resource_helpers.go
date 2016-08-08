package main

import (
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func createResource(d *schema.ResourceData, sync ResourceSync) (e error) {
	var res BareMetalResource
	if res, e = sync.Create(); e != nil {
		return
	}
	d.SetId(res.GetId())
	sync.SetData(res)

	if res.GetState() != baremtlsdk.ResourceCreated {
		res, e = waitForStateRefresh(sync)
	}

	return
}

func readResource(sync ResourceSync) (e error) {
	var res BareMetalResource
	if res, e = sync.Get(); e != nil {
		return
	}
	sync.SetData(res)

	return
}

func updateResource(d *schema.ResourceData, sync ResourceSync) (e error) {
	d.Partial(true)
	var res BareMetalResource
	if res, e = sync.Update(); e != nil {
		return
	}
	d.Partial(false)
	sync.SetData(res)

	return
}

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

func stateRefreshFunc(sync ResourceSync) resource.StateRefreshFunc {
	return func() (res interface{}, s string, e error) {
		if res, e = sync.Get(); e != nil {
			return nil, "", e
		}
		s = res.(BareMetalResource).GetState()
		return
	}
}

func waitForStateRefresh(sync ResourceSync) (res BareMetalResource, e error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{baremtlsdk.ResourceCreating},
		Target:  []string{baremtlsdk.ResourceCreated},
		Refresh: stateRefreshFunc(sync),
		Timeout: 5 * time.Minute,
	}

	raw, err := stateConf.WaitForState()
	res = raw.(BareMetalResource)
	if e = err; e != nil {
		return
	}

	sync.SetData(res)

	return
}
