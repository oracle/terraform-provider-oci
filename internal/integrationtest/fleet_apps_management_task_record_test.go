// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementTaskRecordRequiredOnlyResource = FleetAppsManagementTaskRecordResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_task_record", "test_task_record", acctest.Required, acctest.Create, FleetAppsManagementTaskRecordRepresentation)

	FleetAppsManagementTaskRecordResourceConfig = FleetAppsManagementTaskRecordResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_task_record", "test_task_record", acctest.Optional, acctest.Update, FleetAppsManagementTaskRecordRepresentation)

	FleetAppsManagementTaskRecordSingularDataSourceRepresentation = map[string]interface{}{
		"task_record_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_task_record.test_task_record.id}`},
	}

	FleetAppsManagementTaskRecordDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `USER_DEFINED`},
	}

	FleetAppsManagementTaskRecordRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"details":        acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementTaskRecordDetailsRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
	}
	FleetAppsManagementTaskRecordDetailsRepresentation = map[string]interface{}{
		"execution_details":        acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementTaskRecordDetailsExecutionDetailsRepresentation},
		"os_type":                  acctest.Representation{RepType: acctest.Required, Create: `LINUX`, Update: `WINDOWS`},
		"scope":                    acctest.Representation{RepType: acctest.Required, Create: `LOCAL`, Update: `SHARED`},
		"is_apply_subject_task":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_discovery_output_task": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"properties":               acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementTaskRecordDetailsPropertiesRepresentation},
	}
	FleetAppsManagementTaskRecordDetailsExecutionDetailsRepresentation = map[string]interface{}{
		"execution_type": acctest.Representation{RepType: acctest.Required, Create: `SCRIPT`},
		"command":        acctest.Representation{RepType: acctest.Optional, Create: `command`, Update: `command2`},
		"content":        acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementTaskRecordDetailsExecutionDetailsContentRepresentation},
		"variables":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementTaskRecordDetailsExecutionDetailsVariablesRepresentation},
	}
	FleetAppsManagementTaskRecordDetailsPropertiesRepresentation = map[string]interface{}{
		"num_retries":        acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"timeout_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	FleetAppsManagementTaskRecordDetailsExecutionDetailsContentRepresentation = map[string]interface{}{
		"bucket":      acctest.Representation{RepType: acctest.Required, Create: `bucket`, Update: `bucket2`},
		"checksum":    acctest.Representation{RepType: acctest.Required, Create: `checksum`, Update: `checksum2`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `namespace`, Update: `namespace2`},
		"object":      acctest.Representation{RepType: acctest.Required, Create: `object`, Update: `object2`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_BUCKET`},
	}
	FleetAppsManagementTaskRecordDetailsExecutionDetailsCredentialsRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":           acctest.Representation{RepType: acctest.Optional, Create: `id`, Update: `id2`},
	}
	FleetAppsManagementTaskRecordDetailsExecutionDetailsVariablesRepresentation = map[string]interface{}{
		"input_variables":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementTaskRecordDetailsExecutionDetailsVariablesInputVariablesRepresentation},
		"output_variables": acctest.Representation{RepType: acctest.Optional, Create: []string{`outputVariables`}, Update: []string{`outputVariables2`}},
	}
	FleetAppsManagementTaskRecordDetailsExecutionDetailsVariablesInputVariablesRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"type":        acctest.Representation{RepType: acctest.Required, Create: `STRING`, Update: `OUTPUT_VARIABLE`},
	}

	FleetAppsManagementTaskRecordResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementTaskRecordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementTaskRecordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fleet_apps_management_task_record.test_task_record"

	datasourceName := "data.oci_fleet_apps_management_task_records.test_task_records"
	singularDatasourceName := "data.oci_fleet_apps_management_task_record.test_task_record"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementTaskRecordResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_task_record", "test_task_record", acctest.Optional, acctest.Create, FleetAppsManagementTaskRecordRepresentation), "fleetappsmanagement", "taskRecord", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementTaskRecordDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementTaskRecordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_task_record", "test_task_record", acctest.Required, acctest.Create, FleetAppsManagementTaskRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(resourceName, "details.0.os_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "details.0.scope", "LOCAL"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementTaskRecordResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementTaskRecordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_task_record", "test_task_record", acctest.Optional, acctest.Create, FleetAppsManagementTaskRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.command", "command"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.0.bucket", "bucket"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.0.checksum", "checksum"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.0.object", "object"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.0.source_type", "OBJECT_STORAGE_BUCKET"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.0.input_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.0.input_variables.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.0.input_variables.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.0.input_variables.0.type", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.0.output_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.is_apply_subject_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "details.0.is_discovery_output_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "details.0.os_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.properties.0.num_retries", "10"),
				resource.TestCheckResourceAttr(resourceName, "details.0.properties.0.timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "details.0.scope", "LOCAL"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementTaskRecordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_task_record", "test_task_record", acctest.Optional, acctest.Update, FleetAppsManagementTaskRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.command", "command2"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.0.checksum", "checksum2"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.0.object", "object2"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.content.0.source_type", "OBJECT_STORAGE_BUCKET"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.0.input_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.0.input_variables.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.0.input_variables.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.0.input_variables.0.type", "OUTPUT_VARIABLE"),
				resource.TestCheckResourceAttr(resourceName, "details.0.execution_details.0.variables.0.output_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.is_apply_subject_task", "true"),
				resource.TestCheckResourceAttr(resourceName, "details.0.is_discovery_output_task", "true"),
				resource.TestCheckResourceAttr(resourceName, "details.0.os_type", "WINDOWS"),
				resource.TestCheckResourceAttr(resourceName, "details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.properties.0.num_retries", "11"),
				resource.TestCheckResourceAttr(resourceName, "details.0.properties.0.timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "details.0.scope", "SHARED"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_task_records", "test_task_records", acctest.Optional, acctest.Update, FleetAppsManagementTaskRecordDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementTaskRecordResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_task_record", "test_task_record", acctest.Optional, acctest.Update, FleetAppsManagementTaskRecordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "task_record_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "task_record_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_task_record", "test_task_record", acctest.Required, acctest.Create, FleetAppsManagementTaskRecordSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementTaskRecordResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "task_record_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.command", "command2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.content.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.content.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.content.0.checksum", "checksum2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.content.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.content.0.object", "object2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.content.0.source_type", "OBJECT_STORAGE_BUCKET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.variables.0.input_variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.variables.0.input_variables.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.variables.0.input_variables.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.variables.0.input_variables.0.type", "OUTPUT_VARIABLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.execution_details.0.variables.0.output_variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.is_apply_subject_task", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.is_discovery_output_task", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.os_type", "WINDOWS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.properties.0.num_retries", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.properties.0.timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.scope", "SHARED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementTaskRecordRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementTaskRecordDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementRunbooksClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_task_record" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetTaskRecordRequest{}

			tmp := rs.Primary.ID
			request.TaskRecordId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetTaskRecord(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.TaskRecordLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementTaskRecord") {
		resource.AddTestSweepers("FleetAppsManagementTaskRecord", &resource.Sweeper{
			Name:         "FleetAppsManagementTaskRecord",
			Dependencies: acctest.DependencyGraph["taskRecord"],
			F:            sweepFleetAppsManagementTaskRecordResource,
		})
	}
}

func sweepFleetAppsManagementTaskRecordResource(compartment string) error {
	fleetAppsManagementRunbooksClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementRunbooksClient()
	taskRecordIds, err := getFleetAppsManagementTaskRecordIds(compartment)
	if err != nil {
		return err
	}
	for _, taskRecordId := range taskRecordIds {
		if ok := acctest.SweeperDefaultResourceId[taskRecordId]; !ok {
			deleteTaskRecordRequest := oci_fleet_apps_management.DeleteTaskRecordRequest{}

			deleteTaskRecordRequest.TaskRecordId = &taskRecordId

			deleteTaskRecordRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementRunbooksClient.DeleteTaskRecord(context.Background(), deleteTaskRecordRequest)
			if error != nil {
				fmt.Printf("Error deleting TaskRecord %s %s, It is possible that the resource is already deleted. Please verify manually \n", taskRecordId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &taskRecordId, FleetAppsManagementTaskRecordSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementTaskRecordSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementTaskRecordIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TaskRecordId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementRunbooksClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementRunbooksClient()

	listTaskRecordsRequest := oci_fleet_apps_management.ListTaskRecordsRequest{}
	listTaskRecordsRequest.CompartmentId = &compartmentId
	listTaskRecordsRequest.LifecycleState = oci_fleet_apps_management.TaskRecordLifecycleStateActive
	listTaskRecordsResponse, err := fleetAppsManagementRunbooksClient.ListTaskRecords(context.Background(), listTaskRecordsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting TaskRecord list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, taskRecord := range listTaskRecordsResponse.Items {
		id := *taskRecord.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TaskRecordId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementTaskRecordSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if taskRecordResponse, ok := response.Response.(oci_fleet_apps_management.GetTaskRecordResponse); ok {
		return taskRecordResponse.LifecycleState != oci_fleet_apps_management.TaskRecordLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementTaskRecordSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementRunbooksClient().GetTaskRecord(context.Background(), oci_fleet_apps_management.GetTaskRecordRequest{
		TaskRecordId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
