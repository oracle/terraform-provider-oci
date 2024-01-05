// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafetargetDatabasesTableDataSourceRepresentation = map[string]interface{}{
		"target_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_database.test_target_database.id}`},
		"schema_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`schemaName`}},
		"schema_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `schemaNameContains`},
		"table_name":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_nosql_table.test_table.name}`},
		"table_name_contains":  acctest.Representation{RepType: acctest.Optional, Create: `tableNameContains`},
	}

	DataSafeTargetDatabasesTableResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Required, acctest.Create, NosqlTableRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeTargetDatabasesTableResource_basic(t *testing.T) {
	t.Skip("Skipping this test which queries the database for metadata of tables present in the database. " +
		"This is specifically meant for Console, hence not a use case for terraform")
	httpreplay.SetScenario("TestDataSafeTargetDatabasesTableResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_target_databases_tables.test_target_databases_tables"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_databases_tables", "test_target_databases_tables", acctest.Required, acctest.Create, DataSafetargetDatabasesTableDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeTargetDatabasesTableResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "schema_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name_contains", "schemaNameContains"),
				resource.TestCheckResourceAttr(datasourceName, "table_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "table_name_contains", "tableNameContains"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "tables.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "tables.0.schema_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "tables.0.table_name"),
			),
		},
	})
}
