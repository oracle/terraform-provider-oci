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
	FleetAppsManagementFleetProductDataSourceRepresentation = map[string]interface{}{
		"fleet_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.test_active_fleet}`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_ocid}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `Oracle Java`},
	}

	FleetAppsManagementFleetProductResourceConfig = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementFleetProductResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementFleetProductResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Fleet in ACTIVE state. Fleets require a confirmation action call not supported by Terraform to go active.
	// Thus, this needs to be created and confirmed manually.
	activeFleetId := utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	activeFleetStr := fmt.Sprintf("variable \"test_active_fleet\" { default = \"%s\" }\n", activeFleetId)
	datasourceName := "data.oci_fleet_apps_management_fleet_products.test_fleet_products"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet_products", "test_fleet_products", acctest.Optional, acctest.Create, FleetAppsManagementFleetProductDataSourceRepresentation) +
				activeFleetStr + compartmentIdVariableStr + FleetAppsManagementFleetProductResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Oracle Java"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "fleet_product_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_product_collection.0.items.#", "1"),
			),
		},
	})
}
