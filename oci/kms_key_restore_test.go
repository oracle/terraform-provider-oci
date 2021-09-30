// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	keyRestorekeyRepresentation = map[string]interface{}{
		"restore_trigger":     Representation{RepType: Required, Create: `0`, Update: `1`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":        Representation{RepType: Required, Create: `Key C`, Update: `displayName2`},
		"key_shape":           RepresentationGroup{Required, keyRestoreKeyShapeRepresentation},
		"management_endpoint": Representation{RepType: Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"defined_tags":        Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value-updated")}`},
	}
	keyRestorekeyRepresentationUpdate2 = map[string]interface{}{
		"restore_trigger":     Representation{RepType: Required, Create: `1`, Update: `1`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":        Representation{RepType: Required, Create: `Key C`, Update: `displayName2`},
		"key_shape":           RepresentationGroup{Required, keyRestoreKeyShapeRepresentation},
		"management_endpoint": Representation{RepType: Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"defined_tags":        Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value-updated")}`},
	}
	keyRestoreKeyShapeRepresentation = map[string]interface{}{
		"algorithm": Representation{RepType: Required, Create: `AES`},
		"length":    Representation{RepType: Required, Create: `16`},
	}
	keyrestoreFromObjectBackupLocationRepresentation = map[string]interface{}{
		"destination": Representation{RepType: Required, Create: `BUCKET`},
		"bucket":      Representation{RepType: Optional, Create: `private-vault-kms-backup`},
		"namespace":   Representation{RepType: Optional, Create: `${data.oci_objectstorage_object.test_object.namespace}`},
		"object":      Representation{RepType: Optional, Create: `Key-C-backup`},
	}
	keyrestoreFromObjectUriBackupLocationRepresentation = map[string]interface{}{
		"destination": Representation{RepType: Required, Create: `PRE_AUTHENTICATED_REQUEST_URI`},
		"uri":         Representation{RepType: Optional, Create: `PAR-uri`},
	}
	keyRestoreFromFileRepresentation = map[string]interface{}{
		"restore_key_from_file_details": Representation{RepType: Required, Create: `${data.oci_objectstorage_object.test_object.content}`},
		"content_length":                Representation{RepType: Required, Create: `${data.oci_objectstorage_object.test_object.content_length}`},
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

// issue-routing-tag: kms/default
func TestResourceKmsKeyRestore_basic(t *testing.T) {
	t.Skip("Skip this test till KMS provides a better way of testing this.")

	httpreplay.SetScenario("TestResourceKmsKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_key.test_key"

	ResourceTest(t, nil, []resource.TestStep{
		// verify restore key from file
		{
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies + keyRestoreFileContentObject +
				GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
					RepresentationCopyWithNewProperties(keyRestorekeyRepresentation, map[string]interface{}{
						"restore_from_file": RepresentationGroup{Required, keyRestoreFromFileRepresentation}})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},
		// verify restore from object store
		{
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
					RepresentationCopyWithNewProperties(keyRestorekeyRepresentationUpdate2, map[string]interface{}{
						"restore_from_object_store": RepresentationGroup{Required, keyrestoreFromObjectBackupLocationRepresentation}})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},
		// verify restore from Pre-Authenticated-uri
		{
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
					RepresentationCopyWithNewProperties(keyRepresentation, map[string]interface{}{
						"restore_from_object_store": RepresentationGroup{Required, keyrestoreFromObjectUriBackupLocationRepresentation}})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},
	})
}
