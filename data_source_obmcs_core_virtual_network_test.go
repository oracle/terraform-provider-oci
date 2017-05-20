// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVirtualNetworksTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVirtualNetworksTestSuite) SetupTest() {
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
	display_name = "display_name"
}
data "baremetal_core_virtual_networks" "t" {
	compartment_id = "${baremetal_core_virtual_network.t.compartment_id}"
	limit = 1
}
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_virtual_networks.t"
}

func (s *ResourceCoreVirtualNetworksTestSuite) TestReadVirtualNetworks() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.#"),
				),
			},
		},
	},
	)

}

func TestResourceCoreVirtualNetworksTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVirtualNetworksTestSuite))
}
