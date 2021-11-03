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
	DeployInvokeFunctionStageRequiredOnlyResource = DeployInvokeFunctionStageResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployInvokeFunctionStageRepresentation)

	DeployInvokeFunctionStageResourceConfig = DeployInvokeFunctionStageResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployInvokeFunctionStageRepresentation)

	deployInvokeFunctionStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": Representation{RepType: Required, Create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	deployInvokeFunctionStageRepresentation = GetUpdatedRepresentationCopy("deploy_stage_type", Representation{RepType: Required, Create: `INVOKE_FUNCTION`},
		RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(deployStageRepresentation, []string{"wait_criteria"}), map[string]interface{}{
			"function_deploy_environment_id": Representation{RepType: Required, Create: `${oci_devops_deploy_environment.test_deploy_function_environment_1.id}`},
			"is_async":                       Representation{RepType: Required, Create: `true`, Update: `false`},
			"is_validation_enabled":          Representation{RepType: Required, Create: `false`, Update: `true`},
			"deploy_artifact_id":             Representation{RepType: Optional, Create: `${oci_devops_deploy_artifact.test_deploy_ocir_artifact_1.id}`},
		}))

	DeployInvokeFunctionStageResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_ocir_artifact_1", Required, Create, deployOcirArtifactRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_function_environment_1", Required, Create, deployFunctionEnvironmentRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", Required, Create, deployPipelineRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeployStageResource_invokeFunction(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployStageResource_invokeFunction")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_stage.test_deploy_stage"
	datasourceName := "data.oci_devops_deploy_stages.test_deploy_stages"
	singularDatasourceName := "data.oci_devops_deploy_stage.test_deploy_stage"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DeployInvokeFunctionStageResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployInvokeFunctionStageRepresentation), "devops", "deployStage", t)

	ResourceTest(t, testAccCheckDevopsDeployStageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployInvokeFunctionStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployInvokeFunctionStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "INVOKE_FUNCTION"),
				resource.TestCheckResourceAttrSet(resourceName, "function_deploy_environment_id"),
				resource.TestCheckResourceAttr(resourceName, "is_async", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_validation_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeployInvokeFunctionStageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployInvokeFunctionStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployInvokeFunctionStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "INVOKE_FUNCTION"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "function_deploy_environment_id"),
				resource.TestCheckResourceAttr(resourceName, "is_async", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_validation_enabled", "false"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DeployInvokeFunctionStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployInvokeFunctionStageRepresentation),
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
				resource.TestCheckResourceAttrSet(resourceName, "function_deploy_environment_id"),
				resource.TestCheckResourceAttr(resourceName, "is_async", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_validation_enabled", "true"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stages", "test_deploy_stages", Optional, Update, deployStageDataSourceRepresentation) +
				compartmentIdVariableStr + DeployInvokeFunctionStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployInvokeFunctionStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttr(datasourceName, "deploy_stage_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployInvokeFunctionStageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployInvokeFunctionStageResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_type", "INVOKE_FUNCTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "function_deploy_environment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_async", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_validation_enabled", "true"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DeployInvokeFunctionStageResourceConfig,
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
