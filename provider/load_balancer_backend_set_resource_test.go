// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/loadbalancer"
	"github.com/stretchr/testify/suite"
)

type ResourceLoadBalancerBackendSetTestSuite struct {
	suite.Suite
	Providers           map[string]terraform.ResourceProvider
	Config              string
	ResourceName        string
	BackendResourceName string
}

func (s *ResourceLoadBalancerBackendSetTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
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
	s.BackendResourceName = "oci_load_balancer_backend.t"
}

// todo: BackendSet health_checker appears to be missing 3 properties
// https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/HealthCheckerDetails
func (s *ResourceLoadBalancerBackendSetTestSuite) TestAccResourceLoadBalancerBackendSet_basic() {
	var res, res2 string
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "30000"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "1234"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res, err = fromInstanceState(ts, s.ResourceName, "name")
						return err
					},
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "4321"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "IP_HASH"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.cookie_name", "lb-session1"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "4321"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
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
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.cookie_name", "lb-session2"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.disable_fallback", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.certificate_name", "tf_cert_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "4321"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// test create backend and update backendset (the two operations may happen concurrently)
			// This test is needed because CreateBackend and UpdateBackendSet both update the same backend set
			// resource. Test that both changes are applied sequentially.
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "LEAST_CONNECTIONS"
					health_checker {
						interval_ms = 29999
						port = 8080
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}

					session_persistence_configuration {
						cookie_name = "lb-session2"
						disable_fallback = false
					}

					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.id}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}

				// Create a new backend
				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.id}"
					ip_address = "1.2.3.4"
					port = 8080
				}

				data "oci_load_balancer_backendsets" "t" {
					depends_on = ["oci_load_balancer_backend.t", "oci_load_balancer_backendset.t"]
					load_balancer_id = "${oci_load_balancer.t.id}"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					// The state file could show either 0 or 1 backends in backend_set; depending on the order of operations.
					// If UpdateBackendSet happens first, then you will see 0. If CreateBackend happens first, then you will see 1.
					//resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "1"),
					resource.TestCheckResourceAttr("data.oci_load_balancer_backendsets.t", "backendsets.#", "1"),
					resource.TestCheckResourceAttr("data.oci_load_balancer_backendsets.t", "backendsets.0.backend.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "8080"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					resource.TestCheckResourceAttrSet(s.BackendResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "backendset_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "port", "8080"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "backup", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "drain", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "offline", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "weight", "1"),
					resource.TestCheckResourceAttrSet(s.BackendResourceName, "name"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
				),
				ExpectNonEmptyPlan: true,
			},
			// test ForceNew backend and update backendset (the operations may happen concurrently)
			// This test is needed because DeleteBackend, CreateBackend and UpdateBackendSet all update the same backend set
			// resource. Test that all changes are applied sequentially.
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "LEAST_CONNECTIONS"
					health_checker {
						interval_ms = 29999
						port = 80
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}

					session_persistence_configuration {
						cookie_name = "lb-session2"
						disable_fallback = false
					}

					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.id}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}

				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.id}"
					ip_address = "1.2.3.4"
					port = 80
				}

				data "oci_load_balancer_backendsets" "t" {
					depends_on = ["oci_load_balancer_backend.t", "oci_load_balancer_backendset.t"]
					load_balancer_id = "${oci_load_balancer.t.id}"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					// The state file could show either 0 or 1 backends in backend_set; depending on the order of operations.
					// If UpdateBackendSet happens first, then you will see 0. If CreateBackend happens first, then you will see 1.
					//resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "1"),
					resource.TestCheckResourceAttr("data.oci_load_balancer_backendsets.t", "backendsets.#", "1"),
					resource.TestCheckResourceAttr("data.oci_load_balancer_backendsets.t", "backendsets.0.backend.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "80"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					resource.TestCheckResourceAttrSet(s.BackendResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "backendset_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "port", "80"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "backup", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "drain", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "offline", "false"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "weight", "1"),
					resource.TestCheckResourceAttrSet(s.BackendResourceName, "name"),
					resource.TestCheckResourceAttr(s.BackendResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestResourceLoadBalancerBackendSetTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerBackendSetTestSuite))
}
