// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
		protocol = "${var.security_list_egress_security_rules_protocol}"

		#Optional
		icmp_options {
			#Required
			type = "${var.security_list_egress_security_rules_icmp_options_type}"

			#Optional
			code = "${var.security_list_egress_security_rules_icmp_options_code}"
		}
		stateless = "${var.security_list_egress_security_rules_stateless}"
		tcp_options {

			#Optional
			destination_port_range {
				#Required
				max = "${var.security_list_egress_security_rules_tcp_options_destination_port_range_max}"
				min = "${var.security_list_egress_security_rules_tcp_options_destination_port_range_min}"
			}
			source_port_range {
				#Required
				max = "${var.security_list_egress_security_rules_tcp_options_source_port_range_max}"
				min = "${var.security_list_egress_security_rules_tcp_options_source_port_range_min}"
			}
		}
		udp_options {

			#Optional
			destination_port_range {
				#Required
				max = "${var.security_list_egress_security_rules_udp_options_destination_port_range_max}"
				min = "${var.security_list_egress_security_rules_udp_options_destination_port_range_min}"
			}
			source_port_range {
				#Required
				max = "${var.security_list_egress_security_rules_udp_options_source_port_range_max}"
				min = "${var.security_list_egress_security_rules_udp_options_source_port_range_min}"
			}
		}
	}
	ingress_security_rules {
		#Required
		protocol = "${var.security_list_ingress_security_rules_protocol}"
		source = "${var.security_list_ingress_security_rules_source}"

		#Optional
		icmp_options {
			#Required
			type = "${var.security_list_ingress_security_rules_icmp_options_type}"

			#Optional
			code = "${var.security_list_ingress_security_rules_icmp_options_code}"
		}
		stateless = "${var.security_list_ingress_security_rules_stateless}"
		tcp_options {

			#Optional
			destination_port_range {
				#Required
				max = "${var.security_list_ingress_security_rules_tcp_options_destination_port_range_max}"
				min = "${var.security_list_ingress_security_rules_tcp_options_destination_port_range_min}"
			}
			source_port_range {
				#Required
				max = "${var.security_list_ingress_security_rules_tcp_options_source_port_range_max}"
				min = "${var.security_list_ingress_security_rules_tcp_options_source_port_range_min}"
			}
		}
		udp_options {

			#Optional
			destination_port_range {
				#Required
				max = "${var.security_list_ingress_security_rules_udp_options_destination_port_range_max}"
				min = "${var.security_list_ingress_security_rules_udp_options_destination_port_range_min}"
			}
			source_port_range {
				#Required
				max = "${var.security_list_ingress_security_rules_udp_options_source_port_range_max}"
				min = "${var.security_list_ingress_security_rules_udp_options_source_port_range_min}"
			}
		}
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.security_list_display_name}"
}
`
	SecurityListPropertyVariables = `
variable "security_list_display_name" { default = "MyPrivateSubnetSecurityList" }
variable "security_list_egress_security_rules_destination" { default = "10.0.2.0/24" }
variable "security_list_egress_security_rules_icmp_options_code" { default = 10 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 10 }
variable "security_list_egress_security_rules_protocol" { default = "6" }
variable "security_list_egress_security_rules_stateless" { default = false }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_source_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_source_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 10 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 10 }
variable "security_list_ingress_security_rules_protocol" { default = "6" }
variable "security_list_ingress_security_rules_source" { default = "10.0.1.0/24" }
variable "security_list_ingress_security_rules_stateless" { default = false }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_min" { default = "1521" }
variable "security_list_state" { default = "state" }

`
	SecurityListResourceDependencies = VcnPropertyVariables + VcnResourceConfig
)

func TestCoreSecurityListResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_security_list.test_security_list"
	datasourceName := "data.oci_core_security_lists.test_security_lists"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + SecurityListPropertyVariables + compartmentIdVariableStr + SecurityListRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.destination", "10.0.2.0/24"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.source", "10.0.1.0/24"),
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
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyPrivateSubnetSecurityList"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.destination", "10.0.2.0/24"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.code", "10"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.type", "10"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.code", "10"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.type", "10"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.source", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.min", "1521"),
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
variable "security_list_display_name" { default = "displayName2" }
variable "security_list_egress_security_rules_destination" { default = "10.0.2.0/24" }
variable "security_list_egress_security_rules_icmp_options_code" { default = 11 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 10 }
variable "security_list_egress_security_rules_protocol" { default = "6" }
variable "security_list_egress_security_rules_stateless" { default = true }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_source_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_source_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 11 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 10 }
variable "security_list_ingress_security_rules_protocol" { default = "6" }
variable "security_list_ingress_security_rules_source" { default = "10.0.1.0/24" }
variable "security_list_ingress_security_rules_stateless" { default = true }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_min" { default = "1521" }
variable "security_list_state" { default = "state" }

                ` + compartmentIdVariableStr + SecurityListResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.destination", "10.0.2.0/24"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.code", "11"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.type", "10"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.code", "11"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.type", "10"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.source", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.min", "1521"),
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
			// verify updates to Force New parameters.
			{
				Config: config + `
variable "security_list_display_name" { default = "displayName2" }
variable "security_list_egress_security_rules_destination" { default = "destination2" }
variable "security_list_egress_security_rules_icmp_options_code" { default = 11 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 11 }
variable "security_list_egress_security_rules_protocol" { default = "protocol2" }
variable "security_list_egress_security_rules_stateless" { default = true }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_max" { default = 11 }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_min" { default = 11 }
variable "security_list_egress_security_rules_tcp_options_source_port_range_max" { default = 11 }
variable "security_list_egress_security_rules_tcp_options_source_port_range_min" { default = 11 }
variable "security_list_egress_security_rules_udp_options_destination_port_range_max" { default = 11 }
variable "security_list_egress_security_rules_udp_options_destination_port_range_min" { default = 11 }
variable "security_list_egress_security_rules_udp_options_source_port_range_max" { default = 11 }
variable "security_list_egress_security_rules_udp_options_source_port_range_min" { default = 11 }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 11 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 11 }
variable "security_list_ingress_security_rules_protocol" { default = "protocol2" }
variable "security_list_ingress_security_rules_source" { default = "source2" }
variable "security_list_ingress_security_rules_stateless" { default = true }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_max" { default = 11 }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_min" { default = 11 }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_max" { default = 11 }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_min" { default = 11 }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_max" { default = 11 }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_min" { default = 11 }
variable "security_list_ingress_security_rules_udp_options_source_port_range_max" { default = 11 }
variable "security_list_ingress_security_rules_udp_options_source_port_range_min" { default = 11 }
variable "security_list_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr2 + SecurityListResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.destination", "destination2"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.code", "11"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.type", "11"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.protocol", "protocol2"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.min", "11"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.code", "11"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.type", "11"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.protocol", "protocol2"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.source", "source2"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.min", "11"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "security_list_display_name" { default = "displayName2" }
variable "security_list_egress_security_rules_destination" { default = "destination2" }
variable "security_list_egress_security_rules_icmp_options_code" { default = 11 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 11 }
variable "security_list_egress_security_rules_protocol" { default = "protocol2" }
variable "security_list_egress_security_rules_stateless" { default = true }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_max" { default = 11 }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_min" { default = 11 }
variable "security_list_egress_security_rules_tcp_options_source_port_range_max" { default = 11 }
variable "security_list_egress_security_rules_tcp_options_source_port_range_min" { default = 11 }
variable "security_list_egress_security_rules_udp_options_destination_port_range_max" { default = 11 }
variable "security_list_egress_security_rules_udp_options_destination_port_range_min" { default = 11 }
variable "security_list_egress_security_rules_udp_options_source_port_range_max" { default = 11 }
variable "security_list_egress_security_rules_udp_options_source_port_range_min" { default = 11 }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 11 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 11 }
variable "security_list_ingress_security_rules_protocol" { default = "protocol2" }
variable "security_list_ingress_security_rules_source" { default = "source2" }
variable "security_list_ingress_security_rules_stateless" { default = true }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_max" { default = 11 }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_min" { default = 11 }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_max" { default = 11 }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_min" { default = 11 }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_max" { default = 11 }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_min" { default = 11 }
variable "security_list_ingress_security_rules_udp_options_source_port_range_max" { default = 11 }
variable "security_list_ingress_security_rules_udp_options_source_port_range_min" { default = 11 }
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
                ` + compartmentIdVariableStr2 + SecurityListResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "security_lists.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.destination", "destination2"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.icmp_options.0.code", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.icmp_options.0.type", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.protocol", "protocol2"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.tcp_options.0.destination_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.tcp_options.0.destination_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.tcp_options.0.source_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.tcp_options.0.source_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.udp_options.0.destination_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.udp_options.0.destination_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.udp_options.0.source_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.udp_options.0.source_port_range.0.min", "11"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.icmp_options.0.code", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.icmp_options.0.type", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.protocol", "protocol2"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.source", "source2"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.tcp_options.0.destination_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.tcp_options.0.destination_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.tcp_options.0.source_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.tcp_options.0.source_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.udp_options.0.destination_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.udp_options.0.destination_port_range.0.min", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.udp_options.0.source_port_range.0.max", "11"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.udp_options.0.source_port_range.0.min", "11"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.vcn_id"),
				),
			},
		},
	})
}

func TestCoreSecurityListResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_security_list.test_security_list"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + SecurityListPropertyVariables + compartmentIdVariableStr + SecurityListResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyPrivateSubnetSecurityList"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.destination", "10.0.2.0/24"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.code", "10"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.type", "10"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.code", "10"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.type", "10"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.source", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "security_list_display_name" { default = "MyPrivateSubnetSecurityList" }
variable "security_list_egress_security_rules_destination" { default = "10.0.2.0/24" }
variable "security_list_egress_security_rules_icmp_options_code" { default = 10 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 10 }
variable "security_list_egress_security_rules_protocol" { default = "6" }
variable "security_list_egress_security_rules_stateless" { default = false }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_source_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_source_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 10 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 10 }
variable "security_list_ingress_security_rules_protocol" { default = "6" }
variable "security_list_ingress_security_rules_source" { default = "10.0.1.0/24" }
variable "security_list_ingress_security_rules_stateless" { default = false }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_min" { default = "1521" }
variable "security_list_state" { default = "state" }
				` + compartmentIdVariableStr2 + SecurityListResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyPrivateSubnetSecurityList"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.destination", "10.0.2.0/24"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.code", "10"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.type", "10"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.code", "10"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.type", "10"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.source", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter CompartmentId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "security_list_display_name" { default = "MyPrivateSubnetSecurityList" }
variable "security_list_egress_security_rules_destination" { default = "10.0.2.0/24" }
variable "security_list_egress_security_rules_icmp_options_code" { default = 10 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 10 }
variable "security_list_egress_security_rules_protocol" { default = "6" }
variable "security_list_egress_security_rules_stateless" { default = false }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_destination_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_tcp_options_source_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_destination_port_range_min" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_source_port_range_max" { default = "1521" }
variable "security_list_egress_security_rules_udp_options_source_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 10 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 10 }
variable "security_list_ingress_security_rules_protocol" { default = "6" }
variable "security_list_ingress_security_rules_source" { default = "10.0.1.0/24" }
variable "security_list_ingress_security_rules_stateless" { default = false }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_destination_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_tcp_options_source_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_destination_port_range_min" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_max" { default = "1521" }
variable "security_list_ingress_security_rules_udp_options_source_port_range_min" { default = "1521" }
variable "security_list_state" { default = "state" }
				` + compartmentIdVariableStr2 + SecurityListResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyPrivateSubnetSecurityList"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.destination", "10.0.2.0/24"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.code", "10"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.type", "10"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.tcp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.udp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.code", "10"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.type", "10"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.protocol", "6"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.source", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.tcp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.destination_port_range.0.min", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.max", "1521"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.udp_options.0.source_port_range.0.min", "1521"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter VcnId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
