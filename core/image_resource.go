// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func ImageResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: crud.DefaultTimeout,
		Create:   createImage,
		Read:     readImage,
		Update:   updateImage,
		Delete:   deleteImage,
		Schema: map[string]*schema.Schema{
			"base_image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"create_image_allowed": {
				Type:     schema.TypeBool,
				Computed: true,
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
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_system": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_system_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.DeleteResource(d, sync)
}
