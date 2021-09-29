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
	"github.com/oracle/oci-go-sdk/v48/common"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v48/usageapi"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	QueryResourceConfig = QueryResourceDependencies +
		generateResourceFromRepresentationMap("oci_metering_computation_query", "test_query", Optional, Update, queryRepresentation)

	querySingularDataSourceRepresentation = map[string]interface{}{
		"query_id": Representation{repType: Required, create: `${oci_metering_computation_query.test_query.id}`},
	}

	queryDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_id}`},
		"filter":         RepresentationGroup{Required, queryDataSourceFilterRepresentation}}
	queryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_metering_computation_query.test_query.id}`}},
	}

	queryRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.tenancy_id}`},
		"query_definition": RepresentationGroup{Required, queryQueryDefinitionRepresentation},
	}
	queryQueryDefinitionRepresentation = map[string]interface{}{
		"cost_analysis_ui": RepresentationGroup{Required, queryQueryDefinitionCostAnalysisUIRepresentation},
		"display_name":     Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"report_query":     RepresentationGroup{Required, queryQueryDefinitionReportQueryRepresentation},
		"version":          Representation{repType: Required, create: `1.0`, update: `1.0`},
	}
	queryQueryDefinitionCostAnalysisUIRepresentation = map[string]interface{}{
		"graph":               Representation{repType: Optional, create: `BARS`, update: `LINES`},
		"is_cumulative_graph": Representation{repType: Optional, create: `false`, update: `true`},
	}
	queryQueryDefinitionReportQueryRepresentation = map[string]interface{}{
		"forecast":             RepresentationGroup{Optional, queryQueryDefinitionReportQueryForecastRepresentation},
		"granularity":          Representation{repType: Required, create: `DAILY`, update: `MONTHLY`},
		"tenant_id":            Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"compartment_depth":    Representation{repType: Optional, create: `1.0`, update: `2.0`},
		"filter":               Representation{repType: Optional, create: `{\"operator\":\"AND\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"compartmentNameValue\"}],\"tags\":[],\"filters\":[]}`, update: `{\"operator\":\"AND\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"compartmentNameValue2\"}],\"tags\":[],\"filters\":[]}`},
		"group_by":             Representation{repType: Optional, create: []string{`compartmentPath`}, update: []string{`compartmentName`}},
		"group_by_tag":         RepresentationGroup{Optional, queryQueryDefinitionReportQueryGroupByTagRepresentation},
		"is_aggregate_by_time": Representation{repType: Optional, create: `false`, update: `true`},
		"query_type":           Representation{repType: Optional, create: `USAGE`, update: `COST`},
		"time_usage_ended":     Representation{repType: Required, create: timeUsageEnded.Format(time.RFC3339Nano), update: timeUsageEnded.Format(time.RFC3339Nano)},
		"time_usage_started":   Representation{repType: Required, create: timeUsageStarted.Format(time.RFC3339Nano), update: timeUsageStarted.Format(time.RFC3339Nano)},
	}
	queryQueryDefinitionReportQueryForecastRepresentation = map[string]interface{}{
		"time_forecast_ended":   Representation{repType: Required, create: timeForecastEnded.Format(time.RFC3339Nano), update: timeForecastEnded.Format(time.RFC3339Nano)},
		"forecast_type":         Representation{repType: Optional, create: `BASIC`},
		"time_forecast_started": Representation{repType: Optional, create: timeUsageEnded.Format(time.RFC3339Nano), update: timeUsageEnded.Format(time.RFC3339Nano)},
	}
	queryQueryDefinitionReportQueryGroupByTagRepresentation = map[string]interface{}{
		"key":       Representation{repType: Optional, create: `key`, update: `key2`},
		"namespace": Representation{repType: Optional, create: `namespace`, update: `namespace2`},
		"value":     Representation{repType: Optional, create: `value`, update: `value2`},
	}
	timeUsageStarted  = StartOfDay(time.Now().UTC().Truncate(time.Millisecond))
	timeUsageEnded    = StartOfDay(time.Now().UTC().AddDate(0, 0, 1).Truncate(time.Millisecond))
	timeForecastEnded = StartOfDay(time.Now().UTC().AddDate(0, 0, 2).Truncate(time.Millisecond))

	QueryResourceDependencies = ""
)

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// issue-routing-tag: metering_computation/default
func TestMeteringComputationQueryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationQueryResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)

	resourceName := "oci_metering_computation_query.test_query"
	datasourceName := "data.oci_metering_computation_queries.test_queries"
	singularDatasourceName := "data.oci_metering_computation_query.test_query"

	var resId, resId2 string

	ResourceTest(t, testAccCheckMeteringComputationQueryDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + tenancyIdVariableStr + QueryResourceDependencies +
				generateResourceFromRepresentationMap("oci_metering_computation_query", "test_query", Required, Create, queryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "query_definition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.cost_analysis_ui.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.granularity", "DAILY"),
				resource.TestCheckResourceAttrSet(resourceName, "query_definition.0.report_query.0.tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.version", "1"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + tenancyIdVariableStr + QueryResourceDependencies +
				generateResourceFromRepresentationMap("oci_metering_computation_query", "test_query", Optional, Update, queryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.cost_analysis_ui.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.cost_analysis_ui.0.graph", "LINES"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.cost_analysis_ui.0.is_cumulative_graph", "true"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.forecast.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.forecast.0.forecast_type", "BASIC"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.forecast.0.time_forecast_ended", timeForecastEnded.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.forecast.0.time_forecast_started", timeUsageEnded.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.compartment_depth", "2"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.filter", "{\"operator\":\"AND\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"compartmentNameValue2\"}],\"tags\":[],\"filters\":[]}"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.granularity", "MONTHLY"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.group_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.group_by_tag.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.group_by_tag.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.group_by_tag.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.group_by_tag.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.is_aggregate_by_time", "true"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.query_type", "COST"),
				resource.TestCheckResourceAttrSet(resourceName, "query_definition.0.report_query.0.tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.time_usage_ended", timeUsageEnded.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.time_usage_started", timeUsageStarted.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.version", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
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
				generateDataSourceFromRepresentationMap("oci_metering_computation_queries", "test_queries", Optional, Update, queryDataSourceRepresentation) +
				tenancyIdVariableStr + QueryResourceDependencies +
				generateResourceFromRepresentationMap("oci_metering_computation_query", "test_query", Optional, Update, queryRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttr(datasourceName, "query_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "query_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_metering_computation_query", "test_query", Required, Create, querySingularDataSourceRepresentation) +
				tenancyIdVariableStr + QueryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "query_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.cost_analysis_ui.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.cost_analysis_ui.0.graph", "LINES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.cost_analysis_ui.0.is_cumulative_graph", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.forecast.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.forecast.0.forecast_type", "BASIC"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "query_definition.0.report_query.0.forecast.0.time_forecast_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "query_definition.0.report_query.0.forecast.0.time_forecast_started"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.compartment_depth", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.filter", "{\"operator\":\"AND\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"compartmentNameValue2\"}],\"tags\":[],\"filters\":[]}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.granularity", "MONTHLY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.group_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.group_by_tag.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.group_by_tag.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.group_by_tag.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.group_by_tag.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.is_aggregate_by_time", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.query_type", "COST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.version", "1"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + tenancyIdVariableStr + QueryResourceConfig,
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

func testAccCheckMeteringComputationQueryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).usageapiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_metering_computation_query" {
			noResourceFound = false
			request := oci_metering_computation.GetQueryRequest{}

			tmp := rs.Primary.ID
			request.QueryId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "metering_computation")

			_, err := client.GetQuery(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !inSweeperExcludeList("MeteringComputationQuery") {
		resource.AddTestSweepers("MeteringComputationQuery", &resource.Sweeper{
			Name:         "MeteringComputationQuery",
			Dependencies: DependencyGraph["query"],
			F:            sweepMeteringComputationQueryResource,
		})
	}
}

func sweepMeteringComputationQueryResource(compartment string) error {
	usageapiClient := GetTestClients(&schema.ResourceData{}).usageapiClient()
	queryIds, err := getQueryIds(compartment)
	if err != nil {
		return err
	}
	for _, queryId := range queryIds {
		if ok := SweeperDefaultResourceId[queryId]; !ok {
			deleteQueryRequest := oci_metering_computation.DeleteQueryRequest{}

			deleteQueryRequest.QueryId = &queryId

			deleteQueryRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "metering_computation")
			_, error := usageapiClient.DeleteQuery(context.Background(), deleteQueryRequest)
			if error != nil {
				fmt.Printf("Error deleting Query %s %s, It is possible that the resource is already deleted. Please verify manually \n", queryId, error)
				continue
			}
		}
	}
	return nil
}

func getQueryIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "QueryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	usageapiClient := GetTestClients(&schema.ResourceData{}).usageapiClient()

	listQueriesRequest := oci_metering_computation.ListQueriesRequest{}
	listQueriesRequest.CompartmentId = &compartmentId
	listQueriesResponse, err := usageapiClient.ListQueries(context.Background(), listQueriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Query list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, query := range listQueriesResponse.Items {
		id := *query.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "QueryId", id)
	}
	return resourceIds, nil
}
