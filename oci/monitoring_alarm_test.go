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
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_monitoring "github.com/oracle/oci-go-sdk/v54/monitoring"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AlarmRequiredOnlyResource = AlarmResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Required, Create, alarmRepresentation)

	AlarmResourceConfig = AlarmResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Optional, Update, alarmRepresentation)

	alarmSingularDataSourceRepresentation = map[string]interface{}{
		"alarm_id": Representation{RepType: Required, Create: `${oci_monitoring_alarm.test_alarm.id}`},
	}

	alarmDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{RepType: Optional, Create: `false`},
		"display_name":              Representation{RepType: Optional, Create: `High CPU Utilization`, Update: `displayName2`},
		"state":                     Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":                    RepresentationGroup{Required, alarmDataSourceFilterRepresentation}}
	alarmDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_monitoring_alarm.test_alarm.id}`}},
	}

	alarmRepresentation = map[string]interface{}{
		"compartment_id":                   Representation{RepType: Required, Create: `${var.compartment_id}`},
		"destinations":                     Representation{RepType: Required, Create: []string{`${oci_ons_notification_topic.test_notification_topic.id}`}, Update: []string{`${oci_ons_notification_topic.test_notification_topic2.id}`}},
		"display_name":                     Representation{RepType: Required, Create: `High CPU Utilization`, Update: `displayName2`},
		"is_enabled":                       Representation{RepType: Required, Create: `false`, Update: `true`},
		"metric_compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"namespace":                        Representation{RepType: Required, Create: `oci_computeagent`, Update: `oci_lbaas`},
		"query":                            Representation{RepType: Required, Create: `CpuUtilization[10m].percentile(0.9) < 85`, Update: `AcceptedConnections[10m].count() <= 0`},
		"severity":                         Representation{RepType: Required, Create: `WARNING`, Update: `INFO`},
		"body":                             Representation{RepType: Optional, Create: `CPU utilization has reached high values.`, Update: `body2`},
		"defined_tags":                     Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                    Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"message_format":                   Representation{RepType: Optional, Create: `ONS_OPTIMIZED`, Update: `PRETTY_JSON`},
		"metric_compartment_id_in_subtree": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"pending_duration":                 Representation{RepType: Optional, Create: `PT5M`, Update: `PT10M`},
		"repeat_notification_duration":     Representation{RepType: Optional, Create: `PT2H`, Update: `PT10M`},
		"resolution":                       Representation{RepType: Optional, Create: `1m`},
		"resource_group":                   Representation{RepType: Optional, Create: `resourceGroup`, Update: `resourceGroup2`},
		"suppression":                      RepresentationGroup{Optional, alarmSuppressionRepresentation},
	}
	alarmSuppressionRepresentation = map[string]interface{}{
		"time_suppress_from":  Representation{RepType: Required, Create: `2126-02-01T18:00:00.001Z`, Update: `2125-12-01T18:00:00.001Z`},
		"time_suppress_until": Representation{RepType: Required, Create: `2126-02-01T19:00:00.001Z`, Update: `2125-12-01T19:00:00.001Z`},
		"description":         Representation{RepType: Optional, Create: `System Maintenance`, Update: `description2`},
	}

	AlarmResourceDependencies = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, getTopicRepresentationCopyWithRandomNameOrHttpReplayValue(10, charsetWithoutDigits, "talarm1")) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic2", Required, Create, getTopicRepresentationCopyWithRandomNameOrHttpReplayValue(10, charsetWithoutDigits, "talarm2"))
)

// issue-routing-tag: monitoring/default
func TestMonitoringAlarmResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringAlarmResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_monitoring_alarm.test_alarm"
	datasourceName := "data.oci_monitoring_alarms.test_alarms"
	singularDatasourceName := "data.oci_monitoring_alarm.test_alarm"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+AlarmResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Optional, Create, alarmRepresentation), "monitoring", "alarm", t)

	ResourceTest(t, testAccCheckMonitoringAlarmDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AlarmResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Required, Create, alarmRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "High CPU Utilization"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "metric_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "oci_computeagent"),
				resource.TestCheckResourceAttr(resourceName, "query", "CpuUtilization[10m].percentile(0.9) < 85"),
				resource.TestCheckResourceAttr(resourceName, "severity", "WARNING"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AlarmResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AlarmResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Optional, Create, alarmRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "body", "CPU utilization has reached high values."),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "High CPU Utilization"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "message_format", "ONS_OPTIMIZED"),
				resource.TestCheckResourceAttrSet(resourceName, "metric_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "metric_compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "oci_computeagent"),
				resource.TestCheckResourceAttr(resourceName, "pending_duration", "PT5M"),
				resource.TestCheckResourceAttr(resourceName, "query", "CpuUtilization[10m].percentile(0.9) < 85"),
				resource.TestCheckResourceAttr(resourceName, "repeat_notification_duration", "PT2H"),
				resource.TestCheckResourceAttr(resourceName, "resolution", "1m"),
				resource.TestCheckResourceAttr(resourceName, "resource_group", "resourceGroup"),
				resource.TestCheckResourceAttr(resourceName, "severity", "WARNING"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "suppression.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "suppression.0.description", "System Maintenance"),
				resource.TestCheckResourceAttr(resourceName, "suppression.0.time_suppress_from", "2126-02-01T18:00:00.001Z"),
				resource.TestCheckResourceAttr(resourceName, "suppression.0.time_suppress_until", "2126-02-01T19:00:00.001Z"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AlarmResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Optional, Create,
					RepresentationCopyWithNewProperties(alarmRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "body", "CPU utilization has reached high values."),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "High CPU Utilization"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "message_format", "ONS_OPTIMIZED"),
				resource.TestCheckResourceAttrSet(resourceName, "metric_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "metric_compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "oci_computeagent"),
				resource.TestCheckResourceAttr(resourceName, "pending_duration", "PT5M"),
				resource.TestCheckResourceAttr(resourceName, "query", "CpuUtilization[10m].percentile(0.9) < 85"),
				resource.TestCheckResourceAttr(resourceName, "repeat_notification_duration", "PT2H"),
				resource.TestCheckResourceAttr(resourceName, "resolution", "1m"),
				resource.TestCheckResourceAttr(resourceName, "resource_group", "resourceGroup"),
				resource.TestCheckResourceAttr(resourceName, "severity", "WARNING"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "suppression.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "suppression.0.description", "System Maintenance"),
				resource.TestCheckResourceAttr(resourceName, "suppression.0.time_suppress_from", "2126-02-01T18:00:00.001Z"),
				resource.TestCheckResourceAttr(resourceName, "suppression.0.time_suppress_until", "2126-02-01T19:00:00.001Z"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AlarmResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Optional, Update, alarmRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "body", "body2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "message_format", "PRETTY_JSON"),
				resource.TestCheckResourceAttrSet(resourceName, "metric_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "metric_compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "oci_lbaas"),
				resource.TestCheckResourceAttr(resourceName, "pending_duration", "PT10M"),
				resource.TestCheckResourceAttr(resourceName, "query", "AcceptedConnections[10m].count() <= 0"),
				resource.TestCheckResourceAttr(resourceName, "repeat_notification_duration", "PT10M"),
				resource.TestCheckResourceAttr(resourceName, "resolution", "1m"),
				resource.TestCheckResourceAttr(resourceName, "resource_group", "resourceGroup2"),
				resource.TestCheckResourceAttr(resourceName, "severity", "INFO"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "suppression.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "suppression.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "suppression.0.time_suppress_from", "2125-12-01T18:00:00.001Z"),
				resource.TestCheckResourceAttr(resourceName, "suppression.0.time_suppress_until", "2125-12-01T19:00:00.001Z"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_monitoring_alarms", "test_alarms", Optional, Update, alarmDataSourceRepresentation) +
				compartmentIdVariableStr + AlarmResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Optional, Update, alarmRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "alarms.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.destinations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "alarms.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "alarms.0.metric_compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.namespace", "oci_lbaas"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.query", "AcceptedConnections[10m].count() <= 0"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.severity", "INFO"),
				resource.TestCheckResourceAttrSet(datasourceName, "alarms.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.suppression.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.suppression.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.suppression.0.time_suppress_from", "2125-12-01T18:00:00.001Z"),
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.suppression.0.time_suppress_until", "2125-12-01T19:00:00.001Z"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", Required, Create, alarmSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AlarmResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alarm_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metric_compartment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "body", "body2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "message_format", "PRETTY_JSON"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", "oci_lbaas"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pending_duration", "PT10M"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query", "AcceptedConnections[10m].count() <= 0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "repeat_notification_duration", "PT10M"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resolution", "1m"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_group", "resourceGroup2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "severity", "INFO"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "suppression.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "suppression.0.description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "suppression.0.time_suppress_from"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "suppression.0.time_suppress_until"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AlarmResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMonitoringAlarmDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).monitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_monitoring_alarm" {
			noResourceFound = false
			request := oci_monitoring.GetAlarmRequest{}

			tmp := rs.Primary.ID
			request.AlarmId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "monitoring")

			response, err := client.GetAlarm(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_monitoring.AlarmLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("MonitoringAlarm") {
		resource.AddTestSweepers("MonitoringAlarm", &resource.Sweeper{
			Name:         "MonitoringAlarm",
			Dependencies: DependencyGraph["alarm"],
			F:            sweepMonitoringAlarmResource,
		})
	}
}

func sweepMonitoringAlarmResource(compartment string) error {
	monitoringClient := GetTestClients(&schema.ResourceData{}).monitoringClient()
	alarmIds, err := getAlarmIds(compartment)
	if err != nil {
		return err
	}
	for _, alarmId := range alarmIds {
		if ok := SweeperDefaultResourceId[alarmId]; !ok {
			deleteAlarmRequest := oci_monitoring.DeleteAlarmRequest{}

			deleteAlarmRequest.AlarmId = &alarmId

			deleteAlarmRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "monitoring")
			_, error := monitoringClient.DeleteAlarm(context.Background(), deleteAlarmRequest)
			if error != nil {
				fmt.Printf("Error deleting Alarm %s %s, It is possible that the resource is already deleted. Please verify manually \n", alarmId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &alarmId, alarmSweepWaitCondition, time.Duration(3*time.Minute),
				alarmSweepResponseFetchOperation, "monitoring", true)
		}
	}
	return nil
}

func getAlarmIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "AlarmId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	monitoringClient := GetTestClients(&schema.ResourceData{}).monitoringClient()

	listAlarmsRequest := oci_monitoring.ListAlarmsRequest{}
	listAlarmsRequest.CompartmentId = &compartmentId
	listAlarmsRequest.LifecycleState = oci_monitoring.AlarmLifecycleStateActive
	listAlarmsResponse, err := monitoringClient.ListAlarms(context.Background(), listAlarmsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Alarm list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, alarm := range listAlarmsResponse.Items {
		id := *alarm.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "AlarmId", id)
	}
	return resourceIds, nil
}

func alarmSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if alarmResponse, ok := response.Response.(oci_monitoring.GetAlarmResponse); ok {
		return alarmResponse.LifecycleState != oci_monitoring.AlarmLifecycleStateDeleted
	}
	return false
}

func alarmSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.monitoringClient().GetAlarm(context.Background(), oci_monitoring.GetAlarmRequest{
		AlarmId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
