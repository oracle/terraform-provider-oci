// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
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
		CheckDestroy: testAccCheckIdentitySmtpCredentialDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + SmtpCredentialPropertyVariables + compartmentIdVariableStr + SmtpCredentialResourceConfig,
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

func testAccCheckIdentitySmtpCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_smtp_credential" {
			noResourceFound = false
			request := oci_identity.ListSmtpCredentialsRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			response, err := client.ListSmtpCredentials(context.Background(), request)

			if err == nil {
				id := rs.Primary.Attributes["id"]
				for _, item := range response.Items {
					if *item.Id == id {
						return fmt.Errorf("item still exists")
					}
				}
				// no error and item not found, item is deleted
				return nil
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
