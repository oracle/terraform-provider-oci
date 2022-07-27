// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseDatabaseDbHomePatchHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
	}

	DatabaseDbHomePatchHistoryEntryResourceConfig = DbSystemResourceConfig
)

// issue-routing-tag: database/default
func TestDatabaseDbHomePatchHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbHomePatchHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_home_patch_history_entries.test_db_home_patch_history_entries"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_home_patch_history_entries", "test_db_home_patch_history_entries", acctest.Required, acctest.Create, DatabaseDatabaseDbHomePatchHistoryEntryDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDbHomePatchHistoryEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_home_id"),
				resource.TestCheckResourceAttr(datasourceName, "patch_history_entries.#", "0"),
			),
		},
	})
}
