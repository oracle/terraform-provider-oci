// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	SecurityListRequiredOnlyResource = SecurityListResourceDependencies + `
resource "oci_core_security_list" "test_security_list" {
	#Required
	compartment_id = "${var.compartment_id}"
	egress_security_rules {
		#Required
		destination = "${var.security_list_egress_security_rules_destination}"
		protocol = "${var.security_list_egress_security_rules_protocol}"
	}
	ingress_security_rules {
		#Required
		protocol = "${var.security_list_ingress_security_rules_protocol}"
		source = "${var.security_list_ingress_security_rules_source}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`

	SecurityListResourceConfig = SecurityListResourceDependencies + `
resource "oci_core_security_list" "test_security_list" {
	#Required
	compartment_id = "${var.compartment_id}"
	egress_security_rules {
		#Required
		destination = "${var.security_list_egress_security_rules_destination}"
		protocol = "1"

		#Optional
		destination_type = "${var.security_list_egress_security_rules_destination_type}"
		icmp_options {
			#Required
			type = "${var.security_list_egress_security_rules_icmp_options_type}"

			#Optional
			code = "${var.security_list_egress_security_rules_icmp_options_code}"
		}
		stateless = "${var.security_list_egress_security_rules_stateless}"
	}
	egress_security_rules {
		#Required
		destination = "${var.security_list_egress_security_rules_destination}"
		protocol = "6"

		#Optional
		destination_type = "${var.security_list_egress_security_rules_destination_type}"
		stateless = "${var.security_list_egress_security_rules_stateless}"
		tcp_options {

			#Optional
			max = "${var.security_list_egress_security_rules_tcp_options_destination_port_range_max}"
			min = "${var.security_list_egress_security_rules_tcp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_egress_security_rules_tcp_options_source_port_range_max}"
				min = "${var.security_list_egress_security_rules_tcp_options_source_port_range_min}"
			}
		}
	}
	egress_security_rules {
		#Required
		destination = "${var.security_list_egress_security_rules_destination}"
		protocol = "17"

		#Optional
		destination_type = "${var.security_list_egress_security_rules_destination_type}"
		stateless = "${var.security_list_egress_security_rules_stateless}"
		udp_options {

			#Optional
			max = "${var.security_list_egress_security_rules_udp_options_destination_port_range_max}"
			min = "${var.security_list_egress_security_rules_udp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_egress_security_rules_udp_options_source_port_range_max}"
				min = "${var.security_list_egress_security_rules_udp_options_source_port_range_min}"
			}
		}
	}
	ingress_security_rules {
		#Required
		protocol = "1"
		source = "${var.security_list_ingress_security_rules_source}"

		#Optional
		icmp_options {
			#Required
			type = "${var.security_list_ingress_security_rules_icmp_options_type}"

			#Optional
			code = "${var.security_list_ingress_security_rules_icmp_options_code}"
		}
		source_type = "${var.security_list_ingress_security_rules_source_type}"
		stateless = "${var.security_list_ingress_security_rules_stateless}"
	}
	ingress_security_rules {
		#Required
		protocol = "6"
		source = "${var.security_list_ingress_security_rules_source}"

		#Optional
		source_type = "${var.security_list_ingress_security_rules_source_type}"
		stateless = "${var.security_list_ingress_security_rules_stateless}"
		tcp_options {

			#Optional
			max = "${var.security_list_ingress_security_rules_tcp_options_destination_port_range_max}"
			min = "${var.security_list_ingress_security_rules_tcp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_ingress_security_rules_tcp_options_source_port_range_max}"
				min = "${var.security_list_ingress_security_rules_tcp_options_source_port_range_min}"
			}
		}
	}
	ingress_security_rules {
		#Required
		protocol = "17"
		source = "${var.security_list_ingress_security_rules_source}"

		#Optional
		source_type = "${var.security_list_ingress_security_rules_source_type}"
		stateless = "${var.security_list_ingress_security_rules_stateless}"
		udp_options {

			#Optional
			max = "${var.security_list_ingress_security_rules_udp_options_destination_port_range_max}"
			min = "${var.security_list_ingress_security_rules_udp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_ingress_security_rules_udp_options_source_port_range_max}"
				min = "${var.security_list_ingress_security_rules_udp_options_source_port_range_min}"
			}
		}
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.security_list_defined_tags_value}")}"
	display_name = "${var.security_list_display_name}"
	freeform_tags = "${var.security_list_freeform_tags}"
}
`

	SecurityListResourceConfigWithServiceCidrBlock = SecurityListResourceDependencies + `
resource "oci_core_security_list" "test_security_list" {
	#Required
	compartment_id = "${var.compartment_id}"
	egress_security_rules {
		#Required
		destination = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
		protocol = "1"

		#Optional
		destination_type = "${var.security_list_egress_security_rules_destination_type}"
		icmp_options {
			#Required
			type = "${var.security_list_egress_security_rules_icmp_options_type}"

			#Optional
			code = "${var.security_list_egress_security_rules_icmp_options_code}"
		}
		stateless = "${var.security_list_egress_security_rules_stateless}"
	}
	egress_security_rules {
		#Required
		destination = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
		protocol = "6"

		#Optional
		destination_type = "${var.security_list_egress_security_rules_destination_type}"
		stateless = "${var.security_list_egress_security_rules_stateless}"
		tcp_options {

			#Optional
			max = "${var.security_list_egress_security_rules_tcp_options_destination_port_range_max}"
			min = "${var.security_list_egress_security_rules_tcp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_egress_security_rules_tcp_options_source_port_range_max}"
				min = "${var.security_list_egress_security_rules_tcp_options_source_port_range_min}"
			}
		}
	}
	egress_security_rules {
		#Required
		destination = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
		protocol = "17"

		#Optional
		destination_type = "${var.security_list_egress_security_rules_destination_type}"
		stateless = "${var.security_list_egress_security_rules_stateless}"
		udp_options {

			#Optional
			max = "${var.security_list_egress_security_rules_udp_options_destination_port_range_max}"
			min = "${var.security_list_egress_security_rules_udp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_egress_security_rules_udp_options_source_port_range_max}"
				min = "${var.security_list_egress_security_rules_udp_options_source_port_range_min}"
			}
		}
	}
	ingress_security_rules {
		#Required
		protocol = "1"
		source = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"

		#Optional
		icmp_options {
			#Required
			type = "${var.security_list_ingress_security_rules_icmp_options_type}"

			#Optional
			code = "${var.security_list_ingress_security_rules_icmp_options_code}"
		}
		source_type = "${var.security_list_ingress_security_rules_source_type}"
		stateless = "${var.security_list_ingress_security_rules_stateless}"
	}
	ingress_security_rules {
		#Required
		protocol = "6"
		source = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"

		#Optional
		source_type = "${var.security_list_ingress_security_rules_source_type}"
		stateless = "${var.security_list_ingress_security_rules_stateless}"
		tcp_options {

			#Optional
			max = "${var.security_list_ingress_security_rules_tcp_options_destination_port_range_max}"
			min = "${var.security_list_ingress_security_rules_tcp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_ingress_security_rules_tcp_options_source_port_range_max}"
				min = "${var.security_list_ingress_security_rules_tcp_options_source_port_range_min}"
			}
		}
	}
	ingress_security_rules {
		#Required
		protocol = "17"
		source = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"

		#Optional
		source_type = "${var.security_list_ingress_security_rules_source_type}"
		stateless = "${var.security_list_ingress_security_rules_stateless}"
		udp_options {

			#Optional
			max = "${var.security_list_ingress_security_rules_udp_options_destination_port_range_max}"
			min = "${var.security_list_ingress_security_rules_udp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_ingress_security_rules_udp_options_source_port_range_max}"
				min = "${var.security_list_ingress_security_rules_udp_options_source_port_range_min}"
			}
		}
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.security_list_defined_tags_value}")}"
	display_name = "${var.security_list_display_name}"
	freeform_tags = "${var.security_list_freeform_tags}"
}

data "oci_core_services" "test_services" {
}
`
	SecurityListPropertyVariables = `
variable "security_list_defined_tags_value" { default = "value" }
variable "security_list_display_name" { default = "MyPrivateSubnetSecurityList" }
variable "security_list_egress_security_rules_destination" { default = "10.0.2.0/24" }
variable "security_list_egress_security_rules_destination_type" { default = "CIDR_BLOCK" }
variable "security_list_egress_security_rules_icmp_options_code" { default = 4 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_egress_security_rules_protocol" { default = "1" }
variable "security_list_egress_security_rules_stateless" { default = false }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_source_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_source_port_range_min" { default = "1521" }
variable "security_list_freeform_tags" { default = {"Department"= "Finance"} }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 4 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_ingress_security_rules_protocol" { default = "1" }
variable "security_list_ingress_security_rules_source" { default = "10.0.1.0/24" }
variable "security_list_ingress_security_rules_source_type" { default = "CIDR_BLOCK" }
variable "security_list_ingress_security_rules_stateless" { default = false }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_min" { default = "1521" }
variable "security_list_state" { default = "AVAILABLE" }

`
	SecurityListResourceDependencies = VcnPropertyVariables + VcnResourceConfig
)

func TestCoreSecurityListResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_security_list.test_security_list"
	datasourceName := "data.oci_core_security_lists.test_security_lists"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreSecurityListDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + SecurityListPropertyVariables + compartmentIdVariableStr + SecurityListRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
						"destination": "10.0.2.0/24",
						"protocol":    "1",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
						"protocol": "1",
						"source":   "10.0.1.0/24",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + SecurityListResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + SecurityListPropertyVariables + compartmentIdVariableStr + SecurityListResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyPrivateSubnetSecurityList"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "3"),
					CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
						"destination":         "10.0.2.0/24",
						"destination_type":    "CIDR_BLOCK",
						"icmp_options.#":      "1",
						"icmp_options.0.code": "4",
						"icmp_options.0.type": "3",
						"protocol":            "1",
						"stateless":           "false",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
						"destination":                           "10.0.2.0/24",
						"destination_type":                      "CIDR_BLOCK",
						"protocol":                              "6",
						"stateless":                             "false",
						"tcp_options.#":                         "1",
						"tcp_options.0.max":                     "1521",
						"tcp_options.0.min":                     "1521",
						"tcp_options.0.source_port_range.#":     "1",
						"tcp_options.0.source_port_range.0.max": "1521",
						"tcp_options.0.source_port_range.0.min": "1521",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
						"destination":                           "10.0.2.0/24",
						"destination_type":                      "CIDR_BLOCK",
						"protocol":                              "17",
						"stateless":                             "false",
						"udp_options.#":                         "1",
						"udp_options.0.max":                     "1521",
						"udp_options.0.min":                     "1521",
						"udp_options.0.source_port_range.#":     "1",
						"udp_options.0.source_port_range.0.max": "1521",
						"udp_options.0.source_port_range.0.min": "1521",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "3"),
					CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
						"icmp_options.#":      "1",
						"icmp_options.0.code": "4",
						"icmp_options.0.type": "3",
						"protocol":            "1",
						"source":              "10.0.1.0/24",
						"source_type":         "CIDR_BLOCK",
						"stateless":           "false",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
						"protocol":                              "6",
						"source":                                "10.0.1.0/24",
						"source_type":                           "CIDR_BLOCK",
						"stateless":                             "false",
						"tcp_options.#":                         "1",
						"tcp_options.0.max":                     "1521",
						"tcp_options.0.min":                     "1521",
						"tcp_options.0.source_port_range.#":     "1",
						"tcp_options.0.source_port_range.0.max": "1521",
						"tcp_options.0.source_port_range.0.min": "1521",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
						"protocol":                              "17",
						"source":                                "10.0.1.0/24",
						"source_type":                           "CIDR_BLOCK",
						"stateless":                             "false",
						"udp_options.#":                         "1",
						"udp_options.0.max":                     "1521",
						"udp_options.0.min":                     "1521",
						"udp_options.0.source_port_range.#":     "1",
						"udp_options.0.source_port_range.0.max": "1521",
						"udp_options.0.source_port_range.0.min": "1521",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "security_list_defined_tags_value" { default = "updatedValue" }
variable "security_list_display_name" { default = "displayName2" }
variable "security_list_egress_security_rules_destination" { default = "10.0.2.0/24" }
variable "security_list_egress_security_rules_destination_type" { default = "SERVICE_CIDR_BLOCK" }
variable "security_list_egress_security_rules_icmp_options_code" { default = 0 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_egress_security_rules_protocol" { default = "1" }
variable "security_list_egress_security_rules_stateless" { default = true }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_max" { default = "1522" }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_min" { default = "1522" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_max" { default = "1522" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_min" { default = "1522" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_max" { default = "1522" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_min" { default = "1522" }
variable "security_list_egress_security_rules_udp_options_source_port_range_max" { default = "1522" }
variable "security_list_egress_security_rules_udp_options_source_port_range_min" { default = "1522" }
variable "security_list_freeform_tags" { default = {"Department"= "Accounting"} }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 0 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_ingress_security_rules_protocol" { default = "1" }
variable "security_list_ingress_security_rules_source" { default = "10.0.1.0/24" }
variable "security_list_ingress_security_rules_source_type" { default = "SERVICE_CIDR_BLOCK" }
variable "security_list_ingress_security_rules_stateless" { default = true }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_max" { default = "1522" }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_min" { default = "1522" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_max" { default = "1522" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_min" { default = "1522" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_max" { default = "1522" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_min" { default = "1522" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_max" { default = "1522" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_min" { default = "1522" }
variable "security_list_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + SecurityListResourceConfigWithServiceCidrBlock,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "3"),
					CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
						"destination_type":    "SERVICE_CIDR_BLOCK",
						"icmp_options.#":      "1",
						"icmp_options.0.code": "0",
						"icmp_options.0.type": "3",
						"protocol":            "1",
						"stateless":           "true",
					},
						[]string{
							"destination",
						}),
					CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
						"destination_type":                      "SERVICE_CIDR_BLOCK",
						"protocol":                              "6",
						"stateless":                             "true",
						"tcp_options.#":                         "1",
						"tcp_options.0.max":                     "1522",
						"tcp_options.0.min":                     "1522",
						"tcp_options.0.source_port_range.#":     "1",
						"tcp_options.0.source_port_range.0.max": "1522",
						"tcp_options.0.source_port_range.0.min": "1522",
					},
						[]string{
							"destination",
						}),
					CheckResourceSetContainsElementWithProperties(resourceName, "egress_security_rules", map[string]string{
						"destination_type":                      "SERVICE_CIDR_BLOCK",
						"protocol":                              "17",
						"stateless":                             "true",
						"udp_options.#":                         "1",
						"udp_options.0.max":                     "1522",
						"udp_options.0.min":                     "1522",
						"udp_options.0.source_port_range.#":     "1",
						"udp_options.0.source_port_range.0.max": "1522",
						"udp_options.0.source_port_range.0.min": "1522",
					},
						[]string{
							"destination",
						}),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "3"),
					CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
						"icmp_options.#":      "1",
						"icmp_options.0.code": "0",
						"icmp_options.0.type": "3",
						"protocol":            "1",
						"source_type":         "SERVICE_CIDR_BLOCK",
						"stateless":           "true",
					},
						[]string{
							"source",
						}),
					CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
						"protocol":                              "6",
						"source_type":                           "SERVICE_CIDR_BLOCK",
						"stateless":                             "true",
						"tcp_options.#":                         "1",
						"tcp_options.0.max":                     "1522",
						"tcp_options.0.min":                     "1522",
						"tcp_options.0.source_port_range.#":     "1",
						"tcp_options.0.source_port_range.0.max": "1522",
						"tcp_options.0.source_port_range.0.min": "1522",
					},
						[]string{
							"source",
						}),
					CheckResourceSetContainsElementWithProperties(resourceName, "ingress_security_rules", map[string]string{
						"protocol":                              "17",
						"source_type":                           "SERVICE_CIDR_BLOCK",
						"stateless":                             "true",
						"udp_options.#":                         "1",
						"udp_options.0.max":                     "1522",
						"udp_options.0.min":                     "1522",
						"udp_options.0.source_port_range.#":     "1",
						"udp_options.0.source_port_range.0.max": "1522",
						"udp_options.0.source_port_range.0.min": "1522",
					},
						[]string{
							"source",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "security_list_defined_tags_value" { default = "updatedValue" }
variable "security_list_display_name" { default = "displayName2" }
variable "security_list_egress_security_rules_destination" { default = "10.0.2.0/24" }
variable "security_list_egress_security_rules_destination_type" { default = "SERVICE_CIDR_BLOCK" }
variable "security_list_egress_security_rules_icmp_options_code" { default = 0 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_egress_security_rules_protocol" { default = "1" }
variable "security_list_egress_security_rules_stateless" { default = true }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_max" { default = "1522" }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_min" { default = "1522" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_max" { default = "1522" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_min" { default = "1522" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_max" { default = "1522" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_min" { default = "1522" }
variable "security_list_egress_security_rules_udp_options_source_port_range_max" { default = "1522" }
variable "security_list_egress_security_rules_udp_options_source_port_range_min" { default = "1522" }
variable "security_list_freeform_tags" { default = {"Department"= "Accounting"} }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 0 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_ingress_security_rules_protocol" { default = "1" }
variable "security_list_ingress_security_rules_source" { default = "10.0.1.0/24" }
variable "security_list_ingress_security_rules_source_type" { default = "SERVICE_CIDR_BLOCK" }
variable "security_list_ingress_security_rules_stateless" { default = true }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_max" { default = "1522" }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_min" { default = "1522" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_max" { default = "1522" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_min" { default = "1522" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_max" { default = "1522" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_min" { default = "1522" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_max" { default = "1522" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_min" { default = "1522" }
variable "security_list_state" { default = "AVAILABLE" }

data "oci_core_security_lists" "test_security_lists" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.security_list_display_name}"
	state = "${var.security_list_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_security_list.test_security_list.id}"]
    }
}
                ` + compartmentIdVariableStr + SecurityListResourceConfigWithServiceCidrBlock,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "security_lists.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.#", "3"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.egress_security_rules", map[string]string{
						"destination_type":    "SERVICE_CIDR_BLOCK",
						"icmp_options.#":      "1",
						"icmp_options.0.code": "0",
						"icmp_options.0.type": "3",
						"protocol":            "1",
						"stateless":           "true",
					},
						[]string{
							"destination",
						}),
					CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.egress_security_rules", map[string]string{
						"destination_type":                      "SERVICE_CIDR_BLOCK",
						"protocol":                              "6",
						"stateless":                             "true",
						"tcp_options.#":                         "1",
						"tcp_options.0.max":                     "1522",
						"tcp_options.0.min":                     "1522",
						"tcp_options.0.source_port_range.#":     "1",
						"tcp_options.0.source_port_range.0.max": "1522",
						"tcp_options.0.source_port_range.0.min": "1522",
					},
						[]string{
							"destination",
						}),
					CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.egress_security_rules", map[string]string{
						"destination_type":                      "SERVICE_CIDR_BLOCK",
						"protocol":                              "17",
						"stateless":                             "true",
						"udp_options.#":                         "1",
						"udp_options.0.max":                     "1522",
						"udp_options.0.min":                     "1522",
						"udp_options.0.source_port_range.#":     "1",
						"udp_options.0.source_port_range.0.max": "1522",
						"udp_options.0.source_port_range.0.min": "1522",
					},
						[]string{
							"destination",
						}),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.#", "3"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.ingress_security_rules", map[string]string{
						"icmp_options.#":      "1",
						"icmp_options.0.code": "0",
						"icmp_options.0.type": "3",
						"protocol":            "1",
						"source_type":         "SERVICE_CIDR_BLOCK",
						"stateless":           "true",
					},
						[]string{
							"source",
						}),
					CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.ingress_security_rules", map[string]string{
						"protocol":                              "6",
						"source_type":                           "SERVICE_CIDR_BLOCK",
						"stateless":                             "true",
						"tcp_options.#":                         "1",
						"tcp_options.0.max":                     "1522",
						"tcp_options.0.min":                     "1522",
						"tcp_options.0.source_port_range.#":     "1",
						"tcp_options.0.source_port_range.0.max": "1522",
						"tcp_options.0.source_port_range.0.min": "1522",
					},
						[]string{
							"source",
						}),
					CheckResourceSetContainsElementWithProperties(datasourceName, "security_lists.0.ingress_security_rules", map[string]string{
						"protocol":                              "17",
						"source_type":                           "SERVICE_CIDR_BLOCK",
						"stateless":                             "true",
						"udp_options.#":                         "1",
						"udp_options.0.max":                     "1522",
						"udp_options.0.min":                     "1522",
						"udp_options.0.source_port_range.#":     "1",
						"udp_options.0.source_port_range.0.max": "1522",
						"udp_options.0.source_port_range.0.min": "1522",
					},
						[]string{
							"source",
						}),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.vcn_id"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ResourceName:      resourceName,
			},
		},
	})
}

func testAccCheckCoreSecurityListDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_security_list" {
			noResourceFound = false
			request := oci_core.GetSecurityListRequest{}

			tmp := rs.Primary.ID
			request.SecurityListId = &tmp

			_, err := client.GetSecurityList(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}
			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
