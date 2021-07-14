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
	ModelProvenanceRequiredOnlyResource = ModelProvenanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Required, Create, modelProvenanceRepresentation)

	ModelProvenanceResourceConfig = ModelProvenanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Optional, Update, modelProvenanceRepresentation)

	modelProvenanceSingularDataSourceRepresentation = map[string]interface{}{
		"model_id": Representation{repType: Required, create: `${oci_datascience_model.test_model.id}`},
	}

	modelProvenanceRepresentation = map[string]interface{}{
		"model_id":        Representation{repType: Required, create: `${oci_datascience_model.test_model.id}`},
		"git_branch":      Representation{repType: Optional, create: `gitBranch`, update: `gitBranch2`},
		"git_commit":      Representation{repType: Optional, create: `gitCommit`, update: `gitCommit2`},
		"repository_url":  Representation{repType: Optional, create: `repositoryUrl`, update: `repositoryUrl2`},
		"script_dir":      Representation{repType: Optional, create: `scriptDir`, update: `scriptDir2`},
		"training_script": Representation{repType: Optional, create: `trainingScript`, update: `trainingScript2`},
	}

	ModelProvenanceResourceDependencies = generateResourceFromRepresentationMap("oci_datascience_model", "test_model", Required, Create, modelRepresentation) +
		generateResourceFromRepresentationMap("oci_datascience_project", "test_project", Required, Create, projectRepresentation)
)

func TestDatascienceModelProvenanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelProvenanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_model_provenance.test_model_provenance"

	singularDatasourceName := "data.oci_datascience_model_provenance.test_model_provenance"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ModelProvenanceResourceDependencies+
		generateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Optional, Create, modelProvenanceRepresentation), "datascience", "modelProvenance", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ModelProvenanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Required, Create, modelProvenanceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "model_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			{
				Config: config + compartmentIdVariableStr + ModelProvenanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Optional, Create, modelProvenanceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "git_branch", "gitBranch"),
					resource.TestCheckResourceAttr(resourceName, "git_commit", "gitCommit"),
					resource.TestCheckResourceAttrSet(resourceName, "model_id"),
					resource.TestCheckResourceAttr(resourceName, "repository_url", "repositoryUrl"),
					resource.TestCheckResourceAttr(resourceName, "script_dir", "scriptDir"),
					resource.TestCheckResourceAttr(resourceName, "training_script", "trainingScript"),

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
				Config: config + compartmentIdVariableStr + ModelProvenanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Optional, Update, modelProvenanceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "git_branch", "gitBranch2"),
					resource.TestCheckResourceAttr(resourceName, "git_commit", "gitCommit2"),
					resource.TestCheckResourceAttrSet(resourceName, "model_id"),
					resource.TestCheckResourceAttr(resourceName, "repository_url", "repositoryUrl2"),
					resource.TestCheckResourceAttr(resourceName, "script_dir", "scriptDir2"),
					resource.TestCheckResourceAttr(resourceName, "training_script", "trainingScript2"),

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
					generateDataSourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Required, Create, modelProvenanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ModelProvenanceResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "model_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "git_branch", "gitBranch2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "git_commit", "gitCommit2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "repository_url", "repositoryUrl2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "script_dir", "scriptDir2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "training_script", "trainingScript2"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ModelProvenanceResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportStateIdFunc:       getDatascienceModelProvenanceCompositeIdForImport(resourceName),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func getDatascienceModelProvenanceCompositeIdForImport(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("models/%s/provenance", rs.Primary.Attributes["model_id"]), nil
	}
}
