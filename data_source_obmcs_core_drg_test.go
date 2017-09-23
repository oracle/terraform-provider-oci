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
}

func (s *DatasourceCoreDrgTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-drg"
	}`
	s.ResourceName = "data.oci_core_drgs.t"
}

func (s *DatasourceCoreDrgTestSuite) TestAccDatasourceCoreDrg_basic() {

	resource.Test(s.T(), resource.TestCase{
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
				data "oci_core_drgs" "t" {
					compartment_id = "${var.compartment_id}"
					limit = 1
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "drgs.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drgs.#"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreDrgsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreDrgTestSuite))
}
