// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package database

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/options"
	"github.com/hashicorp/terraform/helper/schema"
)

type DBVersionDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ListDBVersions
}

func (s *DBVersionDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	limit := uint64(s.D.Get("limit").(int))

	opts := &baremetal.PageListOptions{}
	options.SetPageOptions(s.D, opts)

	s.Res = &baremetal.ListDBVersions{}

	for {
		var list *baremetal.ListDBVersions
		if list, e = s.Client.ListDBVersions(compartmentID, limit, opts); e != nil {
			break
		}

		s.Res.DBVersions = append(s.Res.DBVersions, list.DBVersions...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, opts); !hasNextPage {
			break
		}
	}

	return
}

func (s *DBVersionDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.DBVersions {
			res := map[string]interface{}{
				"version": v.Version,
			}
			resources = append(resources, res)
		}
		s.D.Set("db_versions", resources)
	}
	return
}
