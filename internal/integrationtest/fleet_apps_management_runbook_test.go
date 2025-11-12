// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		//"id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_runbook.test_runbook.id}`},
		"operation": acctest.Representation{RepType: acctest.Optional, Create: `PATCH`},
		"platform":  acctest.Representation{RepType: acctest.Optional, Create: `Oracle Java`, Update: `Oracle Linux`},
		//"runbook_relevance": acctest.Representation{RepType: acctest.Optional, Create: `PRODUCT`},
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"type":   acctest.Representation{RepType: acctest.Optional, Create: `USER_DEFINED`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookDataSourceFilterRepresentation},
	}
	FleetAppsManagementRunbookDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_runbook.test_runbook.id}`}},
	}

	FleetAppsManagementRunbookRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"operation":       acctest.Representation{RepType: acctest.Required, Create: `PATCH`},
		"runbook_version": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionRepresentation},
		// "defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("Oracle-Tags.CreatedBy", "value")}`, Update: `${map("Oracle-Tags.CreatedBy", "updatedValue")}`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"estimated_time":        acctest.Representation{RepType: acctest.Optional, Create: `PT1H`, Update: `PT30M`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_default":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_sudo_access_needed": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"os_type":               acctest.Representation{RepType: acctest.Optional, Create: `LINUX`, Update: `WINDOWS`},
		"platform":              acctest.Representation{RepType: acctest.Required, Create: `Oracle Linux`, Update: `Oracle Java`},
		// "export_trigger":        acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}
	FleetAppsManagementRunbookRunbookVersionRepresentation = map[string]interface{}{
		"execution_workflow_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionExecutionWorkflowDetailsRepresentation},
		"groups":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionGroupsRepresentation},
		"tasks":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionTasksRepresentation},
		"is_latest":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		// "rollback_workflow_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionRollbackWorkflowDetailsRepresentation},
		"version": acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	FleetAppsManagementRunbookRunbookVersionExecutionWorkflowDetailsRepresentation = map[string]interface{}{
		"workflow": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionExecutionWorkflowDetailsWorkflowRepresentation},
	}
	FleetAppsManagementRunbookRunbookVersionGroupsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `Parallel_resource_group`},
		//"name":       acctest.Representation{RepType: acctest.Required, Create: `name`},
		//"type":       acctest.Representation{RepType: acctest.Required, Create: `PARALLEL_RESOURCE_GROUP`},
		"type":       acctest.Representation{RepType: acctest.Required, Create: `PARALLEL_RESOURCE_GROUP`},
		"properties": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionGroupsPropertiesRepresentation},
	}

	FleetAppsManagementRunbookRunbookVersionTasksRepresentation = map[string]interface{}{
		//"step_name":                acctest.Representation{RepType: acctest.Required, Create: `Patch_Oracle_Exadata_Database_Service`},
		"step_name":           acctest.Representation{RepType: acctest.Required, Create: `StepName`},
		"task_record_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsRepresentation},
		// "output_variable_mappings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksOutputVariableMappingsRepresentation},
		// "step_properties":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesRepresentation},
	}

	FleetAppsManagementRunbookRunbookVersionRollbackWorkflowDetailsRepresentation = map[string]interface{}{
		// "scope":    acctest.Representation{RepType: acctest.Required, Create: `ACTION_GROUP`},
		"scope":    acctest.Representation{RepType: acctest.Required, Create: `TARGET`},
		"workflow": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionRollbackWorkflowDetailsWorkflowRepresentation},
	}
	FleetAppsManagementRunbookRunbookVersionExecutionWorkflowDetailsWorkflowRepresentation = map[string]interface{}{
		"group_name": acctest.Representation{RepType: acctest.Required, Create: `Parallel_resource_group`},
		//"group_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_group.test_group.name}`},
		"steps": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionExecutionWorkflowDetailsWorkflowStepsRepresentation},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `PARALLEL_RESOURCE_GROUP`},
	}
	FleetAppsManagementRunbookRunbookVersionGroupsPropertiesRepresentation = map[string]interface{}{
		"action_on_failure": acctest.Representation{RepType: acctest.Required, Create: `ABORT`},
		// "notification_preferences": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionGroupsPropertiesNotificationPreferencesRepresentation},
		// "pause_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionGroupsPropertiesPauseDetailsRepresentation},
		// "pre_condition":            acctest.Representation{RepType: acctest.Optional, Create: `preCondition`},
		// "run_on":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionGroupsPropertiesRunOnRepresentation},
	}
	FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsRepresentation = map[string]interface{}{
		//	"scope":                      acctest.Representation{RepType: acctest.Required, Create: `SHARED`},
		"scope":                      acctest.Representation{RepType: acctest.Required, Create: `LOCAL`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"execution_details":          acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsExecutionDetailsRepresentation},
		"is_apply_subject_task":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_copy_to_library_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_discovery_output_task":   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"name":                       acctest.Representation{RepType: acctest.Optional, Create: `StepName`},
		//"os_type":                    acctest.Representation{RepType: acctest.Optional, Create: `WINDOWS`},
		"os_type":    acctest.Representation{RepType: acctest.Optional, Create: `LINUX`},
		"platform":   acctest.Representation{RepType: acctest.Optional, Create: `Oracle Linux`},
		"properties": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsPropertiesRepresentation},
		//"task_record_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.test_task_record_ocid}`},
		// "task_record_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_task_record.test_task_record.id}`},
	}
	FleetAppsManagementRunbookRunbookVersionTasksOutputVariableMappingsRepresentation = map[string]interface{}{
		"name":                    acctest.Representation{RepType: acctest.Required, Create: `name`},
		"output_variable_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionTasksOutputVariableMappingsOutputVariableDetailsRepresentation},
	}
	FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesRepresentation = map[string]interface{}{
		"action_on_failure":        acctest.Representation{RepType: acctest.Required, Create: `ABORT`},
		"notification_preferences": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesNotificationPreferencesRepresentation},
		"pause_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesPauseDetailsRepresentation},
		"pre_condition":            acctest.Representation{RepType: acctest.Optional, Create: `preCondition`},
		"run_on":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesRunOnRepresentation},
	}
	FleetAppsManagementRunbookRunbookVersionRollbackWorkflowDetailsWorkflowRepresentation = map[string]interface{}{
		"group_name": acctest.Representation{RepType: acctest.Required, Create: `Rolling_resource_group`},
		//"group_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_group.test_group.name}`},
		"steps": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookRunbookVersionRollbackWorkflowDetailsWorkflowStepsRepresentation},
		"type":  acctest.Representation{RepType: acctest.Required, Create: `ROLLING_RESOURCE_GROUP`},
	}
	FleetAppsManagementRunbookRunbookVersionExecutionWorkflowDetailsWorkflowStepsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `TASK`},
		//"group_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		// "group_name": acctest.Representation{RepType: acctest.Optional, Create: `Parallel_resource_group`},
		// "step_name":  acctest.Representation{RepType: acctest.Required, Create: `Patch_Oracle_Exadata_Database_Service`},
		"step_name": acctest.Representation{RepType: acctest.Required, Create: `StepName`},
		// "steps":     acctest.Representation{RepType: acctest.Optional, Create: []string{`steps`}},
	}
	FleetAppsManagementRunbookRunbookVersionGroupsPropertiesNotificationPreferencesRepresentation = map[string]interface{}{
		"should_notify_on_pause":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"should_notify_on_task_failure": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"should_notify_on_task_success": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	FleetAppsManagementRunbookRunbookVersionGroupsPropertiesPauseDetailsRepresentation = map[string]interface{}{
		"kind":                acctest.Representation{RepType: acctest.Required, Create: `USER_ACTION`},
		"duration_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	FleetAppsManagementRunbookRunbookVersionGroupsPropertiesRunOnRepresentation = map[string]interface{}{
		"kind":                           acctest.Representation{RepType: acctest.Required, Create: `SCHEDULED_INSTANCES`},
		"condition":                      acctest.Representation{RepType: acctest.Optional, Create: `condition`},
		"host":                           acctest.Representation{RepType: acctest.Optional, Create: `host`},
		"previous_task_instance_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionGroupsPropertiesRunOnPreviousTaskInstanceDetailsRepresentation},
	}
	FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsExecutionDetailsRepresentation = map[string]interface{}{
		"execution_type": acctest.Representation{RepType: acctest.Required, Create: `SCRIPT`},
		// "catalog_id":                      acctest.Representation{RepType: acctest.Optional, Create: `catalogId`},
		"command": acctest.Representation{RepType: acctest.Required, Create: `pwd`},
		// "config_file":                     acctest.Representation{RepType: acctest.Optional, Create: `configFile`},
		// "content":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsExecutionDetailsContentRepresentation},
		// "credentials":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsExecutionDetailsCredentialsRepresentation},
		// "endpoint":                        acctest.Representation{RepType: acctest.Optional, Create: `endpoint`},
		"is_executable_content":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_locked":                       acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_read_output_variable_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		// "system_variables":                acctest.Representation{RepType: acctest.Optional, Create: []string{`$CONFIG_FILE$`}},
		// "target_compartment_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_compartment.test_compartment.id}`},
		// "target_compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"variables": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsExecutionDetailsVariablesRepresentation},
	}
	FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsPropertiesRepresentation = map[string]interface{}{
		"num_retries":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"timeout_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `1000`},
	}
	FleetAppsManagementRunbookRunbookVersionTasksOutputVariableMappingsOutputVariableDetailsRepresentation = map[string]interface{}{
		"output_variable_name": acctest.Representation{RepType: acctest.Required, Create: `outputVariableName`},
		"step_name":            acctest.Representation{RepType: acctest.Required, Create: `Patch_Oracle_Exadata_Database_Service`},
		//"step_name":            acctest.Representation{RepType: acctest.Required, Create: `stepName`},
	}
	FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesNotificationPreferencesRepresentation = map[string]interface{}{
		"should_notify_on_pause":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"should_notify_on_task_failure": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"should_notify_on_task_success": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesPauseDetailsRepresentation = map[string]interface{}{
		"kind":                acctest.Representation{RepType: acctest.Required, Create: `USER_ACTION`},
		"duration_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesRunOnRepresentation = map[string]interface{}{
		"kind":                           acctest.Representation{RepType: acctest.Required, Create: `SCHEDULED_INSTANCES`},
		"condition":                      acctest.Representation{RepType: acctest.Optional, Create: `condition`},
		"host":                           acctest.Representation{RepType: acctest.Optional, Create: `host`},
		"previous_task_instance_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesRunOnPreviousTaskInstanceDetailsRepresentation},
	}
	FleetAppsManagementRunbookRunbookVersionRollbackWorkflowDetailsWorkflowStepsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `TASK`},
		// "group_name": acctest.Representation{RepType: acctest.Optional, Create: `Parallel_resource_container`},
		"step_name": acctest.Representation{RepType: acctest.Optional, Create: `RollbackStepName`},
		// "steps":      acctest.Representation{RepType: acctest.Optional, Create: []string{`steps`}},
	}
	FleetAppsManagementRunbookRunbookVersionGroupsPropertiesRunOnPreviousTaskInstanceDetailsRepresentation = map[string]interface{}{
		"output_variable_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionGroupsPropertiesRunOnPreviousTaskInstanceDetailsOutputVariableDetailsRepresentation},
		"resource_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_resource.test_resource.id}`},
		"resource_type":           acctest.Representation{RepType: acctest.Optional, Create: `resourceType`},
	}
	FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsExecutionDetailsContentRepresentation = map[string]interface{}{
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_BUCKET`},
		"bucket":      acctest.Representation{RepType: acctest.Optional, Create: `bucket`},
		"catalog_id":  acctest.Representation{RepType: acctest.Optional, Create: `catalogId`},
		"checksum":    acctest.Representation{RepType: acctest.Optional, Create: `checksum`},
		"namespace":   acctest.Representation{RepType: acctest.Optional, Create: `namespace`},
		"object":      acctest.Representation{RepType: acctest.Optional, Create: `object`},
	}
	FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsExecutionDetailsCredentialsRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":           acctest.Representation{RepType: acctest.Optional, Create: `id`},
	}
	FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsExecutionDetailsVariablesRepresentation = map[string]interface{}{
		"input_variables":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsExecutionDetailsVariablesInputVariablesRepresentation},
		"output_variables": acctest.Representation{RepType: acctest.Optional, Create: []string{`outputVariables`}},
	}
	FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesRunOnPreviousTaskInstanceDetailsRepresentation = map[string]interface{}{
		"output_variable_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesRunOnPreviousTaskInstanceDetailsOutputVariableDetailsRepresentation},
		"resource_id":             acctest.Representation{RepType: acctest.Optional, Create: `resourceId`},
		"resource_type":           acctest.Representation{RepType: acctest.Optional, Create: `resourceType`},
	}
	FleetAppsManagementRunbookRunbookVersionGroupsPropertiesRunOnPreviousTaskInstanceDetailsOutputVariableDetailsRepresentation = map[string]interface{}{
		"output_variable_name": acctest.Representation{RepType: acctest.Optional, Create: `outputVariableName`},
		"step_name":            acctest.Representation{RepType: acctest.Optional, Create: `stepName`},
	}
	FleetAppsManagementRunbookRunbookVersionTasksTaskRecordDetailsExecutionDetailsVariablesInputVariablesRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"type":        acctest.Representation{RepType: acctest.Optional, Create: `STRING`},
	}
	FleetAppsManagementRunbookRunbookVersionTasksStepPropertiesRunOnPreviousTaskInstanceDetailsOutputVariableDetailsRepresentation = map[string]interface{}{
		"output_variable_name": acctest.Representation{RepType: acctest.Optional, Create: `outputVariableName`},
		"step_name":            acctest.Representation{RepType: acctest.Optional, Create: `stepName`},
	}

	FleetAppsManagementRunbookResourceDependencies = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementRunbookResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementRunbookResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_create", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

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
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "type", "USER_DEFINED"),
				resource.TestCheckResourceAttr(resourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.step_name", "StepName"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.name", "Parallel_resource_group"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.step_name", "StepName"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.scope", "LOCAL"),

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
				resource.TestCheckResourceAttr(resourceName, "runbook_version.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.step_name", "StepName"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.steps.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),

				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.name", "Parallel_resource_group"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.properties.0.action_on_failure", "ABORT"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.groups.0.properties.0.notification_preferences.#"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.groups.0.properties.0.pause_details.#"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.groups.0.properties.0.run_on.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.is_latest", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.rollback_workflow_details.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.output_variable_mappings.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.step_name", "StepName"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.step_properties.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.command"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.content.#"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.credentials.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_executable_content", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_locked", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_read_output_variable_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.system_variables.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.type", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.output_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.is_apply_subject_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.is_copy_to_library_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.is_discovery_output_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.name", "StepName"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.os_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.platform", "Oracle Linux"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.0.num_retries", "10"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.0.timeout_in_seconds", "1000"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.scope", "LOCAL"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.version", "1"),

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "estimated_time", "PT1H"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sudo_access_needed", "false"),
				resource.TestCheckResourceAttr(resourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(resourceName, "os_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "platform", "Oracle Linux"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FleetAppsManagementRunbookResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetAppsManagementRunbookRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "runbook_version.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.step_name", "StepName"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.steps.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.name", "Parallel_resource_group"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.properties.0.action_on_failure", "ABORT"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.groups.0.properties.0.notification_preferences.#"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.groups.0.properties.0.pause_details.#"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.groups.0.properties.0.run_on.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.is_latest", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.rollback_workflow_details.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.output_variable_mappings.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.step_name", "StepName"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.step_properties.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.command"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.content.#"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.credentials.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_executable_content", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_locked", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_read_output_variable_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.system_variables.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.type", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.output_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.is_apply_subject_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.is_copy_to_library_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.is_discovery_output_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.name", "StepName"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.os_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.platform", "Oracle Linux"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.0.num_retries", "10"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.0.timeout_in_seconds", "1000"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.scope", "LOCAL"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.version", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "estimated_time", "PT1H"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sudo_access_needed", "false"),
				resource.TestCheckResourceAttr(resourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(resourceName, "os_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "platform", "Oracle Linux"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
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
			Config: config + testTaskRecordStr + compartmentIdVariableStr + FleetAppsManagementRunbookResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Optional, acctest.Update, FleetAppsManagementRunbookRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "runbook_version.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.step_name", "StepName"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.steps.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.name", "Parallel_resource_group"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.properties.0.action_on_failure", "ABORT"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.groups.0.properties.0.notification_preferences.#"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.groups.0.properties.0.pause_details.#"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.groups.0.properties.0.run_on.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.is_latest", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.rollback_workflow_details.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.output_variable_mappings.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.step_name", "StepName"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.step_properties.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.description"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.command"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.content.#"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.credentials.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_executable_content", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_locked", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_read_output_variable_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.system_variables.#"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.type", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.output_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.is_apply_subject_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.is_copy_to_library_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.is_discovery_output_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.name", "StepName"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.os_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.platform", "Oracle Linux"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.0.num_retries", "10"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.0.timeout_in_seconds", "1000"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.tasks.0.task_record_details.0.scope", "LOCAL"),
				resource.TestCheckResourceAttr(resourceName, "runbook_version.0.version", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "estimated_time", "PT30M"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_sudo_access_needed", "true"),
				resource.TestCheckResourceAttr(resourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(resourceName, "os_type", "WINDOWS"),
				resource.TestCheckResourceAttr(resourceName, "platform", "Oracle Java"),
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
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(datasourceName, "platform", "Oracle Linux"),
				resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type", "USER_DEFINED"),
				resource.TestCheckResourceAttr(datasourceName, "runbook_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "runbook_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Required, acctest.Create, FleetAppsManagementRunbookSingularDataSourceRepresentation) +
				testTaskRecordStr + compartmentIdVariableStr + FleetAppsManagementRunbookResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.step_name", "StepName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.steps.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.groups.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.groups.0.name", "Parallel_resource_group"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.groups.0.properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.groups.0.properties.0.action_on_failure", "ABORT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.groups.0.properties.0.notification_preferences.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.groups.0.properties.0.pause_details.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.groups.0.properties.0.run_on.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.is_latest", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.rollback_workflow_details.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.tasks.0.output_variable_mappings.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.step_name", "StepName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.tasks.0.step_properties.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.command"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.content.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.credentials.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_executable_content", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_locked", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.is_read_output_variable_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.system_variables.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.type", "STRING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.execution_details.0.variables.0.output_variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.is_apply_subject_task", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.is_copy_to_library_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.is_discovery_output_task", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.name", "StepName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.os_type", "LINUX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.platform", "Oracle Linux"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.0.num_retries", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.properties.0.timeout_in_seconds", "1000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.tasks.0.task_record_details.0.scope", "LOCAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runbook_version.0.version", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "estimated_time", "PT30M"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_default", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_sudo_access_needed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "operation", "PATCH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_type", "WINDOWS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "platform", "Oracle Java"),
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
