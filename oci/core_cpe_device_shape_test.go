// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	cpeDeviceShapeSingularDataSourceRepresentation = map[string]interface{}{
		"cpe_device_shape_id": Representation{RepType: Required, Create: `${data.oci_core_cpe_device_shapes.test_cpe_device_shapes.cpe_device_shapes.0.cpe_device_shape_id}`},
	}

	cpeDeviceShapeDataSourceRepresentation = map[string]interface{}{}

	CpeDeviceShapeResourceConfig = ""
)

// issue-routing-tag: core/default
func TestCoreCpeDeviceShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCpeDeviceShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_cpe_device_shapes.test_cpe_device_shapes"
	singularDatasourceName := "data.oci_core_cpe_device_shape.test_cpe_device_shape"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shapes", "test_cpe_device_shapes", Required, Create, cpeDeviceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + CpeDeviceShapeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "cpe_device_shapes.#"),
				resource.TestCheckResourceAttr(datasourceName, "cpe_device_shapes.0.cpe_device_info.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shape", "test_cpe_device_shape", Required, Create, cpeDeviceShapeSingularDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shapes", "test_cpe_device_shapes", Required, Create, cpeDeviceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + CpeDeviceShapeResourceConfig,

			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpe_device_shape_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cpe_device_info.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpe_device_shape_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "template"),
			),
		},
	})
}
