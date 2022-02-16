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
	"github.com/oracle/oci-go-sdk/v58/loadbalancer"
	"github.com/stretchr/testify/suite"
)

type ResourceLoadBalancerBackendTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceLoadBalancerBackendTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
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
			// test Create minimal
			{
				Config: s.Config + `
				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 8080
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
						res, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						return err
					},
				),
			},
			// test partial Update: "weight" only, omitted bool properties should be not be named and null in the Update request
			{
				Config: s.Config + `
				resource "oci_load_balancer_backend" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					backendset_name = "${oci_load_balancer_backendset.t.name}"
					ip_address = "1.2.3.4"
					port = 8080
					weight = 2
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
						res, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						return err
					},
				),
			},
			// test partial Update - previously omitted bools explicitly set to a mix of true and false
			{
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
						res, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
						return err
					},
				),
			},
			// test full Update - invert bools
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
						res2, err = acctest.FromInstanceState(ts, s.ResourceName, "name")
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

// issue-routing-tag: load_balancer/default
func TestResourceLoadBalancerBackendTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceLoadBalancerBackendTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceLoadBalancerBackendTestSuite))
}
