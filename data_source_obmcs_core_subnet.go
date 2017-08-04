// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/oracle/terraform-provider-baremetal/options"
)

func SubnetDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readSubnets,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"subnets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     resourceCoreSubnets(),
			},
		},
	}
}

func resourceCoreSubnets() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cidr_block": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_table_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_list_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"prohibit_public_ip_on_vnic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_router_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_router_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSubnets(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &SubnetDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}

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
				"prohibit_public_ip_on_vnic": v.ProhibitPublicIpOnVnic,
				"state":              v.State,
				"time_created":       v.TimeCreated.String(),
				"virtual_router_ip":  v.VirtualRouterIP,
				"virtual_router_mac": v.VirtualRouterMac,
			}
			resources = append(resources, res)
		}
		s.D.Set("subnets", resources)
	}
	return
}
