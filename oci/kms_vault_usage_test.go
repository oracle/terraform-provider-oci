// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	vaultUsageSingularDataSourceRepresentation = map[string]interface{}{
		"vault_id": Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.id}`},
	}

	VaultUsageResourceConfig = KeyResourceDependencies
)

func TestKmsVaultUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsVaultUsageResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_kms_vault_usage.test_vault_usage"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_vault_usage", "test_vault_usage", Required, Create, vaultUsageSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VaultUsageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_version_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "software_key_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "software_key_version_count"),
				),
			},
		},
	})
}
