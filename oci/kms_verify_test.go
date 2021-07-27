// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VerifyRequiredOnlyResource = VerifyResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_verify", "test_verify", Required, Create, verifyRepresentation)

	verifyRepresentation = map[string]interface{}{
		"crypto_endpoint":   Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":            Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"key_version_id":    Representation{repType: Required, create: `${oci_kms_sign.test_sign.key_version_id}`},
		"message":           Representation{repType: Required, create: `message`},
		"signature":         Representation{repType: Required, create: `${oci_kms_sign.test_sign.signature}`},
		"signing_algorithm": Representation{repType: Required, create: `SHA_224_RSA_PKCS1_V1_5`},
		"message_type":      Representation{repType: Optional, create: `RAW`},
	}

	verifyKeyVersionRepresentation = map[string]interface{}{
		"key_id":              Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"time_of_deletion":    Representation{repType: Required, create: keyVersionDeletionTime.Format(time.RFC3339Nano)},
	}

	VerifyResourceDependencies = generateResourceFromRepresentationMap("oci_kms_sign", "test_sign", Required, Create, signRepresentation) +
		KeyResourceDependencyConfig
)

// issue-routing-tag: kms/default
func TestKmsVerifyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsVerifyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_verify.test_verify"

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+VerifyResourceDependencies+
		generateResourceFromRepresentationMap("oci_kms_verify", "test_verify", Optional, Create, verifyRepresentation), "keymanagement", "verify", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VerifyResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_verify", "test_verify", Required, Create, verifyRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "key_version_id"),
					resource.TestCheckResourceAttr(resourceName, "message", "message"),
					resource.TestCheckResourceAttrSet(resourceName, "signature"),
					resource.TestCheckResourceAttr(resourceName, "signing_algorithm", "SHA_224_RSA_PKCS1_V1_5"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VerifyResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VerifyResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_verify", "test_verify", Optional, Create, verifyRepresentation),
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
		},
	})
}
