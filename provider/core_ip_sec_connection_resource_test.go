// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"fmt"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreIPSecTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreIPSecTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
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
	var resId, resId2 string
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "cpe_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.IpSecConnectionLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "static_routes.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "static_routes.0", "10.0.0.0/16"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_ipsec.t", "id")
						return err
					},
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-ipsec"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "cpe_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.IpSecConnectionLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "static_routes.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "static_routes.0", "10.0.0.0/16"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_core_ipsec.t", "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						resId = resId2
						return err
					},
				),
			},
			// Verify Force New Update
			{
				Config: s.Config + `
					resource "oci_core_drg" "u" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-drg-ipsec-upd"
					}
					resource "oci_core_cpe" "u" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-cpe-ipsec-upd"
						ip_address = "124.124.124.124"
						depends_on = ["oci_core_drg.u"]
					}
					resource "oci_core_ipsec" "t" {
						compartment_id = "${var.compartment_id}"
						cpe_id = "${oci_core_cpe.u.id}"
						drg_id = "${oci_core_drg.u.id}"
						display_name = "-tf-ipsec"
						static_routes = ["10.0.0.0/16"]
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-ipsec"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "cpe_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.IpSecConnectionLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "static_routes.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "static_routes.0", "10.0.0.0/16"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_core_ipsec.t", "id")
						if resId == resId2 {
							return fmt.Errorf("expected new resource to be created but was it was not")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceCoreIPSecTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreIPSecTestSuite))
}
