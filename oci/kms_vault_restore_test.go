// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	vaultRestoreConfig = GenerateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", Optional, Create,
		RepresentationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
			"restore_from_file": RepresentationGroup{Optional, vaultRestoreFromFileRepresentation}}))

	vaultRestoreRepresentation = map[string]interface{}{
		"compartment_id":   Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":     Representation{RepType: Required, Create: `private-vault-kms`},
		"vault_type":       Representation{RepType: Required, Create: `VIRTUAL_PRIVATE`},
		"defined_tags":     Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"time_of_deletion": Representation{RepType: Optional, Create: deletionTime.Format(time.RFC3339Nano)},
	}

	vaultRestoreFromFileRepresentation = map[string]interface{}{
		"restore_vault_from_file_details": Representation{RepType: Optional, Create: `${data.oci_objectstorage_object.test_object.content}`},
		"content_length":                  Representation{RepType: Required, Create: `10`},
		//"content_md5":                     Representation{RepType: Optional, Create: `contentMd5`},
	}

	vaultRestoreFromObjectBackupLocationRepresentation = map[string]interface{}{
		"destination": Representation{RepType: Required, Create: `BUCKET`},
		"bucket":      Representation{RepType: Optional, Create: `${var.bucket_name}`},
		"namespace":   Representation{RepType: Optional, Create: `${oci_objectstorage_object.test_objectnnamespace}`},
		"object":      Representation{RepType: Optional, Create: `${oci_objectstorage_object.test_object.object}`},
	}

	vaultRestoreFromObjecUriBackupLocationRepresentation = map[string]interface{}{
		"destination": Representation{RepType: Required, Create: `PRE_AUTHENTICATED_REQUEST_URI`},
		"uri":         Representation{RepType: Optional, Create: `${var.backup_uri}`},
	}
	vaultRestoreFileContent = `
		data "oci_objectstorage_object" "test_object" {
			#Required
			bucket = "bucket-name"
			namespace = "namespace"
			object = "object"
		}
	`
)

// issue-routing-tag: kms/default
func TestResourceKmsVaultRestore_default(t *testing.T) {
	t.Skip("Skip this test till KMS provides a better way of testing this.")
	httpreplay.SetScenario("TestResourceKmsVaultRestore_virtual")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_vault.private-vault-kms"

	ResourceTest(t, testAccCheckKMSVaultDestroy, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", Optional, Create,
					RepresentationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
						"restore_from_file": RepresentationGroup{Optional, vaultRestoreFromFileRepresentation}})),

			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content_length", "10"),
			),
		},
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", Optional, Create,
					RepresentationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
						"restore_from_object_store": RepresentationGroup{Optional, vaultRestoreFromObjectBackupLocationRepresentation}})),

			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", Optional, Create,
					RepresentationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
						"restore_from_object_store": RepresentationGroup{Optional, vaultRestoreFromObjecUriBackupLocationRepresentation}})),

			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},
	})
}
