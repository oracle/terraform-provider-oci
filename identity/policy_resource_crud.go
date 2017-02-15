// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
)

type PolicyResourceCrud struct {
	*crud.IdentitySync
	crud.BaseCrud
	Res *baremetal.Policy
}

func (s *PolicyResourceCrud) ID() string {
	return s.Res.ID
}

func (s *PolicyResourceCrud) State() string {
	return s.Res.State
}

func (s *PolicyResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *PolicyResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceActive}
}

func (s *PolicyResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDeleting}
}

func (s *PolicyResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func (s *PolicyResourceCrud) toStringArray(vals interface{}) []string {
	arr := vals.([]interface{})
	result := []string{}
	for _, val := range arr {
		result = append(result, val.(string))
	}
	return result
}

func (s *PolicyResourceCrud) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	statements := s.toStringArray(s.D.Get("statements"))

	s.Res, e = s.Client.CreatePolicy(name, description, compartmentID, statements, nil)
	return
}

func (s *PolicyResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetPolicy(s.D.Id())
	return
}

func (s *PolicyResourceCrud) Update() (e error) {
	opts := &baremetal.UpdatePolicyOptions{}
	if description, ok := s.D.GetOk("description"); ok {
		opts.Description = description.(string)
	}

	if rawStatements, ok := s.D.GetOk("statements"); ok {
		statements := s.toStringArray(rawStatements)
		opts.Statements = statements
	}

	s.Res, e = s.Client.UpdatePolicy(s.D.Id(), opts)
	return
}

func (s *PolicyResourceCrud) SetData() {
	s.D.Set("statements", s.Res.Statements)
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *PolicyResourceCrud) Delete() (e error) {
	return s.Client.DeletePolicy(s.D.Id(), nil)
}
