// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func NamespaceDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readNamespace,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readNamespace(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &NamespaceDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}
