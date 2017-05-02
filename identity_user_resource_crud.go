// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type UserResourceCrud struct {
	*crud.IdentitySync
	crud.BaseCrud
	Res *baremetal.User
}

func (s *UserResourceCrud) ID() string {
	return s.Res.ID
}

func (s *UserResourceCrud) State() string {
	return s.Res.State
}

func (s *UserResourceCrud) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.CreateUser(name, description, nil)
	return
}

func (s *UserResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *UserResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceActive}
}

func (s *UserResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDeleting}
}

func (s *UserResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func (s *UserResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetUser(s.D.Id())
	return
}

func (s *UserResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateIdentityOptions{}
	if description, ok := s.D.GetOk("description"); ok {
		opts.Description = description.(string)
	}

	s.Res, e = s.Client.UpdateUser(s.D.Id(), opts)
	return
}

func (s *UserResourceCrud) SetData() {
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *UserResourceCrud) Delete() (e error) {
	return s.Client.DeleteUser(s.D.Id(), nil)
}
