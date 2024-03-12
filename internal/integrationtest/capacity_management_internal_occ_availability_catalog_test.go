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
	CapacityManagementInternalOccAvailabilityCatalogDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"namespace":             acctest.Representation{RepType: acctest.Optional, Create: `COMPUTE`},
		"occ_customer_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.occ_customer_group_id}`},
	}
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementInternalOccAvailabilityCatalogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementInternalOccAvailabilityCatalogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	occCustomerGroupId := utils.GetEnvSettingWithBlankDefault("occ_customer_group_ocid")
	occCustomerGroupIdVariableStr := fmt.Sprintf("variable \"occ_customer_group_id\" { default = \"%s\" }\n", occCustomerGroupId)

	datasourceName := "data.oci_capacity_management_internal_occ_availability_catalogs.test_internal_occ_availability_catalogs"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_internal_occ_availability_catalogs", "test_internal_occ_availability_catalogs", acctest.Required, acctest.Create, CapacityManagementInternalOccAvailabilityCatalogDataSourceRepresentation) +
				compartmentIdVariableStr + occCustomerGroupIdVariableStr, //+ CapacityManagementInternalOccAvailabilityCatalogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "occ_availability_catalog_collection.0.items.0.namespace", "COMPUTE"),
				resource.TestCheckResourceAttrSet(datasourceName, "occ_customer_group_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "occ_availability_catalog_collection.#"),
			),
		},
	})
}
