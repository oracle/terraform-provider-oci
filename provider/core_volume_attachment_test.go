// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	VolumeAttachmentRequiredOnlyResource = VolumeAttachmentResourceDependencies + `
resource "oci_core_volume_attachment" "test_volume_attachment" {
	#Required
	attachment_type = "${var.volume_attachment_attachment_type}"
	instance_id = "${oci_core_instance.test_instance.id}"
	volume_id = "${oci_core_volume.test_volume.id}"
}
`

	VolumeAttachmentResourceConfig = VolumeAttachmentResourceDependencies + `
resource "oci_core_volume_attachment" "test_volume_attachment" {
	#Required
	attachment_type = "${var.volume_attachment_attachment_type}"
	instance_id = "${oci_core_instance.test_instance.id}"
	volume_id = "${oci_core_volume.test_volume.id}"

	#Optional
	display_name = "${var.volume_attachment_display_name}"
	is_read_only = "${var.volume_attachment_is_read_only}"
}
`
	VolumeAttachmentPropertyVariables = `
variable "volume_attachment_attachment_type" { default = "attachmentType" }
variable "volume_attachment_availability_domain" { default = "availabilityDomain" }
variable "volume_attachment_display_name" { default = "displayName" }
variable "volume_attachment_is_read_only" { default = false }

`
	VolumeAttachmentResourceDependencies = "" // Uncomment once defined: InstancePropertyVariables + InstanceResourceConfig + VolumePropertyVariables + VolumeResourceConfig
)

func TestCoreVolumeAttachmentResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_volume_attachment.test_volume_attachment"
	datasourceName := "data.oci_core_volume_attachments.test_volume_attachments"

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
				Config:            config + VolumeAttachmentPropertyVariables + compartmentIdVariableStr + VolumeAttachmentRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "attachmentType"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeAttachmentResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + VolumeAttachmentPropertyVariables + compartmentIdVariableStr + VolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "attachmentType"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `
variable "volume_attachment_attachment_type" { default = "attachmentType2" }
variable "volume_attachment_availability_domain" { default = "availabilityDomain2" }
variable "volume_attachment_display_name" { default = "displayName2" }
variable "volume_attachment_is_read_only" { default = true }

                ` + compartmentIdVariableStr2 + VolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "attachmentType2"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

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
variable "volume_attachment_attachment_type" { default = "attachmentType2" }
variable "volume_attachment_availability_domain" { default = "availabilityDomain2" }
variable "volume_attachment_display_name" { default = "displayName2" }
variable "volume_attachment_is_read_only" { default = true }

data "oci_core_volume_attachments" "test_volume_attachments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.volume_attachment_availability_domain}"
	instance_id = "${oci_core_instance.test_instance.id}"
	volume_id = "${oci_core_volume.test_volume.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_volume_attachment.test_volume_attachment.id}"]
    }
}
                ` + compartmentIdVariableStr2 + VolumeAttachmentResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain2"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_id"),

					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.attachment_type", "attachmentType2"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.is_read_only", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.volume_id"),
				),
			},
		},
	})
}

func TestCoreVolumeAttachmentResource_forcenew(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_attachment.test_volume_attachment"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + VolumeAttachmentPropertyVariables + compartmentIdVariableStr + VolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "attachmentType"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "volume_attachment_attachment_type" { default = "attachmentType2" }
variable "volume_attachment_availability_domain" { default = "availabilityDomain" }
variable "volume_attachment_display_name" { default = "displayName" }
variable "volume_attachment_is_read_only" { default = false }
				` + compartmentIdVariableStr + VolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "attachmentType2"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter AttachmentType but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "volume_attachment_attachment_type" { default = "attachmentType2" }
variable "volume_attachment_availability_domain" { default = "availabilityDomain" }
variable "volume_attachment_display_name" { default = "displayName2" }
variable "volume_attachment_is_read_only" { default = false }
				` + compartmentIdVariableStr + VolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "attachmentType2"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter DisplayName but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "volume_attachment_attachment_type" { default = "attachmentType2" }
variable "volume_attachment_availability_domain" { default = "availabilityDomain" }
variable "volume_attachment_display_name" { default = "displayName2" }
variable "volume_attachment_is_read_only" { default = false }
				` + compartmentIdVariableStr + VolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "attachmentType2"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

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
variable "volume_attachment_attachment_type" { default = "attachmentType2" }
variable "volume_attachment_availability_domain" { default = "availabilityDomain" }
variable "volume_attachment_display_name" { default = "displayName2" }
variable "volume_attachment_is_read_only" { default = true }
				` + compartmentIdVariableStr + VolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "attachmentType2"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter IsReadOnly but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "volume_attachment_attachment_type" { default = "attachmentType2" }
variable "volume_attachment_availability_domain" { default = "availabilityDomain" }
variable "volume_attachment_display_name" { default = "displayName2" }
variable "volume_attachment_is_read_only" { default = true }
				` + compartmentIdVariableStr + VolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "attachmentType2"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter VolumeId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
