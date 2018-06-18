// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	SwiftPasswordResourceConfig = SwiftPasswordResourceDependencies + `
resource "oci_identity_swift_password" "test_swift_password" {
	#Required
	description = "${var.swift_password_description}"
	user_id = "${oci_identity_user.test_user.id}"
}
`
	SwiftPasswordPropertyVariables = `
variable "swift_password_description" { default = "description" }

`
	SwiftPasswordResourceDependencies = UserPropertyVariables + UserResourceConfig
)

func TestIdentitySwiftPasswordResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_swift_password.test_swift_password"
	datasourceName := "data.oci_identity_swift_passwords.test_swift_passwords"

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
				Config:            config + SwiftPasswordPropertyVariables + compartmentIdVariableStr + SwiftPasswordResourceConfig,
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
variable "swift_password_description" { default = "description2" }

                ` + compartmentIdVariableStr + SwiftPasswordResourceConfig,
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
variable "swift_password_description" { default = "description2" }

data "oci_identity_swift_passwords" "test_swift_passwords" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_swift_password.test_swift_password.id}"]
    }
}
                ` + compartmentIdVariableStr + SwiftPasswordResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "passwords.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "passwords.0.description", "description2"),
					resource.TestCheckResourceAttrSet(datasourceName, "passwords.0.user_id"),
				),
			},
		},
	})
}
