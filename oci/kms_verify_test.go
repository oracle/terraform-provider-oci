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
	VerifyRequiredOnlyResource = VerifyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_kms_verify", "test_verify", Required, Create, verifyRepresentation)

	verifyRepresentation = map[string]interface{}{
		"crypto_endpoint":   Representation{RepType: Required, Create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":            Representation{RepType: Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"key_version_id":    Representation{RepType: Required, Create: `${oci_kms_sign.test_sign.key_version_id}`},
		"message":           Representation{RepType: Required, Create: `message`},
		"signature":         Representation{RepType: Required, Create: `${oci_kms_sign.test_sign.signature}`},
		"signing_algorithm": Representation{RepType: Required, Create: `SHA_224_RSA_PKCS1_V1_5`},
		"message_type":      Representation{RepType: Optional, Create: `RAW`},
	}

	verifyKeyVersionRepresentation = map[string]interface{}{
		"key_id":              Representation{RepType: Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"management_endpoint": Representation{RepType: Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"time_of_deletion":    Representation{RepType: Required, Create: keyVersionDeletionTime.Format(time.RFC3339Nano)},
	}

	VerifyResourceDependencies = GenerateResourceFromRepresentationMap("oci_kms_sign", "test_sign", Required, Create, signRepresentation) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: kms/default
func TestKmsVerifyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsVerifyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_verify.test_verify"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+VerifyResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_kms_verify", "test_verify", Optional, Create, verifyRepresentation), "keymanagement", "verify", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VerifyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_verify", "test_verify", Required, Create, verifyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + VerifyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VerifyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_verify", "test_verify", Optional, Create, verifyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
