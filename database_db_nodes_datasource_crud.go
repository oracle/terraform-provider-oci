// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/oracle/terraform-provider-baremetal/options"
)

type DBNodesDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDBNodes
}

func (s *DBNodesDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	dbSystemID := s.D.Get("db_system_id").(string)
	limit := uint64(s.D.Get("limit").(int))

	opts := &baremetal.PageListOptions{}
	options.SetPageOptions(s.D, opts)

	s.Res = &baremetal.ListDBNodes{}

	for {
		var list *baremetal.ListDBNodes
		if list, e = s.Client.ListDBNodes(
			compartmentID, dbSystemID, limit, opts,
		); e != nil {
			break
		}

		s.Res.DBNodes = append(s.Res.DBNodes, list.DBNodes...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, opts); !hasNextPage {
			break
		}
	}

	return
}

func (s *DBNodesDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.DBNodes {
			res := map[string]interface{}{
				"db_system_id": v.DBSystemID,
				"hostname":     v.Hostname,
				"id":           v.ID,
				"state":        v.State,
				"time_created": v.TimeCreated.String(),
				"vnic_id":      v.VnicID,
			}
			resources = append(resources, res)
		}
		s.D.Set("db_nodes", resources)
	}
	return
}
