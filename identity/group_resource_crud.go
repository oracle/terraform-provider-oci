package identity

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

type GroupSync struct {
	*crud.IdentitySync
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.IdentityResource
}

func (s *GroupSync) ID() string {
	return s.Res.ID
}

func (s *GroupSync) State() string {
	return s.Res.State
}

func (s *GroupSync) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.CreateGroup(name, description)
	return
}

func (s *GroupSync) Get() (e error) {
	s.Res, e = s.Client.GetGroup(s.D.Id())
	return
}

func (s *GroupSync) Update() (e error) {
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.UpdateGroup(s.D.Id(), description)
	return
}

func (s *GroupSync) SetData() {
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_modified", s.Res.TimeModified.String())
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *GroupSync) Delete() (e error) {
	return s.Client.DeleteGroup(s.D.Id())
}
