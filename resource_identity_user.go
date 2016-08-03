package main

import (
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

// ResourceIdentityUser exposes a IdentityUser Resource
func ResourceIdentityUser() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: destroy,

		Schema: map[string]*schema.Schema{
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
		},
	}
}

func create(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	var res *baremtlsdk.Resource
	if res, e = client.CreateUser(name, description); e != nil {
		return
	}

	// Set the id and set any fields that were returned by the API.
	d.SetId(res.ID)
	setResourceData(d, res)

	if res.State != baremtlsdk.ResourceCreated {
		res, e = waitForStateRefresh(d, client)
	}

	return
}

func read(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	var res *baremtlsdk.Resource
	if res, e = client.GetUser(d.Id()); e != nil {
		return
	}

	setResourceData(d, res)

	return
}

func update(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	desc := d.Get("description").(string)

	d.Partial(true)
	var res *baremtlsdk.Resource
	if res, e = client.UpdateUser(d.Id(), desc); e != nil {
		return
	}
	d.Partial(false)

	setResourceData(d, res)

	return
}

func destroy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	if e = client.DeleteUser(d.Id()); e != nil {
		return
	}

	return
}

func stateRefreshFunc(client BareMetalClient, d *schema.ResourceData) resource.StateRefreshFunc {
	return func() (res interface{}, s string, e error) {
		if res, e = client.GetUser(d.Id()); e != nil {
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

func waitForStateRefresh(d *schema.ResourceData, c BareMetalClient) (res *baremtlsdk.Resource, e error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{baremtlsdk.ResourceCreating},
		Target:  []string{baremtlsdk.ResourceCreated},
		Refresh: stateRefreshFunc(c, d),
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
