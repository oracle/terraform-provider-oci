// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	QueryRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_query", "test_query", acctest.Required, acctest.Create, MeteringComputationQueryRepresentation)

	MeteringComputationQueryResourceConfig = MeteringComputationQueryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_query", "test_query", acctest.Optional, acctest.Update, MeteringComputationQueryRepresentation)

	MeteringComputationMeteringComputationQuerySingularDataSourceRepresentation = map[string]interface{}{
		"query_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_metering_computation_query.test_query.id}`},
	}

	MeteringComputationMeteringComputationQueryDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationQueryDataSourceFilterRepresentation}}
	MeteringComputationQueryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_metering_computation_query.test_query.id}`}},
	}

	MeteringComputationQueryRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		"query_definition": acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationQueryQueryDefinitionRepresentation},
	}
	MeteringComputationQueryQueryDefinitionRepresentation = map[string]interface{}{
		"cost_analysis_ui": acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationQueryQueryDefinitionCostAnalysisUIRepresentation},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"report_query":     acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationQueryQueryDefinitionReportQueryRepresentation},
		"version":          acctest.Representation{RepType: acctest.Required, Create: `1.0`, Update: `1.0`},
	}
	MeteringComputationQueryQueryDefinitionCostAnalysisUIRepresentation = map[string]interface{}{
		"graph":               acctest.Representation{RepType: acctest.Optional, Create: `BARS`, Update: `LINES`},
		"is_cumulative_graph": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	MeteringComputationQueryQueryDefinitionReportQueryRepresentation = map[string]interface{}{
		"forecast":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: MeteringComputationQueryQueryDefinitionReportQueryForecastRepresentation},
		"granularity":          acctest.Representation{RepType: acctest.Required, Create: `DAILY`, Update: `MONTHLY`},
		"tenant_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"compartment_depth":    acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `2.0`},
		"filter":               acctest.Representation{RepType: acctest.Optional, Create: `{\"operator\":\"AND\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"compartmentNameValue\"}],\"tags\":[],\"filters\":[]}`, Update: `{\"operator\":\"AND\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"compartmentNameValue2\"}],\"tags\":[],\"filters\":[]}`},
		"group_by":             acctest.Representation{RepType: acctest.Optional, Create: []string{`compartmentPath`}, Update: []string{`compartmentName`}},
		"group_by_tag":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: MeteringComputationQueryQueryDefinitionReportQueryGroupByTagRepresentation},
		"is_aggregate_by_time": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"query_type":           acctest.Representation{RepType: acctest.Optional, Create: `USAGE`, Update: `COST`},
		"time_usage_ended":     acctest.Representation{RepType: acctest.Required, Create: timeUsageEnded.Format(time.RFC3339Nano), Update: timeUsageEnded.Format(time.RFC3339Nano)},
		"time_usage_started":   acctest.Representation{RepType: acctest.Required, Create: timeUsageStarted.Format(time.RFC3339Nano), Update: timeUsageStarted.Format(time.RFC3339Nano)},
	}
	MeteringComputationQueryQueryDefinitionReportQueryForecastRepresentation = map[string]interface{}{
		"time_forecast_ended":   acctest.Representation{RepType: acctest.Required, Create: timeForecastEnded.Format(time.RFC3339Nano), Update: timeForecastEnded.Format(time.RFC3339Nano)},
		"forecast_type":         acctest.Representation{RepType: acctest.Optional, Create: `BASIC`},
		"time_forecast_started": acctest.Representation{RepType: acctest.Optional, Create: timeUsageEnded.Format(time.RFC3339Nano), Update: timeUsageEnded.Format(time.RFC3339Nano)},
	}
	MeteringComputationQueryQueryDefinitionReportQueryGroupByTagRepresentation = map[string]interface{}{
		"key":       acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"namespace": acctest.Representation{RepType: acctest.Optional, Create: `namespace`, Update: `namespace2`},
		"value":     acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	timeUsageStarted  = StartOfDay(time.Now().UTC().Truncate(time.Millisecond))
	timeUsageEnded    = StartOfDay(time.Now().UTC().AddDate(0, 0, 1).Truncate(time.Millisecond))
	timeForecastEnded = StartOfDay(time.Now().UTC().AddDate(0, 0, 2).Truncate(time.Millisecond))

	MeteringComputationQueryResourceDependencies = ""
)

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// issue-routing-tag: metering_computation/default
func TestMeteringComputationQueryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationQueryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)

	resourceName := "oci_metering_computation_query.test_query"
	datasourceName := "data.oci_metering_computation_queries.test_queries"
	singularDatasourceName := "data.oci_metering_computation_query.test_query"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckMeteringComputationQueryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + tenancyIdVariableStr + MeteringComputationQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_query", "test_query", acctest.Required, acctest.Create, MeteringComputationQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "query_definition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.cost_analysis_ui.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.granularity", "DAILY"),
				resource.TestCheckResourceAttrSet(resourceName, "query_definition.0.report_query.0.tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.version", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + tenancyIdVariableStr + MeteringComputationQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_query", "test_query", acctest.Optional, acctest.Update, MeteringComputationQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_queries", "test_queries", acctest.Optional, acctest.Update, MeteringComputationMeteringComputationQueryDataSourceRepresentation) +
				tenancyIdVariableStr + MeteringComputationQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_query", "test_query", acctest.Optional, acctest.Update, MeteringComputationQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttr(datasourceName, "query_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "query_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_query", "test_query", acctest.Required, acctest.Create, MeteringComputationMeteringComputationQuerySingularDataSourceRepresentation) +
				tenancyIdVariableStr + MeteringComputationQueryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
		// verify resource import
		{
			Config:                  config + QueryRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMeteringComputationQueryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).UsageapiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_metering_computation_query" {
			noResourceFound = false
			request := oci_metering_computation.GetQueryRequest{}

			tmp := rs.Primary.ID
			request.QueryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "metering_computation")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("MeteringComputationQuery") {
		resource.AddTestSweepers("MeteringComputationQuery", &resource.Sweeper{
			Name:         "MeteringComputationQuery",
			Dependencies: acctest.DependencyGraph["query"],
			F:            sweepMeteringComputationQueryResource,
		})
	}
}

func sweepMeteringComputationQueryResource(compartment string) error {
	usageapiClient := acctest.GetTestClients(&schema.ResourceData{}).UsageapiClient()
	queryIds, err := getMeteringComputationQueryIds(compartment)
	if err != nil {
		return err
	}
	for _, queryId := range queryIds {
		if ok := acctest.SweeperDefaultResourceId[queryId]; !ok {
			deleteQueryRequest := oci_metering_computation.DeleteQueryRequest{}

			deleteQueryRequest.QueryId = &queryId

			deleteQueryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "metering_computation")
			_, error := usageapiClient.DeleteQuery(context.Background(), deleteQueryRequest)
			if error != nil {
				fmt.Printf("Error deleting Query %s %s, It is possible that the resource is already deleted. Please verify manually \n", queryId, error)
				continue
			}
		}
	}
	return nil
}

func getMeteringComputationQueryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "QueryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	usageapiClient := acctest.GetTestClients(&schema.ResourceData{}).UsageapiClient()

	listQueriesRequest := oci_metering_computation.ListQueriesRequest{}
	listQueriesRequest.CompartmentId = &compartmentId
	listQueriesResponse, err := usageapiClient.ListQueries(context.Background(), listQueriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Query list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, query := range listQueriesResponse.Items {
		id := *query.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "QueryId", id)
	}
	return resourceIds, nil
}
