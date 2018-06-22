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
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.image_defined_tags_value}")}"
	display_name = "${var.image_display_name}"
	freeform_tags = "${var.image_freeform_tags}"
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
variable "image_defined_tags_value" { default = "value" }
variable "image_display_name" { default = "MyCustomImage" }
variable "image_freeform_tags" { default = {"Department"= "Finance"} }
variable "image_image_source_details_source_image_type" { default = "sourceImageType" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "launchMode" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_shape" { default = "shape" }
variable "image_state" { default = "AVAILABLE" }

`
	ImageResourceDependencies = DefinedTagsDependencies // Uncomment once defined: InstancePropertyVariables + InstanceResourceConfig
)

func TestCoreImageResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

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
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCustomImage"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
variable "image_defined_tags_value" { default = "updatedValue" }
variable "image_display_name" { default = "displayName2" }
variable "image_freeform_tags" { default = {"Department"= "Accounting"} }
variable "image_image_source_details_source_image_type" { default = "sourceImageType" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "launchMode" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_shape" { default = "shape" }
variable "image_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "create_image_allowed"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
			// verify datasource
			{
				Config: config + `
variable "image_defined_tags_value" { default = "updatedValue" }
variable "image_display_name" { default = "displayName2" }
variable "image_freeform_tags" { default = {"Department"= "Accounting"} }
variable "image_image_source_details_source_image_type" { default = "sourceImageType" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "launchMode" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_shape" { default = "shape" }
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
                ` + compartmentIdVariableStr + ImageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "operating_system", "operatingSystem"),
					resource.TestCheckResourceAttr(datasourceName, "operating_system_version", "operatingSystemVersion"),
					resource.TestCheckResourceAttr(datasourceName, "shape", "shape"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "images.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.create_image_allowed"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.launch_mode", "launchMode"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.operating_system"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.operating_system_version"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.time_created"),
				),
			},
		},
	})
}
