// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	cpeDeviceShapeSingularDataSourceRepresentation = map[string]interface{}{
		"cpe_device_shape_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_cpe_device_shapes.test_cpe_device_shapes.cpe_device_shapes.0.cpe_device_shape_id}`},
	}

	cpeDeviceShapeDataSourceRepresentation = map[string]interface{}{}

	CpeDeviceShapeResourceConfig = ""
)

// issue-routing-tag: core/default
func TestCoreCpeDeviceShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCpeDeviceShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_cpe_device_shapes.test_cpe_device_shapes"
	singularDatasourceName := "data.oci_core_cpe_device_shape.test_cpe_device_shape"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shapes", "test_cpe_device_shapes", acctest.Required, acctest.Create, cpeDeviceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + CpeDeviceShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "cpe_device_shapes.#"),
				resource.TestCheckResourceAttr(datasourceName, "cpe_device_shapes.0.cpe_device_info.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shape", "test_cpe_device_shape", acctest.Required, acctest.Create, cpeDeviceShapeSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cpe_device_shapes", "test_cpe_device_shapes", acctest.Required, acctest.Create, cpeDeviceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + CpeDeviceShapeResourceConfig,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpe_device_shape_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cpe_device_info.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpe_device_shape_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "template"),
			),
		},
	})
}
