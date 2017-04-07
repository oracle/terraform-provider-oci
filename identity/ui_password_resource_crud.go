// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type UIPasswordResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.UIPassword
}

func (s *UIPasswordResourceCrud) ID() string {
	return s.D.Get("version").(string)
}

func (s *UIPasswordResourceCrud) Create() (e error) {
	userID := s.D.Get("user_id").(string)
	s.Res, e = s.Client.CreateOrResetUIPassword(userID, nil)
	return
}

func (s *UIPasswordResourceCrud) SetData() {
	s.D.Set("inactive_status", s.Res.InactiveStatus)
	s.D.Set("state", s.Res.State)
	s.D.Set("password", s.Res.Password)
	s.D.Set("time_created", s.Res.TimeCreated.String())
	s.D.Set("user_id", s.Res.UserID)
}
