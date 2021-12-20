// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	virtualVaultRepresentation = acctest.GetMultipleUpdatedRepresenationCopy([]string{"display_name", "vault_type"},
		[]interface{}{acctest.Representation{RepType: acctest.Required, Create: `DEFAULT_VAULT`, Update: `displayName2`},
			acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`}}, vaultRepresentation)
)

// issue-routing-tag: kms/default
func TestResourceKmsVaultResource_default(t *testing.T) {
	t.Skip("Skip this test till KMS provides a better way of testing this.")

	httpreplay.SetScenario("TestResourceKmsVaultResource_virtual")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_kms_vault.test_vault"
	datasourceName := "data.oci_kms_vaults.test_vaults"
	singularDatasourceName := "data.oci_kms_vault.test_vault"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckKMSVaultDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, virtualVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DEFAULT_VAULT"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Create, virtualVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DEFAULT_VAULT"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(virtualVaultRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "DEFAULT_VAULT"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, virtualVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_vaults", "test_vaults", acctest.Optional, acctest.Update, vaultDataSourceRepresentation) +
				compartmentIdVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, virtualVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "vaults.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vaults.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.crypto_endpoint"),
				resource.TestCheckResourceAttr(datasourceName, "vaults.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "vaults.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.management_endpoint"),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "vaults.0.vault_type", "DEFAULT"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, virtualVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_type", "DEFAULT"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, virtualVaultRepresentation),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"time_of_deletion",
			},
			ResourceName: resourceName,
		},
	})
}
