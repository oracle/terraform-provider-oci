// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatascienceModelCustomMetadataArtifactContentSingularDataSourceRepresentation = map[string]interface{}{
		"metadatum_key_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key.test_key.name}`},
		"model_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
		"range":              acctest.Representation{RepType: acctest.Optional, Create: `range`},
	}

	DatascienceModelCustomMetadataArtifactContentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, DatascienceModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: datascience/default
func TestDatascienceModelCustomMetadataArtifactContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelCustomMetadataArtifactContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_datascience_model_custom_metadata_artifact_content.test_model_custom_metadata_artifact_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_custom_metadata_artifact_content", "test_model_custom_metadata_artifact_content", acctest.Required, acctest.Create, DatascienceModelCustomMetadataArtifactContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelCustomMetadataArtifactContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadatum_key_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "range", "range"),
			),
		},
	})
}
