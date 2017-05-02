// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/options"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type VirtualNetworkDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListVirtualNetworks
}

func (s *VirtualNetworkDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetListOptions(s.D, opts)

	s.Res = &baremetal.ListVirtualNetworks{
		VirtualNetworks: []baremetal.VirtualNetwork{},
	}

	for {
		var list *baremetal.ListVirtualNetworks
		if list, e = s.Client.ListVirtualNetworks(compartmentID, opts); e != nil {
			break
		}

		s.Res.VirtualNetworks = append(s.Res.VirtualNetworks, list.VirtualNetworks...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *VirtualNetworkDatasourceCrud) SetData() {
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]string{}
		for _, v := range s.Res.VirtualNetworks {
			res := map[string]string{
				"cidr_block":               v.CidrBlock,
				"compartment_id":           v.CompartmentID,
				"default_route_table_id":   v.DefaultRouteTableID,
				"default_security_list_id": v.DefaultSecurityListID,
				"display_name":             v.DisplayName,
				"id":                       v.ID,
				"state":                    v.State,
				"time_created":             v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("virtual_networks", resources)
	}
	return
}
