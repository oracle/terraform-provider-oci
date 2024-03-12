// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	CapacityManagementOccAvailabilityCatalogOccAvailabilityDataSourceRepresentation = map[string]interface{}{
		"occ_availability_catalog_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id}`},
		"date_expected_capacity_handover": acctest.Representation{RepType: acctest.Optional, Create: `dateExpectedCapacityHandover`},
		"resource_name":                   acctest.Representation{RepType: acctest.Optional, Create: `test`},
		"resource_type":                   acctest.Representation{RepType: acctest.Optional, Create: `SERVER_HW`},
		"workload_type":                   acctest.Representation{RepType: acctest.Optional, Create: `GENERIC`},
	}

	CapacityManagementOccAvailabilityCatalogOccAvailabilityResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Required, acctest.Create, CapacityManagementOccAvailabilityCatalogRepresentation)
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccAvailabilityCatalogOccAvailabilityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccAvailabilityCatalogOccAvailabilityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	occCustomerGroupId := utils.GetEnvSettingWithBlankDefault("occ_customer_group_ocid")
	occCustomerGroupIdVariableStr := fmt.Sprintf("variable \"occ_customer_group_id\" { default = \"%s\" }\n", occCustomerGroupId)

	datasourceName := "data.oci_capacity_management_occ_availability_catalog_occ_availabilities.test_occ_availability_catalog_occ_availabilities"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog_occ_availabilities", "test_occ_availability_catalog_occ_availabilities", acctest.Required, acctest.Create, CapacityManagementOccAvailabilityCatalogOccAvailabilityDataSourceRepresentation) +
				compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccAvailabilityCatalogOccAvailabilityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "occ_availability_collection.0.items.0.resource_name"),
				resource.TestCheckResourceAttr(datasourceName, "occ_availability_collection.0.items.0.resource_type", "CAPACITY_CONSTRAINT"),
				resource.TestCheckResourceAttr(datasourceName, "occ_availability_collection.0.items.0.workload_type", "US_PROD"),

				resource.TestCheckResourceAttrSet(datasourceName, "occ_availability_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "occ_availability_collection.0.items.#", "24"),
			),
		},
	})
}
