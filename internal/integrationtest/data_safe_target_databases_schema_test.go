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
	DataSafetargetDatabasesSchemaDataSourceRepresentation = map[string]interface{}{
		"target_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_database.test_target_database.id}`},
		"is_oracle_maintained": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"schema_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`schemaName`}},
		"schema_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `schemaNameContains`},
	}

	DataSafeTargetDatabasesSchemaResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Required, acctest.Create, targetDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeTargetDatabasesSchemaResource_basic(t *testing.T) {
	t.Skip("Skipping this test which queries the database for metadata of schema present in the database. " +
		"This is specifically meant for Console, hence not a use case for terraform")
	httpreplay.SetScenario("TestDataSafeTargetDatabasesSchemaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_target_databases_schemas.test_target_databases_schemas"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_databases_schemas", "test_target_databases_schemas", acctest.Required, acctest.Create, DataSafetargetDatabasesSchemaDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeTargetDatabasesSchemaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "is_oracle_maintained", "false"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name_contains", "schemaNameContains"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "schemas.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "schemas.0.is_oracle_maintained"),
				resource.TestCheckResourceAttrSet(datasourceName, "schemas.0.schema_name"),
			),
		},
	})
}
