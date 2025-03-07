// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatascienceModelDefinedMetadataArtifactRequiredOnlyResource = DatascienceModelDefinedMetadataArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_defined_metadata_artifact", "test_model_defined_metadata_artifact", acctest.Required, acctest.Create, DatascienceModelDefinedMetadataArtifactRepresentation)

	DatascienceModelDefinedMetadataArtifactRepresentation = map[string]interface{}{
		"model_defined_metadatum_artifact": acctest.Representation{RepType: acctest.Required, Create: `modelDefinedMetadatumArtifact`, Update: `modelDefinedMetadatumArtifact2`},
		"content_length":                   acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"metadatum_key_name":               acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key.test_key.name}`},
		"model_id":                         acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
		"content_disposition":              acctest.Representation{RepType: acctest.Optional, Create: `contentDisposition`, Update: `contentDisposition2`},
	}

	DatascienceModelDefinedMetadataArtifactResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, DatascienceModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: datascience/default
func TestDatascienceModelDefinedMetadataArtifactResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDefinedMetadataArtifactResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_model_defined_metadata_artifact.test_model_defined_metadata_artifact"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceModelDefinedMetadataArtifactResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_defined_metadata_artifact", "test_model_defined_metadata_artifact", acctest.Optional, acctest.Create, DatascienceModelDefinedMetadataArtifactRepresentation), "datascience", "modelDefinedMetadataArtifact", t)

	acctest.ResourceTest(t, testAccCheckDatascienceModelDefinedMetadataArtifactDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDefinedMetadataArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_defined_metadata_artifact", "test_model_defined_metadata_artifact", acctest.Required, acctest.Create, DatascienceModelDefinedMetadataArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "model_defined_metadatum_artifact", "modelDefinedMetadatumArtifact"),
				resource.TestCheckResourceAttr(resourceName, "content_length", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "metadatum_key_name"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDefinedMetadataArtifactResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDefinedMetadataArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_defined_metadata_artifact", "test_model_defined_metadata_artifact", acctest.Optional, acctest.Create, DatascienceModelDefinedMetadataArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "model_defined_metadatum_artifact", "modelDefinedMetadatumArtifact"),
				resource.TestCheckResourceAttr(resourceName, "content_disposition", "contentDisposition"),
				resource.TestCheckResourceAttr(resourceName, "content_length", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "metadatum_key_name"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),

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
			Config: config + compartmentIdVariableStr + DatascienceModelDefinedMetadataArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_defined_metadata_artifact", "test_model_defined_metadata_artifact", acctest.Optional, acctest.Update, DatascienceModelDefinedMetadataArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "model_defined_metadatum_artifact", "modelDefinedMetadatumArtifact2"),
				resource.TestCheckResourceAttr(resourceName, "content_disposition", "contentDisposition2"),
				resource.TestCheckResourceAttr(resourceName, "content_length", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "metadatum_key_name"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify resource import
		{
			Config:            config + DatascienceModelDefinedMetadataArtifactRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"model_defined_metadatum_artifact",
				"content_disposition",
				"content_length",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatascienceModelDefinedMetadataArtifactDestroy(s *terraform.State) error {
	noResourceFound := true

	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
