package identity

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

type UserSync struct {
	*crud.IdentitySync
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.IdentityResource
}

func (s *UserSync) Id() string {
	return s.Res.ID
}

func (s *UserSync) State() string {
	return s.Res.State
}

func (s *UserSync) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.CreateUser(name, description)
	return
}

func (s *UserSync) Get() (e error) {
	s.Res, e = s.Client.GetUser(s.D.Id())
	return
}

func (s *UserSync) Update() (e error) {
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.UpdateUser(s.D.Id(), description)
	return
}

func (s *UserSync) SetData() {
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_modified", s.Res.TimeModified.String())
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *UserSync) Delete() (e error) {
	return s.Client.DeleteUser(s.D.Id())
}
