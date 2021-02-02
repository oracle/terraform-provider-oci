// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v35/loadbalancer"
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
		ca_certificate = "-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----"
		certificate_name = "tf_cert_name"
		private_key = "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAo83kaUQXpCcSoEuRVFX3jztWDNKtWpjNG240f0RpERI1NnZt\nHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19geIXR6TeavT+W5iRh4goK+N7gubYk\nSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCtCgd4MDlsvLv/YHCLvJL4JgRxKyev\njlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc+Je9EC3MWWxd5jBwXu3vgIYRuGR4\nDPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJKN0NXp5obaQToYqMsvAZyHoEyfCB\nDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t58QIDAQABAoIBADIyHuOPJTt9abzL\nS26vpVw0D6uAR/UyS/Ay9k1ltliv3rSg19DaHlwLjPwqnvCx7jBgTeVCYZhAkvgx\nkSsGDDcCsw+npXiG6wP9dC1jbHdVPUJLqZTPqB6sZCu8bM9RIE4Z/DcUY+HRN3qh\nmoh5wn0HSvJkNokjhx+TfY687uQfDMu0de4V2UPScZ7mboCu9HqK9qu0/krdTMH1\nrtnnFGEnx/Pe38YJl0fWxo8BHKHprwEvWX0MQzQeklnUtxREMuofSAOBe/I2DJGh\n1I94b6I66ypxuX0qAozT1MPbJGuaR+puyKawLNAQmZa9pgrrFK7e8PQUzrGVpVCp\nFtwx420CgYEA0uX/G0ycia0UTdkxkIsKIiLjs12LC0XmYjiWgkoL0PjiZzcPITn6\nvqqqGSz44HwtbrttZPm3Mo79yJ5xFiHCX0vFJykgy6cfS94imMgm8qIOS0bXjX7w\nxH2BOgp0H32LP/Zt7owcWJLEIQCjj0/4+Nvu0GskGVHlE8EYrXWf1E8CgYEAxtWk\nxBo52uNXL712VGDPNxprVGUpWSbkXpE+7wtRso5LnAnAj6dpmZsGe2zaYjRIt3Ls\nGnno4HUmwpQ5yXlHFpDUJvb2soXq3afnuAh5aVu6aKRQoG/5o3cD4pOupNbjDDNs\nTVLtTLIAIYDbph/j7pV/JnJ2WHcdk6WiVJoW/b8CgYAopLZzNmJ8jeR51D+fEYyU\nY5DqQj7Hn2L0zt8CoO6CCVToe03pI1lVYWKCk44rBQNkca51ZUKO9cum3BIDJ+Jj\npyCJmX1+geigIGEefIQ1AlIq464q0Knp1B4RZ25Vm0Y4v28UJ+BWmYI+sfbTaaAb\npZbyh5NfZc717aKp2x9ANQKBgHQpvOkUqVhIGVe6yLbjGCyJMstLjqyXHDRjhvEB\nG+nFWEcBK47Br+Adwdu57JwTD6ida3LMZlE8IDjtgBVE1VNJqahaACasNlrpDWdn\nDAeRn4Yi+TfCM4Zcsdhdj1qecGdgY5WJLTnxhEIOlkSnvPJWRMKhfKKSdKUdz4i9\nvVDhAoGAEHxfhFLVwdTa0RMUq3KYSXa5WqLANRn2e62Cc3eUWekcUjbIATRF5AIo\nm0WS+rURZWy1Fd6fGg8sRHock0+vxwqeP6OlyW4tJMhL33NrNbgyvkXlMMIX6bC4\nUq8aAew0B3j61UUsTqhHMhYwIS3GOIHx/O10wwINPnUMIVER3Wg=\n-----END RSA PRIVATE KEY-----"
		public_certificate = "-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----"
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
