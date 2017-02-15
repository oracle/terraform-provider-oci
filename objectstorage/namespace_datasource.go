// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package objectstorage

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
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
