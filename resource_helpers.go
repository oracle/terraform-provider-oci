package main

import (
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func createResource(d *schema.ResourceData, sync ResourceSync) (e error) {
	var res *baremtlsdk.Resource
	if res, e = sync.Create(); e != nil {
		return
	}
	d.SetId(res.ID)
	sync.SetData(res)

	if res.State != baremtlsdk.ResourceCreated {
		res, e = waitForStateRefresh(sync)
	}

	return
}

func readResource(sync ResourceSync) (e error) {
	if res, e = sync.Get(); e != nil {
		return
	}
	sync.SetData(res)

	return
}

func updateResource() (e error) {
	client := m.(BareMetalClient)
	d.Partial(true)
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

func setResourceData(d *schema.ResourceData, res *baremtlsdk.Resource) {
	d.Set("name", res.Name)
	d.Set("description", res.Description)
	d.Set("compartment_id", res.CompartmentID)
	d.Set("state", res.State)
	d.Set("time_modified", res.TimeModified.String())
	d.Set("time_created", res.TimeCreated.String())
}

func stateRefreshFunc(sync ResourceSync) resource.StateRefreshFunc {
	return func() (res interface{}, s string, e error) {
		if res, e = sync.Get(); e != nil {
			return nil, "", e
		}
		s = res.(*baremtlsdk.Resource).State
		return
	}
}

func waitForStateRefresh(sync ResourceSync) (res *baremtlsdk.Resource, e error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{baremtlsdk.ResourceCreating},
		Target:  []string{baremtlsdk.ResourceCreated},
		Refresh: stateRefreshFunc(sync),
		Timeout: 5 * time.Minute,
	}

	raw, err := stateConf.WaitForState()
	res = raw.(*baremtlsdk.Resource)
	if e = err; e != nil {
		return
	}

	sync.SetData(res)

	return
}
