// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
	"github.com/oracle/terraform-provider-oci/options"
)

func CompartmentDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readCompartments,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     CompartmentResource(),
			},
		},
	}
}

func readCompartments(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &CompartmentDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type CompartmentDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListCompartments
}

func (s *CompartmentDatasourceCrud) Get() (e error) {
	opts := &baremetal.ListOptions{}
	options.SetListOptions(s.D, opts)

	s.Res = &baremetal.ListCompartments{Compartments: []baremetal.Compartment{}}

	for {
		var list *baremetal.ListCompartments
		if list, e = s.Client.ListCompartments(opts); e != nil {
			break
		}

		s.Res.Compartments = append(s.Res.Compartments, list.Compartments...)

		if hasNexPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNexPage {
			break
		}
	}

	return
}

func (s *CompartmentDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Compartments {
			res := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"description":    v.Description,
				"id":             v.ID,
				"inactive_state": v.InactiveStatus,
				"name":           v.Name,
				"state":          v.State,
				"time_created":   v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		if err := s.D.Set("compartments", resources); err != nil {
			panic(err)
		}
	}
	return
}
