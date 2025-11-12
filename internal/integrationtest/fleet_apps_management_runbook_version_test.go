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
	FleetAppsManagementRunbookVersionRequiredOnlyResource = FleetAppsManagementRunbookVersionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook_version", "test_runbook_version", acctest.Required, acctest.Create, FleetAppsManagementRunbookVersionRepresentation)

	FleetAppsManagementRunbookVersionResourceConfig = FleetAppsManagementRunbookVersionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook_version", "test_runbook_version", acctest.Optional, acctest.Update, FleetAppsManagementRunbookVersionRepresentation)

	FleetAppsManagementRunbookVersionSingularDataSourceRepresentation = map[string]interface{}{
		"runbook_version_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_runbook_version.test_runbook_version.id}`},
	}

	FleetAppsManagementRunbookVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_runbook_version.test_runbook_version.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `TestRunbookTFP6`},
		"runbook_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_runbook.test_runbook.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionDataSourceFilterRepresentation}}
	FleetAppsManagementRunbookVersionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_runbook_version.test_runbook_version.id}`}},
	}

	FleetAppsManagementRunbookVersionRepresentation = map[string]interface{}{
		"execution_workflow_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionExecutionWorkflowDetailsRepresentation},
		"groups":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionGroupsRepresentation},
		"runbook_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_runbook.test_runbook.id}`},
		// "tasks":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionTasksRepresentation},
		"tasks": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionTasksRepresentation}, {RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionRollbackTasksRepresentation}},
		// "defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("Oracle-Tags.CreatedBy", "value")}`, Update: `${map("Oracle-Tags.CreatedBy", "updatedValue")}`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"rollback_workflow_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionRollbackWorkflowDetailsRepresentation},
		// "export_trigger":            acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}
	FleetAppsManagementRunbookVersionExecutionWorkflowDetailsRepresentation = map[string]interface{}{
		"workflow": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionExecutionWorkflowDetailsWorkflowRepresentation},
	}
	FleetAppsManagementRunbookVersionGroupsRepresentation = map[string]interface{}{
		"name":       acctest.Representation{RepType: acctest.Required, Create: `Parallel_resource_group`},
		"type":       acctest.Representation{RepType: acctest.Required, Create: `PARALLEL_RESOURCE_GROUP`},
		"properties": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionGroupsPropertiesRepresentation},
	}
	FleetAppsManagementRunbookVersionTasksRepresentation = map[string]interface{}{
		"step_name":           acctest.Representation{RepType: acctest.Required, Create: `stepName`, Update: `stepName2`},
		"task_record_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionTasksTaskRecordDetailsRepresentation},
		"step_properties":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksStepPropertiesRepresentation},
	}
	FleetAppsManagementRunbookVersionRollbackTasksRepresentation = map[string]interface{}{
		"step_name":           acctest.Representation{RepType: acctest.Required, Create: `stepNameRollback`, Update: `stepNameRollback2`},
		"task_record_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionTasksTaskRecordDetailsRepresentation},
		"step_properties":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksStepPropertiesRepresentation},
	}
	FleetAppsManagementRunbookVersionRollbackWorkflowDetailsRepresentation = map[string]interface{}{
		"scope":    acctest.Representation{RepType: acctest.Required, Create: `ACTION_GROUP`, Update: `TARGET`},
		"workflow": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionRollbackWorkflowDetailsWorkflowRepresentation},
	}
	FleetAppsManagementRunbookVersionExecutionWorkflowDetailsWorkflowRepresentation = map[string]interface{}{
		"group_name": acctest.Representation{RepType: acctest.Required, Create: `Parallel_resource_group`},
		"steps":      acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionExecutionWorkflowDetailsWorkflowStepsRepresentation},
		"type":       acctest.Representation{RepType: acctest.Required, Create: `PARALLEL_RESOURCE_GROUP`},
	}
	FleetAppsManagementRunbookVersionGroupsPropertiesRepresentation = map[string]interface{}{
		"action_on_failure":        acctest.Representation{RepType: acctest.Required, Create: `ABORT`, Update: `CONTINUE`},
		"notification_preferences": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionGroupsPropertiesNotificationPreferencesRepresentation},
		"pause_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionGroupsPropertiesPauseDetailsRepresentation},
		"pre_condition":            acctest.Representation{RepType: acctest.Optional, Create: `target.product.name == \"Oracle Weblogic Server\"`, Update: `target.product.name == \"Oracle Linux Server\"`},
		"run_on":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionGroupsPropertiesRunOnRepresentation},
	}
	FleetAppsManagementRunbookVersionTasksTaskRecordDetailsRepresentation = map[string]interface{}{
		// Only of 'LOCAL OR SHARED'
		"scope": acctest.Representation{RepType: acctest.Required, Create: `LOCAL`},
		// descrip only applicable to LOCAL scope
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"execution_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionTasksTaskRecordDetailsExecutionDetailsRepresentation},
		// Lines 97-104 only of scope LOCAL
		"is_apply_subject_task":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_copy_to_library_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_discovery_output_task":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"name":                       acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"os_type":                    acctest.Representation{RepType: acctest.Optional, Create: `WINDOWS`, Update: `LINUX`},
		"platform":                   acctest.Representation{RepType: acctest.Optional, Create: `Oracle Fusion Middleware`, Update: `Oracle Database`},
		"properties":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksTaskRecordDetailsPropertiesRepresentation},
	}
	FleetAppsManagementRunbookVersionTasksOutputVariableMappingsRepresentation = map[string]interface{}{
		"name":                    acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"output_variable_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionTasksOutputVariableMappingsOutputVariableDetailsRepresentation},
	}
	FleetAppsManagementRunbookVersionTasksStepPropertiesRepresentation = map[string]interface{}{
		"action_on_failure":        acctest.Representation{RepType: acctest.Required, Create: `ABORT`, Update: `CONTINUE`},
		"notification_preferences": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksStepPropertiesNotificationPreferencesRepresentation},
		"pause_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksStepPropertiesPauseDetailsRepresentation},
		"pre_condition":            acctest.Representation{RepType: acctest.Optional, Create: `target.product.name == \"Oracle Weblogic Server\"`, Update: `target.product.name == \"Oracle Weblogic Server\"`},
		"run_on":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksStepPropertiesRunOnRepresentation},
	}
	FleetAppsManagementRunbookVersionRollbackWorkflowDetailsWorkflowRepresentation = map[string]interface{}{
		"group_name": acctest.Representation{RepType: acctest.Required, Create: `Parallel_resource_group`},
		"steps":      acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionRollbackWorkflowDetailsWorkflowStepsRepresentation},
		"type":       acctest.Representation{RepType: acctest.Required, Create: `PARALLEL_RESOURCE_GROUP`},
	}
	FleetAppsManagementRunbookVersionExecutionWorkflowDetailsWorkflowStepsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `TASK`},
		//"group_name": acctest.Representation{RepType: acctest.Optional, Create: `Parallel_resource_group`},
		"step_name": acctest.Representation{RepType: acctest.Required, Create: `stepName`, Update: `stepName2`},
	}
	FleetAppsMgmtRunbookVersionsExecutionWorkflowDetailsWorkflowStep1 = map[string]interface{}{
		"type":       acctest.Representation{RepType: acctest.Required, Create: `TASK`, Update: `PARALLEL_TASK_GROUP`},
		"group_name": acctest.Representation{RepType: acctest.Optional, Create: `Parallel_resource_group`},
		"step_name":  acctest.Representation{RepType: acctest.Required, Create: `stepName1`, Update: `stepName1b`},
	}
	FleetAppsMgmtRunbookVersionsExecutionWorkflowDetailsWorkflowStep2 = map[string]interface{}{
		"type":       acctest.Representation{RepType: acctest.Required, Create: `TASK`, Update: `PARALLEL_TASK_GROUP`},
		"group_name": acctest.Representation{RepType: acctest.Optional, Create: `Parallel_resource_group`},
		"step_name":  acctest.Representation{RepType: acctest.Required, Create: `stepName2`, Update: `stepName2b`},
	}
	FleetAppsManagementRunbookVersionGroupsPropertiesNotificationPreferencesRepresentation = map[string]interface{}{
		"should_notify_on_pause":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_notify_on_task_failure": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_notify_on_task_success": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	FleetAppsManagementRunbookVersionGroupsPropertiesPauseDetailsRepresentation = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Required, Create: `USER_ACTION`},
		// When PauseDetails KIND is 'USER_ACTION' there are not duration in minutes, only used for 'TIME_BASED' kind
		//"duration_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	FleetAppsManagementRunbookVersionGroupsPropertiesRunOnRepresentation = map[string]interface{}{
		// KIND of RunOn representation: [SCHEDULED_INSTANCES,PREVIOUS_TASK_INSTANCES,SELF_HOSTED_INSTANCES]
		"kind":      acctest.Representation{RepType: acctest.Required, Create: `SCHEDULED_INSTANCES`},
		"condition": acctest.Representation{RepType: acctest.Optional, Create: `target.product.name == \"Oracle Weblogic Server\"`, Update: `target.product.name == \"Oracle Linux Server\"`},
		// When Kind is SELF_HOSTED_INSTANCES host should be an ocid of the instance. not needed for other 2 kinds
		//"host":      acctest.Representation{RepType: acctest.Optional, Create: `host`, Update: `host2`},
		// Below only used for kind: PREVIOUS_TASK_INSTANCES
		//"previous_task_instance_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionGroupsPropertiesRunOnPreviousTaskInstanceDetailsRepresentation},
	}
	FleetAppsManagementRunbookVersionTasksTaskRecordDetailsExecutionDetailsRepresentation = map[string]interface{}{
		"execution_type": acctest.Representation{RepType: acctest.Required, Create: `SCRIPT`},
		// "catalog_id":                      acctest.Representation{RepType: acctest.Optional, Create: `${var.test_catalog_id}`},
		"command": acctest.Representation{RepType: acctest.Required, Create: `pwd`, Update: `ls -la`},
		// "config_file":                     acctest.Representation{RepType: acctest.Optional, Create: `configFile`, Update: `configFile2`},
		// "content":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksTaskRecordDetailsExecutionDetailsContentRepresentation},
		// "credentials":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksTaskRecordDetailsExecutionDetailsCredentialsRepresentation},
		// "endpoint":                        acctest.Representation{RepType: acctest.Optional, Create: `endpoint`, Update: `endpoint2`},
		"is_executable_content":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_locked":                       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_read_output_variable_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"variables":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksTaskRecordDetailsExecutionDetailsVariablesRepresentation},
	}
	FleetAppsManagementRunbookVersionTasksTaskRecordDetailsPropertiesRepresentation = map[string]interface{}{
		"num_retries":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"timeout_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	FleetAppsManagementRunbookVersionTasksOutputVariableMappingsOutputVariableDetailsRepresentation = map[string]interface{}{
		"output_variable_name": acctest.Representation{RepType: acctest.Required, Create: `outputVariableName`, Update: `outputVariableName2`},
		"step_name":            acctest.Representation{RepType: acctest.Required, Create: `stepName`, Update: `stepName2`},
	}
	FleetAppsManagementRunbookVersionTasksStepPropertiesNotificationPreferencesRepresentation = map[string]interface{}{
		"should_notify_on_pause":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_notify_on_task_failure": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"should_notify_on_task_success": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	FleetAppsManagementRunbookVersionTasksStepPropertiesPauseDetailsRepresentation = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Required, Create: `USER_ACTION`},
		// when KIND is TIME_BASED you have duration, else you don't
		//"duration_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	FleetAppsManagementRunbookVersionTasksStepPropertiesRunOnRepresentation = map[string]interface{}{
		// kind can be one of [SCHEDULE_INSTANCES,SELF_HOSTED_INSTANCES,PREVIOUS_TASK_INSTANCES]
		"kind":      acctest.Representation{RepType: acctest.Required, Create: `SCHEDULED_INSTANCES`, Update: `SCHEDULED_INSTANCES`},
		"condition": acctest.Representation{RepType: acctest.Optional, Create: `target.product.name == \"Oracle Weblogic Server\"`, Update: `target.product.name == \"Oracle Linux Server\"`},
		// When SELF_HOSTED_INSTANCES host is the ocid of the insnce
		//"host":                           acctest.Representation{RepType: acctest.Optional, Create: `host`, Update: `host2`},
		// must be of PREVIOUS_TASK_INSTANCES kind
		//"previous_task_instance_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksStepPropertiesRunOnPreviousTaskInstanceDetailsRepresentation},
	}
	FleetAppsManagementRunbookVersionRollbackWorkflowDetailsWorkflowStepsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `TASK`},
		// "group_name": acctest.Representation{RepType: acctest.Required, Create: `Parallel_resource_group`},
		"step_name": acctest.Representation{RepType: acctest.Required, Create: `stepNameRollback`, Update: `stepNameRollback2`},
	}
	FleetAppsManagementRunbookVersionGroupsPropertiesRunOnPreviousTaskInstanceDetailsRepresentation = map[string]interface{}{
		"output_variable_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionGroupsPropertiesRunOnPreviousTaskInstanceDetailsOutputVariableDetailsRepresentation},
		"resource_id":             acctest.Representation{RepType: acctest.Optional, Create: `test_resource_id`},
		"resource_type":           acctest.Representation{RepType: acctest.Optional, Create: `resourceType`, Update: `resourceType2`},
	}
	FleetAppsManagementRunbookVersionTasksTaskRecordDetailsExecutionDetailsContentRepresentation = map[string]interface{}{
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_BUCKET`, Update: `CATALOG`},
		"bucket":      acctest.Representation{RepType: acctest.Optional, Create: `bucket`, Update: `bucket2`},
		"catalog_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.test_catalog_id}`},
		"checksum":    acctest.Representation{RepType: acctest.Optional, Create: `checksum`, Update: `checksum2`},
		"namespace":   acctest.Representation{RepType: acctest.Optional, Create: `namespace`, Update: `namespace2`},
		"object":      acctest.Representation{RepType: acctest.Optional, Create: `object`, Update: `object2`},
	}
	FleetAppsManagementRunbookVersionTasksTaskRecordDetailsExecutionDetailsCredentialsRepresentation = map[string]interface{}{
		// ID is the ocid of the credentials within the Fleet Apps Mgmt Metadata
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `${var.creds_display_name}`},
		"id":           acctest.Representation{RepType: acctest.Required, Create: `${var.creds_ocid}`},
	}
	FleetAppsManagementRunbookVersionTasksTaskRecordDetailsExecutionDetailsVariablesRepresentation = map[string]interface{}{
		"input_variables":  acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementRunbookVersionTasksTaskRecordDetailsExecutionDetailsVariablesInputVariablesRepresentation},
		"output_variables": acctest.Representation{RepType: acctest.Optional, Create: []string{`outputVariables`}, Update: []string{`outputVariables2`}},
	}
	FleetAppsManagementRunbookVersionTasksStepPropertiesRunOnPreviousTaskInstanceDetailsRepresentation = map[string]interface{}{
		"output_variable_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementRunbookVersionTasksStepPropertiesRunOnPreviousTaskInstanceDetailsOutputVariableDetailsRepresentation},
		"resource_id":             acctest.Representation{RepType: acctest.Optional, Create: `test_resource_id`},
		"resource_type":           acctest.Representation{RepType: acctest.Optional, Create: `resourceType`, Update: `resourceType2`},
	}
	FleetAppsManagementRunbookVersionGroupsPropertiesRunOnPreviousTaskInstanceDetailsOutputVariableDetailsRepresentation = map[string]interface{}{
		"output_variable_name": acctest.Representation{RepType: acctest.Optional, Create: `outputVariableName`, Update: `outputVariableName2`},
		"step_name":            acctest.Representation{RepType: acctest.Optional, Create: `stepName`, Update: `stepName2`},
	}
	FleetAppsManagementRunbookVersionTasksTaskRecordDetailsExecutionDetailsVariablesInputVariablesRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `inputVarDescription`, Update: `inputVarDescription2`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `inputVarName`, Update: `inputVarName2`},
		"type":        acctest.Representation{RepType: acctest.Optional, Create: `STRING`},
	}
	FleetAppsManagementRunbookVersionTasksStepPropertiesRunOnPreviousTaskInstanceDetailsOutputVariableDetailsRepresentation = map[string]interface{}{
		"output_variable_name": acctest.Representation{RepType: acctest.Optional, Create: `outputVariableName`, Update: `outputVariableName2`},
		"step_name":            acctest.Representation{RepType: acctest.Optional, Create: `stepName`, Update: `stepName2`},
	}

	FleetAppsManagementRunbookVersionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Required, acctest.Create, FleetAppsManagementRunbookRepresentation)
	//GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementRunbookVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementRunbookVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	testCatalogId := utils.GetEnvSettingWithBlankDefault("catalog_id")
	credsDisplayName := utils.GetEnvSettingWithBlankDefault("credential_name")
	credsId := utils.GetEnvSettingWithBlankDefault("credential")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	testCatalogIdVariableStr := fmt.Sprintf("variable \"test_catalog_id\" { default = \"%s\" }\n", testCatalogId)
	credsDisplayNameVariableStr := fmt.Sprintf("variable \"creds_display_name\" { default = \"%s\" }\n", credsDisplayName)
	credsIdVariableStr := fmt.Sprintf("variable \"creds_ocid\" { default = \"%s\" }\n", credsId)

	resourceName := "oci_fleet_apps_management_runbook_version.test_runbook_version"
	datasourceName := "data.oci_fleet_apps_management_runbook_versions.test_runbook_versions"
	singularDatasourceName := "data.oci_fleet_apps_management_runbook_version.test_runbook_version"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+testCatalogIdVariableStr+credsDisplayNameVariableStr+credsIdVariableStr+FleetAppsManagementRunbookVersionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook_version", "test_runbook_version", acctest.Optional, acctest.Create, FleetAppsManagementRunbookVersionRepresentation), "fleetappsmanagement", "runbookVersion", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementRunbookVersionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementRunbookVersionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook_version", "test_runbook_version", acctest.Required, acctest.Create, FleetAppsManagementRunbookVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.0.steps.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "execution_workflow_details.0.workflow.0.steps.0.group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "execution_workflow_details.0.workflow.0.steps.0.step_name"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.name", "Parallel_resource_group"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tasks.#"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_name", "stepName"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.#", "1"),
				//resource.TestCheckResourceAttrSet(resourceName, "tasks.0.task_record_details.0.execution_details.0.catalog_id"),
				//resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.endpoint", "endpoint"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.execution_type", "SCRIPT"),
				//resource.TestCheckResourceAttrSet(resourceName, "tasks.0.task_record_details.0.execution_details.0.target_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.scope", "LOCAL"),
				//resource.TestCheckResourceAttrSet(resourceName, "tasks.0.task_record_details.0.task_record_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementRunbookVersionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + testCatalogIdVariableStr + FleetAppsManagementRunbookVersionResourceDependencies + credsDisplayNameVariableStr + credsIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook_version", "test_runbook_version", acctest.Optional, acctest.Create, FleetAppsManagementRunbookVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "execution_workflow_details.0.workflow.0.steps.0.step_name"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.name", "Parallel_resource_group"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.action_on_failure", "ABORT"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.notification_preferences.0.should_notify_on_pause", "false"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.notification_preferences.0.should_notify_on_task_failure", "false"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.notification_preferences.0.should_notify_on_task_success", "false"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.pause_details.0.kind", "USER_ACTION"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.pre_condition", "target.product.name == \"Oracle Weblogic Server\""),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.run_on.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.run_on.0.condition", "target.product.name == \"Oracle Weblogic Server\""),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.run_on.0.kind", "SCHEDULED_INSTANCES"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tasks.#"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_name", "stepName"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.action_on_failure", "ABORT"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.notification_preferences.0.should_notify_on_pause", "false"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_failure", "false"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_success", "false"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.pause_details.0.kind", "USER_ACTION"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.pre_condition", "target.product.name == \"Oracle Weblogic Server\""),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.run_on.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.run_on.0.condition", "target.product.name == \"Oracle Weblogic Server\""),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.run_on.0.kind", "SCHEDULED_INSTANCES"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.command", "pwd"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.is_executable_content", "false"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.is_locked", "false"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.is_read_output_variable_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "tasks.0.task_record_details.0.execution_details.0.system_variables.#"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.description", "inputVarDescription"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.name", "inputVarName"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.type", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.output_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.is_apply_subject_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.is_copy_to_library_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.is_discovery_output_task", "false"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.os_type", "WINDOWS"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.platform", "Oracle Fusion Middleware"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.properties.0.num_retries", "10"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.properties.0.timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.scope", "LOCAL"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//fmt.Printf("FROM STATE: ResourceName %s, Resource Id: %s\n", resourceName, resId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						//fmt.Printf("EXPORT ENABLED:::::\n")
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							//fmt.Printf("errExport is not nil\n\n")
							//fmt.Printf("errExport === %s\n", errExport)
							return errExport
						}
					}
					//fmt.Printf("Error contains %s", err)
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + testCatalogIdVariableStr + FleetAppsManagementRunbookVersionResourceDependencies + credsDisplayNameVariableStr + credsIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook_version", "test_runbook_version", acctest.Optional, acctest.Update, FleetAppsManagementRunbookVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "execution_workflow_details.0.workflow.0.group_name"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.0.steps.0.step_name", "stepName2"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(resourceName, "execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.name", "Parallel_resource_group"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.action_on_failure", "CONTINUE"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.notification_preferences.0.should_notify_on_pause", "true"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.notification_preferences.0.should_notify_on_task_failure", "true"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.notification_preferences.0.should_notify_on_task_success", "true"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.pause_details.0.kind", "USER_ACTION"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.pre_condition", "target.product.name == \"Oracle Linux Server\""),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.run_on.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.run_on.0.condition", "target.product.name == \"Oracle Linux Server\""),
				resource.TestCheckResourceAttr(resourceName, "groups.0.properties.0.run_on.0.kind", "SCHEDULED_INSTANCES"),
				resource.TestCheckResourceAttr(resourceName, "groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "runbook_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tasks.#"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_name", "stepName2"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.action_on_failure", "CONTINUE"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.notification_preferences.0.should_notify_on_pause", "true"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_failure", "true"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_success", "true"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.pause_details.0.kind", "USER_ACTION"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.pre_condition", "target.product.name == \"Oracle Weblogic Server\""),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.run_on.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.run_on.0.condition", "target.product.name == \"Oracle Linux Server\""),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.step_properties.0.run_on.0.kind", "SCHEDULED_INSTANCES"), resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.command", "ls -la"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.is_executable_content", "true"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.is_locked", "true"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.is_read_output_variable_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "tasks.0.task_record_details.0.execution_details.0.system_variables.#"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.description", "inputVarDescription2"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.name", "inputVarName2"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.type", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.output_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.is_apply_subject_task", "true"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.is_copy_to_library_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.is_discovery_output_task", "true"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.os_type", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.platform", "Oracle Database"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.properties.0.num_retries", "11"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.properties.0.timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.task_record_details.0.scope", "LOCAL"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					// if resId != resId2 {
					// 	return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					fmt.Printf("Resource %s updated successfully for ID: %s\n", resourceName, resId2)
					// }
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_runbook_versions", "test_runbook_versions", acctest.Optional, acctest.Update, FleetAppsManagementRunbookVersionDataSourceRepresentation) + credsDisplayNameVariableStr + credsIdVariableStr +
				compartmentIdVariableStr + testCatalogIdVariableStr + FleetAppsManagementRunbookVersionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook_version", "test_runbook_version", acctest.Optional, acctest.Update, FleetAppsManagementRunbookVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "TestRunbookTFP6"),
				resource.TestCheckResourceAttrSet(datasourceName, "runbook_id"),
				//resource.TestCheckResourceAttr(datasourceName, "lifcyclestate", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "runbook_version_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "runbook_version_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_runbook_version", "test_runbook_version", acctest.Required, acctest.Create, FleetAppsManagementRunbookVersionSingularDataSourceRepresentation) + credsDisplayNameVariableStr + credsIdVariableStr +
				compartmentIdVariableStr + testCatalogIdVariableStr + FleetAppsManagementRunbookVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runbook_version_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_workflow_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_workflow_details.0.workflow.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_workflow_details.0.workflow.0.steps.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_workflow_details.0.workflow.0.steps.0.step_name", "stepName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_workflow_details.0.workflow.0.steps.0.type", "TASK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_workflow_details.0.workflow.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.name", "Parallel_resource_group"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.action_on_failure", "CONTINUE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.notification_preferences.0.should_notify_on_pause", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.notification_preferences.0.should_notify_on_task_failure", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.notification_preferences.0.should_notify_on_task_success", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.pause_details.0.kind", "USER_ACTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.pre_condition", "target.product.name == \"Oracle Linux Server\""),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.run_on.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.run_on.0.condition", "target.product.name == \"Oracle Linux Server\""),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.properties.0.run_on.0.kind", "SCHEDULED_INSTANCES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "groups.0.type", "PARALLEL_RESOURCE_GROUP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tasks.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_name", "stepName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.action_on_failure", "CONTINUE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.notification_preferences.0.should_notify_on_pause", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_failure", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.notification_preferences.0.should_notify_on_task_success", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.pause_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.pause_details.0.kind", "USER_ACTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.pre_condition", "target.product.name == \"Oracle Weblogic Server\""),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.run_on.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.run_on.0.condition", "target.product.name == \"Oracle Linux Server\""),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.step_properties.0.run_on.0.kind", "SCHEDULED_INSTANCES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.command", "ls -la"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.execution_type", "SCRIPT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.is_executable_content", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.is_locked", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.is_read_output_variable_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.system_variables.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.description", "inputVarDescription2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.name", "inputVarName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.input_variables.0.type", "STRING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.execution_details.0.variables.0.output_variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.is_apply_subject_task", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.is_copy_to_library_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.is_discovery_output_task", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.os_type", "LINUX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.platform", "Oracle Database"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.properties.0.num_retries", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.properties.0.timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.task_record_details.0.scope", "LOCAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + FleetAppsManagementRunbookVersionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementRunbookVersionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementRunbooksClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_runbook_version" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetRunbookVersionRequest{}

			tmp := rs.Primary.ID
			request.RunbookVersionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetRunbookVersion(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.RunbookVersionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementRunbookVersion") {
		resource.AddTestSweepers("FleetAppsManagementRunbookVersion", &resource.Sweeper{
			Name:         "FleetAppsManagementRunbookVersion",
			Dependencies: acctest.DependencyGraph["runbookVersion"],
			F:            sweepFleetAppsManagementRunbookVersionResource,
		})
	}
}

func sweepFleetAppsManagementRunbookVersionResource(compartment string) error {
	fleetAppsManagementRunbooksClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementRunbooksClient()
	runbookVersionIds, err := getFleetAppsManagementRunbookVersionIds(compartment)
	if err != nil {
		return err
	}
	for _, runbookVersionId := range runbookVersionIds {
		if ok := acctest.SweeperDefaultResourceId[runbookVersionId]; !ok {
			deleteRunbookVersionRequest := oci_fleet_apps_management.DeleteRunbookVersionRequest{}

			deleteRunbookVersionRequest.RunbookVersionId = &runbookVersionId

			deleteRunbookVersionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementRunbooksClient.DeleteRunbookVersion(context.Background(), deleteRunbookVersionRequest)
			if error != nil {
				fmt.Printf("Error deleting RunbookVersion %s %s, It is possible that the resource is already deleted. Please verify manually \n", runbookVersionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &runbookVersionId, FleetAppsManagementRunbookVersionSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementRunbookVersionSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementRunbookVersionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RunbookVersionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementRunbooksClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementRunbooksClient()

	listRunbookVersionsRequest := oci_fleet_apps_management.ListRunbookVersionsRequest{}
	listRunbookVersionsRequest.CompartmentId = &compartmentId
	listRunbookVersionsRequest.LifecycleState = oci_fleet_apps_management.RunbookVersionLifecycleStateActive
	listRunbookVersionsResponse, err := fleetAppsManagementRunbooksClient.ListRunbookVersions(context.Background(), listRunbookVersionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting RunbookVersion list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, runbookVersion := range listRunbookVersionsResponse.Items {
		id := *runbookVersion.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RunbookVersionId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementRunbookVersionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if runbookVersionResponse, ok := response.Response.(oci_fleet_apps_management.GetRunbookVersionResponse); ok {
		return runbookVersionResponse.LifecycleState != oci_fleet_apps_management.RunbookVersionLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementRunbookVersionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementRunbooksClient().GetRunbookVersion(context.Background(), oci_fleet_apps_management.GetRunbookVersionRequest{
		RunbookVersionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
