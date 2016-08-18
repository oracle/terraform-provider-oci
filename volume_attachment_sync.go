package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
)

type VolumeAttachmentSync struct {
	D      *schema.ResourceData
	Client BareMetalClient
	Res    *baremetal.VolumeAttachment
}

func (s *VolumeAttachmentSync) Id() string {
	return s.Res.ID
}

func (s *VolumeAttachmentSync) CreatedPending() []string {
	return []string{baremetal.ResourceAttaching}
}

func (s *VolumeAttachmentSync) CreatedTarget() []string {
	return []string{baremetal.ResourceAttached}
}

func (s *VolumeAttachmentSync) DeletedPending() []string {
	return []string{baremetal.ResourceDetaching}
}

func (s *VolumeAttachmentSync) DeletedTarget() []string {
	return []string{baremetal.ResourceDetached}
}

func (s *VolumeAttachmentSync) State() string {
	return s.Res.State
}

func (s *VolumeAttachmentSync) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	instanceID := s.D.Get("instanceID").(string)
	attachmentType := s.D.Get("attachmentType").(string)
	volumeID := s.D.Get("volumeID").(string)

	s.Res, e = s.Client.AttachVolume(compartmentID, instanceID, attachmentType, volumeID)

	return
}

func (s *VolumeAttachmentSync) Get() (e error) {
	s.Res, e = s.Client.GetVolumeAttachment(s.D.Id())
	return
}

func (s *VolumeAttachmentSync) SetData() {
	s.D.Set("attachment_type", s.Res.AttachmentType)
	s.D.Set("availability_domain", s.Res.AvailabilityDomain)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("instance_id", s.Res.InstanceID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("volume_id", s.Res.VolumeID)
}

func (s *VolumeAttachmentSync) Delete() (e error) {
	return s.Client.DetachVolume(s.D.Id())
}
