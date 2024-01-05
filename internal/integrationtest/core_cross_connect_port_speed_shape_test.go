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
	CoreCoreCrossConnectPortSpeedShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	CoreCrossConnectPortSpeedShapeResourceConfig = ""
)

// issue-routing-tag: core/default
func TestCoreCrossConnectPortSpeedShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCrossConnectPortSpeedShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_cross_connect_port_speed_shapes.test_cross_connect_port_speed_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_port_speed_shapes", "test_cross_connect_port_speed_shapes", acctest.Required, acctest.Create, CoreCoreCrossConnectPortSpeedShapeDataSourceRepresentation) +
				compartmentIdVariableStr + CoreCrossConnectPortSpeedShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_port_speed_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_port_speed_shapes.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_port_speed_shapes.0.port_speed_in_gbps"),
			),
		},
	})
}
