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
	UserGroupMembershipResourceConfig = UserGroupMembershipResourceDependencies + `
resource "oci_identity_user_group_membership" "test_user_group_membership" {
	#Required
	group_id = "${oci_identity_group.test_group.id}"
	user_id = "${oci_identity_user.test_user.id}"
}
`
	UserGroupMembershipPropertyVariables = `

`
	UserGroupMembershipResourceDependencies = GroupPropertyVariables + GroupRequiredOnlyResource + UserPropertyVariables + UserRequiredOnlyResource
)

func TestIdentityUserGroupMembershipResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getRequiredEnvSetting("tenancy_ocid")

	resourceName := "oci_identity_user_group_membership.test_user_group_membership"
	datasourceName := "data.oci_identity_user_group_memberships.test_user_group_memberships"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityUserGroupMembershipDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + UserGroupMembershipPropertyVariables + compartmentIdVariableStr + UserGroupMembershipResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),
				),
			},

			// verify datasource
			{
				Config: config + `

data "oci_identity_user_group_memberships" "test_user_group_memberships" {
	#Required
	compartment_id = "${var.tenancy_ocid}"

	#Optional
	group_id = "${oci_identity_group.test_group.id}"
	user_id = "${oci_identity_user.test_user.id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_user_group_membership.test_user_group_membership.id}"]
    }
}
                ` + compartmentIdVariableStr + UserGroupMembershipResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttrSet(datasourceName, "group_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "memberships.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.group_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "memberships.0.user_id"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ResourceName:      resourceName,
			},
		},
	})
}

func testAccCheckIdentityUserGroupMembershipDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_user_group_membership" {
			noResourceFound = false
			request := oci_identity.GetUserGroupMembershipRequest{}

			tmp := rs.Primary.ID
			request.UserGroupMembershipId = &tmp

			_, err := client.GetUserGroupMembership(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
