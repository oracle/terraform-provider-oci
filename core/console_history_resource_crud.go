package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type ConsoleHistoryResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ConsoleHistoryMetadata
}

func (s *ConsoleHistoryResourceCrud) ID() string {
	return s.Res.ID
}

func (s *ConsoleHistoryResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceRequested}
}

func (s *ConsoleHistoryResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceSucceeded}
}

func (s *ConsoleHistoryResourceCrud) State() string {
	return s.Res.State
}

func (s *ConsoleHistoryResourceCrud) Create() (e error) {
	instanceID := s.D.Get("instance_id").(string)

	s.Res, e = s.Client.CaptureConsoleHistory(instanceID, nil)

	return
}

func (s *ConsoleHistoryResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetConsoleHistory(s.D.Id())
	return
}

func (s *ConsoleHistoryResourceCrud) SetData() {
	s.D.Set("availability_domain", s.Res.AvailabilityDomain)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("instance_id", s.Res.InstanceID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}
