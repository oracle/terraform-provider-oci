// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_email "github.com/oracle/oci-go-sdk/email"
)

const (
	SuppressionResourceConfig = SuppressionResourceDependencies + `
resource "oci_email_suppression" "test_suppression" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	email_address = "${var.suppression_email_address}"
}
`
	SuppressionPropertyVariables = `
variable "suppression_email_address" { default = "JohnSmith@example.com" }
variable "suppression_time_created_greater_than_or_equal_to" { default = "2018-01-01T00:00:00.000Z" }
variable "suppression_time_created_less_than" { default = "2038-01-01T00:00:00.000Z" }

`
	SuppressionResourceDependencies = ""
)

func TestEmailSuppressionResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getRequiredEnvSetting("tenancy_ocid")

	resourceName := "oci_email_suppression.test_suppression"
	datasourceName := "data.oci_email_suppressions.test_suppressions"
	singularDatasourceName := "data.oci_email_suppression.test_suppression"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckEmailSuppressionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + SuppressionPropertyVariables + compartmentIdVariableStr + SuppressionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					// email address is converted to lower case by the service
					resource.TestCheckResourceAttr(resourceName, "email_address", "johnsmith@example.com"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "suppression_email_address" { default = "JohnSmith@example.com" }
variable "suppression_time_created_greater_than_or_equal_to" { default = "2018-01-01T00:00:00.000Z" }
variable "suppression_time_created_less_than" { default = "2038-01-01T00:00:00.000Z" }

data "oci_email_suppressions" "test_suppressions" {
	#Required
	compartment_id = "${var.tenancy_ocid}"

	#Optional
	email_address = "${var.suppression_email_address}"
	time_created_greater_than_or_equal_to = "${var.suppression_time_created_greater_than_or_equal_to}"
	time_created_less_than = "${var.suppression_time_created_less_than}"

    filter {
    	name = "id"
    	values = ["${oci_email_suppression.test_suppression.id}"]
    }
}
                ` + compartmentIdVariableStr + SuppressionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "email_address", "JohnSmith@example.com"),
					resource.TestCheckResourceAttr(datasourceName, "time_created_greater_than_or_equal_to", "2018-01-01T00:00:00.000Z"),
					resource.TestCheckResourceAttr(datasourceName, "time_created_less_than", "2038-01-01T00:00:00.000Z"),

					resource.TestCheckResourceAttr(datasourceName, "suppressions.#", "1"),
					// email address is converted to lower case by the service
					resource.TestCheckResourceAttr(datasourceName, "suppressions.0.email_address", "johnsmith@example.com"),
					resource.TestCheckResourceAttrSet(datasourceName, "suppressions.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "suppressions.0.reason", "MANUAL"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
variable "suppression_email_address" { default = "JohnSmith@example.com" }
variable "suppression_time_created_greater_than_or_equal_to" { default = "2018-01-01T00:00:00.000Z" }
variable "suppression_time_created_less_than" { default = "2038-01-01T00:00:00.000Z" }

data "oci_email_suppression" "test_suppression" {
	#Required
	suppression_id = "${oci_email_suppression.test_suppression.id}"
}
                ` + compartmentIdVariableStr + SuppressionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "suppression_id"),

					// email address is converted to lower case by the service
					resource.TestCheckResourceAttr(singularDatasourceName, "email_address", "johnsmith@example.com"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "reason", "MANUAL"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"compartment_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckEmailSuppressionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).emailClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_suppression" {
			noResourceFound = false
			request := oci_email.GetSuppressionRequest{}

			tmp := rs.Primary.ID
			request.SuppressionId = &tmp

			_, err := client.GetSuppression(context.Background(), request)

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
