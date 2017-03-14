// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type VolumeBackupResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.VolumeBackup
}

func (s *VolumeBackupResourceCrud) ID() string {
	return s.Res.ID
}

func (s *VolumeBackupResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceRequestReceived}
}

// Creating is considered "Created" because it can take some time to finish
// actually creating and uploading the backup.
func (s *VolumeBackupResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceCreating, baremetal.ResourceAvailable}
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
	s.Res, e = s.Client.GetVolumeBackup(s.D.Id())
	return
}

func (s *VolumeBackupResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateBackupOptions{}
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
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("time_request_received", s.Res.TimeCreated.String())
	s.D.Set("unique_size_in_mbs", s.Res.SizeInMBs)
	s.D.Set("volume_id", s.Res.VolumeID)
}

func (s *VolumeBackupResourceCrud) Delete() (e error) {
	return s.Client.DeleteVolumeBackup(s.D.Id(), nil)
}
