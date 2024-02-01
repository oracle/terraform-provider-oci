// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ManagementAgentManagementAgentDataSourceRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_data_source", "test_management_agent_data_source", acctest.Required, acctest.Create, ManagementAgentManagementAgentDataSourceRepresentation)

	ManagementAgentManagementAgentDataSourceResourceConfig                   = acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_data_source", "test_management_agent_data_source", acctest.Optional, acctest.Update, ManagementAgentManagementAgentDataSourceRepresentation)
	ManagementAgentManagementAgentDataSourceSingularDataSourceRepresentation = map[string]interface{}{
		"data_source_key":     acctest.Representation{RepType: acctest.Required, Create: `${oci_management_agent_management_agent_data_source.test_management_agent_data_source.id}`},
		"management_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_agent_id}`},
	}

	ManagementAgentManagementAgentDataSourceDataSourceRepresentation = map[string]interface{}{
		"management_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_agent_id}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `prometheusTerraformTest1`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ManagementAgentManagementAgentDataSourceDataSourceFilterRepresentation}}
	ManagementAgentManagementAgentDataSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_management_agent_management_agent_data_source.test_management_agent_data_source.name}`}},
	}

	ManagementAgentManagementAgentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"management_agent_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.managed_agent_id}`},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: `prometheusTerraformTest1`},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `PROMETHEUS_EMITTER`, Update: `PROMETHEUS_EMITTER`},
		"url":                          acctest.Representation{RepType: acctest.Required, Create: `url`, Update: `url2`},
		"allow_metrics":                acctest.Representation{RepType: acctest.Required, Create: `allowMetrics`, Update: `allowMetrics2`},
		"connection_timeout":           acctest.Representation{RepType: acctest.Optional, Create: "1", Update: "2"},
		"metric_dimensions":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: ManagementAgentManagementAgentDataSourceMetricDimensionsRepresentation},
		"namespace":                    acctest.Representation{RepType: acctest.Required, Create: `namespace`},
		"proxy_url":                    acctest.Representation{RepType: acctest.Optional, Create: `proxyUrl`, Update: `proxyUrl2`},
		"read_data_limit_in_kilobytes": acctest.Representation{RepType: acctest.Optional, Create: "10", Update: "11"},
		"read_timeout":                 acctest.Representation{RepType: acctest.Optional, Create: "20", Update: "21"},
		"resource_group":               acctest.Representation{RepType: acctest.Optional, Create: `resourceGroup`, Update: `resourceGroup2`},
		"schedule_mins":                acctest.Representation{RepType: acctest.Optional, Create: "30", Update: "31"},
	}
	ManagementAgentManagementAgentDataSourceRepresentationKubernetes = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"management_agent_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.managed_agent_id}`},
		"name":                         acctest.Representation{RepType: acctest.Required, Create: `kubernetesTerraformTest1`},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `KUBERNETES_CLUSTER`, Update: `KUBERNETES_CLUSTER`},
		"url":                          acctest.Representation{RepType: acctest.Required, Create: `url`, Update: `url2`},
		"allow_metrics":                acctest.Representation{RepType: acctest.Required, Create: `allowMetrics`, Update: `allowMetrics2`},
		"connection_timeout":           acctest.Representation{RepType: acctest.Optional, Create: "1", Update: "2"},
		"metric_dimensions":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: ManagementAgentManagementAgentDataSourceMetricDimensionsRepresentation},
		"namespace":                    acctest.Representation{RepType: acctest.Required, Create: `namespace`},
		"proxy_url":                    acctest.Representation{RepType: acctest.Optional, Create: `proxyUrl`, Update: `proxyUrl2`},
		"read_data_limit_in_kilobytes": acctest.Representation{RepType: acctest.Optional, Create: "10", Update: "11"},
		"read_timeout":                 acctest.Representation{RepType: acctest.Optional, Create: "20", Update: "21"},
		"resource_group":               acctest.Representation{RepType: acctest.Optional, Create: `resourceGroup`, Update: `resourceGroup2`},
		"schedule_mins":                acctest.Representation{RepType: acctest.Optional, Create: "30", Update: "31"},
	}
	ManagementAgentManagementAgentDataSourceMetricDimensionsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `metricname`, Update: `metricname2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `metricvalue`, Update: `metricvalue2`},
	}
	ManagementAgentManagementAgentRepresentation = map[string]interface{}{
		"managed_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_agent_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `terraformTest`, Update: `terraformTest3`},
		//"deploy_plugins_id": acctest.Representation{RepType: acctest.Optional, Create: []string{`${data.oci_management_agent_management_agent_plugins.test_management_agent_plugins.management_agent_plugins.0.id}`}},
	}
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentDataSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentDataSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentIds, err := getManagementAgentIds(compartmentId)
	if err != nil {
		t.Errorf("Failed to get agents in compartment %s", err)
	}
	if len(managementAgentIds) == 0 {
		t.Errorf("Failed to find any active agents in compartment %s", compartmentId)
	}
	managementAgentId := managementAgentIds[0]
	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_management_agent_management_agent_data_source.test_management_agent_data_source"
	datasourceName := "data.oci_management_agent_management_agent_data_sources.test_management_agent_data_sources"
	singularDatasourceName := "data.oci_management_agent_management_agent_data_source.test_management_agent_data_source"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managementAgentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_data_source", "test_management_agent_data_source", acctest.Optional, acctest.Create, ManagementAgentManagementAgentDataSourceRepresentation), "managementagent", "managementAgentDataSource", t)

	acctest.ResourceTest(t, testAccCheckManagementAgentManagementAgentDataSourceDestroy, []resource.TestStep{
		// verify Create
		// Step 0
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_data_source", "test_management_agent_data_source", acctest.Required, acctest.Create, ManagementAgentManagementAgentDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "prometheusTerraformTest1"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "type", "PROMETHEUS_EMITTER"),
				resource.TestCheckResourceAttr(resourceName, "url", "url"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		// step 1
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr,
		},
		// verify Create with optionals
		// step 2
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_data_source", "test_management_agent_data_source", acctest.Optional, acctest.Create, ManagementAgentManagementAgentDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allow_metrics", "allowMetrics"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_timeout", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "metric_dimensions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metric_dimensions.0.name", "metricname"),
				resource.TestCheckResourceAttr(resourceName, "metric_dimensions.0.value", "metricvalue"),
				resource.TestCheckResourceAttr(resourceName, "name", "prometheusTerraformTest1"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "proxy_url", "proxyUrl"),
				resource.TestCheckResourceAttr(resourceName, "read_data_limit_in_kilobytes", "10"),
				resource.TestCheckResourceAttr(resourceName, "read_timeout", "20"),
				resource.TestCheckResourceAttr(resourceName, "resource_group", "resourceGroup"),
				resource.TestCheckResourceAttr(resourceName, "schedule_mins", "30"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "PROMETHEUS_EMITTER"),
				resource.TestCheckResourceAttr(resourceName, "url", "url"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")

					// Id is the composite, so look for that composite in the export
					var compositeId string
					compositeId = fmt.Sprintf("managementAgents/%s/dataSources/%s", managementAgentId, resId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		// Step 3
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_data_source", "test_management_agent_data_source", acctest.Optional, acctest.Update, ManagementAgentManagementAgentDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allow_metrics", "allowMetrics2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_timeout", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "metric_dimensions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metric_dimensions.0.name", "metricname2"),
				resource.TestCheckResourceAttr(resourceName, "metric_dimensions.0.value", "metricvalue2"),
				resource.TestCheckResourceAttr(resourceName, "name", "prometheusTerraformTest1"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "proxy_url", "proxyUrl2"),
				resource.TestCheckResourceAttr(resourceName, "read_data_limit_in_kilobytes", "11"),
				resource.TestCheckResourceAttr(resourceName, "read_timeout", "21"),
				resource.TestCheckResourceAttr(resourceName, "resource_group", "resourceGroup2"),
				resource.TestCheckResourceAttr(resourceName, "schedule_mins", "31"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "PROMETHEUS_EMITTER"),
				resource.TestCheckResourceAttr(resourceName, "url", "url2"),

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
		// Step 4
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_data_sources", "test_management_agent_data_sources", acctest.Optional, acctest.Update, ManagementAgentManagementAgentDataSourceDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent_data_source", "test_management_agent_data_source", acctest.Optional, acctest.Update, ManagementAgentManagementAgentDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(datasourceName, "data_sources.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_sources.0.data_source_key"),
				resource.TestCheckResourceAttr(datasourceName, "data_sources.0.name", "prometheusTerraformTest1"),
				resource.TestCheckResourceAttr(datasourceName, "data_sources.0.type", "PROMETHEUS_EMITTER"),
			),
		},
		// verify singular datasource
		// Step 5
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_data_source", "test_management_agent_data_source", acctest.Required, acctest.Create, ManagementAgentManagementAgentDataSourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentManagementAgentDataSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_source_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "allow_metrics", "allowMetrics2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_timeout", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_dimensions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_dimensions.0.name", "metricname2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_dimensions.0.value", "metricvalue2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "prometheusTerraformTest1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "proxy_url", "proxyUrl2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "read_data_limit"),
				resource.TestCheckResourceAttr(singularDatasourceName, "read_timeout", "21"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_group", "resourceGroup2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_mins", "31"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "PROMETHEUS_EMITTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "url", "url2"),
			),
		},
		// verify resource import
		// Step 6
		{
			Config:            config + managementAgentIdVariableStr + ManagementAgentManagementAgentDataSourceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getManagementAgentManagementAgentDataSourceKeyCompositeId(resourceName),
			ImportStateVerifyIgnore: []string{
				"read_data_limit_in_kilobytes",
			},
			ResourceName: resourceName,
		},
	})
}
func getManagementAgentManagementAgentDataSourceKeyCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("managementAgents/%s/dataSources/%s", rs.Primary.Attributes["management_agent_id"], rs.Primary.Attributes["id"]), nil
	}
}
func testAccCheckManagementAgentManagementAgentDataSourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagementAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_management_agent_management_agent_data_source" {
			noResourceFound = false
			request := oci_management_agent.GetDataSourceRequest{}

			if value, ok := rs.Primary.Attributes["data_source_key"]; ok {
				request.DataSourceKey = &value
			}

			if value, ok := rs.Primary.Attributes["management_agent_id"]; ok {
				request.ManagementAgentId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_agent")

			response, err := client.GetDataSource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_management_agent.LifecycleStatesTerminated): true, string(oci_management_agent.LifecycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetState())
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
	if !acctest.InSweeperExcludeList("ManagementAgentManagementAgentDataSource") {
		resource.AddTestSweepers("ManagementAgentManagementAgentDataSource", &resource.Sweeper{
			Name:         "ManagementAgentManagementAgentDataSource",
			Dependencies: acctest.DependencyGraph["managementAgentDataSource"],
			F:            sweepManagementAgentManagementAgentDataSourceResource,
		})
	}
}

func sweepManagementAgentManagementAgentDataSourceResource(compartment string) error {
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()
	managementAgentDataSourceIds, err := getManagementAgentManagementAgentDataSourceIds(compartment)
	if err != nil {
		return err
	}
	for _, managementAgentDataSourceId := range managementAgentDataSourceIds {
		if ok := acctest.SweeperDefaultResourceId[managementAgentDataSourceId]; !ok {
			deleteDataSourceRequest := oci_management_agent.DeleteDataSourceRequest{}

			deleteDataSourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_agent")
			_, error := managementAgentClient.DeleteDataSource(context.Background(), deleteDataSourceRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagementAgentDataSource %s %s, It is possible that the resource is already deleted. Please verify manually \n", managementAgentDataSourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &managementAgentDataSourceId, ManagementAgentManagementAgentDataSourceSweepWaitCondition, time.Duration(3*time.Minute),
				ManagementAgentManagementAgentDataSourceSweepResponseFetchOperation, "management_agent", true)
		}
	}
	return nil
}
func getManagementAgentManagementAgentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagementAgentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	terraformTest := "terraformTest"
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()

	listManagementAgentsRequest := oci_management_agent.ListManagementAgentsRequest{}
	listManagementAgentsRequest.CompartmentId = &compartmentId
	listManagementAgentsRequest.LifecycleState = oci_management_agent.ListManagementAgentsLifecycleStateActive
	listManagementAgentsRequest.DisplayName = &terraformTest
	listManagementAgentsResponse, err := managementAgentClient.ListManagementAgents(context.Background(), listManagementAgentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagementAgent list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managementAgent := range listManagementAgentsResponse.Items {
		id := *managementAgent.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagementAgentId", id)
	}
	return resourceIds, nil
}
func getManagementAgentManagementAgentDataSourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagementAgentDataSourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()

	listDataSourcesRequest := oci_management_agent.ListDataSourcesRequest{}

	managementAgentIds, error := getManagementAgentManagementAgentIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting managementAgentId required for ManagementAgentDataSource resource requests \n")
	}
	for _, managementAgentId := range managementAgentIds {
		listDataSourcesRequest.ManagementAgentId = &managementAgentId

		listDataSourcesResponse, err := managementAgentClient.ListDataSources(context.Background(), listDataSourcesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting ManagementAgentDataSource list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, managementAgentDataSource := range listDataSourcesResponse.Items {
			id := *managementAgentDataSource.GetKey()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagementAgentDataSourceId", id)
		}

	}
	return resourceIds, nil
}

func ManagementAgentManagementAgentDataSourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managementAgentDataSourceResponse, ok := response.Response.(oci_management_agent.GetDataSourceResponse); ok {
		return managementAgentDataSourceResponse.GetState() != oci_management_agent.LifecycleStatesTerminated && managementAgentDataSourceResponse.GetState() != oci_management_agent.LifecycleStatesDeleted
	}
	return false
}

func ManagementAgentManagementAgentDataSourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagementAgentClient().GetDataSource(context.Background(), oci_management_agent.GetDataSourceRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
