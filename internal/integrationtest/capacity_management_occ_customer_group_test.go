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
	CapacityManagementOccCustomerGroupSingularDataSourceRepresentation = map[string]interface{}{
		"occ_customer_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.occ_customer_group_id}`},
	}

	CapacityManagementOccCustomerGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.occ_customer_group_id}`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
	}

	CapacityManagementOccCustomerGroupResourceConfig = ""
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccCustomerGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccCustomerGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	occCustomerGroupId := utils.GetEnvSettingWithBlankDefault("occ_customer_group_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	occCustomerGroupIdVariableStr := fmt.Sprintf("variable \"occ_customer_group_id\" { default = \"%s\" }\n", occCustomerGroupId)

	datasourceName := "data.oci_capacity_management_occ_customer_groups.test_occ_customer_groups"
	singularDatasourceName := "data.oci_capacity_management_occ_customer_group.test_occ_customer_group"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_customer_groups", "test_occ_customer_groups", acctest.Required, acctest.Create, CapacityManagementOccCustomerGroupDataSourceRepresentation) +
				compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccCustomerGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "occ_customer_group_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_customer_group", "test_occ_customer_group", acctest.Required, acctest.Create, CapacityManagementOccCustomerGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccCustomerGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "occ_customer_group_id", occCustomerGroupId),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "customers_list.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
