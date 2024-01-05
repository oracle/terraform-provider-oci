// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVirtualNetworkTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
	Token        string
	TokenFn      acctest.TokenFn
}

func (s *DatasourceCoreVirtualNetworkTestSuite) SetupTest() {
	s.Token, s.TokenFn = acctest.TokenizeWithHttpReplay("vcn")
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + s.TokenFn(`
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "{{.token}}"
		cidr_block = "10.0.0.0/16"
		dns_label = "vcn1"
	}
	resource "oci_core_virtual_network" "u" {
		compartment_id = "${var.compartment_id}"
		display_name = "{{.otherToken}}"
		cidr_block = "10.0.0.0/16"
		dns_label = "vcn2"
	}`, map[string]string{"otherToken": s.Token + "-2"})
	s.ResourceName = "data.oci_core_virtual_networks.t"
}

func (s *DatasourceCoreVirtualNetworkTestSuite) TestAccDatasourceCoreVirtualNetwork_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + s.TokenFn(`
					data "oci_core_virtual_networks" "t" {
						compartment_id = "${oci_core_virtual_network.t.compartment_id}"
						filter {
							name = "display_name"
							values = ["{{.token}}"]
						}
					}`, nil),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.display_name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.cidr_block", "10.0.0.0/16"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "virtual_networks.0.id", "oci_core_virtual_network.t", "id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "virtual_networks.0.default_route_table_id", "oci_core_virtual_network.t", "default_route_table_id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "virtual_networks.0.default_security_list_id", "oci_core_virtual_network.t", "default_security_list_id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "virtual_networks.0.default_dhcp_options_id", "oci_core_virtual_network.t", "default_dhcp_options_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.dns_label", "vcn1"),
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.state", string(core.VcnLifecycleStateAvailable)),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "virtual_networks.0.time_created", "oci_core_virtual_network.t", "time_created"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "virtual_networks.0.vcn_domain_name", "oci_core_virtual_network.t", "vcn_domain_name"),
				),
			},
			// Server-side filtering tests.
			{
				Config: s.Config + `
					data "oci_core_virtual_networks" "t" {
						compartment_id = "${oci_core_virtual_network.u.compartment_id}"
						display_name = "${oci_core_virtual_network.u.display_name}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "virtual_networks.0.id", "oci_core_virtual_network.u", "id"),
				),
			},
			{
				Config: s.Config + s.TokenFn(`
					data "oci_core_virtual_networks" "t" {
						compartment_id = "${oci_core_virtual_network.t.compartment_id}"
						display_name = "does-not-exit"
					}`, nil),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.#", "0"),
				),
			},
		},
	},
	)

}

// issue-routing-tag: core/virtualNetwork
func TestDatasourceCoreVirtualNetworkTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreVirtualNetworkTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreVirtualNetworkTestSuite))
}
