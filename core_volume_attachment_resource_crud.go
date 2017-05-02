// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type VolumeAttachmentResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.VolumeAttachment
}

func (s *VolumeAttachmentResourceCrud) ID() string {
	return s.Res.ID
}

func (s *VolumeAttachmentResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceAttaching}
}

func (s *VolumeAttachmentResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAttached}
}

func (s *VolumeAttachmentResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDetaching}
}

func (s *VolumeAttachmentResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDetached}
}

func (s *VolumeAttachmentResourceCrud) State() string {
	return s.Res.State
}

func (s *VolumeAttachmentResourceCrud) Create() (e error) {
	attachmentType := s.D.Get("attachment_type").(string)
	instanceID := s.D.Get("instance_id").(string)
	volumeID := s.D.Get("volume_id").(string)

	s.Res, e = s.Client.AttachVolume(attachmentType, instanceID, volumeID, nil)

	return
}

func (s *VolumeAttachmentResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetVolumeAttachment(s.D.Id())
	return
}

func (s *VolumeAttachmentResourceCrud) SetData() {
	s.D.Set("attachment_type", s.Res.AttachmentType)
	s.D.Set("availability_domain", s.Res.AvailabilityDomain)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("instance_id", s.Res.InstanceID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("volume_id", s.Res.VolumeID)
	s.D.Set("chap_secret", s.Res.CHAPSecret)
	s.D.Set("chap_username", s.Res.CHAPUsername)
	s.D.Set("ipv4", s.Res.IPv4)
	s.D.Set("iqn", s.Res.IQN)
	s.D.Set("port", s.Res.Port)
}

func (s *VolumeAttachmentResourceCrud) Delete() (e error) {
	return s.Client.DetachVolume(s.D.Id(), nil)
}
