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

type ResourceCoreRouteTableTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.RouteTable
	DeletedRes   *baremetal.RouteTable
}

func (s *ResourceCoreRouteTableTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
resource "oci_core_virtual_network" "t" {
	cidr_block = "10.0.0.0/16"
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
}
resource "oci_core_internet_gateway" "CompleteIG" {
    compartment_id = "${var.compartment_id}"
    display_name = "CompleteIG"
    vcn_id = "${oci_core_virtual_network.t.id}"
}
resource "oci_core_route_table" "t" {
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
	route_rules {
		cidr_block = "0.0.0.0/0"
		network_entity_id = "${oci_core_internet_gateway.CompleteIG.id}"
	}
	vcn_id = "${oci_core_virtual_network.t.id}"
}
	`
	s.Config += testProviderConfig()

	s.ResourceName = "oci_core_route_table.t"

}

func (s *ResourceCoreRouteTableTestSuite) TestCreateResourceCoreRouteTable() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "route_rules.0.network_entity_id"),
				),
			},
		},
	})
}

func (s ResourceCoreRouteTableTestSuite) TestUpdateRouteTable() {
	config := `
		resource "oci_core_route_table" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
      route_rules {
				cidr_block = "new_cidr_block"
				network_entity_id = "network_entity_id"
			}
			vcn_id = "vcn_id"
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
				Check:  resource.TestCheckResourceAttr(s.ResourceName, "route_rules.0.cidr_block", "new_cidr_block"),
			},
		},
	})
}

func (s *ResourceCoreRouteTableTestSuite) TestDeleteRouteTable() {

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

func TestResourceCoreRouteTableTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreRouteTableTestSuite))
}
