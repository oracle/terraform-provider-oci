package main

import (
	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/schema"
)

type GroupSync struct {
	d      *schema.ResourceData
	client BareMetalClient
}

func (s *GroupSync) Create() (res BareMetalResource, e error) {
	name := s.d.Get("name").(string)
	description := s.d.Get("description").(string)
	var raw *baremtlsdk.Resource
	raw, e = s.client.CreateGroup(name, description)
	res = &BareMetalResourceAdapter{raw}
	return
}

func (s *GroupSync) Get() (res BareMetalResource, e error) {
	var raw *baremtlsdk.Resource
	raw, e = s.client.GetGroup(s.d.Id())
	res = &BareMetalResourceAdapter{raw}
	return
}

func (s *GroupSync) Update() (res BareMetalResource, e error) {
	description := s.d.Get("description").(string)
	var raw *baremtlsdk.Resource
	raw, e = s.client.UpdateGroup(s.d.Id(), description)
	res = &BareMetalResourceAdapter{raw}
	return
}

func (s *GroupSync) SetData(res BareMetalResource) {
	a := res.(*BareMetalResourceAdapter)
	s.d.Set("name", a.Name)
	s.d.Set("description", a.Description)
	s.d.Set("compartment_id", a.CompartmentID)
	s.d.Set("state", a.State)
	s.d.Set("time_modified", a.TimeModified.String())
	s.d.Set("time_created", a.TimeCreated.String())
}

func (s *GroupSync) Delete() (e error) {
	return s.client.DeleteGroup(s.d.Id())
}
