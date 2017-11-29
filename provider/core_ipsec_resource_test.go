// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreIPSecTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreIPSecTestSuite) SetupTest() {
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
			depends_on = ["oci_core_drg.t"]
		}`

	s.ResourceName = "oci_core_ipsec.t"
}

func (s *ResourceCoreIPSecTestSuite) TestAccResourceCoreIpsec_basic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_core_ipsec" "t" {
					compartment_id = "${var.compartment_id}"
					cpe_id = "${oci_core_cpe.t.id}"
					drg_id = "${oci_core_drg.t.id}"
					static_routes = ["10.0.0.0/16"]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
				),
			},
			// verify update
			{
				Config: s.Config + `
				resource "oci_core_ipsec" "t" {
					compartment_id = "${var.compartment_id}"
					cpe_id = "${oci_core_cpe.t.id}"
					drg_id = "${oci_core_drg.t.id}"
					display_name = "-tf-ipsec"
					static_routes = ["10.0.0.0/16"]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-ipsec"),
				),
			},
		},
	})
}

func TestResourceCoreIPSecTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreIPSecTestSuite))
}
