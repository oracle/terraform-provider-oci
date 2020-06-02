// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	imageShapeSingularDataSourceRepresentation = map[string]interface{}{
		"image_id":   Representation{repType: Required, create: `${var.FlexInstanceImageOCID[var.region]}`},
		"shape_name": Representation{repType: Required, create: `VM.Standard.E3.Flex`},
	}

	imageShapeDataSourceRepresentation = map[string]interface{}{
		"image_id": Representation{repType: Required, create: `${var.FlexInstanceImageOCID[var.region]}`},
	}

	ImageShapeResourceConfig = FlexVmImageIdsVariable
)

func TestCoreImageShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreImageShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_image_shapes.test_image_shapes"
	singularDatasourceName := "data.oci_core_image_shape.test_image_shape"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_image_shapes", "test_image_shapes", Required, Create, imageShapeDataSourceRepresentation) +
					compartmentIdVariableStr + ImageShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "image_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "image_shape_compatibilities.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "image_shape_compatibilities.0.image_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "image_shape_compatibilities.0.shape"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_image_shape", "test_image_shape", Required, Create, imageShapeSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ImageShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_name"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
				),
			},
		},
	})
}
