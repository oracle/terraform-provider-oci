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
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_dataflow "github.com/oracle/oci-go-sdk/v58/dataflow"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InvokeRunSubmitRequiredOnlyResource = InvokeRunSubmitResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", acctest.Required, acctest.Create, invokeRunSubmitRepresentation)

	InvokeRunSubmitResourceConfig = InvokeRunSubmitResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", acctest.Optional, acctest.Update, invokeRunSubmitRepresentation)

	invokeRunSubmitSingularDataSourceRepresentation = map[string]interface{}{
		"run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_invoke_run.test_invoke_run_submit.id}`},
	}

	invokeRunSubmitDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: invokeRunSubmitDataSourceFilterRepresentation}}
	invokeRunSubmitDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataflow_invoke_run.test_invoke_run_submit.id}`}},
	}

	invokeRunSubmitRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"archive_uri":          acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("dataflow_archive_uri")},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `test_wordcount_runsubmit`},
		"driver_shape":         acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard2.1`},
		"execute":              acctest.Representation{RepType: acctest.Required, Create: `--conf spark.shuffle.io.maxRetries=10 ` + utils.GetEnvSettingWithBlankDefault("dataflow_file_uri") + ` arguments`},
		"executor_shape":       acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard2.1`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"logs_bucket_uri":      acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_logs_bucket_uri}`},
		"metastore_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.metastore_id}`},
		"num_executors":        acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"spark_version":        acctest.Representation{RepType: acctest.Optional, Create: `2.4.4`},
		"warehouse_bucket_uri": acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_warehouse_bucket_uri}`},
	}

	InvokeRunSubmitResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dataflow/default
func TestDataflowInvokeRunResource_SparkSubmit(t *testing.T) {
	httpreplay.SetScenario("TestDataflowInvokeRunResource_SparkSubmit")
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

	resourceName := "oci_dataflow_invoke_run.test_invoke_run_submit"
	datasourceName := "data.oci_dataflow_invoke_runs.test_invoke_run_submit"
	singularDatasourceName := "data.oci_dataflow_invoke_run.test_invoke_run_submit"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create run with required execute
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", acctest.Required, acctest.Create, invokeRunSubmitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceDependencies,
		},
		// verify Create with execute, display_name, spark_version
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", acctest.Optional, acctest.Create, invokeRunSubmitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_runsubmit"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "language"),
				resource.TestCheckResourceAttr(resourceName, "logs_bucket_uri", logsBucketUri),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttr(resourceName, "spark_version", "2.4.4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "warehouse_bucket_uri", warehouseBucketUri),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InvokeRunSubmitResourceDependencies + warehouseBucketUriVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + metastoreIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(invokeRunSubmitRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_runsubmit"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "language"),
				resource.TestCheckResourceAttrSet(resourceName, "logs_bucket_uri"),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttr(resourceName, "spark_version", "2.4.4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
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
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", acctest.Optional, acctest.Update, invokeRunSubmitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_runsubmit"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "language"),
				resource.TestCheckResourceAttr(resourceName, "logs_bucket_uri", logsBucketUri),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttr(resourceName, "spark_version", "2.4.4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_invoke_runs", "test_invoke_run_submit", acctest.Optional, acctest.Update, invokeRunSubmitDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", acctest.Optional, acctest.Update, invokeRunSubmitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(datasourceName, "application_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "runs.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.application_id"),
				resource.TestCheckResourceAttr(datasourceName, "runs.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.data_read_in_bytes"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.data_written_in_bytes"),
				resource.TestCheckResourceAttr(datasourceName, "runs.0.display_name", "test_wordcount_runsubmit"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", acctest.Required, acctest.Create, invokeRunSubmitSingularDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(singularDatasourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_read_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_written_in_bytes"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "test_wordcount_runsubmit"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "file_uri"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "language"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logs_bucket_uri", logsBucketUri),
				resource.TestCheckResourceAttr(singularDatasourceName, "num_executors", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opc_request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner_user_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_duration_in_milliseconds"),
				resource.TestCheckResourceAttr(singularDatasourceName, "spark_version", "2.4.4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_ocpu"),
				resource.TestCheckResourceAttr(singularDatasourceName, "warehouse_bucket_uri", warehouseBucketUri),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataflowInvokeRunSubmit") {
		resource.AddTestSweepers("DataflowInvokeRunSubmit", &resource.Sweeper{
			Name:         "DataflowInvokeRun",
			Dependencies: acctest.DependencyGraph["invokeRun"],
			F:            sweepDataflowInvokeRunSubmitResource,
		})
	}
}

func sweepDataflowInvokeRunSubmitResource(compartment string) error {
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
