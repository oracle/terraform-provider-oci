// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DrgDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDrgs,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drgs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DrgResource(),
			},
		},
	}
}

func readDrgs(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DrgDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}
