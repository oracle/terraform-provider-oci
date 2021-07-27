// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v45/common"
	oci_kms "github.com/oracle/oci-go-sdk/v45/keymanagement"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VaultRequiredOnlyResource = VaultResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, vaultRepresentation)

	VaultResourceConfig = VaultResourceDependencies +
		generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Update, vaultRepresentation)

	vaultSingularDataSourceRepresentation = map[string]interface{}{
		"vault_id": Representation{repType: Required, create: `${oci_kms_vault.test_vault.id}`},
	}

	vaultDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"filter":         RepresentationGroup{Required, vaultDataSourceFilterRepresentation}}
	vaultDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_kms_vault.test_vault.id}`}},
	}

	kmsVaultDeletionTime = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)

	vaultRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":     Representation{repType: Required, create: `Vault 1`, update: `displayName2`},
		"vault_type":       Representation{repType: Required, create: `VIRTUAL_PRIVATE`},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"time_of_deletion": Representation{repType: Optional, create: deletionTime.Format(time.RFC3339Nano)},
	}

	VaultResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: kms/default
func TestKmsVaultResource_basic(t *testing.T) {
	t.Skip("Skip this test till KMS provides a better way of testing this.")

	httpreplay.SetScenario("TestKmsVaultResource_basic")
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
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+VaultResourceDependencies+
		generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Create, vaultRepresentation), "keymanagement", "vault", t)

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
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, vaultRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Vault 1"),
					resource.TestCheckResourceAttr(resourceName, "vault_type", "VIRTUAL_PRIVATE"),

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
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Create, vaultRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Vault 1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "vault_type", "VIRTUAL_PRIVATE"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VaultResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Create,
						representationCopyWithNewProperties(vaultRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Vault 1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "vault_type", "VIRTUAL_PRIVATE"),

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
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Update, vaultRepresentation),
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
					resource.TestCheckResourceAttr(resourceName, "vault_type", "VIRTUAL_PRIVATE"),

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
					generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Optional, Update, vaultRepresentation),
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
					resource.TestCheckResourceAttr(datasourceName, "vaults.0.vault_type", "VIRTUAL_PRIVATE"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, vaultSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VaultResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "management_endpoint"),
					resource.TestCheckResourceAttr(singularDatasourceName, "replica_details.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "restored_from_vault_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "vault_type", "VIRTUAL_PRIVATE"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + VaultResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckKMSVaultDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).kmsVaultClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_kms_vault" {
			noResourceFound = false
			request := oci_kms.GetVaultRequest{}

			tmp := rs.Primary.ID
			request.VaultId = &tmp

			response, err := client.GetVault(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_kms.VaultLifecycleStatePendingDeletion): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				if !response.TimeOfDeletion.Equal(kmsVaultDeletionTime) && !httpreplay.ModeRecordReplay() {
					return fmt.Errorf("resource time_of_deletion: %s is not set to %s", response.TimeOfDeletion.Format(time.RFC3339Nano), kmsVaultDeletionTime.Format(time.RFC3339Nano))
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
