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
	oci_demand_signal "github.com/oracle/oci-go-sdk/v65/demandsignal"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DemandSignalOccMetricAlarmRequiredOnlyResource = DemandSignalOccMetricAlarmResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_metric_alarm", "test_occ_metric_alarm", acctest.Required, acctest.Create, DemandSignalOccMetricAlarmRepresentation)

	DemandSignalOccMetricAlarmResourceConfig = DemandSignalOccMetricAlarmResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_metric_alarm", "test_occ_metric_alarm", acctest.Optional, acctest.Update, DemandSignalOccMetricAlarmRepresentation)

	DemandSignalOccMetricAlarmSingularDataSourceRepresentation = map[string]interface{}{
		"occ_metric_alarm_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_demand_signal_occ_metric_alarm.test_occ_metric_alarm.id}`},
	}

	DemandSignalOccMetricAlarmDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"is_active":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DemandSignalOccMetricAlarmDataSourceFilterRepresentation}}
	DemandSignalOccMetricAlarmDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_demand_signal_occ_metric_alarm.test_occ_metric_alarm.id}`}},
	}

	DemandSignalOccMetricAlarmRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"frequency":              acctest.Representation{RepType: acctest.Required, Create: `WEEKLY`, Update: `WEEKLY`},
		"is_active":              acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"resource_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: DemandSignalOccMetricAlarmResourceConfigurationRepresentation},
		"threshold":              acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		//"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		//"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		//"state":                  acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`, Update: `UPDATING`},
		"subscribers":    acctest.Representation{RepType: acctest.Required, Create: []string{`subscribers`}, Update: []string{`subscribers2`}},
		"threshold_type": acctest.Representation{RepType: acctest.Optional, Create: `PERCENTAGE`, Update: `PERCENTAGE`},
	}
	DemandSignalOccMetricAlarmResourceConfigurationRepresentation = map[string]interface{}{
		"resource":              acctest.Representation{RepType: acctest.Required, Create: `COMPUTE`},
		"usage_type":            acctest.Representation{RepType: acctest.Required, Create: `usageType`},
		"compute_hw_generation": acctest.Representation{RepType: acctest.Optional, Create: `computeHwGeneration`},
		"shape":                 acctest.Representation{RepType: acctest.Required, Create: `shape`},
	}

	DemandSignalOccMetricAlarmResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: demand_signal/default
func TestDemandSignalOccMetricAlarmResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDemandSignalOccMetricAlarmResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_demand_signal_occ_metric_alarm.test_occ_metric_alarm"
	datasourceName := "data.oci_demand_signal_occ_metric_alarms.test_occ_metric_alarms"
	singularDatasourceName := "data.oci_demand_signal_occ_metric_alarm.test_occ_metric_alarm"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DemandSignalOccMetricAlarmResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_metric_alarm", "test_occ_metric_alarm", acctest.Optional, acctest.Create, DemandSignalOccMetricAlarmRepresentation), "demandsignal", "occMetricAlarm", t)

	acctest.ResourceTest(t, testAccCheckDemandSignalOccMetricAlarmDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DemandSignalOccMetricAlarmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_metric_alarm", "test_occ_metric_alarm", acctest.Required, acctest.Create, DemandSignalOccMetricAlarmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "frequency", "WEEKLY"),
				resource.TestCheckResourceAttr(resourceName, "is_active", "false"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.node_type", "nodeType"),
				//resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.occ_metric_alarm_provider", "occMetricAlarmProvider"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.resource", "COMPUTE"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.shape", "shape"),
				//resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.storage_type", "storageType"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.usage_type", "usageType"),
				resource.TestCheckResourceAttr(resourceName, "threshold", "10"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DemandSignalOccMetricAlarmResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DemandSignalOccMetricAlarmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_metric_alarm", "test_occ_metric_alarm", acctest.Optional, acctest.Create, DemandSignalOccMetricAlarmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_active", "false"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.compute_hw_generation", "computeHwGeneration"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.resource", "COMPUTE"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.shape", "shape"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.usage_type", "usageType"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "subscribers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "threshold_type", "PERCENTAGE"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DemandSignalOccMetricAlarmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_metric_alarm", "test_occ_metric_alarm", acctest.Optional, acctest.Update, DemandSignalOccMetricAlarmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_active", "true"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.compute_hw_generation", "computeHwGeneration"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.resource", "COMPUTE"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.shape", "shape"),
				resource.TestCheckResourceAttr(resourceName, "resource_configuration.0.usage_type", "usageType"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "subscribers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "threshold_type", "PERCENTAGE"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_demand_signal_occ_metric_alarms", "test_occ_metric_alarms", acctest.Optional, acctest.Update, DemandSignalOccMetricAlarmDataSourceRepresentation) +
				compartmentIdVariableStr + DemandSignalOccMetricAlarmResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_metric_alarm", "test_occ_metric_alarm", acctest.Optional, acctest.Update, DemandSignalOccMetricAlarmRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_active", "true"),

				resource.TestCheckResourceAttr(datasourceName, "occ_metric_alarm_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "occ_metric_alarm_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_demand_signal_occ_metric_alarm", "test_occ_metric_alarm", acctest.Required, acctest.Create, DemandSignalOccMetricAlarmSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DemandSignalOccMetricAlarmResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "occ_metric_alarm_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "frequency", "WEEKLY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_active", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_configuration.0.compute_hw_generation", "computeHwGeneration"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_configuration.0.resource", "COMPUTE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_configuration.0.shape", "shape"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_configuration.0.usage_type", "usageType"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subscribers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "threshold_type", "PERCENTAGE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DemandSignalOccMetricAlarmRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDemandSignalOccMetricAlarmDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OccMetricAlarmClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_demand_signal_occ_metric_alarm" {
			noResourceFound = false
			request := oci_demand_signal.GetOccMetricAlarmRequest{}

			tmp := rs.Primary.ID
			request.OccMetricAlarmId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "demand_signal")

			response, err := client.GetOccMetricAlarm(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_demand_signal.OccMetricAlarmLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DemandSignalOccMetricAlarm") {
		resource.AddTestSweepers("DemandSignalOccMetricAlarm", &resource.Sweeper{
			Name:         "DemandSignalOccMetricAlarm",
			Dependencies: acctest.DependencyGraph["occMetricAlarm"],
			F:            sweepDemandSignalOccMetricAlarmResource,
		})
	}
}

func sweepDemandSignalOccMetricAlarmResource(compartment string) error {
	occMetricAlarmClient := acctest.GetTestClients(&schema.ResourceData{}).OccMetricAlarmClient()
	occMetricAlarmIds, err := getDemandSignalOccMetricAlarmIds(compartment)
	if err != nil {
		return err
	}
	for _, occMetricAlarmId := range occMetricAlarmIds {
		if ok := acctest.SweeperDefaultResourceId[occMetricAlarmId]; !ok {
			deleteOccMetricAlarmRequest := oci_demand_signal.DeleteOccMetricAlarmRequest{}

			deleteOccMetricAlarmRequest.OccMetricAlarmId = &occMetricAlarmId

			deleteOccMetricAlarmRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "demand_signal")
			_, error := occMetricAlarmClient.DeleteOccMetricAlarm(context.Background(), deleteOccMetricAlarmRequest)
			if error != nil {
				fmt.Printf("Error deleting OccMetricAlarm %s %s, It is possible that the resource is already deleted. Please verify manually \n", occMetricAlarmId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &occMetricAlarmId, DemandSignalOccMetricAlarmSweepWaitCondition, time.Duration(3*time.Minute),
				DemandSignalOccMetricAlarmSweepResponseFetchOperation, "demand_signal", true)
		}
	}
	return nil
}

func getDemandSignalOccMetricAlarmIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OccMetricAlarmId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	occMetricAlarmClient := acctest.GetTestClients(&schema.ResourceData{}).OccMetricAlarmClient()

	listOccMetricAlarmsRequest := oci_demand_signal.ListOccMetricAlarmsRequest{}
	listOccMetricAlarmsRequest.CompartmentId = &compartmentId
	listOccMetricAlarmsResponse, err := occMetricAlarmClient.ListOccMetricAlarms(context.Background(), listOccMetricAlarmsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OccMetricAlarm list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, occMetricAlarm := range listOccMetricAlarmsResponse.Items {
		id := *occMetricAlarm.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OccMetricAlarmId", id)
	}
	return resourceIds, nil
}

func DemandSignalOccMetricAlarmSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if occMetricAlarmResponse, ok := response.Response.(oci_demand_signal.GetOccMetricAlarmResponse); ok {
		return occMetricAlarmResponse.LifecycleState != oci_demand_signal.OccMetricAlarmLifecycleStateDeleted
	}
	return false
}

func DemandSignalOccMetricAlarmSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OccMetricAlarmClient().GetOccMetricAlarm(context.Background(), oci_demand_signal.GetOccMetricAlarmRequest{
		OccMetricAlarmId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
