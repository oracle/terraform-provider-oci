package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type VolumeSync struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.Volume
}

func (s *VolumeSync) Id() string {
	return s.Res.ID
}

func (s *VolumeSync) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *VolumeSync) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *VolumeSync) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *VolumeSync) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *VolumeSync) State() string {
	return s.Res.State
}

func (s *VolumeSync) Create() (e error) {
	opts := baremetal.Options{}
	availabilityDosync := s.D.Get("availability_domain").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.CreateVolume(availabilityDosync, compartmentID, opts)

	return
}

func (s *VolumeSync) Get() (e error) {
	s.Res, e = s.Client.GetVolume(s.D.Id())
	return
}

func (s *VolumeSync) Update() (e error) {
	opts := baremetal.Options{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.UpdateVolume(s.D.Id(), opts)

	return
}

func (s *VolumeSync) SetData() {
	s.D.Set("availability_domain", s.Res.AvailabilityDomain)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("size_in_mbs", s.Res.SizeInMBs)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *VolumeSync) Delete() (e error) {
	return s.Client.DeleteVolume(s.D.Id())
}
