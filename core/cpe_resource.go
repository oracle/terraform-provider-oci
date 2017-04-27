// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func CpeResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: crud.DefaultTimeout,
		Create:   createCpe,
		Read:     readCpe,
		Update:   updateCpe,
		Delete:   deleteCpe,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}

}

func createCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &CpeResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &CpeResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &CpeResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.UpdateResource(d, crd)
}

func deleteCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &CpeResourceCrud{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}
