// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreSecurityListTestSuite struct {
	suite.Suite
	Providers           map[string]terraform.ResourceProvider
	Config              string
	ResourceName        string
	DataSourceName      string
	DefaultResourceName string
}

func (s *ResourceCoreSecurityListTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
		resource "oci_core_virtual_network" "t" {
			cidr_block = "10.0.0.0/16"
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-vcn"
		}`
	s.ResourceName = "oci_core_security_list.t"
	s.DefaultResourceName = "oci_core_default_security_list.default"
	s.DataSourceName = "data.oci_core_security_lists.t"
}

var dataSource = `
	data "oci_core_security_lists" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		filter {
			name = "display_name"
			values = ["${oci_core_security_list.t.display_name}"]
		}
	}`

var fullConfig = `
	resource "oci_core_security_list" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-security_list"
		vcn_id = "${oci_core_virtual_network.t.id}"
		egress_security_rules {
			destination = "0.0.0.0/1"
			protocol = "6"
		}
		egress_security_rules {
			destination = "0.0.0.0/2"
			protocol = "1"
			stateless = true
			icmp_options {
				type = 3
				code = 4
			}
		}
		egress_security_rules {
			destination = "0.0.0.0/3"
			protocol = "6"
			stateless = false
			tcp_options {
				min = 10
				max = 11
				source_port_range {
					min = 20
					max = 21
				}
			}
		}
		egress_security_rules {
			destination = "0.0.0.0/4"
			protocol = "17"
			udp_options {
				min = 30
				max = 31
				source_port_range {
					min = 40
					max = 41
				}
			}
		}
		ingress_security_rules {
			protocol = "1"
			source = "0.0.0.0/5"
		}
		ingress_security_rules {
			protocol = "1"
			source = "0.0.0.0/6"
			icmp_options {
				type = 3
				code = 4
			}
		}
		ingress_security_rules {
			protocol = "6"
			stateless = true
			source = "0.0.0.0/7"
			tcp_options {
				min = 50
				max = 51
				source_port_range {
					min = 60
					max = 61
				}
			}
		}
		ingress_security_rules {
			protocol = "17"
			stateless = false
			source = "10.0.0.0/8"
			udp_options {
				min = 70
				max = 71
				source_port_range {
					min = 80
					max = 81
				}
			}
		}
	}
`

// Verifies the contents of fullConfig, with parameters that allows this to be checked via eithier a resource or a data source.
func (s *ResourceCoreSecurityListTestSuite) BuildTestsForFullConfig(resourceName, prefix string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, prefix+"display_name", "-tf-security_list"),
		//resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		resource.TestCheckResourceAttr(resourceName, prefix+"egress_security_rules.#", "4"),
		acctest.CheckResourceSetContainsElementWithProperties(resourceName, prefix+"egress_security_rules", map[string]string{
			"destination":    "0.0.0.0/1",
			"icmp_options.#": "0",
			"tcp_options.#":  "0",
			"udp_options.#":  "0",
			"protocol":       "6",
			"stateless":      "false",
		},
			nil),
		acctest.CheckResourceSetContainsElementWithProperties(resourceName, prefix+"egress_security_rules", map[string]string{
			"destination":         "0.0.0.0/2",
			"icmp_options.#":      "1",
			"icmp_options.0.code": "4",
			"icmp_options.0.type": "3",
			"tcp_options.#":       "0",
			"udp_options.#":       "0",
			"protocol":            "1",
			"stateless":           "true",
		},
			nil),
		acctest.CheckResourceSetContainsElementWithProperties(resourceName, prefix+"egress_security_rules", map[string]string{
			"destination":                           "0.0.0.0/3",
			"icmp_options.#":                        "0",
			"tcp_options.#":                         "1",
			"tcp_options.0.min":                     "10",
			"tcp_options.0.max":                     "11",
			"tcp_options.0.source_port_range.0.min": "20",
			"tcp_options.0.source_port_range.0.max": "21",
			"udp_options.#":                         "0",
			"protocol":                              "6",
			"stateless":                             "false",
		},
			nil),
		acctest.CheckResourceSetContainsElementWithProperties(resourceName, prefix+"egress_security_rules", map[string]string{
			"destination":                           "0.0.0.0/4",
			"icmp_options.#":                        "0",
			"tcp_options.#":                         "0",
			"udp_options.#":                         "1",
			"udp_options.0.min":                     "30",
			"udp_options.0.max":                     "31",
			"udp_options.0.source_port_range.0.min": "40",
			"udp_options.0.source_port_range.0.max": "41",
			"protocol":                              "17",
			"stateless":                             "false",
		},
			nil),
		resource.TestCheckResourceAttr(resourceName, prefix+"ingress_security_rules.#", "4"),
		acctest.CheckResourceSetContainsElementWithProperties(resourceName, prefix+"ingress_security_rules", map[string]string{
			"source":         "0.0.0.0/5",
			"icmp_options.#": "0",
			"tcp_options.#":  "0",
			"udp_options.#":  "0",
			"protocol":       "1",
			"stateless":      "false",
		},
			nil),
		acctest.CheckResourceSetContainsElementWithProperties(resourceName, prefix+"ingress_security_rules", map[string]string{
			"source":              "0.0.0.0/6",
			"icmp_options.#":      "1",
			"icmp_options.0.code": "4",
			"icmp_options.0.type": "3",
			"tcp_options.#":       "0",
			"udp_options.#":       "0",
			"protocol":            "1",
			"stateless":           "false",
		},
			nil),
		acctest.CheckResourceSetContainsElementWithProperties(resourceName, prefix+"ingress_security_rules", map[string]string{
			"source":                                "0.0.0.0/7",
			"icmp_options.#":                        "0",
			"tcp_options.#":                         "1",
			"tcp_options.0.min":                     "50",
			"tcp_options.0.max":                     "51",
			"tcp_options.0.source_port_range.0.min": "60",
			"tcp_options.0.source_port_range.0.max": "61",
			"udp_options.#":                         "0",
			"protocol":                              "6",
			"stateless":                             "true",
		},
			nil),
		acctest.CheckResourceSetContainsElementWithProperties(resourceName, prefix+"ingress_security_rules", map[string]string{
			"source":                                "10.0.0.0/8",
			"icmp_options.#":                        "0",
			"tcp_options.#":                         "0",
			"udp_options.#":                         "1",
			"udp_options.0.min":                     "70",
			"udp_options.0.max":                     "71",
			"udp_options.0.source_port_range.0.min": "80",
			"udp_options.0.source_port_range.0.max": "81",
			"protocol":                              "17",
			"stateless":                             "false",
		},
			nil),
	}
}

func (s *ResourceCoreSecurityListTestSuite) TestAccResourceCoreSecurityList_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify Create with all options
			{
				Config: s.Config + dataSource + fullConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(append(s.BuildTestsForFullConfig(s.ResourceName, ""),
					s.BuildTestsForFullConfig(s.DataSourceName, "security_lists.0.")...)...),
			},
			// Plan with the same config should do nothing
			{
				Config:             s.Config + dataSource + fullConfig,
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			// Update to a single rule
			{
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
						egress_security_rules {
							destination = "0.0.0.0/1"
							protocol = "6"
						}
					}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(s.ResourceName, "egress_security_rules", map[string]string{
						"destination":    "0.0.0.0/1",
						"icmp_options.#": "0",
						"tcp_options.#":  "0",
						"udp_options.#":  "0",
						"protocol":       "6",
						"stateless":      "false",
					},
						nil),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
				),
			},
			// Update to zero rules
			{
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
					}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
				),
			},
			// Update to add ICMP options
			{
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
						ingress_security_rules {
							protocol = "1"
							source = "0.0.0.0/6"
							icmp_options {
								type = 3
								code = 4
							}
						}
						egress_security_rules {
							protocol = "1"
							destination = "0.0.0.0/6"
							icmp_options {
								type = 3
								code = 4
							}
						}
					}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(s.ResourceName, "egress_security_rules", map[string]string{
						"icmp_options.#":      "1",
						"icmp_options.0.code": "4",
						"icmp_options.0.type": "3",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.ResourceName, "ingress_security_rules", map[string]string{
						"icmp_options.#":      "1",
						"icmp_options.0.code": "4",
						"icmp_options.0.type": "3",
					},
						nil),
				),
			},
			// Update to ICMP options that don't contain an optional 'code'
			{
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
						ingress_security_rules {
							protocol = "1"
							source = "0.0.0.0/6"
							icmp_options {
								type = 3
							}
						}
						egress_security_rules {
							protocol = "1"
							destination = "0.0.0.0/6"
							icmp_options {
								type = 3
							}
						}
					}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(s.ResourceName, "egress_security_rules", map[string]string{
						"icmp_options.#":      "1",
						"icmp_options.0.code": "-1",
						"icmp_options.0.type": "3",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.ResourceName, "ingress_security_rules", map[string]string{
						"icmp_options.#":      "1",
						"icmp_options.0.code": "-1",
						"icmp_options.0.type": "3",
					},
						nil),
				),
			},
			// Update to rules that use only source and only destination port ranges
			// Also tests removal of icmp_options
			{
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
						egress_security_rules {
							destination = "0.0.0.0/3"
							protocol = "6"
							stateless = false
							tcp_options {
								source_port_range {
									min = 20
									max = 21
								}
							}
						}
						# Check the maximum range
						egress_security_rules {
							destination = "0.0.0.0/4"
							protocol = "17"
							udp_options {
								min = 1
								max = 65535
							}
						}
					}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
					acctest.CheckResourceSetContainsElementWithProperties(s.ResourceName, "egress_security_rules", map[string]string{
						"destination":                           "0.0.0.0/3",
						"icmp_options.#":                        "0",
						"tcp_options.#":                         "1",
						"tcp_options.0.min":                     "0",
						"tcp_options.0.max":                     "0",
						"tcp_options.0.source_port_range.0.min": "20",
						"tcp_options.0.source_port_range.0.max": "21",
						"udp_options.#":                         "0",
						"protocol":                              "6",
						"stateless":                             "false",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.ResourceName, "egress_security_rules", map[string]string{
						"destination":                       "0.0.0.0/4",
						"icmp_options.#":                    "0",
						"tcp_options.#":                     "0",
						"udp_options.#":                     "1",
						"udp_options.0.min":                 "1",
						"udp_options.0.max":                 "65535",
						"udp_options.0.source_port_range.#": "0",
						"protocol":                          "17",
						"stateless":                         "false",
					},
						nil),
				),
			},
			// Remove source_port_range from tcp_options and add it to udp_options
			{
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
						egress_security_rules {
							destination = "0.0.0.0/3"
							protocol = "6"
							stateless = false
							tcp_options {
								min = 20
								max = 21
							}
						}
						# Check the maximum range
						egress_security_rules {
							destination = "0.0.0.0/4"
							protocol = "17"
							udp_options {
								source_port_range {
									min = 1
									max = 65535
								}
							}
						}
					}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
					acctest.CheckResourceSetContainsElementWithProperties(s.ResourceName, "egress_security_rules", map[string]string{
						"destination":                       "0.0.0.0/3",
						"icmp_options.#":                    "0",
						"tcp_options.#":                     "1",
						"tcp_options.0.min":                 "20",
						"tcp_options.0.max":                 "21",
						"tcp_options.0.source_port_range.#": "0",
						"udp_options.#":                     "0",
						"protocol":                          "6",
						"stateless":                         "false",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.ResourceName, "egress_security_rules", map[string]string{
						"destination":                           "0.0.0.0/4",
						"icmp_options.#":                        "0",
						"tcp_options.#":                         "0",
						"udp_options.#":                         "1",
						"udp_options.0.min":                     "0",
						"udp_options.0.max":                     "0",
						"udp_options.0.source_port_range.0.min": "1",
						"udp_options.0.source_port_range.0.max": "65535",
						"protocol":                              "17",
						"stateless":                             "false",
					},
						nil),
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
				Config: s.Config + `
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
					}

					resource "oci_core_default_security_list" "default" {
						manage_default_resource_id = "${oci_core_virtual_network.t.default_security_list_id}"
						display_name = "default-tf-security_list-updated"
					}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.#", "0"),
				),
			},
			// Update with all options
			{
				Config: s.Config + dataSource + fullConfig,
				Check:  acctest.ComposeAggregateTestCheckFuncWrapper(s.BuildTestsForFullConfig(s.ResourceName, "")...),
			},
			// Apply the same config and check the data source, since the data source will not have updated on the previous apply.
			{
				Config: s.Config + dataSource + fullConfig,
				Check:  acctest.ComposeAggregateTestCheckFuncWrapper(s.BuildTestsForFullConfig(s.DataSourceName, "security_lists.0.")...),
			},
		},
	})
}

func (s *ResourceCoreSecurityListTestSuite) TestAccResourceCoreSecurityList_defaultSecurityList() {
	defaultSecurityList := `
		resource "oci_core_default_security_list" "default" {
			manage_default_resource_id = "${oci_core_virtual_network.t.default_security_list_id}"
			display_name = "default-tf-security_list"
			egress_security_rules {
				destination = "0.0.0.0/0"
				protocol = "6"
			}
			ingress_security_rules {
				protocol = "1"
				source = "0.0.0.0/0"
				icmp_options {
					type = 3
					code = 4
				}
			}
			ingress_security_rules {
				protocol = "6"
				source = "0.0.0.0/0"
				tcp_options {
					min = 80
					max = 80
				}
			}
			ingress_security_rules {
				protocol = "17"
				source = "10.0.0.0/16"
				udp_options {
					min = 319
					max = 320
				}
			}
		}`
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + defaultSecurityList,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.DefaultResourceName, "display_name", "default-tf-security_list"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "egress_security_rules", map[string]string{
						"stateless": "false",
					},
						nil),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.#", "3"),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "ingress_security_rules", map[string]string{
						"icmp_options.0.type": "3",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "ingress_security_rules", map[string]string{
						"tcp_options.0.max": "80",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "ingress_security_rules", map[string]string{
						"udp_options.0.max": "320",
					},
						nil),
				),
			},
			// Update
			{
				Config: compartmentIdUVariableStr + s.Config + `
					resource "oci_core_default_security_list" "default" {
						manage_default_resource_id = "${oci_core_virtual_network.t.default_security_list_id}"
						display_name = "default-tf-security_list-updated"
						compartment_id = "${var.compartment_id_for_update}"
						egress_security_rules {
							destination = "0.0.0.0/0"
							protocol = "17"
							stateless = true
						}
						ingress_security_rules {
							protocol = "1"
							source = "0.0.0.0/0"
							stateless = true
							icmp_options {
								type = 5
								code = 0
							}
						}
						ingress_security_rules {
							protocol = "6"
							source = "0.0.0.0/0"
							stateless = true
							tcp_options {
								source_port_range {
									min = 99
									max = 100
								}
							}
						}
						ingress_security_rules {
							protocol = "17"
							source = "10.0.0.0/16"
							stateless = true
						}
					}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.DefaultResourceName, "display_name", "default-tf-security_list-updated"),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "egress_security_rules", map[string]string{
						"protocol":  "17",
						"stateless": "true",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "ingress_security_rules", map[string]string{
						"icmp_options.0.type": "5",
						"stateless":           "true",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "ingress_security_rules", map[string]string{
						"tcp_options.0.source_port_range.0.max": "100",
						"stateless":                             "true",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "ingress_security_rules", map[string]string{
						"tcp_options.0.source_port_range.0.max": "100",
						"stateless":                             "true",
					},
						nil),
					resource.TestCheckNoResourceAttr(s.DefaultResourceName, "ingress_security_rules.2.udp_options"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "compartment_id", compartmentIdU),
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.DefaultResourceName, "display_name", "default-tf-security_list"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "ingress_security_rules", map[string]string{
						"stateless": "false",
					},
						nil),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.#", "3"),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "ingress_security_rules", map[string]string{
						"icmp_options.0.type": "3",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "ingress_security_rules", map[string]string{
						"tcp_options.0.max": "80",
					},
						nil),
					acctest.CheckResourceSetContainsElementWithProperties(s.DefaultResourceName, "ingress_security_rules", map[string]string{
						"udp_options.0.max": "320",
					},
						nil),
				),
			},
			// Verify lists can be cleared out. Also try adding an additional security list.
			{
				Config: s.Config + `
					resource "oci_core_default_security_list" "default" {
						manage_default_resource_id = "${oci_core_virtual_network.t.default_security_list_id}"
						display_name = "default-tf-security_list-updated"
					}
					resource "oci_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-security_list"
						vcn_id = "${oci_core_virtual_network.t.id}"
					}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "egress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "ingress_security_rules.#", "0"),
				),
			},
		},
	})
}

// issue-routing-tag: core/virtualNetwork
func TestResourceCoreSecurityListTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreSecurityListTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreSecurityListTestSuite))
}
