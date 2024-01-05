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
	LicenseManagerLicenseManagerTopUtilizedResourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"is_compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"resource_unit_type":           acctest.Representation{RepType: acctest.Optional, Create: `OCPU`},
	}

	LicenseManagerTopUtilizedResourceResourceConfig = ""
)

// issue-routing-tag: license_manager/default
func TestLicenseManagerTopUtilizedResourceResource_basic(t *testing.T) {
	t.Skip("The response to this API may take upto 4 hours to populate and there is no work request ID to track it")
	httpreplay.SetScenario("TestLicenseManagerTopUtilizedResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_license_manager_top_utilized_resources.test_top_utilized_resources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_license_manager_top_utilized_resources", "test_top_utilized_resources", acctest.Required, acctest.Create, LicenseManagerLicenseManagerTopUtilizedResourceDataSourceRepresentation) +
				compartmentIdVariableStr + LicenseManagerTopUtilizedResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "items.#", "4"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.resource_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.total_units"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.unit_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.resource_compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.resource_compartment_name"),
			),
		},
	})
}
