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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_monitoring "github.com/oracle/oci-go-sdk/v65/monitoring"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MonitoringAlarmSuppressionRequiredOnlyResource = MonitoringAlarmSuppressionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm_suppression", "test_alarm_suppression", acctest.Required, acctest.Create, MonitoringDimensionAlarmSuppressionRepresentation)

	MonitoringAlarmSuppressionResourceConfig = MonitoringAlarmSuppressionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm_suppression", "test_alarm_suppression", acctest.Optional, acctest.Update, MonitoringDimensionAlarmSuppressionRepresentation)

	MonitoringAlarmSuppressionSingularDataSourceRepresentation = map[string]interface{}{
		"alarm_suppression_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_monitoring_alarm_suppression.test_alarm_suppression.id}`},
	}

	MonitoringAlarmSuppressionDataSourceRepresentation = map[string]interface{}{
		"alarm_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_monitoring_alarm.test_alarm.id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `Suppression for monthly downtime of resource ABC`},
		"is_all_suppressions": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"level":               acctest.Representation{RepType: acctest.Optional, Create: `DIMENSION`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"target_type":         acctest.Representation{RepType: acctest.Optional, Create: `ALARM`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: MonitoringAlarmSuppressionDataSourceFilterRepresentation}}

	MonitoringAlarmSuppressionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_monitoring_alarm_suppression.test_alarm_suppression.id}`}},
	}

	MonitoringDimensionAlarmSuppressionRepresentation = map[string]interface{}{
		"alarm_suppression_target": acctest.RepresentationGroup{RepType: acctest.Required, Group: MonitoringAlarmSuppressionAlarmSuppressionTargetRepresentation},

		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `Suppression for monthly downtime of resource ABC`},
		"time_suppress_from":  acctest.Representation{RepType: acctest.Required, Create: `2025-03-04T05:00:00Z`},
		"time_suppress_until": acctest.Representation{RepType: acctest.Required, Create: `2025-03-31T17:00:00Z`},

		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `Suppression for monthly downtime of resource ABC, for support ticket IT-ABC`},
		"dimensions":             acctest.Representation{RepType: acctest.Required, Create: map[string]string{"resourceId": "instance.instanceId.region1.phx.exampleuniqueID"}},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"level":                  acctest.Representation{RepType: acctest.Optional, Create: `DIMENSION`},
		"suppression_conditions": acctest.RepresentationGroup{RepType: acctest.Optional, Group: MonitoringAlarmSuppressionSuppressionConditionsRepresentation},
	}
	MonitoringAlarmSuppressionAlarmSuppressionTargetRepresentation = map[string]interface{}{
		"target_type": acctest.Representation{RepType: acctest.Required, Create: `ALARM`},
		"alarm_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_monitoring_alarm.test_alarm.id}`},
	}
	MonitoringAlarmSuppressionSuppressionConditionsRepresentation = map[string]interface{}{
		"condition_type":         acctest.Representation{RepType: acctest.Required, Create: `RECURRENCE`},
		"suppression_duration":   acctest.Representation{RepType: acctest.Required, Create: `PT1H`},
		"suppression_recurrence": acctest.Representation{RepType: acctest.Required, Create: `FREQ=WEEKLY;BYDAY=MO,TU,WE,TH;BYHOUR=10;BYMINUTE=00;BYSECOND=00`},
	}

	MonitoringAlarmSuppressionResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm", "test_alarm", acctest.Required, acctest.Create, MonitoringAlarmRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: monitoring/default
func TestMonitoringAlarmSuppressionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMonitoringAlarmSuppressionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("TF_VAR_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_monitoring_alarm_suppression.test_alarm_suppression"
	datasourceName := "data.oci_monitoring_alarm_suppressions.test_alarm_suppressions"
	singularDatasourceName := "data.oci_monitoring_alarm_suppression.test_alarm_suppression"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MonitoringAlarmSuppressionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm_suppression", "test_alarm_suppression", acctest.Optional, acctest.Create, MonitoringDimensionAlarmSuppressionRepresentation), "monitoring", "alarmSuppression", t)

	acctest.ResourceTest(t, testAccCheckMonitoringAlarmSuppressionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MonitoringAlarmSuppressionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm_suppression", "test_alarm_suppression", acctest.Required, acctest.Create, MonitoringDimensionAlarmSuppressionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alarm_suppression_target.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "alarm_suppression_target.0.alarm_id"),
				resource.TestCheckResourceAttr(resourceName, "alarm_suppression_target.0.target_type", "ALARM"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Suppression for monthly downtime of resource ABC"),

				resource.TestCheckResourceAttr(resourceName, "time_suppress_from", "2025-03-04T05:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "time_suppress_until", "2025-03-31T17:00:00Z"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MonitoringAlarmSuppressionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MonitoringAlarmSuppressionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm_suppression", "test_alarm_suppression", acctest.Optional, acctest.Create, MonitoringDimensionAlarmSuppressionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alarm_suppression_target.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "alarm_suppression_target.0.alarm_id"),
				resource.TestCheckResourceAttr(resourceName, "alarm_suppression_target.0.target_type", "ALARM"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "Suppression for monthly downtime of resource ABC, for support ticket IT-ABC"),
				resource.TestCheckResourceAttr(resourceName, "dimensions.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Suppression for monthly downtime of resource ABC"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "level", "DIMENSION"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "suppression_conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "suppression_conditions.0.condition_type", "RECURRENCE"),
				resource.TestCheckResourceAttr(resourceName, "suppression_conditions.0.suppression_duration", "PT1H"),
				resource.TestCheckResourceAttr(resourceName, "suppression_conditions.0.suppression_recurrence", "FREQ=WEEKLY;BYDAY=MO,TU,WE,TH;BYHOUR=10;BYMINUTE=00;BYSECOND=00"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				resource.TestCheckResourceAttr(resourceName, "time_suppress_from", "2025-03-04T05:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "time_suppress_until", "2025-03-31T17:00:00Z"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_monitoring_alarm_suppressions", "test_alarm_suppressions", acctest.Optional, acctest.Update, MonitoringAlarmSuppressionDataSourceRepresentation) +
				compartmentIdVariableStr + MonitoringAlarmSuppressionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_monitoring_alarm_suppression", "test_alarm_suppression", acctest.Optional, acctest.Update, MonitoringDimensionAlarmSuppressionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "alarm_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Suppression for monthly downtime of resource ABC"),

				resource.TestCheckResourceAttr(datasourceName, "is_all_suppressions", "false"),
				resource.TestCheckResourceAttr(datasourceName, "level", "DIMENSION"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "target_type", "ALARM"),

				resource.TestCheckResourceAttr(datasourceName, "alarm_suppression_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "alarm_suppression_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "alarm_suppression_collection.0.items.0.suppression_conditions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "alarm_suppression_collection.0.items.0.suppression_conditions.0.condition_type", "RECURRENCE"),
				resource.TestCheckResourceAttr(datasourceName, "alarm_suppression_collection.0.items.0.suppression_conditions.0.suppression_duration", "PT1H"),
				resource.TestCheckResourceAttr(datasourceName, "alarm_suppression_collection.0.items.0.suppression_conditions.0.suppression_recurrence", "FREQ=WEEKLY;BYDAY=MO,TU,WE,TH;BYHOUR=10;BYMINUTE=00;BYSECOND=00"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_monitoring_alarm_suppression", "test_alarm_suppression", acctest.Required, acctest.Create, MonitoringAlarmSuppressionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MonitoringAlarmSuppressionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alarm_suppression_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "alarm_suppression_target.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "alarm_suppression_target.0.target_type", "ALARM"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "Suppression for monthly downtime of resource ABC, for support ticket IT-ABC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dimensions.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Suppression for monthly downtime of resource ABC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "level", "DIMENSION"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "suppression_conditions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "suppression_conditions.0.condition_type", "RECURRENCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "suppression_conditions.0.suppression_duration", "PT1H"),
				resource.TestCheckResourceAttr(singularDatasourceName, "suppression_conditions.0.suppression_recurrence", "FREQ=WEEKLY;BYDAY=MO,TU,WE,TH;BYHOUR=10;BYMINUTE=00;BYSECOND=00"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_suppress_from"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_suppress_until"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + MonitoringAlarmSuppressionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMonitoringAlarmSuppressionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_monitoring_alarm_suppression" {
			noResourceFound = false
			request := oci_monitoring.GetAlarmSuppressionRequest{}

			tmp := rs.Primary.ID
			request.AlarmSuppressionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "monitoring")

			response, err := client.GetAlarmSuppression(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_monitoring.AlarmSuppressionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MonitoringAlarmSuppression") {
		resource.AddTestSweepers("MonitoringAlarmSuppression", &resource.Sweeper{
			Name:         "MonitoringAlarmSuppression",
			Dependencies: acctest.DependencyGraph["alarmSuppression"],
			F:            sweepMonitoringAlarmSuppressionResource,
		})
	}
}

func sweepMonitoringAlarmSuppressionResource(compartment string) error {
	monitoringClient := acctest.GetTestClients(&schema.ResourceData{}).MonitoringClient()
	alarmSuppressionIds, err := getMonitoringAlarmSuppressionIds(compartment)
	if err != nil {
		return err
	}
	for _, alarmSuppressionId := range alarmSuppressionIds {
		if ok := acctest.SweeperDefaultResourceId[alarmSuppressionId]; !ok {
			deleteAlarmSuppressionRequest := oci_monitoring.DeleteAlarmSuppressionRequest{}

			deleteAlarmSuppressionRequest.AlarmSuppressionId = &alarmSuppressionId

			deleteAlarmSuppressionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "monitoring")
			_, error := monitoringClient.DeleteAlarmSuppression(context.Background(), deleteAlarmSuppressionRequest)
			if error != nil {
				fmt.Printf("Error deleting AlarmSuppression %s %s, It is possible that the resource is already deleted. Please verify manually \n", alarmSuppressionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &alarmSuppressionId, MonitoringAlarmSuppressionSweepWaitCondition, time.Duration(3*time.Minute),
				MonitoringAlarmSuppressionSweepResponseFetchOperation, "monitoring", true)
		}
	}
	return nil
}

func getMonitoringAlarmSuppressionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AlarmSuppressionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	monitoringClient := acctest.GetTestClients(&schema.ResourceData{}).MonitoringClient()

	listAlarmSuppressionsRequest := oci_monitoring.ListAlarmSuppressionsRequest{}

	listAlarmSuppressionsRequest.CompartmentId = &compartmentId
	listAlarmSuppressionsRequest.LifecycleState = oci_monitoring.AlarmSuppressionLifecycleStateActive
	listAlarmSuppressionsResponse, err := monitoringClient.ListAlarmSuppressions(context.Background(), listAlarmSuppressionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AlarmSuppression list for compartment id : %s , %s \n", compartmentId, err)

	}
	for _, alarmSuppression := range listAlarmSuppressionsResponse.Items {
		id := *alarmSuppression.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AlarmSuppressionId", id)
	}
	return resourceIds, nil
}

func MonitoringAlarmSuppressionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if alarmSuppressionResponse, ok := response.Response.(oci_monitoring.GetAlarmSuppressionResponse); ok {
		return alarmSuppressionResponse.LifecycleState != oci_monitoring.AlarmSuppressionLifecycleStateDeleted
	}
	return false
}

func MonitoringAlarmSuppressionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MonitoringClient().GetAlarmSuppression(context.Background(), oci_monitoring.GetAlarmSuppressionRequest{
		AlarmSuppressionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
