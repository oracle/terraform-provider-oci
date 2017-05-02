// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/oracle/terraform-provider-baremetal/options"
)

type DBSystemShapeDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDBSystemShapes
}

func (s *DBSystemShapeDatasourceCrud) Get() (e error) {
	availabilityDomain := s.D.Get("availability_domain").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	limit := uint64(s.D.Get("limit").(int))

	opts := &baremetal.PageListOptions{}
	options.SetPageOptions(s.D, opts)

	s.Res = &baremetal.ListDBSystemShapes{}

	for {
		var list *baremetal.ListDBSystemShapes
		if list, e = s.Client.ListDBSystemShapes(
			availabilityDomain, compartmentID, limit, opts,
		); e != nil {
			break
		}

		s.Res.DBSystemShapes = append(s.Res.DBSystemShapes, list.DBSystemShapes...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, opts); !hasNextPage {
			break
		}
	}

	return
}

func (s *DBSystemShapeDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.DBSystemShapes {
			res := map[string]interface{}{
				"available_core_count": v.AvailableCoreCount,
				"name":                 v.Name,
				"shape":                v.Shape,
			}
			resources = append(resources, res)
		}
		s.D.Set("db_system_shapes", resources)
	}
	return
}
