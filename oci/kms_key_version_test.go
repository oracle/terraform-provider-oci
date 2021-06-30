// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	KeyVersionResourceConfig = KeyVersionResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Required, Create, keyVersionRepresentation)

	keyVersionSingularDataSourceRepresentation = map[string]interface{}{
		"key_id":              Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"key_version_id":      Representation{repType: Required, create: `${oci_kms_key_version.test_key_version.key_version_id}`},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
	}

	keyVersionDataSourceRepresentation = map[string]interface{}{
		"key_id":              Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"filter":              RepresentationGroup{Required, keyVersionDataSourceFilterRepresentation}}
	keyVersionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `key_version_id`},
		"values": Representation{repType: Required, create: []string{`${oci_kms_key_version.test_key_version.key_version_id}`}},
	}

	keyVersionDeletionTime = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)

	keyVersionRepresentation = map[string]interface{}{
		"key_id":              Representation{repType: Required, create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"management_endpoint": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"time_of_deletion":    Representation{repType: Required, create: keyVersionDeletionTime.Format(time.RFC3339Nano)},
	}

	KeyVersionResourceDependencies = KeyResourceDependencyConfig
)

func TestKmsKeyVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsKeyVersionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()
	os.Setenv("disable_kms_version_delete", "true")

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_kms_key_version.test_key_version"
	datasourceName := "data.oci_kms_key_versions.test_key_versions"
	singularDatasourceName := "data.oci_kms_key_version.test_key_version"

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+KeyVersionResourceDependencies+
		generateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Required, Create, keyVersionRepresentation), "keymanagement", "keyVersion", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + KeyVersionResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Required, Create, keyVersionRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_key_versions", "test_key_versions", Optional, Update, keyVersionDataSourceRepresentation) +
					compartmentIdVariableStr + KeyVersionResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Optional, Update, keyVersionRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "key_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_endpoint"),

					resource.TestCheckResourceAttr(datasourceName, "key_versions.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.key_version_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.key_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "key_versions.0.vault_id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Required, Create, keyVersionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + KeyVersionResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_version_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "management_endpoint"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
					resource.TestCheckResourceAttr(singularDatasourceName, "replica_details.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + KeyVersionResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: keyVersionImportId,
				ImportStateVerifyIgnore: []string{
					"management_endpoint",
					"time_of_deletion",
					"replica_details",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func keyVersionImportId(state *terraform.State) (string, error) {
	for _, rs := range state.RootModule().Resources {
		if rs.Type == "oci_kms_key_version" {
			return fmt.Sprintf("managementEndpoint/%s/%s", rs.Primary.Attributes["management_endpoint"], rs.Primary.ID), nil
		}
	}

	return "", fmt.Errorf("unable to create import id as no resource of type oci_kms_key_version in state")
}
