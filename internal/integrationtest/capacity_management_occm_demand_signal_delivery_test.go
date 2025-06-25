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
	CapacityManagementOccmDemandSignalDeliveryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"occm_demand_signal_item_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.demandsignalitem_id}`},
	}

	CapacityManagementOccmDemandSignalDeliveryResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal", "test_occm_demand_signal", acctest.Required, acctest.Create, CapacityManagementOccmDemandSignalRepresentation)

	CapacityManagementOccmDemandSignalDeliveryResourceConfig = CapacityManagementOccmDemandSignalDeliveryResourceDependencies
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccmDemandSignalDeliveryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccmDemandSignalDeliveryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("sp_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	demandsignalitemId := utils.GetEnvSettingWithBlankDefault("demand_signal_item_id")
	demandsignalitemIdVariableStr := fmt.Sprintf("variable \"demandsignalitem_id\" { default = \"%s\" }\n", demandsignalitemId)

	datasourceName := "data.oci_capacity_management_occm_demand_signal_deliveries.test_occm_demand_signal_deliveries"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_deliveries", "test_occm_demand_signal_deliveries", acctest.Required, acctest.Create, CapacityManagementOccmDemandSignalDeliveryDataSourceRepresentation) +
				compartmentIdVariableStr + demandsignalitemIdVariableStr + CapacityManagementOccmDemandSignalDeliveryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				// Removed id check as it's not applicable without specific resource creation
				// resource.TestCheckResourceAttrSet(datasourceName, "occm_demand_signal_item_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "occm_demand_signal_delivery_collection.#"),
				// resource.TestCheckResourceAttr(datasourceName, "occm_demand_signal_delivery_collection.0.items.#", "48"),
			),
		},
	})
}
