// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DatabaseDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabase,
		Schema: map[string]*schema.Schema{
			"character_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_home_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_unique_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_workload": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ncharacter_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pdb_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readDatabase(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &DatabaseDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type DatabaseDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.Database
}

func (s *DatabaseDatasourceCrud) Get() (e error) {
	id := s.D.Get("database_id").(string)
	res, e := s.Client.GetDatabase(id)
	if e == nil {
		s.Res = res
	}
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
		s.D.Set("character_set", s.Res.CharacterSet)
		s.D.Set("ncharacter_set", s.Res.NcharacterSet)
		s.D.Set("pdb_name", s.Res.PDBName)
		s.D.Set("db_workload", s.Res.DBWorkload)
		s.D.Set("lifecycle_details", s.Res.LifecycleDetails)
	}
	return
}
