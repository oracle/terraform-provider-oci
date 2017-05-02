// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type AvailabilityDomainDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListAvailabilityDomains
}

func (s *AvailabilityDomainDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	s.Res, e = s.Client.ListAvailabilityDomains(compartmentID)
	return
}

func (s *AvailabilityDomainDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.AvailabilityDomains {
			res := map[string]interface{}{
				"name":           v.Name,
				"compartment_id": v.CompartmentID,
			}
			resources = append(resources, res)
		}
		s.D.Set("availability_domains", resources)
	}
	return
}
