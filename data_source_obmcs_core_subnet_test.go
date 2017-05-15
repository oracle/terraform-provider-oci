// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreSubnetsTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreSubnetsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
data "baremetal_identity_availability_domains" "ADs" {
  compartment_id = "${var.compartment_id}"
}

resource "baremetal_core_virtual_network" "t" {
	cidr_block = "10.0.0.0/16"
	compartment_id = "${var.compartment_id}"
	display_name = "network_name"
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
  availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
  cidr_block = "10.0.0.0/16"
  display_name = "WebSubnetAD1"
  compartment_id = "${var.compartment_id}"
  vcn_id = "${baremetal_core_virtual_network.t.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.WebSubnet.id}"]
}
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_subnets.s"

}

func (s *ResourceCoreSubnetsTestSuite) TestResourceListSubnets() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				    data "baremetal_core_subnets" "s" {
				      compartment_id = "${var.compartment_id}"
				      vcn_id = "${baremetal_core_virtual_network.t.id}"
				    }`,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnets.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnets.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.#", "1"),
				),
			},
		},
	},
	)
}

func TestResourceCoreSubnetsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreSubnetsTestSuite))
}
