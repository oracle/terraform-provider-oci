// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	JmsFleetJavaMigrationAnalysisResultDataSourceRepresentation = map[string]interface{}{
		"fleet_id":            acctest.Representation{RepType: acctest.Required, Create: JmsFleetId},
		"application_name":    acctest.Representation{RepType: acctest.Optional, Create: `dummy-application-name`},
		"host_name":           acctest.Representation{RepType: acctest.Optional, Create: `dummy-host-name`},
		"managed_instance_id": acctest.Representation{RepType: acctest.Optional, Create: JmsManagedInstanceId},
		"time_start":          acctest.Representation{RepType: acctest.Optional, Create: `2024-01-20T15:15:15.000Z`},
		"time_end":            acctest.Representation{RepType: acctest.Optional, Create: `2024-01-20T16:16:16.000Z`},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetJavaMigrationAnalysisResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetJavaMigrationAnalysisResultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_fleet_java_migration_analysis_results.test_fleet_java_migration_analysis_results"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_java_migration_analysis_results",
					"test_fleet_java_migration_analysis_results",
					acctest.Optional,
					acctest.Create,
					JmsFleetJavaMigrationAnalysisResultDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "application_name", `dummy-application-name`),
				resource.TestCheckResourceAttr(datasourceName, "host_name", `dummy-host-name`),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_id", JmsManagedInstanceId),
				resource.TestCheckResourceAttr(datasourceName, "time_start", `2024-01-20T15:15:15.000Z`),
				resource.TestCheckResourceAttr(datasourceName, "time_end", `2024-01-20T16:16:16.000Z`),

				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_result_collection.#"),
				// we can only verify that response contain zero items because we are using dummy test data values
				// we cannot use actual values because it requires setup of fleet -> compute instance -> management agent -> jms plugin.
				resource.TestCheckResourceAttr(datasourceName, "java_migration_analysis_result_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		// note: we cannot write test to verify singular data source because
		// java migration analysis processing requires setup of fleet -> compute instance -> management agent -> jms plugin.
	})
}
