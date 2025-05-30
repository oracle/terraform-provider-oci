// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DbmulticloudOracleDbAzureKeySingularDataSourceRepresentation = map[string]interface{}{

		"oracle_db_azure_key_id": acctest.Representation{RepType: acctest.Required, Create: os.Getenv("TF_VAR_oracle_db_azure_key_id")},
	}

	DbmulticloudOracleDbAzureKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `MockResourceName`},
		"oracle_db_azure_key_id":   acctest.Representation{RepType: acctest.Required, Create: os.Getenv("TF_VAR_oracle_db_azure_key_id")},
		"oracle_db_azure_vault_id": acctest.Representation{RepType: acctest.Required, Create: os.Getenv("TF_VAR_oracle_db_azure_vault_id")},
	}

	DbmulticloudOracleDbAzureKeyResourceConfig = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `MockResourceName`},
		"oracle_db_azure_key_id":   acctest.Representation{RepType: acctest.Required, Create: os.Getenv("TF_VAR_oracle_db_azure_key_id")},
		"oracle_db_azure_vault_id": acctest.Representation{RepType: acctest.Required, Create: os.Getenv("TF_VAR_oracle_db_azure_vault_id")},
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

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_keys", "test_oracle_db_azure_keys", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureKeyDataSourceRepresentation) + compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "MockResourceName"),
				resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_vault_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_key_summary_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_key", "test_oracle_db_azure_key", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureKeySingularDataSourceRepresentation) + compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_azure_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "azure_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
