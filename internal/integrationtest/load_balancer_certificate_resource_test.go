// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"regexp"
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/suite"
)

type ResourceLoadBalancerCertificateTestSuite struct {
	suite.Suite
	Providers    map[string]*schema.Provider
	Config       string
	ResourceName string
}

func (s *ResourceLoadBalancerCertificateTestSuite) SetupTest() {
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
	}`
	s.ResourceName = "oci_load_balancer_certificate.t"
}

func (s *ResourceLoadBalancerCertificateTestSuite) TestAccResourceLoadBalancerCertificate_basic() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test Create
			{
				Config: s.Config + `
				resource "oci_load_balancer_certificate" "t" {
					load_balancer_id = "${oci_load_balancer.t.id}"
					ca_certificate = "${var.ca_certificate_value}"
					certificate_name = "tf_cert_name"
					private_key = "${var.private_key_value}"
					public_certificate = "${var.ca_certificate_value}"
				}
` + caCertificateVariableStr + privateKeyVariableStr,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "load_balancer_id"),
					resource.TestMatchResourceAttr(s.ResourceName, "ca_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
					resource.TestCheckResourceAttr(s.ResourceName, "certificate_name", "tf_cert_name"),
					resource.TestMatchResourceAttr(s.ResourceName, "private_key", regexp.MustCompile("-----BEGIN RSA.*")),
					resource.TestMatchResourceAttr(s.ResourceName, "public_certificate", regexp.MustCompile("-----BEGIN CERT.*")),
				),
			},
		},
	})
}

// issue-routing-tag: load_balancer/default
func TestResourceLoadBalancerCertificateTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceLoadBalancerCertificateTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceLoadBalancerCertificateTestSuite))
}
