// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	EncryptedDataRequiredOnlyResource = EncryptedDataResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", Required, Create, encryptedDataRepresentation)

	EncryptedDataResourceConfig = EncryptedDataResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", Required, Create, encryptedDataRepresentation)

	encryptedDataSingularDataSourceRepresentation = map[string]interface{}{
		"crypto_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":          Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"plaintext":       Representation{repType: Required, create: `aGVsbG8sIHdvcmxk`},
		"associated_data": Representation{repType: Optional, create: map[string]string{"associatedData": "associatedData"}, update: map[string]string{"associatedData2": "associatedData2"}},
	}

	encryptedDataRepresentation = map[string]interface{}{
		"crypto_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.crypto_endpoint}`},
		"key_id":          Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"plaintext":       Representation{repType: Required, create: `aGVsbG8sIHdvcmxk`},
		"associated_data": Representation{repType: Optional, create: map[string]string{"associatedData": "associatedData"}, update: map[string]string{"associatedData2": "associatedData2"}},
		"logging_context": Representation{repType: Optional, create: map[string]string{"loggingContext": "loggingContext"}, update: map[string]string{"loggingContext2": "loggingContext2"}},
	}

	EncryptedDataResourceDependencies = KeyResourceDependencyConfig
)

func TestKmsEncryptedDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsEncryptedDataResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_encrypted_data.test_encrypted_data"

	singularDatasourceName := "data.oci_kms_encrypted_data.test_encrypted_data"

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + EncryptedDataResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", Required, Create, encryptedDataRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttr(resourceName, "plaintext", "aGVsbG8sIHdvcmxk"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + EncryptedDataResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + EncryptedDataResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", Optional, Create, encryptedDataRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "associated_data.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "ciphertext"),
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttr(resourceName, "logging_context.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "plaintext", "aGVsbG8sIHdvcmxk"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_encrypted_data", "test_encrypted_data", Optional, Create, encryptedDataSingularDataSourceRepresentation) +
					compartmentIdVariableStr + EncryptedDataResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "associated_data.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "plaintext", "aGVsbG8sIHdvcmxk"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "ciphertext"),
				),
			},
		},
	})
}
