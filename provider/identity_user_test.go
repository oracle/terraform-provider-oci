// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + UserPropertyVariables + compartmentIdVariableStr + UserRequiredOnlyResource,
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
		},
	})
}
