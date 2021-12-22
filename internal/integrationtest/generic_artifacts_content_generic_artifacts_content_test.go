// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	genericArtifactsContentSingularDataSourceRepresentation = map[string]interface{}{
		"artifact_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generic_artifacts_content_artifact.test_artifact.id}`},
	}

	GenericArtifactsContentResourceConfig = ""
)

// issue-routing-tag: generic_artifacts_content/default
func TestGenericArtifactsContentGenericArtifactsContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenericArtifactsContentGenericArtifactsContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_generic_artifacts_content_generic_artifacts_content.test_generic_artifacts_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generic_artifacts_content_generic_artifacts_content", "test_generic_artifacts_content", acctest.Required, acctest.Create, genericArtifactsContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenericArtifactsContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "artifact_id"),
			),
		},
	})
}
