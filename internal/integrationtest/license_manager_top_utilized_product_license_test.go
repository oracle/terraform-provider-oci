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
	LicenseManagerLicenseManagerTopUtilizedProductLicenseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"is_compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	LicenseManagerTopUtilizedProductLicenseResourceConfig = ""
)

// issue-routing-tag: license_manager/default
func TestLicenseManagerTopUtilizedProductLicenseResource_basic(t *testing.T) {
	t.Skip("The response to this API may take upto 4 hours to populate and there is no work request ID to track it")
	httpreplay.SetScenario("TestLicenseManagerTopUtilizedProductLicenseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_license_manager_top_utilized_product_licenses.test_top_utilized_product_licenses"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_license_manager_top_utilized_product_licenses", "test_top_utilized_product_licenses", acctest.Required, acctest.Create, LicenseManagerLicenseManagerTopUtilizedProductLicenseDataSourceRepresentation) +
				compartmentIdVariableStr + LicenseManagerTopUtilizedProductLicenseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "items.#", "3"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.is_unlimited"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.product_license_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.product_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.total_license_unit_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.total_units_consumed"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.unit_type"),
			),
		},
	})
}
