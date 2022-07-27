// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	DatabaseDatabaseDbSystemPatchHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system.id}`},
	}

	DatabaseDbSystemPatchHistoryEntryResourceConfig = DbSystemResourceConfig
)

// issue-routing-tag: database/default
func TestDatabaseDbSystemPatchHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbSystemPatchHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_system_patch_history_entries.test_db_system_patch_history_entries"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_system_patch_history_entries", "test_db_system_patch_history_entries", acctest.Required, acctest.Create, DatabaseDatabaseDbSystemPatchHistoryEntryDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDbSystemPatchHistoryEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.patch_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.time_started"),
			),
		},
	})
}
