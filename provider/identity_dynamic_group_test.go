// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"regexp"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	DynamicGroupResourceConfig = DynamicGroupResourceDependencies + `
resource "oci_identity_dynamic_group" "test_dynamic_group" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.dynamic_group_description}"
	matching_rule = "${var.dynamic_group_matching_rule}"
	name = "${var.dynamic_group_name}"
}
`
	DynamicGroupPropertyVariables = `
variable "dynamic_group_description" { default = "Instance group for dev compartment" }
variable "dynamic_group_name" { default = "DevCompartmentDynamicGroup" }

`
	DynamicGroupResourceDependencies = ""
)

func TestIdentityDynamicGroupResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getRequiredEnvSetting("tenancy_ocid")

	matchingRuleValueStr := fmt.Sprintf("instance.compartment_id='%s'", compartmentId)
	matchingRuleVariableStr := fmt.Sprintf("variable \"dynamic_group_matching_rule\" {default = \"%s\" }\n", matchingRuleValueStr)

	matchingRule2ValueStr := fmt.Sprintf("instance.compartment_id='%s'", compartmentId)
	matchingRule2VariableStr := fmt.Sprintf("variable \"dynamic_group_matching_rule\" {default = \"%s\" }\n", matchingRule2ValueStr)
	resourceName := "oci_identity_dynamic_group.test_dynamic_group"
	datasourceName := "data.oci_identity_dynamic_groups.test_dynamic_groups"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify matching rule syntax
			{
				Config: config + `
variable "dynamic_group_description" { default = "description2" }
variable "dynamic_group_matching_rule" { default = "bad_matching_rule" }
variable "dynamic_group_name" { default = "DevCompartmentDynamicGroup" }` + compartmentIdVariableStr + DynamicGroupResourceConfig,
				ExpectError: regexp.MustCompile("Unable to parse matching rule"),
			},
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + DynamicGroupPropertyVariables + compartmentIdVariableStr + matchingRuleVariableStr + DynamicGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "Instance group for dev compartment"),
					resource.TestCheckResourceAttr(resourceName, "matching_rule", matchingRuleValueStr),
					resource.TestCheckResourceAttr(resourceName, "name", "DevCompartmentDynamicGroup"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "dynamic_group_description" { default = "description2" }
variable "dynamic_group_name" { default = "DevCompartmentDynamicGroup" }

                ` + compartmentIdVariableStr + matchingRule2VariableStr + DynamicGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "matching_rule", matchingRule2ValueStr),
					resource.TestCheckResourceAttr(resourceName, "name", "DevCompartmentDynamicGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
variable "dynamic_group_description" { default = "description2" }
variable "dynamic_group_name" { default = "DevCompartmentDynamicGroup" }

data "oci_identity_dynamic_groups" "test_dynamic_groups" {
	#Required
	compartment_id = "${var.tenancy_ocid}"

    filter {
    	name = "id"
    	values = ["${oci_identity_dynamic_group.test_dynamic_group.id}"]
    }
}
                ` + compartmentIdVariableStr + matchingRule2VariableStr + DynamicGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.description", "description2"),
					resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.matching_rule", matchingRule2ValueStr),
					resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.name", "DevCompartmentDynamicGroup"),
					resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.time_created"),
				),
			},
		},
	})
}
