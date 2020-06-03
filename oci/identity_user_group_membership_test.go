// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	userGroupMembershipDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"group_id":       Representation{repType: Optional, create: `${oci_identity_group.test_group.id}`},
		"user_id":        Representation{repType: Optional, create: `${oci_identity_user.test_user.id}`},
		"filter":         RepresentationGroup{Required, userGroupMembershipDataSourceFilterRepresentation}}
	userGroupMembershipDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_user_group_membership.test_user_group_membership.id}`}},
	}

	userGroupMembershipRepresentation = map[string]interface{}{
		"group_id": Representation{repType: Required, create: `${oci_identity_group.test_group.id}`},
		"user_id":  Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
	}

	UserGroupMembershipResourceDependencies = generateResourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create, groupRepresentation) +
		generateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)
)

func TestIdentityUserGroupMembershipResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityUserGroupMembershipResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_user_group_membership.test_user_group_membership"
	datasourceName := "data.oci_identity_user_group_memberships.test_user_group_memberships"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityUserGroupMembershipDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + UserGroupMembershipResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_user_group_membership", "test_user_group_membership", Required, Create, userGroupMembershipRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_user_group_memberships", "test_user_group_memberships", Optional, Update, userGroupMembershipDataSourceRepresentation) +
					compartmentIdVariableStr + UserGroupMembershipResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_user_group_membership", "test_user_group_membership", Optional, Update, userGroupMembershipRepresentation),
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
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckIdentityUserGroupMembershipDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_user_group_membership" {
			noResourceFound = false
			request := oci_identity.GetUserGroupMembershipRequest{}

			tmp := rs.Primary.ID
			request.UserGroupMembershipId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")

			response, err := client.GetUserGroupMembership(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.UserGroupMembershipLifecycleStateDeleted): true,
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
