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
	DeployGenericArtifactRequiredOnlyResource = DeployArtifactResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployGenericArtifactRepresentation)

	DeployGenericArtifactResourceConfig = DeployArtifactResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Optional, Update, deployGenericArtifactRepresentation)

	deployGenericArtifactSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_artifact_id": Representation{repType: Required, create: `${oci_devops_deploy_artifact.test_deploy_artifact.id}`},
	}

	deployGenericArtifactRepresentation                     = getUpdatedRepresentationCopy("deploy_artifact_source", RepresentationGroup{Required, deployGenericArtifactDeployArtifactSourceRepresentation}, deployArtifactRepresentation)
	repository_id                                           = "ocid1.artifactrepository.oc1.iad.0.amaaaaaansx72maa7nbce5ebmsqkan3msgyosvxe5d5a6jghn5su6ykgw7vq"
	repository_id_updated                                   = "ocid1.artifactrepository.oc1.iad.0.amaaaaaansx72maa7nbce5ebmsqkan3msgyosvxe5d5a6jghnfakeartifact2"
	artifact_path                                           = "helloworld-oke.yaml"
	version                                                 = "v1"
	deployGenericArtifactDeployArtifactSourceRepresentation = map[string]interface{}{
		"deploy_artifact_source_type": Representation{repType: Required, create: `GENERIC_ARTIFACT`},
		"repository_id":               Representation{repType: Required, create: repository_id, update: repository_id_updated},
		"deploy_artifact_path":        Representation{repType: Required, create: artifact_path},
		"deploy_artifact_version":     Representation{repType: Required, create: version},
	}
)

// issue-routing-tag: devops/default
func TestDevopsDeployArtifactResource_generic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployArtifactResource_generic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_artifact.test_deploy_artifact"
	datasourceName := "data.oci_devops_deploy_artifacts.test_deploy_artifacts"
	singularDatasourceName := "data.oci_devops_deploy_artifact.test_deploy_artifact"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DeployArtifactResourceDependencies+
		generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Optional, Create, deployGenericArtifactRepresentation), "devops", "deployArtifact", t)

	ResourceTest(t, testAccCheckDevopsDeployArtifactDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + DeployArtifactResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployGenericArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.repository_id", repository_id),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_path", artifact_path),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_version", version),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "GENERIC_ARTIFACT"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + DeployArtifactResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployArtifactResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Optional, Create, deployGenericArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.repository_id", repository_id),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "GENERIC_ARTIFACT"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
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
			Config: config + compartmentIdVariableStr + DeployArtifactResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Optional, Update, deployGenericArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.repository_id", repository_id_updated),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "GENERIC_ARTIFACT"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
				generateDataSourceFromRepresentationMap("oci_devops_deploy_artifacts", "test_deploy_artifacts", Optional, Update, deployArtifactDataSourceRepresentation) +
				compartmentIdVariableStr + DeployArtifactResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Optional, Update, deployGenericArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				generateDataSourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployGenericArtifactSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployGenericArtifactResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_artifact_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.repository_id", repository_id_updated),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "GENERIC_ARTIFACT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DeployGenericArtifactResourceConfig,
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
