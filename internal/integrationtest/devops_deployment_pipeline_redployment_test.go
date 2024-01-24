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
	DevopsPipelineRedeploymentRequiredOnlyResource = DevopsPipelineRedeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", acctest.Required, acctest.Create, devopsPipelineRedeploymentRepresentation)

	DevopsPipelineRedeploymentResourceConfig = DevopsPipelineRedeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", acctest.Optional, acctest.Update, devopsPipelineRedeploymentRepresentation)

	devopsPipelineRedeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deployment.test_pipeline_redeployment.id}`},
	}

	devopsPipelineRedeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"deploy_pipeline_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deployment.test_pipeline_redeployment.id}`},
		"project_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: devopsPipelineRedeploymentDataSourceFilterRepresentation}}
	devopsPipelineRedeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_deployment.test_pipeline_redeployment.id}`}},
	}

	devopsPipelineRedeploymentRepresentation = acctest.GetUpdatedRepresentationCopy("deployment_type", acctest.Representation{RepType: acctest.Required, Create: `PIPELINE_REDEPLOYMENT`},
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DevopsDeploymentRepresentation, []string{"deployment_arguments"}), map[string]interface{}{
			"previous_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deployment.test_deploy_1.id}`},
		}))

	DevopsPipelineRedeploymentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, DevopsDeployArtifactRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, DevopsdeployEnvironmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, DevopsDeployPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deploy_1", acctest.Required, acctest.Create, DevopsDeploymentRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_devops_log_group", acctest.Required, acctest.Create, DevopsLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Create, deployLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeploymentResource_pipelineRedeployment(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeploymentResource_pipelineRedeployment")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deployment.test_pipeline_redeployment"
	datasourceName := "data.oci_devops_deployments.test_pipeline_redeployments"
	singularDatasourceName := "data.oci_devops_deployment.test_pipeline_redeployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsPipelineRedeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", acctest.Optional, acctest.Create, devopsPipelineRedeploymentRepresentation), "devops", "deployment", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsPipelineRedeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", acctest.Required, acctest.Create, devopsPipelineRedeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_REDEPLOYMENT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsPipelineRedeploymentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsPipelineRedeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", acctest.Optional, acctest.Create, devopsPipelineRedeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttrSet(resourceName, "previous_deployment_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_REDEPLOYMENT"),
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
			Config: config + compartmentIdVariableStr + DevopsPipelineRedeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", acctest.Optional, acctest.Update, devopsPipelineRedeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttrSet(resourceName, "previous_deployment_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_REDEPLOYMENT"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deployments", "test_pipeline_redeployments", acctest.Optional, acctest.Update, devopsPipelineRedeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsPipelineRedeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", acctest.Optional, acctest.Update, devopsPipelineRedeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),

				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", acctest.Required, acctest.Create, devopsPipelineRedeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsPipelineRedeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "previous_deployment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_execution_progress.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_type", "PIPELINE_REDEPLOYMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DevopsPipelineRedeploymentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
