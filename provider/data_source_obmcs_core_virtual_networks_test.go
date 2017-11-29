// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVirtualNetworkTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	Token        string
	TokenFn      TokenFn
}

func (s *DatasourceCoreVirtualNetworkTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + s.TokenFn(`
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "{{.token}}"
		cidr_block = "10.0.0.0/16"
		dns_label = "vcn1"
	}`, nil)
	s.ResourceName = "data.oci_core_virtual_networks.t"
}

func (s *DatasourceCoreVirtualNetworkTestSuite) TestAccDatasourceCoreVirtualNetwork_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + s.TokenFn(`
				data "oci_core_virtual_networks" "t" {
					compartment_id = "${oci_core_virtual_network.t.compartment_id}"
					filter {
						name = "display_name"
						values = ["{{.token}}"]
					}
				}`, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.display_name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.0.default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.0.default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.0.default_dhcp_options_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.dns_label", "vcn1"),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.0.time_created"),
				),
			},
		},
	},
	)

}

func TestDatasourceCoreVirtualNetworkTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVirtualNetworkTestSuite))
}
