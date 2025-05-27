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
	testFleetId                                                = utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	testInstanceId                                             = utils.GetEnvSettingWithBlankDefault("test_instance_id")
	FleetAppsManagementInventoryRecordDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"fleet_id":                  acctest.Representation{RepType: acctest.Optional, Create: testFleetId},
		"resource_id":               acctest.Representation{RepType: acctest.Optional, Create: testInstanceId},
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

	datasourceName := "data.oci_fleet_apps_management_inventory_records.test_inventory_records"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_inventory_records", "test_inventory_records", acctest.Optional, acctest.Create, FleetAppsManagementInventoryRecordDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementInventoryRecordResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "inventory_record_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "inventory_record_collection.0.items.#", "1"),
			),
		},
	})
}
