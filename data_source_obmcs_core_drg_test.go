// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreDrgTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	Token        string
	TokenFn      TokenFn
}

func (s *DatasourceCoreDrgTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + s.TokenFn(`
	resource "oci_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "{{.token}}"
	}`, nil)
	s.ResourceName = "data.oci_core_drgs.t"
}

func (s *DatasourceCoreDrgTestSuite) TestAccDatasourceCoreDrg_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			//{
			//	Config:            s.Config,
			//},
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + s.TokenFn(`
				data "oci_core_drgs" "t" {
					compartment_id = "${var.compartment_id}"
					depends_on = ["oci_core_drg.t"]
					filter {
						name = "display_name"
						values = ["{{.token}}"]
					}
				}`, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.0.display_name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.0.state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drgs.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drgs.0.time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drgs.0.compartment_id"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	},
	)
}

func TestDatasourceCoreDrgsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreDrgTestSuite))
}
