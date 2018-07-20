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
	AuthTokenResourceConfig = AuthTokenResourceDependencies + `
resource "oci_identity_auth_token" "test_auth_token" {
	#Required
	description = "${var.auth_token_description}"
	user_id = "${oci_identity_user.test_user.id}"
}
`
	AuthTokenPropertyVariables = `
variable "auth_token_description" { default = "description" }

`
	AuthTokenResourceDependencies = UserPropertyVariables + UserResourceConfig
)

func TestIdentityAuthTokenResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_auth_token.test_auth_token"
	datasourceName := "data.oci_identity_auth_tokens.test_auth_tokens"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityAuthTokenDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + AuthTokenPropertyVariables + compartmentIdVariableStr + AuthTokenResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "auth_token_description" { default = "description2" }

                ` + compartmentIdVariableStr + AuthTokenResourceConfig,
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
variable "auth_token_description" { default = "description2" }

data "oci_identity_auth_tokens" "test_auth_tokens" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_auth_token.test_auth_token.id}"]
    }
}
                ` + compartmentIdVariableStr + AuthTokenResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "tokens.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "tokens.0.description", "description2"),
					resource.TestCheckResourceAttrSet(datasourceName, "tokens.0.user_id"),
				),
			},
		},
	})
}

func testAccCheckIdentityAuthTokenDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_auth_token" {
			noResourceFound = false
			request := oci_identity.ListAuthTokensRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			response, err := client.ListAuthTokens(context.Background(), request)

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
