package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
)

type CompartmentSync struct {
	*IdentitySync
	D      *schema.ResourceData
	Client BareMetalClient
	Res    *baremetal.IdentityResource
}

func (s *CompartmentSync) Id() string {
	return s.Res.ID
}

func (s *CompartmentSync) State() string {
	return s.Res.State
}

func (s *CompartmentSync) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.CreateCompartment(name, description)
	return
}

func (s *CompartmentSync) Get() (e error) {
	s.Res, e = s.Client.GetCompartment(s.D.Id())
	return
}

func (s *CompartmentSync) Update() (e error) {
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.UpdateCompartment(s.D.Id(), description)
	return
}

func (s *CompartmentSync) SetData() {
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_modified", s.Res.TimeModified.String())
	s.D.Set("time_created", s.Res.TimeCreated.String())
}
