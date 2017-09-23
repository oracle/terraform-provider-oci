// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VolumeBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVolumeBackup,
		Read:     readVolumeBackup,
		Update:   updateVolumeBackup,
		Delete:   deleteVolumeBackup,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
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
			"size_in_mbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_request_received": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_size_in_mbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"volume_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createVolumeBackup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &VolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readVolumeBackup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &VolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateVolumeBackup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &VolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteVolumeBackup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &VolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.DeleteResource(d, sync)
}

type VolumeBackupResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.VolumeBackup
}

func (s *VolumeBackupResourceCrud) ID() string {
	return s.Res.ID
}

func (s *VolumeBackupResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceRequestReceived, baremetal.ResourceCreating}
}

// Creating is considered "Created" because it can take some time to finish
// actually creating and uploading the backup.
func (s *VolumeBackupResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *VolumeBackupResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *VolumeBackupResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *VolumeBackupResourceCrud) State() string {
	return s.Res.State
}

func (s *VolumeBackupResourceCrud) Create() (e error) {
	opts := &baremetal.CreateOptions{}
	volumeID := s.D.Get("volume_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.CreateVolumeBackup(volumeID, opts)

	return
}

func (s *VolumeBackupResourceCrud) Get() (e error) {
	res, e := s.Client.GetVolumeBackup(s.D.Id())
	if e == nil {
		s.Res = res
	}
	return
}

func (s *VolumeBackupResourceCrud) Update() (e error) {
	opts := &baremetal.IfMatchDisplayNameOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.UpdateVolumeBackup(s.D.Id(), opts)

	return
}

func (s *VolumeBackupResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("state", s.Res.State)
	s.D.Set("size_in_mbs", s.Res.SizeInMBs)
	if !s.Res.TimeCreated.IsZero() {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}
	s.D.Set("time_request_received", s.Res.TimeCreated.String())
	s.D.Set("unique_size_in_mbs", s.Res.SizeInMBs)
	s.D.Set("volume_id", s.Res.VolumeID)
}

func (s *VolumeBackupResourceCrud) Delete() (e error) {
	return s.Client.DeleteVolumeBackup(s.D.Id(), nil)
}
