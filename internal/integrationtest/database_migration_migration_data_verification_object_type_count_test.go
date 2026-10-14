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
	DatabaseMigrationMigrationDataVerificationObjectTypeCountDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"migration_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.data_verification_migration_id}`},
	}
)

/*
 Note:
  this test requires an existing migration with completed data verification.
  Set the following environment variables in order to make it work:
   TF_VAR_compartment_id=compartment_id_of_the_migration
   TF_VAR_data_verification_migration_id=ocid_of_existing_migration
*/
// issue-routing-tag: database_migration/default
func TestDatabaseMigrationMigrationDataVerificationObjectTypeCountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationMigrationDataVerificationObjectTypeCountResource_basic")
	defer httpreplay.SaveScenario()

	const (
		COMPARTMENT_ID                 = "compartment_id"
		DATA_VERIFICATION_MIGRATION_ID = "data_verification_migration_id"
	)

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t) +
		makeVariableStr(DATA_VERIFICATION_MIGRATION_ID, t)
	compartmentId := utils.GetEnvSettingWithBlankDefault(COMPARTMENT_ID)
	migrationId := utils.GetEnvSettingWithBlankDefault(DATA_VERIFICATION_MIGRATION_ID)

	datasourceName := "data.oci_database_migration_migration_data_verification_object_type_counts.test_migration_data_verification_object_type_counts"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_migration_data_verification_object_type_counts", "test_migration_data_verification_object_type_counts", acctest.Required, acctest.Create, DatabaseMigrationMigrationDataVerificationObjectTypeCountDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "migration_id", migrationId),

				resource.TestCheckResourceAttrSet(datasourceName, "data_verification_object_type_count_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_verification_object_type_count_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_verification_object_type_count_collection.0.total_items"),
			),
		},
	})
}
