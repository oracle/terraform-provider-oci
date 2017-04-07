// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func CpeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readCpeList,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cpes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     CpeResource(),
			},
		},
	}
}

func readCpeList(d *schema.ResourceData, m interface{}) (e error) {
	reader := &CPEDatasourceCrud{}
	reader.D = d
	reader.Client = m.(client.BareMetalClient)
	return crud.ReadResource(reader)

}
