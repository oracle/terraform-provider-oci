// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	vaultReplicaDataSourceRepresentation = map[string]interface{}{
		"vault_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.id}`},
	}

	VaultReplicaResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation)
)

// issue-routing-tag: kms/default
func TestKmsVaultReplicaResource_basic(t *testing.T) {
	t.Skip("Skip this test because virtual private vault is needed")
	httpreplay.SetScenario("TestKmsVaultReplicaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_kms_vault_replicas.test_vault_replicas"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_vault_replicas", "test_vault_replicas", acctest.Required, acctest.Create, vaultReplicaDataSourceRepresentation) +
				compartmentIdVariableStr + VaultReplicaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "vault_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "vault_replicas.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "vault_replicas.0.crypto_endpoint"),
				resource.TestCheckResourceAttrSet(datasourceName, "vault_replicas.0.management_endpoint"),
				resource.TestCheckResourceAttrSet(datasourceName, "vault_replicas.0.region"),
				resource.TestCheckResourceAttrSet(datasourceName, "vault_replicas.0.status"),
			),
		},
	})
}
