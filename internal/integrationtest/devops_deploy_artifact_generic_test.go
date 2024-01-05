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
	DeployGenericArtifactRequiredOnlyResource = DevopsDeployArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployGenericArtifactRepresentation)

	DeployGenericArtifactResourceConfig = DevopsDeployArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, deployGenericArtifactRepresentation)

	deployGenericArtifactSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_artifact_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_artifact.test_deploy_artifact.id}`},
	}

	deployGenericArtifactRepresentation                     = acctest.GetUpdatedRepresentationCopy("deploy_artifact_source", acctest.RepresentationGroup{RepType: acctest.Required, Group: deployGenericArtifactDeployArtifactSourceRepresentation}, DevopsDeployArtifactRepresentation)
	repository_id                                           = "ocid1.artifactrepository.oc1.iad.0.amaaaaaamdb2enyazr72hrxgh4zcffpzrddr7tqe2376vgvfodiavftlavpq"
	repository_id_updated                                   = "ocid1.artifactrepository.oc1.iad.0.amaaaaaansx72maa7nbce5ebmsqkan3msgyosvxe5d5a6jghnfakeartifact2"
	artifact_path                                           = "helloworld-oke.yaml"
	version                                                 = "v1"
	deployGenericArtifactDeployArtifactSourceRepresentation = map[string]interface{}{
		"deploy_artifact_source_type": acctest.Representation{RepType: acctest.Required, Create: `GENERIC_ARTIFACT`},
		"repository_id":               acctest.Representation{RepType: acctest.Required, Create: repository_id, Update: repository_id_updated},
		"deploy_artifact_path":        acctest.Representation{RepType: acctest.Required, Create: artifact_path},
		"deploy_artifact_version":     acctest.Representation{RepType: acctest.Required, Create: version},
	}
)

// issue-routing-tag: devops/default
func TestDevopsDeployArtifactResource_generic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployArtifactResource_generic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_artifact.test_deploy_artifact"
	datasourceName := "data.oci_devops_deploy_artifacts.test_deploy_artifacts"
	singularDatasourceName := "data.oci_devops_deploy_artifact.test_deploy_artifact"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsDeployArtifactResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Create, deployGenericArtifactRepresentation), "devops", "deployArtifact", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployArtifactDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeployArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployGenericArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.repository_id", repository_id),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_path", artifact_path),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_version", version),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "GENERIC_ARTIFACT"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
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
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Create, deployGenericArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.repository_id", repository_id),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "GENERIC_ARTIFACT"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
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
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, deployGenericArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.repository_id", repository_id_updated),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "GENERIC_ARTIFACT"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
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
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, deployGenericArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployGenericArtifactSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployGenericArtifactResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_artifact_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.repository_id", repository_id_updated),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "GENERIC_ARTIFACT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
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
			Config:                  config + DeployGenericArtifactRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
