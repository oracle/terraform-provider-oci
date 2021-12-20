// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	GeneratedKeyRequiredOnlyResource = GeneratedKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_generated_key", "test_generated_key", acctest.Required, acctest.Create, generatedKeyRepresentation)

	generatedKeyRepresentation = map[string]interface{}{
		"crypto_endpoint":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"include_plaintext_key": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"key_id":                acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"key_shape":             acctest.RepresentationGroup{RepType: acctest.Required, Group: generatedKeyKeyShapeRepresentation},
		"associated_data":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"associatedData": "associatedData"}, Update: map[string]string{"associatedData2": "associatedData2"}},
		"logging_context":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"loggingContext": "loggingContext"}, Update: map[string]string{"loggingContext2": "loggingContext2"}},
	}
	generatedKeyKeyShapeRepresentation = map[string]interface{}{
		"algorithm": acctest.Representation{RepType: acctest.Required, Create: `AES`},
		"length":    acctest.Representation{RepType: acctest.Required, Create: `16`},
	}

	GeneratedKeyResourceDependencies = KeyResourceDependencyConfig
)

// issue-routing-tag: kms/default
func TestKmsGeneratedKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsGeneratedKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_generated_key.test_generated_key"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GeneratedKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_generated_key", "test_generated_key", acctest.Optional, acctest.Create, generatedKeyRepresentation), "keymanagement", "generatedKey", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GeneratedKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_generated_key", "test_generated_key", acctest.Required, acctest.Create, generatedKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "include_plaintext_key", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GeneratedKeyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GeneratedKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_generated_key", "test_generated_key", acctest.Optional, acctest.Create, generatedKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "associated_data.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "ciphertext"),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "include_plaintext_key", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "key_id"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.algorithm", "AES"),
				resource.TestCheckResourceAttr(resourceName, "key_shape.0.length", "16"),
				resource.TestCheckResourceAttr(resourceName, "logging_context.%", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
