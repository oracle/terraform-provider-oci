// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DBNodesDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBNodes,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
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
			"db_nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DBNodeDatasource(),
			},
		},
	}
}

func readDBNodes(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DBNodesDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type DBNodesDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDBNodes
}

func (s *DBNodesDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	dbSystemID := s.D.Get("db_system_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetPageOptions(s.D, &opts.PageListOptions)
	options.SetLimitOptions(s.D, &opts.LimitListOptions)

	s.Res = &baremetal.ListDBNodes{}

	for {
		var list *baremetal.ListDBNodes
		if list, e = s.Client.ListDBNodes(
			compartmentID, dbSystemID, opts,
		); e != nil {
			break
		}

		s.Res.DBNodes = append(s.Res.DBNodes, list.DBNodes...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
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
