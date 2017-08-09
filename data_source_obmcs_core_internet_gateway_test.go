// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-baremetal/client"
)

type CoreInternetGatewayDatasourceTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *CoreInternetGatewayDatasourceTestSuite) SetupTest() {
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

resource "baremetal_core_internet_gateway" "t" {
    compartment_id = "${var.compartment_id}"
    display_name = "display_name"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
}
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_internet_gateways.s"

}

func (s *CoreInternetGatewayDatasourceTestSuite) TestResourceListInternetGateways() {
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
				data "baremetal_core_internet_gateways" "s" {
				      compartment_id = "${var.compartment_id}"
				      vcn_id = "${baremetal_core_virtual_network.t.id}"
				    }`,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "gateways.0.display_name", "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "gateways.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "gateways.#", "1"),
				),
			},
		},
	},
	)
}

func TestCoreInternetGatewayDatasource(t *testing.T) {
	suite.Run(t, new(CoreInternetGatewayDatasourceTestSuite))
}
