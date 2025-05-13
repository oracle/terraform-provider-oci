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
	DatabaseManagementManagedMySqlDatabaseInboundReplicationDataSourceRepresentation = map[string]interface{}{
		"managed_my_sql_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id}`},
	}

	DatabaseManagementManagedMySqlDatabaseInboundReplicationResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_databases", "test_managed_my_sql_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedMySqlDatabaseInboundReplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedMySqlDatabaseInboundReplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_my_sql_database_inbound_replications.test_managed_my_sql_database_inbound_replications"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_database_inbound_replications", "test_managed_my_sql_database_inbound_replications", acctest.Required, acctest.Create, DatabaseManagementManagedMySqlDatabaseInboundReplicationDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedMySqlDatabaseInboundReplicationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_inbound_replication_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_inbound_replication_collection.0.inbound_replications_count"),
				resource.TestCheckResourceAttr(datasourceName, "managed_my_sql_database_inbound_replication_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_inbound_replication_collection.0.parallel_workers"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_inbound_replication_collection.0.preserve_commit_order"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_inbound_replication_collection.0.replica_server_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_inbound_replication_collection.0.replica_uuid"),
			),
		},
	})
}
