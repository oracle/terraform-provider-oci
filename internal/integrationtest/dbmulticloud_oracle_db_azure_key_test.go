// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DbmulticloudOracleDbAzureKeyResourceDependencies             = acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureVaultRepresentation)
	DbmulticloudOracleDbAzureKeySingularDataSourceRepresentation = map[string]interface{}{

		"oracle_db_azure_key_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_dbmulticloud_oracle_db_azure_keys.test_oracle_db_azure_keys.oracle_db_azure_key_summary_collection.0.items.0.id}`},
	}

	DbmulticloudOracleDbAzureKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		//"display_name":             acctest.Representation{RepType: acctest.Required, Create: `MockResourceName`},
		//"oracle_db_azure_key_id":   acctest.Representation{RepType: acctest.Required, Create: `MockKeyId`},
		"oracle_db_azure_vault_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault.id}`},
	}
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbAzureKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbAzureKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_dbmulticloud_oracle_db_azure_keys.test_oracle_db_azure_keys"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_azure_key.test_oracle_db_azure_key"
	//associationResourceName := "oci_dbmulticloud_oracle_db_azure_vault_association.test_oracle_db_azure_vault_association"
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + DbmulticloudOracleDbAzureConnectorRequiredOnlyResource + DbmulticloudOracleDbAzureKeyResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_keys", "test_oracle_db_azure_keys", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureKeyDataSourceRepresentation) + compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_vault_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_key_summary_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + DbmulticloudOracleDbAzureConnectorRequiredOnlyResource + DbmulticloudOracleDbAzureKeyResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault_associations", "test_oracle_db_azure_vault_associations", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureVaultAssociationRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_keys", "test_oracle_db_azure_keys", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureKeyDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_key", "test_oracle_db_azure_key", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureKeySingularDataSourceRepresentation) + compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_azure_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "azure_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_properties.%"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				func(s *terraform.State) (err error) {
					//for k, v := range s.RootModule().Resources {
					//	fmt.Printf("Association Resource: %s, ID: %s\n", k, v.Primary.ID)
					//}
					rs, ok := s.RootModule().Resources["oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault"]
					if !ok {
						return fmt.Errorf("vaultId not found in state")
					}
					vaultId := rs.Primary.ID
					return deleteDefaultAssocationCall(compartmentId, vaultId)
				},
			),
		},
	})
}

func deleteDefaultAssocationCall(compartment string, vaultId string) error {
	associationIds, err := getVaultAssocationList(compartment, vaultId)
	if err != nil {
		return err
	}
	for _, associationId := range associationIds {
		oracleDbAzureVaultAssociationClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDbAzureVaultAssociationClient()
		deleteOracleDbAzureVaultAssociationRequest := oci_dbmulticloud.DeleteOracleDbAzureVaultAssociationRequest{}

		deleteOracleDbAzureVaultAssociationRequest.OracleDbAzureVaultAssociationId = &associationId

		deleteOracleDbAzureVaultAssociationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
		_, error := oracleDbAzureVaultAssociationClient.DeleteOracleDbAzureVaultAssociation(context.Background(), deleteOracleDbAzureVaultAssociationRequest)
		if error != nil {
			fmt.Printf("Error deleting OracleDbAzureVaultAssociation %s %s, It is possible that the resource is already deleted. Please verify manually \n", associationId, error)
			continue
		}
	}
	return nil
}

func getVaultAssocationList(compartment string, vaultId string) ([]string, error) {
	var resourceIds []string
	compartmentId := compartment
	oracleDbAzureVaultAssociationClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDbAzureVaultAssociationClient()

	listOracleDbAzureVaultAssociationsRequest := oci_dbmulticloud.ListOracleDbAzureVaultAssociationsRequest{}
	listOracleDbAzureVaultAssociationsRequest.CompartmentId = &compartmentId
	listOracleDbAzureVaultAssociationsRequest.OracleDbAzureVaultId = &vaultId
	listOracleDbAzureVaultAssociationsResponse, err := oracleDbAzureVaultAssociationClient.ListOracleDbAzureVaultAssociations(context.Background(), listOracleDbAzureVaultAssociationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VaultAssocationList list for compartment id : %s  %s \n", compartmentId, err)
	}
	for _, association := range listOracleDbAzureVaultAssociationsResponse.Items {
		id := *association.Id
		resourceIds = append(resourceIds, id)
	}
	return resourceIds, nil
}
