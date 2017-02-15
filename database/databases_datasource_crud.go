// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package database

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/options"
)

type DatabasesDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDatabases
}

func (s *DatabasesDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	dbHomeID := s.D.Get("db_home_id").(string)
	limit := uint64(s.D.Get("limit").(int))

	opts := &baremetal.PageListOptions{}
	options.SetPageOptions(s.D, opts)

	s.Res = &baremetal.ListDatabases{}

	for {
		var list *baremetal.ListDatabases
		if list, e = s.Client.ListDatabases(
			compartmentID, dbHomeID, limit, opts,
		); e != nil {
			break
		}

		s.Res.Databases = append(s.Res.Databases, list.Databases...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, opts); !hasNextPage {
			break
		}
	}

	return
}

func (s *DatabasesDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Databases {
			res := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"db_home_id":     v.DBHomeID,
				"db_name":        v.DBName,
				"db_unique_name": v.DBUniqueName,
				"id":             v.ID,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("databases", resources)
	}
	return
}
