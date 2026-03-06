// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	modelForPublicEgressModelDeploymentRepresentation = map[string]interface{}{
		"artifact_content_length":      acctest.Representation{RepType: acctest.Required, Create: `7532`},
		"model_artifact":               acctest.Representation{RepType: acctest.Required, Create: `../../examples/datascience/model_deployment/artifact.zip`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"artifact_content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `attachment; filename=tfTestArtifact.zip`},
	}

	DatascienceModelDeploymentPublicEgressResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create, modelForPublicEgressModelDeploymentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, logGroupMDRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_access_log", acctest.Required, acctest.Create, customAccessLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_predict_log", acctest.Required, acctest.Create, customPredictLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", acctest.Required, acctest.Create, logGroupUpdateMDRepresentation)

	DatascienceModelDeploymentPublicEgressInstanceConfigurationRepresentation = map[string]interface{}{
		"instance_shape_name":                            acctest.Representation{RepType: acctest.Required, Create: `${var.instance_shape_name}`},
		"model_deployment_instance_shape_config_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfigurationModelDeploymentInstanceShapeConfigDetailsRepresentation},
		"network_access_type":                            acctest.Representation{RepType: acctest.Required, Create: `MANAGED_NETWORKING_NO_INTERNET_ACCESS`, Update: `MANAGED_NETWORKING_INTERNET_ACCESS`},
	}

	DatascienceModelDeploymentPublicEgressModelConfigurationDetailsRepresentation = map[string]interface{}{
		"instance_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentPublicEgressInstanceConfigurationRepresentation},
		"model_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`, Update: `${oci_datascience_model.test_model.id}`},
		"bandwidth_mbps":         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"maximum_bandwidth_mbps": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"scaling_policy":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicyRepresentation},
	}

	DatascienceModelDeploymentPublicEgressModelDeploymentConfigurationDetailsRepresentation = map[string]interface{}{
		"deployment_type":                   acctest.Representation{RepType: acctest.Required, Create: `SINGLE_MODEL`},
		"model_configuration_details":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentPublicEgressModelConfigurationDetailsRepresentation},
		"environment_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentPublicEgressEnvironmentConfigurationDetailsRepresentation},
	}

	DatascienceModelDeploymentPublicEgressEnvironmentConfigurationDetailsRepresentation = map[string]interface{}{
		"environment_configuration_type": acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`, Update: `DEFAULT`},
		"cmd":                            acctest.Representation{RepType: acctest.Optional, Create: []string{`cmd`}, Update: []string{`cmd2`}},
		"default_environment_variables":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"defaultEnvironmentVariables": "defaultEnvironmentVariables"}, Update: map[string]string{"defaultEnvironmentVariables2": "defaultEnvironmentVariables2"}},
		"entrypoint":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`entrypoint`}, Update: []string{`entrypoint2`}},
		"environment_variables":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"environmentVariables": "environmentVariables"}, Update: map[string]string{"environmentVariables2": "environmentVariables2"}},
		"health_check_port":              acctest.Representation{RepType: acctest.Optional, Create: `1024`, Update: `1025`},
		"server_port":                    acctest.Representation{RepType: acctest.Optional, Create: `1024`, Update: `1025`},
		"image":                          acctest.Representation{RepType: acctest.Optional, Create: `${var.image}`, Update: `${var.image}`},
		"image_digest":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.image_digest}`, Update: `${var.image_digest}`},
		"image_signature_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.image_signature_id}`},
	}

	DatascienceModelDeploymentPublicEgressRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_deployment_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentPublicEgressModelDeploymentConfigurationDetailsRepresentation},
		"project_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"category_log_details":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentCategoryLogDetailsRepresentation},
		"defined_tags":                           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                              acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreMDRepresentation},
	}

	DatascienceModelDeploymentPublicEgressSingularDataSourceRepresentation = map[string]interface{}{
		"model_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model_deployment.test_model_deployment_public_egress.id}`},
	}

	DatascienceModelDeploymentPublicEgressDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_deployment.test_model_deployment_public_egress.created_by}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_deployment.test_model_deployment_public_egress.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentPublicEgressDataSourceFilterRepresentation},
	}
	DatascienceModelDeploymentPublicEgressDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_model_deployment.test_model_deployment_public_egress.id}`}},
	}
)

// issue-routing-tag: datascience/default
func TestDatascienceModelDeploymentResource_public_egress_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDeploymentResource_public_egress_basic")
	defer httpreplay.SaveScenario()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	t.Logf("TestDatascienceModelDeploymentResource_public_egress_basic compartment_id=%s", compartmentId)

	instanceShapeName := utils.GetEnvSettingWithBlankDefault("instance_shape_name")
	instanceShapeNameVariableStr := fmt.Sprintf("variable \"instance_shape_name\" { default = \"%s\" }\n", instanceShapeName)

	image := utils.GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	imageDigest := utils.GetEnvSettingWithBlankDefault("image_digest")
	imageDigestVariableStr := fmt.Sprintf("variable \"image_digest\" { default = \"%s\" }\n", imageDigest)

	imageSignatureID := utils.GetEnvSettingWithBlankDefault("image_signature_id")
	imageSignatureIDVariableStr := fmt.Sprintf("variable \"image_signature_id\" { default = \"%s\" }\n", imageSignatureID)

	config := acctest.ProviderTestConfig() + instanceShapeNameVariableStr + imageVariableStr + imageDigestVariableStr + imageSignatureIDVariableStr

	resourceName := "oci_datascience_model_deployment.test_model_deployment_public_egress"
	singularDatasourceName := "data.oci_datascience_model_deployment.test_model_deployment_public_egress"

	acctest.ResourceTest(t, testAccCheckDatascienceModelDeploymentDestroy, []resource.TestStep{
		// create with MANAGED_NETWORKING_INTERNET_ACCESS
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDeploymentPublicEgressResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment_public_egress", acctest.Required, acctest.Create, DatascienceModelDeploymentPublicEgressRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.network_access_type", "MANAGED_NETWORKING_NO_INTERNET_ACCESS"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
			),
		},

		// update to MANAGED_NETWORKING_NO_INTERNET_ACCESS
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDeploymentPublicEgressResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment_public_egress", acctest.Optional, acctest.Update, DatascienceModelDeploymentPublicEgressRepresentation),
			ExpectNonEmptyPlan: true,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.network_access_type", "MANAGED_NETWORKING_INTERNET_ACCESS"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
			),
		},

		// verify singular datasource - get model deployment
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDeploymentPublicEgressResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment_public_egress", acctest.Required, acctest.Create, DatascienceModelDeploymentPublicEgressSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment_public_egress", acctest.Optional, acctest.Update, DatascienceModelDeploymentPublicEgressRepresentation),
			ExpectNonEmptyPlan: true,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
	})
}
