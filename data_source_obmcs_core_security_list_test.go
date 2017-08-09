// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type CoreSecurityListDatasourceTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *CoreSecurityListDatasourceTestSuite) SetupTest() {
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
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_security_lists.t"
}

func (s *CoreSecurityListDatasourceTestSuite) TestReadSecurityLists() {
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
				    data "baremetal_core_security_lists" "t" {
				      compartment_id = "${var.compartment_id}"
				      limit = 1
				      vcn_id = "${baremetal_core_virtual_network.t.id}"
				    }`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "security_lists.0.ingress_security_rules.0.tcp_options.0.max"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "security_lists.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.#", "2"),
				),
			},
		},
	},
	)
}

func TestCoreSecurityListDatasourceTestSuite(t *testing.T) {
	suite.Run(t, new(CoreSecurityListDatasourceTestSuite))
}
