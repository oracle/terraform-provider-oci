// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVirtualNetworkTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreVirtualNetworkTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig()
	s.ResourceName = "oci_core_virtual_network.t"
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestAccResourceCoreVirtualNetwork_basic() {
	var resId string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "dns_label"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_virtual_network.t", "id")
						return err
					},
				),
			},
			// test update
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-vcn"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
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
			// test a destructive update results in a new resource
			{
				Config: s.Config + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/24"
						compartment_id = "${var.compartment_id}"
						dns_label= "MyTestDNSLabel"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_route_table_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "default_security_list_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VcnLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "dns_label", "mytestdnslabel"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_domain_name"),
					func(s *terraform.State) (err error) {
						resId2, err := fromInstanceState(s, "oci_core_virtual_network.t", "id")
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

func TestResourceCoreVirtualNetworkTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVirtualNetworkTestSuite))
}
