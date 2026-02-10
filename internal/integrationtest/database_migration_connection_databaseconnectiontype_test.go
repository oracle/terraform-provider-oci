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
	DatabaseMigrationConnectionDatabaseconnectiontypeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_type":      acctest.Representation{RepType: acctest.Required, Create: []string{`ORACLE`}},
		"source_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_standby_oracle_id}`},
		"technology_type":      acctest.Representation{RepType: acctest.Required, Create: []string{`ORACLE_DATABASE`}},
	}
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationConnectionDatabaseconnectiontypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationConnectionDatabaseconnectiontypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sourceConnectionOracleId := utils.GetEnvSettingWithBlankDefault("source_connection_standby_oracle_id")
	sourceConnectionOracleIdVariableStr := fmt.Sprintf("variable \"source_connection_standby_oracle_id\" { default = \"%s\" }\n", sourceConnectionOracleId)

	datasourceName := "data.oci_database_migration_connection_databaseconnectiontypes.test_connection_databaseconnectiontypes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_connection_databaseconnectiontypes", "test_connection_databaseconnectiontypes", acctest.Required, acctest.Create, DatabaseMigrationConnectionDatabaseconnectiontypeDataSourceRepresentation) +
				compartmentIdVariableStr + sourceConnectionOracleIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "connection_type.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "source_connection_id"),
				resource.TestCheckResourceAttr(datasourceName, "technology_type.#", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "database_connection_type_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "database_connection_type_collection.0.items.#", "1"),
			),
		},
	})
}
