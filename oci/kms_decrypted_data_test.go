// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	decryptedDataSingularDataSourceRepresentation = map[string]interface{}{
		"ciphertext":      Representation{repType: Required, create: `${oci_kms_encrypted_data.test_encrypted_data.ciphertext}`},
		"crypto_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":          Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"associated_data": Representation{repType: Optional, create: map[string]string{"associatedData": "associatedData"}, update: map[string]string{"associatedData2": "associatedData2"}},
	}

	DecryptedDataResourceConfig = EncryptedDataRequiredOnlyResource
)

func TestKmsDecryptedDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsDecryptedDataResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_kms_decrypted_data.test_decrypted_data"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_decrypted_data", "test_decrypted_data", Required, Create, decryptedDataSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DecryptedDataResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ciphertext"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "plaintext", "aGVsbG8sIHdvcmxk"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "plaintext_checksum"),
				),
			},
		},
	})
}
