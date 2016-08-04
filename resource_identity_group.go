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
		Schema: resourceSchema,
	}
}

func createGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return createResource(d, client.CreateGroup, client.GetGroup)
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
