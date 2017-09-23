// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DBHomeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBHome,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_home_id": {
				Type:     schema.TypeString,
				Required: true,
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

func readDBHome(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DBHomeDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type DBHomeDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.DBHome
}

func (s *DBHomeDatasourceCrud) Get() (e error) {
	id := s.D.Get("db_home_id").(string)
	res, e := s.Client.GetDBHome(id)
	if e == nil {
		s.Res = res
	}
	return
}

func (s *DBHomeDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(s.Res.ID)
		s.D.Set("compartment_id", s.Res.CompartmentID)
		s.D.Set("db_system_id", s.Res.DBSystemID)
		s.D.Set("db_version", s.Res.DBVersion)
		s.D.Set("display_name", s.Res.DisplayName)
		s.D.Set("id", s.Res.ID)
		s.D.Set("state", s.Res.State)
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}
	return
}
