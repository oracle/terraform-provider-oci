// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	migrationObjectTypeDataSourceRepresentation = map[string]interface{}{}

	MigrationObjectTypeResourceConfig = ""
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationMigrationObjectTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationMigrationObjectTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_migration_migration_object_types.test_migration_object_types"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_migration_migration_object_types", "test_migration_object_types", Required, Create, migrationObjectTypeDataSourceRepresentation) +
				compartmentIdVariableStr + MigrationObjectTypeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "migration_object_type_summary_collection.#"),
			),
		},
	})
}
