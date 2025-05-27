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
		"compartment_id":                          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                            acctest.Representation{RepType: acctest.Optional, Create: `changed-execution-1`},
		"runbook_version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
		"scheduler_defintion_id":                  acctest.Representation{RepType: acctest.Required, Create: schedulerDefinitionId},
		"time_scheduled_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2025-01-01T00:00:00Z`},
		"time_scheduled_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2029-12-31T00:00:00Z`},
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

	datasourceName := "data.oci_fleet_apps_management_scheduler_executions.test_scheduler_executions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_scheduler_executions", "test_scheduler_executions", acctest.Optional, acctest.Create, FleetAppsManagementSchedulerExecutionDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementSchedulerExecutionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "changed-execution-1"),
				resource.TestCheckResourceAttrSet(datasourceName, "runbook_version_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_defintion_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_scheduled_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_scheduled_less_than"),

				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_execution_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "scheduler_execution_collection.0.items.#", "1"),
			),
		},
	})
}
