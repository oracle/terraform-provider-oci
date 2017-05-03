// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DBNodeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBNode,
		Schema: map[string]*schema.Schema{
			"db_node_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
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
			"vnic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readDBNode(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBNodeDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type DBNodeDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.DBNode
}

func (s *DBNodeDatasourceCrud) Get() (e error) {
	id := s.D.Get("db_node_id").(string)
	s.Res, e = s.Client.GetDBNode(id)
	return
}

func (s *DBNodeDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(s.Res.ID)
		s.D.Set("db_system_id", s.Res.DBSystemID)
		s.D.Set("hostname", s.Res.Hostname)
		s.D.Set("id", s.Res.ID)
		s.D.Set("state", s.Res.State)
		s.D.Set("time_created", s.Res.TimeCreated.String())
		s.D.Set("vnic_id", s.Res.VnicID)
	}
	return
}
