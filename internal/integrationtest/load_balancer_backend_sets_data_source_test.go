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
func TestAccDatasourceLoadBalancerBackendsets_basic(t *testing.T) {
	httpreplay.SetScenario("TestAccDatasourceLoadBalancerBackendsets_basic")
	defer httpreplay.SaveScenario()
	providers := acctest.TestAccProviders
	config := acctest.LegacyTestProviderConfig() + caCertificateVariableStr + privateKeyVariableStr + `
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

	resource "oci_load_balancer_certificate" "t" {
		load_balancer_id = "${oci_load_balancer.t.id}"
		ca_certificate = "${var.ca_certificate_value}"
		certificate_name = "tf_cert_name"
		private_key = "${var.private_key_value}"
		public_certificate = "${var.ca_certificate_value}"
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

		session_persistence_configuration {
			cookie_name = "lb-session2"
			disable_fallback = false
		}

		ssl_configuration {
			certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
			verify_depth = 6
			verify_peer_certificate = false
		}
	}
	
	data "oci_load_balancer_backendsets" "t" {
		load_balancer_id = "${oci_load_balancer.t.id}"
	}`

	resourceName := "data.oci_load_balancer_backendsets.t"

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
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.policy", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.health_checker.0.interval_ms", "30000"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.health_checker.0.port", "1234"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.session_persistence_configuration.0.cookie_name", "lb-session2"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.session_persistence_configuration.0.disable_fallback", "false"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.ssl_configuration.0.certificate_name", "tf_cert_name"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.ssl_configuration.0.verify_peer_certificate", "false"),
				),
			},
		},
	})
}
