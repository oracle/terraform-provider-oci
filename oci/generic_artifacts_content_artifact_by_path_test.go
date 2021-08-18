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
	ArtifactByPathResourceConfig = ArtifactByPathResourceDependencies +
		generateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create, artifactByPathRepresentation)

	artifactByPathSingularDataSourceRepresentation = map[string]interface{}{
		"artifact_path": Representation{repType: Required, create: `artifactPath`},
		"repository_id": Representation{repType: Required, create: `${oci_artifacts_repository.test_repository.id}`},
		"version":       Representation{repType: Required, create: `1.0`},
	}

	artifactByPathRepresentation = map[string]interface{}{
		"artifact_path": Representation{repType: Required, create: `artifactPath`},
		"repository_id": Representation{repType: Required, create: `${oci_artifacts_repository.test_repository.id}`},
		"version":       Representation{repType: Required, create: `1.0`},
		"content":       Representation{repType: Required, create: `<a1>content</a1>`},
	}

	ArtifactByPathResourceDependencies = generateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", Required, Create, repositoryRepresentation)
)

// issue-routing-tag: generic_artifacts_content/default
func TestGenericArtifactsContentArtifactByPathResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenericArtifactsContentArtifactByPathResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_generic_artifacts_content_artifact_by_path.test_artifact_by_path"

	singularDatasourceName := "data.oci_generic_artifacts_content_artifact_by_path.test_artifact_by_path"

	var resId, resId2 string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ArtifactByPathResourceDependencies+
		generateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create, artifactByPathRepresentation), "genericartifactscontent", "artifactByPath", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + ArtifactByPathResourceDependencies +
				generateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create, artifactByPathRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(

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
			Config: config + compartmentIdVariableStr + ArtifactByPathResourceDependencies +
				generateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Optional, Update, artifactByPathRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", Required, Create, artifactByPathSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ArtifactByPathResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_path", "artifactPath"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version", "1.0"),
			),
		},
	})
}
