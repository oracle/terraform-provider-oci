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
	secretSingularDataSourceRepresentation = map[string]interface{}{
		"secret_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_secret.id}`},
	}

	secretDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.name}`},
		"vault_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}

	SecretResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: vault/default
func TestVaultSecretResource_basic(t *testing.T) {
	t.Skip("Skip this test till Secret Management service provides a better way of testing this.")
	httpreplay.SetScenario("TestVaultSecretResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_vault_secrets.test_secrets"
	singularDatasourceName := "data.oci_vault_secret.test_secret"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_vault_secrets", "test_secrets", acctest.Optional, acctest.Update, secretDataSourceRepresentation) +
				compartmentIdVariableStr + SecretResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "name"),
				//resource.TestCheckResourceAttr(datasourceName, "state", "Active"),
				resource.TestCheckResourceAttrSet(datasourceName, "vault_id"),

				resource.TestCheckResourceAttr(datasourceName, "secrets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "secrets.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.secret_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.time_created"),
				//resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.time_of_current_version_expiry"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.time_of_deletion"),
				resource.TestCheckResourceAttrSet(datasourceName, "secrets.0.vault_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, secretSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "secret_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_version_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "secret_rules.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_deletion"),
			),
		},
	})
}
