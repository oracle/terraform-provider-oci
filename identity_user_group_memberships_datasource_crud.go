// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type UserGroupMembershipDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListUserGroupMemberships
}

func (s *UserGroupMembershipDatasourceCrud) Get() (e error) {
	s.Res, e = s.Client.ListUserGroupMemberships(nil)
	return
}

func (s *UserGroupMembershipDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Memberships {
			res := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"id":             v.ID,
				"user_id":        v.UserID,
				"group_id":       v.GroupID,
				"inactive_state": v.InactiveStatus,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		if err := s.D.Set("memberships", resources); err != nil {
			panic(err)
		}
	}
	return
}
