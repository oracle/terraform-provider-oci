// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v41/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v41/databasemanagement"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ManagedDatabaseGroupRequiredOnlyResource = ManagedDatabaseGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Required, Create, managedDatabaseGroupRepresentation)

	ManagedDatabaseGroupResourceConfig = ManagedDatabaseGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Optional, Update, managedDatabaseGroupRepresentation)

	managedDatabaseGroupSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_group_id": Representation{repType: Required, create: `${oci_database_management_managed_database_group.test_managed_database_group.id}`},
	}

	managedDatabaseGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"id":             Representation{repType: Optional, create: `${oci_database_management_managed_database_group.test_managed_database_group.id}`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, managedDatabaseGroupDataSourceFilterRepresentation}}

	managedDatabaseGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_management_managed_database_group.test_managed_database_group.id}`}},
	}

	managedDatabaseGroupRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"name":           Representation{repType: Required, create: `TestGroup`},
		"description":    Representation{repType: Optional, create: `Sales test database group`, update: `description2`},
	}

	managedDatabaseId0Representation = map[string]interface{}{
		"id": Representation{repType: Required, create: `${var.tenancy_ocid}testManagedDatabase0`},
	}

	managedDatabaseId1Representation = map[string]interface{}{
		"id": Representation{repType: Required, create: `${var.tenancy_ocid}testManagedDatabase1`},
	}

	managedDatabaseId2Representation = map[string]interface{}{
		"id": Representation{repType: Required, create: `${var.tenancy_ocid}testManagedDatabase2`},
	}

	managedDatabaseId3Representation = map[string]interface{}{
		"id": Representation{repType: Required, create: `${var.tenancy_ocid}testManagedDatabase3`},
	}

	managedDatabaseId4Representation = map[string]interface{}{
		"id": Representation{repType: Required, create: `${var.tenancy_ocid}testManagedDatabase4`},
	}

	managedDatabaseGroupRepresentationWithManagedDatabases = map[string]interface{}{
		"compartment_id":    Representation{repType: Required, create: `${var.compartment_id}`},
		"name":              Representation{repType: Required, create: `TestGroup`},
		"description":       Representation{repType: Optional, create: `Sales test database group`, update: `description2`},
		"managed_databases": []RepresentationGroup{{Optional, managedDatabaseId0Representation}, {Optional, managedDatabaseId1Representation}},
	}

	ManagedDatabaseGroupResourceDependencies = ""
)

func TestDatabaseManagementManagedDatabaseGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseGroupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_management_managed_database_group.test_managed_database_group"
	datasourceName := "data.oci_database_management_managed_database_groups.test_managed_database_groups"
	singularDatasourceName := "data.oci_database_management_managed_database_group.test_managed_database_group"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ManagedDatabaseGroupResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Optional, Create, managedDatabaseGroupRepresentation), "databasemanagement", "managedDatabaseGroup", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseManagementManagedDatabaseGroupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ManagedDatabaseGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Required, Create, managedDatabaseGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ManagedDatabaseGroupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ManagedDatabaseGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Optional, Create, managedDatabaseGroupRepresentationWithManagedDatabases),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "Sales test database group"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify update with updated managed_databases list
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ManagedDatabaseGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Optional, Create,
						representationCopyWithNewProperties(managedDatabaseGroupRepresentationWithManagedDatabases, map[string]interface{}{
							"managed_databases": []RepresentationGroup{{Optional, managedDatabaseId2Representation}, {Optional, managedDatabaseId3Representation}},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "Sales test database group"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify update after removing entry from managed_databases
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ManagedDatabaseGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Optional, Create,
						representationCopyWithNewProperties(managedDatabaseGroupRepresentationWithManagedDatabases, map[string]interface{}{
							"managed_databases": []RepresentationGroup{{Optional, managedDatabaseId2Representation}, {Optional, managedDatabaseId3Representation}},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "Sales test database group"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify update after adding entry to managed_databases
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ManagedDatabaseGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Optional, Create,
						representationCopyWithNewProperties(managedDatabaseGroupRepresentationWithManagedDatabases, map[string]interface{}{
							"managed_databases": []RepresentationGroup{{Optional, managedDatabaseId2Representation}, {Optional, managedDatabaseId4Representation}},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "Sales test database group"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ManagedDatabaseGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Optional, Create,
						representationCopyWithNewProperties(managedDatabaseGroupRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "description", "Sales test database group"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ManagedDatabaseGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Optional, Update, managedDatabaseGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "managed_databases.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "name", "TestGroup"),
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
					generateDataSourceFromRepresentationMap("oci_database_management_managed_database_groups", "test_managed_database_groups", Optional, Update, managedDatabaseGroupDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedDatabaseGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Optional, Update, managedDatabaseGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_database_management_managed_database_group", "test_managed_database_group", Required, Create, managedDatabaseGroupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedDatabaseGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_group_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "managed_databases.#", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "TestGroup"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ManagedDatabaseGroupResourceConfig,
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

func testAccCheckDatabaseManagementManagedDatabaseGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_managed_database_group" {
			noResourceFound = false
			request := oci_database_management.GetManagedDatabaseGroupRequest{}

			tmp := rs.Primary.ID
			request.ManagedDatabaseGroupId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database_management")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseManagementManagedDatabaseGroup") {
		resource.AddTestSweepers("DatabaseManagementManagedDatabaseGroup", &resource.Sweeper{
			Name:         "DatabaseManagementManagedDatabaseGroup",
			Dependencies: DependencyGraph["managedDatabaseGroup"],
			F:            sweepDatabaseManagementManagedDatabaseGroupResource,
		})
	}
}

func sweepDatabaseManagementManagedDatabaseGroupResource(compartment string) error {
	dbManagementClient := GetTestClients(&schema.ResourceData{}).dbManagementClient()
	managedDatabaseGroupIds, err := getManagedDatabaseGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, managedDatabaseGroupId := range managedDatabaseGroupIds {
		if ok := SweeperDefaultResourceId[managedDatabaseGroupId]; !ok {
			deleteManagedDatabaseGroupRequest := oci_database_management.DeleteManagedDatabaseGroupRequest{}

			deleteManagedDatabaseGroupRequest.ManagedDatabaseGroupId = &managedDatabaseGroupId

			deleteManagedDatabaseGroupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteManagedDatabaseGroup(context.Background(), deleteManagedDatabaseGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagedDatabaseGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", managedDatabaseGroupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &managedDatabaseGroupId, managedDatabaseGroupSweepWaitCondition, time.Duration(3*time.Minute),
				managedDatabaseGroupSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getManagedDatabaseGroupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ManagedDatabaseGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := GetTestClients(&schema.ResourceData{}).dbManagementClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "ManagedDatabaseGroupId", id)
	}
	return resourceIds, nil
}

func managedDatabaseGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managedDatabaseGroupResponse, ok := response.Response.(oci_database_management.GetManagedDatabaseGroupResponse); ok {
		return managedDatabaseGroupResponse.LifecycleState != oci_database_management.LifecycleStatesDeleted
	}
	return false
}

func managedDatabaseGroupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dbManagementClient().GetManagedDatabaseGroup(context.Background(), oci_database_management.GetManagedDatabaseGroupRequest{
		ManagedDatabaseGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
