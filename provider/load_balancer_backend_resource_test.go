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

type ResourceLoadBalancerBackendTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceLoadBalancerBackendTestSuite) SetupTest() {
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
	}`
	s.ResourceName = "oci_load_balancer_backend.t"
}

func (s *ResourceLoadBalancerBackendTestSuite) TestAccResourceLoadBalancerBackend_basic() {
	var res, res2 string
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test create minimal
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 8080
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backendset_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.ResourceName, "ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr(s.ResourceName, "port", "8080"),
					resource.TestCheckResourceAttr(s.ResourceName, "backup", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "drain", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "offline", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "weight", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "name"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res, err = fromInstanceState(ts, s.ResourceName, "name")
						return err
					},
				),
			},
			// test partial update: "weight" only, omitted bool properties should be not be named and null in the update request
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 8080
					weight = 2
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backendset_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.ResourceName, "ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr(s.ResourceName, "port", "8080"),
					resource.TestCheckResourceAttr(s.ResourceName, "backup", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "drain", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "offline", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "weight", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "name"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res, err = fromInstanceState(ts, s.ResourceName, "name")
						return err
					},
				),
			},
			// test partial update - previously omitted bools explicitly set to a mix of true and false
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 8080
					backup = false
					drain = true
					offline = false
					weight = 1
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backendset_name", "-tf-backend-set"),
					resource.TestCheckResourceAttr(s.ResourceName, "ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr(s.ResourceName, "port", "8080"),
					resource.TestCheckResourceAttr(s.ResourceName, "backup", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "drain", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "offline", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "weight", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "name"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res, err = fromInstanceState(ts, s.ResourceName, "name")
						return err
					},
				),
			},
			// test full update - invert bools
			{
				Config: s.Config + `
				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 8080
					backup = true
					drain = false
					offline = true
					weight = 3
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "port", "8080"),
					resource.TestCheckResourceAttr(s.ResourceName, "backup", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "drain", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "offline", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "weight", "3"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "backendset_name", "-tf-backend-set"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "name"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.WorkRequestLifecycleStateSucceeded)),
					func(ts *terraform.State) (err error) {
						res2, err = fromInstanceState(ts, s.ResourceName, "name")
						if res != res2 {
							return fmt.Errorf("resource was unexpectedly recreated while updating updatable properties")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceLoadBalancerBackendTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerBackendTestSuite))
}
