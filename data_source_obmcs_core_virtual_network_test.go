// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"



)

type ResourceCoreVirtualNetworksTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVirtualNetworksTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
resource "baremetal_core_virtual_network" "t" {
	cidr_block = "10.0.0.0/16"
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
}
data "baremetal_core_virtual_networks" "t" {
	compartment_id = "${baremetal_core_virtual_network.t.compartment_id}"
	limit = 1
}
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_virtual_networks.t"
}

func (s *ResourceCoreVirtualNetworksTestSuite) TestReadVirtualNetworks() {
	opts := &baremetal.ListOptions{}
	opts.Limit = 1
	opts.Page = "page"

	s.Client.On(
		"ListVirtualNetworks",
		"compartment_id",
		opts,
	).Return(
		&baremetal.ListVirtualNetworks{
			VirtualNetworks: []baremetal.VirtualNetwork{
				{
					CidrBlock:             "10.0.0.0/16",
					CompartmentID:         "compartment_id",
					DefaultRouteTableID:   "default_route_table_id",
					DefaultSecurityListID: "default_security_list_id",
					DisplayName:           "display_name",
					ID:                    "id1",
					State:                 baremetal.ResourceAttached,
					TimeCreated:           baremetal.Time{Time: time.Now()},
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.#"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVirtualNetworks", "compartment_id", opts)
}

func (s *ResourceCoreVirtualNetworksTestSuite) TestReadVirtualNetworksWithPaging() {
	if IsAccTest() {
		s.T().Skip()
	}
	opts := &baremetal.ListOptions{}
	opts.Limit = 1
	opts.Page = "page"

	res := &baremetal.ListVirtualNetworks{}
	res.NextPage = "nextpage"
	res.VirtualNetworks = []baremetal.VirtualNetwork{
		{
			CidrBlock:             "cidr_block",
			CompartmentID:         "compartment_id",
			DefaultRouteTableID:   "default_route_table_id",
			DefaultSecurityListID: "default_security_list_id",
			DisplayName:           "display_name",
			ID:                    "id1",
			State:                 baremetal.ResourceAttached,
			TimeCreated:           baremetal.Time{Time: time.Now()},
		},
		{
			CidrBlock:             "cidr_block",
			CompartmentID:         "compartment_id",
			DefaultRouteTableID:   "default_route_table_id",
			DefaultSecurityListID: "default_security_list_id",
			DisplayName:           "display_name",
			ID:                    "id2",
			State:                 baremetal.ResourceAttached,
			TimeCreated:           baremetal.Time{Time: time.Now()},
		},
	}

	s.Client.On(
		"ListVirtualNetworks",
		"compartment_id",
		opts,
	).Return(res, nil)

	opts2 := &baremetal.ListOptions{}
	opts2.Limit = 1
	opts2.Page = "nextpage"

	s.Client.On(
		"ListVirtualNetworks",
		"compartment_id",
		opts2,
	).Return(
		&baremetal.ListVirtualNetworks{
			VirtualNetworks: []baremetal.VirtualNetwork{
				{
					CidrBlock:             "cidr_block",
					CompartmentID:         "compartment_id",
					DefaultRouteTableID:   "default_route_table_id",
					DefaultSecurityListID: "default_security_list_id",
					DisplayName:           "display_name",
					ID:                    "id3",
					State:                 baremetal.ResourceAttached,
					TimeCreated:           baremetal.Time{Time: time.Now()},
				},
				{
					CidrBlock:             "cidr_block",
					CompartmentID:         "compartment_id",
					DefaultRouteTableID:   "default_route_table_id",
					DefaultSecurityListID: "default_security_list_id",
					DisplayName:           "display_name",
					ID:                    "id4",
					State:                 baremetal.ResourceAttached,
					TimeCreated:           baremetal.Time{Time: time.Now()},
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.cidr_block", "cidr_block"),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVirtualNetworks", "compartment_id", opts2)
}

func TestResourceCoreVirtualNetworksTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVirtualNetworksTestSuite))
}
