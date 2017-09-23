// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
	
	resource "oci_core_virtual_network" "t" {
		cidr_block     = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name   = "network_name"
	}
	
	resource "oci_core_subnet" "s" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		route_table_id = "${oci_core_virtual_network.t.default_route_table_id}"
		dhcp_options_id = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block = "10.0.2.0/24"
	}`
	
	s.ResourceName = "data.oci_core_subnets.s"
}

func (s *DatasourceCoreSubnetTestSuite) TestAccDatasourceCoreSubnet_basic() {

	resource.Test(s.T(), resource.TestCase{
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
				data "oci_core_subnets" "s" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
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

func TestDatasourceCoreSubnetTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreSubnetTestSuite))
}
