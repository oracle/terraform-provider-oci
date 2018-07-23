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
variable "console_history_state" { default = "SUCCEEDED" }

`
	ConsoleHistoryResourceDependencies = InstancePropertyVariables + InstanceResourceConfig
)

func TestCoreConsoleHistoryResource_basic(t *testing.T) {
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
		CheckDestroy: testAccCheckCoreConsoleHistoryDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + ConsoleHistoryPropertyVariables + compartmentIdVariableStr + ConsoleHistoryRequiredOnlyResource,
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
variable "console_history_state" { default = "SUCCEEDED" }

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
variable "console_history_state" { default = "SUCCEEDED" }

data "oci_core_console_histories" "test_console_histories" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${oci_core_instance.test_instance.availability_domain}"
	instance_id = "${oci_core_instance.test_instance.id}"
	state = "${var.console_history_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_console_history.test_console_history.id}"]
    }
}
                ` + compartmentIdVariableStr + ConsoleHistoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),

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
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreConsoleHistoryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_console_history" {
			noResourceFound = false
			request := oci_core.GetConsoleHistoryRequest{}

			tmp := rs.Primary.ID
			request.InstanceConsoleHistoryId = &tmp

			_, err := client.GetConsoleHistory(context.Background(), request)

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
