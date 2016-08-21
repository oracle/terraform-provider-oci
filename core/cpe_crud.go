package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type CpeSync struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.Cpe
}

func (s *CpeSync) Id() string {
	return s.Resource.ID
}

func (s *CpeSync) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	displayName := s.D.Get("display_name").(string)
	ipAddress := s.D.Get("ip_address").(string)
	s.Resource, e = s.Client.CreateCpe(compartmentID, displayName, ipAddress)
	return
}

func (s *CpeSync) Get() (e error) {
	s.Resource, e = s.Client.GetCpe(s.D.Id())
	return
}

func (s *CpeSync) SetData() {
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("ip_address", s.Resource.IPAddress)
	s.D.Set("time_created", s.Resource.TimeCreated.String())
}

func (s *CpeSync) Delete() (e error) {
	return s.Client.DeleteCpe(s.D.Id())
}
