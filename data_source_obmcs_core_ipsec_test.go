// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreIPSecTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreIPSecTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_ipsec_connections" "s" {
      compartment_id = "${var.compartment_id}"
      cpe_id = "cpeid"
      drg_id = "drgid"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_ipsec_connections.s"

}

func (s *DatasourceCoreIPSecTestSuite) TestResourceListIPConnections() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "drg_id", "drgid"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpe_id", "cpeid"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.0.compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "connections.#", "2"),
				),
			},
		},
	},
	)

}

func TestDatasourceCoreIPSecTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreIPSecTestSuite))
}
