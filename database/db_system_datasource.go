// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package database

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DBSystemDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBSystems,
		Schema: map[string]*schema.Schema{
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
			"db_systems": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DBSystemResource(),
			},
		},
	}
}

func readDBSystems(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBSystemDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}
