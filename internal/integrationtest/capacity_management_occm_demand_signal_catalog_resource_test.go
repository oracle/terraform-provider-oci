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
	CapacityManagementOccmDemandSignalCatalogResourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"demand_signal_namespace": acctest.Representation{RepType: acctest.Optional, Create: `COMPUTE`},
		"name":                    acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	CapacityManagementOccmDemandSignalCatalogResourceResourceConfig = ""
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccmDemandSignalCatalogResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccmDemandSignalCatalogResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("sp_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_capacity_management_occm_demand_signal_catalog_resources.test_occm_demand_signal_catalog_resources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_catalog_resources", "test_occm_demand_signal_catalog_resources", acctest.Required, acctest.Create, CapacityManagementOccmDemandSignalCatalogResourceDataSourceRepresentation) +
				compartmentIdVariableStr + CapacityManagementOccmDemandSignalCatalogResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "occm_demand_signal_catalog_resource_collection.0.items.0.namespace", "GPU"),
				resource.TestCheckResourceAttr(datasourceName, "occm_demand_signal_catalog_resource_collection.0.items.0.name", "BM.GPU.H200.8 (GPU: 8xH200)"),

				resource.TestCheckResourceAttrSet(datasourceName, "occm_demand_signal_catalog_resource_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "occm_demand_signal_catalog_resource_collection.0.items.#", "36"),
			),
		},
	})
}
