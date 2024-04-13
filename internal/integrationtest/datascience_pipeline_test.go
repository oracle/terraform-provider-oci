// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	acctest "github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PipelineRequiredOnlyResource = PipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Required, acctest.Create, pipelineRepresentation)

	PipelineResourceConfig = PipelineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Optional, acctest.Update, pipelineRepresentation)

	pipelineSingularDataSourceRepresentation = map[string]interface{}{
		"pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_pipeline.test_pipeline.id}`},
	}

	pipelineDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_pipeline.test_pipeline.created_by}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_pipeline.test_pipeline.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineDataSourceFilterRepresentation}}
	pipelineDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_pipeline.test_pipeline.id}`}},
	}

	pipelineRepresentation = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"step_details":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineStepDetailsRepresentationCS},
		"configuration_details":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineConfigurationDetailsRepresentation},
		"description":                          acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                         acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"infrastructure_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineInfrastructureConfigurationDetailsRepresentation},
		"log_configuration_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineLogConfigurationDetailsRepresentation},
		"step_artifact":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineStepArtifactRepresentation},
		"delete_related_pipeline_runs":         acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
	}

	pipelineRepresentationContainer = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"step_details":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: pipelineStepDetailsRepresentationContainer},
		"configuration_details":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineConfigurationDetailsRepresentation},
		"description":                          acctest.Representation{RepType: acctest.Optional, Create: `descriptionContainerPipeline`, Update: `description2`},
		"display_name":                         acctest.Representation{RepType: acctest.Optional, Create: `displayNameContainerPipeline`, Update: `displayName2`},
		"freeform_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"infrastructure_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatasciencePipelineInfrastructureConfigurationDetailsRepresentation},
		"log_configuration_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineLogConfigurationDetailsRepresentation},
		"delete_related_pipeline_runs":         acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
	}
	pipelineStepArtifactRepresentation = map[string]interface{}{
		"step_name":                    acctest.Representation{RepType: acctest.Required, Create: `stepName`},
		"pipeline_step_artifact":       acctest.Representation{RepType: acctest.Required, Create: `../../examples/datascience/job-artifact.py`},
		"artifact_content_length":      acctest.Representation{RepType: acctest.Required, Create: `1380`}, // wc -c job-artifact.py
		"artifact_content_disposition": acctest.Representation{RepType: acctest.Required, Create: `attachment; filename=job-artifact.py`},
	}
	// ML_JOB type has been tested locally and is not included in this for brevity as it does not
	// contain most of the optional parameters which are ALL tested in the CUSTOM_SCRIPT path
	pipelineStepDetailsRepresentation = map[string]interface{}{
		"step_name":                  acctest.Representation{RepType: acctest.Required, Create: `stepName`},
		"step_type":                  acctest.Representation{RepType: acctest.Required, Create: `ML_JOB`},
		"depends_on":                 acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"job_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_job.test_job.id}`},
		"step_configuration_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineStepDetailsStepConfigurationDetailsRepresentation},
	}

	pipelineStepDetailsRepresentationCS = map[string]interface{}{
		"step_name":                  acctest.Representation{RepType: acctest.Required, Create: `stepName`},
		"step_type":                  acctest.Representation{RepType: acctest.Required, Create: `CUSTOM_SCRIPT`},
		"depends_on":                 acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"step_configuration_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineStepDetailsStepConfigurationDetailsRepresentation},
		"step_infrastructure_configuration_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineStepDetailsStepInfrastructureConfigurationDetailsRepresentation},
	}

	pipelineStepDetailsRepresentationContainer = map[string]interface{}{
		"step_name":                  acctest.Representation{RepType: acctest.Required, Create: `stepNameContainer`},
		"step_type":                  acctest.Representation{RepType: acctest.Required, Create: `CONTAINER`},
		"depends_on":                 acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `descriptionContainer`},
		"step_configuration_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineStepDetailsStepConfigurationDetailsRepresentation},
		"step_infrastructure_configuration_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineStepDetailsStepInfrastructureConfigurationDetailsRepresentation},
		"step_container_configuration_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatasciencePipelineStepDetailsStepContainerConfigurationDetailsRepresentation},
	}

	pipelineConfigurationDetailsRepresentation = map[string]interface{}{
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"command_line_arguments":     acctest.Representation{RepType: acctest.Optional, Create: `commandLineArguments`},
		"environment_variables":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"environmentVariables": "environmentVariables"}},
		"maximum_runtime_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DatasciencePipelineInfrastructureConfigurationDetailsRepresentation = map[string]interface{}{
		"block_storage_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `50`, Update: `60`},
		"shape_name":                acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		//"shape_config_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineInfrastructureConfigurationDetailsShapeConfigDetailsRepresentation},
		"subnet_id": acctest.Representation{RepType: acctest.Optional, Create: ``},
	}

	pipelineInfrastructureConfigurationDetailsRepresentation = map[string]interface{}{
		"block_storage_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"shape_name":                acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Optional, Create: ``},
	}

	pipelineLogConfigurationDetailsRepresentation = map[string]interface{}{
		"enable_auto_log_creation": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"enable_logging":           acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"log_group_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_logging_log_group.terraform_test_custom_log_group.id}`},
		"log_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_logging_log.terraform_test_custom_log.id}`},
	}
	pipelineStepDetailsStepConfigurationDetailsRepresentation = map[string]interface{}{
		"command_line_arguments":     acctest.Representation{RepType: acctest.Optional, Create: `commandLineArguments`},
		"environment_variables":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"environmentVariables": "environmentVariables"}},
		"maximum_runtime_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DatasciencePipelineStepDetailsStepContainerConfigurationDetailsRepresentation = map[string]interface{}{
		"container_type":     acctest.Representation{RepType: acctest.Required, Create: `OCIR_CONTAINER`},
		"image":              acctest.Representation{RepType: acctest.Required, Create: `iad.ocir.io/ociodscdev/nested-rp-public-python-sdk-1:1.0`},
		"cmd":                acctest.Representation{RepType: acctest.Optional, Create: []string{`hello_world.py`}},
		"entrypoint":         acctest.Representation{RepType: acctest.Optional, Create: []string{`python3`}},
		"image_digest":       acctest.Representation{RepType: acctest.Optional, Create: ``},
		"image_signature_id": acctest.Representation{RepType: acctest.Optional, Create: ``},
	}
	DatasciencePipelineStepDetailsStepInfrastructureConfigurationDetailsRepresentation = map[string]interface{}{
		"block_storage_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `50`},
		//Applicable when step_type=CUSTOM_SCRIPT Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
		//"shape_config_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: pipelineStepDetailsStepInfrastructureConfigurationDetailsShapeConfigDetailsRepresentation},
		"shape_name": acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"subnet_id":  acctest.Representation{RepType: acctest.Optional, Create: ``},
	}

	// shape_configuration_details supported and tested but currently not supported by ml_jobs
	// due to a flex_shape issue - https://jira.oci.oraclecorp.com/browse/ODSC-35288
	pipelineStepDetailsStepInfrastructureConfigurationDetailsRepresentation = map[string]interface{}{
		"block_storage_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"shape_name":                acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Optional, Create: ``},
	}
	pipelineInfrastructureConfigurationDetailsShapeConfigDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1.0`},
	}
	pipelineStepDetailsStepInfrastructureConfigurationDetailsShapeConfigDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1.0`},
	}

	pipelineLogGroupRepresentation = acctest.RepresentationCopyWithNewProperties(LoggingLogGroupRepresentation, map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `terraform_test_custom_log_group`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	})

	pipelineLogRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `terraform_test_custom_log`},
		"log_group_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.terraform_test_custom_log_group.id}`},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"log_type":           acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`},
		"retention_duration": acctest.Representation{RepType: acctest.Optional, Create: `30`},
	}

	PipelineResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "terraform_test_custom_log_group", acctest.Required, acctest.Create, pipelineLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "terraform_test_custom_log", acctest.Required, acctest.Create, pipelineLogRepresentation)
)

// issue-routing-tag: datascience/default
func TestDatasciencePipelineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatasciencePipelineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_pipeline.test_pipeline"
	datasourceName := "data.oci_datascience_pipelines.test_pipelines"
	singularDatasourceName := "data.oci_datascience_pipeline.test_pipeline"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PipelineResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Optional, acctest.Create, pipelineRepresentation), "datascience", "pipeline", t)

	acctest.ResourceTest(t, testAccCheckDatasciencePipelineDestroy, []resource.TestStep{
		// step 0 - verify Create
		{
			Config: config + compartmentIdVariableStr + PipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Required, acctest.Create, pipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "step_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_name", "stepName"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_type", "CUSTOM_SCRIPT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// step 1 - delete before next Create
		{
			Config: config + compartmentIdVariableStr + PipelineResourceDependencies,
		},
		// step 2 - verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Optional, acctest.Create, pipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.command_line_arguments", "commandLineArguments"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "infrastructure_configuration_details.0.shape_name"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_details.0.enable_auto_log_creation", "false"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_details.0.enable_logging", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_details.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_details.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "step_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.depends_on.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.is_artifact_uploaded", "true"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.command_line_arguments", "commandLineArguments"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "step_details.0.step_infrastructure_configuration_details.0.shape_name"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_name", "stepName"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_type", "CUSTOM_SCRIPT"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) { // failing figure out why
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						fmt.Printf("CompartmentId:%v, ResourceName is %s", &compartmentId, resourceName)
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// step 3 - verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(pipelineRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.command_line_arguments", "commandLineArguments"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "infrastructure_configuration_details.0.shape_name"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_details.0.enable_auto_log_creation", "false"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_details.0.enable_logging", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_details.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_details.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "step_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.depends_on.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.is_artifact_uploaded", "true"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.command_line_arguments", "commandLineArguments"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.0.shape_config_details.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "step_details.0.step_infrastructure_configuration_details.0.shape_name"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_name", "stepName"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_type", "CUSTOM_SCRIPT"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// step 4 - verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + PipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Optional, acctest.Update, pipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.command_line_arguments", "commandLineArguments"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttrSet(resourceName, "infrastructure_configuration_details.0.shape_name"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_details.0.enable_auto_log_creation", "false"),
				resource.TestCheckResourceAttr(resourceName, "log_configuration_details.0.enable_logging", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_details.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "log_configuration_details.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "step_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.depends_on.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.is_artifact_uploaded", "true"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.command_line_arguments", "commandLineArguments"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.0.shape_config_details.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_name", "stepName"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_type", "CUSTOM_SCRIPT"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// step 5 - verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_pipelines", "test_pipelines", acctest.Optional, acctest.Update, pipelineDataSourceRepresentation) +
				compartmentIdVariableStr + PipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Optional, acctest.Update, pipelineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "created_by"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "pipelines.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pipelines.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "pipelines.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "pipelines.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "pipelines.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipelines.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipelines.0.project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipelines.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipelines.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipelines.0.time_updated"),
			),
		},
		// step 6 - verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Required, acctest.Create, pipelineSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PipelineResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pipeline_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.0.command_line_arguments", "commandLineArguments"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.0.type", "DEFAULT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_configuration_details.0.enable_auto_log_creation", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_configuration_details.0.enable_logging", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.depends_on.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.is_artifact_uploaded", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.step_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.step_configuration_details.0.command_line_arguments", "commandLineArguments"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.step_configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.step_configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.step_infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.step_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.step_infrastructure_configuration_details.0.shape_config_details.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.step_name", "stepName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "step_details.0.step_type", "CUSTOM_SCRIPT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// step 7 - delete before next Create
		{
			Config: config + compartmentIdVariableStr + PipelineResourceDependencies,
		},
		// step 8 - verify Create with Container Step type
		{
			Config: config + compartmentIdVariableStr + PipelineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_pipeline", "test_pipeline", acctest.Optional, acctest.Create, pipelineRepresentationContainer),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.command_line_arguments", "commandLineArguments"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.type", "DEFAULT"),
				resource.TestCheckResourceAttr(resourceName, "step_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.depends_on.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.description", "descriptionContainer"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.command_line_arguments", "commandLineArguments"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.environment_variables.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_configuration_details.0.maximum_runtime_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_container_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_container_configuration_details.0.cmd.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_container_configuration_details.0.container_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_container_configuration_details.0.entrypoint.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_container_configuration_details.0.image", "iad.ocir.io/ociodscdev/nested-rp-public-python-sdk-1:1.0"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_container_configuration_details.0.image_digest", ""),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_container_configuration_details.0.image_signature_id", ""),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.0.shape_config_details.#", "0"),
				// For flex shape only
				//resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.0.shape_config_details.0.memory_in_gbs", "1.0"),
				//resource.TestCheckResourceAttr(resourceName, "step_details.0.step_infrastructure_configuration_details.0.shape_config_details.0.ocpus", "1.0"),
				resource.TestCheckResourceAttrSet(resourceName, "step_details.0.step_infrastructure_configuration_details.0.shape_name"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_name", "stepNameContainer"),
				resource.TestCheckResourceAttr(resourceName, "step_details.0.step_type", "CONTAINER"),

				func(s *terraform.State) (err error) { // failing figure out why
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					fmt.Printf("ResourceName is %s", resourceName)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						fmt.Printf("CompartmentId:%v, ResourceName is %s", &compartmentId, resourceName)
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify resource import step
		{
			Config:            config + PipelineRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"step_artifact",
				"delete_related_pipeline_runs",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatasciencePipelineDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_pipeline" {
			noResourceFound = false
			request := oci_datascience.GetPipelineRequest{}

			tmp := rs.Primary.ID
			request.PipelineId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetPipeline(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.PipelineLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatasciencePipeline") {
		resource.AddTestSweepers("DatasciencePipeline", &resource.Sweeper{
			Name:         "DatasciencePipeline",
			Dependencies: acctest.DependencyGraph["pipeline"],
			F:            sweepDatasciencePipelineResource,
		})
	}
}

func sweepDatasciencePipelineResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	pipelineIds, err := getPipelineIds(compartment)
	if err != nil {
		return err
	}
	for _, pipelineId := range pipelineIds {
		if ok := acctest.SweeperDefaultResourceId[pipelineId]; !ok {
			deletePipelineRequest := oci_datascience.DeletePipelineRequest{}

			deletePipelineRequest.PipelineId = &pipelineId

			deletePipelineRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeletePipeline(context.Background(), deletePipelineRequest)
			if error != nil {
				fmt.Printf("Error deleting Pipeline %s %s, It is possible that the resource is already deleted. Please verify manually \n", pipelineId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &pipelineId, pipelineSweepWaitCondition, time.Duration(3*time.Minute),
				pipelineSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getPipelineIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PipelineId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listPipelinesRequest := oci_datascience.ListPipelinesRequest{}
	listPipelinesRequest.CompartmentId = &compartmentId
	listPipelinesRequest.LifecycleState = oci_datascience.ListPipelinesLifecycleStateActive
	listPipelinesResponse, err := dataScienceClient.ListPipelines(context.Background(), listPipelinesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Pipeline list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, pipeline := range listPipelinesResponse.Items {
		id := *pipeline.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PipelineId", id)
	}
	return resourceIds, nil
}

func pipelineSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if pipelineResponse, ok := response.Response.(oci_datascience.GetPipelineResponse); ok {
		return pipelineResponse.LifecycleState != oci_datascience.PipelineLifecycleStateDeleted
	}
	return false
}

func pipelineSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetPipeline(context.Background(), oci_datascience.GetPipelineRequest{
		PipelineId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
