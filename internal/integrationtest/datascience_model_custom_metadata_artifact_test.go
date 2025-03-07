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
	DatascienceModelCustomMetadataArtifactRequiredOnlyResource = DatascienceModelCustomMetadataArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_custom_metadata_artifact", "test_model_custom_metadata_artifact", acctest.Required, acctest.Create, DatascienceModelCustomMetadataArtifactRepresentation)

	DatascienceModelCustomMetadataArtifactRepresentation = map[string]interface{}{
		"model_custom_metadatum_artifact": acctest.Representation{RepType: acctest.Required, Create: `modelCustomMetadatumArtifact`, Update: `modelCustomMetadatumArtifact2`},
		"content_length":                  acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"metadatum_key_name":              acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key.test_key.name}`},
		"model_id":                        acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
		"content_disposition":             acctest.Representation{RepType: acctest.Optional, Create: `contentDisposition`, Update: `contentDisposition2`},
	}

	DatascienceModelCustomMetadataArtifactResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, DatascienceModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: datascience/default
func TestDatascienceModelCustomMetadataArtifactResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelCustomMetadataArtifactResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_model_custom_metadata_artifact.test_model_custom_metadata_artifact"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceModelCustomMetadataArtifactResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_custom_metadata_artifact", "test_model_custom_metadata_artifact", acctest.Optional, acctest.Create, DatascienceModelCustomMetadataArtifactRepresentation), "datascience", "modelCustomMetadataArtifact", t)

	acctest.ResourceTest(t, testAccCheckDatascienceModelCustomMetadataArtifactDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelCustomMetadataArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_custom_metadata_artifact", "test_model_custom_metadata_artifact", acctest.Required, acctest.Create, DatascienceModelCustomMetadataArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "model_custom_metadatum_artifact", "modelCustomMetadatumArtifact"),
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
			Config: config + compartmentIdVariableStr + DatascienceModelCustomMetadataArtifactResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceModelCustomMetadataArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_custom_metadata_artifact", "test_model_custom_metadata_artifact", acctest.Optional, acctest.Create, DatascienceModelCustomMetadataArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "model_custom_metadatum_artifact", "modelCustomMetadatumArtifact"),
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
			Config: config + compartmentIdVariableStr + DatascienceModelCustomMetadataArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_custom_metadata_artifact", "test_model_custom_metadata_artifact", acctest.Optional, acctest.Update, DatascienceModelCustomMetadataArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "model_custom_metadatum_artifact", "modelCustomMetadatumArtifact2"),
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
			Config:            config + DatascienceModelCustomMetadataArtifactRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"model_custom_metadatum_artifact",
				"content_disposition",
				"content_length",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatascienceModelCustomMetadataArtifactDestroy(s *terraform.State) error {
	noResourceFound := true

	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
