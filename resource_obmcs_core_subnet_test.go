// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"regexp"
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceCoreSubnetCreate_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig() + `
		data "oci_identity_availability_domains" "ADs" {
			compartment_id = "${var.compartment_id}"
		}
		
		resource "oci_core_virtual_network" "t" {
			cidr_block     = "10.0.0.0/16"
			compartment_id = "${var.compartment_id}"
			display_name   = "network_name"
		}`

	resourceName := "oci_core_subnet.s"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: config + `
				resource "oci_core_subnet" "s" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
					route_table_id = "${oci_core_virtual_network.t.default_route_table_id}"
					dhcp_options_id = "${oci_core_virtual_network.t.default_dhcp_options_id}"
					cidr_block = "10.0.2.0/24"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),
					resource.TestMatchResourceAttr(resourceName, "compartment_id", regexp.MustCompile("ocid1\\.compartment\\.oc1\\..*")),
					resource.TestMatchResourceAttr(resourceName, "vcn_id", regexp.MustCompile("ocid1\\.vcn\\.oc1\\..*")),
					resource.TestMatchResourceAttr(resourceName, "dhcp_options_id", regexp.MustCompile("ocid1\\.dhcpoptions\\.oc1\\..*")),
					resource.TestMatchResourceAttr(resourceName, "route_table_id", regexp.MustCompile("ocid1\\.routetable\\.oc1\\..*")),
					resource.TestCheckResourceAttr(resourceName, "security_list_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.2.0/24"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", ""),
					resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
					resource.TestMatchResourceAttr(resourceName, "id", regexp.MustCompile("ocid1\\.subnet\\.oc1\\..*")),
					resource.TestCheckResourceAttr(resourceName, "state", baremetal.ResourceAvailable),
				),
			},
			// verify update
			{
				Config: config + `
				resource "oci_core_subnet" "s" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
					route_table_id = "${oci_core_virtual_network.t.default_route_table_id}"
					dhcp_options_id = "${oci_core_virtual_network.t.default_dhcp_options_id}"
					cidr_block = "10.0.2.0/24"
					display_name = "-tf-subnet"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "-tf-subnet"),
				),
			},
		},
	})
}
