// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// before running tests, ensure to set up environment variables used below
	JmsFleetCryptoAnalysisResultFleetId       = utils.GetEnvSettingWithBlankDefault("fleet_ocid")
	JmsFleetCryptoAnalysisResultCompartmentId = utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	JmsFleetCryptoAnalysisResultDummyManagedInstanceId = utils.GetEnvSettingWithBlankDefault("managed_instance_ocid")

	JmsFleetCryptoAnalysisResultDataSourceRepresentation = map[string]interface{}{
		"fleet_id":                                 acctest.Representation{RepType: acctest.Required, Create: JmsFleetCryptoAnalysisResultFleetId},
		"aggregation_mode":                         acctest.Representation{RepType: acctest.Optional, Create: `JFR`},
		"finding_count":                            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"finding_count_greater_than":               acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"host_name":                                acctest.Representation{RepType: acctest.Optional, Create: `dummy-host-name`},
		"managed_instance_id":                      acctest.Representation{RepType: acctest.Optional, Create: JmsFleetCryptoAnalysisResultDummyManagedInstanceId},
		"non_compliant_finding_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"non_compliant_finding_count_greater_than": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"time_start":                               acctest.Representation{RepType: acctest.Optional, Create: `2024-01-20T15:15:15.000Z`},
		"time_end":                                 acctest.Representation{RepType: acctest.Optional, Create: `2024-01-20T16:16:16.000Z`},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetCryptoAnalysisResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetCryptoAnalysisResultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_fleet_crypto_analysis_results.test_fleet_crypto_analysis_results"

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_crypto_analysis_results",
					"test_fleet_crypto_analysis_results",
					acctest.Optional,
					acctest.Create,
					JmsFleetCryptoAnalysisResultDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "aggregation_mode", "JFR"),
				resource.TestCheckResourceAttr(datasourceName, "finding_count", `10`),
				resource.TestCheckResourceAttr(datasourceName, "finding_count_greater_than", `10`),
				resource.TestCheckResourceAttr(datasourceName, "host_name", `dummy-host-name`),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_id", JmsFleetCryptoAnalysisResultDummyManagedInstanceId),
				resource.TestCheckResourceAttr(datasourceName, "non_compliant_finding_count", `10`),
				resource.TestCheckResourceAttr(datasourceName, "non_compliant_finding_count_greater_than", `10`),
				resource.TestCheckResourceAttr(datasourceName, "time_start", `2024-01-20T15:15:15.000Z`),
				resource.TestCheckResourceAttr(datasourceName, "time_end", `2024-01-20T16:16:16.000Z`),

				resource.TestCheckResourceAttrSet(datasourceName, "crypto_analysis_result_collection.#"),
				// we can only verify that response contain zero items because we are using dummy test data values
				// we cannot use actual values because it requires setup of fleet -> compute instance -> management agent -> jms plugin.
				resource.TestCheckResourceAttr(datasourceName, "crypto_analysis_result_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		// note: we cannot write test to verify singular data source because
		// crypto analysis processing requires setup of fleet -> compute instance -> management agent -> jms plugin.
	})
}

// clean up Fleet resource after test
func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("JmsFleetCryptoAnalysisResult") {
		resource.AddTestSweepers("JmsFleetCryptoAnalysisResult", &resource.Sweeper{
			Name:         "JmsFleetCryptoAnalysisResult",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepJmsFleetResource,
		})
	}
}
