// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ImageRequiredOnlyResource = ImageResourceDependencies + `
resource "oci_core_image" "test_image" {
	#Required
	compartment_id = "${var.compartment_id}"
}
`

	ImageResourceConfig = ImageResourceDependencies + `
resource "oci_core_image" "test_image" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.image_display_name}"
	image_source_details {
		#Required
		source_type = "${var.image_image_source_details_source_type}"

		#Optional
		source_image_type = "${var.image_image_source_details_source_image_type}"
	}
	instance_id = "${oci_core_instance.test_instance.id}"
	launch_mode = "${var.image_launch_mode}"
}
`
	ImagePropertyVariables = `
variable "image_display_name" { default = "MyCustomImage" }
variable "image_image_source_details_source_image_type" { default = "sourceImageType" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "launchMode" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_shape" { default = "shape" }
variable "image_state" { default = "state" }

`
	ImageResourceDependencies = "" // Uncomment once defined: InstancePropertyVariables + InstanceResourceConfig
)

func TestCoreImageResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_image.test_image"
	datasourceName := "data.oci_core_images.test_images"

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
				Config:            config + ImagePropertyVariables + compartmentIdVariableStr + ImageRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ImageResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + ImagePropertyVariables + compartmentIdVariableStr + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "sourceImageType"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "launchMode"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "image_display_name" { default = "displayName2" }
variable "image_image_source_details_source_image_type" { default = "sourceImageType" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "launchMode" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_shape" { default = "shape" }
variable "image_state" { default = "state" }

                ` + compartmentIdVariableStr + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "sourceImageType"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "launchMode"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates to Force New parameters.
			{
				Config: config + `
variable "image_display_name" { default = "displayName2" }
variable "image_image_source_details_source_image_type" { default = "sourceImageType2" }
variable "image_image_source_details_source_type" { default = "sourceType2" }
variable "image_launch_mode" { default = "launchMode2" }
variable "image_operating_system" { default = "operatingSystem2" }
variable "image_operating_system_version" { default = "operatingSystemVersion2" }
variable "image_shape" { default = "shape2" }
variable "image_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr2 + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "sourceImageType2"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "sourceType2"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "launchMode2"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
variable "image_display_name" { default = "displayName2" }
variable "image_image_source_details_source_image_type" { default = "sourceImageType2" }
variable "image_image_source_details_source_type" { default = "sourceType2" }
variable "image_launch_mode" { default = "launchMode2" }
variable "image_operating_system" { default = "operatingSystem2" }
variable "image_operating_system_version" { default = "operatingSystemVersion2" }
variable "image_shape" { default = "shape2" }
variable "image_state" { default = "AVAILABLE" }

data "oci_core_images" "test_images" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.image_display_name}"
	operating_system = "${var.image_operating_system}"
	operating_system_version = "${var.image_operating_system_version}"
	shape = "${var.image_shape}"
	state = "${var.image_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_image.test_image.id}"]
    }
}
                ` + compartmentIdVariableStr2 + ImageResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "operating_system", "operatingSystem2"),
					resource.TestCheckResourceAttr(datasourceName, "operating_system_version", "operatingSystemVersion2"),
					resource.TestCheckResourceAttr(datasourceName, "shape", "shape2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "images.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.create_image_allowed"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.launch_mode", "launchMode2"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.operating_system"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.operating_system_version"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.time_created"),
				),
			},
		},
	})
}

func TestCoreImageResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_image.test_image"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + ImagePropertyVariables + compartmentIdVariableStr + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "sourceImageType"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "launchMode"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "image_display_name" { default = "MyCustomImage" }
variable "image_image_source_details_source_image_type" { default = "sourceImageType" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "launchMode" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_shape" { default = "shape" }
variable "image_state" { default = "state" }
				` + compartmentIdVariableStr2 + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "sourceImageType"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "launchMode"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter CompartmentId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "image_display_name" { default = "MyCustomImage" }
variable "image_image_source_details_source_image_type" { default = "sourceImageType2" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "launchMode" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_shape" { default = "shape" }
variable "image_state" { default = "state" }
				` + compartmentIdVariableStr2 + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "sourceImageType2"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "launchMode"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter SourceImageType but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "image_display_name" { default = "MyCustomImage" }
variable "image_image_source_details_source_image_type" { default = "sourceImageType2" }
variable "image_image_source_details_source_type" { default = "sourceType2" }
variable "image_launch_mode" { default = "launchMode" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_shape" { default = "shape" }
variable "image_state" { default = "state" }
				` + compartmentIdVariableStr2 + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "sourceImageType2"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "sourceType2"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "launchMode"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter SourceType but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "image_display_name" { default = "MyCustomImage" }
variable "image_image_source_details_source_image_type" { default = "sourceImageType" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "launchMode" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_shape" { default = "shape" }
variable "image_state" { default = "state" }
				` + compartmentIdVariableStr2 + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "sourceImageType"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "launchMode"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter InstanceId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "image_display_name" { default = "MyCustomImage" }
variable "image_image_source_details_source_image_type" { default = "sourceImageType" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "launchMode2" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_shape" { default = "shape" }
variable "image_state" { default = "state" }
				` + compartmentIdVariableStr2 + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "sourceImageType"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "launchMode2"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(resourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter LaunchMode but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
