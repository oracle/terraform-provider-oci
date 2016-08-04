package main

import "github.com/hashicorp/terraform/helper/schema"

// ResourceIdentityUser exposes a IdentityUser Resource
func ResourceIdentityUser() *schema.Resource {
	return &schema.Resource{
		Create: createUser,
		Read:   readUser,
		Update: updateUser,
		Delete: deleteUser,
		Schema: resourceSchema,
	}
}

func createUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return createResource(d, client.CreateUser, client.GetUser)
}

func readUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return readResource(d, client.GetUser)
}

func updateUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return updateResource(d, client.UpdateUser)
}

func deleteUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return destroyResource(d, client.DeleteUser)
}
