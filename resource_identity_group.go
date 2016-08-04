package main

import "github.com/hashicorp/terraform/helper/schema"

// ResourceIdentityGroup exposes an IdentityGroup Resource
func ResourceIdentityGroup() *schema.Resource {
	return &schema.Resource{
		Create: createGroup,
		Read:   readGroup,
		Update: updateGroup,
		Delete: deleteGroup,
		Schema: resourceSchema,
	}
}

func createGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return createResource(d, client.CreateGroup, client.GetGroup)
}

func readGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return readResource(d, client.GetGroup)
}

func updateGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return updateResource(d, client.UpdateGroup)
}

func deleteGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return destroyResource(d, client.DeleteGroup)
}
