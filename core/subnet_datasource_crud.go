// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/options"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type SubnetDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListSubnets
}

func (s *SubnetDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetListOptions(s.D, opts)

	s.Res = &baremetal.ListSubnets{Subnets: []baremetal.Subnet{}}

	for {
		var list *baremetal.ListSubnets
		if list, e = s.Client.ListSubnets(compartmentID, vcnID, opts); e != nil {
			break
		}

		s.Res.Subnets = append(s.Res.Subnets, list.Subnets...)

		if hasNexPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNexPage {
			break
		}
	}

	return
}

func (s *SubnetDatasourceCrud) SetData() {
	if s.Res != nil {

		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Subnets {
			res := map[string]interface{}{
				"availability_domain": v.AvailabilityDomain,
				"cidr_block":          v.CIDRBlock,
				"compartment_id":      v.CompartmentID,
				"route_table_id":      v.RouteTableID,
				"vcn_id":              v.VcnID,
				"security_list_ids":   v.SecurityListIDs,
				"display_name":        v.DisplayName,
				"id":                  v.ID,
				"state":               v.State,
				"time_created":        v.TimeCreated.String(),
				"virtual_router_id":   v.VirtualRouterID,
				"virtual_router_mac":  v.VirtualRouterMac,
			}
			resources = append(resources, res)
		}
		s.D.Set("subnets", resources)
	}
	return
}
