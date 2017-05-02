// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DBSystemShapeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBSystemShapes,
		Schema: map[string]*schema.Schema{
			"db_system_shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"available_core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func readDBSystemShapes(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &DBSystemShapeDatasourceCrud{}
	reader.D = d
	reader.Client = client
	return crud.ReadResource(reader)
}
