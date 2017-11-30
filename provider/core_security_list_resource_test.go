// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreSecurityListTestSuite struct {
	suite.Suite
	Client         *baremetal.Client
	Provider       terraform.ResourceProvider
	Providers      map[string]terraform.ResourceProvider
	Config         string
	ResourceName   string
	DataSourceName string
	FullConfig     string
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
		}
	    data "oci_core_security_lists" "t" {
			compartment_id = "${var.compartment_id}"
			vcn_id = "${oci_core_virtual_network.t.id}"
			filter {
				name = "display_name"
				values = ["${oci_core_security_list.t.display_name}"]
			}
		}`
	s.ResourceName = "oci_core_security_list.t"
	s.DataSourceName = "data.oci_core_security_lists.t"

	s.FullConfig = `
		resource "oci_core_security_list" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-security_list"
			vcn_id = "${oci_core_virtual_network.t.id}"
			egress_security_rules = {
				destination = "0.0.0.0/1"
				protocol = "6"
			}
			egress_security_rules = {
				destination = "0.0.0.0/2"
				protocol = "1"
				stateless = true
				icmp_options {
					"type" = 3
					"code" = 4
				}
			}
			egress_security_rules = {
				destination = "0.0.0.0/3"
				protocol = "6"
				stateless = false
				tcp_options {
					"min" = 10
					"max" = 11
					source_port_range {
						"min" = 20
						"max" = 21
					}
				}
			}
			egress_security_rules = {
				destination = "0.0.0.0/4"
				protocol = "17"
				udp_options {
					"min" = 30
					"max" = 31
					source_port_range {
						"min" = 40
						"max" = 41
					}
				}
			}
			ingress_security_rules = [{
				protocol = "1"
				source = "0.0.0.0/5"
			},
			{
				protocol = "1"
				source = "0.0.0.0/6"
				icmp_options {
					"type" = 3
					"code" = 4
				}
			},
			{
				protocol = "6"
				stateless = true
				source = "0.0.0.0/7"
				tcp_options {
					"min" = 50
					"max" = 51
					source_port_range {
						"min" = 60
						"max" = 61
					}
				}
			},
			{
				protocol = "17"
				stateless = false
				source = "10.0.0.0/8"
				udp_options {
					"min" = 70
					"max" = 71
					source_port_range {
						"min" = 80
						"max" = 81
					}
				}
			}]
		}
	`
}

func (s *ResourceCoreSecurityListTestSuite) BuildTestsForFullConfig(resourceName, prefix string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, prefix+"display_name", "-tf-security_list"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.#", "4"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.0.destination", "0.0.0.0/1"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.0.protocol", "6"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.0.stateless", "false"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.0.tcp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.0.udp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.0.icmp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.1.destination", "0.0.0.0/2"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.1.protocol", "1"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.1.stateless", "true"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.1.tcp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.1.udp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.1.icmp_options.#", "1"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.1.icmp_options.0.type", "3"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.1.icmp_options.0.code", "4"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.2.destination", "0.0.0.0/3"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.2.protocol", "6"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.2.stateless", "false"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.2.tcp_options.#", "1"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.2.tcp_options.0.min", "10"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.2.tcp_options.0.max", "11"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.2.tcp_options.0.source_port_range.0.min", "20"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.2.tcp_options.0.source_port_range.0.max", "21"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.2.udp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.2.icmp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.3.destination", "0.0.0.0/4"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.3.protocol", "17"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.3.stateless", "false"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.3.tcp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.3.udp_options.#", "1"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.3.udp_options.0.min", "30"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.3.udp_options.0.max", "31"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.3.udp_options.0.source_port_range.0.min", "40"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.3.udp_options.0.source_port_range.0.max", "41"),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.3.icmp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.#", "4"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.0.source", "0.0.0.0/5"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.0.protocol", "1"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.0.stateless", "false"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.0.tcp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.0.udp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.0.icmp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.1.source", "0.0.0.0/6"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.1.protocol", "1"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.1.stateless", "false"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.1.tcp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.1.udp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.1.icmp_options.#", "1"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.1.icmp_options.0.type", "3"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.1.icmp_options.0.code", "4"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.2.source", "0.0.0.0/7"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.2.protocol", "6"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.2.stateless", "true"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.2.tcp_options.#", "1"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.2.tcp_options.0.min", "50"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.2.tcp_options.0.max", "51"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.2.tcp_options.0.source_port_range.0.min", "60"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.2.tcp_options.0.source_port_range.0.max", "61"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.2.udp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.2.icmp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.3.source", "10.0.0.0/8"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.3.protocol", "17"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.3.stateless", "false"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.3.tcp_options.#", "0"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.3.udp_options.#", "1"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.3.udp_options.0.min", "70"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.3.udp_options.0.max", "71"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.3.udp_options.0.source_port_range.0.min", "80"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.3.udp_options.0.source_port_range.0.max", "81"),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.3.icmp_options.#", "0"),
	}
}

func (s *ResourceCoreSecurityListTestSuite) TestAccResourceCoreSecurityList_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create with all options
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config + s.FullConfig,
				Check: resource.ComposeTestCheckFunc(append(s.BuildTestsForFullConfig(s.ResourceName, ""),
					s.BuildTestsForFullConfig(s.DataSourceName, "security_lists.0.")...)...),
			},
			// Plan with the same config should do nothing
			{
				Config:             s.Config + s.FullConfig,
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			// Update to a single rule
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
						egress_security_rules = {
							destination = "0.0.0.0/1"
							protocol = "6"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.destination", "0.0.0.0/1"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.tcp_options.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.udp_options.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.icmp_options.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
				),
			},
			// Update to zero rules
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
				),
			},
			// Update to rules that use only source and only destination port ranges
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
						egress_security_rules = {
							destination = "0.0.0.0/3"
							protocol = "6"
							stateless = false
							tcp_options {
								source_port_range {
									"min" = 20
									"max" = 21
								}
							}
						}
						# Check the maximum range
						egress_security_rules = {
							destination = "0.0.0.0/4"
							protocol = "17"
							udp_options {
								"min" = 1
								"max" = 65535
							}
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.destination", "0.0.0.0/3"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.tcp_options.0.min", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.tcp_options.0.max", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.min", "20"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.max", "21"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.udp_options.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.icmp_options.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.1.destination", "0.0.0.0/4"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.1.protocol", "17"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.1.stateless", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.1.tcp_options.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.1.udp_options.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.1.udp_options.0.min", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.1.udp_options.0.max", "65535"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.1.udp_options.0.source_port_range.#", "0"),
				),
			},
		},
	})
}

func (s *ResourceCoreSecurityListTestSuite) TestAccResourceCoreSecurityList_emptyList() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				// Create a security list with no rules (which is different from the earlier test of updating to no rules)
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
				),
			},
			// update with all options
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config + s.FullConfig,
				Check:             resource.ComposeTestCheckFunc(s.BuildTestsForFullConfig(s.ResourceName, "")...),
			},
			// Apply the same config and check the data source, since the data source will not have updated on the previous apply.
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config + s.FullConfig,
				Check:             resource.ComposeTestCheckFunc(s.BuildTestsForFullConfig(s.DataSourceName, "security_lists.0.")...),
			},
		},
	})
}

func TestResourceCoreSecurityListTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreSecurityListTestSuite))
}
