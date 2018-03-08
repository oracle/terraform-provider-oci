// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	VcnRequiredOnlyResource = VcnResourceDependencies + `
resource "oci_core_vcn" "test_vcn" {
	#Required
	cidr_block = "${var.vcn_cidr_block}"
	compartment_id = "${var.compartment_id}"
}
`

	VcnResourceConfig = VcnResourceDependencies + `
resource "oci_core_vcn" "test_vcn" {
	#Required
	cidr_block = "${var.vcn_cidr_block}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.vcn_display_name}"
	dns_label = "${var.vcn_dns_label}"
}
`
	VcnPropertyVariables = `
variable "vcn_cidr_block" { default = "10.0.0.0/16" }
variable "vcn_display_name" { default = "displayName" }
variable "vcn_dns_label" { default = "dnslabel" }
variable "vcn_state" { default = "state" }

`
	VcnResourceDependencies = ""
)

func TestCoreVcnResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_vcn.test_vcn"
	datasourceName := "data.oci_core_vcns.test_vcns"

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
				Config:            config + VcnPropertyVariables + compartmentIdVariableStr + VcnRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VcnResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + VcnPropertyVariables + compartmentIdVariableStr + VcnResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "vcn_cidr_block" { default = "10.0.0.0/16" }
variable "vcn_display_name" { default = "displayName2" }
variable "vcn_dns_label" { default = "dnslabel" }
variable "vcn_state" { default = "state" }

                ` + compartmentIdVariableStr + VcnResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
variable "vcn_cidr_block" { default = "10.0.0.0/24" }
variable "vcn_display_name" { default = "displayName2" }
variable "vcn_dns_label" { default = "dnslabel2" }
variable "vcn_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr2 + VcnResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
variable "vcn_cidr_block" { default = "10.0.0.0/24" }
variable "vcn_display_name" { default = "displayName2" }
variable "vcn_dns_label" { default = "dnslabel2" }
variable "vcn_state" { default = "AVAILABLE" }

data "oci_core_vcns" "test_vcns" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.vcn_display_name}"
	state = "${var.vcn_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_vcn.test_vcn.id}"]
    }
}
                ` + compartmentIdVariableStr2 + VcnResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.cidr_block", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "virtual_networks.0.dns_label", "dnslabel2"),
					resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "virtual_networks.0.state"),
				),
			},
		},
	})
}

func TestCoreVcnResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_vcn.test_vcn"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + VcnPropertyVariables + compartmentIdVariableStr + VcnResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "vcn_cidr_block" { default = "10.0.0.0/24" }
variable "vcn_display_name" { default = "displayName" }
variable "vcn_dns_label" { default = "dnslabel" }
variable "vcn_state" { default = "state" }
				` + compartmentIdVariableStr + VcnResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter CidrBlock but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "vcn_cidr_block" { default = "10.0.0.0/24" }
variable "vcn_display_name" { default = "displayName" }
variable "vcn_dns_label" { default = "dnslabel" }
variable "vcn_state" { default = "state" }
				` + compartmentIdVariableStr2 + VcnResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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
variable "vcn_cidr_block" { default = "10.0.0.0/24" }
variable "vcn_display_name" { default = "displayName" }
variable "vcn_dns_label" { default = "dnslabel2" }
variable "vcn_state" { default = "state" }
				` + compartmentIdVariableStr2 + VcnResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter DnsLabel but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
