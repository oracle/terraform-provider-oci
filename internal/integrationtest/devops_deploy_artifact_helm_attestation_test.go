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
	vaultSecretDefault = "vaultsecret.amaaaaaansx72maakfv2jt6zqypdfv7meqe33we43g6ujlezle3xypj4zsmq"
	vaultSecret        = utils.GetEnvSettingWithDefault("helm_attestation_public_key_secret", vaultSecretDefault)

	DevopsDeployArtifactHelmVerificationKeySourceRepresentation = map[string]interface{}{
		"verification_key_source_type": acctest.Representation{RepType: acctest.Optional, Create: `VAULT_SECRET`, Update: `INLINE_PUBLIC_KEY`},
		"current_public_key":           acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `currentPublicKey`},
		"previous_public_key":          acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `previousPublicKey`},
		"vault_secret_id":              acctest.Representation{RepType: acctest.Optional, Create: vaultSecret, Update: ``},
	}

	deployHelmAttestationArtifactDeployArtifactSourceRepresentation = map[string]interface{}{
		"deploy_artifact_source_type":  acctest.Representation{RepType: acctest.Required, Create: `HELM_CHART`},
		"chart_url":                    acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("helm_chart_url_static_resource"), Update: utils.GetEnvSettingWithBlankDefault("helm_chart_url_update_static_resource")},
		"deploy_artifact_version":      acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("helm_deploy_artifact_version_static_resource"), Update: utils.GetEnvSettingWithBlankDefault("helm_deploy_artifact_version_update_static_resource")},
		"helm_verification_key_source": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DevopsDeployArtifactHelmVerificationKeySourceRepresentation},
	}

	deployHelmAttestationArtifactRepresentation = map[string]interface{}{
		"argument_substitution_mode": acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `SUBSTITUTE_PLACEHOLDERS`},
		"deploy_artifact_source":     acctest.RepresentationGroup{RepType: acctest.Required, Group: deployHelmAttestationArtifactDeployArtifactSourceRepresentation},
		"deploy_artifact_type":       acctest.Representation{RepType: acctest.Required, Create: `HELM_CHART`},
		"project_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	DeployHelmAttestationArtifactRequiredOnlyResource = DevopsDeployArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "helm_chart_pubkey_artifact", acctest.Required, acctest.Create, deployHelmAttestationArtifactRepresentation)

	DeployHelmAttestationArtifactCreateKeySourceResource = DevopsDeployArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "helm_chart_pubkey_artifact", acctest.Optional, acctest.Create, deployHelmAttestationArtifactRepresentation)

	DeployHelmAttestationArtifactUpdateResource = DevopsDeployArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "helm_chart_pubkey_artifact", acctest.Optional, acctest.Update, deployHelmAttestationArtifactRepresentation)

	helmAttestationSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_artifact_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_artifact.helm_chart_pubkey_artifact.id}`},
	}
)

// issue-routing-tag: devops/default
func TestDevopsDeployArtifactResource_helmattestation(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployArtifactResource_helmattestation")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	chartUrl := utils.GetEnvSettingWithBlankDefault("helm_chart_url_static_resource")
	artifactVersion := utils.GetEnvSettingWithBlankDefault("helm_deploy_artifact_version_static_resource")
	chartUrlUpdated := utils.GetEnvSettingWithBlankDefault("helm_chart_url_update_static_resource")
	artifactVersionUpdated := utils.GetEnvSettingWithBlankDefault("helm_deploy_artifact_version_update_static_resource")

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_artifact.helm_chart_pubkey_artifact"
	datasourceName := "data.oci_devops_deploy_artifacts.helm_chart_pubkey_artifacts"
	singularDatasourceName := "data.oci_devops_deploy_artifact.helm_chart_pubkey_artifact"

	var resId, resId2 string
	_ = resId2
	_ = datasourceName
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	tfContentToCreateWithKeySource := config + compartmentIdVariableStr + DeployHelmAttestationArtifactCreateKeySourceResource
	acctest.SaveConfigContent(tfContentToCreateWithKeySource, "devops", "deployHelmChartPubKeyArtifact", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployArtifactDestroy, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + DeployHelmAttestationArtifactRequiredOnlyResource,
			//acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "helm_chart_pubkey_artifact", acctest.Required, acctest.Create, deployHelmAttestationArtifactRepresentation),
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

		{
			Config: config + compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies,
		},

		{
			Config: tfContentToCreateWithKeySource,
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
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_verification_key_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_verification_key_source.0.current_public_key", ""),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_verification_key_source.0.previous_public_key", ""),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_verification_key_source.0.vault_secret_id", vaultSecret),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_verification_key_source.0.verification_key_source_type", "VAULT_SECRET"),

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
			Config: config + compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "helm_chart_pubkey_artifact", acctest.Optional, acctest.Update, deployHelmAttestationArtifactRepresentation),
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
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_verification_key_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_verification_key_source.0.current_public_key", "currentPublicKey"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_verification_key_source.0.previous_public_key", "previousPublicKey"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_verification_key_source.0.vault_secret_id", ""),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.helm_verification_key_source.0.verification_key_source_type", "INLINE_PUBLIC_KEY"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_artifacts", "helm_chart_pubkey_artifacts", acctest.Optional, acctest.Update, DevopsDeployArtifactDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies +
				//acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "helm_chart_pubkey_artifact", acctest.Optional, acctest.Update, deployHelmAttestationArtifactRepresentation),
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, deployHelmAttestationArtifactRepresentation),
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

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_artifact", "helm_chart_pubkey_artifact", acctest.Required, acctest.Create, helmAttestationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployHelmAttestationArtifactUpdateResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_artifact_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.chart_url", chartUrlUpdated),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.deploy_artifact_version", artifactVersionUpdated),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "HELM_CHART"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_type", "HELM_CHART"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.helm_verification_key_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.helm_verification_key_source.0.current_public_key", "currentPublicKey"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.helm_verification_key_source.0.previous_public_key", "previousPublicKey"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.helm_verification_key_source.0.vault_secret_id", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.helm_verification_key_source.0.verification_key_source_type", "INLINE_PUBLIC_KEY"),
			),
		},

		{
			Config:                  config + DeployHelmAttestationArtifactRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
