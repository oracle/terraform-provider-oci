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
	CapacityManagementOccAvailabilityCatalogContentSingularDataSourceRepresentation = map[string]interface{}{
		"occ_availability_catalog_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id}`},
	}

	CapacityManagementOccAvailabilityCatalogContentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Required, acctest.Create, CapacityManagementOccAvailabilityCatalogRepresentation)

	CapacityManagementOccAvailabilityCatalogContentResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog", "test_occ_availability_catalog", acctest.Required, acctest.Create, CapacityManagementOccAvailabilityCatalogRepresentation)
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccAvailabilityCatalogContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccAvailabilityCatalogContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	occCustomerGroupId := utils.GetEnvSettingWithBlankDefault("occ_customer_group_ocid")
	occCustomerGroupIdVariableStr := fmt.Sprintf("variable \"occ_customer_group_id\" { default = \"%s\" }\n", occCustomerGroupId)

	singularDatasourceName := "data.oci_capacity_management_occ_availability_catalog_content.test_occ_availability_catalog_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_availability_catalog_content", "test_occ_availability_catalog_content", acctest.Required, acctest.Create, CapacityManagementOccAvailabilityCatalogContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + occCustomerGroupIdVariableStr + CapacityManagementOccAvailabilityCatalogContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "occ_availability_catalog_id"),
			),
		},
	})
}
