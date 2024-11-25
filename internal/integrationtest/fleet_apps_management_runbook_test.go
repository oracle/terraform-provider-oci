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
	FleetAppsManagementRunbookRequiredOnlyResource = FleetAppsManagementRunbookResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Required, acctest.Create, FleetAppsManagementRunbookRepresentation)

	FleetAppsManagementRunbookResourceConfig = FleetAppsManagementRunbookResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Optional, acctest.Update, FleetAppsManagementRunbookRepresentation)

	FleetAppsManagementRunbookSingularDataSourceRepresentation = map[string]interface{}{
		"runbook_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_runbook.test_runbook.id}`},
	}

	FleetAppsManagementRunbookDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `TestRunbookTFP2`},
		"operation":         acctest.Representation{RepType: acctest.Optional, Create: `PATCH`},
		"platform":          acctest.Representation{RepType: acctest.Optional, Create: `test product`},
		"runbook_relevance": acctest.Representation{RepType: acctest.Optional, Create: `PRODUCT`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"type":              acctest.Representation{RepType: acctest.Optional, Create: `USER_DEFINED`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookDataSourceFilterRepresentation}}
	FleetAppsManagementRunbookDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_runbook.test_runbook.id}`}},
	}

	FleetAppsManagementRunbookRepresentation = map[string]interface{}{
		"associations":      acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsRepresentation},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"operation":         acctest.Representation{RepType: acctest.Required, Create: `PATCH`},
		"os_type":           acctest.Representation{RepType: acctest.Required, Create: `LINUX`, Update: `WINDOWS`},
		"runbook_relevance": acctest.Representation{RepType: acctest.Required, Create: `PRODUCT`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `TestRunbookTFP6`, Update: `TestRunbookTFP2`},
		"estimated_time":    acctest.Representation{RepType: acctest.Required, Create: `PT1H`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_default":        acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"platform":          acctest.Representation{RepType: acctest.Required, Create: `test product`},
	}
	FleetAppsManagementRunbookAssociationsRepresentation = map[string]interface{}{
		"execution_workflow_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsExecutionWorkflowDetailsRepresentation},
		"groups":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsGroupsRepresentation},
		"tasks":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsTasksRepresentation},
	}
	FleetAppsManagementRunbookAssociationsExecutionWorkflowDetailsRepresentation = map[string]interface{}{
		"workflow": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsExecutionWorkflowDetailsWorkflowRepresentation},
	}
	FleetAppsManagementRunbookAssociationsGroupsRepresentation = map[string]interface{}{
		"name":       acctest.Representation{RepType: acctest.Required, Create: `Parallel_resource_container`},
		"type":       acctest.Representation{RepType: acctest.Required, Create: `PARALLEL_RESOURCE_GROUP`},
		"properties": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsGroupsPropertiesRepresentation},
	}
	FleetAppsManagementRunbookAssociationsTasksRepresentation = map[string]interface{}{
		"association_type":    acctest.Representation{RepType: acctest.Required, Create: `TASK`},
		"step_name":           acctest.Representation{RepType: acctest.Required, Create: `Patch_Oracle_Exadata_Database_Service`, Update: `stepName2`},
		"task_record_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsRepresentation},
		"step_properties":     acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsTasksStepPropertiesRepresentation},
	}
	FleetAppsManagementRunbookAssociationsRollbackWorkflowDetailsRepresentation = map[string]interface{}{
		"scope":    acctest.Representation{RepType: acctest.Required, Create: `ACTION_GROUP`, Update: `TARGET`},
		"workflow": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsRollbackWorkflowDetailsWorkflowRepresentation},
	}
	FleetAppsManagementRunbookAssociationsExecutionWorkflowDetailsWorkflowRepresentation = map[string]interface{}{
		"group_name": acctest.Representation{RepType: acctest.Required, Create: `Parallel_resource_container`},
		"steps":      acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsExecutionWorkflowDetailsWorkflowStepsRepresentation},
		"type":       acctest.Representation{RepType: acctest.Required, Create: `PARALLEL_RESOURCE_GROUP`},
	}
	FleetAppsManagementRunbookAssociationsGroupsPropertiesRepresentation = map[string]interface{}{
		"action_on_failure":        acctest.Representation{RepType: acctest.Required, Create: `ABORT`, Update: `CONTINUE`},
		"condition":                acctest.Representation{RepType: acctest.Optional, Create: `condition`, Update: `condition2`},
		"notification_preferences": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookAssociationsGroupsPropertiesNotificationPreferencesRepresentation},
		"pause_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookAssociationsGroupsPropertiesPauseDetailsRepresentation},
		"run_on":                   acctest.Representation{RepType: acctest.Optional, Create: `runOn`, Update: `runOn2`},
	}
	FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsRepresentation = map[string]interface{}{
		"scope":                      acctest.Representation{RepType: acctest.Required, Create: `SHARED`},
		"is_apply_subject_task":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_copy_to_library_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_discovery_output_task":   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"task_record_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.test_task_record_ocid}`},
	}
	FleetAppsManagementRunbookAssociationsTasksOutputVariableMappingsRepresentation = map[string]interface{}{
		"name":                    acctest.Representation{RepType: acctest.Required, Create: `outputVariableName`, Update: `outputVariableName2`},
		"output_variable_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsTasksOutputVariableMappingsOutputVariableDetailsRepresentation},
	}
	FleetAppsManagementRunbookAssociationsTasksStepPropertiesRepresentation = map[string]interface{}{
		"action_on_failure":        acctest.Representation{RepType: acctest.Required, Create: `ABORT`, Update: `CONTINUE`},
		"condition":                acctest.Representation{RepType: acctest.Optional, Create: `condition`, Update: `condition2`},
		"notification_preferences": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookAssociationsTasksStepPropertiesNotificationPreferencesRepresentation},
		"pause_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookAssociationsTasksStepPropertiesPauseDetailsRepresentation},
		"run_on":                   acctest.Representation{RepType: acctest.Optional, Create: `runOn`, Update: `runOn2`},
	}
	FleetAppsManagementRunbookAssociationsRollbackWorkflowDetailsWorkflowRepresentation = map[string]interface{}{
		"group_name": acctest.Representation{RepType: acctest.Required, Create: `Parallel_resource_container`},
		"steps":      acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookAssociationsRollbackWorkflowDetailsWorkflowStepsRepresentation},
		"type":       acctest.Representation{RepType: acctest.Required, Create: `PARALLEL_RESOURCE_GROUP`, Update: `ROLLING_RESOURCE_GROUP`},
	}
	FleetAppsManagementRunbookAssociationsExecutionWorkflowDetailsWorkflowStepsRepresentation = map[string]interface{}{
		"type":      acctest.Representation{RepType: acctest.Required, Create: `TASK`},
		"step_name": acctest.Representation{RepType: acctest.Required, Create: `Patch_Oracle_Exadata_Database_Service`, Update: `stepName2`},
	}
	FleetAppsManagementRunbookAssociationsGroupsPropertiesNotificationPreferencesRepresentation = map[string]interface{}{
		"should_notify_on_pause":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_notify_on_task_failure": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_notify_on_task_success": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	FleetAppsManagementRunbookAssociationsGroupsPropertiesPauseDetailsRepresentation = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Required, Create: `USER_ACTION`, Update: `TIME_BASED`},
	}
	FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsExecutionDetailsRepresentation = map[string]interface{}{
		"execution_type": acctest.Representation{RepType: acctest.Required, Create: `SCRIPT`, Update: `API`},
		"command":        acctest.Representation{RepType: acctest.Optional, Create: `command`, Update: `command2`},
		"content":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsExecutionDetailsContentRepresentation},
		"credentials":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsExecutionDetailsCredentialsRepresentation},
		"endpoint":       acctest.Representation{RepType: acctest.Optional, Create: `endpoint`, Update: `endpoint2`},
		"variables":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsExecutionDetailsVariablesRepresentation},
	}
	FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsPropertiesRepresentation = map[string]interface{}{
		"num_retries":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"timeout_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	FleetAppsManagementRunbookAssociationsTasksOutputVariableMappingsOutputVariableDetailsRepresentation = map[string]interface{}{
		"output_variable_name": acctest.Representation{RepType: acctest.Required, Create: `outputVariableName`, Update: `outputVariableName2`},
		"step_name":            acctest.Representation{RepType: acctest.Required, Create: `Patch_Oracle_Exadata_Database_Service`, Update: `stepName2`},
	}
	FleetAppsManagementRunbookAssociationsTasksStepPropertiesNotificationPreferencesRepresentation = map[string]interface{}{
		"should_notify_on_pause":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_notify_on_task_failure": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_notify_on_task_success": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	FleetAppsManagementRunbookAssociationsTasksStepPropertiesPauseDetailsRepresentation = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Required, Create: `USER_ACTION`, Update: `TIME_BASED`},
	}
	FleetAppsManagementRunbookAssociationsRollbackWorkflowDetailsWorkflowStepsRepresentation = map[string]interface{}{
		"type":       acctest.Representation{RepType: acctest.Required, Create: `TASK`, Update: `PARALLEL_TASK_GROUP`},
		"group_name": acctest.Representation{RepType: acctest.Optional, Create: `Parallel_resource_container`},
		"step_name":  acctest.Representation{RepType: acctest.Optional, Create: `Rollback_Patch_Oracle_Exadata_Database_Service`, Update: `stepName2`},
		"steps":      acctest.Representation{RepType: acctest.Optional, Create: []string{`steps`}, Update: []string{`steps2`}},
	}
	FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsExecutionDetailsContentRepresentation = map[string]interface{}{
		"bucket":      acctest.Representation{RepType: acctest.Required, Create: `bucket`, Update: `bucket2`},
		"checksum":    acctest.Representation{RepType: acctest.Required, Create: `checksum`, Update: `checksum2`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `namespace`, Update: `namespace2`},
		"object":      acctest.Representation{RepType: acctest.Required, Create: `object`, Update: `object2`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_BUCKET`},
	}
	FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsExecutionDetailsCredentialsRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":           acctest.Representation{RepType: acctest.Optional, Create: `id`, Update: `id2`},
	}
	FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsExecutionDetailsVariablesRepresentation = map[string]interface{}{
		"input_variables":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsExecutionDetailsVariablesInputVariablesRepresentation},
		"output_variables": acctest.Representation{RepType: acctest.Optional, Create: []string{`outputVariables`}, Update: []string{`outputVariables2`}},
	}
	FleetAppsManagementRunbookAssociationsTasksTaskRecordDetailsExecutionDetailsVariablesInputVariablesRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"type":        acctest.Representation{RepType: acctest.Optional, Create: `STRING`, Update: `OUTPUT_VARIABLE`},
	}

	FleetAppsManagementRunbookResourceDependencies = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementRunbookResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementRunbookResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fleet_apps_management_runbook.test_runbook"
	taskRecordId := utils.GetEnvSettingWithBlankDefault("test_task_record_ocid")
	testTaskRecordStr := fmt.Sprintf("variable \"test_task_record_ocid\" { default = \"%s\" }\n", taskRecordId)

	datasourceName := "data.oci_fleet_apps_management_runbooks.test_runbooks"
	singularDatasourceName := "data.oci_fleet_apps_management_runbook.test_runbook"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementRunbookResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Optional, acctest.Create, FleetAppsManagementRunbookRepresentation), "fleetappsmanagement", "runbook", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementRunbookDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + testTaskRecordStr + compartmentIdVariableStr + FleetAppsManagementRunbookResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Required, acctest.Create, FleetAppsManagementRunbookRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "associations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "associations.0.execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.steps.0.step_name", "Patch_Oracle_Exadata_Database_Service"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.name", "Parallel_resource_container"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.association_type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_name", "Patch_Oracle_Exadata_Database_Service"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.0.scope", "SHARED"),
				resource.TestCheckResourceAttrSet(resourceName, "associations.0.tasks.0.task_record_details.0.task_record_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(resourceName, "os_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "runbook_relevance", "PRODUCT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementRunbookResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + testTaskRecordStr + compartmentIdVariableStr + FleetAppsManagementRunbookResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Optional, acctest.Create, FleetAppsManagementRunbookRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "associations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "associations.0.execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.steps.0.step_name", "Patch_Oracle_Exadata_Database_Service"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.name", "Parallel_resource_container"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.action_on_failure", "ABORT"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.condition", "condition"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.notification_preferences.0.should_notify_on_pause", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.notification_preferences.0.should_notify_on_task_failure", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.notification_preferences.0.should_notify_on_task_success", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.pause_details.0.kind", "USER_ACTION"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.run_on", "runOn"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.association_type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_name", "Patch_Oracle_Exadata_Database_Service"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.action_on_failure", "ABORT"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.condition", "condition"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.0.should_notify_on_pause", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_failure", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_success", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.pause_details.0.kind", "USER_ACTION"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.run_on", "runOn"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.0.is_apply_subject_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.0.is_copy_to_library_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.0.is_discovery_output_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.0.scope", "SHARED"),
				resource.TestCheckResourceAttrSet(resourceName, "associations.0.tasks.0.task_record_details.0.task_record_id"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.version", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestRunbookTFP6"),
				resource.TestCheckResourceAttr(resourceName, "estimated_time", "PT1H"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				resource.TestCheckResourceAttr(resourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(resourceName, "os_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "platform", "test product"),
				resource.TestCheckResourceAttr(resourceName, "runbook_relevance", "PRODUCT"),
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
			Config: config + testTaskRecordStr + compartmentIdVariableStr + FleetAppsManagementRunbookResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Optional, acctest.Update, FleetAppsManagementRunbookRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "associations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "associations.0.execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.name", "Parallel_resource_container"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.action_on_failure", "CONTINUE"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.condition", "condition2"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.notification_preferences.0.should_notify_on_pause", "true"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.notification_preferences.0.should_notify_on_task_failure", "true"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.notification_preferences.0.should_notify_on_task_success", "true"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.pause_details.0.kind", "TIME_BASED"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.properties.0.run_on", "runOn2"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.association_type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_name", "stepName2"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.action_on_failure", "CONTINUE"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.condition", "condition2"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.0.should_notify_on_pause", "true"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_failure", "true"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_success", "true"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.pause_details.0.kind", "TIME_BASED"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.step_properties.0.run_on", "runOn2"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.0.is_apply_subject_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.0.is_copy_to_library_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.0.is_discovery_output_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "associations.0.tasks.0.task_record_details.0.scope", "SHARED"),
				resource.TestCheckResourceAttrSet(resourceName, "associations.0.tasks.0.task_record_details.0.task_record_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestRunbookTFP2"),
				resource.TestCheckResourceAttr(resourceName, "estimated_time", "PT1H"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				resource.TestCheckResourceAttr(resourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(resourceName, "os_type", "WINDOWS"),
				resource.TestCheckResourceAttr(resourceName, "platform", "test product"),
				resource.TestCheckResourceAttr(resourceName, "runbook_relevance", "PRODUCT"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_runbooks", "test_runbooks", acctest.Optional, acctest.Update, FleetAppsManagementRunbookDataSourceRepresentation) +
				testTaskRecordStr + compartmentIdVariableStr + FleetAppsManagementRunbookResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Optional, acctest.Update, FleetAppsManagementRunbookRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TestRunbookTFP2"),
				resource.TestCheckResourceAttr(datasourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(datasourceName, "platform", "test product"),
				resource.TestCheckResourceAttr(datasourceName, "runbook_relevance", "PRODUCT"),
				resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type", "USER_DEFINED"),

				resource.TestCheckResourceAttr(datasourceName, "runbook_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "runbook_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Required, acctest.Create, FleetAppsManagementRunbookSingularDataSourceRepresentation) +
				testTaskRecordStr + compartmentIdVariableStr + FleetAppsManagementRunbookResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "associations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.execution_workflow_details.0.workflow.0.steps.0.step_name", "stepName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.name", "Parallel_resource_container"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.properties.0.action_on_failure", "CONTINUE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.properties.0.condition", "condition2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.properties.0.notification_preferences.0.should_notify_on_pause", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.properties.0.notification_preferences.0.should_notify_on_task_failure", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.properties.0.notification_preferences.0.should_notify_on_task_success", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.properties.0.pause_details.0.kind", "TIME_BASED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.properties.0.run_on", "runOn2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.association_type", "TASK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_name", "stepName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_properties.0.action_on_failure", "CONTINUE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_properties.0.condition", "condition2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.0.should_notify_on_pause", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_failure", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_success", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_properties.0.pause_details.0.kind", "TIME_BASED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.step_properties.0.run_on", "runOn2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.task_record_details.0.is_apply_subject_task", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.task_record_details.0.is_copy_to_library_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.task_record_details.0.is_discovery_output_task", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "associations.0.tasks.0.task_record_details.0.scope", "SHARED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TestRunbookTFP2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "estimated_time", "PT1H"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_default", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_type", "WINDOWS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "platform", "test product"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_relevance", "PRODUCT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementRunbookRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementRunbookDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementRunbooksClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_runbook" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetRunbookRequest{}

			tmp := rs.Primary.ID
			request.RunbookId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetRunbook(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.RunbookLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementRunbook") {
		resource.AddTestSweepers("FleetAppsManagementRunbook", &resource.Sweeper{
			Name:         "FleetAppsManagementRunbook",
			Dependencies: acctest.DependencyGraph["runbook"],
			F:            sweepFleetAppsManagementRunbookResource,
		})
	}
}

func sweepFleetAppsManagementRunbookResource(compartment string) error {
	fleetAppsManagementRunbooksClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementRunbooksClient()
	runbookIds, err := getFleetAppsManagementRunbookIds(compartment)
	if err != nil {
		return err
	}
	for _, runbookId := range runbookIds {
		if ok := acctest.SweeperDefaultResourceId[runbookId]; !ok {
			deleteRunbookRequest := oci_fleet_apps_management.DeleteRunbookRequest{}

			deleteRunbookRequest.RunbookId = &runbookId

			deleteRunbookRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementRunbooksClient.DeleteRunbook(context.Background(), deleteRunbookRequest)
			if error != nil {
				fmt.Printf("Error deleting Runbook %s %s, It is possible that the resource is already deleted. Please verify manually \n", runbookId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &runbookId, FleetAppsManagementRunbookSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementRunbookSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementRunbookIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RunbookId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementRunbooksClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementRunbooksClient()

	listRunbooksRequest := oci_fleet_apps_management.ListRunbooksRequest{}
	listRunbooksRequest.CompartmentId = &compartmentId
	listRunbooksRequest.LifecycleState = oci_fleet_apps_management.RunbookLifecycleStateActive
	listRunbooksResponse, err := fleetAppsManagementRunbooksClient.ListRunbooks(context.Background(), listRunbooksRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Runbook list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, runbook := range listRunbooksResponse.Items {
		id := *runbook.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RunbookId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementRunbookSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if runbookResponse, ok := response.Response.(oci_fleet_apps_management.GetRunbookResponse); ok {
		return runbookResponse.LifecycleState != oci_fleet_apps_management.RunbookLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementRunbookSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementRunbooksClient().GetRunbook(context.Background(), oci_fleet_apps_management.GetRunbookRequest{
		RunbookId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
