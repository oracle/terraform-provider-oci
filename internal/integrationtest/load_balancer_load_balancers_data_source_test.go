// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// issue-routing-tag: load_balancer/default
func TestAccDatasourceLoadBalancerLB_basic(t *testing.T) {
	httpreplay.SetScenario("TestAccDatasourceLoadBalancerLB_basic")
	defer httpreplay.SaveScenario()
	providers := acctest.TestAccProviders
	config := acctest.LegacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
	}
	
	resource "oci_core_subnet" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		route_table_id = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dhcp_options_id = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block = "10.0.0.0/24"
		display_name = "-tf-subnet"
	}
	
	resource "oci_load_balancer" "t" {
		compartment_id = "${var.compartment_id}"
		subnet_ids = ["${oci_core_subnet.t.id}"]
		shape = "100Mbps"
		display_name = "-tf-lb"
		is_private = true
	}
	
	data "oci_load_balancers" "t" {
		compartment_id = "${var.compartment_id}"
		filter {
			name = "display_name"
			values = ["-tf-lb"]
		}
	}`
	resourceName := "data.oci_load_balancers.t"

	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { acctest.PreCheck(t) },
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
			},
			{
				Config: config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancers.#"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancers.#"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.shape", "100Mbps"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.display_name", "-tf-lb"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.is_private", "true"),
				),
			},
		},
	})
}
