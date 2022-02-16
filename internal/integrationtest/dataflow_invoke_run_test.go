// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/v58/dataflow"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InvokeRunRequiredOnlyResource = InvokeRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Required, acctest.Create, invokeRunRepresentation)

	InvokeRunResourceConfig = InvokeRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Update, invokeRunRepresentation)

	invokeRunSingularDataSourceRepresentation = map[string]interface{}{
		"run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_invoke_run.test_invoke_run.id}`},
	}

	invokeRunDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"application_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_dataflow_application.test_application.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: invokeRunDataSourceFilterRepresentation}}
	invokeRunDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataflow_invoke_run.test_invoke_run.id}`}},
	}

	invokeRunRepresentation = map[string]interface{}{
		"application_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_application.test_application.id}`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: `test_wordcount_run`},
		"arguments":            acctest.Representation{RepType: acctest.Optional, Create: []string{`arguments`}},
		"configuration":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"spark.shuffle.io.maxRetries": "10"}},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"driver_shape":         acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard2.1`},
		"executor_shape":       acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard2.1`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"logs_bucket_uri":      acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_logs_bucket_uri}`},
		"metastore_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.metastore_id}`},
		"num_executors":        acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"parameters":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: invokeRunParametersRepresentation},
		"type":                 acctest.Representation{RepType: acctest.Optional, Create: `BATCH`},
		"warehouse_bucket_uri": acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_warehouse_bucket_uri}`},
	}
	invokeRunParametersRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	InvokeRunResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Create, privateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Optional, acctest.Create, dataFlowApplicationRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dataflow/default
func TestDataflowInvokeRunResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowInvokeRunResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	fileUri := utils.GetEnvSettingWithBlankDefault("dataflow_file_uri")
	fileUriVariableStr := fmt.Sprintf("variable \"dataflow_file_uri\" { default = \"%s\" }\n", fileUri)
	archiveUri := utils.GetEnvSettingWithBlankDefault("dataflow_archive_uri")
	archiveUriVariableStr := fmt.Sprintf("variable \"dataflow_archive_uri\" { default = \"%s\" }\n", archiveUri)
	logsBucketUri := utils.GetEnvSettingWithBlankDefault("dataflow_logs_bucket_uri")
	logsBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_logs_bucket_uri\" { default = \"%s\" }\n", logsBucketUri)
	warehouseBucketUri := utils.GetEnvSettingWithBlankDefault("dataflow_warehouse_bucket_uri")
	warehouseBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_warehouse_bucket_uri\" { default = \"%s\" }\n", warehouseBucketUri)
	metastoreId := utils.GetEnvSettingWithBlankDefault("metastore_id")
	metastoreIdVariableStr := fmt.Sprintf("variable \"metastore_id\" { default = \"%s\" }\n", metastoreId)

	resourceName := "oci_dataflow_invoke_run.test_invoke_run"
	datasourceName := "data.oci_dataflow_invoke_runs.test_invoke_runs"
	singularDatasourceName := "data.oci_dataflow_invoke_run.test_invoke_run"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+InvokeRunResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Create, invokeRunRepresentation), "dataflow", "invokeRun", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Required, acctest.Create, invokeRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Create, invokeRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "language"),
				resource.TestCheckResourceAttrSet(resourceName, "logs_bucket_uri"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "spark_version"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "BATCH"),
				resource.TestCheckResourceAttr(resourceName, "warehouse_bucket_uri", warehouseBucketUri),
				resource.TestCheckResourceAttr(resourceName, "metastore_id", metastoreId),

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
			PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, dataflowRunAvailableShouldWaitCondition, time.Duration(20*time.Minute),
				dataFlowInvokeRunFetchOperation, "dataflow", false),
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InvokeRunResourceDependencies + warehouseBucketUriVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + metastoreIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(invokeRunRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "language"),
				resource.TestCheckResourceAttrSet(resourceName, "logs_bucket_uri"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "spark_version"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "BATCH"),
				resource.TestCheckResourceAttrSet(resourceName, "warehouse_bucket_uri"),

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
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Update, invokeRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "language"),
				resource.TestCheckResourceAttrSet(resourceName, "logs_bucket_uri"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "spark_version"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "BATCH"),
				resource.TestCheckResourceAttr(resourceName, "warehouse_bucket_uri", warehouseBucketUri),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_invoke_runs", "test_invoke_runs", acctest.Optional, acctest.Update, invokeRunDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Update, invokeRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "application_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "runs.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.application_id"),
				resource.TestCheckResourceAttr(datasourceName, "runs.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.data_read_in_bytes"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.data_written_in_bytes"),
				resource.TestCheckResourceAttr(datasourceName, "runs.0.display_name", "test_wordcount_run"),
				resource.TestCheckResourceAttr(datasourceName, "runs.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.language"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.opc_request_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.owner_principal_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.owner_user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.run_duration_in_milliseconds"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.total_ocpu"),
				resource.TestCheckResourceAttr(datasourceName, "runs.0.type", "BATCH"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Required, acctest.Create, invokeRunSingularDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "archive_uri"),
				resource.TestCheckResourceAttr(singularDatasourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_read_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_written_in_bytes"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "test_wordcount_run"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "file_uri"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "language"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logs_bucket_uri", logsBucketUri),
				resource.TestCheckResourceAttr(singularDatasourceName, "num_executors", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opc_request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner_user_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters.0.value", "value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_dns_zones.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_max_host_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_subnet_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_duration_in_milliseconds"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "spark_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_ocpu"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "BATCH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "warehouse_bucket_uri", warehouseBucketUri),
				resource.TestCheckResourceAttr(singularDatasourceName, "metastore_id", metastoreId),
			),
		},
	})
}

func testAccCheckDataflowInvokeRunDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataFlowClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataflow_invoke_run" {
			noResourceFound = false
			request := oci_dataflow.GetRunRequest{}

			tmp := rs.Primary.ID
			request.RunId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")

			_, err := client.GetRun(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataflowInvokeRun") {
		resource.AddTestSweepers("DataflowInvokeRun", &resource.Sweeper{
			Name:         "DataflowInvokeRun",
			Dependencies: acctest.DependencyGraph["invokeRun"],
			F:            sweepDataflowInvokeRunResource,
		})
	}
}

func sweepDataflowInvokeRunResource(compartment string) error {
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()
	invokeRunIds, err := getInvokeRunIds(compartment)
	if err != nil {
		return err
	}
	for _, invokeRunId := range invokeRunIds {
		if ok := acctest.SweeperDefaultResourceId[invokeRunId]; !ok {
			deleteRunRequest := oci_dataflow.DeleteRunRequest{}

			deleteRunRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")
			_, error := dataFlowClient.DeleteRun(context.Background(), deleteRunRequest)
			if error != nil {
				fmt.Printf("Error deleting InvokeRun %s %s, It is possible that the resource is already deleted. Please verify manually \n", invokeRunId, error)
				continue
			}
		}
	}
	return nil
}

func getInvokeRunIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InvokeRunId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()

	listRunsRequest := oci_dataflow.ListRunsRequest{}
	listRunsRequest.CompartmentId = &compartmentId
	listRunsResponse, err := dataFlowClient.ListRuns(context.Background(), listRunsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting InvokeRun list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, invokeRun := range listRunsResponse.Items {
		id := *invokeRun.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InvokeRunId", id)
	}
	return resourceIds, nil
}

func dataflowRunAvailableShouldWaitCondition(response common.OCIOperationResponse) bool {
	if runResponse, ok := response.Response.(oci_dataflow.GetRunResponse); ok {
		return runResponse.LifecycleState != oci_dataflow.RunLifecycleStateCanceled && runResponse.LifecycleState != oci_dataflow.RunLifecycleStateFailed &&
			runResponse.LifecycleState != oci_dataflow.RunLifecycleStateSucceeded
	}
	return false
}

func dataFlowInvokeRunFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataFlowClient().GetRun(context.Background(), oci_dataflow.GetRunRequest{
		RunId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
