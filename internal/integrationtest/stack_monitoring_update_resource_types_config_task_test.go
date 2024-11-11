// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringUpdateResourceTypesConfigTaskRequiredOnlyResource = StackMonitoringUpdateResourceTypesConfigTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_resource_types_config_task", acctest.Required, acctest.Create, StackMonitoringUpdateResourceTypesConfigTaskRepresentation)

	StackMonitoringUpdateResourceTypesConfigTaskResourceConfig = StackMonitoringUpdateResourceTypesConfigTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_resource_types_config_task", acctest.Optional, acctest.Update, StackMonitoringUpdateResourceTypesConfigTaskRepresentation)

	StackMonitoringUpdateResourceTypesConfigTaskSingularDataSourceRepresentation = map[string]interface{}{
		"monitored_resource_task_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitored_resource_task.test_update_resource_types_config_task.id}`},
	}

	StackMonitoringUpdateResourceTypesConfigTaskDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `ACCEPTED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringUpdateResourceTypesConfigTaskDataSourceFilterRepresentation}}
	StackMonitoringUpdateResourceTypesConfigTaskDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_monitored_resource_task.test_update_resource_types_config_task.id}`}},
	}

	StackMonitoringUpdateResourceTypesConfigTaskRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"task_details":   acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringUpdateResourceTypesConfigTaskDetailsRepresentation},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `${var.task_name}`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreUpdateResourceTypesConfigTaskSensitiveDataRepresentation},
	}
	StackMonitoringUpdateResourceTypesConfigTaskDetailsRepresentation = map[string]interface{}{
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `UPDATE_RESOURCE_TYPE_CONFIGS`},
		"handler_type":                 acctest.Representation{RepType: acctest.Required, Create: `TELEGRAF`},
		"resource_types_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringUpdateResourceTypesConfigurationRepresentation},
	}

	StackMonitoringUpdateResourceTypesConfigurationRepresentation = map[string]interface{}{
		"resource_type":               acctest.Representation{RepType: acctest.Required, Create: `test_resource_type`},
		"handler_config":              acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringUpdateResourceTypesConfigHandlerConfigRepresentation},
		"availability_metrics_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringUpdateResourceTypesConfigAvailabilityMetricsConfigRepresentation},
	}
	StackMonitoringUpdateResourceTypesConfigAvailabilityMetricsConfigRepresentation = map[string]interface{}{
		"collection_interval_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `300`},
		"metrics":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`metrics`}},
	}
	StackMonitoringUpdateResourceTypesConfigHandlerConfigRepresentation = map[string]interface{}{
		"collector_types":                   acctest.Representation{RepType: acctest.Required, Create: []string{`tsh`}},
		"telemetry_resource_group":          acctest.Representation{RepType: acctest.Required, Create: `telemetryResourceGroup`},
		"metric_upload_interval_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `300`},
		"handler_properties":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringUpdateResourceTypesConfigHandlerConfigHandlerPropertiesRepresentation},
		"metric_mappings":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringUpdateResourceTypesConfigHandlerConfigMetricMappingsRepresentation},
		"metric_name_config":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringUpdateResourceTypesConfigHandlerConfigMetricNameConfigRepresentation},
		"telegraf_resource_name_config":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringUpdateResourceTypesConfigHandlerConfigTelegrafResourceNameConfigRepresentation},
	}
	StackMonitoringUpdateResourceTypesConfigHandlerConfigHandlerPropertiesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`},
	}
	StackMonitoringUpdateResourceTypesConfigHandlerConfigMetricMappingsRepresentation = map[string]interface{}{
		"collector_metric_name":             acctest.Representation{RepType: acctest.Optional, Create: `active_counts`},
		"telemetry_metric_name":             acctest.Representation{RepType: acctest.Optional, Create: `activecounts`},
		"is_skip_upload":                    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"metric_upload_interval_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `180`},
	}
	StackMonitoringUpdateResourceTypesConfigHandlerConfigMetricNameConfigRepresentation = map[string]interface{}{
		"exclude_pattern_on_prefix":     acctest.Representation{RepType: acctest.Optional, Create: `excludePatternOnPrefix`},
		"is_prefix_with_collector_type": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	StackMonitoringUpdateResourceTypesConfigHandlerConfigTelegrafResourceNameConfigRepresentation = map[string]interface{}{
		"exclude_tags":     acctest.Representation{RepType: acctest.Optional, Create: []string{`excludeTags`}},
		"include_tags":     acctest.Representation{RepType: acctest.Optional, Create: []string{`includeTags`}},
		"is_use_tags_only": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	//Get API does not return sensitive data, it returns null
	ignoreUpdateResourceTypesConfigTaskSensitiveDataRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{
			RepType: acctest.Required, Create: []string{
				`freeform_tags`, `defined_tags`, `system_tags`, `task_details`,
			}},
	}

	StackMonitoringUpdateResourceTypesConfigTaskResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringUpdateResourceTypesConfigTaskResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringUpdateResourceTypesConfigTaskResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	managementAgentId1 := utils.GetEnvSettingWithBlankDefault("stack_mon_management_agent_id_resource1")
	if managementAgentId1 == "" {
		t.Skip("Setting environmental variable stack_mon_management_agent_id_resource1 that represents management agent with resource monitoring plugin is pre-requisite for this test")
	}
	//managementAgentId1VariableStr := fmt.Sprintf("variable \"stack_mon_management_agent_id_resource1\" { default = \"%s\" }\n", managementAgentId1)

	resourceName := "oci_stack_monitoring_monitored_resource_task.test_update_resource_types_config_task"
	datasourceName := "data.oci_stack_monitoring_monitored_resource_tasks.test_update_resource_types_config_tasks"
	singularDatasourceName := "data.oci_stack_monitoring_monitored_resource_task.test_update_resource_types_config_task"

	currentTime, _ := time.Now().UTC().MarshalText()
	taskName := "update_resource_types_config_task_name_" + string(currentTime)
	taskNameVariableStr := fmt.Sprintf("variable \"task_name\" { default = \"%s\" }\n", taskName)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+taskNameVariableStr+StackMonitoringUpdateResourceTypesConfigTaskResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_resource_types_config_task", acctest.Optional, acctest.Create, StackMonitoringUpdateResourceTypesConfigTaskRepresentation), "stackmonitoring", "monitoredResourceTask", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringUpdateResourceTypesConfigTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_resource_types_config_task", acctest.Required, acctest.Create, StackMonitoringUpdateResourceTypesConfigTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "UPDATE_RESOURCE_TYPE_CONFIGS"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.handler_type", "TELEGRAF"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringUpdateResourceTypesConfigTaskResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringUpdateResourceTypesConfigTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_resource_types_config_task", acctest.Optional, acctest.Create, StackMonitoringUpdateResourceTypesConfigTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "UPDATE_RESOURCE_TYPE_CONFIGS"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + taskNameVariableStr + StackMonitoringUpdateResourceTypesConfigTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_resource_types_config_task", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(StackMonitoringUpdateResourceTypesConfigTaskRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "UPDATE_RESOURCE_TYPE_CONFIGS"),

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
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringUpdateResourceTypesConfigTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_resource_types_config_task", acctest.Optional, acctest.Update, StackMonitoringUpdateResourceTypesConfigTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "UPDATE_RESOURCE_TYPE_CONFIGS"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_tasks", "test_update_resource_types_config_tasks", acctest.Optional, acctest.Update, StackMonitoringUpdateResourceTypesConfigTaskDataSourceRepresentation) +
				compartmentIdVariableStr + taskNameVariableStr + StackMonitoringUpdateResourceTypesConfigTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_resource_types_config_task", acctest.Optional, acctest.Update, StackMonitoringUpdateResourceTypesConfigTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "status", "ACCEPTED"),

				resource.TestCheckResourceAttr(datasourceName, "monitored_resource_tasks_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "monitored_resource_tasks_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_resource_types_config_task", acctest.Required, acctest.Create, StackMonitoringUpdateResourceTypesConfigTaskSingularDataSourceRepresentation) +
				compartmentIdVariableStr + taskNameVariableStr + StackMonitoringUpdateResourceTypesConfigTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitored_resource_task_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", taskName),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.type", "UPDATE_RESOURCE_TYPE_CONFIGS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "work_request_ids.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringUpdateResourceTypesConfigTaskRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"defined_tags"},
			ResourceName:            resourceName,
		},
	})
}
