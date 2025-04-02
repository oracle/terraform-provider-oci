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
	StackMonitoringImportResourcesTaskRequiredOnlyResource = StackMonitoringImportResourcesTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Required, acctest.Create, StackMonitoringImportResourcesTaskRepresentation)

	StackMonitoringImportResourcesTaskResourceConfig = StackMonitoringImportResourcesTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Update, StackMonitoringImportResourcesTaskRepresentation)

	StackMonitoringImportResourcesTaskSingularDataSourceRepresentation = map[string]interface{}{
		"monitored_resource_task_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitored_resource_task.test_monitored_resource_task.id}`},
	}

	StackMonitoringImportResourcesTaskDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `ACCEPTED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringImportResourcesTaskDataSourceFilterRepresentation}}
	StackMonitoringImportResourcesTaskDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_monitored_resource_task.test_monitored_resource_task.id}`}},
	}

	StackMonitoringImportResourcesTaskRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"task_details":   acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringImportResourcesTaskTaskDetailsRepresentation},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `${var.task_name}`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreTaskSensitiveDataRepresentation},
	}
	StackMonitoringImportResourcesTaskTaskDetailsRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `oci_terraform_namespace`},
		"source":    acctest.Representation{RepType: acctest.Required, Create: `OCI_TELEMETRY_NATIVE`},
		"type":      acctest.Representation{RepType: acctest.Required, Create: `IMPORT_OCI_TELEMETRY_RESOURCES`},
		"availability_proxy_metric_collection_interval": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_proxy_metrics":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`availabilityProxyMetrics`}},
		"console_path_prefix":                           acctest.Representation{RepType: acctest.Optional, Create: `consolePathPrefix`},
		"external_id_mapping":                           acctest.Representation{RepType: acctest.Optional, Create: `id`},
		"lifecycle_status_mappings_for_up_status":       acctest.Representation{RepType: acctest.Optional, Create: []string{`lifecycleStatusMappingsForUpStatus`}},
		"resource_group":                                acctest.Representation{RepType: acctest.Optional, Create: `tf_group`},
		"resource_name_filter":                          acctest.Representation{RepType: acctest.Optional, Create: `resourceNameFilter`},
		"resource_name_mapping":                         acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"resource_type_filter":                          acctest.Representation{RepType: acctest.Optional, Create: `resourceTypeFilter`},
		"resource_type_mapping":                         acctest.Representation{RepType: acctest.Optional, Create: `resourceTypeMapping`},
		"service_base_url":                              acctest.Representation{RepType: acctest.Optional, Create: `serviceBaseUrl`},
		"should_use_metrics_flow_for_status":            acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	//Get API does not return sensitive data, it returns null
	ignoreTaskSensitiveDataRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{
			RepType: acctest.Required, Create: []string{
				`freeform_tags`, `defined_tags`, `system_tags`, `task_details`,
			}},
	}

	StackMonitoringImportResourcesTaskResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringImportResourcesTaskResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringImportResourcesTaskResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_stack_monitoring_monitored_resource_task.test_monitored_resource_task"
	datasourceName := "data.oci_stack_monitoring_monitored_resource_tasks.test_monitored_resource_tasks"
	singularDatasourceName := "data.oci_stack_monitoring_monitored_resource_task.test_monitored_resource_task"

	currentTime, _ := time.Now().UTC().MarshalText()
	taskName := "terraform_task_name_" + string(currentTime)
	taskNameVariableStr := fmt.Sprintf("variable \"task_name\" { default = \"%s\" }\n", taskName)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+taskNameVariableStr+StackMonitoringImportResourcesTaskResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Create, StackMonitoringImportResourcesTaskRepresentation), "stackmonitoring", "ImportResourcesTask", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringImportResourcesTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Required, acctest.Create, StackMonitoringImportResourcesTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.namespace", "oci_terraform_namespace"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.source", "OCI_TELEMETRY_NATIVE"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "IMPORT_OCI_TELEMETRY_RESOURCES"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringImportResourcesTaskResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringImportResourcesTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Create, StackMonitoringImportResourcesTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.availability_proxy_metric_collection_interval", "10"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.availability_proxy_metrics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.console_path_prefix", "consolePathPrefix"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.external_id_mapping", "id"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.lifecycle_status_mappings_for_up_status.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_name_filter", "resourceNameFilter"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_name_mapping", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_type_filter", "resourceTypeFilter"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_type_mapping", "resourceTypeMapping"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.service_base_url", "serviceBaseUrl"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.should_use_metrics_flow_for_status", "false"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.namespace", "oci_terraform_namespace"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_group", "tf_group"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.source", "OCI_TELEMETRY_NATIVE"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "IMPORT_OCI_TELEMETRY_RESOURCES"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + taskNameVariableStr + StackMonitoringImportResourcesTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(StackMonitoringImportResourcesTaskRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.availability_proxy_metric_collection_interval", "10"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.availability_proxy_metrics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.console_path_prefix", "consolePathPrefix"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.external_id_mapping", "id"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.lifecycle_status_mappings_for_up_status.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_name_filter", "resourceNameFilter"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_name_mapping", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_type_filter", "resourceTypeFilter"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_type_mapping", "resourceTypeMapping"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.service_base_url", "serviceBaseUrl"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.should_use_metrics_flow_for_status", "false"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.namespace", "oci_terraform_namespace"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_group", "tf_group"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.source", "OCI_TELEMETRY_NATIVE"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "IMPORT_OCI_TELEMETRY_RESOURCES"),

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
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringImportResourcesTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Update, StackMonitoringImportResourcesTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.availability_proxy_metric_collection_interval", "10"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.availability_proxy_metrics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.console_path_prefix", "consolePathPrefix"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.external_id_mapping", "id"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.lifecycle_status_mappings_for_up_status.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_name_filter", "resourceNameFilter"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_name_mapping", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_type_filter", "resourceTypeFilter"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_type_mapping", "resourceTypeMapping"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.service_base_url", "serviceBaseUrl"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.should_use_metrics_flow_for_status", "false"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.namespace", "oci_terraform_namespace"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_group", "tf_group"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.source", "OCI_TELEMETRY_NATIVE"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "IMPORT_OCI_TELEMETRY_RESOURCES"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_tasks", "test_monitored_resource_tasks", acctest.Optional, acctest.Update, StackMonitoringImportResourcesTaskDataSourceRepresentation) +
				compartmentIdVariableStr + taskNameVariableStr + StackMonitoringImportResourcesTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Update, StackMonitoringImportResourcesTaskRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Required, acctest.Create, StackMonitoringImportResourcesTaskSingularDataSourceRepresentation) +
				compartmentIdVariableStr + taskNameVariableStr + StackMonitoringImportResourcesTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitored_resource_task_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", taskName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.availability_proxy_metric_collection_interval", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.availability_proxy_metrics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.console_path_prefix", "consolePathPrefix"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.external_id_mapping", "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.lifecycle_status_mappings_for_up_status.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.resource_name_filter", "resourceNameFilter"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.resource_name_mapping", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.resource_type_filter", "resourceTypeFilter"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.resource_type_mapping", "resourceTypeMapping"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.service_base_url", "serviceBaseUrl"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.should_use_metrics_flow_for_status", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.namespace", "oci_terraform_namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.resource_group", "tf_group"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.source", "OCI_TELEMETRY_NATIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.type", "IMPORT_OCI_TELEMETRY_RESOURCES"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "work_request_ids.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringImportResourcesTaskRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"defined_tags"},
			ResourceName:            resourceName,
		},
	})
}
