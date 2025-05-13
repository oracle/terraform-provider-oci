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
	DatabaseManagementManagedMySqlDatabaseHighAvailabilityMemberDataSourceRepresentation = map[string]interface{}{
		"managed_my_sql_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id}`},
	}

	DatabaseManagementManagedMySqlDatabaseHighAvailabilityMemberResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_databases", "test_managed_my_sql_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedMySqlDatabaseHighAvailabilityMemberResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedMySqlDatabaseHighAvailabilityMemberResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_my_sql_database_high_availability_members.test_managed_my_sql_database_high_availability_members"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_database_high_availability_members", "test_managed_my_sql_database_high_availability_members", acctest.Required, acctest.Create, DatabaseManagementManagedMySqlDatabaseHighAvailabilityMemberDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedMySqlDatabaseHighAvailabilityMemberResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_high_availability_member_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_high_availability_member_collection.0.flow_control"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_high_availability_member_collection.0.group_auto_increment"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_high_availability_member_collection.0.group_name"),
				resource.TestCheckResourceAttr(datasourceName, "managed_my_sql_database_high_availability_member_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_high_availability_member_collection.0.member_role"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_high_availability_member_collection.0.member_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_high_availability_member_collection.0.single_primary_mode"),
				resource.TestCheckResourceAttr(datasourceName, "managed_my_sql_database_high_availability_member_collection.0.status_summary.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_high_availability_member_collection.0.transactions_in_gtid_executed"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_high_availability_member_collection.0.view_id"),
			),
		},
	})
}
