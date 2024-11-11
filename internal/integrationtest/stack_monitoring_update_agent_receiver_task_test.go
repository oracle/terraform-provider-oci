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
	StackMonitoringUpdateAgentReceiverTaskRequiredOnlyResource = StackMonitoringUpdateAgentReceiverTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_agent_receiver_task", acctest.Required, acctest.Create, StackMonitoringUpdateAgentReceiverTaskRepresentation)

	StackMonitoringUpdateAgentReceiverTaskResourceConfig = StackMonitoringUpdateAgentReceiverTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_agent_receiver_task", acctest.Optional, acctest.Update, StackMonitoringUpdateAgentReceiverTaskRepresentation)

	StackMonitoringUpdateAgentReceiverTaskSingularDataSourceRepresentation = map[string]interface{}{
		"monitored_resource_task_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitored_resource_task.test_update_agent_receiver_task.id}`},
	}

	StackMonitoringUpdateAgentReceiverTaskDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `ACCEPTED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringUpdateAgentReceiverTaskDataSourceFilterRepresentation}}
	StackMonitoringUpdateAgentReceiverTaskDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_monitored_resource_task.test_update_agent_receiver_task.id}`}},
	}

	StackMonitoringUpdateAgentReceiverTaskRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"task_details":   acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringUpdateAgentReceiverTaskDetailsRepresentation},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `${var.task_name}`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreUpdateAgentReceiverTaskSensitiveDataRepresentation},
	}
	StackMonitoringUpdateAgentReceiverTaskDetailsRepresentation = map[string]interface{}{
		"type":                acctest.Representation{RepType: acctest.Required, Create: `UPDATE_AGENT_RECEIVER`},
		"agent_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.stack_mon_management_agent_id_resource1}`},
		"handler_type":        acctest.Representation{RepType: acctest.Required, Create: `TELEGRAF`},
		"is_enable":           acctest.Representation{RepType: acctest.Required, Create: `true`},
		"receiver_properties": acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringUpdateAgentReceiverTaskReceiverPropertiesRepresentation},
	}
	StackMonitoringUpdateAgentReceiverTaskReceiverPropertiesRepresentation = map[string]interface{}{
		"listener_port": acctest.Representation{RepType: acctest.Required, Create: `3322`},
	}

	//Get API does not return sensitive data, it returns null
	ignoreUpdateAgentReceiverTaskSensitiveDataRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{
			RepType: acctest.Required, Create: []string{
				`freeform_tags`, `defined_tags`, `system_tags`, `task_details`,
			}},
	}

	StackMonitoringUpdateAgentReceiverTaskResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringUpdateAgentReceiverTaskResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringUpdateAgentReceiverTaskResource_basic")
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
	managementAgentId1VariableStr := fmt.Sprintf("variable \"stack_mon_management_agent_id_resource1\" { default = \"%s\" }\n", managementAgentId1)

	resourceName := "oci_stack_monitoring_monitored_resource_task.test_update_agent_receiver_task"
	datasourceName := "data.oci_stack_monitoring_monitored_resource_tasks.test_update_agent_receiver_tasks"
	singularDatasourceName := "data.oci_stack_monitoring_monitored_resource_task.test_update_agent_receiver_task"

	currentTime, _ := time.Now().UTC().MarshalText()
	taskName := "update_agent_receiver_task_name_" + string(currentTime)
	taskNameVariableStr := fmt.Sprintf("variable \"task_name\" { default = \"%s\" }\n", taskName)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managementAgentId1VariableStr+taskNameVariableStr+StackMonitoringUpdateAgentReceiverTaskResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_agent_receiver_task", acctest.Optional, acctest.Create, StackMonitoringUpdateAgentReceiverTaskRepresentation), "stackmonitoring", "monitoredResourceTask", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + taskNameVariableStr + StackMonitoringUpdateAgentReceiverTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_agent_receiver_task", acctest.Required, acctest.Create, StackMonitoringUpdateAgentReceiverTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "UPDATE_AGENT_RECEIVER"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.handler_type", "TELEGRAF"),
				resource.TestCheckResourceAttrSet(resourceName, "task_details.0.agent_id"),
				//resource.TestCheckResourceAttr(resourceName, "task_details.0.agent_id", managementAgentId1VariableStr),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.is_enable", `true`),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + taskNameVariableStr + StackMonitoringUpdateAgentReceiverTaskResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + taskNameVariableStr + StackMonitoringUpdateAgentReceiverTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_agent_receiver_task", acctest.Optional, acctest.Create, StackMonitoringUpdateAgentReceiverTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "UPDATE_AGENT_RECEIVER"),

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
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + compartmentIdUVariableStr + taskNameVariableStr + StackMonitoringUpdateAgentReceiverTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_agent_receiver_task", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(StackMonitoringUpdateAgentReceiverTaskRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "UPDATE_AGENT_RECEIVER"),

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
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + taskNameVariableStr + StackMonitoringUpdateAgentReceiverTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_agent_receiver_task", acctest.Optional, acctest.Update, StackMonitoringUpdateAgentReceiverTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", taskName),
				resource.TestCheckResourceAttr(resourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "task_details.0.type", "UPDATE_AGENT_RECEIVER"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_tasks", "test_update_agent_receiver_tasks", acctest.Optional, acctest.Update, StackMonitoringUpdateAgentReceiverTaskDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentId1VariableStr + taskNameVariableStr + StackMonitoringUpdateAgentReceiverTaskResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_agent_receiver_task", acctest.Optional, acctest.Update, StackMonitoringUpdateAgentReceiverTaskRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource_task", "test_update_agent_receiver_task", acctest.Required, acctest.Create, StackMonitoringUpdateAgentReceiverTaskSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentId1VariableStr + taskNameVariableStr + StackMonitoringUpdateAgentReceiverTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitored_resource_task_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", taskName),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_details.0.type", "UPDATE_AGENT_RECEIVER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "work_request_ids.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + managementAgentId1VariableStr + taskNameVariableStr + StackMonitoringUpdateAgentReceiverTaskRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"defined_tags"},
			ResourceName:            resourceName,
		},
	})
}
