// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataFlowApplicationSubmitRequiredOnlyResource = dataFlowApplicationSubmitResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", acctest.Required, acctest.Create, dataFlowApplicationSubmitRepresentation)

	DataFlowApplicationSubmitResourceConfig = dataFlowApplicationSubmitResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", acctest.Optional, acctest.Update, dataFlowApplicationSubmitRepresentation)

	dataFlowApplicationSubmitSingularDataSourceRepresentation = map[string]interface{}{
		"application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_application.test_application_submit.id}`},
	}

	dataFlowApplicationSubmitDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `test_wordcount_app_submit`, Update: `test_wordcount_app_submit2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: dataFlowApplicationSubmitDataSourceFilterRepresentation}}
	dataFlowApplicationSubmitDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataflow_application.test_application_submit.id}`}},
	}

	dataFlowApplicationSubmitRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: `test_wordcount_app_submit`, Update: `test_wordcount_app_submit2`},
		"driver_shape":         acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"executor_shape":       acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"file_uri":             acctest.Representation{RepType: acctest.Required, Create: `${var.dataflow_file_uri}`, Update: `${var.dataflow_file_uri}`},
		"language":             acctest.Representation{RepType: acctest.Required, Create: `PYTHON`, Update: `PYTHON`},
		"num_executors":        acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"spark_version":        acctest.Representation{RepType: acctest.Required, Create: `2.4`, Update: `2.4.4`},
		"archive_uri":          acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_archive_uri}`, Update: `${var.dataflow_archive_uri}`},
		"description":          acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"execute":              acctest.Representation{RepType: acctest.Required, Create: `--conf spark.shuffle.io.maxRetries=10 ` + utils.GetEnvSettingWithBlankDefault("dataflow_file_uri") + ` arguments`, Update: `--conf spark.shuffle.io.maxRetries=10 ` + utils.GetEnvSettingWithBlankDefault("dataflow_file_uri") + ` arguments2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"logs_bucket_uri":      acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_logs_bucket_uri}`},
		"warehouse_bucket_uri": acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_warehouse_bucket_uri}`},
		"metastore_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.metastore_id}`},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForDataFlowResource},
	}

	dataFlowApplicationSubmitResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dataflow/default
func TestDataflowApplicationResource_SparkSubmit(t *testing.T) {
	httpreplay.SetScenario("TestDataflowApplicationResource_SparkSubmit")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	fileUri := utils.GetEnvSettingWithBlankDefault("dataflow_file_uri")
	fileUriUpdated := utils.GetEnvSettingWithBlankDefault("dataflow_file_uri_updated")
	fileUriVariableStr := fmt.Sprintf("variable \"dataflow_file_uri\" { default = \"%s\" }\n", fileUri)
	fileUriVariableStrUpdated := fmt.Sprintf("variable \"dataflow_file_uri_updated\" { default = \"%s\" }\n", fileUriUpdated)
	archiveUri := utils.GetEnvSettingWithBlankDefault("dataflow_archive_uri")
	archiveUriVariableStr := fmt.Sprintf("variable \"dataflow_archive_uri\" { default = \"%s\" }\n", archiveUri)

	logsBucketUri := utils.GetEnvSettingWithBlankDefault("dataflow_logs_bucket_uri")
	logsBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_logs_bucket_uri\" { default = \"%s\" }\n", logsBucketUri)
	warehouseBucketUri := utils.GetEnvSettingWithBlankDefault("dataflow_warehouse_bucket_uri")
	warehouseBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_warehouse_bucket_uri\" { default = \"%s\" }\n", warehouseBucketUri)
	resourceName := "oci_dataflow_application.test_application_submit"
	datasourceName := "data.oci_dataflow_applications.test_applications_submit"
	singularDatasourceName := "data.oci_dataflow_application.test_application_submit"

	metastoreId := utils.GetEnvSettingWithBlankDefault("metastore_id")
	metastoreIdVariableStr := fmt.Sprintf("variable \"metastore_id\" { default = \"%s\" }\n", metastoreId)

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDataflowApplicationDestroy, []resource.TestStep{
		// verify Create with execute only
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + dataFlowApplicationSubmitResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", acctest.Required, acctest.Create, dataFlowApplicationSubmitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_app_submit"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments"),
				resource.TestCheckResourceAttr(resourceName, "file_uri", fileUri),
				resource.TestCheckResourceAttr(resourceName, "language", "PYTHON"),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttr(resourceName, "spark_version", "2.4"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + dataFlowApplicationSubmitResourceDependencies,
		},
		// verify Create with execute, and other optionals
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + dataFlowApplicationSubmitResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
					"execute": acctest.Representation{RepType: acctest.Optional, Create: "--conf spark.shuffle.io.maxRetries=10 " + fileUri + " arguments"}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_app_submit"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "file_uri", fileUri),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "language", "PYTHON"),
				resource.TestCheckResourceAttr(resourceName, "logs_bucket_uri", logsBucketUri),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttr(resourceName, "spark_version", "2.4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataFlowApplicationResourceDependencies + fileUriVariableStr + archiveUriVariableStr + warehouseBucketUriVariableStr + logsBucketUriVariableStr + metastoreIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						"execute":        acctest.Representation{RepType: acctest.Optional, Create: "--conf spark.shuffle.io.maxRetries=10 " + fileUri + " arguments"},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "archive_uri"),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_app_submit"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "language", "PYTHON"),
				resource.TestCheckResourceAttrSet(resourceName, "logs_bucket_uri"),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttr(resourceName, "spark_version", "2.4"),
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
			Config: config + compartmentIdVariableStr + fileUriVariableStr + fileUriVariableStrUpdated + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataFlowApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", acctest.Optional, acctest.Update,
					dataFlowApplicationSubmitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_app_submit2"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments2"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "file_uri", fileUri),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "language", "PYTHON"),
				resource.TestCheckResourceAttr(resourceName, "logs_bucket_uri", logsBucketUri),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_applications", "test_applications_submit", acctest.Optional, acctest.Update, dataFlowApplicationSubmitDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + fileUriVariableStrUpdated + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + dataFlowApplicationSubmitResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", acctest.Optional, acctest.Update, dataFlowApplicationSubmitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "test_wordcount_app_submit2"),
				resource.TestCheckResourceAttr(datasourceName, "applications.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.display_name", "test_wordcount_app_submit2"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.language", "PYTHON"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.owner_principal_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.owner_user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", acctest.Required, acctest.Create, dataFlowApplicationSubmitSingularDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + fileUriVariableStrUpdated + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + dataFlowApplicationSubmitResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", acctest.Optional, acctest.Update, dataFlowApplicationSubmitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "archive_uri"),
				resource.TestCheckResourceAttr(singularDatasourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "test_wordcount_app_submit2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execute", "--conf spark.shuffle.io.maxRetries=10 "+fileUri+" arguments2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_uri", fileUri),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "language", "PYTHON"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logs_bucket_uri", logsBucketUri),
				resource.TestCheckResourceAttr(singularDatasourceName, "num_executors", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner_user_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "spark_version", "2.4.4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "warehouse_bucket_uri", warehouseBucketUri),
			),
		},
		// verify resource import
		{
			Config:                  config + DataFlowApplicationSubmitRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataflowApplicationSubmit") {
		resource.AddTestSweepers("DataflowApplicationSubmit", &resource.Sweeper{
			Name:         "DataflowApplicationSubmit",
			Dependencies: acctest.DependencyGraph["application"],
			F:            sweepDataflowApplicationResource,
		})
	}
}
