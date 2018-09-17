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
	SenderResourceConfig = SenderResourceDependencies + `
resource "oci_email_sender" "test_sender" {
	#Required
	compartment_id = "${var.compartment_id}"
	email_address = "${var.sender_email_address}"
}
`
	SenderPropertyVariables = `
variable "sender_email_address" { default = "JohnSmith@example.com" }
variable "sender_state" { default = "ACTIVE" }

`
	SenderResourceDependencies = ""
)

func TestEmailSenderResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_email_sender.test_sender"
	datasourceName := "data.oci_email_senders.test_senders"
	singularDatasourceName := "data.oci_email_sender.test_sender"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckEmailSenderDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + SenderPropertyVariables + compartmentIdVariableStr + SenderResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "email_address", "JohnSmith@example.com"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "sender_email_address" { default = "JohnSmith@example.com" }
variable "sender_state" { default = "ACTIVE" }

data "oci_email_senders" "test_senders" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	email_address = "${var.sender_email_address}"
	state = "${var.sender_state}"

    filter {
    	name = "id"
    	values = ["${oci_email_sender.test_sender.id}"]
    }
}
                ` + compartmentIdVariableStr + SenderResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "email_address", "JohnSmith@example.com"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "senders.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "senders.0.email_address", "JohnSmith@example.com"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
variable "sender_email_address" { default = "JohnSmith@example.com" }
variable "sender_state" { default = "ACTIVE" }

data "oci_email_sender" "test_sender" {
	#Required
	sender_id = "${oci_email_sender.test_sender.id}"
}
                ` + compartmentIdVariableStr + SenderResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "sender_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "email_address", "JohnSmith@example.com"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_spf", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + `
variable "sender_email_address" { default = "JohnSmith@example.com" }
variable "sender_state" { default = "ACTIVE" }

                ` + compartmentIdVariableStr + SenderResourceConfig,
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

func testAccCheckEmailSenderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).emailClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_sender" {
			noResourceFound = false
			request := oci_email.GetSenderRequest{}

			tmp := rs.Primary.ID
			request.SenderId = &tmp

			response, err := client.GetSender(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_email.SenderLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
