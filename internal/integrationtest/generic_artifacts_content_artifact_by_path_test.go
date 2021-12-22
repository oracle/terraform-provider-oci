// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ArtifactByPathResourceConfig = ArtifactByPathResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", acctest.Required, acctest.Create, artifactByPathRepresentation)

	artifactByPathSingularDataSourceRepresentation = map[string]interface{}{
		"artifact_path": acctest.Representation{RepType: acctest.Required, Create: `artifactPath`},
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_artifacts_repository.test_repository.id}`},
		"version":       acctest.Representation{RepType: acctest.Required, Create: `1.0`},
	}

	artifactByPathRepresentation = map[string]interface{}{
		"artifact_path": acctest.Representation{RepType: acctest.Required, Create: `artifactPath`},
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_artifacts_repository.test_repository.id}`},
		"version":       acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"content":       acctest.Representation{RepType: acctest.Required, Create: `<a1>content</a1>`},
	}

	ArtifactByPathResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_artifacts_repository", "test_repository", acctest.Required, acctest.Create, repositoryRepresentation)
	// the deletion of oci_generic_artifacts_content_artifact_by_path is done by oci_artifacts_generic_artifact
	GenericArtifactManager = acctest.GenerateResourceFromRepresentationMap("oci_artifacts_generic_artifact", "test_generic_artifact", acctest.Required, acctest.Create, genericArtifactRepresentation)
)

// issue-routing-tag: generic_artifacts_content/default
func TestGenericArtifactsContentArtifactByPathResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenericArtifactsContentArtifactByPathResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_generic_artifacts_content_artifact_by_path.test_artifact_by_path"

	singularDatasourceName := "data.oci_generic_artifacts_content_artifact_by_path.test_artifact_by_path"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ArtifactByPathResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", acctest.Required, acctest.Create, artifactByPathRepresentation), "genericartifactscontent", "artifactByPath", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ArtifactByPathResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", acctest.Required, acctest.Create, artifactByPathRepresentation) + GenericArtifactManager,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

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
			Config: config + compartmentIdVariableStr + ArtifactByPathResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", acctest.Optional, acctest.Update, artifactByPathRepresentation) + GenericArtifactManager,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", acctest.Required, acctest.Create, artifactByPathSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ArtifactByPathResourceConfig + GenericArtifactManager,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "artifact_path", "artifactPath"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version", "1.0"),
			),
		},
	})
}

const (
	tempFilePrefix = "small-"
	tempFileSize   = 2e5
	tempFileSha256 = "4cbbd9be0cba685835755f827758705db5a413c5494c34262cd25946a73e7582"
)

func createTmpFile() (string, error) {
	tempFile, err := ioutil.TempFile(os.TempDir(), tempFilePrefix)
	if err != nil {
		return "", err
	}
	if err := tempFile.Truncate(tempFileSize); err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

var (
	artifactByPathSourceRepresentation = map[string]interface{}{
		"artifact_path": acctest.Representation{RepType: acctest.Required, Create: `artifactPath`},
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_artifacts_repository.test_repository.id}`},
		"version":       acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":        acctest.Representation{RepType: acctest.Required, Create: ``},
	}
)

// issue-routing-tag: generic_artifacts_content/default
func TestGenericArtifactsContentArtifactByPathResource_uploadFile(t *testing.T) {
	httpreplay.SetScenario("TestGenericArtifactsContentArtifactByPathResource_uploadFile")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_generic_artifacts_content_artifact_by_path.test_artifact_by_path"

	tempFilePath, err := createTmpFile()
	if err != nil {
		t.Fatalf("Unable to Create file to upload. Error: %q", err)
	}

	var resId, _ string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ArtifactByPathResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", acctest.Required, acctest.Create, artifactByPathRepresentation), "genericartifactscontent", "artifactByPath", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + ArtifactByPathResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_generic_artifacts_content_artifact_by_path", "test_artifact_by_path", acctest.Required, acctest.Create,
						acctest.GetUpdatedRepresentationCopy("source", acctest.Representation{RepType: acctest.Required, Create: tempFilePath}, artifactByPathSourceRepresentation)) + GenericArtifactManager,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "sha256", tempFileSha256),
					resource.TestCheckResourceAttr(resourceName, "size_in_bytes", strconv.Itoa(tempFileSize)),
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
		},
	})
}
