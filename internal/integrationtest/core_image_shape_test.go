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
	CoreCoreImageShapeSingularDataSourceRepresentation = map[string]interface{}{
		"image_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
		"shape_name": acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E3.Flex`},
	}

	CoreCoreImageShapeDataSourceRepresentation = map[string]interface{}{
		"image_id": acctest.Representation{RepType: acctest.Required, Create: `${var.FlexInstanceImageOCID[var.region]}`},
	}

	CoreImageShapeResourceConfig = utils.FlexVmImageIdsVariable
)

// issue-routing-tag: core/computeImaging
func TestCoreImageShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreImageShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_image_shapes.test_image_shapes"
	singularDatasourceName := "data.oci_core_image_shape.test_image_shape"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_image_shapes", "test_image_shapes", acctest.Required, acctest.Create, CoreCoreImageShapeDataSourceRepresentation) +
				compartmentIdVariableStr + CoreImageShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "image_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "image_shape_compatibilities.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "image_shape_compatibilities.0.image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "image_shape_compatibilities.0.shape"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_image_shape", "test_image_shape", acctest.Required, acctest.Create, CoreCoreImageShapeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreImageShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_name"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
			),
		},
	})
}
