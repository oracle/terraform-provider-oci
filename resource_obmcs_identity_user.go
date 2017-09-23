// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

// ResourceIdentityUser exposes a IdentityUser Resource
func UserResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createUser,
		Read:     readUser,
		Update:   updateUser,
		Delete:   deleteUser,
		Schema:   baseIdentitySchemaWithID,
	}
}

func createUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteUser(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}

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
	res, e := s.Client.GetUser(s.D.Id())
	if e == nil {
		s.Res = res
	}
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
