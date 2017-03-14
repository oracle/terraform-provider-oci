// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package database

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/oracle/terraform-provider-baremetal/options"
)

type DBHomesDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDBHomes
}

func (s *DBHomesDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	dbSystemID := s.D.Get("db_system_id").(string)
	limit := uint64(s.D.Get("limit").(int))

	opts := &baremetal.PageListOptions{}
	options.SetPageOptions(s.D, opts)

	s.Res = &baremetal.ListDBHomes{}

	for {
		var list *baremetal.ListDBHomes
		if list, e = s.Client.ListDBHomes(
			compartmentID, dbSystemID, limit, opts,
		); e != nil {
			break
		}

		s.Res.DBHomes = append(s.Res.DBHomes, list.DBHomes...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, opts); !hasNextPage {
			break
		}
	}

	return
}

func (s *DBHomesDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.DBHomes {
			res := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"db_system_id":   v.DBSystemID,
				"display_name":   v.DisplayName,
				"id":             v.ID,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("db_homes", resources)
	}
	return
}
