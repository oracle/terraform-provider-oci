// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreInternetGatewayTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreInternetGatewayTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
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
				Config: s.Config + `
				data "oci_core_internet_gateways" "s" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
				
					filter {
						name = "display_name"
						values = ["${oci_core_internet_gateway.t.display_name}"]
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.display_name", "-tf-internet-gateway"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.state", string(core.InternetGatewayLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.enabled", "true"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "gateways.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "gateways.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "gateways.0.vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "gateways.0.time_created"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreInternetGatewayTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreInternetGatewayTestSuite))
}
