package main

import (
	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/schema"
)

// ResourceIdentityGroup exposes an IdentityGroup Resource
func ResourceIdentityGroup() *schema.Resource {
	return &schema.Resource{
		Create: createGroup,
		Read:   readGroup,
		Update: updateGroup,
		Delete: destroyGroup,

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

func createGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	var res *baremtlsdk.Resource
	if res, e = client.CreateGroup(name, description); e != nil {
		return
	}

	// Set the id and set any fields that were returned by the API.
	d.SetId(res.ID)
	setResourceData(d, res)

	if res.State != baremtlsdk.ResourceCreated {
		res, e = waitForStateRefresh(d, client, client.GetGroup)
	}

	return
}

func readGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	var res *baremtlsdk.Resource
	if res, e = client.GetGroup(d.Id()); e != nil {
		return
	}

	setResourceData(d, res)

	return
}

func updateGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	desc := d.Get("description").(string)

	d.Partial(true)
	var res *baremtlsdk.Resource
	if res, e = client.UpdateGroup(d.Id(), desc); e != nil {
		return
	}
	d.Partial(false)

	setResourceData(d, res)

	return
}

func destroyGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	if e = client.DeleteGroup(d.Id()); e != nil {
		return
	}

	return
}
