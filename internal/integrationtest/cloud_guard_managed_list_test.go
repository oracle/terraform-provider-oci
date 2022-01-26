// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v56/cloudguard"
	"github.com/oracle/oci-go-sdk/v56/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ManagedListRequiredOnlyResource = ManagedListResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_managed_list", "test_managed_list", acctest.Required, acctest.Create, managedListRepresentation)

	ManagedListResourceConfig = ManagedListResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_managed_list", "test_managed_list", acctest.Optional, acctest.Update, managedListRepresentation)

	managedListSingularDataSourceRepresentation = map[string]interface{}{
		"managed_list_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_managed_list.test_managed_list.id}`},
	}

	managedListDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		//access_level has acceptable values as RESTRICTED and ACCESSIBLE, latter providing lenient access check.
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		//Valid list Type is required
		"list_type":              acctest.Representation{RepType: acctest.Optional, Create: `USERS`},
		"resource_metadata_only": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		//Valid lifecyclestate is required
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: managedListDataSourceFilterRepresentation}}
	managedListDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_guard_managed_list.test_managed_list.id}`}},
	}

	managedListRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		//Below 2 params are marked as optional from api-spec but for testing purpose we will have that marked as required.
		//These 2 params are required for making CUSTOMER managed recipes but not for ORACLE managed recipes.
		"list_items":             acctest.Representation{RepType: acctest.Required, Create: []string{`listItems`}, Update: []string{`listItems2`}},
		"list_type":              acctest.Representation{RepType: acctest.Required, Create: `USERS`},
		"source_managed_list_id": acctest.Representation{RepType: acctest.Optional, Create: nil},
	}

	ManagedListResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardManagedListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardManagedListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_guard_managed_list.test_managed_list"
	datasourceName := "data.oci_cloud_guard_managed_lists.test_managed_lists"
	singularDatasourceName := "data.oci_cloud_guard_managed_list.test_managed_list"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ManagedListResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_managed_list", "test_managed_list", acctest.Optional, acctest.Create, managedListRepresentation), "cloudguard", "managedList", t)

	acctest.ResourceTest(t, testAccCheckCloudGuardManagedListDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ManagedListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_managed_list", "test_managed_list", acctest.Required, acctest.Create, managedListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ManagedListResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ManagedListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_managed_list", "test_managed_list", acctest.Optional, acctest.Create, managedListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "list_items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "list_type", "USERS"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ManagedListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_managed_list", "test_managed_list", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(managedListRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "list_items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "list_type", "USERS"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ManagedListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_managed_list", "test_managed_list", acctest.Optional, acctest.Update, managedListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "list_items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "list_type", "USERS"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_managed_lists", "test_managed_lists", acctest.Optional, acctest.Update, managedListDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_managed_list", "test_managed_list", acctest.Optional, acctest.Update, managedListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "list_type", "USERS"),
				resource.TestCheckResourceAttr(datasourceName, "resource_metadata_only", "false"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "managed_list_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "managed_list_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_managed_list", "test_managed_list", acctest.Required, acctest.Create, managedListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_list_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "feed_provider"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_editable"),
				//No life cycle details associated with ManagedList Resource
				resource.TestCheckResourceAttr(singularDatasourceName, "list_items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "list_type", "USERS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ManagedListResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudGuardManagedListDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_managed_list" {
			noResourceFound = false
			request := oci_cloud_guard.GetManagedListRequest{}

			tmp := rs.Primary.ID
			request.ManagedListId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")

			response, err := client.GetManagedList(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_guard.LifecycleStateDeleted): true,
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CloudGuardManagedList") {
		resource.AddTestSweepers("CloudGuardManagedList", &resource.Sweeper{
			Name:         "CloudGuardManagedList",
			Dependencies: acctest.DependencyGraph["managedList"],
			F:            sweepCloudGuardManagedListResource,
		})
	}
}

func sweepCloudGuardManagedListResource(compartment string) error {
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()
	managedListIds, err := getManagedListIds(compartment)
	if err != nil {
		return err
	}
	for _, managedListId := range managedListIds {
		if ok := acctest.SweeperDefaultResourceId[managedListId]; !ok {
			deleteManagedListRequest := oci_cloud_guard.DeleteManagedListRequest{}

			deleteManagedListRequest.ManagedListId = &managedListId

			deleteManagedListRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteManagedList(context.Background(), deleteManagedListRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagedList %s %s, It is possible that the resource is already deleted. Please verify manually \n", managedListId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &managedListId, managedListSweepWaitCondition, time.Duration(3*time.Minute),
				managedListSweepResponseFetchOperation, "cloud_guard", true)
		}
	}
	return nil
}

func getManagedListIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagedListId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()

	listManagedListsRequest := oci_cloud_guard.ListManagedListsRequest{}
	listManagedListsRequest.CompartmentId = &compartmentId
	listManagedListsRequest.LifecycleState = oci_cloud_guard.ListManagedListsLifecycleStateActive
	listManagedListsResponse, err := cloudGuardClient.ListManagedLists(context.Background(), listManagedListsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagedList list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managedList := range listManagedListsResponse.Items {
		id := *managedList.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagedListId", id)
	}
	return resourceIds, nil
}

func managedListSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managedListResponse, ok := response.Response.(oci_cloud_guard.GetManagedListResponse); ok {
		return managedListResponse.LifecycleState != oci_cloud_guard.LifecycleStateDeleted
	}
	return false
}

func managedListSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CloudGuardClient().GetManagedList(context.Background(), oci_cloud_guard.GetManagedListRequest{
		ManagedListId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
