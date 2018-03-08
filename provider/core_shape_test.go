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
resource "oci_core_shape" "test_shape" {
}
`
	ShapePropertyVariables = `
variable "shape_availability_domain" { default = "availabilityDomain" }
variable "shape_image_id" { default = "imageId" }

`
	ShapeResourceDependencies = ""
)

func TestCoreShapeResource_basic(t *testing.T) {
	t.Skip("Creating shape resource is not supported. Data source test is covered by legacy test. Need to enable data source only tests in generator. https://jira.aka.lgl.grungy.us/browse/ORCH-708")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_shape.test_shape"
	datasourceName := "data.oci_core_shapes.test_shapes"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + ShapePropertyVariables + compartmentIdVariableStr + ShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `
variable "shape_availability_domain" { default = "availabilityDomain2" }
variable "shape_image_id" { default = "imageId2" }

                ` + compartmentIdVariableStr2 + ShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "name"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "shape_availability_domain" { default = "availabilityDomain2" }
variable "shape_image_id" { default = "imageId2" }

data "oci_core_shapes" "test_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.shape_availability_domain}"
	image_id = "${var.shape_image_id}"

    filter {
    	name = "id"
    	values = ["${oci_core_shape.test_shape.id}"]
    }
}
                ` + compartmentIdVariableStr2 + ShapeResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "image_id", "imageId2"),

					resource.TestCheckResourceAttr(datasourceName, "shapes.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.name"),
				),
			},
		},
	})
}
