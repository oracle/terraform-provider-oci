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
	dbSystemComputePerformanceDataSourceRepresentation = map[string]interface{}{
		"db_system_shape": acctest.Representation{RepType: acctest.Required, Create: `BM.DenseIO2.52`},
	}
)

// issue-routing-tag: database/default
func TestDatabaseDbSystemComputePerformanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbSystemComputePerformanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	datasourceName := "data.oci_database_db_system_compute_performances.test_db_system_compute_performances"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_system_compute_performances", "test_db_system_compute_performances", acctest.Required, acctest.Create, dbSystemComputePerformanceDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "db_system_compute_performances.0.shape", "BM.DenseIO2.52"),

				resource.TestCheckResourceAttrSet(datasourceName, "db_system_compute_performances.#"),
				resource.TestCheckResourceAttr(datasourceName, "db_system_compute_performances.0.compute_performance_list.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_compute_performances.0.shape"),
			),
		},
	})
}
