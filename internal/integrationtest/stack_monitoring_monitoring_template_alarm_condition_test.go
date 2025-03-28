// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringMonitoringTemplateAlarmConditionRequiredOnlyResource = StackMonitoringMonitoringTemplateAlarmConditionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_alarm_condition", "test_monitoring_template_alarm_condition", acctest.Required, acctest.Create, StackMonitoringMonitoringTemplateAlarmConditionRepresentation)

	StackMonitoringMonitoringTemplateAlarmConditionResourceConfig = StackMonitoringMonitoringTemplateAlarmConditionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_alarm_condition", "test_monitoring_template_alarm_condition", acctest.Optional, acctest.Update, StackMonitoringMonitoringTemplateAlarmConditionRepresentation)

	StackMonitoringMonitoringTemplateAlarmConditionSingularDataSourceRepresentation = map[string]interface{}{
		"alarm_condition_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitoring_template_alarm_condition.test_monitoring_template_alarm_condition.id}`},
		"monitoring_template_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitoring_template_alarm_condition.test_monitoring_template_alarm_condition.monitoring_template_id}`},
	}

	monitoringTemplateId = utils.GetEnvSettingWithBlankDefault("monitoring_template_ocid")

	StackMonitoringMonitoringTemplateAlarmConditionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"alarm_condition_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitoring_template_alarm_condition.test_monitoring_template_alarm_condition.id}`},
		"monitoring_template_id": acctest.Representation{RepType: acctest.Required, Create: monitoringTemplateId},
		"metric_name":            acctest.Representation{RepType: acctest.Required, Create: []string{`CpuUtilization`}},
		"resource_types":         acctest.Representation{RepType: acctest.Optional, Create: []string{`resourceTypes`}},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"status":                 acctest.Representation{RepType: acctest.Optional, Create: `NOT_APPLIED`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMonitoringTemplateAlarmConditionDataSourceFilterRepresentation}}

	StackMonitoringMonitoringTemplateAlarmConditionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_monitoring_template_alarm_condition.test_monitoring_template_alarm_condition.id}`}},
	}

	StackMonitoringMonitoringTemplateAlarmConditionRepresentation = map[string]interface{}{
		"condition_type":         acctest.Representation{RepType: acctest.Required, Create: `FIXED`, Update: `AVAILABILITY`},
		"conditions":             acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMonitoringTemplateAlarmConditionConditionsRepresentation},
		"metric_name":            acctest.Representation{RepType: acctest.Required, Create: `GarbageCollectionThroughput`},
		"monitoring_template_id": acctest.Representation{RepType: acctest.Required, Create: monitoringTemplateId},
		"namespace":              acctest.Representation{RepType: acctest.Required, Create: `oracle_appmgmt`, Update: `oracle_appmgmt`},
		"resource_type":          acctest.Representation{RepType: acctest.Required, Create: `ocid1.stackmonitoringresourcetype.apache_tomcat`, Update: `ocid1.stackmonitoringresourcetype.apache_tomcat`},
		"lifecycle":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreAlarmConditionSensitiveDataRepresentation},
	}

	StackMonitoringMonitoringTemplateAlarmConditionConditionsRepresentation = map[string]interface{}{
		"query":              acctest.Representation{RepType: acctest.Required, Create: `GarbageCollectionThroughput[10m].mean() > 0.3`, Update: `GarbageCollectionThroughput[1m].mean() > 0.8`},
		"severity":           acctest.Representation{RepType: acctest.Required, Create: `WARNING`, Update: `CRITICAL`},
		"body":               acctest.Representation{RepType: acctest.Optional, Create: `Garbage collection throughput exceeds the warning threshold value`, Update: `Garbage collection throughput exceeds the critical threshold value`},
		"should_append_note": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"should_append_url":  acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"trigger_delay":      acctest.Representation{RepType: acctest.Optional, Create: `PT5M`, Update: `PT10M`},
	}

	ignoreAlarmConditionSensitiveDataRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`system_tags`}},
	}

	StackMonitoringMonitoringTemplateAlarmConditionResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoringTemplateAlarmConditionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoringTemplateAlarmConditionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_stack_monitoring_monitoring_template_alarm_condition.test_monitoring_template_alarm_condition"
	datasourceName := "data.oci_stack_monitoring_monitoring_template_alarm_conditions.test_monitoring_template_alarm_conditions"
	singularDatasourceName := "data.oci_stack_monitoring_monitoring_template_alarm_condition.test_monitoring_template_alarm_condition"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StackMonitoringMonitoringTemplateAlarmConditionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_alarm_condition", "test_monitoring_template_alarm_condition", acctest.Optional, acctest.Create, StackMonitoringMonitoringTemplateAlarmConditionRepresentation), "stackmonitoring", "monitoringTemplateAlarmCondition", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringMonitoringTemplateAlarmConditionDestroy, []resource.TestStep{
		// Verify Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMonitoringTemplateAlarmConditionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_alarm_condition", "test_monitoring_template_alarm_condition", acctest.Required, acctest.Create, StackMonitoringMonitoringTemplateAlarmConditionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "condition_type", "FIXED"),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.query", "GarbageCollectionThroughput[10m].mean() > 0.3"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.severity", "WARNING"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "oracle_appmgmt"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "ocid1.stackmonitoringresourcetype.apache_tomcat"),
				resource.TestCheckResourceAttrSet(resourceName, "metric_name"),
				resource.TestCheckResourceAttrSet(resourceName, "monitoring_template_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Delete before next Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMonitoringTemplateAlarmConditionResourceDependencies,
		},

		// Verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMonitoringTemplateAlarmConditionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_alarm_condition", "test_monitoring_template_alarm_condition", acctest.Optional, acctest.Create, StackMonitoringMonitoringTemplateAlarmConditionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "condition_type", "FIXED"),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.body", "Garbage collection throughput exceeds the warning threshold value"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.query", "GarbageCollectionThroughput[10m].mean() > 0.3"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.severity", "WARNING"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.should_append_note", "true"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.should_append_url", "true"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.trigger_delay", "PT5M"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "oracle_appmgmt"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "ocid1.stackmonitoringresourcetype.apache_tomcat"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "metric_name"),
				resource.TestCheckResourceAttrSet(resourceName, "monitoring_template_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),

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

		// Verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMonitoringTemplateAlarmConditionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_alarm_condition", "test_monitoring_template_alarm_condition", acctest.Optional, acctest.Update, StackMonitoringMonitoringTemplateAlarmConditionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "condition_type", "AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.body", "Garbage collection throughput exceeds the critical threshold value"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.query", "GarbageCollectionThroughput[1m].mean() > 0.8"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.severity", "CRITICAL"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.should_append_note", "false"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.should_append_url", "false"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.trigger_delay", "PT10M"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "oracle_appmgmt"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "ocid1.stackmonitoringresourcetype.apache_tomcat"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "metric_name"),
				resource.TestCheckResourceAttrSet(resourceName, "monitoring_template_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// Verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_alarm_conditions", "test_monitoring_template_alarm_conditions", acctest.Optional, acctest.Update, StackMonitoringMonitoringTemplateAlarmConditionDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringMonitoringTemplateAlarmConditionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_alarm_condition", "test_monitoring_template_alarm_condition", acctest.Optional, acctest.Update, StackMonitoringMonitoringTemplateAlarmConditionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "metric_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "resource_types.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "status", "NOT_APPLIED"),
				resource.TestCheckResourceAttrSet(datasourceName, "monitoring_template_id"),
			),
		},

		// Verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitoring_template_alarm_condition", "test_monitoring_template_alarm_condition", acctest.Required, acctest.Create, StackMonitoringMonitoringTemplateAlarmConditionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringMonitoringTemplateAlarmConditionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.body", "Garbage collection throughput exceeds the critical threshold value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.query", "GarbageCollectionThroughput[1m].mean() > 0.8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.severity", "CRITICAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.should_append_note", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.should_append_url", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.trigger_delay", "PT10M"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", "oracle_appmgmt"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_type", "ocid1.stackmonitoringresourcetype.apache_tomcat"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alarm_condition_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitoring_template_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},

		// Verify resource import
		{
			Config:                  config + StackMonitoringMonitoringTemplateAlarmConditionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},

		// Delete dangling resource, if any
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMonitoringTemplateAlarmConditionResourceDependencies,
		},
	})
}

func testAccCheckStackMonitoringMonitoringTemplateAlarmConditionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_monitoring_template_alarm_condition" {
			noResourceFound = false
			request := oci_stack_monitoring.GetAlarmConditionRequest{}

			if value, ok := rs.Primary.Attributes["alarm_condition_id"]; ok {
				request.AlarmConditionId = &value
			}

			if value, ok := rs.Primary.Attributes["monitoring_template_id"]; ok {
				request.MonitoringTemplateId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetAlarmCondition(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.AlarmConditionLifeCycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					// Resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				// Resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			// Verify that exception is for '404 not found'.
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
	if !acctest.InSweeperExcludeList("StackMonitoringMonitoringTemplateAlarmCondition") {
		resource.AddTestSweepers("StackMonitoringMonitoringTemplateAlarmCondition", &resource.Sweeper{
			Name:         "StackMonitoringMonitoringTemplateAlarmCondition",
			Dependencies: acctest.DependencyGraph["monitoringTemplateAlarmCondition"],
			F:            sweepStackMonitoringMonitoringTemplateAlarmConditionResource,
		})
	}
}

func sweepStackMonitoringMonitoringTemplateAlarmConditionResource(compartment string) error {
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()
	monitoringTemplateAlarmConditionIds, err := getStackMonitoringMonitoringTemplateAlarmConditionIds(compartment)
	if err != nil {
		return err
	}
	for _, monitoringTemplateAlarmConditionId := range monitoringTemplateAlarmConditionIds {
		if ok := acctest.SweeperDefaultResourceId[monitoringTemplateAlarmConditionId]; !ok {
			deleteAlarmConditionRequest := oci_stack_monitoring.DeleteAlarmConditionRequest{}

			deleteAlarmConditionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")
			_, error := stackMonitoringClient.DeleteAlarmCondition(context.Background(), deleteAlarmConditionRequest)
			if error != nil {
				fmt.Printf("Error deleting MonitoringTemplateAlarmCondition %s %s, It is possible that the resource is already deleted. Please verify manually \n", monitoringTemplateAlarmConditionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &monitoringTemplateAlarmConditionId, StackMonitoringMonitoringTemplateAlarmConditionSweepWaitCondition, time.Duration(3*time.Minute),
				StackMonitoringMonitoringTemplateAlarmConditionSweepResponseFetchOperation, "stack_monitoring", true)
		}
	}
	return nil
}

func getStackMonitoringMonitoringTemplateAlarmConditionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MonitoringTemplateAlarmConditionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()

	listAlarmConditionsRequest := oci_stack_monitoring.ListAlarmConditionsRequest{}

	monitoringTemplateIds, error := getStackMonitoringMonitoringTemplateIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting monitoringTemplateId required for MonitoringTemplateAlarmCondition resource requests \n")
	}
	for _, monitoringTemplateId := range monitoringTemplateIds {
		listAlarmConditionsRequest.MonitoringTemplateId = &monitoringTemplateId

		listAlarmConditionsRequest.LifecycleState = oci_stack_monitoring.ListAlarmConditionsLifecycleStateActive
		listAlarmConditionsResponse, err := stackMonitoringClient.ListAlarmConditions(context.Background(), listAlarmConditionsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting MonitoringTemplateAlarmCondition list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, monitoringTemplateAlarmCondition := range listAlarmConditionsResponse.Items {
			id := *monitoringTemplateAlarmCondition.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MonitoringTemplateAlarmConditionId", id)
		}

	}
	return resourceIds, nil
}

func StackMonitoringMonitoringTemplateAlarmConditionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if monitoringTemplateAlarmConditionResponse, ok := response.Response.(oci_stack_monitoring.GetAlarmConditionResponse); ok {
		return monitoringTemplateAlarmConditionResponse.LifecycleState != oci_stack_monitoring.AlarmConditionLifeCycleStatesDeleted
	}
	return false
}

func StackMonitoringMonitoringTemplateAlarmConditionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackMonitoringClient().GetAlarmCondition(context.Background(), oci_stack_monitoring.GetAlarmConditionRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
