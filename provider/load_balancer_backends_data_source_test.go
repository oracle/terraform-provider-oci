// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDatasourceLoadBalancerBackends_basic(t *testing.T) {
	providers := testAccProviders
	config := legacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
	}
	
	resource "oci_core_subnet" "t" {
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block          = "10.0.0.0/24"
		display_name        = "-tf-subnet"
	}
	
	resource "oci_load_balancer" "t" {
		shape = "100Mbps"
		compartment_id = "${var.compartment_id}"
		subnet_ids = ["${oci_core_subnet.t.id}"]
		display_name = "-tf-lb"
		is_private = true
	}
	
	resource "oci_load_balancer_backendset" "t" {
		load_balancer_id = "${oci_load_balancer.t.id}"
		name = "-tf-backend-set"
		policy = "ROUND_ROBIN"
		health_checker {
			interval_ms = 30000
			port = 1234
			protocol = "TCP"
			response_body_regex = ".*"
			url_path = "/"
		}
	}
	
	resource "oci_load_balancer_backend" "t" {
		load_balancer_id = "${oci_load_balancer.t.id}"
		backendset_name = "${oci_load_balancer_backendset.t.name}"
		ip_address = "1.2.3.4"
		port = 8080
		backup = false
		drain = false
		offline = false
		weight = 1
	}
	
	data "oci_load_balancer_backends" "t" {
		load_balancer_id = "${oci_load_balancer.t.id}"
		backendset_name  = "${oci_load_balancer_backendset.t.name}"
	}`

	resourceName := "data.oci_load_balancer_backends.t"

	resource.Test(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config,
			},
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttrSet(resourceName, "backendset_name"),
					resource.TestCheckResourceAttr(resourceName, "backends.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backends.0.ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr(resourceName, "backends.0.port", "8080"),
					resource.TestCheckResourceAttr(resourceName, "backends.0.backup", "false"),
					resource.TestCheckResourceAttr(resourceName, "backends.0.drain", "false"),
					resource.TestCheckResourceAttr(resourceName, "backends.0.offline", "false"),
					resource.TestCheckResourceAttr(resourceName, "backends.0.weight", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "backends.0.name"),
				),
			},
		},
	})
}
