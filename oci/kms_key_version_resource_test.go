// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	KeyVersionVirtualResourceConfig = KeyVersionVirtualResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Required, Create, keyVersionVirtualRepresentation)

	keyVersionVirtualSingularDataSourceRepresentation = getMultipleUpdatedRepresenationCopy([]string{"key_id", "management_endpoint"}, []interface{}{
		Representation{repType: Required, create: `${oci_kms_key.test_key.id}`},
		Representation{repType: Required, create: `${data.oci_kms_vault.test_virtual_vault.management_endpoint}`}}, keyVersionSingularDataSourceRepresentation)

	keyVersionVirtualDataSourceRepresentation = getMultipleUpdatedRepresenationCopy([]string{"key_id", "management_endpoint"}, []interface{}{
		Representation{repType: Required, create: `${oci_kms_key.test_key.id}`},
		Representation{repType: Required, create: `${data.oci_kms_vault.test_virtual_vault.management_endpoint}`}}, keyVersionDataSourceRepresentation)

	keyVersionVirtualRepresentation = getMultipleUpdatedRepresenationCopy([]string{"key_id", "management_endpoint"}, []interface{}{
		Representation{repType: Required, create: `${oci_kms_key.test_key.id}`},
		Representation{repType: Required, create: `${data.oci_kms_vault.test_virtual_vault.management_endpoint}`}}, keyVersionRepresentation)

	KeyVersionVirtualResourceDependencies = KeyResourceVirtualDependencies + DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_kms_key", "test_key", Optional, Create, representationCopyWithRemovedProperties(keyVirtualRepresentation, []string{"desired_state"}))
)

func TestResourceKmsKeyVersionVirtualResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceKmsKeyVersionVirtualResource_basic")
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + KeyVersionVirtualResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Required, Create, keyVersionVirtualRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_key_versions", "test_key_versions", Optional, Update, keyVersionVirtualDataSourceRepresentation) +
					compartmentIdVariableStr + KeyVersionVirtualResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Optional, Update, keyVersionVirtualRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_kms_key_version", "test_key_version", Required, Create, keyVersionVirtualSingularDataSourceRepresentation) +
					compartmentIdVariableStr + KeyVersionVirtualResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_version_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "management_endpoint"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + KeyVersionVirtualResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: keyVersionImportId,
				ImportStateVerifyIgnore: []string{
					"time_of_deletion",
				},
				ResourceName: resourceName,
			},
		},
	})
}
