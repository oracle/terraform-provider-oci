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

type DatasourceCoreIPSecStatusTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
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
	resource "baremetal_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}
	resource "baremetal_core_cpe" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "displayname"
		ip_address = "123.123.123.123"
	}
	resource "baremetal_core_ipsec" "t" {
		compartment_id = "${var.compartment_id}"
		cpe_id = "${baremetal_core_cpe.t.id}"
		drg_id = "${baremetal_core_drg.t.id}"
		display_name = "display_name"
		static_routes = ["10.0.0.0/16"]
	}
    	data "baremetal_core_ipsec_status" "s" {
      		ipsec_id = "${baremetal_core_ipsec.t.id}"
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
					resource.TestCheckResourceAttrSet(s.ResourceName, "tunnels.#"),
				),
			},
		},
	},
	)

}

func TestDatasourceCoreIPSecStatusTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreIPSecStatusTestSuite))
}
