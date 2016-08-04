package main

import (
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func createResource(d *schema.ResourceData, client BareMetalClient, create CreateResourceFn, get GetResourceFn) (e error) {
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	var res *baremtlsdk.Resource
	if res, e = create(name, description); e != nil {
		return
	}

	d.SetId(res.ID)
	setResourceData(d, res)

	if res.State != baremtlsdk.ResourceCreated {
		res, e = waitForStateRefresh(d, client, get)
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

func stateRefreshFunc(client BareMetalClient, d *schema.ResourceData, get GetResourceFn) resource.StateRefreshFunc {
	return func() (res interface{}, s string, e error) {
		if res, e = get(d.Id()); e != nil {
			return nil, "", e
		}
		s = res.(*baremtlsdk.Resource).State
		return
	}
}

func setResourceData(d *schema.ResourceData, res *baremtlsdk.Resource) {
	d.Set("name", res.Name)
	d.Set("description", res.Description)
	d.Set("compartment_id", res.CompartmentID)
	d.Set("state", res.State)
	d.Set("time_modified", res.TimeModified.String())
	d.Set("time_created", res.TimeCreated.String())
}

func waitForStateRefresh(d *schema.ResourceData, c BareMetalClient, get GetResourceFn) (res *baremtlsdk.Resource, e error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{baremtlsdk.ResourceCreating},
		Target:  []string{baremtlsdk.ResourceCreated},
		Refresh: stateRefreshFunc(c, d, get),
		Timeout: 5 * time.Minute,
	}

	raw, err := stateConf.WaitForState()
	res = raw.(*baremtlsdk.Resource)
	if e = err; e != nil {
		return
	}

	// Fields may have changed during polling, set them again.
	setResourceData(d, res)

	return
}
