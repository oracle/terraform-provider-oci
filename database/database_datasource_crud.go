// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package database

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type DatabaseDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.Database
}

func (s *DatabaseDatasourceCrud) Get() (e error) {
	id := s.D.Get("database_id").(string)
	s.Res, e = s.Client.GetDatabase(id)
	return
}

func (s *DatabaseDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(s.Res.ID)
		s.D.Set("compartment_id", s.Res.CompartmentID)
		s.D.Set("db_home_id", s.Res.DBHomeID)
		s.D.Set("db_name", s.Res.DBName)
		s.D.Set("db_unique_name", s.Res.DBUniqueName)
		s.D.Set("id", s.Res.ID)
		s.D.Set("state", s.Res.State)
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}
	return
}
