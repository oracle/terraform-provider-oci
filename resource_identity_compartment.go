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
	sync := &CompartmentSync{D: d, Client: client}
	return createResource(d, sync)
}

func readCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &CompartmentSync{D: d, Client: client}
	return readResource(sync)
}

func updateCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &CompartmentSync{D: d, Client: client}
	return updateResource(d, sync)
}

func deleteCompartment(d *schema.ResourceData, m interface{}) (e error) {
	return fmt.Errorf("compartment resource: compartment %v cannot be deleted", d.Id())
}
