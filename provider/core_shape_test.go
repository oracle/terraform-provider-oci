// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ShapeResourceConfig = ShapeResourceDependencies + `

`
	ShapePropertyVariables = `
variable "shape_availability_domain" { default = "availabilityDomain" }
variable "shape_image_id" { default = "imageId" }

`
	ShapeResourceDependencies = ""
)

func TestCoreShapeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	datasourceName := "data.oci_core_shapes.test_shapes"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `
variable "shape_availability_domain" { default = "availabilityDomain2" }
variable "shape_image_id" { default = "imageId2" }

data "oci_core_shapes" "test_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	#availability_domain = "${var.shape_availability_domain}"
	#image_id = "${var.shape_image_id}"
}
                ` + compartmentIdVariableStr2 + ShapeResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					//resource.TestCheckResourceAttr(datasourceName, "image_id", "imageId2"),

					resource.TestCheckResourceAttrSet(datasourceName, "shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.name"),
				),
			},
		},
	})
}
