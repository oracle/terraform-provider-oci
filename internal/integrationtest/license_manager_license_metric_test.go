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
	LicenseManagerLicenseManagerLicenseMetricSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"is_compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	LicenseManagerLicenseMetricResourceConfig = ""
)

// issue-routing-tag: license_manager/default
func TestLicenseManagerLicenseMetricResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLicenseManagerLicenseMetricResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_license_manager_license_metric.test_license_metric"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_license_manager_license_metric", "test_license_metric", acctest.Required, acctest.Create, LicenseManagerLicenseManagerLicenseMetricSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LicenseManagerLicenseMetricResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "license_record_expiring_soon_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_byol_instance_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_license_included_instance_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_product_license_count"),
			),
		},
	})
}
