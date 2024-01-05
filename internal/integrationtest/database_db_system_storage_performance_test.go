// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	dbSystemStoragePerformanceDataSourceRepresentation = map[string]interface{}{
		"storage_management": acctest.Representation{RepType: acctest.Required, Create: `ASM`},
		"shape_type":         acctest.Representation{RepType: acctest.Required, Create: `AMD`},
	}
)

// issue-routing-tag: database/default
func TestDatabaseDbSystemStoragePerformanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbSystemStoragePerformanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_database_db_system_storage_performances.test_db_system_storage_performances"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_system_storage_performances", "test_db_system_storage_performances", acctest.Required, acctest.Create, dbSystemStoragePerformanceDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_storage_performances.#"),
				resource.TestCheckResourceAttr(datasourceName, "db_system_storage_performances.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_storage_performances.0.data_storage_performance_list.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_storage_performances.0.reco_storage_performance_list.#"),
				resource.TestCheckResourceAttr(datasourceName, "db_system_storage_performances.0.shape_type", "AMD"),
			),
		},
	})
}
