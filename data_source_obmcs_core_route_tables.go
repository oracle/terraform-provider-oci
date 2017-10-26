// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func RouteTableDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readRouteTables,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"route_tables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     RouteTableResource(),
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func readRouteTables(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	reader := &RouteTableDatasourceCrud{}
	reader.D = d
	reader.Client = client.client

	return crud.ReadResource(reader)
}

type RouteTableDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListRouteTables
}

func (s *RouteTableDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)
	opts := &baremetal.ListOptions{}
	options.SetListOptions(s.D, opts)

	s.Res = &baremetal.ListRouteTables{RouteTables: []baremetal.RouteTable{}}

	for {
		var list *baremetal.ListRouteTables
		if list, e = s.Client.ListRouteTables(compartmentID, vcnID, opts); e != nil {
			break
		}

		s.Res.RouteTables = append(s.Res.RouteTables, list.RouteTables...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *RouteTableDatasourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(time.Now().UTC().String())

	resources := []map[string]interface{}{}
	for _, v := range s.Res.RouteTables {

		rules := []map[string]interface{}{}
		for _, val := range v.RouteRules {
			rule := map[string]interface{}{
				"cidr_block":        val.CidrBlock,
				"network_entity_id": val.NetworkEntityID,
			}
			rules = append(rules, rule)
		}

		res := map[string]interface{}{
			"compartment_id": v.CompartmentID,
			"display_name":   v.DisplayName,
			"id":             v.ID,
			"route_rules":    rules,
			"state":          v.State,
			"time_created":   v.TimeCreated.String(),
			"time_modified":  v.TimeModified.String(),
			"vcn_id":         s.D.Get("vcn_id").(string), // todo: get this off the route table result
		}
		resources = append(resources, res)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("route_tables", resources); err != nil {
		panic(err)
	}

	return
}
