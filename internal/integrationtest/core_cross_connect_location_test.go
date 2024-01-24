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
	CoreCoreCrossConnectLocationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	CoreCrossConnectLocationResourceConfig = ""
)

// issue-routing-tag: core/default
func TestCoreCrossConnectLocationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCrossConnectLocationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_cross_connect_locations.test_cross_connect_locations"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_locations", "test_cross_connect_locations", acctest.Required, acctest.Create, CoreCoreCrossConnectLocationDataSourceRepresentation) +
				compartmentIdVariableStr + CoreCrossConnectLocationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_locations.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_locations.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_locations.0.name"),
			),
		},
	})
}
