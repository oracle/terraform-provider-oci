// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	vaultRestoreConfig = acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
			"restore_from_file": acctest.RepresentationGroup{RepType: acctest.Optional, Group: vaultRestoreFromFileRepresentation}}))

	vaultRestoreRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `private-vault-kms`},
		"vault_type":       acctest.Representation{RepType: acctest.Required, Create: `VIRTUAL_PRIVATE`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"time_of_deletion": acctest.Representation{RepType: acctest.Optional, Create: deletionTime.Format(time.RFC3339Nano)},
	}

	vaultRestoreFromFileRepresentation = map[string]interface{}{
		"restore_vault_from_file_details": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_objectstorage_object.test_object.content}`},
		"content_length":                  acctest.Representation{RepType: acctest.Required, Create: `10`},
		//"content_md5":                     acctest.Representation{RepType: acctest.Optional, Create: `contentMd5`},
	}

	vaultRestoreFromObjectBackupLocationRepresentation = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `BUCKET`},
		"bucket":      acctest.Representation{RepType: acctest.Optional, Create: `${var.bucket_name}`},
		"namespace":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_object.test_objectnnamespace}`},
		"object":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_object.test_object.object}`},
	}

	vaultRestoreFromObjecUriBackupLocationRepresentation = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `PRE_AUTHENTICATED_REQUEST_URI`},
		"uri":         acctest.Representation{RepType: acctest.Optional, Create: `${var.backup_uri}`},
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

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_vault.private-vault-kms"

	acctest.ResourceTest(t, testAccCheckKMSVaultDestroy, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
						"restore_from_file": acctest.RepresentationGroup{RepType: acctest.Optional, Group: vaultRestoreFromFileRepresentation}})),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content_length", "10"),
			),
		},
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
						"restore_from_object_store": acctest.RepresentationGroup{RepType: acctest.Optional, Group: vaultRestoreFromObjectBackupLocationRepresentation}})),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "private-vault-kms", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(vaultRestoreRepresentation, map[string]interface{}{
						"restore_from_object_store": acctest.RepresentationGroup{RepType: acctest.Optional, Group: vaultRestoreFromObjecUriBackupLocationRepresentation}})),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			),
		},
	})
}
