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
	CapacityManagementInternalOccmDemandSignalDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"occ_customer_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.customergroup_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementInternalOccmDemandSignalResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementInternalOccmDemandSignalResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("prod_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	customerGroupId := utils.GetEnvSettingWithBlankDefault("customergroup_id")
	customerGroupIdVariableStr := fmt.Sprintf("variable \"customergroup_id\" { default = \"%s\" }\n", customerGroupId)

	datasourceName := "data.oci_capacity_management_internal_occm_demand_signals.test_internal_occm_demand_signals"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+customerGroupIdVariableStr, "capacitymanagement", "internalOccmDemandSignal", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signals", "test_internal_occm_demand_signals", acctest.Optional, acctest.Create, CapacityManagementInternalOccmDemandSignalDataSourceRepresentation) +
				compartmentIdVariableStr + customerGroupIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "occ_customer_group_id", customerGroupId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
			),
		},
	})
}
