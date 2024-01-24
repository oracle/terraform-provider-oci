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
	DataFlowApplicationRequiredOnlyResource = DataFlowApplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Required, acctest.Create, DataflowApplicationRepresentation)

	DataFlowApplicationResourceConfig = DataFlowApplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Optional, acctest.Update, DataflowApplicationRepresentation)

	DataflowDataflowApplicationSingularDataSourceRepresentation = map[string]interface{}{
		"application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_application.test_application.id}`},
	}

	DataflowDataflowApplicationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `test_wordcount_app`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowApplicationDataSourceFilterRepresentation}}
	DataflowApplicationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataflow_application.test_application.id}`}},
	}
	DataflowApplicationApplicationLogConfigRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
		"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_log.id}`},
	}

	DataflowApplicationRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `test_wordcount_app`, Update: `displayName2`},
		"driver_shape":           acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`, Update: `VM.Standard2.2`},
		"executor_shape":         acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`, Update: `VM.Standard2.2`},
		"file_uri":               acctest.Representation{RepType: acctest.Required, Create: `${var.dataflow_file_uri}`, Update: `${var.dataflow_file_uri_updated}`},
		"language":               acctest.Representation{RepType: acctest.Required, Create: `PYTHON`, Update: `SCALA`},
		"num_executors":          acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"spark_version":          acctest.Representation{RepType: acctest.Required, Create: `2.4`, Update: `2.4.4`},
		"application_log_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowApplicationApplicationLogConfigRepresentation},
		"archive_uri":            acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_archive_uri}`},
		"arguments":              acctest.Representation{RepType: acctest.Optional, Create: []string{`arguments`}, Update: []string{`arguments2`}},
		"configuration":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"configuration": "configuration"}, Update: map[string]string{"configuration2": "configuration2"}},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"driver_shape_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowApplicationDriverShapeConfigRepresentation},
		"executor_shape_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowApplicationExecutorShapeConfigRepresentation},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"logs_bucket_uri":        acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_logs_bucket_uri}`},
		"metastore_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.metastore_id}`},
		"parameters":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowApplicationParametersRepresentation},
		"private_endpoint_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_dataflow_private_endpoint.test_private_endpoint.id}`},
		"pool_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_dataflow_pool.test_pool.id}`},
		"type":                   acctest.Representation{RepType: acctest.Optional, Create: `BATCH`},
		"warehouse_bucket_uri":   acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_warehouse_bucket_uri}`},
		"lifecycle":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForDataFlowResource},
	}

	ignoreDefinedTagsChangesForDataFlowResource = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"defined_tags"}},
	}

	DataflowApplicationDriverShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `15`, Update: `30`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	DataflowApplicationExecutorShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `15`, Update: `30`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	DataflowApplicationParametersRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	DataflowApplicationLogGroupRepresentation = acctest.RepresentationCopyWithNewProperties(LoggingLogGroupRepresentation, map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `terraform_test_custom_log_group`},
	})

	DataflowApplicationLogRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `terraform_test_custom_log`},
		"log_group_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
		"log_type":           acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"retention_duration": acctest.Representation{RepType: acctest.Optional, Create: `30`},
	}

	DataFlowApplicationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Required, acctest.Create, DataflowPoolRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, DataflowApplicationLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, DataflowApplicationLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: dataflow/default
func TestDataflowApplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowApplicationResource_basic")
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
	classNameUpdated := utils.GetEnvSettingWithBlankDefault("dataflow_class_name_updated")
	classNameStrUpdated := fmt.Sprintf("variable \"dataflow_class_name_updated\" { default = \"%s\" }\n", classNameUpdated)
	resourceName := "oci_dataflow_application.test_application"
	datasourceName := "data.oci_dataflow_applications.test_applications"
	singularDatasourceName := "data.oci_dataflow_application.test_application"

	metastoreId := utils.GetEnvSettingWithBlankDefault("metastore_id")
	metastoreIdVariableStr := fmt.Sprintf("variable \"metastore_id\" { default = \"%s\" }\n", metastoreId)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataFlowApplicationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Optional, acctest.Create, DataflowApplicationRepresentation), "dataflow", "application", t)

	acctest.ResourceTest(t, testAccCheckDataflowApplicationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataFlowApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Required, acctest.Create, DataflowApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_app"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.1"),
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
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataFlowApplicationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataFlowApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Optional, acctest.Create, DataflowApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "application_log_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "archive_uri"),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_app"),
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
				resource.TestCheckResourceAttr(resourceName, "language", "PYTHON"),
				resource.TestCheckResourceAttr(resourceName, "logs_bucket_uri", logsBucketUri),
				resource.TestCheckResourceAttr(resourceName, "metastore_id", metastoreId),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "spark_version", "2.4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "BATCH"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataFlowApplicationResourceDependencies + fileUriVariableStr + archiveUriVariableStr + warehouseBucketUriVariableStr + logsBucketUriVariableStr + metastoreIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataflowApplicationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "application_log_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "archive_uri"),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test_wordcount_app"),
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
				resource.TestCheckResourceAttr(resourceName, "language", "PYTHON"),
				resource.TestCheckResourceAttrSet(resourceName, "logs_bucket_uri"),
				resource.TestCheckResourceAttr(resourceName, "metastore_id", metastoreId),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "value"),
				resource.TestCheckResourceAttrSet(resourceName, "pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "spark_version", "2.4"),
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
			Config: config + compartmentIdVariableStr + fileUriVariableStr + classNameStrUpdated + fileUriVariableStrUpdated + archiveUriVariableStr + logsBucketUriVariableStr + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataFlowApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataflowApplicationRepresentation, map[string]interface{}{
						"class_name": acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_class_name_updated}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "application_log_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "application_log_config.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "archive_uri"),
				resource.TestCheckResourceAttr(resourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.memory_in_gbs", "30"),
				resource.TestCheckResourceAttr(resourceName, "driver_shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.memory_in_gbs", "30"),
				resource.TestCheckResourceAttr(resourceName, "executor_shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "file_uri"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "language", "SCALA"),
				resource.TestCheckResourceAttr(resourceName, "logs_bucket_uri", logsBucketUri),
				resource.TestCheckResourceAttr(resourceName, "metastore_id", metastoreId),
				resource.TestCheckResourceAttr(resourceName, "num_executors", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "spark_version", "2.4.4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "BATCH"),
				resource.TestCheckResourceAttr(resourceName, "warehouse_bucket_uri", warehouseBucketUri),
				resource.TestCheckResourceAttr(resourceName, "class_name", classNameUpdated),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_applications", "test_applications", acctest.Optional, acctest.Update, DataflowDataflowApplicationDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + fileUriVariableStrUpdated + logsBucketUriVariableStr + classNameStrUpdated + warehouseBucketUriVariableStr + metastoreIdVariableStr +
				DataFlowApplicationResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(DataflowApplicationRepresentation, map[string]interface{}{
				"class_name": acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_class_name_updated}`},
			})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "applications.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.language", "SCALA"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.owner_principal_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.owner_user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.pool_id"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.spark_version", "2.4.4"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "applications.0.time_updated"),
				resource.TestCheckResourceAttr(datasourceName, "applications.0.type", "BATCH"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Required, acctest.Create, DataflowDataflowApplicationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + archiveUriVariableStr + fileUriVariableStrUpdated + logsBucketUriVariableStr + classNameStrUpdated + warehouseBucketUriVariableStr + metastoreIdVariableStr + DataFlowApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(DataflowApplicationRepresentation, map[string]interface{}{
					"class_name": acctest.Representation{RepType: acctest.Optional, Create: `${var.dataflow_class_name_updated}`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "application_log_config.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "archive_uri"),
				resource.TestCheckResourceAttr(singularDatasourceName, "arguments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape_config.0.memory_in_gbs", "30"),
				resource.TestCheckResourceAttr(singularDatasourceName, "driver_shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape", "VM.Standard2.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape_config.0.memory_in_gbs", "30"),
				resource.TestCheckResourceAttr(singularDatasourceName, "executor_shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "file_uri"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "language", "SCALA"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logs_bucket_uri", logsBucketUri),
				resource.TestCheckResourceAttr(singularDatasourceName, "metastore_id", metastoreId),
				resource.TestCheckResourceAttr(singularDatasourceName, "num_executors", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner_user_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parameters.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "spark_version", "2.4.4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "BATCH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "warehouse_bucket_uri", warehouseBucketUri),
			),
		},
		// verify resource import
		{
			Config:                  config + DataFlowApplicationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataflowApplicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataFlowClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataflow_application" {
			noResourceFound = false
			request := oci_dataflow.GetApplicationRequest{}

			tmp := rs.Primary.ID
			request.ApplicationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")

			response, err := client.GetApplication(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dataflow.ApplicationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataflowApplication") {
		resource.AddTestSweepers("DataflowApplication", &resource.Sweeper{
			Name:         "DataflowApplication",
			Dependencies: acctest.DependencyGraph["application"],
			F:            sweepDataflowApplicationResource,
		})
	}
}

func sweepDataflowApplicationResource(compartment string) error {
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()
	applicationIds, err := getDataflowApplicationIds(compartment)
	if err != nil {
		return err
	}
	for _, applicationId := range applicationIds {
		if ok := acctest.SweeperDefaultResourceId[applicationId]; !ok {
			deleteApplicationRequest := oci_dataflow.DeleteApplicationRequest{}

			deleteApplicationRequest.ApplicationId = &applicationId

			deleteApplicationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")
			_, error := dataFlowClient.DeleteApplication(context.Background(), deleteApplicationRequest)
			if error != nil {
				fmt.Printf("Error deleting Application %s %s, It is possible that the resource is already deleted. Please verify manually \n", applicationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &applicationId, DataflowApplicationSweepWaitCondition, time.Duration(3*time.Minute),
				DataflowApplicationSweepResponseFetchOperation, "dataflow", true)
		}
	}
	return nil
}

func getDataflowApplicationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ApplicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()

	listApplicationsRequest := oci_dataflow.ListApplicationsRequest{}
	listApplicationsRequest.CompartmentId = &compartmentId
	listApplicationsResponse, err := dataFlowClient.ListApplications(context.Background(), listApplicationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Application list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, application := range listApplicationsResponse.Items {
		id := *application.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ApplicationId", id)
	}
	return resourceIds, nil
}

func DataflowApplicationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if applicationResponse, ok := response.Response.(oci_dataflow.GetApplicationResponse); ok {
		return applicationResponse.LifecycleState != oci_dataflow.ApplicationLifecycleStateDeleted
	}
	return false
}

func DataflowApplicationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataFlowClient().GetApplication(context.Background(), oci_dataflow.GetApplicationRequest{
		ApplicationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
