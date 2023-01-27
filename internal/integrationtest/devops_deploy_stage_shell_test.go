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
	DeployShellStageRequiredOnlyResource = DeployShellStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployShellStageRepresentation)

	DeployShellStageResourceConfig = DeployShellStageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployShellStageRepresentation)

	deployShellStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	DevopsDeployStageContainerConfigNetworkChannelRepresentation = map[string]interface{}{
		"network_channel_type": acctest.Representation{RepType: acctest.Required, Create: `SERVICE_VNIC_CHANNEL`},
		"subnet_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"nsg_ids":              acctest.Representation{RepType: acctest.Optional, Create: []string{}},
	}
	DevopsDeployStageContainerConfigShapeConfigRepresentation = map[string]interface{}{
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1.0`, Update: `2.0`},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `2.0`},
	}
	DevopsDeployStageContainerInstanceConfigRepresentation = map[string]interface{}{
		"container_config_type": acctest.Representation{RepType: acctest.Required, Create: `CONTAINER_INSTANCE_CONFIG`},
		"shape_name":            acctest.Representation{RepType: acctest.Required, Create: `CI.Standard.E3.Flex`, Update: `CI.Standard.E4.Flex`},
		"shape_config":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsDeployStageContainerConfigShapeConfigRepresentation},
		"network_channel":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsDeployStageContainerConfigNetworkChannelRepresentation},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	deployShellStageRepresentation = acctest.GetUpdatedRepresentationCopy("deploy_stage_type", acctest.Representation{RepType: acctest.Required, Create: `SHELL`},
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DevopsDeployStageRepresentation, []string{"wait_criteria"}), map[string]interface{}{
			"command_spec_deploy_artifact_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_artifact.test_deploy_command_spec_artifact.id}`},
			"timeout_in_seconds":              acctest.Representation{RepType: acctest.Optional, Create: `36000`},
			"container_config":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsDeployStageContainerInstanceConfigRepresentation},
		}))

	DeployShellStageResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_command_spec_artifact", acctest.Required, acctest.Create, DevopsDeployCommandSpecArtifactRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, DevopsDeployPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeployStageResource_deployShell(t *testing.T) {
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeployShellStageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Create, deployShellStageRepresentation), "devops", "deployStage", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployStageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployShellStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployShellStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "SHELL"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttrSet(resourceName, "command_spec_deploy_artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "container_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.shape_name", "CI.Standard.E3.Flex"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.shape_config.0.memory_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.network_channel.0.network_channel_type", "SERVICE_VNIC_CHANNEL"),
				resource.TestCheckResourceAttrSet(resourceName, "container_config.0.network_channel.0.subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeployShellStageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployShellStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Create, deployShellStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "SHELL"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "container_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.container_config_type", "CONTAINER_INSTANCE_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.shape_name", "CI.Standard.E3.Flex"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.shape_config.0.memory_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.network_channel.0.network_channel_type", "SERVICE_VNIC_CHANNEL"),
				resource.TestCheckResourceAttrSet(resourceName, "container_config.0.network_channel.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.network_channel.0.nsg_ids.#", "0"),

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
			Config: config + compartmentIdVariableStr + DeployShellStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployShellStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "SHELL"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "container_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.container_config_type", "CONTAINER_INSTANCE_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.shape_name", "CI.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.shape_config.0.memory_in_gbs", "2"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.network_channel.0.network_channel_type", "SERVICE_VNIC_CHANNEL"),
				resource.TestCheckResourceAttrSet(resourceName, "container_config.0.network_channel.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "container_config.0.network_channel.0.nsg_ids.#", "0"),

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
				compartmentIdVariableStr + DeployShellStageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Optional, acctest.Update, deployShellStageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", acctest.Required, acctest.Create, deployShellStageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployShellStageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "SHELL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "container_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "container_config.0.container_config_type", "CONTAINER_INSTANCE_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "container_config.0.shape_name", "CI.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(singularDatasourceName, "container_config.0.shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "container_config.0.shape_config.0.memory_in_gbs", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "container_config.0.network_channel.0.network_channel_type", "SERVICE_VNIC_CHANNEL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "container_config.0.network_channel.0.subnet_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "container_config.0.network_channel.0.nsg_ids.#", "0"),
			),
		},
		// verify resource import
		{
			Config:                  config + DevopsDeployStageRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
