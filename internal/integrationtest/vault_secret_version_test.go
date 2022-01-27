// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	secretVersionSingularDataSourceRepresentation = map[string]interface{}{
		"secret_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_secret.id`},
		"secret_version_number": acctest.Representation{RepType: acctest.Required, Create: `10`},
	}

	SecretVersionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, secretRepresentation)
)

// issue-routing-tag: vault/default
func TestVaultSecretVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestVaultSecretVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_vault_secret_version.test_secret_version"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_vault_secret_version", "test_secret_version", acctest.Required, acctest.Create, secretVersionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SecretVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "secret_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "secret_version_number", "1"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "content_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.#", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_current_version_expiry"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_deletion"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version_number"),
			),
		},
	})
}
