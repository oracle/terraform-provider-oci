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
	FleetAppsManagementSchedulerJobJobActivityResourceDataSourceRepresentation = map[string]interface{}{
		"job_activity_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.job_activity_id}`},
		"scheduler_job_id": acctest.Representation{RepType: acctest.Required, Create: `${var.scheduler_job_id}`},
		"resource_task_id": acctest.Representation{RepType: acctest.Optional, Create: `testResourceTaskId`},
		"sequence":         acctest.Representation{RepType: acctest.Optional, Create: `sequence`},
		"step_name":        acctest.Representation{RepType: acctest.Optional, Create: `stepName`},
		"target_name":      acctest.Representation{RepType: acctest.Optional, Create: `targetName`},
	}

	FleetAppsManagementSchedulerJobJobActivityResourceResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementSchedulerJobJobActivityResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementSchedulerJobJobActivityResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Scheduler Jobs and Job Activities are created as a side-effect of creating Schedules,
	// there is no separate Create API, so we need to pass a variable for the id's.
	schedulerJobId := utils.GetEnvSettingWithBlankDefault("scheduler_job_id")
	schedulerJobIdStr := fmt.Sprintf("variable \"scheduler_job_id\" { default = \"%s\" }\n", schedulerJobId)
	jobActivityId := utils.GetEnvSettingWithBlankDefault("job_activity_id")
	jobActivityIdStr := fmt.Sprintf("variable \"job_activity_id\" { default = \"%s\" }\n", jobActivityId)

	datasourceName := "data.oci_fleet_apps_management_scheduler_job_job_activity_resources.test_scheduler_job_job_activity_resources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_scheduler_job_job_activity_resources", "test_scheduler_job_job_activity_resources", acctest.Required, acctest.Create, FleetAppsManagementSchedulerJobJobActivityResourceDataSourceRepresentation) +
				schedulerJobIdStr + jobActivityIdStr + compartmentIdVariableStr + FleetAppsManagementSchedulerJobJobActivityResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "job_activity_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_job_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "resource_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "resource_collection.0.items.#", "1"),
			),
		},
	})
}
