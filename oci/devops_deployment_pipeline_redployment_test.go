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
	DevopsPipelineRedeploymentRequiredOnlyResource = DevopsPipelineRedeploymentResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", Required, Create, devopsPipelineRedeploymentRepresentation)

	DevopsPipelineRedeploymentResourceConfig = DevopsPipelineRedeploymentResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", Optional, Update, devopsPipelineRedeploymentRepresentation)

	devopsPipelineRedeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": Representation{repType: Required, create: `${oci_devops_deployment.test_pipeline_redeployment.id}`},
	}

	devopsPipelineRedeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Optional, create: `${var.compartment_id}`},
		"deploy_pipeline_id": Representation{repType: Optional, create: `${oci_devops_deploy_pipeline.test_deploy_pipeline.id}`},
		"display_name":       Representation{repType: Optional, create: `displayName`},
		"id":                 Representation{repType: Optional, create: `${oci_devops_deployment.test_pipeline_redeployment.id}`},
		"project_id":         Representation{repType: Optional, create: `${oci_devops_project.test_project.id}`},
		"filter":             RepresentationGroup{Required, devopsPipelineRedeploymentDataSourceFilterRepresentation}}
	devopsPipelineRedeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_devops_deployment.test_pipeline_redeployment.id}`}},
	}

	devopsPipelineRedeploymentRepresentation = getUpdatedRepresentationCopy("deployment_type", Representation{repType: Required, create: `PIPELINE_REDEPLOYMENT`},
		representationCopyWithNewProperties(representationCopyWithRemovedProperties(devopsDeploymentRepresentation, []string{"deployment_arguments"}), map[string]interface{}{
			"previous_deployment_id": Representation{repType: Required, create: `${oci_devops_deployment.test_deploy_1.id}`},
		}))

	DevopsPipelineRedeploymentResourceDependencies = generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployArtifactRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Required, Create, deployEnvironmentRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", Required, Create, deployPipelineRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		generateResourceFromRepresentationMap("oci_devops_deployment", "test_deploy_1", Required, Create, devopsDeploymentRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_logging_log", "test_log", Optional, Create, deployLogRepresentation) +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

func TestDevopsDeploymentResource_pipelineRedeployment(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeploymentResource_pipelineRedeployment")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deployment.test_pipeline_redeployment"
	datasourceName := "data.oci_devops_deployments.test_pipeline_redeployments"
	singularDatasourceName := "data.oci_devops_deployment.test_pipeline_redeployment"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DevopsPipelineRedeploymentResourceDependencies+
		generateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", Optional, Create, devopsPipelineRedeploymentRepresentation), "devops", "deployment", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DevopsPipelineRedeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", Required, Create, devopsPipelineRedeploymentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_REDEPLOYMENT"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DevopsPipelineRedeploymentResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DevopsPipelineRedeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", Optional, Create, devopsPipelineRedeploymentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttrSet(resourceName, "previous_deployment_id"),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_REDEPLOYMENT"),
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
				Config: config + compartmentIdVariableStr + DevopsPipelineRedeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", Optional, Update, devopsPipelineRedeploymentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
					resource.TestCheckResourceAttrSet(resourceName, "previous_deployment_id"),
					resource.TestCheckResourceAttr(resourceName, "deployment_type", "PIPELINE_REDEPLOYMENT"),
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
					generateDataSourceFromRepresentationMap("oci_devops_deployments", "test_pipeline_redeployments", Optional, Update, devopsPipelineRedeploymentDataSourceRepresentation) +
					compartmentIdVariableStr + DevopsPipelineRedeploymentResourceDependencies +
					generateResourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", Optional, Update, devopsPipelineRedeploymentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_devops_deployment", "test_pipeline_redeployment", Required, Create, devopsPipelineRedeploymentSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DevopsPipelineRedeploymentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "previous_deployment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "deployment_execution_progress.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "deployment_type", "PIPELINE_REDEPLOYMENT"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + DevopsPipelineRedeploymentResourceConfig,
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
