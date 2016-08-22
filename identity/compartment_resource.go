package identity

import (
	"fmt"

	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

// ResourceIdentityCompartment exposes an IdentityCompartment Resource
func ResourceIdentityCompartment() *schema.Resource {
	return &schema.Resource{
		Create: createCompartment,
		Read:   readCompartment,
		Update: updateCompartment,
		Delete: deleteCompartment,
		Schema: identitySchema,
	}
}

func createCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &CompartmentSync{D: d, Client: client}
	return crud.CreateResource(d, sync)
}

func readCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &CompartmentSync{D: d, Client: client}
	return crud.ReadResource(sync)
}

func updateCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &CompartmentSync{D: d, Client: client}
	return crud.UpdateResource(d, sync)
}

func deleteCompartment(d *schema.ResourceData, m interface{}) (e error) {
	return fmt.Errorf("compartment resource: compartment %v cannot be deleted", d.Id())
}
