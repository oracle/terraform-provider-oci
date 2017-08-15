// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/terraform-provider-baremetal/options"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DrgDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDrgs,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
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
			"drgs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DrgResource(),
			},
		},
	}
}

func readDrgs(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DrgDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type DrgDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListDrgs
}

func (s *DrgDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetListOptions(s.D, opts)

	s.Res = &baremetal.ListDrgs{Drgs: []baremetal.Drg{}}

	for {
		var list *baremetal.ListDrgs
		if list, e = s.Client.ListDrgs(compartmentID, opts); e != nil {
			break
		}

		s.Res.Drgs = append(s.Res.Drgs, list.Drgs...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *DrgDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]string{}
		for _, v := range s.Res.Drgs {
			res := map[string]string{
				"compartment_id": v.CompartmentID,
				"display_name":   v.DisplayName,
				"id":             v.ID,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("drgs", resources)
	}
	return
}
