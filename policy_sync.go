package main

import (
	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/schema"
)

type PolicySync struct {
	D      *schema.ResourceData
	Client BareMetalClient
	Res    *baremtlsdk.Policy
}

func (s *PolicySync) Id() string {
	return s.Res.ID
}

func (s *PolicySync) State() string {
	return s.Res.State
}

func (s *PolicySync) toStringArray(vals interface{}) []string {
	arr := vals.([]interface{})
	result := []string{}
	for _, val := range arr {
		result = append(result, val.(string))
	}
	return result
}

func (s *PolicySync) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	statements := s.toStringArray(s.D.Get("statements"))

	s.Res, e = s.Client.CreatePolicy(name, description, statements)
	return
}

func (s *PolicySync) Get() (e error) {
	s.Res, e = s.Client.GetPolicy(s.D.Id())
	return
}

func (s *PolicySync) Update() (e error) {
	description := s.D.Get("description").(string)
	statements := s.toStringArray(s.D.Get("statements"))
	s.Res, e = s.Client.UpdatePolicy(s.D.Id(), description, statements)
	return
}

func (s *PolicySync) SetData() {
	s.D.Set("statements", s.Res.Statements)
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_modified", s.Res.TimeModified.String())
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *PolicySync) Delete() (e error) {
	return s.Client.DeletePolicy(s.D.Id())
}
