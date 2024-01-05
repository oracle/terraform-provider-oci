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
	LicenseManagerLicenseManagerProductLicenseConsumerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"product_license_id":           acctest.Representation{RepType: acctest.Required, Create: ``},
		"is_compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	LicenseManagerProductLicenseConsumerResourceConfig = ""
)

// issue-routing-tag: license_manager/default
func TestLicenseManagerProductLicenseConsumerResource_basic(t *testing.T) {
	t.Skip("The response to this API may take upto 4 hours to populate and there is no work request ID to track it")
	httpreplay.SetScenario("TestLicenseManagerProductLicenseConsumerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_license_manager_product_license_consumers.test_product_license_consumers"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_license_manager_product_license_consumers", "test_product_license_consumers", acctest.Required, acctest.Create, LicenseManagerLicenseManagerProductLicenseConsumerDataSourceRepresentation) +
				compartmentIdVariableStr + LicenseManagerProductLicenseConsumerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "product_license_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.are_all_options_available"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.is_base_license_available"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.license_unit_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.license_units_consumed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.missing_products.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.0.missing_products.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.product_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.resource_compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.resource_compartment_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.resource_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.resource_unit_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.resource_unit_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.resource_id"),
			),
		},
	})
}
