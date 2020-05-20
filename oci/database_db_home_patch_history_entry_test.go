// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	dbHomePatchHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"db_home_id": Representation{repType: Required, create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
	}

	DbHomePatchHistoryEntryResourceConfig = DbSystemResourceConfig
)

func TestDatabaseDbHomePatchHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbHomePatchHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_home_patch_history_entries.test_db_home_patch_history_entries"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_db_home_patch_history_entries", "test_db_home_patch_history_entries", Required, Create, dbHomePatchHistoryEntryDataSourceRepresentation) +
					compartmentIdVariableStr + DbHomePatchHistoryEntryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "db_home_id"),
					resource.TestCheckResourceAttr(datasourceName, "patch_history_entries.#", "0"),
				),
			},
		},
	})
}
