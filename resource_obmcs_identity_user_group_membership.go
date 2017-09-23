// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func UserGroupMembershipResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createUserGroupMembership,
		Read:     readUserGroupMembership,
		Delete:   deleteUserGroupMembership,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createUserGroupMembership(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &UserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readUserGroupMembership(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &UserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func deleteUserGroupMembership(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &UserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}

type UserGroupMembershipResourceCrud struct {
	*crud.IdentitySync
	crud.BaseCrud
	Res *baremetal.UserGroupMembership
}

func (s *UserGroupMembershipResourceCrud) ID() string {
	return s.Res.ID
}

func (s *UserGroupMembershipResourceCrud) State() string {
	return s.Res.State
}

func (s *UserGroupMembershipResourceCrud) Create() (e error) {
	userID := s.D.Get("user_id").(string)
	groupID := s.D.Get("group_id").(string)
	s.Res, e = s.Client.AddUserToGroup(userID, groupID, nil)
	return
}

func (s *UserGroupMembershipResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *UserGroupMembershipResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceActive}
}

func (s *UserGroupMembershipResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDeleting}
}

func (s *UserGroupMembershipResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func (s *UserGroupMembershipResourceCrud) Get() (e error) {
	res, e := s.Client.GetUserGroupMembership(s.D.Id())
	if e == nil {
		s.Res = res
	}
	return
}

func (s *UserGroupMembershipResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("group_id", s.Res.GroupID)
	s.D.Set("user_id", s.Res.UserID)
	s.D.Set("inactive_state", s.Res.InactiveStatus)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *UserGroupMembershipResourceCrud) Delete() (e error) {
	return s.Client.DeleteUserGroupMembership(s.D.Id(), nil)
}
