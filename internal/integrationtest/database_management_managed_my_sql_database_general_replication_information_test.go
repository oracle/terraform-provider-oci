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
	DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationSingularDataSourceRepresentation = map[string]interface{}{
		"managed_my_sql_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id}`},
	}

	DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_databases", "test_managed_my_sql_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_management_managed_my_sql_database_general_replication_information.test_managed_my_sql_database_general_replication_information"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_database_general_replication_information", "test_managed_my_sql_database_general_replication_information", acctest.Required, acctest.Create, DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_my_sql_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "apply_status_summary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "binary_log_format"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "binary_logging"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "executed_gtid_set"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fetch_status_summary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gtid_mode"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "high_availability_member_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "inbound_replications_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_high_availability_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "outbound_replications_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "port"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "read_only"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "seconds_behind_source_max"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "server_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "server_uuid"),
			),
		},
	})
}
