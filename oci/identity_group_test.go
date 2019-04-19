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

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	GroupRequiredOnlyResource = GroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create, groupRepresentation)

	groupSingularDataSourceRepresentation = map[string]interface{}{
		"group_id": Representation{repType: Required, create: `${oci_identity_group.test_group.id}`},
	}

	groupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"filter":         RepresentationGroup{Required, groupDataSourceFilterRepresentation}}
	groupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_group.test_group.id}`}},
	}

	groupRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"description":    Representation{repType: Required, create: `Group for network administrators`, update: `description2`},
		"name":           Representation{repType: Required, create: `NetworkAdmins`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	GroupResourceDependencies = DefinedTagsDependencies
	GroupResourceConfig       = generateResourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create, groupRepresentation)
)

func TestIdentityGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityGroupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_group.test_group"
	datasourceName := "data.oci_identity_groups.test_groups"
	singularDatasourceName := "data.oci_identity_group.test_group"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityGroupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + GroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create, groupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "description", "Group for network administrators"),
					resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + GroupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + GroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Create, groupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "Group for network administrators"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),
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
				Config: config + compartmentIdVariableStr + GroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Update, groupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "NetworkAdmins"),
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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_groups", "test_groups", Optional, Update, groupDataSourceRepresentation) +
					compartmentIdVariableStr + GroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Update, groupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttr(datasourceName, "groups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "groups.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "groups.0.name", "NetworkAdmins"),
					resource.TestCheckResourceAttrSet(datasourceName, "groups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "groups.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create, groupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + GroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Update, groupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "group_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "NetworkAdmins"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + GroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_group", "test_group", Optional, Update, groupRepresentation),
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

func testAccCheckIdentityGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_group" {
			noResourceFound = false
			request := oci_identity.GetGroupRequest{}

			tmp := rs.Primary.ID
			request.GroupId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")

			response, err := client.GetGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.GroupLifecycleStateDeleted): true,
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
	resource.AddTestSweepers("IdentityGroup", &resource.Sweeper{
		Name:         "IdentityGroup",
		Dependencies: DependencyGraph["group"],
		F:            sweepIdentityGroupResource,
	})
}

func sweepIdentityGroupResource(compartment string) error {
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient
	groupIds, err := getGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, groupId := range groupIds {
		if ok := SweeperDefaultResourceId[groupId]; !ok {
			deleteGroupRequest := oci_identity.DeleteGroupRequest{}

			deleteGroupRequest.GroupId = &groupId

			deleteGroupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			_, error := identityClient.DeleteGroup(context.Background(), deleteGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting Group %s %s, It is possible that the resource is already deleted. Please verify manually \n", groupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &groupId, groupSweepWaitCondition, time.Duration(3*time.Minute),
				groupSweepResponseFetchOperation, "identity", true)
		}
	}
	return nil
}

func getGroupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "GroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityClient := GetTestClients(&schema.ResourceData{}).identityClient

	listGroupsRequest := oci_identity.ListGroupsRequest{}
	listGroupsRequest.CompartmentId = &compartmentId
	listGroupsResponse, err := identityClient.ListGroups(context.Background(), listGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Group list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, group := range listGroupsResponse.Items {
		id := *group.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "GroupId", id)
	}
	return resourceIds, nil
}

func groupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if groupResponse, ok := response.Response.(oci_identity.GetGroupResponse); ok {
		return groupResponse.LifecycleState != oci_identity.GroupLifecycleStateDeleted
	}
	return false
}

func groupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.identityClient.GetGroup(context.Background(), oci_identity.GetGroupRequest{
		GroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
