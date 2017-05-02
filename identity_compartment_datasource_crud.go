// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

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
