// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ignoreManagedDatabaseGroupDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DatabaseManagementManagedDatabaseGroupRequiredOnlyResource = DatabaseManagementManagedDatabaseGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Required, acctest.Create, DatabaseManagementManagedDatabaseGroupRepresentation)

	DatabaseManagementManagedDatabaseGroupResourceConfig = DatabaseManagementManagedDatabaseGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Optional, acctest.Update, DatabaseManagementManagedDatabaseGroupRepresentation)

	DatabaseManagementDatabaseManagementManagedDatabaseGroupSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database_group.test_managed_database_group.id}`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_managed_database_group.test_managed_database_group.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementManagedDatabaseGroupDataSourceFilterRepresentation}}

	DatabaseManagementManagedDatabaseGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_managed_database_group.test_managed_database_group.id}`}},
	}

	DatabaseManagementManagedDatabaseGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `TestGroup`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Sales test database Group`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreManagedDatabaseGroupDefinedTagsChangesRepresentation},
	}

	managedDatabaseId0Representation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase0`},
	}

	managedDatabaseId1Representation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase1`},
	}

	managedDatabaseId2Representation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase2`},
	}

	managedDatabaseId3Representation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase3`},
	}

	managedDatabaseId4Representation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase4`},
	}

	managedDatabaseGroupRepresentationWithManagedDatabases = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":              acctest.Representation{RepType: acctest.Required, Create: `TestGroup`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `Sales test database Group`, Update: `description2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"managed_databases": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: managedDatabaseId0Representation}, {RepType: acctest.Optional, Group: managedDatabaseId1Representation}},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreManagedDatabaseGroupDefinedTagsChangesRepresentation},
	}

	DatabaseManagementManagedDatabaseGroupResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_management_managed_database_group.test_managed_database_group"
	datasourceName := "data.oci_database_management_managed_database_groups.test_managed_database_groups"
	singularDatasourceName := "data.oci_database_management_managed_database_group.test_managed_database_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseManagementManagedDatabaseGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Optional, acctest.Create, DatabaseManagementManagedDatabaseGroupRepresentation), "databasemanagement", "managedDatabaseGroup", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementManagedDatabaseGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseManagementManagedDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Required, acctest.Create, DatabaseManagementManagedDatabaseGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseManagementManagedDatabaseGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseManagementManagedDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Optional, acctest.Create, managedDatabaseGroupRepresentationWithManagedDatabases),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Sales test database Group"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Update with updated managed_databases list
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseManagementManagedDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(managedDatabaseGroupRepresentationWithManagedDatabases, map[string]interface{}{
						"managed_databases": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: managedDatabaseId2Representation}, {RepType: acctest.Optional, Group: managedDatabaseId3Representation}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Sales test database Group"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Update after removing entry from managed_databases
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseManagementManagedDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(managedDatabaseGroupRepresentationWithManagedDatabases, map[string]interface{}{
						"managed_databases": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: managedDatabaseId2Representation}, {RepType: acctest.Optional, Group: managedDatabaseId3Representation}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Sales test database Group"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Update after adding entry to managed_databases
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseManagementManagedDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(managedDatabaseGroupRepresentationWithManagedDatabases, map[string]interface{}{
						"managed_databases": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: managedDatabaseId2Representation}, {RepType: acctest.Optional, Group: managedDatabaseId4Representation}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Sales test database Group"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseManagementManagedDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseManagementManagedDatabaseGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Sales test database Group"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + DatabaseManagementManagedDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Optional, acctest.Update, DatabaseManagementManagedDatabaseGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_groups", "test_managed_database_groups", acctest.Optional, acctest.Update, DatabaseManagementDatabaseManagementManagedDatabaseGroupDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedDatabaseGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Optional, acctest.Update, DatabaseManagementManagedDatabaseGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "managed_database_group_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "managed_database_group_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedDatabaseGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "managed_databases.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TestGroup"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseManagementManagedDatabaseGroupRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseManagementManagedDatabaseGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_managed_database_group" {
			noResourceFound = false
			request := oci_database_management.GetManagedDatabaseGroupRequest{}

			tmp := rs.Primary.ID
			request.ManagedDatabaseGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetManagedDatabaseGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementManagedDatabaseGroup") {
		resource.AddTestSweepers("DatabaseManagementManagedDatabaseGroup", &resource.Sweeper{
			Name:         "DatabaseManagementManagedDatabaseGroup",
			Dependencies: acctest.DependencyGraph["managedDatabaseGroup"],
			F:            sweepDatabaseManagementManagedDatabaseGroupResource,
		})
	}
}

func sweepDatabaseManagementManagedDatabaseGroupResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	managedDatabaseGroupIds, err := getDatabaseManagementManagedDatabaseGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, managedDatabaseGroupId := range managedDatabaseGroupIds {
		if ok := acctest.SweeperDefaultResourceId[managedDatabaseGroupId]; !ok {
			deleteManagedDatabaseGroupRequest := oci_database_management.DeleteManagedDatabaseGroupRequest{}

			deleteManagedDatabaseGroupRequest.ManagedDatabaseGroupId = &managedDatabaseGroupId

			deleteManagedDatabaseGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteManagedDatabaseGroup(context.Background(), deleteManagedDatabaseGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagedDatabaseGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", managedDatabaseGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &managedDatabaseGroupId, DatabaseManagementManagedDatabaseGroupSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementManagedDatabaseGroupSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementManagedDatabaseGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagedDatabaseGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listManagedDatabaseGroupsRequest := oci_database_management.ListManagedDatabaseGroupsRequest{}
	listManagedDatabaseGroupsRequest.CompartmentId = &compartmentId
	listManagedDatabaseGroupsRequest.LifecycleState = oci_database_management.ListManagedDatabaseGroupsLifecycleStateActive
	listManagedDatabaseGroupsResponse, err := dbManagementClient.ListManagedDatabaseGroups(context.Background(), listManagedDatabaseGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagedDatabaseGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managedDatabaseGroup := range listManagedDatabaseGroupsResponse.Items {
		id := *managedDatabaseGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagedDatabaseGroupId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementManagedDatabaseGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managedDatabaseGroupResponse, ok := response.Response.(oci_database_management.GetManagedDatabaseGroupResponse); ok {
		return managedDatabaseGroupResponse.LifecycleState != oci_database_management.LifecycleStatesDeleted
	}
	return false
}

func DatabaseManagementManagedDatabaseGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetManagedDatabaseGroup(context.Background(), oci_database_management.GetManagedDatabaseGroupRequest{
		ManagedDatabaseGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
