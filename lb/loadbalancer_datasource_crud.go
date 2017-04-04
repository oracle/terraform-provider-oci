// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package lb

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type LoadBalancerDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListLoadBalancers
}

func (s *LoadBalancerDatasourceCrud) Get() (e error) {
	cID := s.D.Get("compartment_id").(string)
	s.Res, e = s.Client.ListLoadBalancers(cID, nil)
	return
}

func (s *LoadBalancerDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Res.LoadBalancers {
			res := map[string]interface{}{
				"id":             v.ID,
				"compartment_id": v.CompartmentID,
				"display_name":   v.DisplayName,
				"ip_addresses":   v.IPAddresses,
				"shape":          v.Shape,
				"state":          v.State,
				"subnet_ids":     v.SubnetIDs,
				"time_created":   v.TimeCreated,
			}
			resources = append(resources, res)

		}
		s.D.Set("load_balancers", resources)
	}
	return
}
