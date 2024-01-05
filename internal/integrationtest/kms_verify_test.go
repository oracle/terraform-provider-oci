// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	KmsVerifyRequiredOnlyResource = KmsVerifyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_verify", "test_verify", acctest.Required, acctest.Create, KmsVerifyRepresentation)

	KmsVerifyRepresentation = map[string]interface{}{
		"crypto_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":            acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"key_version_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_sign.test_sign.key_version_id}`},
		"message":           acctest.Representation{RepType: acctest.Required, Create: `message`},
		"signature":         acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_sign.test_sign.signature}`},
		"signing_algorithm": acctest.Representation{RepType: acctest.Required, Create: `SHA_224_RSA_PKCS1_V1_5`},
		"message_type":      acctest.Representation{RepType: acctest.Optional, Create: `RAW`},
	}

	verifyKeyVersionRepresentation = map[string]interface{}{
		"key_id":              acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"time_of_deletion":    acctest.Representation{RepType: acctest.Required, Create: keyVersionDeletionTime.Format(time.RFC3339Nano)},
	}

	KmsVerifyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_kms_sign", "test_sign", acctest.Required, acctest.Create, KmsSignRepresentation) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: kms/default
func TestKmsVerifyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsVerifyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_verify.test_verify"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+KmsVerifyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_verify", "test_verify", acctest.Optional, acctest.Create, KmsVerifyRepresentation), "keymanagement", "verify", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + KmsVerifyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_verify", "test_verify", acctest.Required, acctest.Create, KmsVerifyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_version_id"),
				resource.TestCheckResourceAttr(resourceName, "message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", "SHA_224_RSA_PKCS1_V1_5"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + KmsVerifyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + KmsVerifyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_verify", "test_verify", acctest.Optional, acctest.Create, KmsVerifyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "is_signature_valid"),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_version_id"),
				resource.TestCheckResourceAttr(resourceName, "message", "message"),
				resource.TestCheckResourceAttr(resourceName, "message_type", "RAW"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", "SHA_224_RSA_PKCS1_V1_5"),
			),
		},
	})
}
