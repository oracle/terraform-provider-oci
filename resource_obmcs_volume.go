// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/oracle/terraform-provider-baremetal/client"
)

func VolumeResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceOBMCSVolumeCreate,
		Read:   resourceOBMCSVolumeRead,
		Update: resourceOBMCSVolumeUpdate,
		Delete: resourceOBMCSVolumeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
			"size_in_mbs": {
				Type:     schema.TypeInt,
				Optional: true,
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
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_backup_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceOBMCSVolumeCreate(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)

	// Ideally we would build a createVolumeInput struct and pass that into CreateVolume
	availabilityDomain := d.Get("availability_domain").(string)
	compartmentID := d.Get("compartment_id").(string)

	opts := &baremetal.CreateVolumeOptions{}
	displayName, ok := d.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}
	sizeInMBs, ok := d.GetOk("size_in_mbs")
	if ok {
		opts.SizeInMBs = sizeInMBs.(int)
	}
	volumeBackupID, ok := d.GetOk("volume_backup_id")
	if ok {
		opts.VolumeBackupID = volumeBackupID.(string)
	}

	result, e := client.CreateVolume(availabilityDomain, compartmentID, opts)
	if e != nil {
		return e
	}

	d.SetId(result.ID)

	return resourceOBMCSVolumeRead(d, m)
}

func resourceOBMCSVolumeRead(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	name := d.Id()
	volume, e := client.GetVolume(name)
	if e != nil {
		// Check if error was because we coudn't find the resource
		// Set Id to nil and return nil
		// Else
		return fmt.Errorf("Error reading volume %s: %s", name, e)
	}

	d.Set("availability_domain", volume.AvailabilityDomain)
	d.Set("compartment_id", volume.CompartmentID)
	d.Set("display_name", volume.DisplayName)
	d.Set("size_in_mbs", volume.SizeInMBs)
	d.Set("state", volume.State)
	d.Set("time_created", volume.TimeCreated.String())

	return nil
}

func resourceOBMCSVolumeUpdate(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	name := d.Id()

	// Ideally we would build a updateVolumeInput struct and pass that into CreateVolume
	opts := &baremetal.UpdateOptions{}
	displayName, ok := d.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	_, e = client.UpdateVolume(name, opts)
	if e != nil {
		return e
	}
	return resourceOBMCSVolumeRead(d, m)
}

func resourceOBMCSVolumeDelete(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	name := d.Id()

  err := client.DeleteVolume(d.Id(), nil)
	if err != nil {
		return fmt.Errorf("Error deleting volume %s: %s", name, err)
	}
	return nil
}
