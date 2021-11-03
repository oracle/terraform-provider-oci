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
	DevopsDeploymentRequiredOnlyResource = DevopsDeploymentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Required, Create, devopsDeploymentRepresentation)

	DevopsDeploymentResourceConfig = DevopsDeploymentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Update, devopsDeploymentRepresentation)

	devopsDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": Representation{RepType: Required, Create: `${oci_devops_deployment.test_deployment.id}`},
	}

	devopsDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"deploy_pipeline_id": Representation{RepType: Optional, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"display_name":       Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":                 Representation{RepType: Optional, Create: `${oci_devops_deployment.test_deployment.id}`},
		"project_id":         Representation{RepType: Optional, Create: `${oci_devops_project.test_project.id}`},
		"filter":             RepresentationGroup{Required, devopsDeploymentDataSourceFilterRepresentation}}
	devopsDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_devops_deployment.test_deployment.id}`}},
	}

	devopsDeploymentRepresentation = map[string]interface{}{
		"deploy_pipeline_id":   Representation{RepType: Required, Create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"deployment_type":      Representation{RepType: Required, Create: `PIPELINE_DEPLOYMENT`},
		"defined_tags":         Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"deployment_arguments": RepresentationGroup{Optional, deploymentDeploymentArgumentsRepresentation},
		"display_name":         Representation{RepType: Optional, Create: `displayName`},
		"freeform_tags":        Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}},
		"lifecycle":            RepresentationGroup{Required, ignoreDefinedTagsDifferencesRepresentation},
	}
	deploymentDeploymentArgumentsRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Optional, deploymentDeploymentArgumentsItemsRepresentation},
	}
	deploymentDeploymentArgumentsItemsRepresentation = map[string]interface{}{
		"name":  Representation{RepType: Optional, Create: `name`},
		"value": Representation{RepType: Optional, Create: `value`},
	}

	devopLogConfigurationRepresentation = map[string]interface{}{
		"source": RepresentationGroup{Required, devopLogConfigurationSourceRepresentation},
	}
	devopLogConfigurationSourceRepresentation = map[string]interface{}{
		"category":    Representation{RepType: Required, Create: `all`},
		"resource":    Representation{RepType: Required, Create: `${oci_devops_project.test_project.id}`},
		"service":     Representation{RepType: Required, Create: `devops`},
		"source_type": Representation{RepType: Required, Create: `OCISERVICE`},
	}

	DevopsDeploymentResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployArtifactRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Required, Create, deployEnvironmentRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", Required, Create, deployPipelineRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_devops_log_group", Required, Create, devopsLogGroupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", Optional, Create, deployLogRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deployment.test_deployment"
	datasourceName := "data.oci_devops_deployments.test_deployments"
	singularDatasourceName := "data.oci_devops_deployment.test_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DevopsDeploymentResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Create, devopsDeploymentRepresentation), "devops", "deployment", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Required, Create, devopsDeploymentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_DEPLOYMENT"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Create, devopsDeploymentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_DEPLOYMENT"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
			Config: config + compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Update, devopsDeploymentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "deployment_arguments.0.items.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_DEPLOYMENT"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
				GenerateDataSourceFromRepresentationMap("oci_devops_deployments", "test_deployments", Optional, Update, devopsDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Update, devopsDeploymentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Required, Create, devopsDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsDeploymentResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
