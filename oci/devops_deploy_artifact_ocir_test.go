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
	DeployOcirArtifactRequiredOnlyResource = DeployArtifactResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployOcirArtifactRepresentation)

	DeployOcirArtifactResourceConfig = DeployArtifactResourceDependencies +
		generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Optional, Update, deployOcirArtifactRepresentation)

	deployOcirArtifactSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_artifact_id": Representation{repType: Required, create: `${oci_devops_deploy_artifact.test_deploy_artifact.id}`},
	}

	deployOcirArtifactRepresentation = getMultipleUpdatedRepresenationCopy(
		[]string{"deploy_artifact_type", "deploy_artifact_source"},
		[]interface{}{Representation{repType: Required, create: `DOCKER_IMAGE`}, RepresentationGroup{Required, deployOcirArtifactDeployArtifactSourceRepresentation}},
		deployArtifactRepresentation)
	image_uri                                            = "iad.ocir.io/ax022wvgmjpq/fake/hello-java:0.0.2"
	image_uri_updated                                    = "iad.ocir.io/ax022wvgmjpq/fake/hello-java:0.0.3"
	image_digest                                         = "38598585.fakedigest1"
	image_digest_updated                                 = "38598585.fakedigest2"
	deployOcirArtifactDeployArtifactSourceRepresentation = map[string]interface{}{
		"deploy_artifact_source_type": Representation{repType: Required, create: `OCIR`},
		"image_uri":                   Representation{repType: Required, create: image_uri, update: image_uri_updated},
		"image_digest":                Representation{repType: Optional, create: image_digest, update: image_digest_updated},
	}
)

// issue-routing-tag: devops/default
func TestDevopsDeployArtifactResource_ocir(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployArtifactResource_ocir")
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
		generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Optional, Create, deployOcirArtifactRepresentation), "devops", "deployArtifact", t)

	ResourceTest(t, testAccCheckDevopsDeployArtifactDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + DeployArtifactResourceDependencies +
				generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployOcirArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.image_uri", image_uri),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "OCIR"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "DOCKER_IMAGE"),
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
				generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Optional, Create, deployOcirArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.image_uri", image_uri),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.image_digest", image_digest),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "OCIR"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "DOCKER_IMAGE"),
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
				generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Optional, Update, deployOcirArtifactRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.image_uri", image_uri_updated),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.image_digest", image_digest_updated),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "OCIR"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "DOCKER_IMAGE"),
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
				generateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Optional, Update, deployOcirArtifactRepresentation),
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
				generateDataSourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", Required, Create, deployOcirArtifactSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployOcirArtifactResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_artifact_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.image_uri", image_uri_updated),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.image_digest", image_digest_updated),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "OCIR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_type", "DOCKER_IMAGE"),
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
			Config: config + compartmentIdVariableStr + DeployOcirArtifactResourceConfig,
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
