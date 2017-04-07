// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func SubnetResource() *schema.Resource {
	return &schema.Resource{
		Create: createSubnet,
		Read:   readSubnet,
		Update: updateSubnet,
		Delete: deleteSubnet,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cidr_block": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"route_table_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_list_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_options_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
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
			"virtual_router_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_router_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createSubnet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.CreateResource(d, sync)
}

func readSubnet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.ReadResource(sync)
}

func updateSubnet(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(sync.D, sync)
}

func deleteSubnet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(sync)
}
