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
	CapacityManagementInternalOccmDemandSignalItemDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"occ_customer_group_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.customergroup_id}`},
		"demand_signal_namespace": acctest.Representation{RepType: acctest.Optional, Create: `COMPUTE`},
		"occm_demand_signal_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.demand_signal_id}`},
		"resource_name":           acctest.Representation{RepType: acctest.Optional, Create: `BM.Standard.E5.192`},
	}
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementInternalOccmDemandSignalItemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementInternalOccmDemandSignalItemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("prod_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	customerGroupId := utils.GetEnvSettingWithBlankDefault("customergroup_id")
	customerGroupIdVariableStr := fmt.Sprintf("variable \"customergroup_id\" { default = \"%s\" }\n", customerGroupId)

	demandSignalId := utils.GetEnvSettingWithBlankDefault("demand_signal_id")
	demandSignalIdVariableStr := fmt.Sprintf("variable \"demand_signal_id\" { default = \"%s\" }\n", demandSignalId)

	datasourceName := "data.oci_capacity_management_internal_occm_demand_signal_items.test_internal_occm_demand_signal_items"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+customerGroupIdVariableStr+demandSignalIdVariableStr, "capacitymanagement", "internalOccmDemandSignalItem", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_items", "test_internal_occm_demand_signal_items", acctest.Optional, acctest.Create, CapacityManagementInternalOccmDemandSignalItemDataSourceRepresentation) +
				compartmentIdVariableStr + customerGroupIdVariableStr + demandSignalIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "occ_customer_group_id", customerGroupId),
				resource.TestCheckResourceAttr(datasourceName, "demand_signal_namespace", "COMPUTE"),
			),
		},
	})
}
