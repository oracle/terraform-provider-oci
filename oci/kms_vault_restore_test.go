// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	vaultRestoreConfig = generateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", Optional, Create,
		representationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
			"restore_from_file": RepresentationGroup{Optional, vaultRestoreFromFileRepresentation}}))

	vaultRestoreRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":     Representation{repType: Required, create: `private-vault-kms`},
		"vault_type":       Representation{repType: Required, create: `VIRTUAL_PRIVATE`},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"time_of_deletion": Representation{repType: Optional, create: deletionTime.Format(time.RFC3339Nano)},
	}

	vaultRestoreFromFileRepresentation = map[string]interface{}{
		"restore_vault_from_file_details": Representation{repType: Optional, create: `${data.oci_objectstorage_object.test_object.content}`},
		"content_length":                  Representation{repType: Required, create: `10`},
		//"content_md5":                     Representation{repType: Optional, create: `contentMd5`},
	}

	vaultRestoreFromObjectBackupLocationRepresentation = map[string]interface{}{
		"destination": Representation{repType: Required, create: `BUCKET`},
		"bucket":      Representation{repType: Optional, create: `${var.bucket_name}`},
		"namespace":   Representation{repType: Optional, create: `${oci_objectstorage_object.test_objectnnamespace}`},
		"object":      Representation{repType: Optional, create: `${oci_objectstorage_object.test_object.object}`},
	}

	vaultRestoreFromObjecUriBackupLocationRepresentation = map[string]interface{}{
		"destination": Representation{repType: Required, create: `PRE_AUTHENTICATED_REQUEST_URI`},
		"uri":         Representation{repType: Optional, create: `${var.backup_uri}`},
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

func TestResourceKmsVaultRestore_default(t *testing.T) {
	t.Skip("Skip this test till KMS provides a better way of testing this.")
	httpreplay.SetScenario("TestResourceKmsVaultRestore_virtual")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_vault.private-vault-kms"
	var c = config + compartmentIdVariableStr + VaultResourceDependencies + vaultRestoreConfig
	fmt.Println(c)
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckKMSVaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: config + compartmentIdVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", Optional, Create,
						representationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
							"restore_from_file": RepresentationGroup{Optional, vaultRestoreFromFileRepresentation}})),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_length", "10"),
				),
			},
			{
				Config: config + compartmentIdVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", Optional, Create,
						representationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
							"restore_from_object_store": RepresentationGroup{Optional, vaultRestoreFromObjectBackupLocationRepresentation}})),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				),
			},
			{
				Config: config + compartmentIdVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", Optional, Create,
						representationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
							"restore_from_object_store": RepresentationGroup{Optional, vaultRestoreFromObjecUriBackupLocationRepresentation}})),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				),
			},
		},
	})
}
