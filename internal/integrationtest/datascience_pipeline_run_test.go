// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	acctest "github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PipelineRunRequiredOnlyResource = PipelineRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Required, acctest.Create, pipelineRunRepresentation)

	PipelineRunResourceConfig = PipelineRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Optional, acctest.Update, pipelineRunRepresentation)

	pipelineRunSingularDataSourceRepresentation = map[string]interface{}{
		"pipeline_run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_pipeline_run.test_pipeline_run.id}`},
	}

	pipelineRunDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_pipeline.test_pipeline.created_by}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_pipeline_run.test_pipeline_run.id}`},
		"pipeline_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_pipeline.test_pipeline.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `SUCCEEDED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineRunDataSourceFilterRepresentation}}
	pipelineRunDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_pipeline_run.test_pipeline_run.id}`}},
	}

	// change projectId to optional after creating the new branches
	pipelineRunRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"pipeline_id":                    acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_pipeline.test_pipeline.id}`},
		"project_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"configuration_override_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineRunConfigurationOverrideDetailsRepresentation},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"delete_related_job_runs":        acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"parameters_override":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"parametersOverride": "parametersOverride"}},
		"infrastructure_configuration_override_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineRunInfrastructureConfigurationOverrideDetailsRepresentation},
		"log_configuration_override_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineRunLogConfigurationOverrideDetailsRepresentation},
		"step_override_details": []acctest.RepresentationGroup{
			{RepType: acctest.Optional, Group: pipelineRunStepOverrideDetailsRepresentation},
			{RepType: acctest.Optional, Group: pipelineRunStepOverrideDetailsRepresentation2},
		},
	}

	pipelineRunContainerRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"pipeline_id":                        acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_pipeline.test_pipeline.id}`},
		"project_id":                         acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"configuration_override_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineRunConfigurationOverrideDetailsRepresentation},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"delete_related_job_runs":            acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"log_configuration_override_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineRunLogConfigurationOverrideDetailsRepresentation},
		"step_override_details":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineRunStepOverrideDetailsContainerRepresentation},
	}

	pipelineRunDataflowRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"pipeline_id":                        acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_pipeline.test_pipeline.id}`},
		"project_id":                         acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"configuration_override_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineRunDataflowConfigurationOverrideDetailsRepresentation},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"delete_related_job_runs":            acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"log_configuration_override_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineRunLogConfigurationOverrideDetailsRepresentation},
		"step_override_details":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineRunStepOverrideDetailsDataflowRepresentation},
	}

	pipelineRunConfigurationOverrideDetailsRepresentation = map[string]interface{}{
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"command_line_arguments":     acctest.Representation{RepType: acctest.Optional, Create: `commandLineArgumentsOverriden`},
		"environment_variables":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"environmentVariablesOverriden": "environmentVariablesOverriden"}},
		"maximum_runtime_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DatasciencePipelineRunInfrastructureConfigurationOverrideDetailsRepresentation = map[string]interface{}{
		"block_storage_size_in_gbs":               acctest.Representation{RepType: acctest.Required, Create: `10`},
		"shape_name":                              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_shape.test_shape.name}`},
		"block_storage_size_in_gbs_parameterized": acctest.Representation{RepType: acctest.Optional, Create: `blockStorageSizeInGBsParameterized`},
		"shape_config_details":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineRunInfrastructureConfigurationOverrideDetailsShapeConfigDetailsRepresentation},
		"subnet_id":                               acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	pipelineRunDataflowConfigurationOverrideDetailsRepresentation = map[string]interface{}{
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"maximum_runtime_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `60`},
	}
	pipelineRunLogConfigurationOverrideDetailsRepresentation = map[string]interface{}{
		"enable_auto_log_creation": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"enable_logging":           acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"log_group_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_logging_log_group.terraform_test_custom_log_group.id}`},
		"log_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_logging_log.terraform_test_custom_log.id}`},
	}
	pipelineRunStepOverrideDetailsRepresentation = map[string]interface{}{
		"step_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineRunStepOverrideDetailsStepConfigurationDetailsRepresentation},
		"step_name":                  acctest.Representation{RepType: acctest.Required, Create: `stepName1`},
	}
	pipelineRunStepOverrideDetailsRepresentation2 = map[string]interface{}{
		"step_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineRunStepOverrideDetailsStepConfigurationDetailsRepresentation},
		"step_name":                  acctest.Representation{RepType: acctest.Required, Create: `stepName2`},
	}
	DatasciencePipelineRunStepOverrideDetailsRepresentation = map[string]interface{}{
		"step_configuration_details":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineRunStepOverrideDetailsStepConfigurationDetailsRepresentation},
		"step_name":                                     acctest.Representation{RepType: acctest.Required, Create: `stepName`},
		"step_container_configuration_details":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineRunStepOverrideDetailsStepConfigurationDetailsRepresentation},
		"step_dataflow_configuration_details":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineRunStepOverrideDetailsStepDataflowConfigurationDetailsRepresentation},
		"step_infrastructure_configuration_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineRunStepOverrideDetailsStepInfrastructureConfigurationDetailsRepresentation},
		"step_storage_mount_configuration_details_list": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineRunStepOverrideDetailsStepStorageMountConfigurationDetailsListRepresentation},
	}
	DatasciencePipelineRunInfrastructureConfigurationOverrideDetailsShapeConfigDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
	}
	DatasciencePipelineRunStorageMountConfigurationOverrideDetailsListRepresentation = map[string]interface{}{
		"destination_directory_name": acctest.Representation{RepType: acctest.Required, Create: `destinationDirectoryName`},
		"storage_type":               acctest.Representation{RepType: acctest.Required, Create: `FILE_STORAGE`},
		"bucket":                     acctest.Representation{RepType: acctest.Optional, Create: `bucket`},
		"destination_path":           acctest.Representation{RepType: acctest.Optional, Create: `destinationPath`},
		"export_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_export.test_export.id}`},
		"mount_target_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_mount_target.test_mount_target.id}`},
		"namespace":                  acctest.Representation{RepType: acctest.Optional, Create: `namespace`},
		"prefix":                     acctest.Representation{RepType: acctest.Optional, Create: `prefix`},
	}
	pipelineRunStepOverrideDetailsContainerRepresentation = map[string]interface{}{
		"step_configuration_details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineRunStepOverrideDetailsStepConfigurationDetailsRepresentation},
		"step_name":                            acctest.Representation{RepType: acctest.Required, Create: `stepNameContainer`},
		"step_container_configuration_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineRunStepOverrideDetailsContainerConfigurationDetailsRepresentation},
	}
	pipelineRunStepOverrideDetailsDataflowRepresentation = map[string]interface{}{
		"step_configuration_details":          acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineRunStepOverrideDetailsStepConfigurationDetailsRepresentationDataflow},
		"step_name":                           acctest.Representation{RepType: acctest.Required, Create: `stepNameDataflow`},
		"step_dataflow_configuration_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineRunStepOverrideDetailsStepDataflowConfigurationDetailsRepresentation},
	}
	pipelineRunStepOverrideDetailsStepConfigurationDetailsRepresentation = map[string]interface{}{
		"command_line_arguments":     acctest.Representation{RepType: acctest.Optional, Create: `commandLineArgumentsOverriden`},
		"environment_variables":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"environmentVariables": "environmentVariablesOverriden"}},
		"maximum_runtime_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	pipelineRunStepOverrideDetailsStepConfigurationDetailsRepresentationDataflow = map[string]interface{}{
		"maximum_runtime_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `60`},
	}
	DatasciencePipelineRunStepOverrideDetailsContainerConfigurationDetailsRepresentation = map[string]interface{}{
		"container_type":     acctest.Representation{RepType: acctest.Required, Create: `OCIR_CONTAINER`},
		"image":              acctest.Representation{RepType: acctest.Required, Create: `iad.ocir.io/idtlxnfdweil/byod/test-hello-world:1.0`},
		"cmd":                acctest.Representation{RepType: acctest.Optional, Create: []string{``}},
		"entrypoint":         acctest.Representation{RepType: acctest.Optional, Create: []string{``}},
		"image_digest":       acctest.Representation{RepType: acctest.Optional, Create: ``},
		"image_signature_id": acctest.Representation{RepType: acctest.Optional, Create: ``},
	}
	DatasciencePipelineRunStepOverrideDetailsStepDataflowConfigurationDetailsRepresentation = map[string]interface{}{
		"configuration":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"dataflow.auth": "3.1.0"}},
		"driver_shape":                  acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard.E5.Flex`},
		"driver_shape_config_details":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineStepDetailsStepDataflowConfigurationDetailsDriverShapeConfigDetailsRepresentation},
		"executor_shape":                acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard.E5.Flex`},
		"executor_shape_config_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineStepDetailsStepDataflowConfigurationDetailsExecutorShapeConfigDetailsRepresentation},
		"logs_bucket_uri":               acctest.Representation{RepType: acctest.Optional, Create: `oci://xuejuzha-test@idtlxnfdweil/`},
		"num_executors":                 acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"warehouse_bucket_uri":          acctest.Representation{RepType: acctest.Optional, Create: `oci://xuejuzha-test@idtlxnfdweil/`},
	}
	DatasciencePipelineRunStepOverrideDetailsStepInfrastructureConfigurationDetailsRepresentation = map[string]interface{}{
		"block_storage_size_in_gbs":               acctest.Representation{RepType: acctest.Required, Create: `10`},
		"shape_name":                              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_shape.test_shape.name}`},
		"block_storage_size_in_gbs_parameterized": acctest.Representation{RepType: acctest.Optional, Create: `blockStorageSizeInGBsParameterized`},
		"shape_config_details":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineRunStepOverrideDetailsStepInfrastructureConfigurationDetailsShapeConfigDetailsRepresentation},
		"subnet_id":                               acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	DatasciencePipelineRunStepOverrideDetailsStepStorageMountConfigurationDetailsListRepresentation = map[string]interface{}{
		"destination_directory_name": acctest.Representation{RepType: acctest.Required, Create: `destinationDirectoryName`},
		"storage_type":               acctest.Representation{RepType: acctest.Required, Create: `FILE_STORAGE`},
		"bucket":                     acctest.Representation{RepType: acctest.Optional, Create: `bucket`},
		"destination_path":           acctest.Representation{RepType: acctest.Optional, Create: `destinationPath`},
		"export_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_export.test_export.id}`},
		"mount_target_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_mount_target.test_mount_target.id}`},
		"namespace":                  acctest.Representation{RepType: acctest.Optional, Create: `namespace`},
		"prefix":                     acctest.Representation{RepType: acctest.Optional, Create: `prefix`},
	}
	DatasciencePipelineRunStepOverrideDetailsStepDataflowConfigurationDetailsDriverShapeConfigDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `8.0`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
	}
	DatasciencePipelineRunStepOverrideDetailsStepDataflowConfigurationDetailsExecutorShapeConfigDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `8.0`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
	}
	DatasciencePipelineRunStepOverrideDetailsStepInfrastructureConfigurationDetailsShapeConfigDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
	}

	// DatasciencePipelineRunResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", acctest.Required, acctest.Create, CoreShapeDataSourceRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Required, acctest.Create, DataflowApplicationRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Required, acctest.Create, DatascienceJobRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Required, acctest.Create, DatasciencePipelineRepresentation)

	PipelineRunResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Required, acctest.Create, pipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "terraform_test_custom_log_group", acctest.Required, acctest.Create, pipelineLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "terraform_test_custom_log", acctest.Required, acctest.Create, pipelineLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation)

	PipelineRunContainerResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Optional, acctest.Create, pipelineRepresentationContainer) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "terraform_test_custom_log_group", acctest.Required, acctest.Create, pipelineLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "terraform_test_custom_log", acctest.Required, acctest.Create, pipelineLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation)

	PipelineRunDataflowResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Optional, acctest.Create, pipelineRepresentationDataflow) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "terraform_test_custom_log_group", acctest.Required, acctest.Create, pipelineLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "terraform_test_custom_log", acctest.Required, acctest.Create, pipelineLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation)
)

// issue-routing-tag: datascience/default
func TestDatasciencePipelineRunResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatasciencePipelineRunResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_pipeline_run.test_pipeline_run"
	datasourceName := "data.oci_datascience_pipeline_runs.test_pipeline_runs"
	singularDatasourceName := "data.oci_datascience_pipeline_run.test_pipeline_run"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PipelineRunContainerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Optional, acctest.Create, pipelineRunRepresentation), "datascience", "pipelineRun", t)

	acctest.ResourceTest(t, testAccCheckDatasciencePipelineRunDestroy, []resource.TestStep{
		// Step 0 - Verify Create Pipeline Run with Container
		//{
		//	Config: config + compartmentIdVariableStr + PipelineRunDataflowResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Optional, acctest.Create, pipelineRunDataflowRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		//		resource.TestCheckResourceAttrSet(resourceName, "pipeline_id"),
		//		resource.TestCheckResourceAttrSet(resourceName, "project_id"),
		//		resource.TestCheckResourceAttrSet(resourceName, "id"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_name", "stepNameDataflow"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.driver_shape", "VM.Standard.E5.Flex"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.driver_shape_config_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.driver_shape_config_details.0.memory_in_gbs", "14"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.driver_shape_config_details.0.ocpus", "2"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.executor_shape", "VM.Standard.E5.Flex"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.executor_shape_config_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.executor_shape_config_details.0.memory_in_gbs", "14"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.executor_shape_config_details.0.ocpus", "2"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.logs_bucket_uri", "oci://xuejuzha-test@idtlxnfdweil/"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.num_executors", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_dataflow_configuration_details.0.warehouse_bucket_uri", "oci://xuejuzha-test@idtlxnfdweil/"),
		//
		//		func(s *terraform.State) (err error) {
		//			resId, err = acctest.FromInstanceState(s, resourceName, "id")
		//			return err
		//		},
		//	),
		//},
		//{
		//	Config: config + compartmentIdVariableStr + PipelineRunContainerResourceDependencies, // current pipeline state = ACCEPTED and DELETE after SUCCEEDED/CANCELED/FAILED
		//},
		//{
		//	Config: config + compartmentIdVariableStr + PipelineRunContainerResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Optional, acctest.Create, pipelineRunContainerRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		//		resource.TestCheckResourceAttrSet(resourceName, "pipeline_id"),
		//		resource.TestCheckResourceAttrSet(resourceName, "project_id"),
		//		resource.TestCheckResourceAttrSet(resourceName, "id"),
		//		resource.TestCheckResourceAttr(resourceName, "configuration_override_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.command_line_arguments", "commandLineArgumentsOverriden"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.environment_variables.%", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.maximum_runtime_in_minutes", "10"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_name", "stepNameContainer"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_container_configuration_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_container_configuration_details.0.container_type", "OCIR_CONTAINER"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_container_configuration_details.0.image", "iad.ocir.io/idtlxnfdweil/byod/test-hello-world:1.0"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_container_configuration_details.0.cmd.#", "0"),
		//		resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_container_configuration_details.0.entrypoint.#", "0"),
		//
		//		func(s *terraform.State) (err error) {
		//			resId, err = acctest.FromInstanceState(s, resourceName, "id")
		//			return err
		//		},
		//	),
		//},
		//// Step 1 - delete before next Create
		//{
		//	Config: config + compartmentIdVariableStr + PipelineRunContainerResourceDependencies, // current pipeline state = ACCEPTED and DELETE after SUCCEEDED/CANCELED/FAILED
		//},
		// Step 2 - verify Create
		{
			Config: config + compartmentIdVariableStr + PipelineRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Required, acctest.Create, pipelineRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "pipeline_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// Step 3 - delete before next Create
		{
			Config: config + compartmentIdVariableStr + PipelineRunResourceDependencies, // current pipeline state = ACCEPTED and DELETE after SUCCEEDED/CANCELED/FAILED
		},
		// Step 4 - verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PipelineRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Optional, acctest.Create, pipelineRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.command_line_arguments", "commandLineArgumentsOverriden"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.block_storage_size_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.shape_config_details.0.memory_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.shape_config_details.0.ocpus", "1.0"),
				resource.TestCheckResourceAttrSet(resourceName, "infrastructure_configuration_override_details.0.shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "infrastructure_configuration_override_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_override_details.0.enable_auto_log_creation", "false"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_override_details.0.enable_logging", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_override_details.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_override_details.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "pipeline_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.command_line_arguments", "commandLineArgumentsOverriden"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_accepted"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_name", "stepName1"),
				resource.TestCheckResourceAttr(resourceName, "step_runs.#", "2"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// Step 5 - verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PipelineRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(pipelineRunRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.command_line_arguments", "commandLineArgumentsOverriden"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.block_storage_size_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.shape_config_details.0.memory_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.shape_config_details.0.ocpus", "1.0"),
				resource.TestCheckResourceAttrSet(resourceName, "infrastructure_configuration_override_details.0.shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "infrastructure_configuration_override_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_override_details.0.enable_auto_log_creation", "false"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_override_details.0.enable_logging", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_override_details.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_override_details.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "pipeline_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.command_line_arguments", "commandLineArgumentsOverriden"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.environment_variables.%", "1"),
				//resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "step_runs.#", "2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Step 6 - verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + PipelineRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Optional, acctest.Update, pipelineRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.command_line_arguments", "commandLineArgumentsOverriden"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration_override_details.0.type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.block_storage_size_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.shape_config_details.0.memory_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_override_details.0.shape_config_details.0.ocpus", "1.0"),
				resource.TestCheckResourceAttrSet(resourceName, "infrastructure_configuration_override_details.0.shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "infrastructure_configuration_override_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_override_details.0.enable_auto_log_creation", "false"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_override_details.0.enable_logging", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_override_details.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_override_details.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "pipeline_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.command_line_arguments", "commandLineArgumentsOverriden"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_accepted"),
				resource.TestCheckResourceAttr(resourceName, "step_override_details.0.step_name", "stepName1"),
				resource.TestCheckResourceAttr(resourceName, "step_runs.#", "2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// Step 7 - verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_pipeline_runs", "test_pipeline_runs", acctest.Optional, acctest.Update, pipelineRunDataSourceRepresentation) +
				compartmentIdVariableStr + PipelineRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Optional, acctest.Update, pipelineRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "created_by"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),
				resource.TestCheckResourceAttr(datasourceName, "pipeline_runs.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pipeline_runs.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_runs.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "pipeline_runs.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "pipeline_runs.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_runs.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_runs.0.pipeline_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_runs.0.project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_runs.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "pipeline_runs.0.system_tags.#", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_runs.0.time_accepted"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_runs.0.time_finished"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_runs.0.time_started"),
			),
		},
		// Step 8 - verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_pipeline_run", "test_pipeline_run", acctest.Required, acctest.Create, pipelineRunSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PipelineRunResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pipeline_run_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_override_details.0.command_line_arguments", "commandLineArgumentsOverriden"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_override_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_override_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_override_details.0.type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_configuration_override_details.0.block_storage_size_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_configuration_override_details.0.shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_configuration_override_details.0.shape_config_details.0.memory_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_configuration_override_details.0.shape_config_details.0.ocpus", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_configuration_override_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_configuration_override_details.0.enable_auto_log_creation", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_configuration_override_details.0.enable_logging", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_override_details.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_override_details.0.step_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_override_details.0.step_configuration_details.0.command_line_arguments", "commandLineArgumentsOverriden"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_override_details.0.step_configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_override_details.0.step_configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_override_details.0.step_name", "stepName1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_runs.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
			),
		},
		// Step 9 - verify resource import
		{
			Config:            config + PipelineRunRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"delete_related_job_runs",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatasciencePipelineRunDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_pipeline_run" {
			noResourceFound = false
			request := oci_datascience.GetPipelineRunRequest{}

			tmp := rs.Primary.ID
			request.PipelineRunId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetPipelineRun(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.PipelineRunLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.PipelineRun.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.PipelineRun.LifecycleState)
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
	if !acctest.InSweeperExcludeList("DatasciencePipelineRun") {
		resource.AddTestSweepers("DatasciencePipelineRun", &resource.Sweeper{
			Name:         "DatasciencePipelineRun",
			Dependencies: acctest.DependencyGraph["pipelineRun"],
			F:            sweepDatasciencePipelineRunResource,
		})
	}
}

func sweepDatasciencePipelineRunResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	pipelineRunIds, err := getPipelineRunIds(compartment)
	if err != nil {
		return err
	}
	for _, pipelineRunId := range pipelineRunIds {
		if ok := acctest.SweeperDefaultResourceId[pipelineRunId]; !ok {
			deletePipelineRunRequest := oci_datascience.DeletePipelineRunRequest{}

			deletePipelineRunRequest.PipelineRunId = &pipelineRunId

			*deletePipelineRunRequest.DeleteRelatedJobRuns = true

			deletePipelineRunRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeletePipelineRun(context.Background(), deletePipelineRunRequest)
			if error != nil {
				fmt.Printf("Error deleting PipelineRun %s %s, It is possible that the resource is already deleted. Please verify manually \n", pipelineRunId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &pipelineRunId, pipelineRunSweepWaitCondition, time.Duration(3*time.Minute),
				pipelineRunSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getPipelineRunIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PipelineRunId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listPipelineRunsRequest := oci_datascience.ListPipelineRunsRequest{}
	listPipelineRunsRequest.CompartmentId = &compartmentId
	listPipelineRunsRequest.LifecycleState = oci_datascience.ListPipelineRunsLifecycleStateSucceeded
	listPipelineRunsResponse, err := dataScienceClient.ListPipelineRuns(context.Background(), listPipelineRunsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PipelineRun list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, pipelineRun := range listPipelineRunsResponse.Items {
		id := *pipelineRun.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PipelineRunId", id)
	}
	return resourceIds, nil
}

func pipelineRunSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if pipelineRunResponse, ok := response.Response.(oci_datascience.GetPipelineRunResponse); ok {
		return pipelineRunResponse.PipelineRun.LifecycleState != oci_datascience.PipelineRunLifecycleStateDeleted
	}
	return false
}

func pipelineRunSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetPipelineRun(context.Background(), oci_datascience.GetPipelineRunRequest{
		PipelineRunId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
