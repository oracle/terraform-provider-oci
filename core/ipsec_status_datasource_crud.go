// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/MustWin/baremetal-sdk-go"

	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
)

type IPSecStatusDatasourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.IPSecConnectionDeviceStatus
}

func (s *IPSecStatusDatasourceCrud) Get() (e error) {
	ipsecID := s.D.Get("ipsec_id").(string)
	s.Resource, e = s.Client.GetIPSecConnectionDeviceStatus(ipsecID)
	return
}

func (s *IPSecStatusDatasourceCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(s.Resource.ID)
		s.D.Set("compartment_id", s.Resource.CompartmentID)
		s.D.Set("id", s.Resource.ID)
		s.D.Set("time_created", s.Resource.TimeCreated)

		tunnels := []map[string]interface{}{}

		for _, val := range s.Resource.Tunnels {
			tunnel := map[string]interface{}{
				"ip_address":         val.IPAddress,
				"state":              val.State,
				"time_created":       val.TimeCreated.String(),
				"time_state_modifed": val.TimeStateModified.String(),
			}

			tunnels = append(tunnels, tunnel)
		}

		s.D.Set("tunnels", tunnels)

	}
}
