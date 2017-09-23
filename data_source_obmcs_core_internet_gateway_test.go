// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreInternetGatewayTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreInternetGatewayTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-vcn"
	}
	resource "oci_core_internet_gateway" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-internet-gateway"
		vcn_id = "${oci_core_virtual_network.t.id}"
	}`
	s.ResourceName = "data.oci_core_internet_gateways.s"
}

func (s *DatasourceCoreInternetGatewayTestSuite) TestAccDatasourceCoreInternetGateway_basic() {
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
				data "oci_core_internet_gateways" "s" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.display_name", "-tf-internet-gateway"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "gateways.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.#", "1"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreInternetGatewayTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreInternetGatewayTestSuite))
}
