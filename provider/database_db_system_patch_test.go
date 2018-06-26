// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	DbSystemPatchResourceConfig = DbSystemPatchResourceDependencies + `

`
	DbSystemPatchPropertyVariables = `

`
	DbSystemPatchResourceDependencies = DbHomePatchResourceDependencies
)

func TestDatabaseDbSystemPatchResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_system_patches.test_db_system_patches"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `
data "oci_database_db_system_patches" "test_db_system_patches" {
	#Required
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
                ` + compartmentIdVariableStr + DbSystemPatchResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "patches.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.time_released"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.version"),
				),
			},
		},
	})
}
