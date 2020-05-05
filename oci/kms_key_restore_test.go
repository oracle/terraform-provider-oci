// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	keyRestoreFromFileRepresentation = map[string]interface{}{
		"restore_key_from_file_details": Representation{repType: Optional, create: `${data.oci_objectstorage_object.test_object.content}`},
		"content_length":                Representation{repType: Required, create: `10`},
	}

	keyrestoreFromObjectBackupLocationRepresentation = map[string]interface{}{
		"destination": Representation{repType: Required, create: `BUCKET`},
		"bucket":      Representation{repType: Optional, create: `private-vault-kms-backup`},
		"namespace":   Representation{repType: Optional, create: `${data.oci_objectstorage_object.test_object.namespace}`},
		"object":      Representation{repType: Optional, create: `Key-C-backup`},
		"uri":         Representation{repType: Optional, create: `https://objectstorage.us-phoenix-1.oraclecloud.com/p/5FhcAi6WKBBXLneoMYB51sM1mIgJllbhX5OdgwPGp-Y/n/dxterraformdev/b/private-vault-kms-backup/o/example-key2-backup`},
	}
	keyrestoreFromObjectUriBackupLocationRepresentation = map[string]interface{}{
		"destination": Representation{repType: Required, create: `PRE_AUTHENTICATED_REQUEST_URI`},
		"bucket":      Representation{repType: Optional, create: `private-vault-kms-backup`},
		"namespace":   Representation{repType: Optional, create: `${data.oci_objectstorage_object.test_object.namespace}`},
		"object":      Representation{repType: Optional, create: `test-key-backup`},
		"uri":         Representation{repType: Optional, create: `https://objectstorage.us-phoenix-1.oraclecloud.com/p/R1m0iYINWXwjq1BjxXWypMt0XKj3WCIT4M05hGZ9baw/n/dxterraformdev/b/private-vault-kms-backup/o/test-key-backup`},
	}
	keyRestoreFileContent = `
		data "oci_objectstorage_object" "test_object" {
			#Required
			bucket = "private-vault-kms-backup"
			namespace = "dxterraformdev"
			object = "example-key2-backup"
		}
	`
)

func TestResourceKmsKeyRestore_basic(t *testing.T) {
	//t.Skip("Skip this test till KMS provides a better way of testing this.")
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
			// verify restore from object storage
			{
				Config: config + compartmentIdVariableStr + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
						representationCopyWithNewProperties(keyRepresentation, map[string]interface{}{
							"restore_from_object_store": RepresentationGroup{Required, keyrestoreFromObjectBackupLocationRepresentation}})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				),
			},
			// verify restore from Pre-Authenticated-uri
			{
				Config: config + compartmentIdVariableStr + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
						representationCopyWithNewProperties(keyRepresentation, map[string]interface{}{
							"restore_from_object_store": RepresentationGroup{Required, keyrestoreFromObjectUriBackupLocationRepresentation}})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				),
			},
			// verify restore key from file
			{
				Config: config + compartmentIdVariableStr + DefinedTagsDependencies +
					generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create,
						representationCopyWithNewProperties(keyRepresentation, map[string]interface{}{
							"backup_location_file": RepresentationGroup{Required, keyRestoreFromFileRepresentation}})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_length", "10"),
				),
			},
		},
	})
}
