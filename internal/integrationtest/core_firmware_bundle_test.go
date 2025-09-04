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
	CoreFirmwareBundleSingularDataSourceRepresentation = map[string]interface{}{
		"firmware_bundle_id": acctest.Representation{RepType: acctest.Required, Create: `${var.firmware_bundle_id}`},
	}

	CoreFirmwareBundleDataSourceRepresentation = map[string]interface{}{
		"platform":          acctest.Representation{RepType: acctest.Required, Create: `Reg_Comp_A1-2C_Server.01`},
		"compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"is_default_bundle": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle_state":   acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}
)

// issue-routing-tag: core/firmwareRepository
func TestCoreFirmwareBundleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreFirmwareBundleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	firmwareBundleId := utils.GetEnvSettingWithBlankDefault("firmware_bundle_id")
	firmwareBundleIdVariableStr := fmt.Sprintf("variable \"firmware_bundle_id\" { default = \"%s\" }\n", firmwareBundleId)

	datasourceName := "data.oci_core_firmware_bundles.test_firmware_bundles"
	singularDatasourceName := "data.oci_core_firmware_bundle.test_firmware_bundle"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_firmware_bundles", "test_firmware_bundles", acctest.Optional, acctest.Create, CoreFirmwareBundleDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "is_default_bundle", "false"),
				resource.TestCheckResourceAttr(datasourceName, "platform", "Reg_Comp_A1-2C_Server.01"),
				resource.TestCheckResourceAttr(datasourceName, "firmware_bundles_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "firmware_bundles_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + firmwareBundleIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_firmware_bundle", "test_firmware_bundle", acctest.Required, acctest.Create, CoreFirmwareBundleSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "firmware_bundle_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "platforms.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
