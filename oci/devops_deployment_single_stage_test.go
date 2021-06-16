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
	DevopsSingleStageDeploymentRequiredOnlyResource = DevopsSingleStageDeploymentResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Required, Create, devopsSingleStageDeploymentRepresentation)

	DevopsSingleStageDeploymentResourceConfig = DevopsSingleStageDeploymentResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Update, devopsSingleStageDeploymentRepresentation)

	devopsSingleStageDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": Representation{repType: Required, create: `${oci_devops_deployment.test_deployment.id}`},
	}

	deployOkeSingleStageRepresentation = map[string]interface{}{
		"deploy_pipeline_id":                      Representation{repType: Required, create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"deploy_stage_predecessor_collection":     RepresentationGroup{Required, deployStageDeployStagePredecessorCollectionRepresentation},
		"deploy_stage_type":                       Representation{repType: Required, create: `OKE_DEPLOYMENT`},
		"defined_tags":                            Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                             Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":                            Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":                           Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"oke_cluster_deploy_environment_id":       Representation{repType: Required, create: `${oci_devops_deploy_environment.test_deploy_kubernetes_environment.id}`},
		"kubernetes_manifest_deploy_artifact_ids": Representation{repType: Required, create: []string{`${oci_devops_deploy_artifact.test_deploy_inline_artifact.id}`}},
		"namespace":                               Representation{repType: Optional, create: `helloworld-demo`},
		"rollback_policy":                         RepresentationGroup{Optional, deployStageRollbackPolicyRepresentation},
	}

	deployOkeEnvironmentRepresentation = map[string]interface{}{
		"deploy_environment_type": Representation{repType: Required, create: `OKE_CLUSTER`},
		"project_id":              Representation{repType: Required, create: `${oci_devops_project.test_project.id}`},
		"cluster_id":              Representation{repType: Required, create: `ocid1.cluster.oc1.phx.aaaaaaaaqu6xnflkdfghjvcp3dw7qwliqygtfmdw3yvbapudjcwkwfecjxxq`}, // TODO: Need to create via terraform
		"defined_tags":            Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":            Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":           Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}

	deployGenericArtifactSingleStageRepresentation = map[string]interface{}{
		"argument_substitution_mode": Representation{repType: Required, create: `NONE`, update: `SUBSTITUTE_PLACEHOLDERS`},
		"deploy_artifact_source":     RepresentationGroup{Required, deployGenericArtifactDeployArtifactSingleStageSourceRepresentation},
		"deploy_artifact_type":       Representation{repType: Required, create: `KUBERNETES_MANIFEST`},
		"project_id":                 Representation{repType: Required, create: `${oci_devops_project.test_project.id}`},
		"defined_tags":               Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":               Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":              Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}
	repository_id_single_stage                                         = "ocid1.artifactrepository.oc1.phx.0.amaaaaaansx72maaj7g2xjtiuffscp7jouvkpttxnbjxuxr6q45mt7b5lqfq" // TODO: Need to create via terraform
	deployGenericArtifactDeployArtifactSingleStageSourceRepresentation = map[string]interface{}{
		"deploy_artifact_source_type": Representation{repType: Required, create: `GENERIC_ARTIFACT`},
		"repository_id":               Representation{repType: Required, create: repository_id},
		"deploy_artifact_path":        Representation{repType: Required, create: artifact_path},
		"deploy_artifact_version":     Representation{repType: Required, create: version},
	}

	devopsSingleStageDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Optional, create: `${var.compartment_id}`},
		"deploy_pipeline_id": Representation{repType: Optional, create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"display_name":       Representation{repType: Optional, create: `displayName`},
		"id":                 Representation{repType: Optional, create: `${oci_devops_deployment.test_deployment.id}`},
		"project_id":         Representation{repType: Optional, create: `${oci_devops_project.test_project.id}`},
		"state":              Representation{repType: Optional, create: `Active`},
		"filter":             RepresentationGroup{Required, devopsSingleStageDeploymentDataSourceFilterRepresentation}}
	devopsSingleStageDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_devops_deployment.test_deployment.id}`}},
	}

	devopsSingleStageDeploymentRepresentation = map[string]interface{}{
		"deploy_pipeline_id": Representation{repType: Required, create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"deployment_type":    Representation{repType: Required, create: `SINGLE_STAGE_DEPLOYMENT`},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":       Representation{repType: Optional, create: `displayName`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"deploy_stage_id":    Representation{repType: Required, create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
		"lifecycle":          RepresentationGroup{Required, ignoreDefinedTagsDifferencesRepresentation},
	}

	deployLogRepresentation = map[string]interface{}{
		"display_name":       Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"log_group_id":       Representation{repType: Required, create: `${oci_logging_log_group.test_log_group.id}`, update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_type":           Representation{repType: Required, create: `SERVICE`},
		"configuration":      RepresentationGroup{Required, devopLogConfigurationRepresentation},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         Representation{repType: Optional, create: `true`},
		"retention_duration": Representation{repType: Optional, create: `30`},
		"lifecycle":          RepresentationGroup{Required, ignoreDefinedTagsDifferencesRepresentation},
	}

	ignoreDefinedTagsDifferencesRepresentation = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}

	DevopsSingleStageDeploymentResourceDependencies = generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_inline_artifact", Required, Create, deployGenericArtifactSingleStageRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_kubernetes_environment", Required, Create, deployOkeEnvironmentRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", Required, Create, deployPipelineRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployOkeSingleStageRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_logging_log", "test_log", Optional, Create, deployLogRepresentation) +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

func TestDevopsDeploymentResource_singleStageDeployment(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeploymentResource_singleStageDeployment")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deployment.test_deployment"
	datasourceName := "data.oci_devops_deployments.test_deployments"
	singularDatasourceName := "data.oci_devops_deployment.test_deployment"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DevopsSingleStageDeploymentResourceDependencies+
		generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Create, devopsSingleStageDeploymentRepresentation), "devops", "deployment", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DevopsSingleStageDeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Required, Create, devopsSingleStageDeploymentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "SINGLE_STAGE_DEPLOYMENT"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DevopsSingleStageDeploymentResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DevopsSingleStageDeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Create, devopsSingleStageDeploymentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_id"),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "SINGLE_STAGE_DEPLOYMENT"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DevopsSingleStageDeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Update, devopsSingleStageDeploymentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_id"),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "SINGLE_STAGE_DEPLOYMENT"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
					generateDataSourceFromRepresentationMap("oci_devops_deployments", "test_deployments", Optional, Update, devopsDeploymentDataSourceRepresentation) +
					compartmentIdVariableStr + DevopsSingleStageDeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Update, devopsSingleStageDeploymentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "Accepted"),

					resource.TestCheckResourceAttr(datasourceName, "deployment_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "deployment_collection.0.items.#", "0"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Required, Create, devopsSingleStageDeploymentSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DevopsSingleStageDeploymentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "deployment_type", "SINGLE_STAGE_DEPLOYMENT"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
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
		},
	})
}
