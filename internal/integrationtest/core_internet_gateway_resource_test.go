// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreInternetGatewayTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreInternetGatewayTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
	}
	`

	s.ResourceName = "oci_core_internet_gateway.t"
}

func (s *ResourceCoreInternetGatewayTestSuite) TestAccResourceCoreInternetGateway_basic() {
	var resId, resId2 string
	compartmentId := acctest.GetCompartmentIDForLegacyTests()
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: s.Config +
					`resource "oci_core_internet_gateway" "t" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.InternetGatewayLifecycleStateAvailable)),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_core_internet_gateway.t", "id")
						return err
					},
				),
			},
			// verify Update
			{
				Config: s.Config + `
					resource "oci_core_internet_gateway" "t" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						display_name = "-tf-internet-gateway"
						enabled = false
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-internet-gateway"),
					resource.TestCheckResourceAttr(s.ResourceName, "enabled", "false"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.InternetGatewayLifecycleStateAvailable)),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, "oci_core_internet_gateway.t", "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						resId = resId2
						return err
					},
				),
			},
			// verify destructive Update
			{
				Config: acctest.LegacyTestProviderConfig() + `
					resource "oci_core_virtual_network" "t2" {
						compartment_id = "${var.compartment_id}"
						cidr_block = "10.0.0.0/16"
						display_name = "-tf-vcn-igate-upd"
					}
					resource "oci_core_internet_gateway" "t" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t2.id}"
						display_name = "-tf-internet-gateway"
						enabled = false
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-internet-gateway"),
					resource.TestCheckResourceAttr(s.ResourceName, "enabled", "false"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.InternetGatewayLifecycleStateAvailable)),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, "oci_core_internet_gateway.t", "id")
						if resId == resId2 {
							return fmt.Errorf("expected resource to be recreated but was not")
						}
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: core/virtualNetwork
func TestResourceCoreInternetGatewayTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreInternetGatewayTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreInternetGatewayTestSuite))
}
