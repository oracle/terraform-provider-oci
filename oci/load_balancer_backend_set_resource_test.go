// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/v26/loadbalancer"
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
	testAccPreCheck(s.T())
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
		ca_certificate = "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"
		certificate_name = "tf_cert_name"
		private_key = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA30+wt7OlUB/YpmWbTRkxnLG0lKWiV+oupNKj8luXmC5jvOFT\nUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU+DWVV2So2B/obYxpiiyWF2tcF/cY\n\ni1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oCMQ2985/MTdCXONgnbmePU64GrJwf\nvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOOjLKRM68KXC5us4879IrSA77NQr1K\nwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6ytM66P/1CTpk1YpbI4gqiG0HBbuX\nG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc2wIDAQABAoIBAGQznukfG/uS/qTT\njNcQifl0p8HXfLwUIa/lsJkMTj6D+k8DkF59tVMGjv3NQSQ26JVX4J1L8XiAj+fc\nUtYr1Ap4CLX5PeYUkzesvKK6lPKXQvCh+Ip2eq9PVrvL2WcdDpb5695cy7suXD7c\n05aUtS0LrINH3eXAxkpEe5UHtQFni5YLrCLEXd+SSA3OKdCB+23HRELc1iCTvqjK\n5AtR916uHTBhtREHRMvWIdu4InRUsedlJhaJOLJ8G8r64JUtfm3wLUK1U8HFOsd0\nLAx9ZURU6cXl4osTWiy1vigGaM8Xuish2HkOLNYZADDUiDBB3SshmW5IDAJ5XTn5\nqVrszRECgYEA79j1y+WLTyV7yz7XkWk3OqoQXG4b2JfKItJI1M95UwllzQ8U/krM\n+QZjP3NTtB9i1YoHyaEfic103hV9Fkgz8jvKS5ocLGJulpN4CgqbHN6v9EJ3dqTk\no6X8mpx2eP2E0ngRekFyC/OCp0Zhe2KR9PXhijMa5eB2LTeCMIS/tzkCgYEA7lmk\nIdVjcpfqY7UFJ2R8zqPJHOne2+llrl9vzo6N5kx4DzAg7MP6XO9MekOvfmD1X1Lm\nFckXWFEF+0TlN5YvCTR/+OmVufYM3xp4GBT8RZdLFbyI4+xpAAeSC4SeM0ZkC9Jt\nrKqCS24+Kqy/+qSqtkxiPLQrXSdCSfCUlmn0ALMCgYBB7SLy3q+CG82BOk7Km18g\n8un4XhOtX1uiYqa+SCETH/wpd0HP/AOHV6gkIrEZS59BDuXBGFaw7BZ5jPKLE2Gj\n7adXTI797Dh1jydpqyyjrNo0i6iGpiBqkw9x+Bvged7ucy5qql6MxmxdSk01Owzf\nhk5uTEnScfZJy34vk+2WkQKBgBXx5uy+iuN4HTqE5i6UT/FunwusdLpmqNf/LXol\nIed8TumHEuD5wklgNvhi1vuZzb2zEkAbPa0B+L0DwN73UulUDhxK1WBDyTeZZklB\nVWDK5zzfGPNzRs+b4tRwp2gtKPT1sOde45QyWELxmNNo6dbS/ZB9Pijbfnz0S5n1\ns2OFAoGBAJUohI1+d2hKlkSUzpCorQarDe8lFVEbGMu0kX0JNDI7QU+H8vDp9NOl\nGqLm3sCVBYypT8sWfchgZpcVaLRLQCQtWy4+CbMN6DT3j/uBWeDpayU5Gvqt0/no\nvwqbG6b0NEYLRPLEdsS/c8TV9mMlvb0EW+GXfmkpTrTNt3hyXniu\n-----END RSA PRIVATE KEY-----"
		public_certificate = "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"
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
			// update prop which is not force new
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// update session persistence attribute
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
				Check: resource.ComposeAggregateTestCheckFunc(
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

// todo: BackendSet health_checker appears to be missing 3 properties
// https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/HealthCheckerDetails
func (s *ResourceLoadBalancerBackendSetTestSuite) TestAccResourceLoadBalancerBackendSetLBCookie_basic() {
	var res, res2 string
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test create
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
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

                    lb_cookie_session_persistence_configuration {
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
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

					lb_cookie_session_persistence_configuration {
                        cookie_name = "example_cookie"
						domain = "example.oracle.com"
		                max_age_in_seconds = 10
		                path = "/tmp/example"
						disable_fallback = true
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// update non force-new property without update lb cookie session
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("new resource created when it should not have been")
						}
						return err
					},
				),
			},
			// update lb cookie session persistence attribute
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
				Check: resource.ComposeAggregateTestCheckFunc(
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
	httpreplay.SetScenario("TestResourceLoadBalancerBackendSetTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceLoadBalancerBackendSetTestSuite))
}
