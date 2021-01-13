// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	crossConnectPortSpeedShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	CrossConnectPortSpeedShapeResourceConfig = ""
)

func TestCoreCrossConnectPortSpeedShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCrossConnectPortSpeedShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_cross_connect_port_speed_shapes.test_cross_connect_port_speed_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_cross_connect_port_speed_shapes", "test_cross_connect_port_speed_shapes", Required, Create, crossConnectPortSpeedShapeDataSourceRepresentation) +
					compartmentIdVariableStr + CrossConnectPortSpeedShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_port_speed_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_port_speed_shapes.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_port_speed_shapes.0.port_speed_in_gbps"),
				),
			},
		},
	})
}
