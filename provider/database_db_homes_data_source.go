// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DBHomesDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBHomes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"db_homes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DBHomeResource(),
			},
		},
	}
}

func readDBHomes(d *schema.ResourceData, m interface{}) (e error) {
	sync := &DBHomesDatasourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

type DBHomesDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDBHomes
}

func (s *DBHomesDatasourceCrud) Get() (e error) {
	opts := &baremetal.ListDBHomesOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)
	compartmentID := s.D.Get("compartment_id").(string)
	dbSystemID := s.D.Get("db_system_id").(string)

	s.Res = &baremetal.ListDBHomes{
		DBHomes: []baremetal.DBHome{},
	}

	for {
		var list *baremetal.ListDBHomes
		if list, e = s.Client.ListDBHomes(compartmentID, dbSystemID, opts); e != nil {
			break
		}

		s.Res.DBHomes = append(s.Res.DBHomes, list.DBHomes...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}
	return
}

func (s *DBHomesDatasourceCrud) SetData() {
	if s.Res == nil {
		return
	}
	s.D.SetId(time.Now().UTC().String())
	resources := []map[string]interface{}{}
	for _, r := range s.Res.DBHomes {
		dbHome := map[string]interface{}{
			"compartment_id": r.CompartmentID,
			"db_system_id":   r.DBSystemID,
			"db_version":     r.DBVersion,
			"display_name":   r.DisplayName,
			"id":             r.ID,
			"last_patch_history_entry_id": r.LastPatchHistoryEntryID,
			"state":        r.State,
			"time_created": r.TimeCreated.String(),
		}
		resources = append(resources, dbHome)
	}
	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}
	if err := s.D.Set("db_homes", resources); err != nil {
		panic(err)
	}
	return
}
