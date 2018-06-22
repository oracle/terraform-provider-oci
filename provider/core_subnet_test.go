// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	SubnetRequiredOnlyResource = SubnetRequiredOnlyResourceDependencies + `
resource "oci_core_subnet" "test_subnet" {
	#Required
	availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
	cidr_block = "${var.subnet_cidr_block}"
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`

	SubnetResourceConfig = SubnetResourceDependencies + `
resource "oci_core_subnet" "test_subnet" {
	#Required
	availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
	cidr_block = "${var.subnet_cidr_block}"
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.subnet_defined_tags_value}")}"
	dhcp_options_id = "${oci_core_dhcp_options.test_dhcp_options.id}"
	display_name = "${var.subnet_display_name}"
	dns_label = "${var.subnet_dns_label}"
	freeform_tags = "${var.subnet_freeform_tags}"
	prohibit_public_ip_on_vnic = "${var.subnet_prohibit_public_ip_on_vnic}"
	route_table_id = "${oci_core_route_table.test_route_table.id}"
	security_list_ids = ["${oci_core_vcn.test_vcn.default_security_list_id}"]
}
`
	SubnetPropertyVariables = `
variable "subnet_cidr_block" { default = "10.0.0.0/16" }
variable "subnet_defined_tags_value" { default = "value" }
variable "subnet_display_name" { default = "MySubnet" }
variable "subnet_dns_label" { default = "dnslabel" }
variable "subnet_freeform_tags" { default = {"Department"= "Finance"} }
variable "subnet_prohibit_public_ip_on_vnic" { default = false }
variable "subnet_security_list_ids" { default = [] }
variable "subnet_state" { default = "AVAILABLE" }

`
	SubnetRequiredOnlyResourceDependencies = AvailabilityDomainConfig + VcnPropertyVariables + VcnRequiredOnlyResource
	SubnetResourceDependencies             = AvailabilityDomainConfig + DhcpOptionsPropertyVariables + DhcpOptionsResourceConfigOnly + RouteTablePropertyVariables + RouteTableResourceConfig
)

func TestCoreSubnetResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_subnet.test_subnet"
	datasourceName := "data.oci_core_subnets.test_subnets"

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
				Config:            config + SubnetPropertyVariables + compartmentIdVariableStr + SubnetRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
					resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + SubnetResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + SubnetPropertyVariables + compartmentIdVariableStr + SubnetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MySubnet"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
					resource.TestCheckResourceAttr(resourceName, "security_list_ids.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "subnet_cidr_block" { default = "10.0.0.0/16" }
variable "subnet_defined_tags_value" { default = "updatedValue" }
variable "subnet_display_name" { default = "displayName2" }
variable "subnet_dns_label" { default = "dnslabel" }
variable "subnet_freeform_tags" { default = {"Department"= "Accounting"} }
variable "subnet_prohibit_public_ip_on_vnic" { default = false }
variable "subnet_security_list_ids" { default = [] }
variable "subnet_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + SubnetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
					resource.TestCheckResourceAttr(resourceName, "security_list_ids.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

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
variable "subnet_cidr_block" { default = "10.0.0.0/16" }
variable "subnet_defined_tags_value" { default = "updatedValue" }
variable "subnet_display_name" { default = "displayName2" }
variable "subnet_dns_label" { default = "dnslabel" }
variable "subnet_freeform_tags" { default = {"Department"= "Accounting"} }
variable "subnet_prohibit_public_ip_on_vnic" { default = false }
variable "subnet_security_list_ids" { default = [] }
variable "subnet_state" { default = "AVAILABLE" }

data "oci_core_subnets" "test_subnets" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.subnet_display_name}"
	state = "${var.subnet_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_subnet.test_subnet.id}"]
    }
}
                ` + compartmentIdVariableStr + SubnetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "subnets.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "subnets.0.cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(datasourceName, "subnets.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "subnets.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.dhcp_options_id"),
					resource.TestCheckResourceAttr(datasourceName, "subnets.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "subnets.0.dns_label", "dnslabel"),
					resource.TestCheckResourceAttr(datasourceName, "subnets.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "subnets.0.prohibit_public_ip_on_vnic", "false"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.route_table_id"),
					resource.TestCheckResourceAttr(datasourceName, "subnets.0.security_list_ids.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.vcn_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.virtual_router_ip"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.virtual_router_mac"),
				),
			},
		},
	})
}
