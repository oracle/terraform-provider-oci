// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"fmt"

	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

// ResourceIdentityCompartment exposes an IdentityCompartment Resource
func CompartmentResource() *schema.Resource {
	return &schema.Resource{
		Create: createCompartment,
		Read:   readCompartment,
		Update: updateCompartment,
		Delete: deleteCompartment,
		Schema: baseIdentitySchema,
	}
}

func createCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &CompartmentResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &CompartmentResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &CompartmentResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteCompartment(d *schema.ResourceData, m interface{}) (e error) {
	sync := &CompartmentResourceCrud{}
	sync.D = d
	return crud.DeleteResource(sync)
	return fmt.Errorf("compartment resource: compartment %v cannot be deleted", d.Id())
}
