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
	dbHomePatchDataSourceRepresentation = map[string]interface{}{
		"db_home_id": Representation{repType: Required, create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
	}

	DbHomePatchResourceConfig = DbSystemResourceConfig
)

func TestDatabaseDbHomePatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbHomePatchResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_home_patches.test_db_home_patches"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_db_home_patches", "test_db_home_patches", Required, Create, dbHomePatchDataSourceRepresentation) +
					compartmentIdVariableStr + DbHomePatchResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "db_home_id"),

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
