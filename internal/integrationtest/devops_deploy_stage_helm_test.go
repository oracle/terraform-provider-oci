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
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DeployHelmStageRequiredOnlyResource = DeployHelmStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployHelmStageRepresentation)

	DeployHelmStageResourceConfig = DeployHelmStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployHelmStageRepresentation)

	deployHelmStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	deployHelmStageRollbackPolicyRepresentation = map[string]interface{}{
		"policy_type": acctest.Representation{RepType: acctest.Optional, Create: `AUTOMATED_STAGE_ROLLBACK_POLICY`},
	}

	DevopsDeployStageSetStringItemsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	DevopsDeployStageSetValuesItemsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	DevopsDeployStageSetStringRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsDeployStageSetStringItemsRepresentation},
	}
	DevopsDeployStageSetValuesRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsDeployStageSetValuesItemsRepresentation},
	}

	deployHelmStageRepresentation = acctest.GetUpdatedRepresentationCopy("deploy_stage_type", acctest.Representation{RepType: acctest.Required, Create: `OKE_HELM_CHART_DEPLOYMENT`},
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DevopsDeployStageRepresentation, []string{"wait_criteria"}), map[string]interface{}{
			"oke_cluster_deploy_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_environment.test_deploy_kubernetes_environment.id}`},
			"helm_chart_deploy_artifact_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_artifact.test_deploy_helm_artifact.id}`},
			"release_name":                      acctest.Representation{RepType: acctest.Required, Create: `release-name`},
			"purpose":                           acctest.Representation{RepType: acctest.Required, Create: `EXECUTE_HELM_UPGRADE`, Update: `EXECUTE_HELM_COMMAND`},
			"namespace":                         acctest.Representation{RepType: acctest.Optional, Create: `namespace`},
			"rollback_policy":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: deployHelmStageRollbackPolicyRepresentation},
			"are_hooks_enabled":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"is_debug_enabled":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"is_force_enabled":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"max_history":                       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
			"set_string":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsDeployStageSetStringRepresentation},
			"set_values":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsDeployStageSetValuesRepresentation},
			"should_cleanup_on_fail":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"should_not_wait":                   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"should_reset_values":               acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"should_reuse_values":               acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"should_skip_crds":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"should_skip_render_subchart_notes": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
			"is_uninstall_on_stage_delete":      acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		}))

	DeployHelmStageResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_helm_artifact", acctest.Required, acctest.Create, deployHelmChartArtifactRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_kubernetes_environment", acctest.Required, acctest.Create, DevopsdeployEnvironmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, DevopsDeployPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: devops/default
func TestDevopsDeployStageResource_deployHelm(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployStageResource_deployHelm")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_stage.test_deploy_stage"
	datasourceName := "data.oci_devops_deploy_stages.test_deploy_stages"
	singularDatasourceName := "data.oci_devops_deploy_stage.test_deploy_stage"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeployHelmStageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Create, deployHelmStageRepresentation), "devops", "deployStage", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployStageDestroy, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + DeployHelmStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployHelmStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "OKE_HELM_CHART_DEPLOYMENT"),
				resource.TestCheckResourceAttr(resourceName, "release_name", "release-name"),
				resource.TestCheckResourceAttr(resourceName, "purpose", "EXECUTE_HELM_UPGRADE"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttrSet(resourceName, "helm_chart_deploy_artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "oke_cluster_deploy_environment_id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "default"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "set_string.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "set_values.#", "0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		{
			Config: config + compartmentIdVariableStr + DeployHelmStageResourceDependencies,
		},
		{
			Config: config + compartmentIdVariableStr + DeployHelmStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Create, deployHelmStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "OKE_HELM_CHART_DEPLOYMENT"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "release_name", "release-name"),
				resource.TestCheckResourceAttr(resourceName, "purpose", "EXECUTE_HELM_UPGRADE"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oke_cluster_deploy_environment_id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),

				// helm args validation
				resource.TestCheckResourceAttr(resourceName, "are_hooks_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_debug_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_force_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_uninstall_on_stage_delete", "true"),
				resource.TestCheckResourceAttr(resourceName, "max_history", "10"),
				resource.TestCheckResourceAttr(resourceName, "set_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "set_string.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "set_string.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "set_string.0.items.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "set_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "set_values.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "set_values.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "set_values.0.items.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "should_cleanup_on_fail", "false"),
				resource.TestCheckResourceAttr(resourceName, "should_not_wait", "false"),
				resource.TestCheckResourceAttr(resourceName, "should_reset_values", "false"),
				resource.TestCheckResourceAttr(resourceName, "should_reuse_values", "false"),
				resource.TestCheckResourceAttr(resourceName, "should_skip_crds", "false"),
				resource.TestCheckResourceAttr(resourceName, "should_skip_render_subchart_notes", "false"),

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

		{
			Config: config + compartmentIdVariableStr + DeployHelmStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployHelmStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oke_cluster_deploy_environment_id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),

				// // helm args validation
				resource.TestCheckResourceAttr(resourceName, "are_hooks_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_debug_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_force_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "max_history", "11"),
				resource.TestCheckResourceAttr(resourceName, "set_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "set_string.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "set_string.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "set_string.0.items.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "set_values.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "set_values.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "set_values.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "set_values.0.items.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "should_cleanup_on_fail", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_not_wait", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_reset_values", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_reuse_values", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_skip_crds", "true"),
				resource.TestCheckResourceAttr(resourceName, "should_skip_render_subchart_notes", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_uninstall_on_stage_delete", "false"),
				resource.TestCheckResourceAttr(resourceName, "purpose", "EXECUTE_HELM_COMMAND"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stages", "test_deploy_stages", acctest.Optional, acctest.Update, DevopsDevopsDeployStageDataSourceRepresentation) +
				compartmentIdVariableStr + DeployHelmStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployHelmStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployHelmStageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployHelmStageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_type", "OKE_HELM_CHART_DEPLOYMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "release_name", "release-name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oke_cluster_deploy_environment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollback_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),

				// // helm args validation
				resource.TestCheckResourceAttr(singularDatasourceName, "are_hooks_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_debug_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_force_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_history", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "set_string.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "set_string.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "set_string.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "set_string.0.items.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "set_values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "set_values.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "set_values.0.items.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "set_values.0.items.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "should_cleanup_on_fail", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "should_not_wait", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "should_reset_values", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "should_reuse_values", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "should_skip_crds", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "should_skip_render_subchart_notes", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_uninstall_on_stage_delete", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "purpose", "EXECUTE_HELM_COMMAND"),
			),
		},
		{
			Config:                  config + DevopsDeployStageRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
