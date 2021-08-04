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
		generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Required, Create, devopsDeploymentRepresentation)

	DevopsDeploymentResourceConfig = DevopsDeploymentResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Update, devopsDeploymentRepresentation)

	devopsDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": Representation{repType: Required, create: `${oci_devops_deployment.test_deployment.id}`},
	}

	devopsDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Optional, create: `${var.compartment_id}`},
		"deploy_pipeline_id": Representation{repType: Optional, create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"display_name":       Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"id":                 Representation{repType: Optional, create: `${oci_devops_deployment.test_deployment.id}`},
		"project_id":         Representation{repType: Optional, create: `${oci_devops_project.test_project.id}`},
		"filter":             RepresentationGroup{Required, devopsDeploymentDataSourceFilterRepresentation}}
	devopsDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_devops_deployment.test_deployment.id}`}},
	}

	devopsDeploymentRepresentation = map[string]interface{}{
		"deploy_pipeline_id":   Representation{repType: Required, create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"deployment_type":      Representation{repType: Required, create: `PIPELINE_DEPLOYMENT`},
		"defined_tags":         Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"deployment_arguments": RepresentationGroup{Optional, deploymentDeploymentArgumentsRepresentation},
		"display_name":         Representation{repType: Optional, create: `displayName`},
		"freeform_tags":        Representation{repType: Optional, create: map[string]string{"bar-key": "value"}},
		"lifecycle":            RepresentationGroup{Required, ignoreDefinedTagsDifferencesRepresentation},
	}
	deploymentDeploymentArgumentsRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Optional, deploymentDeploymentArgumentsItemsRepresentation},
	}
	deploymentDeploymentArgumentsItemsRepresentation = map[string]interface{}{
		"name":  Representation{repType: Optional, create: `name`},
		"value": Representation{repType: Optional, create: `value`},
	}

	devopLogConfigurationRepresentation = map[string]interface{}{
		"source": RepresentationGroup{Required, devopLogConfigurationSourceRepresentation},
	}
	devopLogConfigurationSourceRepresentation = map[string]interface{}{
		"category":    Representation{repType: Required, create: `all`},
		"resource":    Representation{repType: Required, create: `${oci_devops_project.test_project.id}`},
		"service":     Representation{repType: Required, create: `devops`},
		"source_type": Representation{repType: Required, create: `OCISERVICE`},
	}

	DevopsDeploymentResourceDependencies = generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployArtifactRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Required, Create, deployEnvironmentRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", Required, Create, deployPipelineRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_logging_log", "test_log", Optional, Create, deployLogRepresentation) +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deployment.test_deployment"
	datasourceName := "data.oci_devops_deployments.test_deployments"
	singularDatasourceName := "data.oci_devops_deployment.test_deployment"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DevopsDeploymentResourceDependencies+
		generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Create, devopsDeploymentRepresentation), "devops", "deployment", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Required, Create, devopsDeploymentRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_DEPLOYMENT"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DevopsDeploymentResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Create, devopsDeploymentRepresentation),
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
				Config: config + compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Update, devopsDeploymentRepresentation),
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
					generateDataSourceFromRepresentationMap("oci_devops_deployments", "test_deployments", Optional, Update, devopsDeploymentDataSourceRepresentation) +
					compartmentIdVariableStr + DevopsDeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Optional, Update, devopsDeploymentRepresentation),
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
					generateDataSourceFromRepresentationMap("oci_devops_deployment", "test_deployment", Required, Create, devopsDeploymentSingularDataSourceRepresentation) +
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
		},
	})
}
