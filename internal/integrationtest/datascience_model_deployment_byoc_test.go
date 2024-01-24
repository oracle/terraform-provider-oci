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
	//"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	//"github.com/oracle/oci-go-sdk/v65/common"
	//oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatascienceModelDeploymentBYOCRequiredOnlyResource = DatascienceModelDeploymentBYOCResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatascienceModelDeploymentBYOCRepresentation)

	DatascienceModelDeploymentBYOCResourceConfig = DatascienceModelDeploymentBYOCResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatascienceModelDeploymentBYOCRepresentation)

	modelForModelDeploymentBYOCRepresentation = map[string]interface{}{
		"artifact_content_length":      acctest.Representation{RepType: acctest.Required, Create: `1772`},
		"model_artifact":               acctest.Representation{RepType: acctest.Required, Create: `../../examples/datascience/model_deployment/artifact_byoc.zip`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"artifact_content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `attachment; filename=tfTestArtifact.zip`},
	}

	modelForUpdateModelDeploymentBYOCRepresentation = map[string]interface{}{
		"artifact_content_length":      acctest.Representation{RepType: acctest.Required, Create: `1772`},
		"model_artifact":               acctest.Representation{RepType: acctest.Required, Create: `../../examples/datascience/model_deployment/artifact_byoc.zip`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"artifact_content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `attachment; filename=tfTestArtifact.zip`},
	}

	DatascienceModelDeploymentBYOCRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_deployment_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentBYOCModelDeploymentConfigurationDetailsRepresentation},
		"project_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"category_log_details":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentBYOCCategoryLogDetailsRepresentation},
		"defined_tags":                           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatascienceModelDeploymentBYOCModelDeploymentConfigurationDetailsRepresentation = map[string]interface{}{
		"deployment_type":                   acctest.Representation{RepType: acctest.Required, Create: `SINGLE_MODEL`},
		"model_configuration_details":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsRepresentation},
		"environment_configuration_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentBYOCModelDeploymentConfigurationDetailsEnvironmentConfigurationDetailsRepresentation},
	}

	DatascienceModelDeploymentBYOCModelDeploymentConfigurationDetailsEnvironmentConfigurationDetailsRepresentation = map[string]interface{}{
		"environment_configuration_type": acctest.Representation{RepType: acctest.Required, Create: `OCIR_CONTAINER`, Update: `OCIR_CONTAINER`},
		"cmd":                            acctest.Representation{RepType: acctest.Optional, Create: []string{`python`, `-m`, `uvicorn`, `local_server_main:app`, `--port`, `5000`, `--host`, `0.0.0.0`}, Update: []string{`python`, `-m`, `uvicorn`, `local_server_main:app`, `--port`, `5000`, `--host`, `0.0.0.0`}},
		"entrypoint":                     acctest.Representation{RepType: acctest.Optional, Create: []string{``}, Update: []string{``}},
		"environment_variables":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"environmentVariables": ""}, Update: map[string]string{"environmentVariables": "1"}},
		"health_check_port":              acctest.Representation{RepType: acctest.Optional, Create: `5000`, Update: `5000`},
		"image":                          acctest.Representation{RepType: acctest.Optional, Create: `iad.ocir.io/ociodscdev/onnx_demo:1.0.3`, Update: `iad.ocir.io/ociodscdev/onnx_demo:1.0.3`},
		"image_digest":                   acctest.Representation{RepType: acctest.Optional, Create: ``, Update: ``},
		"server_port":                    acctest.Representation{RepType: acctest.Optional, Create: `5000`, Update: `5000`},
	}

	DatascienceModelDeploymentBYOCCategoryLogDetailsRepresentation = map[string]interface{}{
		"access":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentBYOCCategoryLogDetailsAccessRepresentation},
		"predict": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentBYOCCategoryLogDetailsPredictRepresentation},
	}

	DatascienceModelDeploymentBYOCCategoryLogDetailsAccessRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
		"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_access_log.id}`},
	}

	DatascienceModelDeploymentBYOCCategoryLogDetailsPredictRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
		"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_predict_log.id}`},
	}

	logGroupMDBYOCRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tf_testing_byoc_log_group`, Update: `tf_testing_byoc_log_group_update`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	logGroupUpdateMDBYOCRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tf_update_testing_byoc_log_group`, Update: `tf_update_testing_byoc_log_group_update`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	customAccessLogBYOCRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `tf-testing-access-log`, Update: `tf-testing-Update-access-log`},
		"log_group_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_type":           acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"retention_duration": acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `60`},
	}

	customPredictLogBYOCRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `tf-testing-predict-log`, Update: `tf-testing-Update-predict-log`},
		"log_group_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_type":           acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"retention_duration": acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `60`},
	}

	DatascienceModelDeploymentBYOCResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create, modelForModelDeploymentBYOCRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model_update", acctest.Optional, acctest.Create, modelForUpdateModelDeploymentBYOCRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, logGroupMDBYOCRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_access_log", acctest.Required, acctest.Create, customAccessLogBYOCRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_predict_log", acctest.Required, acctest.Create, customPredictLogBYOCRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", acctest.Required, acctest.Create, logGroupUpdateMDBYOCRepresentation)
)

// issue-routing-tag: datascience/default
func TestDatascienceModelDeploymentBYOCResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDeploymentBYOCResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_model_deployment.test_model_deployment"
	singularDatasourceName := "data.oci_datascience_model_deployment.test_model_deployment"
	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatascienceModelDeploymentDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDeploymentBYOCResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create, DatascienceModelDeploymentBYOCRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.cmd.0", "python"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.cmd.6", "--host"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.entrypoint.0", ""),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_variables.environmentVariables", ""),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "5000"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/onnx_demo:1.0.3"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:243590ea099af4019b6afc104b8a70b9552f0b001b37d0442f8b5a399244681c"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "5000"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify updates to updatable parameters including environmental config
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDeploymentBYOCResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatascienceModelDeploymentBYOCRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.cmd.0", "python"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.cmd.6", "--host"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.entrypoint.0", ""),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_variables.environmentVariables", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "5000"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/onnx_demo:1.0.3"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:243590ea099af4019b6afc104b8a70b9552f0b001b37d0442f8b5a399244681c"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "5000"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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

		// verify singular datasource - get model deployment
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatascienceDatascienceModelDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelDeploymentBYOCResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.cmd.0", "python"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.cmd.6", "--host"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.entrypoint.0", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_variables.environmentVariables", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "5000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/onnx_demo:1.0.3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:243590ea099af4019b6afc104b8a70b9552f0b001b37d0442f8b5a399244681c"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "5000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},

		// verify resource import
		{
			Config:            config + DatascienceModelDeploymentBYOCRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"description",
			},
			ResourceName: resourceName,
		},
	})
}
