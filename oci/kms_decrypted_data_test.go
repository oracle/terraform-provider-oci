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
	decryptedDataSingularDataSourceRepresentation = map[string]interface{}{
		"ciphertext":      Representation{RepType: Required, Create: `${oci_kms_encrypted_data.test_encrypted_data.ciphertext}`},
		"crypto_endpoint": Representation{RepType: Required, Create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":          Representation{RepType: Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"associated_data": Representation{RepType: Optional, Create: map[string]string{"associatedData": "associatedData"}, Update: map[string]string{"associatedData2": "associatedData2"}},
	}

	DecryptedDataResourceConfig = GenerateResourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", Required, Create, encryptedDataRepresentation) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: kms/default
func TestKmsDecryptedDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsDecryptedDataResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_kms_decrypted_data.test_decrypted_data"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_kms_decrypted_data", "test_decrypted_data", Required, Create, decryptedDataSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DecryptedDataResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ciphertext"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "plaintext", "aGVsbG8sIHdvcmxk"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plaintext_checksum"),
			),
		},
	})
}
