// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	DbSystemPatchHistoryEntryResourceConfig = DbSystemPatchHistoryEntryResourceDependencies + `

`
	DbSystemPatchHistoryEntryPropertyVariables = `

`
	DbSystemPatchHistoryEntryResourceDependencies = DbHomePatchResourceDependencies
)

func TestDatabaseDbSystemPatchHistoryEntryResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_system_patch_history_entries.test_db_system_patch_history_entries"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `
data "oci_database_db_system_patch_history_entries" "test_db_system_patch_history_entries" {
	#Required
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
                ` + compartmentIdVariableStr + DbSystemPatchHistoryEntryResourceConfig,
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
