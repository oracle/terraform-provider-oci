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
	DeployFunctionStageRequiredOnlyResource = DeployFunctionStageResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployFunctionStageRepresentation)

	DeployFunctionStageResourceConfig = DeployFunctionStageResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployFunctionStageRepresentation)

	deployFunctionStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": Representation{repType: Required, create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	deployFunctionStageRepresentation = getUpdatedRepresentationCopy("deploy_stage_type", Representation{repType: Required, create: `DEPLOY_FUNCTION`},
		representationCopyWithNewProperties(representationCopyWithRemovedProperties(deployStageRepresentation, []string{"wait_criteria"}), map[string]interface{}{
			"function_deploy_environment_id":  Representation{repType: Required, create: `${oci_devops_deploy_environment.test_deploy_function_environment.id}`},
			"docker_image_deploy_artifact_id": Representation{repType: Required, create: `${oci_devops_deploy_artifact.test_deploy_ocir_artifact.id}`},
			"function_timeout_in_seconds":     Representation{repType: Required, create: `30`, update: `20`},
			"max_memory_in_mbs":               Representation{repType: Required, create: `128`, update: `256`},
		}))

	DeployFunctionStageResourceDependencies = generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_ocir_artifact", Required, Create, deployOcirArtifactRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_function_environment", Required, Create, deployFunctionEnvironmentRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", Required, Create, deployPipelineRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

func TestDevopsDeployStageResource_deployFunction(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployStageResource_deployFunction")
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
	saveConfigContent(config+compartmentIdVariableStr+DeployFunctionStageResourceDependencies+
		generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployFunctionStageRepresentation), "devops", "deployStage", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDevopsDeployStageDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DeployFunctionStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployFunctionStageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "DEPLOY_FUNCTION"),
					resource.TestCheckResourceAttrSet(resourceName, "function_deploy_environment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "docker_image_deploy_artifact_id"),
					resource.TestCheckResourceAttr(resourceName, "function_timeout_in_seconds", "30"),
					resource.TestCheckResourceAttr(resourceName, "max_memory_in_mbs", "128"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DeployFunctionStageResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DeployFunctionStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployFunctionStageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "DEPLOY_FUNCTION"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "function_deploy_environment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "docker_image_deploy_artifact_id"),
					resource.TestCheckResourceAttr(resourceName, "function_timeout_in_seconds", "30"),
					resource.TestCheckResourceAttr(resourceName, "max_memory_in_mbs", "128"),

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
				Config: config + compartmentIdVariableStr + DeployFunctionStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployFunctionStageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttrSet(resourceName, "function_deploy_environment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "docker_image_deploy_artifact_id"),
					resource.TestCheckResourceAttr(resourceName, "function_timeout_in_seconds", "20"),
					resource.TestCheckResourceAttr(resourceName, "max_memory_in_mbs", "256"),

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
					compartmentIdVariableStr + DeployFunctionStageResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployFunctionStageRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployFunctionStageSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DeployFunctionStageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_type", "DEPLOY_FUNCTION"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "function_deploy_environment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "docker_image_deploy_artifact_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "function_timeout_in_seconds", "20"),
					resource.TestCheckResourceAttr(singularDatasourceName, "max_memory_in_mbs", "256"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + DeployFunctionStageResourceConfig,
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
