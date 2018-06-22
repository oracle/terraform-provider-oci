// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ConsoleHistoryRequiredOnlyResource = ConsoleHistoryResourceDependencies + `
resource "oci_core_console_history" "test_console_history" {
	#Required
	instance_id = "${oci_core_instance.test_instance.id}"
}
`

	ConsoleHistoryResourceConfig = ConsoleHistoryResourceDependencies + `
resource "oci_core_console_history" "test_console_history" {
	#Required
	instance_id = "${oci_core_instance.test_instance.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.console_history_defined_tags_value}")}"
	display_name = "${var.console_history_display_name}"
	freeform_tags = "${var.console_history_freeform_tags}"
}
`
	ConsoleHistoryPropertyVariables = `
variable "console_history_availability_domain" { default = "availabilityDomain" }
variable "console_history_defined_tags_value" { default = "value" }
variable "console_history_display_name" { default = "displayName" }
variable "console_history_freeform_tags" { default = {"Department"= "Finance"} }
variable "console_history_state" { default = "AVAILABLE" }

`
	ConsoleHistoryResourceDependencies = DefinedTagsDependencies // Uncomment once defined: InstancePropertyVariables + InstanceResourceConfig
)

func TestCoreConsoleHistoryResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_console_history.test_console_history"
	datasourceName := "data.oci_core_console_histories.test_console_histories"

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
				Config:            config + ConsoleHistoryPropertyVariables + compartmentIdVariableStr + ConsoleHistoryRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ConsoleHistoryResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + ConsoleHistoryPropertyVariables + compartmentIdVariableStr + ConsoleHistoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "console_history_availability_domain" { default = "availabilityDomain" }
variable "console_history_defined_tags_value" { default = "updatedValue" }
variable "console_history_display_name" { default = "displayName2" }
variable "console_history_freeform_tags" { default = {"Department"= "Accounting"} }
variable "console_history_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + ConsoleHistoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
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
variable "console_history_availability_domain" { default = "availabilityDomain" }
variable "console_history_defined_tags_value" { default = "updatedValue" }
variable "console_history_display_name" { default = "displayName2" }
variable "console_history_freeform_tags" { default = {"Department"= "Accounting"} }
variable "console_history_state" { default = "AVAILABLE" }

data "oci_core_console_histories" "test_console_histories" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.console_history_availability_domain}"
	instance_id = "${oci_core_instance.test_instance.id}"
	state = "${var.console_history_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_console_history.test_console_history.id}"]
    }
}
                ` + compartmentIdVariableStr + ConsoleHistoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "console_histories.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "console_histories.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "console_histories.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "console_histories.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.instance_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.time_created"),
				),
			},
		},
	})
}
