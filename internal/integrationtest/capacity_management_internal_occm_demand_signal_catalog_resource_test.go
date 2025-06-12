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
	CapacityManagementInternalOccmDemandSignalCatalogResourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"occ_customer_group_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.customergroup_id}`},
		"occm_demand_signal_catalog_id": acctest.Representation{RepType: acctest.Required, Create: `${var.occm_demand_signal_catalog_id}`},
		"demand_signal_namespace":       acctest.Representation{RepType: acctest.Optional, Create: `COMPUTE`},
		"name":                          acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementInternalOccmDemandSignalCatalogResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementInternalOccmDemandSignalCatalogResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("prod_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	customerGroupId := utils.GetEnvSettingWithBlankDefault("customergroup_id")
	customerGroupIdVariableStr := fmt.Sprintf("variable \"customergroup_id\" { default = \"%s\" }\n", customerGroupId)

	catalogId := utils.GetEnvSettingWithBlankDefault("occm_demand_signal_catalog_id")
	catalogIdVariableStr := fmt.Sprintf("variable \"occm_demand_signal_catalog_id\" { default = \"%s\" }\n", catalogId)

	datasourceName := "data.oci_capacity_management_internal_occm_demand_signal_catalog_resources.test_internal_occm_demand_signal_catalog_resources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_catalog_resources", "test_internal_occm_demand_signal_catalog_resources", acctest.Required, acctest.Create, CapacityManagementInternalOccmDemandSignalCatalogResourceDataSourceRepresentation) +
				compartmentIdVariableStr + customerGroupIdVariableStr + catalogIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttr(datasourceName, "namespace", "COMPUTE"),
				// resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(datasourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "occm_demand_signal_catalog_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "internal_occm_demand_signal_catalog_resource_collection.#"),
				// resource.TestCheckResourceAttr(datasourceName, "internal_occm_demand_signal_catalog_resource_collection.0.items.#", "1"),
			),
		},
	})
}
