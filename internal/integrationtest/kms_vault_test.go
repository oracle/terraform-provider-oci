// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_kms "github.com/oracle/oci-go-sdk/v58/keymanagement"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	VaultRequiredOnlyResource = VaultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation)

	VaultResourceConfig = VaultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, vaultRepresentation)

	vaultSingularDataSourceRepresentation = map[string]interface{}{
		"vault_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_vault.test_vault.id}`},
	}

	vaultDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: vaultDataSourceFilterRepresentation}}
	vaultDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_kms_vault.test_vault.id}`}},
	}

	kmsVaultDeletionTime = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)

	vaultRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `Vault 1`, Update: `displayName2`},
		"vault_type":       acctest.Representation{RepType: acctest.Required, Create: `VIRTUAL_PRIVATE`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"time_of_deletion": acctest.Representation{RepType: acctest.Optional, Create: deletionTime.Format(time.RFC3339Nano)},
	}

	VaultResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: kms/default
func TestKmsVaultResource_basic(t *testing.T) {
	t.Skip("Skip this test till KMS provides a better way of testing this.")

	httpreplay.SetScenario("TestKmsVaultResource_basic")
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
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+VaultResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Create, vaultRepresentation), "keymanagement", "vault", t)

	acctest.ResourceTest(t, testAccCheckKMSVaultDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Vault 1"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "VIRTUAL_PRIVATE"),

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
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Create, vaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Vault 1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "VIRTUAL_PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(vaultRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Vault 1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "VIRTUAL_PRIVATE"),

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
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, vaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "VIRTUAL_PRIVATE"),

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
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, vaultRepresentation),
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
				resource.TestCheckResourceAttr(datasourceName, "vaults.0.vault_type", "VIRTUAL_PRIVATE"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VaultResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
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
	})
}

func testAccCheckKMSVaultDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).KmsVaultClient()
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
