package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type VolumeAttachmentResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.VolumeAttachment
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
	compartmentID := s.D.Get("compartment_id").(string)
	instanceID := s.D.Get("instance_id").(string)
	volumeID := s.D.Get("volume_id").(string)

	s.Res, e = s.Client.AttachVolume(compartmentID, instanceID, attachmentType, volumeID)

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
}

func (s *VolumeAttachmentResourceCrud) Delete() (e error) {
	return s.Client.DetachVolume(s.D.Id())
}
