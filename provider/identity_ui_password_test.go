// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	UiPasswordResourceConfig = UiPasswordResourceDependencies + `
resource "oci_identity_ui_password" "test_ui_password" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"
}
`
	UiPasswordPropertyVariables = `

`
	UiPasswordResourceDependencies = UserPropertyVariables + UserResourceConfig
)

func TestIdentityUiPasswordResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_ui_password.test_ui_password"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + UiPasswordPropertyVariables + compartmentIdVariableStr + UiPasswordResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),
				),
			},
		},
	})
}
