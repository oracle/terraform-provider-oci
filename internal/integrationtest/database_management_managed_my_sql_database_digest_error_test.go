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
	DatabaseManagementManagedMySqlDatabaseDigestErrorDataSourceRepresentation = map[string]interface{}{
		"digest":                     acctest.Representation{RepType: acctest.Required, Create: `digest`},
		"managed_my_sql_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id}`},
	}

	DatabaseManagementManagedMySqlDatabaseDigestErrorResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_databases", "test_managed_my_sql_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedMySqlDatabaseDigestErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedMySqlDatabaseDigestErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_my_sql_database_digest_errors.test_managed_my_sql_database_digest_errors"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_database_digest_errors", "test_managed_my_sql_database_digest_errors", acctest.Required, acctest.Create, DatabaseManagementManagedMySqlDatabaseDigestErrorDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedMySqlDatabaseDigestErrorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "digest", "digest"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "my_sql_digest_errors_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "my_sql_digest_errors_collection.0.items.#", "1"),
			),
		},
	})
}
