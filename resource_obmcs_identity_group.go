// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"time"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

// ResourceIdentityGroup exposes an IdentityGroup Resource
func GroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createGroup,
		Read:     readGroup,
		Update:   updateGroup,
		Delete:   deleteGroup,
		Schema:   baseIdentitySchemaWithID,
	}
}

func createGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &GroupSync{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &GroupSync{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &GroupSync{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &GroupSync{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}

type GroupSync struct {
	*crud.IdentitySync
	crud.BaseCrud
	Res *baremetal.Group
}

func (s *GroupSync) ID() string {
	return s.Res.ID
}

func (s *GroupSync) State() string {
	return s.Res.State
}

func (s *GroupSync) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *GroupSync) CreatedTarget() []string {
	return []string{baremetal.ResourceActive}
}

func (s *GroupSync) DeletedPending() []string {
	return []string{baremetal.ResourceDeleting}
}

func (s *GroupSync) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func (s *GroupSync) ExtraWaitPostCreateDelete() time.Duration {
	return time.Duration(20 * time.Second)
}

func (s *GroupSync) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.CreateGroup(name, description, nil)
	return
}

func (s *GroupSync) Get() (e error) {
	s.Res, e = s.Client.GetGroup(s.D.Id())
	return
}

func (s *GroupSync) Update() (e error) {
	opts := &baremetal.UpdateIdentityOptions{}
	if description, ok := s.D.GetOk("description"); ok {
		opts.Description = description.(string)
	}

	s.Res, e = s.Client.UpdateGroup(s.D.Id(), opts)
	return
}

func (s *GroupSync) SetData() {
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *GroupSync) Delete() (e error) {
	return s.Client.DeleteGroup(s.D.Id(), nil)
}
