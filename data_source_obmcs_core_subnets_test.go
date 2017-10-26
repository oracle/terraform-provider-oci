// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreSubnetTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreSubnetTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "vcn1" {
		cidr_block     = "10.1.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name   = "vcn1"
		dns_label = "vnc1"
	}
	
	resource "oci_core_subnet" "s" {	
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.vcn1.id}"
		count = "${length(data.oci_identity_availability_domains.ADs.availability_domains)}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[count.index], "name")}"
		cidr_block = "10.1.2${count.index}.0/24"
		display_name = "subnet${count.index}"
		dns_label = "subnet${count.index}"
		security_list_ids = ["${oci_core_virtual_network.vcn1.default_security_list_id}"]
		route_table_id = "${oci_core_virtual_network.vcn1.default_route_table_id}"
		dhcp_options_id = "${oci_core_virtual_network.vcn1.default_dhcp_options_id}"
	}`

	s.ResourceName = "data.oci_core_subnets.s"
}

func (s *DatasourceCoreSubnetTestSuite) TestAccDatasourceCoreSubnet_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			// check properties on the subnet created in AD2
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_core_subnets" "s" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.vcn1.id}"
					depends_on = ["oci_core_subnet.s"]
					filter {
						name   = "availability_domain"
						values = ["${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1], "name")}"]
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnets.0.availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.cidr_block", "10.1.21.0/24"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.display_name", "subnet1"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.dns_label", "subnet1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnets.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.prohibit_public_ip_on_vnic", "false"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnets.0.route_table_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.security_list_ids.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnets.0.time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnets.0.vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnets.0.dhcp_options_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnets.0.virtual_router_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnets.0.virtual_router_mac"),
				),
				// Work around for terraform bug where depends_on results in "plan was not empty"
				// https://github.com/hashicorp/terraform/issues/11139
				ExpectNonEmptyPlan: true,
			},
		},
	},
	)
}

func TestDatasourceCoreSubnetTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreSubnetTestSuite))
}
