// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/loadbalancer"
	"github.com/stretchr/testify/suite"
)

type ResourceLoadBalancerListenerTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceLoadBalancerListenerTestSuite) SetupTest() {
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
	
	resource "oci_load_balancer_certificate" "t" {
		load_balancer_id = "${oci_load_balancer.t.id}"
		ca_certificate = "${var.ca_certificate_value}"
		certificate_name = "tf_cert_name"
		private_key = "${var.private_key_value}"
		public_certificate = "${var.ca_certificate_value}"
	}`
	s.ResourceName = "oci_load_balancer_listener.t"
}

func (s *ResourceLoadBalancerListenerTestSuite) TestAccResourceLoadBalancerListener_basic() {
	var resId, resId2 string
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test Create
			{
				Config: s.Config + `
				resource "oci_load_balancer_listener" "t" {
					load_balancer_id  = "${oci_load_balancer.t.id}"
					name = "-tf-listener"
					default_backend_set_name = "${oci_load_balancer_backendset.t.name}"
					port = 8080
					protocol = "TCP"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-listener"),
					resource.TestCheckResourceAttr(s.ResourceName, "default_backend_set_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.ResourceName, "port", "8080"),
					resource.TestCheckResourceAttr(s.ResourceName, "protocol", "TCP"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// test Update
			{
				Config: s.Config + `
				resource "oci_load_balancer_listener" "t" {
					load_balancer_id  = "${oci_load_balancer.t.id}"
					name = "-tf-listener-updated"
					default_backend_set_name = "${oci_load_balancer_backendset.t.name}"
					port = 80
					protocol = "HTTP"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-listener-updated"),
					resource.TestCheckResourceAttr(s.ResourceName, "default_backend_set_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.ResourceName, "port", "80"),
					resource.TestCheckResourceAttr(s.ResourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId2 == resId {
							return fmt.Errorf("resource expected to be recreated but was not")
						}
						resId = resId2
						return err
					},
				),
			},
			// test add ssl configuration
			{
				Config: s.Config + `
				resource "oci_load_balancer_listener" "t" {
					load_balancer_id  = "${oci_load_balancer.t.id}"
					name = "-tf-listener-updated"
					default_backend_set_name = "${oci_load_balancer_backendset.t.name}"
					port = 443
					protocol = "HTTP"
				
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-listener-updated"),
					resource.TestCheckResourceAttr(s.ResourceName, "default_backend_set_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.ResourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttr(s.ResourceName, "port", "443"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.certificate_name", "tf_cert_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_depth", "6"),
					resource.TestCheckResourceAttr(s.ResourceName, "ssl_configuration.0.verify_peer_certificate", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId2 != resId {
							return fmt.Errorf("resource recreated when it should not have been")
						}
						resId = resId2
						return err
					},
				),
			},
			// test remove ssl configuration
			{
				Config: s.Config + `
				resource "oci_load_balancer_listener" "t" {
					load_balancer_id  = "${oci_load_balancer.t.id}"
					name = "-tf-listener-updated"
					default_backend_set_name = "${oci_load_balancer_backendset.t.name}"
					port = 443
					protocol = "HTTP"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-listener-updated"),
					resource.TestCheckResourceAttr(s.ResourceName, "default_backend_set_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.ResourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttr(s.ResourceName, "port", "443"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "ssl_configuration.0.certificate_name"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "ssl_configuration.0.verify_depth"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "ssl_configuration.0.verify_peer_certificate"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId2 != resId {
							return fmt.Errorf("resource recreated when it should not have been")
						}
						resId = resId2
						return err
					},
				),
			},
			// verify resource import
			{
				Config: s.Config + `
				resource "oci_load_balancer_listener" "t" {
					load_balancer_id  = "${oci_load_balancer.t.id}"
					name = "-tf-listener-updated"
					default_backend_set_name = "${oci_load_balancer_backendset.t.name}"
					port = 80
					protocol = "HTTP"
	
					ssl_configuration {
						certificate_name = "${oci_load_balancer_certificate.t.certificate_name}"
						verify_depth = 6
						verify_peer_certificate = false
					}
				}`,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"load_balancer_id",
					"passphrase",
					"private_key",
					"state",
				},
				ResourceName: "oci_load_balancer_listener.t",
			},
		},
	})
}

// issue-routing-tag: load_balancer/default
func TestResourceLoadBalancerListenerTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceLoadBalancerListenerTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceLoadBalancerListenerTestSuite))
}
