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
	secretVersionSingularDataSourceRepresentation = map[string]interface{}{
		"secret_id":             Representation{repType: Required, create: `${oci_vault_secret.test_secret.id`},
		"secret_version_number": Representation{repType: Required, create: `1`},
	}

	SecretVersionResourceConfig = ``
)

func TestVaultSecretVersionResource_basic(t *testing.T) {
	t.Skip("Skip this test till Secret Management service provides a better way of testing this.")
	httpreplay.SetScenario("TestVaultSecretVersionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_vault_secret_version.test_secret_version"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_vault_secret_version", "test_secret_version", Required, Create, secretVersionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + SecretVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "secret_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "secret_version_number", "1"),

					//resource.TestCheckResourceAttrSet(singularDatasourceName, "content_type"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
					resource.TestCheckResourceAttr(singularDatasourceName, "stages.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_current_version_expiry"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_deletion"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "version_number"),
				),
			},
		},
	})
}
