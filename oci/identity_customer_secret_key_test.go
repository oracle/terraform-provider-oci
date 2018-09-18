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
	CustomerSecretKeyResourceConfig = CustomerSecretKeyResourceDependencies + `
resource "oci_identity_customer_secret_key" "test_customer_secret_key" {
	#Required
	display_name = "${var.customer_secret_key_display_name}"
	user_id = "${oci_identity_user.test_user.id}"
}
`
	CustomerSecretKeyPropertyVariables = `
variable "customer_secret_key_display_name" { default = "displayName" }

`
	CustomerSecretKeyResourceDependencies = UserPropertyVariables + UserResourceConfig
)

func TestIdentityCustomerSecretKeyResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_customer_secret_key.test_customer_secret_key"
	datasourceName := "data.oci_identity_customer_secret_keys.test_customer_secret_keys"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityCustomerSecretKeyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + CustomerSecretKeyPropertyVariables + compartmentIdVariableStr + CustomerSecretKeyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "key"),
					//resource.TestCheckResourceAttrSet(resourceName, "inactive_state"), // not set by service
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					//resource.TestCheckResourceAttrSet(resourceName, "time_expires"), // not set by service

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "customer_secret_key_display_name" { default = "displayName2" }

                ` + compartmentIdVariableStr + CustomerSecretKeyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
variable "customer_secret_key_display_name" { default = "displayName2" }

data "oci_identity_customer_secret_keys" "test_customer_secret_keys" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_customer_secret_key.test_customer_secret_key.id}"]
    }
}
                ` + compartmentIdVariableStr + CustomerSecretKeyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "customer_secret_keys.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "customer_secret_keys.0.display_name", "displayName2"),
					TestCheckResourceAttributesEqual(datasourceName, "customer_secret_keys.0.user_id", "oci_identity_customer_secret_key.test_customer_secret_key", "user_id"),
				),
			},
		},
	})
}

func testAccCheckIdentityCustomerSecretKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_customer_secret_key" {
			noResourceFound = false
			request := oci_identity.ListCustomerSecretKeysRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			response, err := client.ListCustomerSecretKeys(context.Background(), request)

			if err == nil {
				id := rs.Primary.Attributes["id"]
				for _, item := range response.Items {
					if *item.Id == id {
						return fmt.Errorf("item still exists")
					}
				}
				// no error and item not found, that means item is deleted. continue checking next one
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
