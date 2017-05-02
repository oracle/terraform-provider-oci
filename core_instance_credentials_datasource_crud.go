// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/oracle/terraform-provider-baremetal/crud"
	"time"
)

type InstanceCredentialsDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.InstanceCredentials
}

func (s *InstanceCredentialsDatasourceCrud) Get() (e error) {
	instanceId := s.D.Get("instance_id").(string)
	s.Res, e = s.Client.GetWindowsInstanceInitialCredentials(instanceId)
	return
}

func (s *InstanceCredentialsDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		s.D.Set("username", s.Res.Username)
		s.D.Set("password", s.Res.Password)
	}
	return
}
