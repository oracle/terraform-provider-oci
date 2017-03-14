// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DHCPOptionsResource() *schema.Resource {
	return &schema.Resource{
		Create: createDHCPOptions,
		Read:   readDHCPOptions,
		Update: updateDHCPOptions,
		Delete: deleteDHCPOptions,
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
			"options": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"custom_dns_servers": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"server_type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
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

func createDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.CreateResource(d, crd)
}

func readDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.ReadResource(crd)
}

func updateDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.UpdateResource(d, crd)
}

func deleteDHCPOptions(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &DHCPOptionsResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.DeleteResource(crd)
}
