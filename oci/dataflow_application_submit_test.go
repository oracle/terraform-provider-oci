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
		generateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Required, Create, dataFlowApplicationSubmitRepresentation)

	DataFlowApplicationSubmitResourceConfig = dataFlowApplicationSubmitResourceDependencies +
		generateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Update, dataFlowApplicationSubmitRepresentation)

	dataFlowApplicationSubmitSingularDataSourceRepresentation = map[string]interface{}{
		"application_id": Representation{repType: Required, create: `${oci_dataflow_application.test_application_submit.id}`},
	}

	dataFlowApplicationSubmitDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `test_wordcount_app_submit`, update: `test_wordcount_app_submit2`},
		"filter":         RepresentationGroup{Required, dataFlowApplicationSubmitDataSourceFilterRepresentation}}
	dataFlowApplicationSubmitDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_dataflow_application.test_application_submit.id}`}},
	}

	dataFlowApplicationSubmitRepresentation = map[string]interface{}{
		"compartment_id":       Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":         Representation{repType: Required, create: `test_wordcount_app_submit`, update: `test_wordcount_app_submit2`},
		"driver_shape":         Representation{repType: Required, create: `VM.Standard2.1`},
		"executor_shape":       Representation{repType: Required, create: `VM.Standard2.1`},
		"file_uri":             Representation{repType: Required, create: `${var.dataflow_file_uri}`, update: `${var.dataflow_file_uri}`},
		"language":             Representation{repType: Required, create: `PYTHON`, update: `PYTHON`},
		"num_executors":        Representation{repType: Required, create: `1`, update: `2`},
		"spark_version":        Representation{repType: Required, create: `2.4`, update: `2.4.4`},
		"archive_uri":          Representation{repType: Optional, create: `${var.dataflow_archive_uri}`, update: `${var.dataflow_archive_uri}`},
		"defined_tags":         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":          Representation{repType: Optional, create: `description`, update: `description2`},
		"execute":              Representation{repType: Required, create: `--conf spark.shuffle.io.maxRetries=10 ` + getEnvSettingWithBlankDefault("dataflow_file_uri") + ` arguments`, update: `--conf spark.shuffle.io.maxRetries=10 ` + getEnvSettingWithBlankDefault("dataflow_file_uri") + ` arguments2`},
		"freeform_tags":        Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"logs_bucket_uri":      Representation{repType: Optional, create: `${var.dataflow_logs_bucket_uri}`},
		"warehouse_bucket_uri": Representation{repType: Optional, create: `${var.dataflow_warehouse_bucket_uri}`},
	}

	dataFlowApplicationSubmitResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestDataflowApplicationResource_SparkSubmit(t *testing.T) {
	httpreplay.SetScenario("TestDataflowApplicationResource_SparkSubmit")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	fileUri := getEnvSettingWithBlankDefault("dataflow_file_uri")
	fileUriUpdated := getEnvSettingWithBlankDefault("dataflow_file_uri_updated")
	fileUriVariableStr := fmt.Sprintf("variable \"dataflow_file_uri\" { default = \"%s\" }\n", fileUri)
	fileUriVariableStrUpdated := fmt.Sprintf("variable \"dataflow_file_uri_updated\" { default = \"%s\" }\n", fileUriUpdated)
	archiveUri := getEnvSettingWithBlankDefault("dataflow_archive_uri")
	archiveUriVariableStr := fmt.Sprintf("variable \"dataflow_archive_uri\" { default = \"%s\" }\n", archiveUri)

	logsBucketUri := getEnvSettingWithBlankDefault("dataflow_logs_bucket_uri")
	logsBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_logs_bucket_uri\" { default = \"%s\" }\n", logsBucketUri)
	warehouseBucketUri := getEnvSettingWithBlankDefault("dataflow_warehouse_bucket_uri")
	warehouseBucketUriVariableStr := fmt.Sprintf("variable \"dataflow_warehouse_bucket_uri\" { default = \"%s\" }\n", warehouseBucketUri)
	classNameUpdated := ""
	classNameStrUpdated := fmt.Sprintf("variable \"dataflow_class_name_updated\" { default = \"%s\" }\n", classNameUpdated)
	resourceName := "oci_dataflow_application.test_application_submit"
	datasourceName := "data.oci_dataflow_applications.test_applications_submit"
	singularDatasourceName := "data.oci_dataflow_application.test_application_submit"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDataflowApplicationDestroy,
		Steps: []resource.TestStep{
			// verify create with execute only
			{
				Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + dataFlowApplicationSubmitResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Required, Create, dataFlowApplicationSubmitRepresentation),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + dataFlowApplicationSubmitResourceDependencies,
			},
			// verify create with execute, and other optionals
			{
				Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + dataFlowApplicationSubmitResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Create, representationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
						"execute": Representation{repType: Optional, create: "--conf spark.shuffle.io.maxRetries=10 " + fileUri + " arguments"}})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
					resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + dataFlowApplicationResourceDependencies + fileUriVariableStr + archiveUriVariableStr + warehouseBucketUriVariableStr + logsBucketUriVariableStr +
					generateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Create,
						representationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
							"execute":        Representation{repType: Optional, create: "--conf spark.shuffle.io.maxRetries=10 " + fileUri + " arguments"},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "archive_uri"),
					resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
				Config: config + compartmentIdVariableStr + fileUriVariableStr + classNameStrUpdated + fileUriVariableStrUpdated + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + dataFlowApplicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Update,
						dataFlowApplicationSubmitRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "archive_uri", archiveUri),
					resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_dataflow_applications", "test_applications_submit", Optional, Update, dataFlowApplicationSubmitDataSourceRepresentation) +
					compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + fileUriVariableStrUpdated + logsBucketUriVariableStr + classNameStrUpdated + warehouseBucketUriVariableStr + dataFlowApplicationSubmitResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Update, representationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
						"class_name": Representation{repType: Optional, create: `${var.dataflow_class_name_updated}`},
					})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "test_wordcount_app_submit2"),
					resource.TestCheckResourceAttr(datasourceName, "applications.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "applications.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "applications.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Required, Create, dataFlowApplicationSubmitSingularDataSourceRepresentation) +
					compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + fileUriVariableStrUpdated + logsBucketUriVariableStr + classNameStrUpdated + warehouseBucketUriVariableStr + dataFlowApplicationSubmitResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Update, representationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
						"class_name": Representation{repType: Optional, create: `${var.dataflow_class_name_updated}`},
					})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "application_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "archive_uri"),
					resource.TestCheckResourceAttr(singularDatasourceName, "arguments.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "configuration.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
				Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + fileUriVariableStrUpdated + classNameStrUpdated + logsBucketUriVariableStr + warehouseBucketUriVariableStr + dataFlowApplicationSubmitResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_application", "test_application_submit", Optional, Update, representationCopyWithNewProperties(dataFlowApplicationSubmitRepresentation, map[string]interface{}{
						"class_name": Representation{repType: Optional, create: `${var.dataflow_class_name_updated}`},
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
		},
	})
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DataflowApplicationSubmit") {
		resource.AddTestSweepers("DataflowApplicationSubmit", &resource.Sweeper{
			Name:         "DataflowApplicationSubmit",
			Dependencies: DependencyGraph["application"],
			F:            sweepDataflowApplicationResource,
		})
	}
}
