// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DbmulticloudOracleDbGcpKeySingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_gcp_key_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_dbmulticloud_oracle_db_gcp_keys.test_oracle_db_gcp_keys.oracle_db_gcp_key_summary_collection.0.items.0.id}`},
	}

	DbmulticloudOracleDbGcpKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		//"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		//"oracle_db_gcp_key_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_dbmulticloud_oracle_db_gcp_key.test_oracle_db_gcp_key.id}`},
		"oracle_db_gcp_key_ring_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring.id}`},
	}
	//DbmulticloudOracleDbGcpKeyResourceConfig = ""
	DbmulticloudOracleDbGcpKeyResourceConfig = DbmulticloudOracleDbGcpIdentityConnectorRequiredOnlyResource + acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key_ring", "test_oracle_db_gcp_key_ring", acctest.Required, acctest.Create, DbmulticloudOracleDbGcpKeyRingRepresentation)
	//acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_keys", "test_oracle_db_gcp_keys", acctest.Required, acctest.Create, DbmulticloudOracleDbGcpKeyDataSourceRepresentation)
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbGcpKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbGcpKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_dbmulticloud_oracle_db_gcp_keys.test_oracle_db_gcp_keys"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_gcp_key.test_oracle_db_gcp_key"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_keys", "test_oracle_db_gcp_keys", acctest.Required, acctest.Create, DbmulticloudOracleDbGcpKeyDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbGcpKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_gcp_key_summary_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_keys", "test_oracle_db_gcp_keys", acctest.Required, acctest.Create, DbmulticloudOracleDbGcpKeyDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key", "test_oracle_db_gcp_key", acctest.Required, acctest.Create, DbmulticloudOracleDbGcpKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbGcpKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_gcp_key_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gcp_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gcp_key_properties.%"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
