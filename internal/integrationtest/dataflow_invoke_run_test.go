// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataflowInvokeRunRequiredOnlyResource = DataflowInvokeRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Required, acctest.Create, DataflowInvokeRunRepresentation)

	DataflowInvokeRunResourceConfig = DataflowInvokeRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Update, DataflowInvokeRunRepresentation)

	DataflowDataflowInvokeRunSingularDataSourceRepresentation = map[string]interface{}{
		"run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_invoke_run.test_invoke_run.id}`},
	}

	DataflowDataflowInvokeRunDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"application_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_dataflow_application.test_application.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowInvokeRunDataSourceFilterRepresentation}}
	DataflowInvokeRunDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataflow_invoke_run.test_invoke_run.id}`}},
	}

	DataflowInvokeRunRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"application_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_application.test_application.id}`},
		"application_log_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowInvokeRunApplicationLogConfigRepresentation},
		"arguments":              acctest.Representation{RepType: acctest.Optional, Create: []string{`arguments`}},
		"configuration":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"spark.shuffle.io.maxRetries": "10"}},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `test_wordcount_run`},
		"driver_shape":           acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard2.1`},
		"driver_shape_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowInvokeRunDriverShapeConfigRepresentation},
		"executor_shape":         acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard2.1`},
		"executor_shape_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowInvokeRunExecutorShapeConfigRepresentation},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"logs_bucket_uri":        acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_logs_bucket_uri}`},
		"metastore_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.metastore_id}`},
		"num_executors":          acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"parameters":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowInvokeRunParametersRepresentation},
		"type":                   acctest.Representation{RepType: acctest.Optional, Create: `BATCH`},
		"warehouse_bucket_uri":   acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_warehouse_bucket_uri}`},
		"lifecycle":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForDataFlowResource},
		"pool_id":                acctest.Representation{RepType: acctest.Optional, Create: ``},
	}
	DataflowInvokeRunApplicationLogConfigRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
		"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_log.id}`},
	}
	DataflowInvokeRunDriverShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `15`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}
	DataflowInvokeRunExecutorShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `15`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}
	DataflowInvokeRunParametersRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	DataflowInvokeRunResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Create, DataflowPrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Optional, acctest.Create, DataflowApplicationRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, DataflowApplicationLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, DataflowApplicationLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Optional, acctest.Create, DataflowPoolRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataflowInvokeRunResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Create, DataflowInvokeRunRepresentation), "dataflow", "invokeRun", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			PreConfig: func() {
				fmt.Println("step 1")
			},
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataflowInvokeRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Required, acctest.Create, DataflowInvokeRunRepresentation),
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
			PreConfig: func() {
				fmt.Println("step 2, in delete")
			},
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataflowInvokeRunResourceDependencies,
		},
		// verify Create with optionals
		{
			PreConfig: func() {
				fmt.Println("step 3, create with optionals")
			},
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataflowInvokeRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Create, DataflowInvokeRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "application_log_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.ocpus", "1"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataflowInvokeRunResourceDependencies + warehouseBucketUriVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + metastoreIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataflowInvokeRunRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "application_log_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.ocpus", "1"),
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
			PreConfig: func() {
				fmt.Println("step 5 updatable params")
			},
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataflowInvokeRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Update, DataflowInvokeRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "application_log_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_run"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.ocpus", "1"),
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
			PreConfig: func() {
				fmt.Println("step 6, verify datasource")
			},
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_invoke_runs", "test_invoke_runs", acctest.Optional, acctest.Update, DataflowDataflowInvokeRunDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataflowInvokeRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Optional, acctest.Update, DataflowInvokeRunRepresentation),
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

			PreConfig: func() {
				fmt.Println("step 7, singular datasource")

				acctest.WaitTillCondition(acctest.TestAccProvider, &resId, dataflowRunAvailableShouldWaitCondition, time.Duration(20*time.Minute),
					dataFlowInvokeRunFetchOperation, "dataflow", false)()
				client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataFlowClient()
				run, _ := client.GetRun(context.Background(), oci_dataflow.GetRunRequest{
					RunId: &resId,
					RequestMetadata: common.RequestMetadata{
						RetryPolicy: tfresource.GetRetryPolicy(true, "dataflow"),
					},
				})

				stopPoolRequest := oci_dataflow.StopPoolRequest{}
				stopPoolRequest.PoolId = run.PoolId
				stopPoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")
				_, errorStopping := client.StopPool(context.Background(), stopPoolRequest)
				if errorStopping != nil {
					fmt.Printf("Error stopping Pool %s, It is possible that the resource is already stopped. Please verify manually \n", errorStopping)
				}
				acctest.WaitTillCondition(acctest.TestAccProvider, run.PoolId, DataflowPoolStopWaitCondition, time.Duration(10*time.Minute),
					DataflowPoolSweepResponseFetchOperation, "dataflow", true)()
			},
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Required, acctest.Create, DataflowDataflowInvokeRunSingularDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataflowInvokeRunResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "application_log_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(singularDatasourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_read_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_written_in_bytes"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "test_wordcount_run"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape_config.0.memory_in_gbs", "15"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape_config.0.ocpus", "1"),
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
	invokeRunIds, err := getDataflowInvokeRunIds(compartment)
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

func getDataflowInvokeRunIds(compartment string) ([]string, error) {
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
