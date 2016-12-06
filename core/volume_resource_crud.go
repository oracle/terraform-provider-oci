package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type VolumeResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.Volume
}

func (s *VolumeResourceCrud) ID() string {
	return s.Res.ID
}

func (s *VolumeResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
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

	s.Res, e = s.Client.CreateVolume(availabilityDomain, compartmentID, opts)

	return
}

func (s *VolumeResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetVolume(s.D.Id())
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
