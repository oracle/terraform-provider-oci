package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DeployComputeInstanceCanaryApprovalStageRequiredOnlyResource = DeployComputeInstanceCanaryApprovalStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployComputeInstanceCanaryApprovalStageRepresentation)

	DeployInstanceGroupCanaryApprovalStageResourceConfig = DeployComputeInstanceCanaryApprovalStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployComputeInstanceCanaryApprovalStageRepresentation)

	deployComputeInstanceCanaryApprovalStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	deployComputeInstanceCanaryApprovalStageRepresentation = acctest.GetUpdatedRepresentationCopy("deploy_stage_type", acctest.Representation{RepType: acctest.Required, Create: `COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL`},
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(deployStageRepresentation, []string{"wait_criteria", "deploy_stage_predecessor_collection"}), map[string]interface{}{
			"approval_policy": acctest.RepresentationGroup{RepType: acctest.Required, Group: deployComputeInstanceCanaryStageApprovalPolicyRepresentation},
			"compute_instance_group_canary_traffic_shift_deploy_stage_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_stage.test_compute_instance_canary_traffic_shift_stage.id}`},
			"deploy_stage_predecessor_collection":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: deployStageCanaryComputeInstanceApprovalPredecessorCollectionRepresentation},
		}))

	deployComputeInstanceCanaryStageApprovalPolicyRepresentation = map[string]interface{}{
		"approval_policy_type":         acctest.Representation{RepType: acctest.Required, Create: `COUNT_BASED_APPROVAL`},
		"number_of_approvals_required": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
	}

	deployStageCanaryComputeInstanceApprovalPredecessorCollectionRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: deployStageCanaryComputeInstanceApprovalPredecessorCollectionItemsRepresentation},
	}

	deployStageCanaryComputeInstanceApprovalPredecessorCollectionItemsRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_stage.test_compute_instance_canary_traffic_shift_stage.id}`},
	}

	DeployComputeInstanceCanaryApprovalStageResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_compute_instance_canary_deploy_stage", acctest.Required, acctest.Create, computeInstanceGroupCanaryStageRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_compute_instance_canary_traffic_shift_stage", acctest.Required, acctest.Create, computeInstanceGroupCanaryTrafficShiftStageRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_generic_artifact", acctest.Required, acctest.Create, deployGenericArtifactRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_instance_group_environment", acctest.Required, acctest.Create, deployInstanceGroupEnvironmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, deployPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer_1", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer_2", acctest.Optional, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerResourceDependencies
)

// issue-routing-tag: devops/default
func TestDevopsDeployStageResource_computeInstanceGroupCanaryApproval(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployStageResource_computeInstanceGroupCanaryApproval")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_stage.test_deploy_stage"
	datasourceName := "data.oci_devops_deploy_stages.test_deploy_stages"
	singularDatasourceName := "data.oci_devops_deploy_stage.test_deploy_stage"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeployComputeInstanceCanaryApprovalStageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Create, deployComputeInstanceCanaryApprovalStageRepresentation), "devops", "deployStage", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployStageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployComputeInstanceCanaryApprovalStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployComputeInstanceCanaryApprovalStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.approval_policy_type", "COUNT_BASED_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.number_of_approvals_required", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_group_canary_traffic_shift_deploy_stage_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeployComputeInstanceCanaryApprovalStageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployComputeInstanceCanaryApprovalStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Create, deployComputeInstanceCanaryApprovalStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.approval_policy_type", "COUNT_BASED_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.number_of_approvals_required", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_group_canary_traffic_shift_deploy_stage_id"),

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
			Config: config + compartmentIdVariableStr + DeployComputeInstanceCanaryApprovalStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployComputeInstanceCanaryApprovalStageRepresentation),
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
				resource.TestCheckResourceAttr(resourceName, "approval_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.approval_policy_type", "COUNT_BASED_APPROVAL"),
				resource.TestCheckResourceAttr(resourceName, "approval_policy.0.number_of_approvals_required", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_group_canary_traffic_shift_deploy_stage_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stages", "test_deploy_stages", acctest.Optional, acctest.Update, deployStageDataSourceRepresentation) +
				compartmentIdVariableStr + DeployComputeInstanceCanaryApprovalStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployComputeInstanceCanaryApprovalStageRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployComputeInstanceCanaryApprovalStageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployInstanceGroupCanaryApprovalStageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_type", "COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approval_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approval_policy.0.approval_policy_type", "COUNT_BASED_APPROVAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approval_policy.0.number_of_approvals_required", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_instance_group_canary_traffic_shift_deploy_stage_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + DeployComputeInstanceCanaryApprovalStageRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
