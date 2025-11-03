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
	FleetAppsManagementRecommendedPatchDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"patch_id":       acctest.Representation{RepType: acctest.Optional, Create: `26541`},
		"patch_level":    acctest.Representation{RepType: acctest.Optional, Create: `LATEST`},
		"patch_type":     acctest.Representation{RepType: acctest.Optional, Create: `Critical Patch Update`},
		"severity":       acctest.Representation{RepType: acctest.Optional, Create: `CRITICAL`},
		"target_id":      acctest.Representation{RepType: acctest.Optional, Create: `301`},
		"target_name":    acctest.Representation{RepType: acctest.Optional, Create: `/home/oracle/Oracle/Middleware/Oracle_Home/wlserver`},
	}
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementRecommendedPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementRecommendedPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid_1")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_fleet_apps_management_recommended_patches.test_recommended_patches"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_recommended_patches", "test_recommended_patches", acctest.Required, acctest.Create, FleetAppsManagementRecommendedPatchDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "recommended_patch_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "recommended_patch_collection.0.items.#", "0"),
			),
		},
	})
}
