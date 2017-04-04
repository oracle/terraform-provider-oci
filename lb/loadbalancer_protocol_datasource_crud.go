// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package lb

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type ProtocolDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListLoadBalancerProtocols
}

func (s *ProtocolDatasourceCrud) Get() (e error) {
	cID := s.D.Get("compartment_id").(string)
	s.Res, e = s.Client.ListLoadBalancerProtocols(cID, nil)
	return
}

func (s *ProtocolDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Res.LoadBalancerProtocols {
			res := map[string]interface{}{
				"name": v.Name,
			}
			resources = append(resources, res)

		}
		s.D.Set("protocols", resources)
	}
	return
}
