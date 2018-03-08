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
	display_name = "${var.console_history_display_name}"
}
`
	ConsoleHistoryPropertyVariables = `
variable "console_history_availability_domain" { default = "availabilityDomain" }
variable "console_history_display_name" { default = "displayName" }
variable "console_history_state" { default = "state" }

`
	ConsoleHistoryResourceDependencies = "" // Uncomment once defined: InstancePropertyVariables + InstanceResourceConfig
)

func TestCoreConsoleHistoryResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

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
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
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
variable "console_history_display_name" { default = "displayName2" }
variable "console_history_state" { default = "state" }

                ` + compartmentIdVariableStr + ConsoleHistoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
			// verify updates to Force New parameters.
			{
				Config: config + `
variable "console_history_availability_domain" { default = "availabilityDomain2" }
variable "console_history_display_name" { default = "displayName2" }
variable "console_history_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr2 + ConsoleHistoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
variable "console_history_availability_domain" { default = "availabilityDomain2" }
variable "console_history_display_name" { default = "displayName2" }
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
                ` + compartmentIdVariableStr2 + ConsoleHistoryResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "console_histories.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "console_histories.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.instance_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "console_histories.0.time_created"),
				),
			},
		},
	})
}

func TestCoreConsoleHistoryResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_console_history.test_console_history"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + ConsoleHistoryPropertyVariables + compartmentIdVariableStr + ConsoleHistoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
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
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "console_history_availability_domain" { default = "availabilityDomain" }
variable "console_history_display_name" { default = "displayName" }
variable "console_history_state" { default = "state" }
				` + compartmentIdVariableStr + ConsoleHistoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter InstanceId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
