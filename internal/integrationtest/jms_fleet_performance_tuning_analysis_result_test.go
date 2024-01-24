// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// before running tests, ensure to set up environment variables used below
	JmsFleetPerformanceTuningAnalysisResultCompartmentId  = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	JmsFleetPerformanceTuningAnalysisResultLogGroupId     = utils.GetEnvSettingWithBlankDefault("fleet_log_group_ocid")
	JmsFleetPerformanceTuningAnalysisResultInventoryLogId = utils.GetEnvSettingWithBlankDefault("fleet_inventory_log_ocid")
	JmsFleetPerformanceTuningAnalysisResultOperationLogId = utils.GetEnvSettingWithBlankDefault("fleet_operation_log_ocid")

	JmsFleetPerformanceTuningAnalysisResultResourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: JmsFleetPerformanceTuningAnalysisResultCompartmentId},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `Created Fleet for Crypto Analysis Result`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Created Fleet for Crypto Analysis Result`},
		"inventory_log": acctest.RepresentationGroup{
			RepType: acctest.Required,
			Group: map[string]interface{}{
				"log_group_id": acctest.Representation{
					RepType: acctest.Required,
					Create:  JmsFleetPerformanceTuningAnalysisResultLogGroupId,
					Update:  JmsFleetPerformanceTuningAnalysisResultLogGroupId,
				},
				"log_id": acctest.Representation{
					RepType: acctest.Required,
					Create:  JmsFleetPerformanceTuningAnalysisResultInventoryLogId,
					Update:  JmsFleetPerformanceTuningAnalysisResultInventoryLogId,
				},
			}},
		"operation_log": acctest.RepresentationGroup{
			RepType: acctest.Optional,
			Group: map[string]interface{}{
				"log_group_id": acctest.Representation{
					RepType: acctest.Required,
					Create:  JmsFleetPerformanceTuningAnalysisResultLogGroupId,
					Update:  JmsFleetPerformanceTuningAnalysisResultLogGroupId,
				},
				"log_id": acctest.Representation{
					RepType: acctest.Required,
					Create:  JmsFleetPerformanceTuningAnalysisResultOperationLogId,
					Update:  JmsFleetPerformanceTuningAnalysisResultOperationLogId,
				},
			}},
	}

	JmsFleetPerformanceTuningAnalysisResultDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetPerformanceTuningAnalysisResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetPerformanceTuningAnalysisResultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_fleet_performance_tuning_analysis_results.test_fleet_performance_tuning_analysis_results"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_fleet",
					"test_fleet",
					acctest.Optional,
					acctest.Create,
					JmsFleetPerformanceTuningAnalysisResultResourceRepresentation,
				) +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_performance_tuning_analysis_results",
					"test_fleet_performance_tuning_analysis_results",
					acctest.Optional,
					acctest.Create,
					JmsFleetPerformanceTuningAnalysisResultDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_result_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "performance_tuning_analysis_result_collection.0.items.#", "0"),
			),
		},
	})
}

// clean up Fleet resource after test
func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("JmsFleetPerformanceTuningAnalysisResult") {
		resource.AddTestSweepers("JmsFleetPerformanceTuningAnalysisResult", &resource.Sweeper{
			Name:         "JmsFleetPerformanceTuningAnalysisResult",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepJmsFleetResource,
		})
	}
}
