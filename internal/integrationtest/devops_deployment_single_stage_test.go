// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DevopsSingleStageDeploymentRequiredOnlyResource = DevopsSingleStageDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Required, acctest.Create, devopsSingleStageDeploymentRepresentation)

	DevopsSingleStageDeploymentResourceConfig = DevopsSingleStageDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Optional, acctest.Update, devopsSingleStageDeploymentRepresentation)

	devopsSingleStageDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deployment.test_deployment.id}`},
	}

	deployOkeSingleStageRepresentation = map[string]interface{}{
		"deploy_pipeline_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"deploy_stage_predecessor_collection":     acctest.RepresentationGroup{RepType: acctest.Required, Group: deployStageDeployStagePredecessorCollectionRepresentation},
		"deploy_stage_type":                       acctest.Representation{RepType: acctest.Required, Create: `OKE_DEPLOYMENT`},
		"defined_tags":                            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
		"oke_cluster_deploy_environment_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_environment.test_deploy_kubernetes_environment.id}`},
		"kubernetes_manifest_deploy_artifact_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_deploy_artifact.test_deploy_inline_artifact.id}`}},
		"namespace":                               acctest.Representation{RepType: acctest.Optional, Create: `helloworld-demo`},
		"rollback_policy":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: deployStageRollbackPolicyRepresentation},
	}

	deployOkeEnvironmentRepresentation = map[string]interface{}{
		"deploy_environment_type": acctest.Representation{RepType: acctest.Required, Create: `OKE_CLUSTER`},
		"project_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"cluster_id":              acctest.Representation{RepType: acctest.Required, Create: `ocid1.cluster.oc1.iad.aaaaaaaalsoirfmjo7kiyneawyxoucafh2qzn2cik45bx7p3fc5f4wtseuca`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	deployGenericArtifactSingleStageRepresentation = map[string]interface{}{
		"argument_substitution_mode": acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `SUBSTITUTE_PLACEHOLDERS`},
		"deploy_artifact_source":     acctest.RepresentationGroup{RepType: acctest.Required, Group: deployGenericArtifactDeployArtifactSingleStageSourceRepresentation},
		"deploy_artifact_type":       acctest.Representation{RepType: acctest.Required, Create: `KUBERNETES_MANIFEST`},
		"project_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	deployGenericArtifactDeployArtifactSingleStageSourceRepresentation = map[string]interface{}{
		"deploy_artifact_source_type": acctest.Representation{RepType: acctest.Required, Create: `GENERIC_ARTIFACT`},
		"repository_id":               acctest.Representation{RepType: acctest.Required, Create: repository_id},
		"deploy_artifact_path":        acctest.Representation{RepType: acctest.Required, Create: artifact_path},
		"deploy_artifact_version":     acctest.Representation{RepType: acctest.Required, Create: version},
	}

	devopsSingleStageDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"deploy_pipeline_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deployment.test_deployment.id}`},
		"project_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `Accepted`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: devopsSingleStageDeploymentDataSourceFilterRepresentation}}
	devopsSingleStageDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_deployment.test_deployment.id}`}},
	}

	devopsSingleStageDeploymentRepresentation = map[string]interface{}{
		"deploy_pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"deployment_type":    acctest.Representation{RepType: acctest.Required, Create: `SINGLE_STAGE_DEPLOYMENT`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"deploy_stage_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreExecutionProgressDifferencesRepresentation},
	}

	deployLogRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"log_group_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_devops_log_group.id}`},
		"log_type":           acctest.Representation{RepType: acctest.Required, Create: `SERVICE`},
		"configuration":      acctest.RepresentationGroup{RepType: acctest.Required, Group: devopLogConfigurationRepresentation},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"retention_duration": acctest.Representation{RepType: acctest.Optional, Create: `30`},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	ignoreExecutionProgressDifferencesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`deployment_execution_progress`, `defined_tags`}},
	}

	DevopsSingleStageDeploymentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_inline_artifact", acctest.Required, acctest.Create, deployGenericArtifactSingleStageRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_kubernetes_environment", acctest.Required, acctest.Create, deployOkeEnvironmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, deployPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployOkeSingleStageRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_devops_log_group", acctest.Required, acctest.Create, devopsLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Create, deployLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeploymentResource_singleStageDeployment(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeploymentResource_singleStageDeployment")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deployment.test_deployment"
	datasourceName := "data.oci_devops_deployments.test_deployments"
	singularDatasourceName := "data.oci_devops_deployment.test_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsSingleStageDeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Optional, acctest.Create, devopsSingleStageDeploymentRepresentation), "devops", "deployment", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsSingleStageDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Required, acctest.Create, devopsSingleStageDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "SINGLE_STAGE_DEPLOYMENT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsSingleStageDeploymentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsSingleStageDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Optional, acctest.Create, devopsSingleStageDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "SINGLE_STAGE_DEPLOYMENT"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DevopsSingleStageDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Optional, acctest.Update, devopsSingleStageDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "SINGLE_STAGE_DEPLOYMENT"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deployments", "test_deployments", acctest.Optional, acctest.Update, devopsDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsSingleStageDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Optional, acctest.Update, devopsSingleStageDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),

				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Required, acctest.Create, devopsSingleStageDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsSingleStageDeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_type", "SINGLE_STAGE_DEPLOYMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DevopsSingleStageDeploymentResourceConfig,
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
