// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	KmsVaultRequiredOnlyResource = KmsVaultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation)

	KmsVaultResourceConfig = KmsVaultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, KmsVaultRepresentation)

	KmsKmsVaultSingularDataSourceRepresentation = map[string]interface{}{
		"vault_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_vault.test_vault.id}`},
	}

	KmsKmsVaultDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsVaultDataSourceFilterRepresentation}}
	KmsVaultDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_kms_vault.test_vault.id}`}},
	}

	kmsVaultDeletionTime = time.Now().UTC().AddDate(0, 0, 8).Truncate(time.Millisecond)

	privateEndpointId            = utils.GetEnvSettingWithBlankDefault("ekms_private_endpoint_id")
	privateEndpointIdVariableStr = fmt.Sprintf("variable \"private_endpoint_id\" { default = \"%s\" }\n", privateEndpointId)

	KmsVaultExternalKeyManagerMetadataRepresentation = map[string]interface{}{
		"external_vault_endpoint_url": acctest.Representation{RepType: acctest.Required, Create: `https://10.0.0.31/api/v1/cckm/oci/ekm/v1/vaults/af872d6e-52f2-4c6b-9694-5b4821d1b5b6`},
		"oauth_metadata":              acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsVaultExternalKeyManagerMetadataOauthMetadataRepresentation},
		"private_endpoint_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.private_endpoint_id}`},
	}

	KmsVaultExternalKeyManagerMetadataOauthMetadataRepresentation = map[string]interface{}{
		"client_app_id":         acctest.Representation{RepType: acctest.Required, Create: `3977f2b65fca4c569f31142959867127`},
		"client_app_secret":     acctest.Representation{RepType: acctest.Required, Create: `d82452e5-f5e3-4363-b7a9-0d74052d1236`},
		"idcs_account_name_url": acctest.Representation{RepType: acctest.Required, Create: `https://idcs-87920edcd339458790351b0e4d415385.identity.oraclecloud.com`},
	}

	KmsVaultRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `Vault 1`, Update: `displayName2`},
		"vault_type":       acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"time_of_deletion": acctest.Representation{RepType: acctest.Optional, Create: deletionTime.Format(time.RFC3339Nano)},
	}

	KmsExternalVaultRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `External Vault 1`, Update: `displayName2`},
		"vault_type":                    acctest.Representation{RepType: acctest.Required, Create: `EXTERNAL`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"time_of_deletion":              acctest.Representation{RepType: acctest.Optional, Create: deletionTime.Format(time.RFC3339Nano)},
		"external_key_manager_metadata": acctest.RepresentationGroup{RepType: acctest.Required, Group: KmsVaultExternalKeyManagerMetadataRepresentation},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesExtVaultRepresentation},
	}

	ignoreChangesExtVaultRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}
	KmsVaultResourceDependencies = DefinedTagsDependencies
)

func TestKmsExternalVaultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestKmsExternalVaultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_vault.test_vault"
	datasourceName := "data.oci_kms_vaults.test_vaults"
	singularDatasourceName := "data.oci_kms_vault.test_vault"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Create, KmsExternalVaultRepresentation), "keymanagement", "vault", t)

	acctest.ResourceTest(t, testAccCheckKMSVaultDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + privateEndpointIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsExternalVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "External Vault 1"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "EXTERNAL"),
				resource.TestCheckResourceAttr(resourceName, "external_key_manager_metadata.#", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + privateEndpointIdVariableStr + KmsVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Create, KmsExternalVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "External Vault 1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "external_key_manager_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "EXTERNAL"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + privateEndpointIdVariableStr + KmsVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, KmsExternalVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "EXTERNAL"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_vaults", "test_vaults", acctest.Optional, acctest.Update, KmsKmsVaultDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, KmsVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "vaults.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vaults.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.crypto_endpoint"),
				resource.TestCheckResourceAttr(datasourceName, "vaults.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.management_endpoint"),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "vaults.0.vault_type", "EXTERNAL"),
				resource.TestCheckResourceAttr(resourceName, "external_key_manager_metadata.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsKmsVaultSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_endpoint"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_type", "EXTERNAL"),
				resource.TestCheckResourceAttr(resourceName, "external_key_manager_metadata.#", "1"),
			),
		},
	})
}

// issue-routing-tag: kms/default
func TestKmsVaultResource_basic(t *testing.T) {
	//t.Skip("Skip this test till KMS provides a better way of testing this.")

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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+KmsVaultResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Create, KmsVaultRepresentation), "keymanagement", "vault", t)

	acctest.ResourceTest(t, testAccCheckKMSVaultDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + KmsVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Vault 1"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + KmsVaultResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + KmsVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Create, KmsVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Vault 1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "vault_type", "DEFAULT"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + KmsVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(KmsVaultRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Vault 1"),
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
			Config: config + compartmentIdVariableStr + KmsVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, KmsVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_vaults", "test_vaults", acctest.Optional, acctest.Update, KmsKmsVaultDataSourceRepresentation) +
				compartmentIdVariableStr + KmsVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Optional, acctest.Update, KmsVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "vaults.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vaults.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "vaults.0.crypto_endpoint"),
				resource.TestCheckResourceAttr(datasourceName, "vaults.0.display_name", "displayName2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsKmsVaultSingularDataSourceRepresentation) +
				compartmentIdVariableStr + KmsVaultResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_vault_replicable"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_endpoint"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_type", "DEFAULT"),
			),
		},
		// verify resource import
		{
			Config:                  config + KmsVaultRequiredOnlyResource,
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
