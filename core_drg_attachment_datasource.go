// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DrgAttachmentDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDrgAttachments,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DrgAttachmentResource(),
			},
		},
	}
}

func readDrgAttachments(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DrgAttachmentDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}
