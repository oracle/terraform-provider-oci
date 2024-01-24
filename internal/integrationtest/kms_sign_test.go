// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	KmsSignRequiredOnlyResource = KmsSignResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_sign", "test_sign", acctest.Required, acctest.Create, KmsSignRepresentation)

	KmsSignRepresentation = map[string]interface{}{
		"crypto_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":            acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency_RSA.keys[0], "id")}`},
		"message":           acctest.Representation{RepType: acctest.Required, Create: `message`},
		"signing_algorithm": acctest.Representation{RepType: acctest.Required, Create: `SHA_224_RSA_PKCS1_V1_5`},
		"message_type":      acctest.Representation{RepType: acctest.Optional, Create: `RAW`},
	}

	KmsSignResourceDependencies = KeyResourceDependencyConfig
)

// issue-routing-tag: kms/default
func TestKmsSignResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsSignResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_sign.test_sign"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+KmsSignResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_sign", "test_sign", acctest.Optional, acctest.Create, KmsSignRepresentation), "keymanagement", "sign", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + KmsSignResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_sign", "test_sign", acctest.Required, acctest.Create, KmsSignRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "message", "message"),
				resource.TestCheckResourceAttr(resourceName, "signing_algorithm", "SHA_224_RSA_PKCS1_V1_5"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + KmsSignResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + KmsSignResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_sign", "test_sign", acctest.Optional, acctest.Create, KmsSignRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
