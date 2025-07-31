// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVirtualNetworkTestSuite struct {
	suite.Suite
	Providers    map[string]*schema.Provider
	Config       string
	ResourceName string
}

func (s *ResourceCoreVirtualNetworkTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig()
	s.ResourceName = "oci_core_virtual_network.t"
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestAccResourceCoreVirtualNetwork_basic() {
	var resId string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test Create with cidr_block
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test Create with cidr_blocks
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_blocks = ["10.0.0.0/16", "11.0.0.0/16"]
						compartment_id = "${var.compartment_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_blocks.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test add cidr with cidr_blocks
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_blocks = ["10.0.0.0/16", "11.0.0.0/16", "12.0.0.0/16"]
						compartment_id = "${var.compartment_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_blocks.#", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test remove cidr with cidr_blocks
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_blocks = ["10.0.0.0/16", "12.0.0.0/16"]
						compartment_id = "${var.compartment_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_blocks.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test modify cidr with cidr_blocks
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_blocks = ["10.0.0.0/16", "11.0.0.0/16"]
						compartment_id = "${var.compartment_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_blocks.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test Update
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-vcn"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-vcn"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
				),
			},
			// test a destructive Update results in a new resource
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/24"
						compartment_id = "${var.compartment_id}"
						dns_label= "MyTestDNSLabel"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "dns_label", "mytestdnslabel"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId2, err := acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
						if resId == resId2 {
							return fmt.Errorf("Expected new vcn ocid, got the same")
						}
						return err
					},
				),
			},
			// DNS capitalization changes should be ignored.
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/24"
						compartment_id = "${var.compartment_id}"
						dns_label= "mYtESTdnsLABEL"
					}`,
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			// DNS label change should cause a change
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/24"
						compartment_id = "${var.compartment_id}"
						dns_label= "mynewlabel"
					}`,
				ExpectNonEmptyPlan: true,
				PlanOnly:           true,
			},
		},
	})
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestAccResourceCoreVirtualNetwork_ipv6() {
	var resId string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test Create non ipv6enabled vcn
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "is_ipv6enabled", "false"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test Update with ULA prefix
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = false
						ipv6private_cidr_blocks = ["fc00::/48"]
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "ipv6private_cidr_blocks.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_ipv6enabled", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "ipv6cidr_blocks.#", "0"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId2, err := acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
						if resId != resId2 {
							return fmt.Errorf("expected same vcn ocid, got different")
						}
						return err
					},
				),
			},
			// Step add GUA cidr
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = true
						ipv6private_cidr_blocks = ["fc00::/48"]
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "ipv6private_cidr_blocks.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "ipv6cidr_blocks.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_ipv6enabled", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId2, err := acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
						if resId != resId2 {
							return fmt.Errorf("expected same vcn ocid, got different")
						}
						return err
					},
				),
			},
			// Step remove ULA cidr
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = true
						ipv6private_cidr_blocks = []
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "ipv6private_cidr_blocks.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "ipv6cidr_blocks.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_ipv6enabled", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId2, err := acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
						if resId != resId2 {
							return fmt.Errorf("expected same vcn ocid, got different")
						}
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: core/virtualNetwork
func TestResourceCoreVirtualNetworkTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreVirtualNetworkTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreVirtualNetworkTestSuite))
}
