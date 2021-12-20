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
	EncryptedDataResourceConfig = EncryptedDataResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", acctest.Optional, acctest.Create, encryptedDataRepresentation)

	encryptedDataSingularDataSourceRepresentation = map[string]interface{}{
		"crypto_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":          acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"plaintext":       acctest.Representation{RepType: acctest.Required, Create: `aGVsbG8sIHdvcmxk`},
		"associated_data": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"associatedData": "associatedData"}, Update: map[string]string{"associatedData2": "associatedData2"}},
	}

	encryptedDataRepresentation = map[string]interface{}{
		"crypto_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":          acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"plaintext":       acctest.Representation{RepType: acctest.Required, Create: `aGVsbG8sIHdvcmxk`},
		"associated_data": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"associatedData": "associatedData"}, Update: map[string]string{"associatedData2": "associatedData2"}},
		"logging_context": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"loggingContext": "loggingContext"}, Update: map[string]string{"loggingContext2": "loggingContext2"}},
	}

	EncryptedDataResourceDependencies = KeyResourceDependencyConfig
)

// issue-routing-tag: kms/default
func TestKmsEncryptedDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsEncryptedDataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_encrypted_data.test_encrypted_data"
	singularDatasourceName := "data.oci_kms_encrypted_data.test_encrypted_data"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+EncryptedDataResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", acctest.Optional, acctest.Create, encryptedDataRepresentation), "keymanagement", "encryptedData", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + EncryptedDataResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", acctest.Required, acctest.Create, encryptedDataRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "plaintext", "aGVsbG8sIHdvcmxk"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + EncryptedDataResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + EncryptedDataResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", acctest.Optional, acctest.Create, encryptedDataRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "associated_data.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "ciphertext"),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "logging_context.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "plaintext", "aGVsbG8sIHdvcmxk"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", acctest.Optional, acctest.Create, encryptedDataSingularDataSourceRepresentation) +
				compartmentIdVariableStr + EncryptedDataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "associated_data.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plaintext", "aGVsbG8sIHdvcmxk"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "ciphertext"),
			),
		},
	})
}
