// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringBaselineableMetricResourceConfig = StackMonitoringBaselineableMetricResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_baselineable_metric", "test_baselineable_metric", acctest.Optional, acctest.Update, StackMonitoringBaselineableMetricRepresentation)

	StackMonitoringBaselineableMetricSingularDataSourceRepresentation = map[string]interface{}{
		"baselineable_metric_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_baselineable_metric.test_baselineable_metric.id}`},
	}

	StackMonitoringBaselineableMetricDataSourceRepresentation = map[string]interface{}{
		"baselineable_metric_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_stack_monitoring_baselineable_metric.test_baselineable_metric.id}`},
		"compartment_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"metric_namespace":       acctest.Representation{RepType: acctest.Optional, Create: `metricNamespace`},
		"name":                   acctest.Representation{RepType: acctest.Optional, Create: `CPU`, Update: `name2`},
		"resource_group":         acctest.Representation{RepType: acctest.Optional, Create: `oracle_database`, Update: `resourceGroup2`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringBaselineableMetricDataSourceFilterRepresentation}}
	StackMonitoringBaselineableMetricDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_baselineable_metric.test_baselineable_metric.id}`}},
	}

	StackMonitoringBaselineableMetricRepresentation = map[string]interface{}{
		"column":         acctest.Representation{RepType: acctest.Required, Create: `CPUUtilization`, Update: `column2`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `CPU`, Update: `name2`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `oracle_appmgmt`, Update: `namespace2`},
		"resource_group": acctest.Representation{RepType: acctest.Required, Create: `oracle_database`, Update: `resourceGroup2`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSensitiveBaselineableMetricRepresentation},
	}
	//Get API does not return sensitive data, it returns null
	ignoreSensitiveBaselineableMetricRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`}},
	}

	StackMonitoringBaselineableMetricResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringBaselineableMetricResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringBaselineableMetricResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_stack_monitoring_baselineable_metric.test_baselineable_metric"
	datasourceName := "data.oci_stack_monitoring_baselineable_metrics.test_baselineable_metrics"
	singularDatasourceName := "data.oci_stack_monitoring_baselineable_metric.test_baselineable_metric"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StackMonitoringBaselineableMetricResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_baselineable_metric", "test_baselineable_metric", acctest.Required, acctest.Create, StackMonitoringBaselineableMetricRepresentation), "stackmonitoring", "baselineableMetric", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringBaselineableMetricDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringBaselineableMetricResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_baselineable_metric", "test_baselineable_metric", acctest.Required, acctest.Create, StackMonitoringBaselineableMetricRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column", "CPUUtilization"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "CPU"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "oracle_appmgmt"),
				resource.TestCheckResourceAttr(resourceName, "resource_group", "oracle_database"),

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
			Config: config + compartmentIdVariableStr + StackMonitoringBaselineableMetricResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_baselineable_metric", "test_baselineable_metric", acctest.Optional, acctest.Update, StackMonitoringBaselineableMetricRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "column", "column2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_out_of_box"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "resource_group", "resourceGroup2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_baselineable_metrics", "test_baselineable_metrics", acctest.Optional, acctest.Update, StackMonitoringBaselineableMetricDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringBaselineableMetricResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_baselineable_metric", "test_baselineable_metric", acctest.Optional, acctest.Update, StackMonitoringBaselineableMetricRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "baselineable_metric_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "metric_namespace", "metricNamespace"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(datasourceName, "resource_group", "resourceGroup2"),

				resource.TestCheckResourceAttr(datasourceName, "baselineable_metric_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "baselineable_metric_summary_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_baselineable_metric", "test_baselineable_metric", acctest.Required, acctest.Create, StackMonitoringBaselineableMetricSingularDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringBaselineableMetricResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "baselineable_metric_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "column", "column2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_out_of_box"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_updated_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_group", "resourceGroup2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + StackMonitoringBaselineableMetricResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckStackMonitoringBaselineableMetricDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_baselineable_metric" {
			noResourceFound = false
			request := oci_stack_monitoring.GetBaselineableMetricRequest{}

			tmp := rs.Primary.ID
			request.BaselineableMetricId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetBaselineableMetric(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.BaselineableMetricLifeCycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("StackMonitoringBaselineableMetric") {
		resource.AddTestSweepers("StackMonitoringBaselineableMetric", &resource.Sweeper{
			Name:         "StackMonitoringBaselineableMetric",
			Dependencies: acctest.DependencyGraph["baselineableMetric"],
			F:            sweepStackMonitoringBaselineableMetricResource,
		})
	}
}

func sweepStackMonitoringBaselineableMetricResource(compartment string) error {
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()
	baselineableMetricIds, err := getStackMonitoringBaselineableMetricIds(compartment)
	if err != nil {
		return err
	}
	for _, baselineableMetricId := range baselineableMetricIds {
		if ok := acctest.SweeperDefaultResourceId[baselineableMetricId]; !ok {
			deleteBaselineableMetricRequest := oci_stack_monitoring.DeleteBaselineableMetricRequest{}

			deleteBaselineableMetricRequest.BaselineableMetricId = &baselineableMetricId

			deleteBaselineableMetricRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")
			_, error := stackMonitoringClient.DeleteBaselineableMetric(context.Background(), deleteBaselineableMetricRequest)
			if error != nil {
				fmt.Printf("Error deleting BaselineableMetric %s %s, It is possible that the resource is already deleted. Please verify manually \n", baselineableMetricId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &baselineableMetricId, StackMonitoringBaselineableMetricSweepWaitCondition, time.Duration(3*time.Minute),
				StackMonitoringBaselineableMetricSweepResponseFetchOperation, "stack_monitoring", true)
		}
	}
	return nil
}

func getStackMonitoringBaselineableMetricIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BaselineableMetricId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()

	listBaselineableMetricsRequest := oci_stack_monitoring.ListBaselineableMetricsRequest{}
	listBaselineableMetricsRequest.CompartmentId = &compartmentId
	listBaselineableMetricsResponse, err := stackMonitoringClient.ListBaselineableMetrics(context.Background(), listBaselineableMetricsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BaselineableMetric list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, baselineableMetric := range listBaselineableMetricsResponse.Items {
		id := *baselineableMetric.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BaselineableMetricId", id)
	}
	return resourceIds, nil
}

func StackMonitoringBaselineableMetricSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if baselineableMetricResponse, ok := response.Response.(oci_stack_monitoring.GetBaselineableMetricResponse); ok {
		return baselineableMetricResponse.LifecycleState != oci_stack_monitoring.BaselineableMetricLifeCycleStatesDeleted
	}
	return false
}

func StackMonitoringBaselineableMetricSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackMonitoringClient().GetBaselineableMetric(context.Background(), oci_stack_monitoring.GetBaselineableMetricRequest{
		BaselineableMetricId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
