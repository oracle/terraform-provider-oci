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
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.security_list_defined_tags_value}")}"
	display_name = "${var.security_list_display_name}"
	freeform_tags = "${var.security_list_freeform_tags}"
}
`
	SecurityListPropertyVariables = `
variable "security_list_defined_tags_value" { default = "value" }
variable "security_list_display_name" { default = "MyPrivateSubnetSecurityList" }
variable "security_list_egress_security_rules_destination" { default = "10.0.2.0/24" }
variable "security_list_egress_security_rules_icmp_options_code" { default = 4 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_egress_security_rules_protocol" { default = "1" }
variable "security_list_egress_security_rules_stateless" { default = false }
variable "security_list_freeform_tags" { default = {"Department"= "Finance"} }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 4 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_ingress_security_rules_protocol" { default = "1" }
variable "security_list_ingress_security_rules_source" { default = "10.0.1.0/24" }
variable "security_list_ingress_security_rules_stateless" { default = false }
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
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.protocol", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.protocol", "1"),
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
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyPrivateSubnetSecurityList"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.destination", "10.0.2.0/24"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.code", "4"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.type", "3"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.protocol", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.code", "4"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.type", "3"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.protocol", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.source", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.stateless", "false"),
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
variable "security_list_egress_security_rules_icmp_options_code" { default = 0 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_egress_security_rules_protocol" { default = "1" }
variable "security_list_egress_security_rules_stateless" { default = true }
variable "security_list_freeform_tags" { default = {"Department"= "Accounting"} }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 0 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_ingress_security_rules_protocol" { default = "1" }
variable "security_list_ingress_security_rules_source" { default = "10.0.1.0/24" }
variable "security_list_ingress_security_rules_stateless" { default = true }
variable "security_list_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + SecurityListResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.destination", "10.0.2.0/24"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.code", "0"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.icmp_options.0.type", "3"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.protocol", "1"),
					resource.TestCheckResourceAttr(resourceName, "egress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.code", "0"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.icmp_options.0.type", "3"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.protocol", "1"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.source", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(resourceName, "ingress_security_rules.0.stateless", "true"),
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
variable "security_list_egress_security_rules_icmp_options_code" { default = 0 }
variable "security_list_egress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_egress_security_rules_protocol" { default = "1" }
variable "security_list_egress_security_rules_stateless" { default = true }
variable "security_list_freeform_tags" { default = {"Department"= "Accounting"} }
variable "security_list_ingress_security_rules_icmp_options_code" { default = 0 }
variable "security_list_ingress_security_rules_icmp_options_type" { default = 3 }
variable "security_list_ingress_security_rules_protocol" { default = "1" }
variable "security_list_ingress_security_rules_source" { default = "10.0.1.0/24" }
variable "security_list_ingress_security_rules_stateless" { default = true }
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
                ` + compartmentIdVariableStr + SecurityListResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "security_lists.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.destination", "10.0.2.0/24"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.icmp_options.0.code", "0"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.icmp_options.0.type", "3"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.protocol", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.egress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.icmp_options.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.icmp_options.0.code", "0"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.icmp_options.0.type", "3"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.protocol", "1"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.source", "10.0.1.0/24"),
					resource.TestCheckResourceAttr(datasourceName, "security_lists.0.ingress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "security_lists.0.vcn_id"),
				),
			},
		},
	})
}
