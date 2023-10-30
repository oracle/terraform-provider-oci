// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringMonitoredResourceTaskRequiredOnlyResource = StackMonitoringMonitoredResourceTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Required, acctest.Create, StackMonitoringMonitoredResourceTaskRepresentation)

	StackMonitoringMonitoredResourceTaskResourceConfig = StackMonitoringMonitoredResourceTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceTaskRepresentation)

	StackMonitoringMonitoredResourceTaskSingularDataSourceRepresentation = map[string]interface{}{
		"monitored_resource_task_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitored_resource_task.test_monitored_resource_task.id}`},
	}

	StackMonitoringMonitoredResourceTaskDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `ACCEPTED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMonitoredResourceTaskDataSourceFilterRepresentation}}
	StackMonitoringMonitoredResourceTaskDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_monitored_resource_task.test_monitored_resource_task.id}`}},
	}

	StackMonitoringMonitoredResourceTaskRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"task_details":   acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMonitoredResourceTaskTaskDetailsRepresentation},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `${var.task_name}`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreTaskSensitiveDataRepresentation},
	}
	StackMonitoringMonitoredResourceTaskTaskDetailsRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `oci_terraform_namespace`},
		"source":    acctest.Representation{RepType: acctest.Required, Create: `OCI_TELEMETRY_NATIVE`},
		"type":      acctest.Representation{RepType: acctest.Required, Create: `IMPORT_OCI_TELEMETRY_RESOURCES`},
	}

	//Get API does not return sensitive data, it returns null
	ignoreTaskSensitiveDataRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{
			RepType: acctest.Required, Create: []string{
				`freeform_tags`, `defined_tags`, `system_tags`,
			}},
	}

	StackMonitoringMonitoredResourceTaskResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoredResourceTaskResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoredResourceTaskResource_basic")
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+taskNameVariableStr+StackMonitoringMonitoredResourceTaskResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourceTaskRepresentation), "stackmonitoring", "monitoredResourceTask", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringMonitoredResourceTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Required, acctest.Create, StackMonitoringMonitoredResourceTaskRepresentation),
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
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringMonitoredResourceTaskResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringMonitoredResourceTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Create, StackMonitoringMonitoredResourceTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.identifying_properties.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.namespace", "oci_terraform_namespace"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_group", ""),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + taskNameVariableStr + StackMonitoringMonitoredResourceTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(StackMonitoringMonitoredResourceTaskRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.identifying_properties.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.namespace", "oci_terraform_namespace"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_group", ""),
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
			Config: config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringMonitoredResourceTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.identifying_properties.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.namespace", "oci_terraform_namespace"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.resource_group", ""),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_tasks", "test_monitored_resource_tasks", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceTaskDataSourceRepresentation) +
				compartmentIdVariableStr + taskNameVariableStr + StackMonitoringMonitoredResourceTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Optional, acctest.Update, StackMonitoringMonitoredResourceTaskRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_monitored_resource_task", acctest.Required, acctest.Create, StackMonitoringMonitoredResourceTaskSingularDataSourceRepresentation) +
				compartmentIdVariableStr + taskNameVariableStr + StackMonitoringMonitoredResourceTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitored_resource_task_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", taskName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.identifying_properties.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.namespace", "oci_terraform_namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.resource_group", ""),
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
			Config:                  config + compartmentIdVariableStr + taskNameVariableStr + StackMonitoringMonitoredResourceTaskRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"defined_tags"},
			ResourceName:            resourceName,
		},
	})
}
