// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func CompartmentDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readCompartments,
		Schema: map[string]*schema.Schema{
			"compartment_id":{
				Type:     schema.TypeString,
				Required: true,
			},
			"compartments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     CompartmentResource(),
			},
		},
	}
}

func readCompartments(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &CompartmentDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}
