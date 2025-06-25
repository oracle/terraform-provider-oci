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
	CapacityManagementInternalOccmDemandSignalCatalogSingularDataSourceRepresentation = map[string]interface{}{
		"occm_demand_signal_catalog_id": acctest.Representation{RepType: acctest.Required, Create: `{}`},
	}

	CapacityManagementInternalOccmDemandSignalCatalogDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"occ_customer_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.customergroup_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementInternalOccmDemandSignalCatalogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementInternalOccmDemandSignalCatalogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("prod_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	customerGroupId := utils.GetEnvSettingWithBlankDefault("customergroup_id")
	customerGroupIdVariableStr := fmt.Sprintf("variable \"customergroup_id\" { default = \"%s\" }\n", customerGroupId)

	datasourceName := "data.oci_capacity_management_internal_occm_demand_signal_catalogs.test_internal_occm_demand_signal_catalogs"
	// singularDatasourceName := "data.oci_capacity_management_internal_occm_demand_signal_catalog.test_internal_occm_demand_signal_catalog"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_catalogs", "test_internal_occm_demand_signal_catalogs", acctest.Required, acctest.Create, CapacityManagementInternalOccmDemandSignalCatalogDataSourceRepresentation) +
				compartmentIdVariableStr + customerGroupIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "occ_customer_group_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "occm_demand_signal_catalog_collection.#"),
			),
		},
		// verify singular datasource
		//{
		//	Config: config +
		//		acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_catalog", "test_internal_occm_demand_signal_catalog", acctest.Required, acctest.Create, CapacityManagementInternalOccmDemandSignalCatalogSingularDataSourceRepresentation) +
		//		compartmentIdVariableStr + customerGroupIdVariableStr + CapacityManagementInternalOccmDemandSignalCatalogResourceConfig,
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "occm_demand_signal_catalog_id"),
		//
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
		//	),
		//},
	})
}
