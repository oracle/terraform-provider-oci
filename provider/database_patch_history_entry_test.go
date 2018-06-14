// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	PatchHistoryEntryResourceConfig = PatchResourceDependencies
)

func TestDatabasePatchHistoryEntryResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_patch_history_entries.test_patch_history_entries"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `
data "oci_database_patch_history_entries" "test_patch_history_entries" {
	#Required
	db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
}
                ` + compartmentIdVariableStr + PatchHistoryEntryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "db_home_id"),
					resource.TestCheckResourceAttr(datasourceName, "patch_history_entries.#", "0"),
				),
			},
		},
	})
}
