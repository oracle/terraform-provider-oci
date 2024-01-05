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
	SecretsSecretsSecretbundleSingularDataSourceRepresentation = map[string]interface{}{
		"secret_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_secret.id}`},
		"secret_version_name": acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"stage":               acctest.Representation{RepType: acctest.Optional, Create: `CURRENT`},
		"version_number":      acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	SecretsSecretbundleResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation)
)

// issue-routing-tag: secrets/default
func TestSecretsSecretbundleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSecretsSecretbundleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	vaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_ocid")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	keyId := utils.GetEnvSettingWithBlankDefault("kms_key_ocid")
	keyIdVariableStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)

	singularDatasourceName := "data.oci_secrets_secretbundle.test_secretbundle"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + vaultIdVariableStr + keyIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_secrets_secretbundle", "test_secretbundle", acctest.Required, acctest.Create, SecretsSecretsSecretbundleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SecretsSecretbundleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "secret_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "secret_bundle_content.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
