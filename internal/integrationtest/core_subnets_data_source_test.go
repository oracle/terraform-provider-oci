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

type DatasourceCoreSubnetTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
}

func (s *DatasourceCoreSubnetTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "vcn1" {
		cidr_block     = "10.1.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name   = "vcn1"
		dns_label = "vnc1"
	}
	
	resource "oci_core_subnet" "s" {	
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.vcn1.id}"
		count = "${length(data.oci_identity_availability_domains.ADs.availability_domains)}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[count.index], "name")}"
		cidr_block = "10.1.2${count.index}.0/24"
		display_name = "subnet${count.index}"
		dns_label = "subnet${count.index}"
		security_list_ids = ["${oci_core_virtual_network.vcn1.default_security_list_id}"]
		route_table_id = "${oci_core_virtual_network.vcn1.default_route_table_id}"
		dhcp_options_id = "${oci_core_virtual_network.vcn1.default_dhcp_options_id}"
	}`

	s.ResourceName = "data.oci_core_subnets.s"
}

func (s *DatasourceCoreSubnetTestSuite) TestAccDatasourceCoreSubnet_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			// check properties on the subnet created in AD2
			{
				Config: s.Config + `
				data "oci_core_subnets" "s" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.vcn1.id}"
					filter {
						name   = "availability_domain"
						values = ["${oci_core_subnet.s.*.availability_domain[1]}"]
					}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.availability_domain", "oci_core_subnet.s.1", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.cidr_block", "10.1.21.0/24"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.display_name", "subnet1"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.dns_label", "subnet1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.id", "oci_core_subnet.s.1", "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.prohibit_public_ip_on_vnic", "false"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.route_table_id", "oci_core_subnet.s.1", "route_table_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.security_list_ids.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.state", string(core.SubnetLifecycleStateAvailable)),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.time_created", "oci_core_subnet.s.1", "time_created"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.vcn_id", "oci_core_subnet.s.1", "vcn_id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.dhcp_options_id", "oci_core_subnet.s.1", "dhcp_options_id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.virtual_router_ip", "oci_core_subnet.s.1", "virtual_router_ip"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.virtual_router_mac", "oci_core_subnet.s.1", "virtual_router_mac"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.subnet_domain_name", "oci_core_subnet.s.1", "subnet_domain_name"),
				),
			},
			// Server-side filtering tests.
			{
				Config: s.Config + `
				data "oci_core_subnets" "s" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.vcn1.id}"
					display_name = "${oci_core_subnet.s.2.display_name}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.id", "oci_core_subnet.s.2", "id"),
				),
			},
			{
				Config: s.Config + `
				data "oci_core_subnets" "s" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.vcn1.id}"
					state = "${oci_core_subnet.s.0.state}" # Adding implicit dependency
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.#", "3"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: core/virtualNetwork
func TestDatasourceCoreSubnetTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreSubnetTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreSubnetTestSuite))
}
