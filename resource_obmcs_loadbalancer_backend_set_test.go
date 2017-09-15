package main

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

type ResourceLoadBalancerBackendSetTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceLoadBalancerBackendSetTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
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
		ca_certificate = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"
		certificate_name = "tf_cert_name"
		private_key = "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----"
		public_certificate = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"
	}`
	s.ResourceName = "oci_load_balancer_backendset.t"
}

// todo: BackendSet health_checker appears to be missing 3 properties
// https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/HealthCheckerDetails
func (s *ResourceLoadBalancerBackendSetTestSuite) TestAccResourceLoadBalancerBackendSet_basic() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
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
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "30000"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "1234"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
				),
			},
			// test update
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "LEAST_CONNECTIONS"
					health_checker {
						interval_ms = 29999
						port = 4321
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "4321"),
				),
			},
			// test add session persistence
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "IP_HASH"
					health_checker {
						interval_ms = 29999
						port = 4321
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}

					session_persistence_configuration {
						cookie_name = "lb-session1"
						disable_fallback = true
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "IP_HASH"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.cookie_name", "lb-session1"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.disable_fallback", "true"),
				),
			},
			// test add ssl
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "LEAST_CONNECTIONS"
					health_checker {
						interval_ms = 29999
						port = 4321
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.certificate_name", "tf_cert_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
				),
			},
		},
	})
}

func TestResourceLoadBalancerBackendSetTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerBackendSetTestSuite))
}
