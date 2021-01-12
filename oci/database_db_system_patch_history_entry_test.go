// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	dbSystemPatchHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"db_system_id": Representation{repType: Required, create: `${oci_database_db_system.test_db_system.id}`},
	}

	DbSystemPatchHistoryEntryResourceConfig = DbSystemResourceConfig
)

func TestDatabaseDbSystemPatchHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbSystemPatchHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_system_patch_history_entries.test_db_system_patch_history_entries"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_db_system_patch_history_entries", "test_db_system_patch_history_entries", Required, Create, dbSystemPatchHistoryEntryDataSourceRepresentation) +
					compartmentIdVariableStr + DbSystemPatchHistoryEntryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.patch_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.time_started"),
				),
			},
		},
	})
}
