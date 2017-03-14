// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package database

import (
	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type DBNodeDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.DBNode
}

func (s *DBNodeDatasourceCrud) Get() (e error) {
	id := s.D.Get("db_node_id").(string)
	s.Res, e = s.Client.GetDBNode(id)
	return
}

func (s *DBNodeDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(s.Res.ID)
		s.D.Set("db_system_id", s.Res.DBSystemID)
		s.D.Set("hostname", s.Res.Hostname)
		s.D.Set("id", s.Res.ID)
		s.D.Set("state", s.Res.State)
		s.D.Set("time_created", s.Res.TimeCreated.String())
		s.D.Set("vnic_id", s.Res.VnicID)
	}
	return
}
