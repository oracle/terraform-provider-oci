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
	SecretsSecretsSecretbundleVersionDataSourceRepresentation = map[string]interface{}{
		"secret_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_secret.id}`},
	}

	SecretsSecretbundleVersionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation)
)

// issue-routing-tag: secrets/default
func TestSecretsSecretbundleVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSecretsSecretbundleVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	vaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_ocid")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	keyId := utils.GetEnvSettingWithBlankDefault("kms_key_ocid")
	keyIdVariableStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)

	datasourceName := "data.oci_secrets_secretbundle_versions.test_secretbundle_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + vaultIdVariableStr + keyIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_secrets_secretbundle_versions", "test_secretbundle_versions", acctest.Required, acctest.Create, SecretsSecretsSecretbundleVersionDataSourceRepresentation) +
				compartmentIdVariableStr + SecretsSecretbundleVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "secret_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "secret_bundle_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "secret_bundle_versions.0.secret_id"),
				resource.TestCheckResourceAttr(datasourceName, "secret_bundle_versions.0.stages.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "secret_bundle_versions.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "secret_bundle_versions.0.version_number"),
			),
		},
	})
}
