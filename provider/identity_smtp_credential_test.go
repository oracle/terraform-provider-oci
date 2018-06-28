// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	SmtpCredentialResourceConfig = SmtpCredentialResourceDependencies + `
resource "oci_identity_smtp_credential" "test_smtp_credential" {
	#Required
	description = "${var.smtp_credential_description}"
	user_id = "${oci_identity_user.test_user.id}"
}
`
	SmtpCredentialPropertyVariables = `
variable "smtp_credential_description" { default = "description" }

`
	SmtpCredentialResourceDependencies = UserPropertyVariables + UserResourceConfig
)

func TestIdentitySmtpCredentialResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_smtp_credential.test_smtp_credential"
	datasourceName := "data.oci_identity_smtp_credentials.test_smtp_credentials"

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
				Config:            config + SmtpCredentialPropertyVariables + compartmentIdVariableStr + SmtpCredentialResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "password"),
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
variable "smtp_credential_description" { default = "description2" }

                ` + compartmentIdVariableStr + SmtpCredentialResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

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
variable "smtp_credential_description" { default = "description2" }

data "oci_identity_smtp_credentials" "test_smtp_credentials" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_smtp_credential.test_smtp_credential.id}"]
    }
}
                ` + compartmentIdVariableStr + SmtpCredentialResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "smtp_credentials.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "smtp_credentials.0.description", "description2"),
					resource.TestCheckResourceAttrSet(datasourceName, "smtp_credentials.0.user_id"),
					TestCheckResourceAttributesEqual(datasourceName, "smtp_credentials.0.user_id", "oci_identity_smtp_credential.test_smtp_credential", "user_id"),
				),
			},
		},
	})
}
