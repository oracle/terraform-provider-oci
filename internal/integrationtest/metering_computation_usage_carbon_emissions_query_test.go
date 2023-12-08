// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

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
	CarbonEmissionsQueryRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_carbon_emissions_query", "test_usage_carbon_emissions_query", acctest.Required, acctest.Create, MeteringComputationUsageCarbonEmissionsQueryRepresentation)

	MeteringComputationUsageCarbonEmissionsQueryResourceConfig = MeteringComputationUsageCarbonEmissionsQueryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_carbon_emissions_query", "test_usage_carbon_emissions_query", acctest.Optional, acctest.Update, MeteringComputationUsageCarbonEmissionsQueryRepresentation)

	MeteringComputationUsageCarbonEmissionsQuerySingularDataSourceRepresentation = map[string]interface{}{
		"usage_carbon_emissions_query_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_metering_computation_usage_carbon_emissions_query.test_usage_carbon_emissions_query.id}`},
	}

	MeteringComputationUsageCarbonEmissionsQueryDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationUsageCarbonEmissionsQueryDataSourceFilterRepresentation}}
	MeteringComputationUsageCarbonEmissionsQueryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_metering_computation_usage_carbon_emissions_query.test_usage_carbon_emissions_query.id}`}},
	}

	MeteringComputationUsageCarbonEmissionsQueryRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		"query_definition": acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationUsageCarbonEmissionsQueryQueryDefinitionRepresentation},
	}
	MeteringComputationUsageCarbonEmissionsQueryQueryDefinitionRepresentation = map[string]interface{}{
		"cost_analysis_ui": acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationUsageCarbonEmissionsQueryQueryDefinitionCostAnalysisUIRepresentation},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"report_query":     acctest.RepresentationGroup{RepType: acctest.Required, Group: MeteringComputationUsageCarbonEmissionsQueryQueryDefinitionReportQueryRepresentation},
		"version":          acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
	}
	MeteringComputationUsageCarbonEmissionsQueryQueryDefinitionCostAnalysisUIRepresentation = map[string]interface{}{
		"graph":               acctest.Representation{RepType: acctest.Optional, Create: `BARS`, Update: `LINES`},
		"is_cumulative_graph": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	MeteringComputationUsageCarbonEmissionsQueryQueryDefinitionReportQueryRepresentation = map[string]interface{}{
		"tenant_id":                           acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"compartment_depth":                   acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `2.0`},
		"group_by":                            acctest.Representation{RepType: acctest.Optional, Create: []string{`compartmentPath`}, Update: []string{`compartmentName`}},
		"group_by_tag":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: MeteringComputationUsageCarbonEmissionsQueryQueryDefinitionReportQueryGroupByTagRepresentation},
		"is_aggregate_by_time":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"time_usage_ended":                    acctest.Representation{RepType: acctest.Required, Create: `2023-07-01T00:00:00Z`},
		"time_usage_started":                  acctest.Representation{RepType: acctest.Required, Create: `2023-01-01T00:00:00Z`},
		"usage_carbon_emissions_query_filter": acctest.Representation{RepType: acctest.Optional, Create: `{\"operator\":\"AND\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"compartmentNameValue\"}],\"tags\":[],\"filters\":[]}`, Update: `{\"operator\":\"AND\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"compartmentNameValue2\"}],\"tags\":[],\"filters\":[]}`},
	}
	MeteringComputationUsageCarbonEmissionsQueryQueryDefinitionReportQueryGroupByTagRepresentation = map[string]interface{}{
		"key":       acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"namespace": acctest.Representation{RepType: acctest.Optional, Create: `namespace`, Update: `namespace2`},
		"value":     acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}

	MeteringComputationUsageCarbonEmissionsQueryResourceDependencies = ""
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationUsageCarbonEmissionsQueryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationUsageCarbonEmissionsQueryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)

	resourceName := "oci_metering_computation_usage_carbon_emissions_query.test_usage_carbon_emissions_query"
	datasourceName := "data.oci_metering_computation_usage_carbon_emissions_queries.test_usage_carbon_emissions_queries"
	singularDatasourceName := "data.oci_metering_computation_usage_carbon_emissions_query.test_usage_carbon_emissions_query"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+tenancyIdVariableStr+MeteringComputationUsageCarbonEmissionsQueryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_carbon_emissions_query", "test_usage_carbon_emissions_query", acctest.Required, acctest.Create, MeteringComputationUsageCarbonEmissionsQueryRepresentation), "usageapi", "usageCarbonEmissionsQuery", t)

	acctest.ResourceTest(t, testAccCheckMeteringComputationUsageCarbonEmissionsQueryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + tenancyIdVariableStr + MeteringComputationUsageCarbonEmissionsQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_carbon_emissions_query", "test_usage_carbon_emissions_query", acctest.Required, acctest.Create, MeteringComputationUsageCarbonEmissionsQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "query_definition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.cost_analysis_ui.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "query_definition.0.report_query.0.tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.version", "1"),

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
			Config: config + tenancyIdVariableStr + MeteringComputationUsageCarbonEmissionsQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_carbon_emissions_query", "test_usage_carbon_emissions_query", acctest.Optional, acctest.Update, MeteringComputationUsageCarbonEmissionsQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", "ocid1.tenancy.oc1..aaaaaaaattzudpdasr7fnrskmyk6lepuc54anjzq7jvmqjxfrwshale4lelq"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.cost_analysis_ui.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.cost_analysis_ui.0.graph", "LINES"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.cost_analysis_ui.0.is_cumulative_graph", "true"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.compartment_depth", "2"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.group_by.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.group_by_tag.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.group_by_tag.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.group_by_tag.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.group_by_tag.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.is_aggregate_by_time", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "query_definition.0.report_query.0.tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.time_usage_ended", "2023-07-01T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.time_usage_started", "2023-01-01T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "query_definition.0.report_query.0.usage_carbon_emissions_query_filter", "{\"operator\":\"AND\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"compartmentNameValue2\"}],\"tags\":[],\"filters\":[]}"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_usage_carbon_emissions_queries", "test_usage_carbon_emissions_queries", acctest.Optional, acctest.Update, MeteringComputationUsageCarbonEmissionsQueryDataSourceRepresentation) +
				tenancyIdVariableStr + MeteringComputationUsageCarbonEmissionsQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_metering_computation_usage_carbon_emissions_query", "test_usage_carbon_emissions_query", acctest.Optional, acctest.Update, MeteringComputationUsageCarbonEmissionsQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttr(datasourceName, "usage_carbon_emissions_query_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "usage_carbon_emissions_query_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_usage_carbon_emissions_query", "test_usage_carbon_emissions_query", acctest.Required, acctest.Create, MeteringComputationUsageCarbonEmissionsQuerySingularDataSourceRepresentation) +
				tenancyIdVariableStr + MeteringComputationUsageCarbonEmissionsQueryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "usage_carbon_emissions_query_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.cost_analysis_ui.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.cost_analysis_ui.0.graph", "LINES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.cost_analysis_ui.0.is_cumulative_graph", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.compartment_depth", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.group_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.group_by_tag.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.group_by_tag.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.group_by_tag.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.group_by_tag.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.is_aggregate_by_time", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.report_query.0.usage_carbon_emissions_query_filter", "{\"operator\":\"AND\",\"dimensions\":[{\"key\":\"compartmentName\",\"value\":\"compartmentNameValue2\"}],\"tags\":[],\"filters\":[]}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_definition.0.version", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + MeteringComputationUsageCarbonEmissionsQueryResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMeteringComputationUsageCarbonEmissionsQueryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).UsageapiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_metering_computation_usage_carbon_emissions_query" {
			noResourceFound = false
			request := oci_metering_computation.GetUsageCarbonEmissionsQueryRequest{}

			tmp := rs.Primary.ID
			request.UsageCarbonEmissionsQueryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "metering_computation")

			_, err := client.GetUsageCarbonEmissionsQuery(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("MeteringComputationUsageCarbonEmissionsQuery") {
		resource.AddTestSweepers("MeteringComputationUsageCarbonEmissionsQuery", &resource.Sweeper{
			Name:         "MeteringComputationUsageCarbonEmissionsQuery",
			Dependencies: acctest.DependencyGraph["usageCarbonEmissionsQuery"],
			F:            sweepMeteringComputationUsageCarbonEmissionsQueryResource,
		})
	}
}

func sweepMeteringComputationUsageCarbonEmissionsQueryResource(compartment string) error {
	usageapiClient := acctest.GetTestClients(&schema.ResourceData{}).UsageapiClient()
	usageCarbonEmissionsQueryIds, err := getMeteringComputationUsageCarbonEmissionsQueryIds(compartment)
	if err != nil {
		return err
	}
	for _, usageCarbonEmissionsQueryId := range usageCarbonEmissionsQueryIds {
		if ok := acctest.SweeperDefaultResourceId[usageCarbonEmissionsQueryId]; !ok {
			deleteUsageCarbonEmissionsQueryRequest := oci_metering_computation.DeleteUsageCarbonEmissionsQueryRequest{}

			deleteUsageCarbonEmissionsQueryRequest.UsageCarbonEmissionsQueryId = &usageCarbonEmissionsQueryId

			deleteUsageCarbonEmissionsQueryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "metering_computation")
			_, error := usageapiClient.DeleteUsageCarbonEmissionsQuery(context.Background(), deleteUsageCarbonEmissionsQueryRequest)
			if error != nil {
				fmt.Printf("Error deleting UsageCarbonEmissionsQuery %s %s, It is possible that the resource is already deleted. Please verify manually \n", usageCarbonEmissionsQueryId, error)
				continue
			}
		}
	}
	return nil
}

func getMeteringComputationUsageCarbonEmissionsQueryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "UsageCarbonEmissionsQueryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	usageapiClient := acctest.GetTestClients(&schema.ResourceData{}).UsageapiClient()

	listUsageCarbonEmissionsQueriesRequest := oci_metering_computation.ListUsageCarbonEmissionsQueriesRequest{}
	listUsageCarbonEmissionsQueriesRequest.CompartmentId = &compartmentId
	listUsageCarbonEmissionsQueriesResponse, err := usageapiClient.ListUsageCarbonEmissionsQueries(context.Background(), listUsageCarbonEmissionsQueriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting UsageCarbonEmissionsQuery list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, usageCarbonEmissionsQuery := range listUsageCarbonEmissionsQueriesResponse.Items {
		id := *usageCarbonEmissionsQuery.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "UsageCarbonEmissionsQueryId", id)
	}
	return resourceIds, nil
}
