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
	decryptedDataSingularDataSourceRepresentation = map[string]interface{}{
		"ciphertext":      acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_encrypted_data.test_encrypted_data.ciphertext}`},
		"crypto_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":          acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"associated_data": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"associatedData": "associatedData"}, Update: map[string]string{"associatedData2": "associatedData2"}},
	}

	DecryptedDataResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", acctest.Required, acctest.Create, encryptedDataRepresentation) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: kms/default
func TestKmsDecryptedDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsDecryptedDataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_kms_decrypted_data.test_decrypted_data"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_decrypted_data", "test_decrypted_data", acctest.Required, acctest.Create, decryptedDataSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DecryptedDataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ciphertext"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "plaintext", "aGVsbG8sIHdvcmxk"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plaintext_checksum"),
			),
		},
	})
}
