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

type ResourceCoreVirtualNetworkTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.VirtualNetwork
	DeletedRes   *baremetal.VirtualNetwork
	DeletingRes  *baremetal.VirtualNetwork
}

func (s *ResourceCoreVirtualNetworkTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
		resource "baremetal_core_virtual_network" "t" {
			cidr_block = "10.0.0.0/16"
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_virtual_network.t"

}

func (s *ResourceCoreVirtualNetworkTestSuite) TestCreateResourceCoreVirtualNetwork() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", s.Res.CidrBlock),

					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
				),
			},
		},
	})
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestDeleteResourceCoreVirtualNetwork() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
				),
			},
			{
				Config: testProviderConfig(),
				Check: resource.ComposeTestCheckFunc(
					testNoInstanceState("baremetal_core_virtual_network"),
				),
			},
		},
	})
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestCreateResourceCoreVirtualNetworkWithoutDisplayName() {
	if IsAccTest() {
		s.T().Skip()
	}

	s.Config = `
		resource "baremetal_core_virtual_network" "t" {
			cidr_block = "10.0.0.0/16"
			compartment_id = "${var.compartment_id}"
		}
	`
	s.Config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
				),
			},
		},
	})
}

func (s ResourceCoreVirtualNetworkTestSuite) TestUpdateCidrBlockForcesNewVirtualNetwork() {
	// Step 1 uses the mocking in Setup plus the following two Get mocks to create
	// and then destroy the original resource.

	config := `
		resource "baremetal_core_virtual_network" "t" {
			cidr_block = "10.0.0.0/24"
			compartment_id = "${var.compartment_id}"
		}
  `
	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/24"),
				),
			},
		},
	})
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestDeleteVirtualNetwork() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

}

func TestResourceCoreVirtualNetworkTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVirtualNetworkTestSuite))
}
