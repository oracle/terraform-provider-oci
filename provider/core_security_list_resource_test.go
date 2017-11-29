// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-oci/crud"
)

type ResourceCoreSecurityListTestSuite struct {
	suite.Suite
	Client              *baremetal.Client
	Provider            terraform.ResourceProvider
	Providers           map[string]terraform.ResourceProvider
	Config              string
	ResourceName        string
	DefaultResourceName string
}

var defaultSecurityList = `
resource "oci_core_default_security_list" "default" {
	manage_default_resource_id = "${oci_core_virtual_network.t.default_security_list_id}"
	display_name = "default-tf-security_list"
	egress_security_rules = [{
		destination = "0.0.0.0/0"
		protocol = "6"
	}]
	ingress_security_rules = [{
		protocol = "1"
		source = "0.0.0.0/0"
		icmp_options {
			"type" = 3
			"code" = 4
		}
	},
	{
		protocol = "6"
		source = "0.0.0.0/0"
		tcp_options {
			"min" = 80
			"max" = 80
		}
	},
	{
		protocol = "17"
		source = "10.0.0.0/16"
		udp_options {
			"min" = 319
			"max" = 320
		}
	}]
}`

func extraWait(ew crud.ExtraWaitPostCreateDelete) {
	return
}

func (s *ResourceCoreSecurityListTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
		resource "oci_core_virtual_network" "t" {
			cidr_block = "10.0.0.0/16"
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-vcn"
		}`
	s.ResourceName = "oci_core_security_list.t"
	s.DefaultResourceName = "oci_core_default_security_list.default"
}

func (s *ResourceCoreSecurityListTestSuite) TestAccResourceCoreSecurityList_basic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
						egress_security_rules = [{
							destination = "0.0.0.0/0"
							protocol = "6"
						}]
						ingress_security_rules = [{
							protocol = "1"
							source = "0.0.0.0/0"
							icmp_options {
								"type" = 3
								"code" = 4
							}
						},
						{
							protocol = "6"
							source = "0.0.0.0/0"
							tcp_options {
								"min" = 80
								"max" = 80
							}
						},
						{
							protocol = "17"
							source = "10.0.0.0/16"
							udp_options {
								"min" = 319
								"max" = 320
							}
						}]
					}` + defaultSecurityList,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-security_list"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.icmp_options.0.type", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.1.tcp_options.0.max", "80"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.2.udp_options.0.max", "320"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "display_name", "default-tf-security_list"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.#", "3"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.0.icmp_options.0.type", "3"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.1.tcp_options.0.max", "80"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.2.udp_options.0.max", "320"),
				),
			},
			// verify update
			{
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list-updated"
						vcn_id = "${oci_core_virtual_network.t.id}"
						egress_security_rules = [{
							destination = "0.0.0.0/0"
							protocol = "17"
							stateless = true
						}]
						ingress_security_rules = [{
							protocol = "1"
							source = "0.0.0.0/0"
							stateless = true
							icmp_options {
								"type" = 5
								"code" = 0
							}
						},
						{
							protocol = "6"
							source = "0.0.0.0/0"
							stateless = true
							tcp_options {
								"min" = 80
								"max" = 82
							}
						},
						{
							protocol = "17"
							source = "10.0.0.0/16"
							stateless = true
						}]
					}

					resource "oci_core_default_security_list" "default" {
						manage_default_resource_id = "${oci_core_virtual_network.t.default_security_list_id}"
						display_name = "default-tf-security_list-updated"
						egress_security_rules = [{
							destination = "0.0.0.0/0"
							protocol = "17"
							stateless = true
						}]
						ingress_security_rules = [{
							protocol = "1"
							source = "0.0.0.0/0"
							stateless = true
							icmp_options {
								"type" = 5
								"code" = 0
							}
						},
						{
							protocol = "6"
							source = "0.0.0.0/0"
							stateless = true
							tcp_options {
								"min" = 80
								"max" = 82
							}
						},
						{
							protocol = "17"
							source = "10.0.0.0/16"
							stateless = true
						}]
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-security_list-updated"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.protocol", "17"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.icmp_options.0.type", "5"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.1.tcp_options.0.max", "82"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.1.stateless", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.2.stateless", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "ingress_security_rules.2.udp_options"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "display_name", "default-tf-security_list-updated"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.0.protocol", "17"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.0.icmp_options.0.type", "5"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.1.tcp_options.0.max", "82"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.1.stateless", "true"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.2.stateless", "true"),
					resource.TestCheckNoResourceAttr(s.DefaultResourceName, "ingress_security_rules.2.udp_options"),
				),
			},
			// Verify removing the default resource
			{
				Config: s.Config,
				Check:  nil,
			},
			// verify adding the default resource again
			{
				Config: s.Config + defaultSecurityList,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.DefaultResourceName, "display_name", "default-tf-security_list"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.#", "3"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.0.icmp_options.0.type", "3"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.1.tcp_options.0.max", "80"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.2.udp_options.0.max", "320"),
				),
			},
			// todo: consistent 500 error from server without this step
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			// verify lists can be cleared out
			{
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list-updated"
						vcn_id = "${oci_core_virtual_network.t.id}"
						egress_security_rules = []
						ingress_security_rules = []
					}

					resource "oci_core_default_security_list" "default" {
						manage_default_resource_id = "${oci_core_virtual_network.t.default_security_list_id}"
						display_name = "default-tf-security_list-updated"
						egress_security_rules = []
						ingress_security_rules = []
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.#", "0"),
				),
			},
		},
	})
}

func TestResourceCoreSecurityListTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreSecurityListTestSuite))
}
