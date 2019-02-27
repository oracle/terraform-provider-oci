// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
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

	UserGroupMembershipResourceDependencies = GroupResourceConfig + UserRequiredOnlyResource
)

func TestIdentityUserGroupMembershipResource_basic(t *testing.T) {
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
	client := testAccProvider.Meta().(*OracleClients).identityClient
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

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	resource.AddTestSweepers("IdentityUserGroupMembership", &resource.Sweeper{
		Name:         "IdentityUserGroupMembership",
		Dependencies: DependencyGraph["userGroupMembership"],
		F:            sweepIdentityUserGroupMembershipResource,
	})
}

func sweepIdentityUserGroupMembershipResource(compartment string) error {
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient
	userGroupMembershipIds, err := getUserGroupMembershipIds(compartment)
	if err != nil {
		return err
	}
	for _, userGroupMembershipId := range userGroupMembershipIds {
		if ok := SweeperDefaultResourceId[userGroupMembershipId]; !ok {
			removeUserFromGroupRequest := oci_identity.RemoveUserFromGroupRequest{}

			removeUserFromGroupRequest.UserGroupMembershipId = &userGroupMembershipId

			removeUserFromGroupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			_, error := identityClient.RemoveUserFromGroup(context.Background(), removeUserFromGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting UserGroupMembership %s %s, It is possible that the resource is already deleted. Please verify manually \n", userGroupMembershipId, error)
				continue
			}
			waitTillCondition(testAccProvider, &userGroupMembershipId, userGroupMembershipSweepWaitCondition, time.Duration(3*time.Minute),
				userGroupMembershipSweepResponseFetchOperation, "identity", true)
		}
	}
	return nil
}

func getUserGroupMembershipIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "UserGroupMembershipId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient

	listUserGroupMembershipsRequest := oci_identity.ListUserGroupMembershipsRequest{}
	listUserGroupMembershipsRequest.CompartmentId = &compartmentId
	listUserGroupMembershipsResponse, err := identityClient.ListUserGroupMemberships(context.Background(), listUserGroupMembershipsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting UserGroupMembership list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, userGroupMembership := range listUserGroupMembershipsResponse.Items {
		id := *userGroupMembership.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "UserGroupMembershipId", id)
	}
	return resourceIds, nil
}

func userGroupMembershipSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if userGroupMembershipResponse, ok := response.Response.(oci_identity.GetUserGroupMembershipResponse); ok {
		return userGroupMembershipResponse.LifecycleState == oci_identity.UserGroupMembershipLifecycleStateDeleted
	}
	return false
}

func userGroupMembershipSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.identityClient.GetUserGroupMembership(context.Background(), oci_identity.GetUserGroupMembershipRequest{
		UserGroupMembershipId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
