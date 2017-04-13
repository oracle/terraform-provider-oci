// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type PolicyDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListPolicies
}

func (s *PolicyDatasourceCrud) Get() (e error) {
	compartment_id := s.D.Get("compartment_id").(string)

	s.Res, e = s.Client.ListPolicies(compartment_id, nil)
	return
}

func (s *PolicyDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Policies {
			res := map[string]interface{}{
				"id":             v.ID,
				"compartment_id": v.CompartmentID,
				"name":           v.Name,
				"statements":     v.Statements,
				"description":    v.Description,
				"time_created":   v.TimeCreated.String(),
				"state":          v.State,
				"inactive_state": v.InactiveStatus,
				"version_date":   v.VersionDate.String(),
			}
			resources = append(resources, res)
		}
		if err := s.D.Set("policies", resources); err != nil {
			panic(err)
		}
	}
	return
}
