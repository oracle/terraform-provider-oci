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
	virtualVaultRepresentation = getMultipleUpdatedRepresenationCopy([]string{"display_name", "vault_type"},
		[]interface{}{Representation{repType: Required, create: `DEFAULT_VAULT`, update: `displayName2`},
			Representation{repType: Required, create: `DEFAULT`}}, vaultRepresentation)
)

// issue-routing-tag: kms/default
func TestResourceKmsVaultResource_default(t *testing.T) {
	t.Skip("Skip this test till KMS provides a better way of testing this.")

	httpreplay.SetScenario("TestResourceKmsVaultResource_virtual")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_kms_vault.test_vault"
	datasourceName := "data.oci_kms_vaults.test_vaults"
	singularDatasourceName := "data.oci_kms_vault.test_vault"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckKMSVaultDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, virtualVaultRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "DEFAULT_VAULT"),
					resource.TestCheckResourceAttr(resourceName, "vault_type", "DEFAULT"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VaultResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Create, virtualVaultRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "DEFAULT_VAULT"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "vault_type", "DEFAULT"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Create,
						representationCopyWithNewProperties(virtualVaultRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "DEFAULT_VAULT"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "vault_type", "DEFAULT"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Update, virtualVaultRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "vault_type", "DEFAULT"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_kms_vaults", "test_vaults", Optional, Update, vaultDataSourceRepresentation) +
					compartmentIdVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Update, virtualVaultRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "vaults.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "vaults.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.crypto_endpoint"),
					resource.TestCheckResourceAttr(datasourceName, "vaults.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, vaultSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Update, virtualVaultRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Update, virtualVaultRepresentation),
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
		},
	})
}
