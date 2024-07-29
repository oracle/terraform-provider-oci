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
	CapacityManagementOccHandoverResourceBlockDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"handover_date_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2023-04-06T16:10:30.999Z`},
		"handover_date_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `2027-04-06T16:10:30.999Z`},
		"namespace":                              acctest.Representation{RepType: acctest.Optional, Create: `COMPUTE`},
	}

	CapacityManagementOccHandoverResourceBlockResourceConfig = ""
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccHandoverResourceBlockResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccHandoverResourceBlockResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_sp_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_capacity_management_occ_handover_resource_blocks.test_occ_handover_resource_blocks"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_handover_resource_blocks", "test_occ_handover_resource_blocks", acctest.Optional, acctest.Create, CapacityManagementOccHandoverResourceBlockDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "namespace", "COMPUTE"),

				resource.TestCheckResourceAttr(datasourceName, "occ_handover_resource_block_collection.0.items.#", "47"),
			),
		},
	})
}
