// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseMigrationMigrationDataVerificationDetailSingularDataSourceRepresentation = map[string]interface{}{
		"migration_id": acctest.Representation{RepType: acctest.Required, Create: `${var.data_verification_migration_id}`},
	}
)

/*
 Note:
  this test requires an existing migration with data verification details.
  Set the following environment variable in order to make it work:
   TF_VAR_data_verification_migration_id=ocid_of_existing_migration
*/
// issue-routing-tag: database_migration/default
func TestDatabaseMigrationMigrationDataVerificationDetailResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationMigrationDataVerificationDetailResource_basic")
	defer httpreplay.SaveScenario()

	const DATA_VERIFICATION_MIGRATION_ID = "data_verification_migration_id"

	config := acctest.ProviderTestConfig() + makeVariableStr(DATA_VERIFICATION_MIGRATION_ID, t)
	migrationId := utils.GetEnvSettingWithBlankDefault(DATA_VERIFICATION_MIGRATION_ID)

	singularDatasourceName := "data.oci_database_migration_migration_data_verification_detail.test_migration_data_verification_detail"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_migration_data_verification_detail", "test_migration_data_verification_detail", acctest.Required, acctest.Create, DatabaseMigrationMigrationDataVerificationDetailSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "migration_id", migrationId),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_completed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
