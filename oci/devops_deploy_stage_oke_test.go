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
	DeployOkeStageRequiredOnlyResource = DeployOkeStageResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployOkeStageRepresentation)

	DeployOkeStageResourceConfig = DeployOkeStageResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployOkeStageRepresentation)

	deployOkeStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": Representation{repType: Required, create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	deployOkeStageRepresentation = getUpdatedRepresentationCopy("deploy_stage_type", Representation{repType: Required, create: `OKE_DEPLOYMENT`},
		representationCopyWithNewProperties(representationCopyWithRemovedProperties(deployStageRepresentation, []string{"wait_criteria"}), map[string]interface{}{
			"oke_cluster_deploy_environment_id":       Representation{repType: Required, create: `${oci_devops_deploy_environment.test_deploy_kubernetes_environment.id}`},
			"kubernetes_manifest_deploy_artifact_ids": Representation{repType: Required, create: []string{`${oci_devops_deploy_artifact.test_deploy_inline_artifact.id}`}},
			"namespace":       Representation{repType: Optional, create: `helloworld-demo`},
			"rollback_policy": RepresentationGroup{Optional, deployStageRollbackPolicyRepresentation},
		}))

	deployStageRollbackPolicyRepresentation = map[string]interface{}{
		"policy_type": Representation{repType: Optional, create: `AUTOMATED_STAGE_ROLLBACK_POLICY`},
	}

	DeployOkeStageResourceDependencies = generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_inline_artifact", Required, Create, deployArtifactRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_kubernetes_environment", Required, Create, deployEnvironmentRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", Required, Create, deployPipelineRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeployStageResource_deployOke(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployStageResource_deployOke")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_stage.test_deploy_stage"
	datasourceName := "data.oci_devops_deploy_stages.test_deploy_stages"
	singularDatasourceName := "data.oci_devops_deploy_stage.test_deploy_stage"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DeployOkeStageResourceDependencies+
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployOkeStageRepresentation), "devops", "deployStage", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDevopsDeployStageDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DeployOkeStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployOkeStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "OKE_DEPLOYMENT"),
					resource.TestCheckResourceAttrSet(resourceName, "oke_cluster_deploy_environment_id"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "default"),
					resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "0"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DeployOkeStageResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DeployOkeStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployOkeStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "OKE_DEPLOYMENT"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "oke_cluster_deploy_environment_id"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "helloworld-demo"),
					resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),

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
				Config: config + compartmentIdVariableStr + DeployOkeStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployOkeStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "oke_cluster_deploy_environment_id"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "helloworld-demo"),
					resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),

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
					generateDataSourceFromRepresentationMap("oci_devops_deploy_stages", "test_deploy_stages", Optional, Update, deployStageDataSourceRepresentation) +
					compartmentIdVariableStr + DeployOkeStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployOkeStageRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployOkeStageSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DeployOkeStageResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_type", "OKE_DEPLOYMENT"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "oke_cluster_deploy_environment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "namespace", "helloworld-demo"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rollback_policy.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + DeployOkeStageResourceConfig,
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
