// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VolumeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVolume,
		Read:     readVolume,
		Update:   updateVolume,
		Delete:   deleteVolume,
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

func createVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}

type VolumeResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.Volume
}

func (s *VolumeResourceCrud) ID() string {
	return s.Res.ID
}

func (s *VolumeResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning, baremetal.ResourceRestoring}
}

func (s *VolumeResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *VolumeResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *VolumeResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *VolumeResourceCrud) State() string {
	return s.Res.State
}

func (s *VolumeResourceCrud) Create() (e error) {
	availabilityDomain := s.D.Get("availability_domain").(string)
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.CreateVolumeOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}
	sizeInMBs, ok := s.D.GetOk("size_in_mbs")
	if ok {
		opts.SizeInMBs = sizeInMBs.(int)
	}
	volumeBackupID, ok := s.D.GetOk("volume_backup_id")
	if ok {
		opts.VolumeBackupID = volumeBackupID.(string)
	}

	s.Res, e = s.Client.CreateVolume(availabilityDomain, compartmentID, opts)

	return
}

func (s *VolumeResourceCrud) Get() (e error) {
	res, e := s.Client.GetVolume(s.D.Id())
	if e == nil {
		s.Res = res
	}
	return
}

func (s *VolumeResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.UpdateVolume(s.D.Id(), opts)

	return
}

func (s *VolumeResourceCrud) SetData() {
	s.D.Set("availability_domain", s.Res.AvailabilityDomain)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("size_in_mbs", s.Res.SizeInMBs)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *VolumeResourceCrud) Delete() (e error) {
	return s.Client.DeleteVolume(s.D.Id(), nil)
}
