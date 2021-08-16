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
	databasePdbConversionHistoryEntrySingularDataSourceRepresentation = map[string]interface{}{
		"database_id":                     Representation{RepType: Required, Create: getEnvSettingWithBlankDefault("converted_to_pdb_db_id")},
		"pdb_conversion_history_entry_id": Representation{RepType: Required, Create: getEnvSettingWithBlankDefault("pdb_conversion_history_entry_id")},
	}

	databasePdbConversionHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"database_id": Representation{RepType: Required, Create: getEnvSettingWithBlankDefault("converted_to_pdb_db_id")},
	}
)

// issue-routing-tag: database/default
func TestDatabaseDatabasePdbConversionHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabasePdbConversionHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_database_pdb_conversion_history_entries.test_database_pdb_conversion_history_entries"
	singularDatasourceName := "data.oci_database_database_pdb_conversion_history_entry.test_database_pdb_conversion_history_entry"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_database_pdb_conversion_history_entries", "test_database_pdb_conversion_history_entries", Required, Create, databasePdbConversionHistoryEntryDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pdb_conversion_history_entries.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "pdb_conversion_history_entries.0.action"),
				resource.TestCheckResourceAttrSet(datasourceName, "pdb_conversion_history_entries.0.cdb_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "pdb_conversion_history_entries.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pdb_conversion_history_entries.0.source_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pdb_conversion_history_entries.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "pdb_conversion_history_entries.0.target"),
				resource.TestCheckResourceAttrSet(datasourceName, "pdb_conversion_history_entries.0.time_ended"),
				resource.TestCheckResourceAttrSet(datasourceName, "pdb_conversion_history_entries.0.time_started"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_database_pdb_conversion_history_entry", "test_database_pdb_conversion_history_entry", Required, Create, databasePdbConversionHistoryEntrySingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pdb_conversion_history_entry_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cdb_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
			),
		},
	})
}
