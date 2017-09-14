// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
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
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-drg"
	}
	resource "oci_core_cpe" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-cpe"
		ip_address = "123.123.123.123"
	}
	resource "oci_core_ipsec" "t" {
		compartment_id = "${var.compartment_id}"
		cpe_id = "${oci_core_cpe.t.id}"
		drg_id = "${oci_core_drg.t.id}"
		display_name = "-tf-ipsec"
		static_routes = ["10.0.0.0/16"]
	}
	data "oci_core_ipsec_status" "s" {
		ipsec_id = "${oci_core_ipsec.t.id}"
	}`
	s.ResourceName = "data.oci_core_ipsec_status.s"
}

func (s *DatasourceCoreIPSecStatusTestSuite) TestAccDatasourceCoreIPSecStatus_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
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
