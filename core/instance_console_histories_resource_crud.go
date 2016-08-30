package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type InstanceConsoleHistoriesResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ConsoleHistoryMetadata
}

func (s *InstanceConsoleHistoriesResourceCrud) ID() string {
	return s.Res.ID
}

func (s *InstanceConsoleHistoriesResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceRequested}
}

func (s *InstanceConsoleHistoriesResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceSucceeded}
}

func (s *InstanceConsoleHistoriesResourceCrud) State() string {
	return s.Res.State
}

func (s *InstanceConsoleHistoriesResourceCrud) Create() (e error) {
	opts := baremetal.Options{}
	instanceID := s.D.Get("instance_id").(string)

	s.Res, e = s.Client.CaptureConsoleHistory(instanceID, opts)

	return
}

func (s *InstanceConsoleHistoriesResourceCrud) Get() (e error) {
	opts := baremetal.Options{}
	s.Res, e = s.Client.GetConsoleHistory(s.D.Id(), opts)
	return
}

func (s *InstanceConsoleHistoriesResourceCrud) Delete() (e error) {
	return s.Client.DeleteConsoleHistory(s.D.Id())
}

func (s *InstanceConsoleHistoriesResourceCrud) SetData() {
	s.D.Set("availability_domain", s.Res.AvailabilityDomain)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}
