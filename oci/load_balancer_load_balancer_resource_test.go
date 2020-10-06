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

type ResourceLoadBalancerLBTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceLoadBalancerLBTestSuite) SetupTest() {
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
	
	resource "oci_core_subnet" "t2" {
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet2"
	}
	
	data "oci_load_balancer_shapes" "t" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_network_security_group" "test_network_security_group" {
         compartment_id  = "${var.compartment_id}"
		 vcn_id            = "${oci_core_virtual_network.t.id}"
         display_name      =  "displayName"
    }

	resource "oci_core_network_security_group" "test_network_security_group2" {
		compartment_id = "${var.compartment_id}"
		vcn_id            = "${oci_core_virtual_network.t.id}"
	}
	`
	s.ResourceName = "oci_load_balancer.t"
}

func (s *ResourceLoadBalancerLBTestSuite) TestAccResourceLoadBalancerLB_basicPrivate() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test create
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids = ["${oci_core_subnet.t.id}"]
					display_name = "-tf-lb"
					is_private = true
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_mode"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "nsg_ids.#", "0"),
					func(ts *terraform.State) (err error) {
						resId, err = fromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// test update without nsgIds
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape          = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids     = ["${oci_core_subnet.t.id}"]
					display_name   = "-tf-lb-updated"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb-updated"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
						if resId2 != resId {
							return fmt.Errorf("resource recreated when it should not have been")
						}
						return err
					},
				),
			},
			// test update with nsgIds
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape          = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids     = ["${oci_core_subnet.t.id}"]
					display_name   = "-tf-lb-updated"
					network_security_group_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb-updated"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "network_security_group_ids.#", "1"),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
						if resId2 != resId {
							return fmt.Errorf("resource recreated when it should not have been")
						}
						return err
					},
				),
			},
			// test update with removing nsgIds
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape          = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids     = ["${oci_core_subnet.t.id}"]
					display_name   = "-tf-lb-updated"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb-updated"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "network_security_group_ids.#", "0"),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
						if resId2 != resId {
							return fmt.Errorf("resource recreated when it should not have been")
						}
						return err
					},
				),
			},
			// verify force update
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape          = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids     = ["${oci_core_subnet.t.id}", "${oci_core_subnet.t2.id}"]
					display_name   = "-tf-lb-updated"
					is_private 	   = false
					network_security_group_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb-updated"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_mode"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "network_security_group_ids.#", "1"),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
						if resId2 == resId {
							return fmt.Errorf("resource was not recreated as expected")
						}
						return err
					},
				),
			},
		},
	})
}

func (s *ResourceLoadBalancerLBTestSuite) TestAccResourceLoadBalancerLB_basicPublic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test create
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids = ["${oci_core_subnet.t.id}", "${oci_core_subnet.t2.id}"]
					display_name = "-tf-lb"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_mode"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func TestResourceLoadBalancerLBTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceLoadBalancerLBTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceLoadBalancerLBTestSuite))
}
