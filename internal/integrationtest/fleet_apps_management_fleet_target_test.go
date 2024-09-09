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
	FleetAppsManagementFleetTargetDataSourceRepresentation = map[string]interface{}{
		"fleet_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.test_active_fleet}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `/u01/app/oracle/product/jdk`},
		"product":               acctest.Representation{RepType: acctest.Optional, Create: `Oracle Java`},
		"resource_display_name": acctest.Representation{RepType: acctest.Optional, Create: `weblogic-1`},
	}

	FleetAppsManagementFleetTargetResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementFleetTargetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementFleetTargetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Fleet in ACTIVE state. Fleets require a confirmation action call not supported by Terraform to go active.
	// Thus, this needs to be created and confirmed manually.
	activeFleetId := utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	activeFleetStr := fmt.Sprintf("variable \"test_active_fleet\" { default = \"%s\" }\n", activeFleetId)
	datasourceName := "data.oci_fleet_apps_management_fleet_targets.test_fleet_targets"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet_targets", "test_fleet_targets", acctest.Optional, acctest.Create, FleetAppsManagementFleetTargetDataSourceRepresentation) +
				activeFleetStr + compartmentIdVariableStr + FleetAppsManagementFleetTargetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "/u01/app/oracle/product/jdk"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "product", "Oracle Java"),
				resource.TestCheckResourceAttr(datasourceName, "resource_display_name", "weblogic-1"),

				resource.TestCheckResourceAttrSet(datasourceName, "fleet_target_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_target_collection.0.items.#", "1"),
			),
		},
	})
}
