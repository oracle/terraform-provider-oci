// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/dataflow"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InvokeRunRequiredOnlyResource = InvokeRunResourceDependencies +
		generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", Required, Create, invokeRunRepresentation)

	InvokeRunResourceConfig = InvokeRunResourceDependencies +
		generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", Optional, Update, invokeRunRepresentation)

	invokeRunSingularDataSourceRepresentation = map[string]interface{}{
		"run_id": Representation{repType: Required, create: `${oci_dataflow_invoke_run.test_invoke_run.id}`},
	}

	invokeRunDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"application_id": Representation{repType: Optional, create: `${oci_dataflow_application.test_application.id}`},
		"filter":         RepresentationGroup{Required, invokeRunDataSourceFilterRepresentation}}
	invokeRunDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_dataflow_invoke_run.test_invoke_run.id}`}},
	}

	invokeRunRepresentation = map[string]interface{}{
		"application_id":       Representation{repType: Required, create: `${oci_dataflow_application.test_application.id}`},
		"compartment_id":       Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":         Representation{repType: Required, create: `test_wordcount_run`},
		"arguments":            Representation{repType: Optional, create: []string{`arguments`}},
		"configuration":        Representation{repType: Optional, create: map[string]string{"spark.shuffle.io.maxRetries": "10"}},
		"defined_tags":         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"driver_shape":         Representation{repType: Optional, create: `VM.Standard2.1`},
		"executor_shape":       Representation{repType: Optional, create: `VM.Standard2.1`},
		"freeform_tags":        Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"logs_bucket_uri":      Representation{repType: Optional, create: `${var.dataflow_logs_bucket_uri}`},
		"num_executors":        Representation{repType: Optional, create: `1`},
		"parameters":           RepresentationGroup{Optional, invokeRunParametersRepresentation},
		"warehouse_bucket_uri": Representation{repType: Optional, create: `${var.dataflow_warehouse_bucket_uri}`},
	}
	invokeRunParametersRepresentation = map[string]interface{}{
		"name":  Representation{repType: Required, create: `name`},
		"value": Representation{repType: Required, create: `value`},
	}

	InvokeRunResourceDependencies = generateResourceFromRepresentationMap("oci_dataflow_application", "test_application", Optional, Create, dataFlowApplicationRepresentation) +
		DefinedTagsDependencies
)

func TestDataflowInvokeRunResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowInvokeRunResource_basic")
	defer httpreplay.SaveScenario()
	const DataFlowWaitConditionDuration = time.Duration(20 * time.Minute)

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	fileUri := getEnvSettingWithBlankDefault("dataflow_file_uri")
	fileUriVariableStr := fmt.Sprintf("variable \"dataflow_file_uri\" { default = \"%s\" }\n", fileUri)
	logsBucketUri := getEnvSettingWithBlankDefault("dataflow_logs_bucket_uri")
	logsBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_logs_bucket_uri\" { default = \"%s\" }\n", logsBucketUri)
	warehouseBucketUri := getEnvSettingWithBlankDefault("dataflow_warehouse_bucket_uri")
	warehouseBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_warehouse_bucket_uri\" { default = \"%s\" }\n", warehouseBucketUri)

	resourceName := "oci_dataflow_invoke_run.test_invoke_run"
	datasourceName := "data.oci_dataflow_invoke_runs.test_invoke_runs"
	singularDatasourceName := "data.oci_dataflow_invoke_run.test_invoke_run"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + fileUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + InvokeRunResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", Required, Create, invokeRunRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + fileUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + InvokeRunResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + fileUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + InvokeRunResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", Optional, Create, invokeRunRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),
					resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "language"),
					resource.TestCheckResourceAttr(resourceName, "logs_bucket_uri", logsBucketUri),
					resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
					resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "name"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "value"),
					resource.TestCheckResourceAttrSet(resourceName, "spark_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "warehouse_bucket_uri", warehouseBucketUri),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				PreConfig: waitTillCondition(testAccProvider, &resId, dataflowRunAvailableShouldWaitCondition, DataFlowWaitConditionDuration,
					dataFlowInvokeRunFetchOperation, "dataflow", false),
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InvokeRunResourceDependencies + warehouseBucketUriVariableStr + fileUriVariableStr + logsBucketUriVariableStr +
					generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", Optional, Create,
						representationCopyWithNewProperties(invokeRunRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),
					resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "language"),
					resource.TestCheckResourceAttrSet(resourceName, "logs_bucket_uri"),
					resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
					resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "name"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "value"),
					resource.TestCheckResourceAttrSet(resourceName, "spark_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "warehouse_bucket_uri"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + fileUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + InvokeRunResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", Optional, Update, invokeRunRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),
					resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
					resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "language"),
					resource.TestCheckResourceAttr(resourceName, "logs_bucket_uri", logsBucketUri),
					resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
					resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "name"),
					resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "value"),
					resource.TestCheckResourceAttrSet(resourceName, "spark_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "warehouse_bucket_uri", warehouseBucketUri),

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
					generateDataSourceFromRepresentationMap("oci_dataflow_invoke_runs", "test_invoke_runs", Optional, Update, invokeRunDataSourceRepresentation) +
					compartmentIdVariableStr + fileUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + InvokeRunResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", Optional, Update, invokeRunRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "application_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "runs.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "runs.0.application_id"),
					resource.TestCheckResourceAttr(datasourceName, "runs.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "runs.0.data_read_in_bytes"),
					resource.TestCheckResourceAttrSet(datasourceName, "runs.0.data_written_in_bytes"),
					resource.TestCheckResourceAttr(datasourceName, "runs.0.defined_tags.%", "1"),
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
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", Required, Create, invokeRunSingularDataSourceRepresentation) +
					compartmentIdVariableStr + fileUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + InvokeRunResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "run_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "arguments.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "configuration.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_read_in_bytes"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_written_in_bytes"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
					resource.TestCheckResourceAttrSet(singularDatasourceName, "run_duration_in_milliseconds"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "spark_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "total_ocpu"),
					resource.TestCheckResourceAttr(singularDatasourceName, "warehouse_bucket_uri", warehouseBucketUri),
				),
			},
		},
	})
}

func testAccCheckDataflowInvokeRunDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataFlowClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataflow_invoke_run" {
			noResourceFound = false
			request := oci_dataflow.GetRunRequest{}

			tmp := rs.Primary.ID
			request.RunId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dataflow")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DataflowInvokeRun") {
		resource.AddTestSweepers("DataflowInvokeRun", &resource.Sweeper{
			Name:         "DataflowInvokeRun",
			Dependencies: DependencyGraph["invokeRun"],
			F:            sweepDataflowInvokeRunResource,
		})
	}
}

func sweepDataflowInvokeRunResource(compartment string) error {
	dataFlowClient := GetTestClients(&schema.ResourceData{}).dataFlowClient()
	invokeRunIds, err := getInvokeRunIds(compartment)
	if err != nil {
		return err
	}
	for _, invokeRunId := range invokeRunIds {
		if ok := SweeperDefaultResourceId[invokeRunId]; !ok {
			deleteRunRequest := oci_dataflow.DeleteRunRequest{}

			deleteRunRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dataflow")
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
	ids := getResourceIdsToSweep(compartment, "InvokeRunId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataFlowClient := GetTestClients(&schema.ResourceData{}).dataFlowClient()

	listRunsRequest := oci_dataflow.ListRunsRequest{}
	listRunsRequest.CompartmentId = &compartmentId
	listRunsResponse, err := dataFlowClient.ListRuns(context.Background(), listRunsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting InvokeRun list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, invokeRun := range listRunsResponse.Items {
		id := *invokeRun.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "InvokeRunId", id)
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

func dataFlowInvokeRunFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataFlowClient().GetRun(context.Background(), oci_dataflow.GetRunRequest{
		RunId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
