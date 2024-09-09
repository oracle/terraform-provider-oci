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
	FleetAppsManagementSchedulerDefinitionScheduledFleetDataSourceRepresentation = map[string]interface{}{
		"scheduler_definition_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_scheduler_definition.test_scheduler_definition.id}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `testWeblogicFleet`},
	}

	FleetAppsManagementSchedulerDefinitionScheduledFleetResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Required, acctest.Create, FleetAppsManagementMaintenanceWindowRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Required, acctest.Create, FleetAppsManagementSchedulerDefinitionRepresentation)
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementSchedulerDefinitionScheduledFleetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementSchedulerDefinitionScheduledFleetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Runbooks are currently created by Oracle, and read-only. There is no Create API.
	runbookId := utils.GetEnvSettingWithBlankDefault("test_runbook_ocid")
	testRunbookStr := fmt.Sprintf("variable \"test_runbook_ocid\" { default = \"%s\" }\n", runbookId)

	// Fleet in ACTIVE state. Fleets require a confirmation action call not supported by Terraform to go active.
	// Thus, this needs to be created and confirmed manually.
	activeFleetId := utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	activeFleetStr := fmt.Sprintf("variable \"test_active_fleet\" { default = \"%s\" }\n", activeFleetId)
	datasourceName := "data.oci_fleet_apps_management_scheduler_definition_scheduled_fleets.test_scheduler_definition_scheduled_fleets"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition_scheduled_fleets", "test_scheduler_definition_scheduled_fleets", acctest.Optional, acctest.Create, FleetAppsManagementSchedulerDefinitionScheduledFleetDataSourceRepresentation) +
				testRunbookStr + activeFleetStr + compartmentIdVariableStr + FleetAppsManagementSchedulerDefinitionScheduledFleetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "testWeblogicFleet"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_definition_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "scheduled_fleet_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "scheduled_fleet_collection.0.items.#", "1"),
			),
		},
	})
}
