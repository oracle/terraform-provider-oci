// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreDHCPOptionsDatasourceTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListDHCPOptions
}

func (s *ResourceCoreDHCPOptionsDatasourceTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
	resource "baremetal_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "network_name"
	}
	resource "baremetal_core_dhcp_options" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
     		options {
			type = "DomainNameServer"
			custom_dns_servers = [ "8.8.8.8" ]
			server_type = "CustomDnsServer"
		}
     		vcn_id = "${baremetal_core_virtual_network.t.id}"
	}
    data "baremetal_core_dhcp_options" "t" {
      compartment_id = "${var.compartment_id}"
      limit = 1
      vcn_id = "${baremetal_core_virtual_network.t.id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_dhcp_options.t"
}

func (s *ResourceCoreDHCPOptionsDatasourceTestSuite) TestReadDHCPOptions() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "options.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "options.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
				),
			},
		},
	},
	)

}

func TestResourceCoreDHCPOptionsDatasourceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDHCPOptionsDatasourceTestSuite))
}
