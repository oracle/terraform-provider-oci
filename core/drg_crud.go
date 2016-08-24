package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type DrgSync struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.Drg
}

func (s *DrgSync) ID() string {
	return s.Res.ID
}

func (s *DrgSync) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *DrgSync) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *DrgSync) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *DrgSync) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *DrgSync) State() string {
	return s.Res.State
}

func (s *DrgSync) Create() (e error) {
	opts := baremetal.Options{}
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.CreateDrg(compartmentID, opts)

	return
}

func (s *DrgSync) Get() (e error) {
	s.Res, e = s.Client.GetDrg(s.D.Id())
	return
}

func (s *DrgSync) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *DrgSync) Delete() (e error) {
	return s.Client.DeleteDrg(s.D.Id())
}
