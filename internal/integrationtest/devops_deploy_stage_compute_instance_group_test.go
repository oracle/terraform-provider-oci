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
	DeployComputeInstanceGroupStageRequiredOnlyResource = DeployComputeInstanceGroupStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployComputeInstanceGroupStageRepresentation)

	DeployComputeInstanceGroupStageResourceConfig = DeployComputeInstanceGroupStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployComputeInstanceGroupStageRepresentation)

	deployComputeInstanceGroupStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	deployComputeInstanceGroupStageRepresentation = acctest.GetUpdatedRepresentationCopy("deploy_stage_type", acctest.Representation{RepType: acctest.Required, Create: `COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT`},
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DevopsDeployStageRepresentation, []string{"wait_criteria"}), map[string]interface{}{
			"compute_instance_group_deploy_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_environment.test_deploy_instance_group_environment.id}`},
			"deployment_spec_deploy_artifact_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_artifact.test_deploy_generic_artifact.id}`},
			"rollout_policy":                               acctest.RepresentationGroup{RepType: acctest.Required, Group: deployComputeInstanceStageRolloutPolicyRepresentation},
			"rollback_policy":                              acctest.RepresentationGroup{RepType: acctest.Optional, Group: deployStageRollbackPolicyRepresentation},
			"load_balancer_config":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: deployStageLoadBalancerInstanceGroupConfigRepresentation},
			"failure_policy":                               acctest.RepresentationGroup{RepType: acctest.Optional, Group: deployComputeInstanceStageFailurePolicyRepresentation},
			"lifecycle":                                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
		}))

	deployComputeInstanceStageRolloutPolicyRepresentation = map[string]interface{}{
		"policy_type":            acctest.Representation{RepType: acctest.Required, Create: `COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT`},
		"batch_delay_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `5`},
		"batch_count":            acctest.Representation{RepType: acctest.Required, Create: `5`},
	}

	deployComputeInstanceStageFailurePolicyRepresentation = map[string]interface{}{
		"policy_type":   acctest.Representation{RepType: acctest.Required, Create: `COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT`},
		"failure_count": acctest.Representation{RepType: acctest.Required, Create: `1`},
	}

	deployStageLoadBalancerInstanceGroupConfigRepresentation = map[string]interface{}{
		"backend_port":     acctest.Representation{RepType: acctest.Optional, Create: `8080`},
		"listener_name":    acctest.Representation{RepType: acctest.Required, Create: `LoadBalancerListener`, Update: `LoadBalancerListener2`},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer_1.id}`, Update: `${oci_load_balancer_load_balancer.test_load_balancer_2.id}`},
	}

	DeployComputeInstanceGroupStageResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_generic_artifact", acctest.Required, acctest.Create, deployGenericArtifactRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_instance_group_environment", acctest.Required, acctest.Create, deployInstanceGroupEnvironmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, DevopsDeployPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer_1", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer_2", acctest.Optional, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerResourceDependencies
)

// issue-routing-tag: devops/default
func TestDevopsDeployStageResource_computeInstanceGroupDeploy(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployStageResource_computeInstanceGroupDeploy")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_stage.test_deploy_stage"
	datasourceName := "data.oci_devops_deploy_stages.test_deploy_stages"
	singularDatasourceName := "data.oci_devops_deploy_stage.test_deploy_stage"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeployComputeInstanceGroupStageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Create, deployComputeInstanceGroupStageRepresentation), "devops", "deployStage", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployStageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployComputeInstanceGroupStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.batch_count", "5"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_group_deploy_environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_spec_deploy_artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Create, deployComputeInstanceGroupStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.batch_count", "5"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.batch_delay_in_seconds", "5"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_group_deploy_environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_spec_deploy_artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_config.0.listener_name", "LoadBalancerListener"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_config.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.0.failure_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"),

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
			Config: config + compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployComputeInstanceGroupStageRepresentation),
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
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.batch_count", "5"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.batch_delay_in_seconds", "5"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_group_deploy_environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_spec_deploy_artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_config.0.listener_name", "LoadBalancerListener2"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_config.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.0.failure_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stages", "test_deploy_stages", acctest.Optional, acctest.Update, DevopsDevopsDeployStageDataSourceRepresentation) +
				compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployComputeInstanceGroupStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployComputeInstanceGroupStageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_type", "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollout_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollout_policy.0.batch_count", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollout_policy.0.batch_delay_in_seconds", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollout_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_instance_group_deploy_environment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_spec_deploy_artifact_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollback_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancer_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancer_config.0.listener_name", "LoadBalancerListener2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_config.0.load_balancer_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "failure_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "failure_policy.0.failure_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "failure_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"),
			),
		},
		// verify resource import
		{
			Config:                  config + DeployComputeInstanceGroupStageRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
