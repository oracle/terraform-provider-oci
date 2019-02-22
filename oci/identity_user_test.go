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
	UserRequiredOnlyResource = UserResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)

	userDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"filter":         RepresentationGroup{Required, userDataSourceFilterRepresentation}}
	userDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_user.test_user.id}`}},
	}

	userRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"description":    Representation{repType: Required, create: `John Smith`, update: `description2`},
		"name":           Representation{repType: Required, create: `JohnSmith@example.com`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"email":          Representation{repType: Optional, create: `email`, update: `email2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	UserResourceDependencies = DefinedTagsDependencies
)

func TestIdentityUserResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_user.test_user"
	datasourceName := "data.oci_identity_users.test_users"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityUserDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + UserResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation),
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
				Config: config + compartmentIdVariableStr + UserResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_user", "test_user", Optional, Create, userRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "John Smith"),
					resource.TestCheckResourceAttr(resourceName, "email", "email"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "JohnSmith@example.com"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "capabilities.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + UserResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_user", "test_user", Optional, Update, userRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "email", "email2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "JohnSmith@example.com"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "capabilities.#", "1"),

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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_users", "test_users", Optional, Update, userDataSourceRepresentation) +
					compartmentIdVariableStr + UserResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_user", "test_user", Optional, Update, userRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttr(datasourceName, "users.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "users.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "users.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "users.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "users.0.email", "email2"),
					resource.TestCheckResourceAttr(datasourceName, "users.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "users.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "users.0.name", "JohnSmith@example.com"),
					resource.TestCheckResourceAttrSet(datasourceName, "users.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "users.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "users.0.capabilities.#", "1"),
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")

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

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	resource.AddTestSweepers("IdentityUser", &resource.Sweeper{
		Name:         "IdentityUser",
		Dependencies: DependencyGraph["user"],
		F:            sweepIdentityUserResource,
	})
}

func sweepIdentityUserResource(compartment string) error {
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient
	userIds, err := getUserIds(compartment)
	if err != nil {
		return err
	}
	for _, userId := range userIds {
		if ok := SweeperDefaultResourceId[userId]; !ok {
			deleteUserRequest := oci_identity.DeleteUserRequest{}

			deleteUserRequest.UserId = &userId

			deleteUserRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			_, error := identityClient.DeleteUser(context.Background(), deleteUserRequest)
			if error != nil {
				fmt.Printf("Error deleting User %s %s, It is possible that the resource is already deleted. Please verify manually \n", userId, error)
				continue
			}
			waitTillCondition(testAccProvider, &userId, userSweepWaitCondition, time.Duration(3*time.Minute),
				userSweepResponseFetchOperation, "identity", true)
		}
	}
	return nil
}

func getUserIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "UserId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient

	listUsersRequest := oci_identity.ListUsersRequest{}
	listUsersRequest.CompartmentId = &compartmentId
	listUsersResponse, err := identityClient.ListUsers(context.Background(), listUsersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting User list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, user := range listUsersResponse.Items {
		id := *user.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "UserId", id)
	}
	return resourceIds, nil
}

func userSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if userResponse, ok := response.Response.(oci_identity.GetUserResponse); ok {
		return userResponse.LifecycleState == oci_identity.UserLifecycleStateDeleted
	}
	return false
}

func userSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.identityClient.GetUser(context.Background(), oci_identity.GetUserRequest{
		UserId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
