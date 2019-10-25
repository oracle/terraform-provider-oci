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
	GeneratedKeyRequiredOnlyResource = GeneratedKeyResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_generated_key", "test_generated_key", Required, Create, generatedKeyRepresentation)

	generatedKeyRepresentation = map[string]interface{}{
		"crypto_endpoint":       Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"include_plaintext_key": Representation{repType: Required, create: `false`},
		"key_id":                Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"key_shape":             RepresentationGroup{Required, generatedKeyKeyShapeRepresentation},
		"associated_data":       Representation{repType: Optional, create: map[string]string{"associatedData": "associatedData"}, update: map[string]string{"associatedData2": "associatedData2"}},
		"logging_context":       Representation{repType: Optional, create: map[string]string{"loggingContext": "loggingContext"}, update: map[string]string{"loggingContext2": "loggingContext2"}},
	}
	generatedKeyKeyShapeRepresentation = map[string]interface{}{
		"algorithm": Representation{repType: Required, create: `AES`},
		"length":    Representation{repType: Required, create: `16`},
	}

	GeneratedKeyResourceDependencies = KeyResourceDependencyConfig
)

func TestKmsGeneratedKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsGeneratedKeyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_generated_key.test_generated_key"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + GeneratedKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_generated_key", "test_generated_key", Required, Create, generatedKeyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "include_plaintext_key", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + GeneratedKeyResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + GeneratedKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_generated_key", "test_generated_key", Optional, Create, generatedKeyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "associated_data.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "ciphertext"),
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "include_plaintext_key", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
					resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
					resource.TestCheckResourceAttr(resourceName, "logging_context.%", "1"),
				),
			},
		},
	})
}
