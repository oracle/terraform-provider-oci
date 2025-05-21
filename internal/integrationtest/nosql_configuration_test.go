// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	NosqlConfigurationRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_nosql_configuration", "test_configuration",
		acctest.Optional, acctest.Create, NosqlConfigurationRepresentation)

	NosqlConfigurationResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_nosql_configuration", "test_configuration",
		acctest.Optional, acctest.Update, NosqlConfigurationRepresentation)

	NosqlConfigurationDataResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_configuration", "test_configuration",
		acctest.Required, acctest.Create, NosqlConfigurationSingularDataSourceRepresentation)

	UnassignKeyResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_nosql_configuration", "test_configuration",
		acctest.Required, acctest.Update, UnassignKeyRepresentation)

	NosqlConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	NosqlConfigurationRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id}`},
		"environment":    acctest.Representation{RepType: acctest.Required, Create: `HOSTED`, Update: `HOSTED`},
		"is_opc_dry_run": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"kms_key":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: NosqlConfigurationKmsKeyRepresentation},
	}

	NosqlConfigurationKmsKeyRepresentation = map[string]interface{}{
		"id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.key_id}`, Update: `${var.key_id2}`},
		"kms_vault_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`, Update: `${var.vault_id}`},
	}

	UnassignKeyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Update: `${var.compartment_id}`},
		"environment":    acctest.Representation{RepType: acctest.Required, Update: `HOSTED`},
		"kms_key":        acctest.RepresentationGroup{RepType: acctest.Required, Group: UnassignKeyKmsKeyRepresentation},
	}

	UnassignKeyKmsKeyRepresentation = map[string]interface{}{
		"id":           acctest.Representation{RepType: acctest.Required, Update: nil},
		"kms_vault_id": acctest.Representation{RepType: acctest.Required, Update: nil},
	}
)

// issue-routing-tag: nosql/default
func TestNosqlConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNosqlConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// Compartment_id is used to identify the tenancy.
	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Vault Id
	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	// Key Id
	keyId := utils.GetEnvSettingWithBlankDefault("key_id")
	keyIdVariableStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)

	// Key Id for rotating key
	keyId2 := utils.GetEnvSettingWithBlankDefault("key_id2")
	keyId2VariableStr := fmt.Sprintf("variable \"key_id2\" { default = \"%s\" }\n", keyId2)

	nosqlConfigurationResourceDependencies := keyIdVariableStr + keyId2VariableStr + vaultIdVariableStr

	resourceName := "oci_nosql_configuration.test_configuration"

	singularDatasourceName := "data.oci_nosql_configuration.test_configuration"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	assignkey_config := config + compartmentIdVariableStr + nosqlConfigurationResourceDependencies + NosqlConfigurationRequiredOnlyResource
	updatekey_config := config + compartmentIdVariableStr + nosqlConfigurationResourceDependencies + NosqlConfigurationResourceConfig

	get_configuration_config := config + compartmentIdVariableStr + NosqlConfigurationDataResourceConfig
	get_configuration_after_updatekey_config := get_configuration_config + nosqlConfigurationResourceDependencies + NosqlConfigurationResourceConfig
	get_configuration_after_unassignkey_config := get_configuration_config + UnassignKeyResourceConfig

	acctest.SaveConfigContent(assignkey_config, "nosql", "configuration", t)

	acctest.ResourceTest(
		t,
		nil,
		[]resource.TestStep{
			// verify Assign global encryption key
			{
				Config: assignkey_config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "environment", "HOSTED"),
					resource.TestCheckResourceAttr(resourceName, "kms_key.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kms_key.0.id", keyId),
					resource.TestCheckResourceAttr(resourceName, "kms_key.0.kms_key_state", "ACTIVE"),
					resource.TestCheckResourceAttr(resourceName, "kms_key.0.kms_vault_id", vaultId),

					func(s *terraform.State) (err error) {
						_, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify Update (rotate) global encryption key
			{
				Config: updatekey_config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "environment", "HOSTED"),
					resource.TestCheckResourceAttr(resourceName, "kms_key.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kms_key.0.id", keyId2),
					resource.TestCheckResourceAttr(resourceName, "kms_key.0.kms_key_state", "ACTIVE"),
					resource.TestCheckResourceAttr(resourceName, "kms_key.0.kms_vault_id", vaultId),

					func(s *terraform.State) (err error) {
						_, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify singular datasource
			{
				Config: get_configuration_after_updatekey_config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "environment", "HOSTED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "kms_key.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "kms_key.0.id", keyId2),
					resource.TestCheckResourceAttr(singularDatasourceName, "kms_key.0.kms_key_state", "ACTIVE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "kms_key.0.kms_vault_id", vaultId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "kms_key.0.time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "kms_key.0.time_updated"),
				),
			},

			// verify resource import
			{
				Config:                  updatekey_config,
				ImportState:             true,
				ImportStateIdFunc:       getConfigurationCompositeId(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_opc_dry_run"},
				ResourceName:            resourceName,
			},

			// verify key unassignment
			{
				Config: config + compartmentIdVariableStr + UnassignKeyResourceConfig,
			},

			// verify singular datasource after unassign key
			{
				Config: get_configuration_after_unassignkey_config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "environment", "HOSTED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "kms_key.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "kms_key.0.id", ""),
					resource.TestCheckResourceAttr(singularDatasourceName, "kms_key.0.kms_vault_id", ""),
					resource.TestCheckResourceAttr(singularDatasourceName, "kms_key.0.kms_key_state", "ACTIVE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "kms_key.0.time_created", ""),
					resource.TestCheckResourceAttr(singularDatasourceName, "kms_key.0.time_updated", ""),
				),
			},
		})
}

// Gets the composite Id of the configuration resource
func getConfigurationCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("configuration/compartmentId/%s", rs.Primary.Attributes["compartment_id"]), nil
	}
}
