// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func RouteTableResource() *schema.Resource {
	return &schema.Resource{
		Create: createRouteTable,
		Read:   readRouteTable,
		Update: updateRouteTable,
		Delete: deleteRouteTable,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr_block": {
							Type:     schema.TypeString,
							Required: true,
						},
						"network_entity_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"time_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createRouteTable(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &RouteTableResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.CreateResource(d, crd)
}

func readRouteTable(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &RouteTableResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.ReadResource(crd)
}

func updateRouteTable(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &RouteTableResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.UpdateResource(d, crd)
}

func deleteRouteTable(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &RouteTableResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.DeleteResource(crd)
}
