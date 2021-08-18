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
	oci_dataflow "github.com/oracle/oci-go-sdk/v46/dataflow"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	InvokeRunSubmitRequiredOnlyResource = InvokeRunSubmitResourceDependencies +
		generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", Required, Create, invokeRunSubmitRepresentation)

	InvokeRunSubmitResourceConfig = InvokeRunSubmitResourceDependencies +
		generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", Optional, Update, invokeRunSubmitRepresentation)

	invokeRunSubmitSingularDataSourceRepresentation = map[string]interface{}{
		"run_id": Representation{repType: Required, create: `${oci_dataflow_invoke_run.test_invoke_run_submit.id}`},
	}

	invokeRunSubmitDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"filter":         RepresentationGroup{Required, invokeRunSubmitDataSourceFilterRepresentation}}
	invokeRunSubmitDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_dataflow_invoke_run.test_invoke_run_submit.id}`}},
	}

	invokeRunSubmitRepresentation = map[string]interface{}{
		"compartment_id":       Representation{repType: Required, create: `${var.compartment_id}`},
		"archive_uri":          Representation{repType: Optional, create: getEnvSettingWithBlankDefault("dataflow_archive_uri")},
		"defined_tags":         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         Representation{repType: Optional, create: `test_wordcount_runsubmit`},
		"driver_shape":         Representation{repType: Optional, create: `VM.Standard2.1`},
		"execute":              Representation{repType: Required, create: `--conf spark.shuffle.io.maxRetries=10 ` + getEnvSettingWithBlankDefault("dataflow_file_uri") + ` arguments`},
		"executor_shape":       Representation{repType: Optional, create: `VM.Standard2.1`},
		"freeform_tags":        Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"logs_bucket_uri":      Representation{repType: Optional, create: `${var.dataflow_logs_bucket_uri}`},
		"metastore_id":         Representation{repType: Optional, create: `${var.metastore_id}`},
		"num_executors":        Representation{repType: Optional, create: `1`},
		"spark_version":        Representation{repType: Optional, create: `2.4.4`},
		"warehouse_bucket_uri": Representation{repType: Optional, create: `${var.dataflow_warehouse_bucket_uri}`},
	}

	InvokeRunSubmitResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dataflow/default
func TestDataflowInvokeRunResource_SparkSubmit(t *testing.T) {
	httpreplay.SetScenario("TestDataflowInvokeRunResource_SparkSubmit")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	fileUri := getEnvSettingWithBlankDefault("dataflow_file_uri")
	fileUriVariableStr := fmt.Sprintf("variable \"dataflow_file_uri\" { default = \"%s\" }\n", fileUri)
	archiveUri := getEnvSettingWithBlankDefault("dataflow_archive_uri")
	archiveUriVariableStr := fmt.Sprintf("variable \"dataflow_archive_uri\" { default = \"%s\" }\n", archiveUri)
	logsBucketUri := getEnvSettingWithBlankDefault("dataflow_logs_bucket_uri")
	logsBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_logs_bucket_uri\" { default = \"%s\" }\n", logsBucketUri)
	warehouseBucketUri := getEnvSettingWithBlankDefault("dataflow_warehouse_bucket_uri")
	warehouseBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_warehouse_bucket_uri\" { default = \"%s\" }\n", warehouseBucketUri)
	metastoreId := getEnvSettingWithBlankDefault("metastore_id")
	metastoreIdVariableStr := fmt.Sprintf("variable \"metastore_id\" { default = \"%s\" }\n", metastoreId)

	resourceName := "oci_dataflow_invoke_run.test_invoke_run_submit"
	datasourceName := "data.oci_dataflow_invoke_runs.test_invoke_run_submit"
	singularDatasourceName := "data.oci_dataflow_invoke_run.test_invoke_run_submit"

	var resId, resId2 string

	ResourceTest(t, nil, []resource.TestStep{
		// verify create run with required execute
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceDependencies +
				generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", Required, Create, invokeRunSubmitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments"),
				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceDependencies,
		},
		// verify create with execute, display_name, spark_version
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceDependencies +
				generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", Optional, Create, invokeRunSubmitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
			PreConfig: waitTillCondition(testAccProvider, &resId, dataflowRunAvailableShouldWaitCondition, time.Duration(20*time.Minute),
				dataFlowInvokeRunFetchOperation, "dataflow", false),
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + InvokeRunSubmitResourceDependencies + warehouseBucketUriVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + metastoreIdVariableStr +
				generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", Optional, Create,
					representationCopyWithNewProperties(invokeRunSubmitRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceDependencies +
				generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", Optional, Update, invokeRunSubmitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
				generateDataSourceFromRepresentationMap("oci_dataflow_invoke_runs", "test_invoke_run_submit", Optional, Update, invokeRunSubmitDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceDependencies +
				generateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", Optional, Update, invokeRunSubmitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(datasourceName, "application_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "runs.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.application_id"),
				resource.TestCheckResourceAttr(datasourceName, "runs.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.data_read_in_bytes"),
				resource.TestCheckResourceAttrSet(datasourceName, "runs.0.data_written_in_bytes"),
				resource.TestCheckResourceAttr(datasourceName, "runs.0.defined_tags.%", "1"),
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
				generateDataSourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run_submit", Required, Create, invokeRunSubmitSingularDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + InvokeRunSubmitResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(singularDatasourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_read_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_written_in_bytes"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DataflowInvokeRunSubmit") {
		resource.AddTestSweepers("DataflowInvokeRunSubmit", &resource.Sweeper{
			Name:         "DataflowInvokeRun",
			Dependencies: DependencyGraph["invokeRun"],
			F:            sweepDataflowInvokeRunSubmitResource,
		})
	}
}

func sweepDataflowInvokeRunSubmitResource(compartment string) error {
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
