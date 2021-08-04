// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	genericArtifactsContentSingularDataSourceRepresentation = map[string]interface{}{
		"artifact_id": Representation{repType: Required, create: `${oci_generic_artifacts_content_artifact.test_artifact.id}`},
	}

	GenericArtifactsContentResourceConfig = ""
)

// issue-routing-tag: generic_artifacts_content/default
func TestGenericArtifactsContentGenericArtifactsContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenericArtifactsContentGenericArtifactsContentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_generic_artifacts_content_generic_artifacts_content.test_generic_artifacts_content"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_generic_artifacts_content_generic_artifacts_content", "test_generic_artifacts_content", Required, Create, genericArtifactsContentSingularDataSourceRepresentation) +
					compartmentIdVariableStr + GenericArtifactsContentResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "artifact_id"),
				),
			},
		},
	})
}
