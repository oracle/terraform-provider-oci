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
	keyRestorekeyRepresentation = map[string]interface{}{
		"restore_trigger":     acctest.Representation{RepType: acctest.Required, Create: `0`, Update: `1`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `Key C`, Update: `displayName2`},
		"key_shape":           acctest.RepresentationGroup{RepType: acctest.Required, Group: keyRestoreKeyShapeRepresentation},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value-updated")}`},
	}
	keyRestorekeyRepresentationUpdate2 = map[string]interface{}{
		"restore_trigger":     acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `Key C`, Update: `displayName2`},
		"key_shape":           acctest.RepresentationGroup{RepType: acctest.Required, Group: keyRestoreKeyShapeRepresentation},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value-updated")}`},
	}
	keyRestoreKeyShapeRepresentation = map[string]interface{}{
		"algorithm": acctest.Representation{RepType: acctest.Required, Create: `AES`},
		"length":    acctest.Representation{RepType: acctest.Required, Create: `16`},
	}
	keyrestoreFromObjectBackupLocationRepresentation = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `BUCKET`},
		"bucket":      acctest.Representation{RepType: acctest.Optional, Create: `private-vault-kms-backup`},
		"namespace":   acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_objectstorage_object.test_object.namespace}`},
		"object":      acctest.Representation{RepType: acctest.Optional, Create: `Key-C-backup`},
	}
	keyrestoreFromObjectUriBackupLocationRepresentation = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `PRE_AUTHENTICATED_REQUEST_URI`},
		"uri":         acctest.Representation{RepType: acctest.Optional, Create: `PAR-uri`},
	}
	keyRestoreFromFileRepresentation = map[string]interface{}{
		"restore_key_from_file_details": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_object.test_object.content}`},
		"content_length":                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_object.test_object.content_length}`},
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

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_key.test_key"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify restore key from file
		{
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies + keyRestoreFileContentObject +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(keyRestorekeyRepresentation, map[string]interface{}{
						"restore_from_file": acctest.RepresentationGroup{RepType: acctest.Required, Group: keyRestoreFromFileRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},
		// verify restore from object store
		{
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(keyRestorekeyRepresentationUpdate2, map[string]interface{}{
						"restore_from_object_store": acctest.RepresentationGroup{RepType: acctest.Required, Group: keyrestoreFromObjectBackupLocationRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},
		// verify restore from Pre-Authenticated-uri
		{
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(keyRepresentation, map[string]interface{}{
						"restore_from_object_store": acctest.RepresentationGroup{RepType: acctest.Required, Group: keyrestoreFromObjectUriBackupLocationRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},
	})
}
