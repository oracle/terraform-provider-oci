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
	FleetAppsManagementInventoryRecordDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"fleet_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${var.fleet_id}`},
		"is_details_required":       acctest.Representation{RepType: acctest.Optional, Create: `false`},
		// "resource_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_resource.test_resource.id}`},
		"resource_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.instance_id}`},
	}

	FleetAppsManagementInventoryRecordResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementInventoryRecordResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementInventoryRecordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	fleetId := utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	fleetIdStr := fmt.Sprintf("variable \"fleet_id\" { default = \"%s\" }\n", fleetId)

	testInstanceId := utils.GetEnvSettingWithBlankDefault("self_hosted_instance_id")
	testInstanceIdStr := fmt.Sprintf("variable \"instance_id\" { default = \"%s\" }\n", testInstanceId)

	datasourceName := "data.oci_fleet_apps_management_inventory_records.test_inventory_records"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_inventory_records", "test_inventory_records", acctest.Optional, acctest.Create, FleetAppsManagementInventoryRecordDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementInventoryRecordResourceConfig + fleetIdStr + testInstanceIdStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "is_details_required", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "inventory_record_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "inventory_record_collection.0.items.#"),
			),
		},
	})
}
