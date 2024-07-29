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
	CapacityManagementInternalOccHandoverResourceBlockRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":                              acctest.Representation{RepType: acctest.Required, Create: `COMPUTE`},
		"occ_customer_group_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.occ_customer_group_id}`},
		"handover_date_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2023-04-06T16:10:30.999Z`},
		"handover_date_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `2027-04-06T16:10:30.999Z`},
	}

	CapacityManagementInternalOccHandoverResourceBlockResourceConfig = ""
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementInternalOccHandoverResourceBlockResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementInternalOccHandoverResourceBlockResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	occCustomerGroupId := utils.GetEnvSettingWithBlankDefault("occ_customer_group_id")
	occCustomerGroupIdVariableStr := fmt.Sprintf("variable \"occ_customer_group_id\" { default = \"%s\" }\n", occCustomerGroupId)

	datasourceName := "data.oci_capacity_management_internal_occ_handover_resource_blocks.test_internal_occ_handover_resource_blocks"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + occCustomerGroupIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_internal_occ_handover_resource_blocks", "test_internal_occ_handover_resource_blocks", acctest.Optional, acctest.Create, CapacityManagementInternalOccHandoverResourceBlockRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "namespace", "COMPUTE"),
				resource.TestCheckResourceAttr(datasourceName, "occ_customer_group_id", occCustomerGroupId),
				// Update the below items list size if the test fails
				resource.TestCheckResourceAttr(datasourceName, "occ_handover_resource_block_collection.0.items.#", "47"),
			),
		},
	})
}
