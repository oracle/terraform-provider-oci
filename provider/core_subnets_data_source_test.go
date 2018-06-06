// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreSubnetTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreSubnetTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
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
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_core_subnets" "s" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.vcn1.id}"
					filter {
						name   = "availability_domain"
						values = ["${oci_core_subnet.s.*.availability_domain[1]}"]
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.#", "1"),
					TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.availability_domain", "oci_core_subnet.s.1", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.cidr_block", "10.1.21.0/24"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.display_name", "subnet1"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.dns_label", "subnet1"),
					TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.id", "oci_core_subnet.s.1", "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.prohibit_public_ip_on_vnic", "false"),
					TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.route_table_id", "oci_core_subnet.s.1", "route_table_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.security_list_ids.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.0.state", string(core.SubnetLifecycleStateAvailable)),
					TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.time_created", "oci_core_subnet.s.1", "time_created"),
					TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.vcn_id", "oci_core_subnet.s.1", "vcn_id"),
					TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.dhcp_options_id", "oci_core_subnet.s.1", "dhcp_options_id"),
					TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.virtual_router_ip", "oci_core_subnet.s.1", "virtual_router_ip"),
					TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.virtual_router_mac", "oci_core_subnet.s.1", "virtual_router_mac"),
					TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.subnet_domain_name", "oci_core_subnet.s.1", "subnet_domain_name"),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.#", "1"),
					TestCheckResourceAttributesEqual(s.ResourceName, "subnets.0.id", "oci_core_subnet.s.2", "id"),
				),
			},
			{
				Config: s.Config + `
				data "oci_core_subnets" "s" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.vcn1.id}"
					state = "${oci_core_subnet.s.0.state}" # Adding implicit dependency
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnets.#", "3"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreSubnetTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreSubnetTestSuite))
}
