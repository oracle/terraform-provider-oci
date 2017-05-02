// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func SupportedOperationDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readSupportedOperations,
		Schema: map[string]*schema.Schema{
			"supported_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     SupportedOperationResource(),
			},
		},
	}
}

func SupportedOperationResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSupportedOperations(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &SupportedOperationDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}
