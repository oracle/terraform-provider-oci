package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

// ResourceIdentityCompartment exposes an IdentityCompartment Resource
func ResourceIdentityCompartment() *schema.Resource {
	return &schema.Resource{
		Create: createCompartment,
		Read:   readCompartment,
		Update: updateCompartment,
		Delete: deleteCompartment,
		Schema: resourceSchema,
	}
}

func createCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return createResource(d, client.CreateCompartment, client.GetCompartment)
}

func readCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return readResource(d, client.GetCompartment)
}

func updateCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	return updateResource(d, client.UpdateCompartment)
}

func deleteCompartment(d *schema.ResourceData, m interface{}) (e error) {
	return fmt.Errorf("delete compartment %v: compartments cannot be deleted", d.Id())
}
