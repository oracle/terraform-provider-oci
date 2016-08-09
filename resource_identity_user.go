package main

import "github.com/hashicorp/terraform/helper/schema"

// ResourceIdentityUser exposes a IdentityUser Resource
func ResourceIdentityUser() *schema.Resource {
	return &schema.Resource{
		Create: createUser,
		Read:   readUser,
		Update: updateUser,
		Delete: deleteUser,
		Schema: identitySchema,
	}
}

func createUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &UserSync{D: d, Client: client}
	return createResource(d, sync)
}

func readUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &UserSync{D: d, Client: client}
	return readResource(sync)
}

func updateUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &UserSync{D: d, Client: client}
	return updateResource(d, sync)
}

func deleteUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &UserSync{D: d, Client: client}
	return sync.Delete()
}
