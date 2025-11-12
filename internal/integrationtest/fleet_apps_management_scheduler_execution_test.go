// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	schedulerDefinitionId                                         = utils.GetEnvSettingWithBlankDefault("scheduler_definition_id")
	FleetAppsManagementSchedulerExecutionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `execution-tersi-schedule-1`},
		// "lifecycle_operation":                     acctest.Representation{RepType: acctest.Optional, Create: `PROVISION`},
		// "resource_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_resource.test_resource.id}`},
		"resource_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.fleet_id}`},
		// "runbook_id":                              acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_runbook.test_runbook.id}`},
		"runbook_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.runbook_id}`},
		"runbook_version_name": acctest.Representation{RepType: acctest.Optional, Create: `1`},
		// "scheduler_defintion_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_scheduler_defintion.test_scheduler_defintion.id}`},
		"scheduler_defintion_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.schedular_definition_id}`},
		// "scheduler_job_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_migration_job.test_job.id}`},
		"scheduler_job_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.schedular_job_id}`},
		"substate":         acctest.Representation{RepType: acctest.Optional, Create: `FAILED`},
		"time_scheduled_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2025-10-20T00:00:00.000Z`},
		"time_scheduled_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2025-10-24T00:00:00.000Z`},
	}

	FleetAppsManagementSchedulerExecutionResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementSchedulerExecutionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementSchedulerExecutionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	fleetId := utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	fleetIdStr := fmt.Sprintf("variable \"fleet_id\" { default = \"%s\" }\n", fleetId)

	runbookId := utils.GetEnvSettingWithBlankDefault("test_runbook_ocid")
	runbokIdStr := fmt.Sprintf("variable \"runbook_id\" { default = \"%s\" }\n", runbookId)

	scedularDefinitionId := utils.GetEnvSettingWithBlankDefault("schedular_definition")
	scedularDefinitionIdStr := fmt.Sprintf("variable \"schedular_definition_id\" { default = \"%s\" }\n", scedularDefinitionId)

	scedularJobId := utils.GetEnvSettingWithBlankDefault("schedular_job_id")
	scedularJobIdStr := fmt.Sprintf("variable \"schedular_job_id\" { default = \"%s\" }\n", scedularJobId)

	datasourceName := "data.oci_fleet_apps_management_scheduler_executions.test_scheduler_executions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + fleetIdStr + runbokIdStr + scedularDefinitionIdStr + scedularJobIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_scheduler_executions", "test_scheduler_executions", acctest.Optional, acctest.Create, FleetAppsManagementSchedulerExecutionDataSourceRepresentation) +
				FleetAppsManagementSchedulerExecutionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "execution-tersi-schedule-1"),
				// resource.TestCheckResourceAttrSet(datasourceName, "lifecycle_operation"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "runbook_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "runbook_version_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_defintion_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_scheduled_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_scheduled_less_than"),

				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_execution_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_execution_collection.0.items.#"),
			),
		},
	})
}
