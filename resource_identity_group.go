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
	sync := &GroupSync{D: d, Client: client}
	return createResource(d, sync)
}

func readGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &GroupSync{D: d, Client: client}
	return readResource(sync)
}

func updateGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &GroupSync{D: d, Client: client}
	return updateResource(d, sync)
}

func deleteGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &GroupSync{D: d, Client: client}
	return sync.Delete()
}
