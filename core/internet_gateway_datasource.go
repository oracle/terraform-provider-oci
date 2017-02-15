// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func InternetGatewayDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readInternetGateways,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": {
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
			"gateways": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     InternetGatewayResource(),
			},
		},
	}
}

func readInternetGateways(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &InternetGatewayDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}
