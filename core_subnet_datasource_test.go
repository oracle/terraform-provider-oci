// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreSubnetsTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreSubnetsTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_subnets" "s" {
      compartment_id = "compartmentid"
      vcn_id = "vcnid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_subnets.s"

}

func (s *ResourceCoreSubnetsTestSuite) TestResourceListSubnets() {
	opts := &baremetal.ListOptions{}
	s.Client.On(
		"ListSubnets",
		"compartmentid",
		"vcnid",
		opts,
	).Return(
		&baremetal.ListSubnets{
			Subnets: []baremetal.Subnet{
				{
					AvailabilityDomain: "availabilitydomainid",
					CIDRBlock:          "10.10.10.0/24",
					CompartmentID:      "compartmentid",
					DisplayName:        "display_name",
					ID:                 "id1",
					RouteTableID:       "routetableid",
					SecurityListIDs: []string{
						"slid1",
						"slid2",
					},
					State: baremetal.ResourceAvailable,
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
					VcnID:            "vcnid",
					VirtualRouterIP:  "virtualrouterip",
					VirtualRouterMac: "virtualroutermac",
				},
				{
					AvailabilityDomain: "availabilitydomainid",
					CIDRBlock:          "10.10.11.0/24",
					CompartmentID:      "compartmentid",
					DisplayName:        "display_name",
					ID:                 "id2",
					RouteTableID:       "routetableid",
					SecurityListIDs: []string{
						"slid1",
						"slid2",
					},
					State: baremetal.ResourceAvailable,
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
					VcnID:            "vcnid",
					VirtualRouterIP:  "virtualrouterip",
					VirtualRouterMac: "virtualroutermac",
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
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", "vcnid"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.availability_domain", "availabilitydomainid"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListSubnets", "compartmentid", "vcnid", opts)

}

func (s *ResourceCoreSubnetsTestSuite) TestResourceListSubnetsWithPagination() {
	opts := &baremetal.ListOptions{}
	res := &baremetal.ListSubnets{}
	res.NextPage = "nextpage"
	res.Subnets = []baremetal.Subnet{
		{
			AvailabilityDomain: "availabilitydomainid",
			CIDRBlock:          "10.10.10.0/24",
			CompartmentID:      "compartmentid",
			DisplayName:        "display_name",
			ID:                 "id1",
			RouteTableID:       "routetableid",
			SecurityListIDs: []string{
				"slid1",
				"slid2",
			},
			State: baremetal.ResourceAvailable,
			TimeCreated: baremetal.Time{
				Time: time.Now(),
			},
			VcnID:            "vcnid",
			VirtualRouterIP:  "virtualrouterip",
			VirtualRouterMac: "virtualroutermac",
		},
		{
			AvailabilityDomain: "availabilitydomainid",
			CIDRBlock:          "10.10.11.0/24",
			CompartmentID:      "compartmentid",
			DisplayName:        "display_name",
			ID:                 "id2",
			RouteTableID:       "routetableid",
			SecurityListIDs: []string{
				"slid1",
				"slid2",
			},
			State: baremetal.ResourceAvailable,
			TimeCreated: baremetal.Time{
				Time: time.Now(),
			},
			VcnID:            "vcnid",
			VirtualRouterIP:  "virtualrouterip",
			VirtualRouterMac: "virtualroutermac",
		},
	}
	s.Client.On(
		"ListSubnets",
		"compartmentid",
		"vcnid",
		opts,
	).Return(res, nil)

	opts2 := &baremetal.ListOptions{}
	opts2.Page = "nextpage"

	s.Client.On(
		"ListSubnets",
		"compartmentid",
		"vcnid",
		opts2,
	).Return(
		&baremetal.ListSubnets{
			Subnets: []baremetal.Subnet{
				{
					AvailabilityDomain: "availabilitydomainid",
					CIDRBlock:          "10.10.10.0/24",
					CompartmentID:      "compartmentid",
					DisplayName:        "display_name",
					ID:                 "id3",
					RouteTableID:       "routetableid",
					SecurityListIDs: []string{
						"slid1",
						"slid2",
					},
					State: baremetal.ResourceAvailable,
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
					VcnID:            "vcnid",
					VirtualRouterIP:  "virtualrouterip",
					VirtualRouterMac: "virtualroutermac",
				},
				{
					AvailabilityDomain: "availabilitydomainid",
					CIDRBlock:          "10.10.11.0/24",
					CompartmentID:      "compartmentid",
					DisplayName:        "display_name",
					ID:                 "id4",
					RouteTableID:       "routetableid",
					SecurityListIDs: []string{
						"slid1",
						"slid2",
					},
					State: baremetal.ResourceAvailable,
					TimeCreated: baremetal.Time{
						Time: time.Now(),
					},
					VcnID:            "vcnid",
					VirtualRouterIP:  "virtualrouterip",
					VirtualRouterMac: "virtualroutermac",
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
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", "vcnid"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.availability_domain", "availabilitydomainid"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListSubnets", "compartmentid", "vcnid", opts2)

}

func TestResourceCoreSubnetsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreSubnetsTestSuite))
}
