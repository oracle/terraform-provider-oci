// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// issue-routing-tag: load_balancer/default
func TestAccDatasourceLoadBalancerCertificates_basic(t *testing.T) {
	httpreplay.SetScenario("TestAccDatasourceLoadBalancerCertificates_basic")
	defer httpreplay.SaveScenario()
	providers := testAccProviders
	config := legacyTestProviderConfig() + caCertificateVariableStr + privateKeyVariableStr + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
	}
	
	resource "oci_core_subnet" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		route_table_id = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dhcp_options_id = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block = "10.0.0.0/24"
		display_name = "-tf-subnet"
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
	}
	
	data "oci_load_balancer_certificates" "t" {
		load_balancer_id = "${oci_load_balancer.t.id}"
	}`

	resourceName := "data.oci_load_balancer_certificates.t"

	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
			},
			{
				Config: config,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "certificates.#", "1"),
					resource.TestMatchResourceAttr(resourceName, "certificates.0.ca_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
					resource.TestCheckResourceAttr(resourceName, "certificates.0.certificate_name", "tf_cert_name"),
					resource.TestMatchResourceAttr(resourceName, "certificates.0.public_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
				),
			},
		},
	})
}
