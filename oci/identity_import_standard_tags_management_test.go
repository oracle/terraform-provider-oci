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
	importStandardTagsManagementRepresentation = map[string]interface{}{
		"compartment_id":              Representation{RepType: Required, Create: `${var.compartment_id}`},
		"standard_tag_namespace_name": Representation{RepType: Required, Create: `${oci_identity_tag_namespace.test_tag_namespace.name}`},
	}

	standardTagNamespaceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"description":    Representation{RepType: Required, Create: `Oracle recommended tags`},
		"name":           Representation{RepType: Required, Create: `Oracle-Standard`},
	}

	ImportStandardTagsManagementResourceDependencies = GenerateResourceFromRepresentationMap("oci_identity_tag_namespace", "test_tag_namespace", Required, Create, standardTagNamespaceRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityImportStandardTagsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityImportStandardTagsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_import_standard_tags_management.test_import_standard_tags_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ImportStandardTagsManagementResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_identity_import_standard_tags_management", "test_import_standard_tags_management", Required, Create, importStandardTagsManagementRepresentation), "identity", "importStandardTagsManagement", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ImportStandardTagsManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_import_standard_tags_management", "test_import_standard_tags_management", Required, Create, importStandardTagsManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "standard_tag_namespace_name"),

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
	})
}
