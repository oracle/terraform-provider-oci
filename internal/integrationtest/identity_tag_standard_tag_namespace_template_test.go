// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	tagStandardTagNamespaceTemplateSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"standard_tag_namespace_name": acctest.Representation{RepType: acctest.Required, Create: `Oracle-Standard`},
	}

	tagStandardTagNamespaceTemplateDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	TagStandardTagNamespaceTemplateResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityTagStandardTagNamespaceTemplateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTagStandardTagNamespaceTemplateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.LegacyTestProviderConfig()

	datasourceName := "data.oci_identity_tag_standard_tag_namespace_templates.test_tag_standard_tag_namespace_template"
	singularDatasourceName := "data.oci_identity_tag_standard_tag_namespace_template.test_tag_standard_tag_namespace_template"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_tag_standard_tag_namespace_templates", "test_tag_standard_tag_namespace_template", acctest.Required, acctest.Create, tagStandardTagNamespaceTemplateDataSourceRepresentation) +
				TagStandardTagNamespaceTemplateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "standard_tag_namespace_templates.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "standard_tag_namespace_templates.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "standard_tag_namespace_templates.0.standard_tag_namespace_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "standard_tag_namespace_templates.0.status"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_tag_standard_tag_namespace_template", "test_tag_standard_tag_namespace_template", acctest.Required, acctest.Create, tagStandardTagNamespaceTemplateSingularDataSourceRepresentation) +
				TagStandardTagNamespaceTemplateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "standard_tag_namespace_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tag_definition_templates.#", "8"),
			),
		},
	})
}
