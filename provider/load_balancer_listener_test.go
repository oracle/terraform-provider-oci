// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ListenerRequiredOnlyResource = ListenerResourceDependencies + `
resource "oci_load_balancer_listener" "test_listener" {
	#Required
	default_backend_set_name = "${oci_load_balancer_backendset.test_backend_set.name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.listener_name}"
	port = "${var.listener_port}"
	protocol = "${var.listener_protocol}"
}
`

	ListenerResourceConfig = ListenerResourceDependencies + `
resource "oci_load_balancer_listener" "test_listener" {
	#Required
	default_backend_set_name = "${oci_load_balancer_backendset.test_backend_set.name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.listener_name}"
	port = "${var.listener_port}"
	protocol = "${var.listener_protocol}"

	#Optional
	connection_configuration {
		#Required
		idle_timeout_in_seconds = "${var.listener_connection_configuration_idle_timeout_in_seconds}"
	}
	hostname_names = ["${oci_load_balancer_hostname.test_hostname.name}"]
	path_route_set_name = "${oci_load_balancer_path_route_set.test_path_route_set.name}"
	ssl_configuration {
		#Required
		certificate_name = "${oci_load_balancer_certificate.test_certificate.certificate_name}"

		#Optional
		verify_depth = "${var.listener_ssl_configuration_verify_depth}"
		verify_peer_certificate = "${var.listener_ssl_configuration_verify_peer_certificate}"
	}
}
`

	ListenerWithTwoHostnamesResourceConfig = ListenerResourceDependencies + `
resource "oci_load_balancer_listener" "test_listener" {
	#Required
	default_backend_set_name = "${oci_load_balancer_backendset.test_backend_set.name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.listener_name}"
	port = "${var.listener_port}"
	protocol = "${var.listener_protocol}"

	#Optional
	connection_configuration {
		#Required
		idle_timeout_in_seconds = "${var.listener_connection_configuration_idle_timeout_in_seconds}"
	}
	hostname_names = ["${oci_load_balancer_hostname.test_hostname.name}", "${oci_load_balancer_hostname.test_hostname2.name}"]
	path_route_set_name = "${oci_load_balancer_path_route_set.test_path_route_set.name}"
	ssl_configuration {
		#Required
		certificate_name = "${oci_load_balancer_certificate.test_certificate.certificate_name}"

		#Optional
		verify_depth = "${var.listener_ssl_configuration_verify_depth}"
		verify_peer_certificate = "${var.listener_ssl_configuration_verify_peer_certificate}"
	}
}
`
	ListenerPropertyVariables = `
variable "listener_connection_configuration_idle_timeout_in_seconds" { default = 10 }
variable "listener_default_backend_set_name" { default = "example_backend_set" }
variable "listener_name" { default = "mylistener" }
variable "listener_port" { default = 10 }
variable "listener_protocol" { default = "HTTP" }
variable "listener_ssl_configuration_certificate_name" { default = "example_certificate_bundle" }
variable "listener_ssl_configuration_verify_depth" { default = 10 }
variable "listener_ssl_configuration_verify_peer_certificate" { default = false }

`
	ListenerResourceDependencies = PathRouteSetPropertyVariables + HostnamePropertyVariables + `
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
	
	resource "oci_load_balancer_load_balancer" "test_load_balancer" {
		shape = "100Mbps"
		compartment_id = "${var.compartment_id}"
		subnet_ids = ["${oci_core_subnet.t.id}"]
		display_name = "-tf-lb"
		is_private = true
	}
	
	resource "oci_load_balancer_backendset" "test_backend_set" {
		load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
		name = "example_backend_set"
		policy = "ROUND_ROBIN"
		health_checker {
			interval_ms = 30000
			port = 1234
			protocol = "HTTP"
			response_body_regex = ".*"
			url_path = "/"
		}
	}
	
	resource "oci_load_balancer_certificate" "test_certificate" {
		load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
		ca_certificate = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"
		certificate_name = "example_certificate_bundle"
		private_key = "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----"
		public_certificate = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"
	}

	resource "oci_load_balancer_path_route_set" "test_path_route_set" {
		#Required
		load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
		name = "${var.path_route_set_name}"
		path_routes {
			#Required
			backend_set_name = "${oci_load_balancer_backendset.test_backend_set.name}"
			path = "${var.path_route_set_path_routes_path}"
			path_match_type {
				#Required
				match_type = "${var.path_route_set_path_routes_path_match_type_match_type}"
			}
		}
	}

	resource "oci_load_balancer_hostname" "test_hostname" {
		#Required
		hostname = "${var.hostname_hostname}"
		load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
		name = "${var.hostname_name}"
	}

	resource "oci_load_balancer_hostname" "test_hostname2" {
		#Required
		hostname = "${var.hostname_hostname}2"
		load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
		name = "${var.hostname_name}2"
	}
`
)

func TestLoadBalancerListenerResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_listener.test_listener"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + ListenerPropertyVariables + compartmentIdVariableStr + ListenerRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", "example_backend_set"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "mylistener"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ListenerResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + ListenerPropertyVariables + compartmentIdVariableStr + ListenerWithTwoHostnamesResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.idle_timeout_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", "example_backend_set"),
					resource.TestCheckResourceAttr(resourceName, "hostname_names.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "hostname_names.0", "example_hostname_001"),
					resource.TestCheckResourceAttr(resourceName, "hostname_names.1", "example_hostname_0012"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "mylistener"),
					resource.TestCheckResourceAttrSet(resourceName, "path_route_set_name"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "10"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "false"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "listener_connection_configuration_idle_timeout_in_seconds" { default = 10 }
variable "listener_default_backend_set_name" { default = "example_backend_set" }
variable "listener_hostname_names" { default = [] }
variable "listener_name" { default = "mylistener" }
variable "listener_port" { default = 10 }
variable "listener_protocol" { default = "HTTP" }
variable "listener_ssl_configuration_certificate_name" { default = "example_certificate_bundle" }
variable "listener_ssl_configuration_verify_depth" { default = 11 }
variable "listener_ssl_configuration_verify_peer_certificate" { default = true }

                ` + compartmentIdVariableStr + ListenerResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_configuration.0.idle_timeout_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", "example_backend_set"),
					resource.TestCheckResourceAttr(resourceName, "hostname_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostname_names.0", "example_hostname_001"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "mylistener"),
					resource.TestCheckResourceAttrSet(resourceName, "path_route_set_name"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", "11"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", "true"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		},
	})
}
