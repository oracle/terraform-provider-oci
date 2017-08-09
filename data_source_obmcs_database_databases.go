// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/terraform-provider-baremetal/options"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DatabasesDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabases,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_home_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DatabaseDatasource(),
			},
		},
	}
}

func readDatabases(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DatabasesDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

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
