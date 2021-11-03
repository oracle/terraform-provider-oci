// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DataFlowApplicationSubmitRequiredOnlyResource = dataFlowApplicationSubmitResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Required, Create, dataFlowApplicationSubmitRepresentation)

	DataFlowApplicationSubmitResourceConfig = dataFlowApplicationSubmitResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Update, dataFlowApplicationSubmitRepresentation)

	dataFlowApplicationSubmitSingularDataSourceRepresentation = map[string]interface{}{
		"application_id": Representation{RepType: Required, Create: `${oci_dataflow_application.test_application_submit.id}`},
	}

	dataFlowApplicationSubmitDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `test_wordcount_app_submit`, Update: `test_wordcount_app_submit2`},
		"filter":         RepresentationGroup{Required, dataFlowApplicationSubmitDataSourceFilterRepresentation}}
	dataFlowApplicationSubmitDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_dataflow_application.test_application_submit.id}`}},
	}

	dataFlowApplicationSubmitRepresentation = map[string]interface{}{
		"compartment_id":       Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":         Representation{RepType: Required, Create: `test_wordcount_app_submit`, Update: `test_wordcount_app_submit2`},
		"driver_shape":         Representation{RepType: Required, Create: `VM.Standard2.1`},
		"executor_shape":       Representation{RepType: Required, Create: `VM.Standard2.1`},
		"file_uri":             Representation{RepType: Required, Create: `${var.dataflow_file_uri}`, Update: `${var.dataflow_file_uri}`},
		"language":             Representation{RepType: Required, Create: `PYTHON`, Update: `PYTHON`},
		"num_executors":        Representation{RepType: Required, Create: `1`, Update: `2`},
		"spark_version":        Representation{RepType: Required, Create: `2.4`, Update: `2.4.4`},
		"archive_uri":          Representation{RepType: Optional, Create: `${var.dataflow_archive_uri}`, Update: `${var.dataflow_archive_uri}`},
		"defined_tags":         Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":          Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"execute":              Representation{RepType: Required, Create: `--conf spark.shuffle.io.maxRetries=10 ` + GetEnvSettingWithBlankDefault("dataflow_file_uri") + ` arguments`, Update: `--conf spark.shuffle.io.maxRetries=10 ` + GetEnvSettingWithBlankDefault("dataflow_file_uri") + ` arguments2`},
		"freeform_tags":        Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"logs_bucket_uri":      Representation{RepType: Optional, Create: `${var.dataflow_logs_bucket_uri}`},
		"warehouse_bucket_uri": Representation{RepType: Optional, Create: `${var.dataflow_warehouse_bucket_uri}`},
		"metastore_id":         Representation{RepType: Optional, Create: `${var.metastore_id}`},
	}

	dataFlowApplicationSubmitResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, SubnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dataflow/default
func TestDataflowApplicationResource_SparkSubmit(t *testing.T) {
	httpreplay.SetScenario("TestDataflowApplicationResource_SparkSubmit")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	fileUri := GetEnvSettingWithBlankDefault("dataflow_file_uri")
	fileUriUpdated := GetEnvSettingWithBlankDefault("dataflow_file_uri_updated")
	fileUriVariableStr := fmt.Sprintf("variable \"dataflow_file_uri\" { default = \"%s\" }\n", fileUri)
	fileUriVariableStrUpdated := fmt.Sprintf("variable \"dataflow_file_uri_updated\" { default = \"%s\" }\n", fileUriUpdated)
	archiveUri := GetEnvSettingWithBlankDefault("dataflow_archive_uri")
	archiveUriVariableStr := fmt.Sprintf("variable \"dataflow_archive_uri\" { default = \"%s\" }\n", archiveUri)

	logsBucketUri := GetEnvSettingWithBlankDefault("dataflow_logs_bucket_uri")
	logsBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_logs_bucket_uri\" { default = \"%s\" }\n", logsBucketUri)
	warehouseBucketUri := GetEnvSettingWithBlankDefault("dataflow_warehouse_bucket_uri")
	warehouseBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_warehouse_bucket_uri\" { default = \"%s\" }\n", warehouseBucketUri)
	classNameUpdated := ""
	classNameStrUpdated := fmt.Sprintf("variable \"dataflow_class_name_updated\" { default = \"%s\" }\n", classNameUpdated)
	resourceName := "oci_dataflow_application.test_application_submit"
	datasourceName := "data.oci_dataflow_applications.test_applications_submit"
	singularDatasourceName := "data.oci_dataflow_application.test_application_submit"

	metastoreId := GetEnvSettingWithBlankDefault("metastore_id")
	metastoreIdVariableStr := fmt.Sprintf("variable \"metastore_id\" { default = \"%s\" }\n", metastoreId)

	var resId, resId2 string

	ResourceTest(t, testAccCheckDataflowApplicationDestroy, []resource.TestStep{
		// verify Create with execute only
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + dataFlowApplicationSubmitResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Required, Create, dataFlowApplicationSubmitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Create, RepresentationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
					"execute": Representation{RepType: Optional, Create: "--conf spark.shuffle.io.maxRetries=10 " + fileUri + " arguments"}})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + dataFlowApplicationResourceDependencies + fileUriVariableStr + archiveUriVariableStr + warehouseBucketUriVariableStr + logsBucketUriVariableStr + metastoreIdVariableStr +
				GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Create,
					RepresentationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
						"execute":        Representation{RepType: Optional, Create: "--conf spark.shuffle.io.maxRetries=10 " + fileUri + " arguments"},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + classNameStrUpdated + fileUriVariableStrUpdated + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + dataFlowApplicationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Update,
					dataFlowApplicationSubmitRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_dataflow_applications", "test_applications_submit", Optional, Update, dataFlowApplicationSubmitDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + fileUriVariableStrUpdated + logsBucketUriVariableStr + classNameStrUpdated + warehouseBucketUriVariableStr + metastoreIdVariableStr + dataFlowApplicationSubmitResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Update, RepresentationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
					"class_name": Representation{RepType: Optional, Create: `${var.dataflow_class_name_updated}`},
				})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Required, Create, dataFlowApplicationSubmitSingularDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + fileUriVariableStrUpdated + logsBucketUriVariableStr + classNameStrUpdated + warehouseBucketUriVariableStr + metastoreIdVariableStr + dataFlowApplicationSubmitResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Update, RepresentationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
					"class_name": Representation{RepType: Optional, Create: `${var.dataflow_class_name_updated}`},
				})),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + fileUriVariableStrUpdated + classNameStrUpdated + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + dataFlowApplicationSubmitResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Update, RepresentationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
					"class_name": Representation{RepType: Optional, Create: `${var.dataflow_class_name_updated}`},
				})),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func init() {
	if DependencyGraph == nil {
		InitDependencyGraph()
	}
	if !InSweeperExcludeList("DataflowApplicationSubmit") {
		resource.AddTestSweepers("DataflowApplicationSubmit", &resource.Sweeper{
			Name:         "DataflowApplicationSubmit",
			Dependencies: DependencyGraph["application"],
			F:            sweepDataflowApplicationResource,
		})
	}
}
