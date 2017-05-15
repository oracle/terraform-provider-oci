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

type ResourceCoreSubnetTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Subnet
	DeletedRes   *baremetal.Subnet
}

func (s *ResourceCoreSubnetTestSuite) SetupTest() {
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

resource "baremetal_core_internet_gateway" "CompleteIG" {
    compartment_id = "${var.compartment_id}"
    display_name = "CompleteIG"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
}

resource "baremetal_core_route_table" "RouteForComplete" {
    compartment_id = "${var.compartment_id}"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
    display_name = "RouteTableForComplete"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${baremetal_core_internet_gateway.CompleteIG.id}"
    }
}

resource "baremetal_core_security_list" "WebSubnet" {
    compartment_id = "${var.compartment_id}"
    display_name = "Public"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
    egress_security_rules = [{
        destination = "0.0.0.0/0"
        protocol = "6"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 80
            "min" = 80
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
	{
	protocol = "6"
	source = "10.0.0.0/16"
    }]
}


resource "baremetal_core_subnet" "WebSubnetAD1" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block = "10.0.0.0/16"
  display_name = "WebSubnetAD1"
  compartment_id = "${var.compartment_id}"
  vcn_id = "${baremetal_core_virtual_network.t.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.WebSubnet.id}"]
}

	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_subnet.t"

}

func (s *ResourceCoreSubnetTestSuite) TestCreateResourceCoreSubnet() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", s.Res.AvailabilityDomain),

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func (s *ResourceCoreSubnetTestSuite) TestTerminateSubnet() {
	if IsAccTest() {
		s.T().Skip()
	}

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

func TestResourceCoreSubnetTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreSubnetTestSuite))
}
