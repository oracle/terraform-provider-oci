// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type RouteTableResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.RouteTable
}

func (s *RouteTableResourceCrud) ID() string {
	return s.Res.ID
}

func (s *RouteTableResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *RouteTableResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *RouteTableResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *RouteTableResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *RouteTableResourceCrud) State() string {
	return s.Res.State
}

func (s *RouteTableResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = s.D.Get("display_name").(string)

	s.Res, e = s.Client.CreateRouteTable(compartmentID, vcnID, s.buildRouteRules(), opts)

	return
}

func (s *RouteTableResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetRouteTable(s.D.Id())
	return
}

func (s *RouteTableResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateRouteTableOptions{}
	opts.RouteRules = s.buildRouteRules()

	s.Res, e = s.Client.UpdateRouteTable(s.D.Id(), opts)
	return
}

func (s *RouteTableResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)

	rules := []map[string]interface{}{}
	for _, val := range s.Res.RouteRules {
		rule := map[string]interface{}{
			"cidr_block":          val.CidrBlock,
			"network_entity_id":   val.NetworkEntityID,
		}
		rules = append(rules, rule)
	}
	s.D.Set("route_rules", rules)

	s.D.Set("time_modified", s.Res.TimeModified.String())
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *RouteTableResourceCrud) Delete() (e error) {
	return s.Client.DeleteRouteTable(s.D.Id(), nil)
}

func (s *RouteTableResourceCrud) buildRouteRules() (routeRules []baremetal.RouteRule) {
	routeRules = []baremetal.RouteRule{}
	for _, val := range s.D.Get("route_rules").([]interface{}) {
		data := val.(map[string]interface{})
		routeRule := baremetal.RouteRule{
			CidrBlock:         data["cidr_block"].(string),
			NetworkEntityID:   data["network_entity_id"].(string),
		}
		routeRules = append(routeRules, routeRule)
	}
	return
}
