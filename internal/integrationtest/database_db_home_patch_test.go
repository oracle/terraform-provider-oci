// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	dbHomePatchDataSourceRepresentation = map[string]interface{}{
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
	}

	DbHomePatchResourceConfig = DbSystemResourceConfig
)

// issue-routing-tag: database/default
func TestDatabaseDbHomePatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbHomePatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_home_patches.test_db_home_patches"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_home_patches", "test_db_home_patches", acctest.Required, acctest.Create, dbHomePatchDataSourceRepresentation) +
				compartmentIdVariableStr + DbHomePatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_home_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "patches.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.time_released"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.version"),
			),
		},
	})
}
