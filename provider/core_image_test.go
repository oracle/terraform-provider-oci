// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	ImageRequiredOnlyResource = ImageResourceDependencies + `
resource "oci_core_image" "test_image" {
	#Required
	compartment_id = "${var.compartment_id}"
	instance_id = "${oci_core_instance.test_instance.id}"
	timeouts {
		create = "30m"
	}
}
`

	ImageResourceConfig = ImageResourceDependencies + `
data "oci_objectstorage_namespace" "t" {
}
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
		namespace_name = "${data.oci_objectstorage_namespace.t.namespace}"
		bucket_name = "test-artifacts"
		object_name = "test-image-export"
	}
	launch_mode = "${var.image_launch_mode}"
	timeouts {
		create = "30m"
	}
}
`
	ImagePropertyVariables = `
variable "image_defined_tags_value" { default = "value" }
variable "image_display_name" { default = "MyCustomImage" }
variable "image_freeform_tags" { default = {"Department"= "Finance"} }
variable "image_image_source_details_source_image_type" { default = "QCOW2" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "NATIVE" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_state" { default = "AVAILABLE" }

`
	ImageResourceDependencies = DefinedTagsDependencies + InstancePropertyVariables + InstanceResourceAsDependencyConfig
)

func TestCoreImageResource_basic(t *testing.T) {
	t.Skip("Long running test")
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
		CheckDestroy: testAccCheckCoreImageDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + ImagePropertyVariables + compartmentIdVariableStr + ImageRequiredOnlyResource,
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
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "QCOW2"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "NATIVE"),
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
variable "image_image_source_details_source_image_type" { default = "QCOW2" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "NATIVE" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
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
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_image_type", "QCOW2"),
					resource.TestCheckResourceAttr(resourceName, "image_source_details.0.source_type", "objectStorageTuple"),
					resource.TestCheckResourceAttr(resourceName, "launch_mode", "NATIVE"),
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
variable "image_image_source_details_source_image_type" { default = "QCOW2" }
variable "image_image_source_details_source_type" { default = "objectStorageTuple" }
variable "image_launch_mode" { default = "NATIVE" }
variable "image_operating_system" { default = "operatingSystem" }
variable "image_operating_system_version" { default = "operatingSystemVersion" }
variable "image_state" { default = "AVAILABLE" }

data "oci_core_images" "test_images" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.image_display_name}"
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
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "images.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.create_image_allowed"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "images.0.launch_mode", "NATIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.operating_system"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.operating_system_version"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "images.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ResourceName:      resourceName,
			},
		},
	})
}

func testAccCheckCoreImageDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_image" {
			noResourceFound = false
			request := oci_core.GetImageRequest{}

			tmp := rs.Primary.ID
			request.ImageId = &tmp

			_, err := client.GetImage(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}
			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
