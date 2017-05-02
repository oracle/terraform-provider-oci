// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DBVersionDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBVersions,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func readDBVersions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &DBVersionDatasourceCrud{}
	reader.D = d
	reader.Client = client
	return crud.ReadResource(reader)
}
