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
	DatabaseDatabaseDatabasePdbConversionHistoryEntrySingularDataSourceRepresentation = map[string]interface{}{
		"database_id":                     acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("converted_to_pdb_db_id")},
		"pdb_conversion_history_entry_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("pdb_conversion_history_entry_id")},
	}

	DatabaseDatabaseDatabasePdbConversionHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("converted_to_pdb_db_id")},
	}
)

// issue-routing-tag: database/default
func TestDatabaseDatabasePdbConversionHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabasePdbConversionHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_database_pdb_conversion_history_entries.test_database_pdb_conversion_history_entries"
	singularDatasourceName := "data.oci_database_database_pdb_conversion_history_entry.test_database_pdb_conversion_history_entry"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_pdb_conversion_history_entries", "test_database_pdb_conversion_history_entries", acctest.Required, acctest.Create, DatabaseDatabaseDatabasePdbConversionHistoryEntryDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_pdb_conversion_history_entry", "test_database_pdb_conversion_history_entry", acctest.Required, acctest.Create, DatabaseDatabaseDatabasePdbConversionHistoryEntrySingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
