// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_monitoring "github.com/oracle/oci-go-sdk/v65/monitoring"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	MonitoringAlarmRequiredOnlyResource = MonitoringAlarmResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Required, acctest.Create, MonitoringAlarmRepresentation)

	MonitoringAlarmResourceConfig = MonitoringAlarmResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Optional, acctest.Update, MonitoringAlarmRepresentation)

	MonitoringAlarmSingularDataSourceRepresentation = map[string]interface{}{
		"alarm_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_monitoring_alarm.test_alarm.id}`},
	}

	MonitoringAlarmDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `High CPU Utilization`, Update: `displayName2`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: MonitoringAlarmDataSourceFilterRepresentation}}
	MonitoringAlarmDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_monitoring_alarm.test_alarm.id}`}},
	}

	MonitoringAlarmRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"destinations":          acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ons_notification_topic.test_notification_topic.id}`}, Update: []string{`${oci_ons_notification_topic.test_notification_topic2.id}`}},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `High CPU Utilization`, Update: `displayName2`},
		"is_enabled":            acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"metric_compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":             acctest.Representation{RepType: acctest.Required, Create: `oci_computeagent`, Update: `oci_lbaas`},
		"query":                 acctest.Representation{RepType: acctest.Required, Create: `CpuUtilization[10m].percentile(0.9) < 85`, Update: `AcceptedConnections[10m].count() <= 0`},
		"severity":              acctest.Representation{RepType: acctest.Required, Create: `WARNING`, Update: `INFO`},
		"body":                  acctest.Representation{RepType: acctest.Optional, Create: `CPU utilization has reached high values.`, Update: `body2`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_notifications_per_metric_dimension_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"message_format":                   acctest.Representation{RepType: acctest.Optional, Create: `ONS_OPTIMIZED`, Update: `PRETTY_JSON`},
		"metric_compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"pending_duration":                 acctest.Representation{RepType: acctest.Optional, Create: `PT5M`, Update: `PT10M`},
		"repeat_notification_duration":     acctest.Representation{RepType: acctest.Optional, Create: `PT2H`, Update: `PT10M`},
		"resolution":                       acctest.Representation{RepType: acctest.Optional, Create: `1m`},
		"resource_group":                   acctest.Representation{RepType: acctest.Optional, Create: `resourceGroup`, Update: `resourceGroup2`},
		"suppression":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: MonitoringAlarmSuppressionRepresentation},
	}
	MonitoringAlarmSuppressionRepresentation = map[string]interface{}{
		"time_suppress_from":  acctest.Representation{RepType: acctest.Required, Create: `2126-02-01T18:00:00.001Z`, Update: `2125-12-01T18:00:00.001Z`},
		"time_suppress_until": acctest.Representation{RepType: acctest.Required, Create: `2126-02-01T19:00:00.001Z`, Update: `2125-12-01T19:00:00.001Z`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `System Maintenance`, Update: `description2`},
	}

	MonitoringAlarmResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, getONSTopicRepresentationCopyWithRandomNameOrHttpReplayValue(10, "talarm1")) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic2", acctest.Required, acctest.Create, getONSTopicRepresentationCopyWithRandomNameOrHttpReplayValue(10, "talarm2"))
)

func getONSTopicRepresentationCopyWithRandomNameOrHttpReplayValue(length int, httpReplayValue string) map[string]interface{} {
	return acctest.RepresentationCopyWithNewProperties(OnsNotificationTopicRepresentation, map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: utils.RandomStringOrHttpReplayValue(length, utils.CharsetWithoutDigits, httpReplayValue)},
	})
}

// issue-routing-tag: monitoring/default
func TestMonitoringAlarmResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringAlarmResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_monitoring_alarm.test_alarm"
	datasourceName := "data.oci_monitoring_alarms.test_alarms"
	singularDatasourceName := "data.oci_monitoring_alarm.test_alarm"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MonitoringAlarmResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Optional, acctest.Create, MonitoringAlarmRepresentation), "monitoring", "alarm", t)

	acctest.ResourceTest(t, testAccCheckMonitoringAlarmDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MonitoringAlarmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Required, acctest.Create, MonitoringAlarmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "High CPU Utilization"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_notifications_per_metric_dimension_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "metric_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "oci_computeagent"),
				resource.TestCheckResourceAttr(resourceName, "query", "CpuUtilization[10m].percentile(0.9) < 85"),
				resource.TestCheckResourceAttr(resourceName, "severity", "WARNING"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MonitoringAlarmResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MonitoringAlarmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Optional, acctest.Create, MonitoringAlarmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "body", "CPU utilization has reached high values."),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "High CPU Utilization"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_notifications_per_metric_dimension_enabled", "true"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MonitoringAlarmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(MonitoringAlarmRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "body", "CPU utilization has reached high values."),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "High CPU Utilization"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_notifications_per_metric_dimension_enabled", "true"),
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
			Config: config + compartmentIdVariableStr + MonitoringAlarmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Optional,
					acctest.Update, MonitoringAlarmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "body", "body2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_notifications_per_metric_dimension_enabled", "true"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_monitoring_alarms", "test_alarms", acctest.Optional, acctest.Update, MonitoringAlarmDataSourceRepresentation) +
				compartmentIdVariableStr + MonitoringAlarmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Optional, acctest.Update, MonitoringAlarmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(datasourceName, "alarms.0.is_notifications_per_metric_dimension_enabled", "true"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Required, acctest.Create, MonitoringAlarmSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MonitoringAlarmResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alarm_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metric_compartment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "body", "body2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "destinations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_notifications_per_metric_dimension_enabled", "true"),
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
		// verify resource import
		{
			Config:                  config + MonitoringAlarmRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMonitoringAlarmDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_monitoring_alarm" {
			noResourceFound = false
			request := oci_monitoring.GetAlarmRequest{}

			tmp := rs.Primary.ID
			request.AlarmId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "monitoring")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("MonitoringAlarm") {
		resource.AddTestSweepers("MonitoringAlarm", &resource.Sweeper{
			Name:         "MonitoringAlarm",
			Dependencies: acctest.DependencyGraph["alarm"],
			F:            sweepMonitoringAlarmResource,
		})
	}
}

func sweepMonitoringAlarmResource(compartment string) error {
	monitoringClient := acctest.GetTestClients(&schema.ResourceData{}).MonitoringClient()
	alarmIds, err := getMonitoringAlarmIds(compartment)
	if err != nil {
		return err
	}
	for _, alarmId := range alarmIds {
		if ok := acctest.SweeperDefaultResourceId[alarmId]; !ok {
			deleteAlarmRequest := oci_monitoring.DeleteAlarmRequest{}

			deleteAlarmRequest.AlarmId = &alarmId

			deleteAlarmRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "monitoring")
			_, error := monitoringClient.DeleteAlarm(context.Background(), deleteAlarmRequest)
			if error != nil {
				fmt.Printf("Error deleting Alarm %s %s, It is possible that the resource is already deleted. Please verify manually \n", alarmId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &alarmId, MonitoringAlarmSweepWaitCondition, time.Duration(3*time.Minute),
				MonitoringAlarmSweepResponseFetchOperation, "monitoring", true)
		}
	}
	return nil
}

func getMonitoringAlarmIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AlarmId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	monitoringClient := acctest.GetTestClients(&schema.ResourceData{}).MonitoringClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AlarmId", id)
	}
	return resourceIds, nil
}

func MonitoringAlarmSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if alarmResponse, ok := response.Response.(oci_monitoring.GetAlarmResponse); ok {
		return alarmResponse.LifecycleState != oci_monitoring.AlarmLifecycleStateDeleted
	}
	return false
}

func MonitoringAlarmSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MonitoringClient().GetAlarm(context.Background(), oci_monitoring.GetAlarmRequest{
		AlarmId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
