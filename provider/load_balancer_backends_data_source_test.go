// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
				Config: config + `
				data "oci_load_balancer_backends" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name  = "${oci_load_balancer_backendset.t.name}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					TestCheckResourceAttributesEqual(resourceName, "load_balancer_id", "oci_load_balancer.t", "id"),
					TestCheckResourceAttributesEqual(resourceName, "backendset_name", "oci_load_balancer_backendset.t", "name"),
					resource.TestCheckResourceAttr(resourceName, "backends.#", "1"),
					TestCheckResourceAttributesEqual(resourceName, "backends.0.ip_address", "oci_load_balancer_backend.t", "ip_address"),
					TestCheckResourceAttributesEqual(resourceName, "backends.0.port", "oci_load_balancer_backend.t", "port"),
					TestCheckResourceAttributesEqual(resourceName, "backends.0.backup", "oci_load_balancer_backend.t", "backup"),
					TestCheckResourceAttributesEqual(resourceName, "backends.0.drain", "oci_load_balancer_backend.t", "drain"),
					TestCheckResourceAttributesEqual(resourceName, "backends.0.offline", "oci_load_balancer_backend.t", "offline"),
					TestCheckResourceAttributesEqual(resourceName, "backends.0.weight", "oci_load_balancer_backend.t", "weight"),
					resource.TestCheckResourceAttrSet(resourceName, "backends.0.name"),
					validateBackendName(resourceName),
				),
			},
			// Client-side filtering.
			{
				Config: config + `
				resource "oci_load_balancer_backend" "u" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "5.6.7.8"
					port = 80
					backup = false
					drain = false
					offline = false
					weight = 1
				}
				
				data "oci_load_balancer_backends" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name  = "${oci_load_balancer_backendset.t.name}"
					filter {
						name = "ip_address"
						values = ["1.2.3.4"]
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					TestCheckResourceAttributesEqual(resourceName, "load_balancer_id", "oci_load_balancer.t", "id"),
					TestCheckResourceAttributesEqual(resourceName, "backendset_name", "oci_load_balancer_backendset.t", "name"),
					resource.TestCheckResourceAttr(resourceName, "backends.#", "1"),
					TestCheckResourceAttributesEqual(resourceName, "backends.0.ip_address", "oci_load_balancer_backend.t", "ip_address"),
					TestCheckResourceAttributesEqual(resourceName, "backends.0.port", "oci_load_balancer_backend.t", "port"),
					validateBackendName(resourceName),
				),
			},
		},
	})
}

func validateBackendName(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ipAddress, err := fromInstanceState(s, resourceName, "backends.0.ip_address")
		if err != nil {
			return err
		}
		port, err := fromInstanceState(s, resourceName, "backends.0.port")
		if err != nil {
			return err
		}
		actualName, err := fromInstanceState(s, resourceName, "backends.0.name")
		if err != nil {
			return err
		}
		expectedName := ipAddress + ":" + port
		if expectedName != actualName {
			return fmt.Errorf("invalid name: expected %s, got %s", expectedName, actualName)
		}
		return nil
	}
}
