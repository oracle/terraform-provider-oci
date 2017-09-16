// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
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
	client := m.(*baremetal.Client)
	sync := &CompartmentDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type CompartmentDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListCompartments
}

func (s *CompartmentDatasourceCrud) Get() (e error) {
	s.Res, e = s.Client.ListCompartments(nil)
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
