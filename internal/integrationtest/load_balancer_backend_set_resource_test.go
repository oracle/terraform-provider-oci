// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/loadbalancer"
	"github.com/stretchr/testify/suite"
)

type ResourceLoadBalancerBackendSetTestSuite struct {
	suite.Suite
	Providers           map[string]*schema.Provider
	Config              string
	ResourceName        string
	BackendResourceName string
}

func (s *ResourceLoadBalancerBackendSetTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + caCertificateVariableStr + privateKeyVariableStr + `
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
			// test Create
			{
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
						res, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						return err
					},
				),
			},
			// test Update
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// Update prop which is not force new
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
						url_path = "/abc"
					}
				
					lb_cookie_session_persistence_configuration {
                        cookie_name = "example_cookie"
						domain = "example.oracle.com"
		                max_age_in_seconds = 10
		                path = "/tmp/example"
						disable_fallback = true
					}
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "10"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.path", "/tmp/example"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.certificate_name", "tf_cert_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "4321"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/abc"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// switch from session persistence to LB cookie session persistence
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
				
					lb_cookie_session_persistence_configuration {
                        cookie_name = "example_cookie"
						domain = "example.oracle.com"
		                max_age_in_seconds = 10
		                path = "/tmp/example"
						disable_fallback = true
					}
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "10"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.path", "/tmp/example"),
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// switch from LB cookie session persistence back to session persistence
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "0"),
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// Update session persistence attribute
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
						disable_fallback = true
					}
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.cookie_name", "lb-session2"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.disable_fallback", "true"),
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// test Create backend and Update backendset (the two operations may happen concurrently)
			// This test is needed because CreateBackend and UpdateBackendSet both Update the same backend set
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
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}

				// Create a new backend
				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 8080
				}

				data "oci_load_balancer_backendsets" "t" {
					depends_on = ["oci_load_balancer_backend.t", "oci_load_balancer_backendset.t"]
					load_balancer_id = "${oci_load_balancer.t.id}"
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			// test ForceNew backend and Update backendset (the operations may happen concurrently)
			// This test is needed because DeleteBackend, CreateBackend and UpdateBackendSet all Update the same backend set
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
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}

				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 80
				}

				data "oci_load_balancer_backendsets" "t" {
					depends_on = ["oci_load_balancer_backend.t", "oci_load_balancer_backendset.t"]
					load_balancer_id = "${oci_load_balancer.t.id}"
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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

// todo: BackendSet health_checker appears to be missing 3 properties
// https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/HealthCheckerDetails
func (s *ResourceLoadBalancerBackendSetTestSuite) TestAccResourceLoadBalancerBackendSetLBCookie_basic() {
	var res, res2 string
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test Create
			{
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
						res, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						return err
					},
				),
			},
			// test add session persistence without optionals
			{
				Config: s.Config + `
				resource "oci_load_balancer_backendset" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					name = "-tf-backend-set"
					policy = "IP_HASH"
					health_checker {
						interval_ms = 30000
						port = 1234
						protocol = "TCP"
						response_body_regex = ".*"
						url_path = "/"
					}

					lb_cookie_session_persistence_configuration {
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "IP_HASH"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "1"),
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// test Update
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

                    lb_cookie_session_persistence_configuration {
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "1"),
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
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

					lb_cookie_session_persistence_configuration {
                        cookie_name = "example_cookie"
						domain = "example.oracle.com"
		                max_age_in_seconds = 10
		                path = "/tmp/example"
						disable_fallback = true
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "IP_HASH"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "10"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.path", "/tmp/example"),
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
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
				
					lb_cookie_session_persistence_configuration {
                        cookie_name = "example_cookie"
						domain = "example.oracle.com"
		                max_age_in_seconds = 10
		                path = "/tmp/example"
						disable_fallback = true
					}
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.disable_fallback", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "10"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.path", "/tmp/example"),
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// test switching from LB cookie to session cookie
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
						cookie_name = "lb-session1"
						disable_fallback = true
					}
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.cookie_name", "lb-session1"),
					resource.TestCheckResourceAttr(s.ResourceName, "session_persistence_configuration.0.disable_fallback", "true"),
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// test switching back to LB session cookie
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
				
					lb_cookie_session_persistence_configuration {
                        cookie_name = "example_cookie"
						domain = "example.oracle.com"
		                max_age_in_seconds = 10
					}
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "10"),
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// Update non force-new property without Update lb cookie session
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
						url_path = "/abc"
					}
				
					lb_cookie_session_persistence_configuration {
                        cookie_name = "example_cookie"
						domain = "example.oracle.com"
		                max_age_in_seconds = 10
					}
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "10"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.certificate_name", "tf_cert_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "4321"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/abc"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// Update lb cookie session persistence attribute
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
						url_path = "/abc"
					}
				
					lb_cookie_session_persistence_configuration {
                        cookie_name = "example_cookie"
						domain = "example.oracle.com"
		                max_age_in_seconds = 20
					}
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policy", "LEAST_CONNECTIONS"),
					resource.TestCheckResourceAttr(s.ResourceName, "backend.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.cookie_name", "example_cookie"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.domain", "example.oracle.com"),
					resource.TestCheckResourceAttr(s.ResourceName, "lb_cookie_session_persistence_configuration.0.max_age_in_seconds", "20"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.certificate_name", "tf_cert_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.interval_ms", "29999"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.port", "4321"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.url_path", "/abc"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.retries", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.return_code", "200"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.response_body_regex", ".*"),
					resource.TestCheckResourceAttr(s.ResourceName, "health_checker.0.timeout_in_millis", "3000"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// test Create backend and Update backendset (the two operations may happen concurrently)
			// This test is needed because CreateBackend and UpdateBackendSet both Update the same backend set
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

					lb_cookie_session_persistence_configuration {
                        cookie_name = "example_cookie"
						domain = "example.oracle.com"
		                max_age_in_seconds = 10
		                path = "/tmp/example"
						disable_fallback = true
					}

					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}

				// Create a new backend
				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 8080
				}

				data "oci_load_balancer_backendsets" "t" {
					depends_on = ["oci_load_balancer_backend.t", "oci_load_balancer_backendset.t"]
					load_balancer_id = "${oci_load_balancer.t.id}"
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			// test ForceNew backend and Update backendset (the operations may happen concurrently)
			// This test is needed because DeleteBackend, CreateBackend and UpdateBackendSet all Update the same backend set
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

					lb_cookie_session_persistence_configuration {
                        cookie_name = "example_cookie"
						domain = "example.oracle.com"
		                max_age_in_seconds = 10
		                path = "/tmp/example"
						disable_fallback = true
					}

					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}

				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 80
				}

				data "oci_load_balancer_backendsets" "t" {
					depends_on = ["oci_load_balancer_backend.t", "oci_load_balancer_backendset.t"]
					load_balancer_id = "${oci_load_balancer.t.id}"
				}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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

// issue-routing-tag: load_balancer/default
func TestResourceLoadBalancerBackendSetTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceLoadBalancerBackendSetTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceLoadBalancerBackendSetTestSuite))
}
