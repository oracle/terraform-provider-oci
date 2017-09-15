// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreDrgTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Res          *baremetal.Drg
	DeletedRes   *baremetal.Drg
}

func (s *ResourceCoreDrgTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig()
	s.ResourceName = "oci_core_drg.t"
}

func (s *ResourceCoreDrgTestSuite) TestAccResourceCoreDrg_basic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify a drg can be created
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: testProviderConfig() + `
				resource "oci_core_drg" "t" {
					compartment_id = "${var.compartment_id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("oci_core_drg.t", "id"),
					resource.TestCheckResourceAttrSet("oci_core_drg.t", "time_created"),
					resource.TestCheckResourceAttrSet("oci_core_drg.t", "display_name"),
					resource.TestCheckResourceAttr("oci_core_drg.t", "state", baremetal.ResourceAvailable),
				),
			},
			// verify drg update
			{
				Config: testProviderConfig() + `
				resource "oci_core_drg" "t" {
					compartment_id = "${var.compartment_id}"
					display_name = "-tf-drg"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("oci_core_drg.t", "display_name", "-tf-drg"),
				),
			},
		},
	})
}

func TestResourceCoreDrgTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDrgTestSuite))
}
