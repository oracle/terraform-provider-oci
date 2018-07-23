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
	UserRequiredOnlyResource = UserRequiredOnlyResourceDependencies + `
resource "oci_identity_user" "test_user" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.user_description}"
	name = "${var.user_name}"
}
`

	UserResourceConfig = UserResourceDependencies + `
resource "oci_identity_user" "test_user" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.user_description}"
	name = "${var.user_name}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.user_defined_tags_value}")}"
	freeform_tags = "${var.user_freeform_tags}"
}
`
	UserPropertyVariables = `
variable "user_defined_tags_value" { default = "value" }
variable "user_description" { default = "John Smith" }
variable "user_freeform_tags" { default = {"Department"= "Finance"} }
variable "user_name" { default = "JohnSmith@example.com" }

`
	UserRequiredOnlyResourceDependencies = ``
	UserResourceDependencies             = DefinedTagsDependencies
)

func TestIdentityUserResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getRequiredEnvSetting("tenancy_ocid")

	resourceName := "oci_identity_user.test_user"
	datasourceName := "data.oci_identity_users.test_users"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityUserDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + UserPropertyVariables + compartmentIdVariableStr + UserRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "John Smith"),
					resource.TestCheckResourceAttr(resourceName, "name", "JohnSmith@example.com"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + UserResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + UserPropertyVariables + compartmentIdVariableStr + UserResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "John Smith"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "JohnSmith@example.com"),
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
variable "user_defined_tags_value" { default = "updatedValue" }
variable "user_description" { default = "description2" }
variable "user_freeform_tags" { default = {"Department"= "Accounting"} }
variable "user_name" { default = "JohnSmith@example.com" }

                ` + compartmentIdVariableStr + UserResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "JohnSmith@example.com"),
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
variable "user_defined_tags_value" { default = "updatedValue" }
variable "user_description" { default = "description2" }
variable "user_freeform_tags" { default = {"Department"= "Accounting"} }
variable "user_name" { default = "JohnSmith@example.com" }

data "oci_identity_users" "test_users" {
	#Required
	compartment_id = "${var.tenancy_ocid}"

    filter {
    	name = "id"
    	values = ["${oci_identity_user.test_user.id}"]
    }
}
                ` + compartmentIdVariableStr + UserResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttr(datasourceName, "users.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "users.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "users.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "users.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "users.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "users.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "users.0.name", "JohnSmith@example.com"),
					resource.TestCheckResourceAttrSet(datasourceName, "users.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "users.0.time_created"),
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

func testAccCheckIdentityUserDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_user" {
			noResourceFound = false
			request := oci_identity.GetUserRequest{}

			tmp := rs.Primary.ID
			request.UserId = &tmp

			response, err := client.GetUser(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.UserLifecycleStateDeleted): true,
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
