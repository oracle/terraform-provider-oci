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
	dbSystemPatchDataSourceRepresentation = map[string]interface{}{
		"db_system_id": Representation{repType: Required, create: `${oci_database_db_system.test_db_system.id}`},
	}

	DbSystemPatchResourceConfig = DbSystemResourceConfig
)

func TestDatabaseDbSystemPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbSystemPatchResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_system_patches.test_db_system_patches"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_db_system_patches", "test_db_system_patches", Required, Create, dbSystemPatchDataSourceRepresentation) +
					compartmentIdVariableStr + DbSystemPatchResourceConfig,
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
