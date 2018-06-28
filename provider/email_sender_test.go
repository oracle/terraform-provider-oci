// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_email_sender.test_sender"
	datasourceName := "data.oci_email_senders.test_senders"
	singularDatasourceName := "data.oci_email_sender.test_sender"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + SenderPropertyVariables + compartmentIdVariableStr + SenderResourceConfig,
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
		},
	})
}
