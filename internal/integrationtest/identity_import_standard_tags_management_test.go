// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	IdentityImportStandardTagsManagementRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"standard_tag_namespace_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_tag_namespace.test_tag_namespace.name}`},
	}

	IdentityStandardTagNamespaceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `Oracle recommended tags`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `Oracle-Standard`},
	}

	IdentityImportStandardTagsManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", acctest.Required, acctest.Create, IdentityStandardTagNamespaceRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityImportStandardTagsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityImportStandardTagsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.LegacyTestProviderConfig()

	resourceName := "oci_identity_import_standard_tags_management.test_import_standard_tags_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+IdentityImportStandardTagsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_import_standard_tags_management", "test_import_standard_tags_management", acctest.Required, acctest.Create, IdentityImportStandardTagsManagementRepresentation), "identity", "importStandardTagsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + IdentityImportStandardTagsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_import_standard_tags_management", "test_import_standard_tags_management", acctest.Required, acctest.Create, IdentityImportStandardTagsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "standard_tag_namespace_name"),

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
	})
}
