// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataflowSqlEndpointRequiredOnlyResource = DataflowSqlEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Required, acctest.Create, DataflowSqlEndpointRepresentation)

	DataflowSqlEndpointResourceConfig = DataflowSqlEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Update, DataflowSqlEndpointRepresentation)

	DataflowSqlEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"sql_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_sql_endpoint.test_sql_endpoint.id}`},
	}

	DataflowSqlEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DataflowSqlEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataflow_sql_endpoint.test_sql_endpoint.id}`}},
	}

	DataflowSqlEndpointRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `test_sql_endpoint_terraform`, Update: `test_sql_endpoint_terraform_updated`},
		"driver_shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"executor_shape":        acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"max_executor_count":    acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `2`},
		"metastore_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.metastore_id}`},
		"min_executor_count":    acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"network_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowSqlEndpointNetworkConfigurationRepresentation},
		"sql_endpoint_version":  acctest.Representation{RepType: acctest.Required, Create: `3.2.1`},
		//"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description updated`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSqlEndpointDefinedTagsRepresentation},
		"spark_advanced_configurations": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{}, Update: map[string]string{"testConfig": "testValue"}},
	}

	DataflowSqlEndpointNetworkConfigurationRepresentation = map[string]interface{}{
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `SECURE_ACCESS`},
		// "access_control_rules": acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		"access_control_rules": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowSqlEndpointNetworkConfigurationAccessControlRulesRepresentation},
	}

	DataflowSqlEndpointDriverShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `15`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}
	DataflowSqlEndpointExecutorShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `15`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	ignoreSqlEndpointDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
	}

	DataflowSqlEndpointSparkConfigurationRepresentation = map[string]interface{}{
		"testConfig": acctest.Representation{RepType: acctest.Required, Create: `testValue`},
	}

	DataflowSqlEndpointResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: dataflow/default
func TestDataflowSqlEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowSqlEndpointResource_basic")

	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	metastoreId := utils.GetEnvSettingWithBlankDefault("metastore_id")
	metastoreIdVariableStr := fmt.Sprintf("variable \"metastore_id\" { default = \"%s\" }\n", metastoreId)

	resourceName := "oci_dataflow_sql_endpoint.test_sql_endpoint"
	datasourceName := "data.oci_dataflow_sql_endpoints.test_sql_endpoints"
	singularDatasourceName := "data.oci_dataflow_sql_endpoint.test_sql_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+metastoreIdVariableStr+DataflowSqlEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Create, DataflowSqlEndpointRepresentation), "dataflow", "sqlEndpoint", t)

	acctest.ResourceTest(t, testAccCheckDataflowSqlEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + DataflowSqlEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Required, acctest.Create, DataflowSqlEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_terraform"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "SECURE_ACCESS"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + DataflowSqlEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + DataflowSqlEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Create, DataflowSqlEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_terraform"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "SECURE_ACCESS"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.public_endpoint_ip"),
				resource.TestCheckResourceAttr(resourceName, "spark_advanced_configurations.%", "0"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},
		// Stop the sql endpoint
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + DataflowSqlEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataflowSqlEndpointRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Required, Create: `INACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Start the SQL Endpoint again
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + DataflowSqlEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataflowSqlEndpointRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Required, Create: `ACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify Update to the updateable parameters including compartment
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + metastoreIdVariableStr + DataflowSqlEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataflowSqlEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`, Update: `${var.compartment_id_for_update}`},
					}),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description updated"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_terraform_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "spark_advanced_configurations.%", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Switch back to the original compartment
		{
			Config: config + compartmentIdVariableStr + metastoreIdVariableStr + DataflowSqlEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataflowSqlEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id}`},
					}),
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description updated"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_sql_endpoint_terraform_updated"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "min_executor_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_configuration.0.network_type", "SECURE_ACCESS"),
				resource.TestCheckResourceAttrSet(resourceName, "network_configuration.0.public_endpoint_ip"),
				resource.TestCheckResourceAttr(resourceName, "spark_advanced_configurations.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "sql_endpoint_version", "3.2.1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_sql_endpoints", "test_sql_endpoints", acctest.Optional, acctest.Update, DataflowSqlEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + metastoreIdVariableStr + DataflowSqlEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Update, DataflowSqlEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttr(datasourceName, "sql_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_endpoint_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_sql_endpoint", "test_sql_endpoint", acctest.Optional, acctest.Update, DataflowSqlEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + metastoreIdVariableStr + DataflowSqlEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "test_sql_endpoint_terraform_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape_config.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape_config.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jdbc_endpoint_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_executor_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_executor_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_configuration.0.network_type", "SECURE_ACCESS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_configuration.0.public_endpoint_ip"),
				resource.TestCheckResourceAttr(singularDatasourceName, "spark_advanced_configurations.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sql_endpoint_version", "3.2.1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataflowSqlEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataflowSqlEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataFlowClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataflow_sql_endpoint" {
			noResourceFound = false
			request := oci_dataflow.GetSqlEndpointRequest{}

			tmp := rs.Primary.ID
			request.SqlEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")

			response, err := client.GetSqlEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dataflow.SqlEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataflowSqlEndpoint") {
		resource.AddTestSweepers("DataflowSqlEndpoint", &resource.Sweeper{
			Name:         "DataflowSqlEndpoint",
			Dependencies: acctest.DependencyGraph["sqlEndpoint"],
			F:            sweepDataflowSqlEndpointResource,
		})
	}
}

func sweepDataflowSqlEndpointResource(compartment string) error {
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()
	sqlEndpointIds, err := getDataflowSqlEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, sqlEndpointId := range sqlEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[sqlEndpointId]; !ok {
			deleteSqlEndpointRequest := oci_dataflow.DeleteSqlEndpointRequest{}

			deleteSqlEndpointRequest.SqlEndpointId = &sqlEndpointId

			deleteSqlEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")
			_, error := dataFlowClient.DeleteSqlEndpoint(context.Background(), deleteSqlEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting SqlEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", sqlEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sqlEndpointId, DataflowSqlEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				DataflowSqlEndpointSweepResponseFetchOperation, "dataflow", true)
		}
	}
	return nil
}

func getDataflowSqlEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SqlEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()

	listSqlEndpointsRequest := oci_dataflow.ListSqlEndpointsRequest{}
	listSqlEndpointsRequest.CompartmentId = &compartmentId
	listSqlEndpointsRequest.LifecycleState = oci_dataflow.ListSqlEndpointsLifecycleStateActive
	listSqlEndpointsResponse, err := dataFlowClient.ListSqlEndpoints(context.Background(), listSqlEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SqlEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sqlEndpoint := range listSqlEndpointsResponse.Items {
		id := *sqlEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SqlEndpointId", id)
	}
	return resourceIds, nil
}

func DataflowSqlEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sqlEndpointResponse, ok := response.Response.(oci_dataflow.GetSqlEndpointResponse); ok {
		return sqlEndpointResponse.LifecycleState != oci_dataflow.SqlEndpointLifecycleStateDeleted
	}
	return false
}

func DataflowSqlEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataFlowClient().GetSqlEndpoint(context.Background(), oci_dataflow.GetSqlEndpointRequest{
		SqlEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
