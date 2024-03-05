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
	DeployHelmArtifactRequiredOnlyResource = DevopsDeployArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployHelmChartArtifactRepresentation)

	DeployHelmArtifactResourceConfig = DevopsDeployArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, deployHelmChartArtifactRepresentation)

	DeployHelmCommandSpecArtifactResourceConfig = DevopsDeployArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployHelmCommandSpecArtifactRepresentation)

	deployHelmArtifactSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_artifact_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_artifact.test_deploy_artifact.id}`},
	}

	deployHelmChartArtifactRepresentation = map[string]interface{}{
		"argument_substitution_mode": acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `SUBSTITUTE_PLACEHOLDERS`},
		"deploy_artifact_source":     acctest.RepresentationGroup{RepType: acctest.Required, Group: deployHelmChartArtifactDeployArtifactSourceRepresentation},
		"deploy_artifact_type":       acctest.Representation{RepType: acctest.Required, Create: `HELM_CHART`},
		"project_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	deployHelmChartArtifactDeployArtifactSourceRepresentation = map[string]interface{}{
		"deploy_artifact_source_type": acctest.Representation{RepType: acctest.Required, Create: `HELM_CHART`},
		"chart_url":                   acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithDefault("helm_chart_url_static_resource", "oci://us-ashburn-1.ocir.io/idkxrdu9epyt/helm-test-chart"), Update: utils.GetEnvSettingWithDefault("helm_chart_url_update_static_resource", "oci://iad.ocir.io/ax022wvgmjpq/helm-charts-testing/helm-chart-nginx")},
		"deploy_artifact_version":     acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithDefault("helm_deploy_artifact_version_static_resource", "0.0.1"), Update: utils.GetEnvSettingWithDefault("helm_deploy_artifact_version_update_static_resource", "1.0.6")},
	}

	deployHelmCommandSpecArtifactRepresentation = map[string]interface{}{
		"argument_substitution_mode": acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `SUBSTITUTE_PLACEHOLDERS`},
		"deploy_artifact_source":     acctest.RepresentationGroup{RepType: acctest.Required, Group: deployHelmCommandSpecArtifactDeployArtifactSourceRepresentation},
		"deploy_artifact_type":       acctest.Representation{RepType: acctest.Required, Create: `HELM_COMMAND_SPEC`},
		"project_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	base64EncodeHelmCommandSpec                                     = "aGVsbSBscyAtYQpoZWxtIHN0YXR1cyAtYQ=="
	deployHelmCommandSpecArtifactDeployArtifactSourceRepresentation = map[string]interface{}{
		"deploy_artifact_source_type": acctest.Representation{RepType: acctest.Required, Create: `INLINE`},
		"helm_artifact_source_type":   acctest.Representation{RepType: acctest.Required, Create: `INLINE`},
		"base64encoded_content":       acctest.Representation{RepType: acctest.Required, Create: base64EncodeHelmCommandSpec},
	}
)

// issue-routing-tag: devops/default
func TestDevopsDeployArtifactResource_helm(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployArtifactResource_helm")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	chartUrl := utils.GetEnvSettingWithDefault("helm_chart_url_static_resource", "oci://us-ashburn-1.ocir.io/idkxrdu9epyt/helm-test-chart")
	artifactVersion := utils.GetEnvSettingWithDefault("helm_deploy_artifact_version_static_resource", "0.0.1")
	chartUrlUpdated := utils.GetEnvSettingWithDefault("helm_chart_url_update_static_resource", "oci://iad.ocir.io/ax022wvgmjpq/helm-charts-testing/helm-chart-nginx")
	artifactVersionUpdated := utils.GetEnvSettingWithDefault("helm_deploy_artifact_version_update_static_resource", "1.0.6")

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_artifact.test_deploy_artifact"
	datasourceName := "data.oci_devops_deploy_artifacts.test_deploy_artifacts"
	singularDatasourceName := "data.oci_devops_deploy_artifact.test_deploy_artifact"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsDeployArtifactResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Create, deployHelmChartArtifactRepresentation), "devops", "deployArtifact", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployArtifactDestroy, []resource.TestStep{
		/*
			// verify Create (Helm command Spec)
			{
				Config: config + compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployHelmCommandSpecArtifactRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.base64encoded_content", base64EncodeHelmCommandSpec),
					resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "HELM_COMMAND_SPEC"),
					resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_artifact_source_type", "INLINE"),
					resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "HELM_COMMAND_SPEC"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies,
			},
		*/
		// verify Create (Helm Chart)
		{
			Config: config + compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployHelmChartArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.chart_url", chartUrl),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_version", artifactVersion),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "HELM_CHART"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "HELM_CHART"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Create, deployHelmChartArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.chart_url", chartUrl),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_version", artifactVersion),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "HELM_CHART"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "HELM_CHART"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
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
			Config: config + compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, deployHelmChartArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.chart_url", chartUrlUpdated),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_version", artifactVersionUpdated),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "HELM_CHART"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "HELM_CHART"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_artifacts", "test_deploy_artifacts", acctest.Optional, acctest.Update, DevopsDevopsDeployArtifactDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, deployHelmChartArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "deploy_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deploy_artifact_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployHelmArtifactSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployHelmArtifactResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_artifact_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.chart_url", chartUrlUpdated),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_version", artifactVersionUpdated),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "HELM_CHART"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "HELM_CHART"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DevopsDeployArtifactRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
