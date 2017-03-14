// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
)

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
	s.Res, e = s.Client.GetUserGroupMembership(s.D.Id())
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
