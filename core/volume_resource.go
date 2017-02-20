// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func VolumeResource() *schema.Resource {
	return &schema.Resource{
		Create: createVolume,
		Read:   readVolume,
		Update: updateVolume,
		Delete: deleteVolume,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
<<<<<<< HEAD
			"size_in_mbs": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
=======
>>>>>>> b138c42f8af0008f6e01c2a7333a372a4ade6853
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
<<<<<<< HEAD
=======
			"size_in_mbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
>>>>>>> b138c42f8af0008f6e01c2a7333a372a4ade6853
			"state": {
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

func createVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}
