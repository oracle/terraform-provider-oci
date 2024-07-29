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
	CapacityManagementInternalOccHandoverResourceBlockDetailDataSourceRepresentation = map[string]interface{}{
		"occ_handover_resource_block_id": acctest.Representation{RepType: acctest.Required, Create: `${var.occ_handover_resource_block_id}`},
		"page":                           acctest.Representation{RepType: acctest.Optional, Create: `null`},
	}

	CapacityManagementInternalOccHandoverResourceBlockDetailResourceConfig = ""
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementInternalOccHandoverResourceBlockDetailResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementInternalOccHandoverResourceBlockDetailResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	occHandoverResourceBlockId := utils.GetEnvSettingWithBlankDefault("occ_handover_resource_block_id")
	occHandoverResourceBlockIdVariableStr := fmt.Sprintf("variable \"occ_handover_resource_block_id\" { default = \"%s\" }\n", occHandoverResourceBlockId)

	datasourceName := "data.oci_capacity_management_internal_occ_handover_resource_block_details.test_internal_occ_handover_resource_block_details"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + occHandoverResourceBlockIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_internal_occ_handover_resource_block_details", "test_internal_occ_handover_resource_block_details", acctest.Required, acctest.Create, CapacityManagementInternalOccHandoverResourceBlockDetailDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "page"),
				resource.TestCheckResourceAttr(datasourceName, "occ_handover_resource_block_id", occHandoverResourceBlockId),
				resource.TestCheckResourceAttr(datasourceName, "namespace", "COMPUTE"),
				resource.TestCheckResourceAttr(datasourceName, "occ_handover_resource_block_detail_collection.0.items.#", "1"),
			),
		},
	})
}
