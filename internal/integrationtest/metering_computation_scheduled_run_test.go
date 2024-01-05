// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MeteringComputationMeteringComputationScheduledRunSingularDataSourceRepresentation = map[string]interface{}{
		"scheduled_run_id": acctest.Representation{RepType: acctest.Required, Create: `${var.scheduled_run_ocid}`},
	}

	MeteringComputationMeteringComputationScheduledRunDataSourceRepresentation = map[string]interface{}{
		"schedule_id": acctest.Representation{RepType: acctest.Required, Create: `${var.schedule_ocid}`},
	}
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationScheduledRunResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationScheduledRunResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	scheduledRunId := utils.GetEnvSettingWithBlankDefault("scheduled_run_id")
	scheduledRunIdVariableStr := fmt.Sprintf("variable \"scheduled_run_ocid\" { default = \"%s\" }\n", scheduledRunId)

	scheduleId := utils.GetEnvSettingWithBlankDefault("schedule_id")
	scheduleIdVariableStr := fmt.Sprintf("variable \"schedule_ocid\" { default = \"%s\" }\n", scheduleId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_metering_computation_scheduled_runs.test_scheduled_runs"
	singularDatasourceName := "data.oci_metering_computation_scheduled_run.test_scheduled_run"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + scheduleIdVariableStr + scheduledRunIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_scheduled_runs", "test_scheduled_runs", acctest.Required, acctest.Create, MeteringComputationMeteringComputationScheduledRunDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "schedule_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "scheduled_run_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + scheduleIdVariableStr + scheduledRunIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_scheduled_run", "test_scheduled_run", acctest.Required, acctest.Create, MeteringComputationMeteringComputationScheduledRunSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_run_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
			),
		},
	})
}
