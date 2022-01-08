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
	DevopsDeploymentRequiredOnlyResource = DevopsDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Required, acctest.Create, devopsDeploymentRepresentation)

	DevopsDeploymentResourceConfig = DevopsDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Optional, acctest.Update, devopsDeploymentRepresentation)

	devopsDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deployment.test_deployment.id}`},
	}

	devopsDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"deploy_pipeline_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deployment.test_deployment.id}`},
		"project_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: devopsDeploymentDataSourceFilterRepresentation}}
	devopsDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_deployment.test_deployment.id}`}},
	}

	devopsDeploymentRepresentation = map[string]interface{}{
		"deploy_pipeline_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"deployment_type":      acctest.Representation{RepType: acctest.Required, Create: `PIPELINE_DEPLOYMENT`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"deployment_arguments": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentDeploymentArgumentsRepresentation},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	deploymentDeploymentArgumentsRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Optional, Group: deploymentDeploymentArgumentsItemsRepresentation},
	}
	deploymentDeploymentArgumentsItemsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`},
	}

	devopLogConfigurationRepresentation = map[string]interface{}{
		"source": acctest.RepresentationGroup{RepType: acctest.Required, Group: devopLogConfigurationSourceRepresentation},
	}
	devopLogConfigurationSourceRepresentation = map[string]interface{}{
		"category":    acctest.Representation{RepType: acctest.Required, Create: `all`},
		"resource":    acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"service":     acctest.Representation{RepType: acctest.Required, Create: `devops`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OCISERVICE`},
	}

	DevopsDeploymentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployArtifactRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, deployEnvironmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", acctest.Required, acctest.Create, deployPipelineRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_devops_log_group", acctest.Required, acctest.Create, devopsLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Optional, acctest.Create, deployLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deployment.test_deployment"
	datasourceName := "data.oci_devops_deployments.test_deployments"
	singularDatasourceName := "data.oci_devops_deployment.test_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsDeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Optional, acctest.Create, devopsDeploymentRepresentation), "devops", "deployment", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Required, acctest.Create, devopsDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_DEPLOYMENT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeploymentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Optional, acctest.Create, devopsDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_DEPLOYMENT"),
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
			Config: config + compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Optional, acctest.Update, devopsDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_DEPLOYMENT"),
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
				compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Optional, acctest.Update, devopsDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),

				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deployment", "test_deployment", acctest.Required, acctest.Create, devopsDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsDeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_arguments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_arguments.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_arguments.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_arguments.0.items.0.value", "value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_execution_progress.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_type", "PIPELINE_DEPLOYMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DevopsDeploymentResourceConfig,
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
