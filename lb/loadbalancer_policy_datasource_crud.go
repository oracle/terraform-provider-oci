// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package lb

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type PoliciesDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListLoadBalancerPolicies
}

func (s *PoliciesDatasourceCrud) Get() (e error) {
	cID := s.D.Get("compartment_id").(string)
	s.Res, e = s.Client.ListLoadBalancerPolicies(cID, nil)
	return
}

func (s *PoliciesDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Res.LoadBalancerPolicies {
			res := map[string]interface{}{
				"name": v.Name,
			}
			resources = append(resources, res)

		}
		s.D.Set("policies", resources)
	}
	return
}
