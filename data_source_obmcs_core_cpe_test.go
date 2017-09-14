// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreCpeTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreCpeTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_core_cpe" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-cpe"
		ip_address = "142.10.10.1"
	}
	data "oci_core_cpes" "s" {
		compartment_id = "${oci_core_cpe.t.compartment_id}"
		limit = 1
	}`
	s.ResourceName = "data.oci_core_cpes.s"
}

func (s *DatasourceCoreCpeTestSuite) TestAccDatasourceCoreCpe_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.0.ip_address", "142.10.10.1"),
					resource.TestCheckResourceAttr(s.ResourceName, "cpes.0.display_name", "-tf-cpe"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "cpes.#"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreCpeTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreCpeTestSuite))
}
