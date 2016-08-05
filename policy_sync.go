package main

import (
	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/schema"
)

type PolicySync struct {
	d      *schema.ResourceData
	client BareMetalClient
}

func (s *PolicySync) toStringArray(vals interface{}) []string {
	arr := vals.([]interface{})
	result := []string{}
	for _, val := range arr {
		result = append(result, val.(string))
	}
	return result
}

func (s *PolicySync) Create() (res *baremtlsdk.Resource, e error) {
	name := s.d.Get("name").(string)
	description := s.d.Get("description").(string)
	statements := s.toStringArray(s.d.Get("statements"))
	res, e = s.client.CreatePolicy(name, description, statements)
	return
}

func (s *PolicySync) Get() (res *baremtlsdk.Resource, e error) {
	res, e = s.client.GetPolicy(s.d.Id())
	return
}

func (s *PolicySync) Update() {
	description := d.Get("description").(string)
	statements := s.toStringArray(s.d.Get("statements"))
	res, e = s.client.UpdatePolicy(s.d.Id(), description, statements)
	return
}

func (s *PolicySync) SetData(res *baremtlsdk.Resource) {
	s.d.Set("statements", res.Statements)
	setResourceData(s.d, res)
}

func (s *PolicySync) Delete() (e error) {
	return s.client.DeletePolicy(s.d.Id())
}
