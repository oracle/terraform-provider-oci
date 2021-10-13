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
		GenerateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Required, Create, modelProvenanceRepresentation)

	ModelProvenanceResourceConfig = ModelProvenanceResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Optional, Update, modelProvenanceRepresentation)

	modelProvenanceSingularDataSourceRepresentation = map[string]interface{}{
		"model_id": Representation{RepType: Required, Create: `${oci_datascience_model.test_model.id}`},
	}

	modelProvenanceRepresentation = map[string]interface{}{
		"model_id":        Representation{RepType: Required, Create: `${oci_datascience_model.test_model.id}`},
		"git_branch":      Representation{RepType: Optional, Create: `gitBranch`, Update: `gitBranch2`},
		"git_commit":      Representation{RepType: Optional, Create: `gitCommit`, Update: `gitCommit2`},
		"repository_url":  Representation{RepType: Optional, Create: `repositoryUrl`, Update: `repositoryUrl2`},
		"script_dir":      Representation{RepType: Optional, Create: `scriptDir`, Update: `scriptDir2`},
		"training_id":     Representation{RepType: Optional, Create: `ocid1.datasciencenotebooksession.oc1.iad.amaaaaaav66vvniaiasz7njfotab2z7i6yvnk4tnbaamqqc7g45tvme3doyq`},
		"training_script": Representation{RepType: Optional, Create: `trainingScript`, Update: `trainingScript2`},
	}

	ModelProvenanceResourceDependencies = GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", Required, Create, modelRepresentation) +
		GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", Required, Create, projectRepresentation)
)

// issue-routing-tag: datascience/default
func TestDatascienceModelProvenanceResource_basic(t *testing.T) {
	t.Skip("Skip this test till data science team provides an execution for the test within few hours")
	httpreplay.SetScenario("TestDatascienceModelProvenanceResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_model_provenance.test_model_provenance"

	singularDatasourceName := "data.oci_datascience_model_provenance.test_model_provenance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ModelProvenanceResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Optional, Create, modelProvenanceRepresentation), "datascience", "modelProvenance", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ModelProvenanceResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Required, Create, modelProvenanceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		{
			Config: config + compartmentIdVariableStr + ModelProvenanceResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Optional, Create, modelProvenanceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "git_branch", "gitBranch"),
				resource.TestCheckResourceAttr(resourceName, "git_commit", "gitCommit"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttr(resourceName, "repository_url", "repositoryUrl"),
				resource.TestCheckResourceAttr(resourceName, "script_dir", "scriptDir"),
				resource.TestCheckResourceAttrSet(resourceName, "training_id"),
				resource.TestCheckResourceAttr(resourceName, "training_script", "trainingScript"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Optional, Update, modelProvenanceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "git_branch", "gitBranch2"),
				resource.TestCheckResourceAttr(resourceName, "git_commit", "gitCommit2"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttr(resourceName, "repository_url", "repositoryUrl2"),
				resource.TestCheckResourceAttr(resourceName, "script_dir", "scriptDir2"),
				resource.TestCheckResourceAttrSet(resourceName, "training_id"),
				resource.TestCheckResourceAttr(resourceName, "training_script", "trainingScript2"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_datascience_model_provenance", "test_model_provenance", Required, Create, modelProvenanceSingularDataSourceRepresentation) +
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
