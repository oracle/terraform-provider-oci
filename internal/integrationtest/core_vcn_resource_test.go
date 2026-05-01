// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strings"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

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

func (s *ResourceCoreVirtualNetworkTestSuite) TestAccResourceCoreVirtualNetwork_byoipv6() {
	// Using an already defined byoip range id
	byoipv6RangeId := acctest.GetEnvSettingWithDefaultVar("byoipv6_range_id", "unknown")
	if byoipv6RangeId != "unknown" {
		resource.Test(s.T(), resource.TestCase{
			Providers: s.Providers,
			Steps: []resource.TestStep{
				{
					Config: s.Config + fmt.Sprintf(`
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = false
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2200::/64"
                        }
					}`, byoipv6RangeId),
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
						resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
						resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
						resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
						resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/16"),
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "1"),
						func(s *terraform.State) (err error) {
							resId, err = acctest.FromInstanceState(s, "oci_core_virtual_network.t", "id")
							return err
						},
					),
				},
				// Step to force a drift and validate that the plan shows the drift
				{
					PreConfig: func() {
						err := addIpv6CidrToVcn(acctest.GetTestClients(&schema.ResourceData{}), resId, "2607:f590:0000:2201::/64", byoipv6RangeId)
						if err != nil {
							return
						}
					},
					Config: s.Config + fmt.Sprintf(`
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = false
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2200::/64"
                        }
					}`, byoipv6RangeId),
				},
				{
					RefreshState:       true,
					ExpectNonEmptyPlan: true,
					PlanOnly:           true,
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "1"),
					),
				},
				// Step to add the missing cidr to config and validate that the plan does not show any diff
				{
					Config: s.Config + fmt.Sprintf(`
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = false
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2200::/64"
                        }
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2201::/64"
                        }
					}`, byoipv6RangeId, byoipv6RangeId),
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "2"),
					),
				},
				{
					RefreshState:       true,
					ExpectNonEmptyPlan: false,
					PlanOnly:           true,
				},

				// Step to replace range id with a random range id and validate that the plan does not show any diff
				{
					ExpectNonEmptyPlan: false,
					PlanOnly:           true,
					Config: s.Config + fmt.Sprintf(`
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = false
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2200::/64"
                        }
						byoipv6cidr_details {
          					byoipv6range_id = "(known_after_apply)"
          					ipv6cidr_block  = "2607:f590:0000:2201::/64"
                        }
					}`, byoipv6RangeId),
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "2"),
					),
				},
				// Test - Drift created when byoipv6 cidr removed from resource
				// Step 1 - validate force removal of cidr creates a drift
				{
					PreConfig: func() {
						err := removeIpv6CidrFromVcn(acctest.GetTestClients(&schema.ResourceData{}), resId, "2607:f590:0000:2201::/64", 1)
						if err != nil {
							return
						}
					},
					RefreshState:       true,
					ExpectNonEmptyPlan: true,
					PlanOnly:           true,
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "1"),
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_details.#", "1"),
					),
				},
				// Step 2 - Validate that Intact / unchanged config should lead to application of the missing cidr to the VCN
				{
					Config: s.Config + fmt.Sprintf(`
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = false
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2200::/64"
                        }
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2201::/64"
                        }
					}`, byoipv6RangeId, byoipv6RangeId),
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "2"),
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_details.#", "2"),
					),
				},
				// Step 3 - Post application of the new cidr, terraform should show an empty plan
				{
					RefreshState:       true,
					ExpectNonEmptyPlan: false,
					PlanOnly:           true,
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "2"),
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_details.#", "2"),
					),
				},
				// Test - Drift created with one change when byoipv6 cidr removed from the middle of the list
				// Step 1 - Add one more cidr at the end
				{
					Config: s.Config + fmt.Sprintf(`
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = false
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2200::/64"
                        }
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2201::/64"
                        }
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2202::/64"
                        }
					}`, byoipv6RangeId, byoipv6RangeId, byoipv6RangeId),
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "3"),
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_details.#", "3"),
					),
				},
				// Step 2 - Out of band Remove cidr in the middle from the list
				{
					PreConfig: func() {
						err := removeIpv6CidrFromVcn(acctest.GetTestClients(&schema.ResourceData{}), resId, "2607:f590:0000:2201::/64", 1)
						if err != nil {
							return
						}
					},
					RefreshState:       true,
					ExpectNonEmptyPlan: true,
					PlanOnly:           true,
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "2"),
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_details.#", "2"),
					),
				},
				// Step 3 - Post removal of the new cidr from config, terraform should show an empty plan
				{
					ExpectNonEmptyPlan: false,
					PlanOnly:           true,
					Config: s.Config + fmt.Sprintf(`
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = false
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2200::/64"
                        }
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2202::/64"
                        }
					}`, byoipv6RangeId, byoipv6RangeId),
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "2"),
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_details.#", "2"),
					),
				},
				// Step 4 - Post removal the stale cidr from config, terraform apply should succeed
				{
					Config: s.Config + fmt.Sprintf(`
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = false
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2200::/64"
                        }
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2202::/64"
                        }
					}`, byoipv6RangeId, byoipv6RangeId),
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "2"),
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_details.#", "2"),
					),
				},
				// Test - Drift created when byoipv6 cidr removed from resource
				// Step 1 - validate force removal of cidr creates a drift
				{
					PreConfig: func() {
						err := removeIpv6CidrFromVcn(acctest.GetTestClients(&schema.ResourceData{}), resId, "2607:f590:0000:2202::/64", 1)
						if err != nil {
							return
						}
					},
					RefreshState:       true,
					ExpectNonEmptyPlan: true,
					PlanOnly:           true,
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "1"),
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_details.#", "1"),
					),
				},
				// Step 2 - Post removal of the new cidr from config, terraform should show an empty plan
				{
					ExpectNonEmptyPlan: false,
					PlanOnly:           true,
					Config: s.Config + fmt.Sprintf(`
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						is_ipv6enabled = true
  						is_oracle_gua_allocation_enabled = false
						byoipv6cidr_details {
          					byoipv6range_id = %q
          					ipv6cidr_block  = "2607:f590:0000:2200::/64"
                        }
					}`, byoipv6RangeId),
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_blocks.#", "1"),
						resource.TestCheckResourceAttr(s.ResourceName, "byoipv6cidr_details.#", "1"),
					),
				},
			},
		})
	}
}

// issue-routing-tag: core/virtualNetwork
func TestAccResourceCoreVirtualNetworkPatch_ipv6(t *testing.T) {
	httpreplay.SetScenario("TestAccResourceCoreVirtualNetworkPatch_ipv6")
	defer httpreplay.SaveScenario()
	acctest.PreCheck(t)

	byoipv6RangeId := acctest.GetEnvSettingWithDefaultVar("byoipv6_range_id", "unknown")
	if byoipv6RangeId == "unknown" {
		t.Skip("TF_VAR_byoipv6_range_id must be set for VCN IPv6 patch acceptance tests")
	}

	resourceName := "oci_core_virtual_network.patch_vcn"

	byoCidrs := coreVcnPatchTestByoIpv6Cidrs()
	initialByo := []string{byoCidrs[0]}
	byoAfterBulkAdd := append([]string{}, byoCidrs[:14]...)
	byoAfterBeginningReplace := append([]string{byoCidrs[14]}, byoAfterBulkAdd[1:]...)
	byoMiddleReplaceIndex := len(byoAfterBeginningReplace) / 2
	byoMiddleReplacement := byoAfterBulkAdd[0]
	byoEndReplacement := byoAfterBeginningReplace[byoMiddleReplaceIndex]
	byoAfterMiddleReplace := append([]string{}, byoAfterBeginningReplace...)
	byoAfterMiddleReplace[byoMiddleReplaceIndex] = byoMiddleReplacement
	byoAfterEndReplace := append([]string{}, byoAfterMiddleReplace[:len(byoAfterMiddleReplace)-1]...)
	byoAfterEndReplace = append(byoAfterEndReplace, byoEndReplacement)
	byoAfterMultiReplace := append([]string{}, byoAfterEndReplace...)
	byoAfterMultiReplace[0] = byoCidrs[15]
	byoAfterMultiReplace[len(byoAfterMultiReplace)/2-1] = byoCidrs[16]
	byoAfterMultiReplace[len(byoAfterMultiReplace)/2] = byoCidrs[17]
	byoAfterReorder := append([]string{byoAfterMultiReplace[1], byoAfterMultiReplace[2], byoAfterMultiReplace[0]}, byoAfterMultiReplace[3:]...)
	removedBeginningByo := byoAfterReorder[0]
	byoAfterBeginningRemove := append([]string{}, byoAfterReorder[1:]...)
	byoMiddleRemoveIndex := len(byoAfterBeginningRemove) / 2
	removedMiddleByo := byoAfterBeginningRemove[byoMiddleRemoveIndex]
	byoAfterMiddleRemove := append([]string{}, byoAfterBeginningRemove[:byoMiddleRemoveIndex]...)
	byoAfterMiddleRemove = append(byoAfterMiddleRemove, byoAfterBeginningRemove[byoMiddleRemoveIndex+1:]...)
	removedEndByo := byoAfterMiddleRemove[len(byoAfterMiddleRemove)-1]
	byoAfterEndRemove := append([]string{}, byoAfterMiddleRemove[:len(byoAfterMiddleRemove)-1]...)
	removedByo := []string{removedBeginningByo, removedMiddleByo, removedEndByo}

	ulaCidrs := coreVcnPatchTestUlaCidrs()
	initialUla := []string{ulaCidrs[0]}
	ulaAfterBulkAdd := append([]string{}, ulaCidrs[:14]...)
	ulaAfterBeginningReplace := append([]string{ulaCidrs[14]}, ulaAfterBulkAdd[1:]...)
	ulaMiddleReplaceIndex := len(ulaAfterBeginningReplace) / 2
	ulaMiddleReplacement := ulaAfterBulkAdd[0]
	ulaEndReplacement := ulaAfterBeginningReplace[ulaMiddleReplaceIndex]
	ulaAfterMiddleReplace := append([]string{}, ulaAfterBeginningReplace...)
	ulaAfterMiddleReplace[ulaMiddleReplaceIndex] = ulaMiddleReplacement
	ulaAfterEndReplace := append([]string{}, ulaAfterMiddleReplace[:len(ulaAfterMiddleReplace)-1]...)
	ulaAfterEndReplace = append(ulaAfterEndReplace, ulaEndReplacement)
	ulaAfterMultiReplace := append([]string{}, ulaAfterEndReplace...)
	ulaAfterMultiReplace[0] = ulaCidrs[15]
	ulaAfterMultiReplace[len(ulaAfterMultiReplace)/2-1] = ulaCidrs[16]
	ulaAfterMultiReplace[len(ulaAfterMultiReplace)/2] = ulaCidrs[17]
	ulaAfterReorder := append([]string{ulaAfterMultiReplace[1], ulaAfterMultiReplace[2], ulaAfterMultiReplace[0]}, ulaAfterMultiReplace[3:]...)
	removedBeginningUla := ulaAfterReorder[0]
	ulaAfterBeginningRemove := append([]string{}, ulaAfterReorder[1:]...)
	ulaMiddleRemoveIndex := len(ulaAfterBeginningRemove) / 2
	removedMiddleUla := ulaAfterBeginningRemove[ulaMiddleRemoveIndex]
	ulaAfterMiddleRemove := append([]string{}, ulaAfterBeginningRemove[:ulaMiddleRemoveIndex]...)
	ulaAfterMiddleRemove = append(ulaAfterMiddleRemove, ulaAfterBeginningRemove[ulaMiddleRemoveIndex+1:]...)
	removedEndUla := ulaAfterMiddleRemove[len(ulaAfterMiddleRemove)-1]
	ulaAfterEndRemove := append([]string{}, ulaAfterMiddleRemove[:len(ulaAfterMiddleRemove)-1]...)
	removedUla := []string{removedBeginningUla, removedMiddleUla, removedEndUla}

	checkByoOrder := func(expected []string) resource.TestCheckFunc {
		return acctest.ComposeAggregateTestCheckFuncWrapper(
			testCheckCanonicalListEquals(resourceName, "byoipv6cidr_blocks", expected),
			testCheckCanonicalNestedIpv6CidrBlockListEquals(resourceName, "byoipv6cidr_details", expected),
		)
	}
	checkUlaOrder := func(expected []string) resource.TestCheckFunc {
		return testCheckCanonicalListEquals(resourceName, "ipv6private_cidr_blocks", expected)
	}

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Step 1 - Create a VCN with 1 byoipv6 cidr and 1 ULA ipv6 cidr (via the ipv6private_cidr_blocks field) and validate.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, initialByo, initialUla),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", initialByo),
				checkByoOrder(initialByo),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", initialUla),
				checkUlaOrder(initialUla),
				resource.TestCheckResourceAttr(resourceName, "is_ipv6enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "state", string(core.VcnLifecycleStateAvailable)),
			),
		},
		// Step 2 - Update VCN with 13 more byoipv6 cidr and validate that terraform plan is non-empty.
		{
			Config:             coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterBulkAdd, initialUla),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
		},
		// Step 3 - Update VCN with 13 more byoipv6 cidr and validate that terraform apply shows just 13 new additions for the byoipv6 cidr. Apply the change and validate that the final state has 14 byoipv6 cidrs in total.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterBulkAdd, initialUla),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							byoFieldChanges: 1,
							byoAdditions:    13,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "14"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterBulkAdd),
				checkByoOrder(byoAfterBulkAdd),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", initialUla),
				checkUlaOrder(initialUla),
			),
		},
		// Step 4 - Update VCN with 13 more byoipv6 cidr and validate that terraform plan shows no new changes.
		{
			Config:             coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterBulkAdd, initialUla),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},
		// Step 5 - Replace a byoipv6 cidr at the beginning of the VCN's byoipv6cidr_details field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterBeginningReplace, initialUla),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							byoFieldChanges:      1,
							byoAdditions:         1,
							byoRemovals:          1,
							byoReplacementGroups: 1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "14"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterBeginningReplace),
				checkByoOrder(byoAfterBeginningReplace),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", initialUla),
				checkUlaOrder(initialUla),
			),
		},
		// Step 6 - Replace a byoipv6 cidr in the middle of the VCN's byoipv6cidr_details field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterMiddleReplace, initialUla),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							byoFieldChanges:      1,
							byoAdditions:         1,
							byoRemovals:          1,
							byoReplacementGroups: 1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "14"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterMiddleReplace),
				checkByoOrder(byoAfterMiddleReplace),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", initialUla),
				checkUlaOrder(initialUla),
			),
		},
		// Step 7 - Replace a byoipv6 cidr at the end of the VCN's byoipv6cidr_details field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndReplace, initialUla),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							byoFieldChanges:      1,
							byoAdditions:         1,
							byoRemovals:          1,
							byoReplacementGroups: 1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "14"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndReplace),
				checkByoOrder(byoAfterEndReplace),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", initialUla),
				checkUlaOrder(initialUla),
			),
		},
		// Step 8 - Replace three byoipv6 cidrs with entirely new cidrs while preserving config order.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterMultiReplace, initialUla),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							byoFieldChanges:      1,
							byoAdditions:         3,
							byoRemovals:          3,
							byoReplacementGroups: 3,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "14"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterMultiReplace),
				checkByoOrder(byoAfterMultiReplace),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", initialUla),
				checkUlaOrder(initialUla),
			),
		},
		// Step 9 - Reorder byoipv6 cidrs without changing membership.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterReorder, initialUla),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation:     vcnIpv6PlanExpectation{},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "14"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterReorder),
				checkByoOrder(byoAfterReorder),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", initialUla),
				checkUlaOrder(initialUla),
			),
		},
		// Step 10 - Remove a byoipv6 cidr from the beginning of the VCN's byoipv6cidr_details field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterBeginningRemove, initialUla),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							byoFieldChanges: 1,
							byoRemovals:     1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "13"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "13"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterBeginningRemove),
				checkByoOrder(byoAfterBeginningRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", initialUla),
				checkUlaOrder(initialUla),
				testCheckCanonicalTypeSetExcludes(resourceName, "byoipv6cidr_blocks", []string{removedBeginningByo}),
			),
		},
		// Step 11 - Remove a byoipv6 cidr from the middle of the VCN's byoipv6cidr_details field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterMiddleRemove, initialUla),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							byoFieldChanges: 1,
							byoRemovals:     1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "12"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "12"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterMiddleRemove),
				checkByoOrder(byoAfterMiddleRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", initialUla),
				checkUlaOrder(initialUla),
				testCheckCanonicalTypeSetExcludes(resourceName, "byoipv6cidr_blocks", []string{removedBeginningByo, removedMiddleByo}),
			),
		},
		// Step 12 - Remove a byoipv6 cidr from the end of the VCN's byoipv6cidr_details field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, initialUla),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							byoFieldChanges: 1,
							byoRemovals:     1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "11"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndRemove),
				checkByoOrder(byoAfterEndRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "1"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", initialUla),
				checkUlaOrder(initialUla),
				testCheckCanonicalTypeSetExcludes(resourceName, "byoipv6cidr_blocks", removedByo),
			),
		},
		// Step 13 - Validate that terraform plan shows no new changes after byoipv6 replacements and removals.
		{
			Config:             coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, initialUla),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},
		// Step 14 - Update VCN with 13 more blocks under ipv6private_cidr_blocks and validate that terraform plan is non-empty.
		{
			Config:             coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterBulkAdd),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
		},
		// Step 15 - Update VCN with 13 more blocks under ipv6private_cidr_blocks and validate that terraform apply shows just 13 new additions for the ipv6private_cidr_blocks. Apply the change and validate that the final state has 14 ipv6private_cidr_blocks in total.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterBulkAdd),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							privateFieldChanges: 1,
							privateAdditions:    13,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "11"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndRemove),
				checkByoOrder(byoAfterEndRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", ulaAfterBulkAdd),
				checkUlaOrder(ulaAfterBulkAdd),
			),
		},
		// Step 16 - Update VCN with 13 more blocks under ipv6private_cidr_blocks and validate that terraform plan shows no new changes.
		{
			Config:             coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterBulkAdd),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},
		// Step 17 - Replace a cidr at the beginning of the VCN's ipv6private_cidr_blocks field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterBeginningReplace),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							privateFieldChanges:      1,
							privateAdditions:         1,
							privateRemovals:          1,
							privateReplacementGroups: 1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "11"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndRemove),
				checkByoOrder(byoAfterEndRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", ulaAfterBeginningReplace),
				checkUlaOrder(ulaAfterBeginningReplace),
			),
		},
		// Step 18 - Replace a cidr in the middle of the VCN's ipv6private_cidr_blocks field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterMiddleReplace),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							privateFieldChanges:      1,
							privateAdditions:         1,
							privateRemovals:          1,
							privateReplacementGroups: 1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "11"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndRemove),
				checkByoOrder(byoAfterEndRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", ulaAfterMiddleReplace),
				checkUlaOrder(ulaAfterMiddleReplace),
			),
		},
		// Step 19 - Replace a cidr at the end of the VCN's ipv6private_cidr_blocks field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterEndReplace),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							privateFieldChanges:      1,
							privateAdditions:         1,
							privateRemovals:          1,
							privateReplacementGroups: 1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "11"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndRemove),
				checkByoOrder(byoAfterEndRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", ulaAfterEndReplace),
				checkUlaOrder(ulaAfterEndReplace),
			),
		},
		// Step 20 - Replace three private cidrs with entirely new cidrs while preserving config order.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterMultiReplace),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							privateFieldChanges:      1,
							privateAdditions:         3,
							privateRemovals:          3,
							privateReplacementGroups: 3,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "11"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndRemove),
				checkByoOrder(byoAfterEndRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", ulaAfterMultiReplace),
				checkUlaOrder(ulaAfterMultiReplace),
			),
		},
		// Step 21 - Reorder private cidrs without changing membership.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterReorder),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation:     vcnIpv6PlanExpectation{},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "11"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndRemove),
				checkByoOrder(byoAfterEndRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "14"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", ulaAfterReorder),
				checkUlaOrder(ulaAfterReorder),
			),
		},
		// Step 22 - Remove a cidr from the beginning of the VCN's ipv6private_cidr_blocks field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterBeginningRemove),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							privateFieldChanges: 1,
							privateRemovals:     1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "11"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndRemove),
				checkByoOrder(byoAfterEndRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "13"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", ulaAfterBeginningRemove),
				checkUlaOrder(ulaAfterBeginningRemove),
				testCheckCanonicalTypeSetExcludes(resourceName, "ipv6private_cidr_blocks", []string{removedBeginningUla}),
			),
		},
		// Step 23 - Remove a cidr from the middle of the VCN's ipv6private_cidr_blocks field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterMiddleRemove),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							privateFieldChanges: 1,
							privateRemovals:     1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "11"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndRemove),
				checkByoOrder(byoAfterEndRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "12"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", ulaAfterMiddleRemove),
				checkUlaOrder(ulaAfterMiddleRemove),
				testCheckCanonicalTypeSetExcludes(resourceName, "ipv6private_cidr_blocks", []string{removedBeginningUla, removedMiddleUla}),
			),
		},
		// Step 24 - Remove a cidr from the end of the VCN's ipv6private_cidr_blocks field.
		{
			Config: coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterEndRemove),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					vcnIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: vcnIpv6PlanExpectation{
							privateFieldChanges: 1,
							privateRemovals:     1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_details.#", "11"),
				resource.TestCheckResourceAttr(resourceName, "byoipv6cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "byoipv6cidr_blocks", byoAfterEndRemove),
				checkByoOrder(byoAfterEndRemove),
				resource.TestCheckResourceAttr(resourceName, "ipv6private_cidr_blocks.#", "11"),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6private_cidr_blocks", ulaAfterEndRemove),
				checkUlaOrder(ulaAfterEndRemove),
				testCheckCanonicalTypeSetExcludes(resourceName, "ipv6private_cidr_blocks", removedUla),
			),
		},
		// Step 25 - Validate that terraform plan shows no new changes after ipv6private_cidr_blocks replacements and removals.
		{
			Config:             coreVcnPatchIpv6Config(byoipv6RangeId, byoAfterEndRemove, ulaAfterEndRemove),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},
	})
}

func coreVcnPatchTestByoIpv6Cidrs() []string {
	cidrs := make([]string, 0, 18)
	for i := 0; i < 18; i++ {
		cidrs = append(cidrs, fmt.Sprintf("2607:f590:0000:%04x::/64", 0x2200+i))
	}
	return cidrs
}

func coreVcnPatchTestUlaCidrs() []string {
	cidrs := make([]string, 0, 18)
	for i := 0; i < 18; i++ {
		cidrs = append(cidrs, fmt.Sprintf("fc00:%04x::/48", i))
	}
	return cidrs
}

func coreVcnPatchIpv6Config(byoipv6RangeId string, byoCidrs []string, ulaCidrs []string) string {
	var config strings.Builder

	config.WriteString(acctest.LegacyTestProviderConfig())
	config.WriteString(`
		resource "oci_core_virtual_network" "patch_vcn" {
			cidr_block                       = "10.0.0.0/16"
			compartment_id                   = "${var.compartment_id}"
			display_name                     = "patch-vcn"
			is_ipv6enabled                   = true
			is_oracle_gua_allocation_enabled = false
`)

	for _, cidr := range byoCidrs {
		fmt.Fprintf(&config, `
			byoipv6cidr_details {
				byoipv6range_id = %q
				ipv6cidr_block  = %q
			}
`, byoipv6RangeId, cidr)
	}

	if ulaCidrs != nil {
		config.WriteString("			ipv6private_cidr_blocks = [")
		for i, cidr := range ulaCidrs {
			if i > 0 {
				config.WriteString(", ")
			}
			fmt.Fprintf(&config, "%q", cidr)
		}
		config.WriteString("]\n")
	}

	config.WriteString(`
		}
`)

	return config.String()
}

type vcnIpv6PlanExpectation struct {
	byoFieldChanges          int
	byoAdditions             int
	byoRemovals              int
	byoReplacementGroups     int
	privateFieldChanges      int
	privateAdditions         int
	privateRemovals          int
	privateReplacementGroups int
}

type vcnIpv6PlanCheck struct {
	resourceAddress string
	expectation     vcnIpv6PlanExpectation
}

func (c vcnIpv6PlanCheck) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	change, err := findResourceChange(req.Plan, c.resourceAddress)
	if err != nil {
		resp.Error = err
		return
	}

	summary, err := summarizeVcnIpv6PlanChange(change)
	if err != nil {
		resp.Error = err
		return
	}

	if summary.byoFieldChanges != c.expectation.byoFieldChanges {
		resp.Error = fmt.Errorf("resource %s byoipv6 field changes = %d, want %d", c.resourceAddress, summary.byoFieldChanges, c.expectation.byoFieldChanges)
		return
	}
	if summary.byoAdditions != c.expectation.byoAdditions {
		resp.Error = fmt.Errorf("resource %s byoipv6 additions = %d, want %d", c.resourceAddress, summary.byoAdditions, c.expectation.byoAdditions)
		return
	}
	if summary.byoRemovals != c.expectation.byoRemovals {
		resp.Error = fmt.Errorf("resource %s byoipv6 removals = %d, want %d", c.resourceAddress, summary.byoRemovals, c.expectation.byoRemovals)
		return
	}
	if summary.byoReplacementGroups != c.expectation.byoReplacementGroups {
		resp.Error = fmt.Errorf("resource %s byoipv6 replacements = %d, want %d", c.resourceAddress, summary.byoReplacementGroups, c.expectation.byoReplacementGroups)
		return
	}
	if summary.privateFieldChanges != c.expectation.privateFieldChanges {
		resp.Error = fmt.Errorf("resource %s ipv6private_cidr_blocks field changes = %d, want %d", c.resourceAddress, summary.privateFieldChanges, c.expectation.privateFieldChanges)
		return
	}
	if summary.privateAdditions != c.expectation.privateAdditions {
		resp.Error = fmt.Errorf("resource %s ipv6private_cidr_blocks additions = %d, want %d", c.resourceAddress, summary.privateAdditions, c.expectation.privateAdditions)
		return
	}
	if summary.privateRemovals != c.expectation.privateRemovals {
		resp.Error = fmt.Errorf("resource %s ipv6private_cidr_blocks removals = %d, want %d", c.resourceAddress, summary.privateRemovals, c.expectation.privateRemovals)
		return
	}
	if summary.privateReplacementGroups != c.expectation.privateReplacementGroups {
		resp.Error = fmt.Errorf("resource %s ipv6private_cidr_blocks replacements = %d, want %d", c.resourceAddress, summary.privateReplacementGroups, c.expectation.privateReplacementGroups)
	}
}

type vcnIpv6PlanSummary struct {
	byoFieldChanges          int
	byoAdditions             int
	byoRemovals              int
	byoReplacementGroups     int
	privateFieldChanges      int
	privateAdditions         int
	privateRemovals          int
	privateReplacementGroups int
}

func summarizeVcnIpv6PlanChange(change *tfjson.ResourceChange) (vcnIpv6PlanSummary, error) {
	summary := vcnIpv6PlanSummary{}
	if change == nil || change.Change == nil {
		return summary, fmt.Errorf("resource change or nested change is nil")
	}

	before, err := planValueMap(change.Change.Before)
	if err != nil {
		return summary, err
	}
	after, err := planValueMap(change.Change.After)
	if err != nil {
		return summary, err
	}

	beforeByo, err := canonicalByoipv6CidrList(before["byoipv6cidr_details"])
	if err != nil {
		return summary, err
	}
	afterByo, err := canonicalByoipv6CidrList(after["byoipv6cidr_details"])
	if err != nil {
		return summary, err
	}

	removedByo := ipv6CidrsMissingFrom(beforeByo, afterByo)
	addedByo := ipv6CidrsMissingFrom(afterByo, beforeByo)
	if len(removedByo) > 0 || len(addedByo) > 0 {
		summary.byoFieldChanges = 1
	}
	summary.byoAdditions = len(addedByo)
	summary.byoRemovals = len(removedByo)
	if len(addedByo) < len(removedByo) {
		summary.byoReplacementGroups = len(addedByo)
	} else {
		summary.byoReplacementGroups = len(removedByo)
	}

	beforePrivate, err := canonicalIpv6CidrList(before["ipv6private_cidr_blocks"])
	if err != nil {
		return summary, err
	}
	afterPrivate, err := canonicalIpv6CidrList(after["ipv6private_cidr_blocks"])
	if err != nil {
		return summary, err
	}

	removedPrivate := ipv6CidrsMissingFrom(beforePrivate, afterPrivate)
	addedPrivate := ipv6CidrsMissingFrom(afterPrivate, beforePrivate)
	if len(removedPrivate) > 0 || len(addedPrivate) > 0 {
		summary.privateFieldChanges = 1
	}
	summary.privateAdditions = len(addedPrivate)
	summary.privateRemovals = len(removedPrivate)
	if len(addedPrivate) < len(removedPrivate) {
		summary.privateReplacementGroups = len(addedPrivate)
	} else {
		summary.privateReplacementGroups = len(removedPrivate)
	}

	return summary, nil
}

func canonicalByoipv6CidrList(value interface{}) ([]string, error) {
	if value == nil {
		return []string{}, nil
	}

	typed, ok := value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected byoipv6cidr_details list, got %T", value)
	}

	result := make([]string, 0, len(typed))
	for _, entry := range typed {
		entryMap, ok := entry.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected byoipv6cidr_details entry map, got %T", entry)
		}

		cidrValue, ok := entryMap["ipv6cidr_block"]
		if !ok {
			return nil, fmt.Errorf("missing ipv6cidr_block in byoipv6cidr_details entry")
		}

		cidr, ok := cidrValue.(string)
		if !ok {
			return nil, fmt.Errorf("expected ipv6cidr_block string, got %T", cidrValue)
		}

		canonical, err := canonicalIpv6CidrLiteral(cidr)
		if err != nil {
			return nil, err
		}
		result = append(result, canonical)
	}

	return result, nil
}

func testCheckCanonicalTypeSetExcludes(resourceName, attr string, excluded []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found in state: %s", resourceName)
		}

		excludedCanonical := make(map[string]struct{}, len(excluded))
		for _, block := range excluded {
			canonical, err := canonicalIpv6CidrLiteral(block)
			if err != nil {
				return err
			}
			excludedCanonical[canonical] = struct{}{}
		}

		prefix := attr + "."
		for key, value := range rs.Primary.Attributes {
			if !strings.HasPrefix(key, prefix) || key == attr+".#" {
				continue
			}

			canonical, err := canonicalIpv6CidrLiteral(value)
			if err != nil {
				return err
			}
			if _, ok := excludedCanonical[canonical]; ok {
				return fmt.Errorf("expected %s to exclude canonical block %s", attr, canonical)
			}
		}

		return nil
	}
}

// issue-routing-tag: core/virtualNetwork
func TestResourceCoreVirtualNetworkTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreVirtualNetworkTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreVirtualNetworkTestSuite))
}
