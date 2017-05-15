// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreIPSecStatusTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreIPSecStatusTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_ipsec_status" "s" {
      ipsec_id = "ipsecid"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_ipsec_status.s"

}

func (s *DatasourceCoreIPSecStatusTestSuite) TestIPSecStatus() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "id", "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "tunnels.0.ip_address", "10.10.10.2"),
					resource.TestCheckResourceAttr(s.ResourceName, "tunnels.1.ip_address", "10.10.10.3"),
				),
			},
		},
	},
	)

}

func TestDatasourceCoreIPSecStatusTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreIPSecStatusTestSuite))
}
