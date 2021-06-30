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
	keyRestorekeyRepresentation = map[string]interface{}{
		"restore_trigger":     Representation{repType: Required, create: `0`, update: `1`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":        Representation{repType: Required, create: `Key C`, update: `displayName2`},
		"key_shape":           RepresentationGroup{Required, keyRestoreKeyShapeRepresentation},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value-updated")}`},
	}
	keyRestorekeyRepresentationUpdate2 = map[string]interface{}{
		"restore_trigger":     Representation{repType: Required, create: `1`, update: `1`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":        Representation{repType: Required, create: `Key C`, update: `displayName2`},
		"key_shape":           RepresentationGroup{Required, keyRestoreKeyShapeRepresentation},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value-updated")}`},
	}
	keyRestoreKeyShapeRepresentation = map[string]interface{}{
		"algorithm": Representation{repType: Required, create: `AES`},
		"length":    Representation{repType: Required, create: `16`},
	}
	keyrestoreFromObjectBackupLocationRepresentation = map[string]interface{}{
		"destination": Representation{repType: Required, create: `BUCKET`},
		"bucket":      Representation{repType: Optional, create: `private-vault-kms-backup`},
		"namespace":   Representation{repType: Optional, create: `${data.oci_objectstorage_object.test_object.namespace}`},
		"object":      Representation{repType: Optional, create: `Key-C-backup`},
	}
	keyrestoreFromObjectUriBackupLocationRepresentation = map[string]interface{}{
		"destination": Representation{repType: Required, create: `PRE_AUTHENTICATED_REQUEST_URI`},
		"uri":         Representation{repType: Optional, create: `PAR-uri`},
	}
	keyRestoreFromFileRepresentation = map[string]interface{}{
		"restore_key_from_file_details": Representation{repType: Required, create: `${data.oci_objectstorage_object.test_object.content}`},
		"content_length":                Representation{repType: Required, create: `${data.oci_objectstorage_object.test_object.content_length}`},
	}
	keyRestoreFileContentObject = `
		data "oci_objectstorage_object" "test_object" {
			#Required
			bucket = "bucket"
			namespace = "object-namespace"
			object = "KeyC-backup"
			base64_encode_content = "true"
		}`
	vaultData = `
	data "oci_kms_vault" "test_vault" {
		vault_id = "${oci_kms_vault.test_vault.id}"
	}`
)

func TestResourceKmsKeyRestore_basic(t *testing.T) {
	t.Skip("Skip this test till KMS provides a better way of testing this.")

	httpreplay.SetScenario("TestResourceKmsKeyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_key.test_key"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify restore key from file
			{
				Config: config + compartmentIdVariableStr + DefinedTagsDependencies + keyRestoreFileContentObject +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
						representationCopyWithNewProperties(keyRestorekeyRepresentation, map[string]interface{}{
							"restore_from_file": RepresentationGroup{Required, keyRestoreFromFileRepresentation}})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				),
			},
			// verify restore from object store
			{
				Config: config + compartmentIdVariableStr + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
						representationCopyWithNewProperties(keyRestorekeyRepresentationUpdate2, map[string]interface{}{
							"restore_from_object_store": RepresentationGroup{Required, keyrestoreFromObjectBackupLocationRepresentation}})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				),
			},
			// verify restore from Pre-Authenticated-uri
			{
				Config: config + compartmentIdVariableStr + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
						representationCopyWithNewProperties(keyRepresentation, map[string]interface{}{
							"restore_from_object_store": RepresentationGroup{Required, keyrestoreFromObjectUriBackupLocationRepresentation}})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				),
			},
		},
	})
}
