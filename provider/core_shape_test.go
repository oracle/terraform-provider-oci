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

`
	ShapeResourceDependencies = ""
)

func TestCoreShapeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_shapes.test_shapes"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `
variable "shape_availability_domain" { default = "availabilityDomain" }

data "oci_core_shapes" "test_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	#availability_domain = "${var.shape_availability_domain}"
	#image_id = "${var.shape_image_id}"
}
                ` + compartmentIdVariableStr + ShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					//resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttrSet(datasourceName, "image_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.name"),
				),
			},
		},
	})
}
