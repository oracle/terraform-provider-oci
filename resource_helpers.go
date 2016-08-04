package main

import (
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func createResource(d *schema.ResourceData, create CreateResourceFn, get GetResourceFn) (e error) {
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	var res *baremtlsdk.Resource
	if res, e = create(name, description); e != nil {
		return
	}

	d.SetId(res.ID)
	setResourceData(d, res)

	if res.State != baremtlsdk.ResourceCreated {
		res, e = waitForStateRefresh(d, get)
	}

	return
}

func readResource(d *schema.ResourceData, get GetResourceFn) (e error) {
	var res *baremtlsdk.Resource
	if res, e = get(d.Id()); e != nil {
		return
	}

	setResourceData(d, res)

	return
}

func updateResource(d *schema.ResourceData, update UpdateResourceFn) (e error) {
	desc := d.Get("description").(string)

	d.Partial(true)
	var res *baremtlsdk.Resource
	if res, e = update(d.Id(), desc); e != nil {
		return
	}
	d.Partial(false)

	setResourceData(d, res)

	return
}

func destroyResource(d *schema.ResourceData, del DeleteResourceFn) (e error) {
	if e = del(d.Id()); e != nil {
		return
	}

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

func stateRefreshFunc(d *schema.ResourceData, get GetResourceFn) resource.StateRefreshFunc {
	return func() (res interface{}, s string, e error) {
		if res, e = get(d.Id()); e != nil {
			return nil, "", e
		}
		s = res.(*baremtlsdk.Resource).State
		return
	}
}

func waitForStateRefresh(d *schema.ResourceData, get GetResourceFn) (res *baremtlsdk.Resource, e error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{baremtlsdk.ResourceCreating},
		Target:  []string{baremtlsdk.ResourceCreated},
		Refresh: stateRefreshFunc(d, get),
		Timeout: 5 * time.Minute,
	}

	raw, err := stateConf.WaitForState()
	res = raw.(*baremtlsdk.Resource)
	if e = err; e != nil {
		return
	}

	setResourceData(d, res)

	return
}
