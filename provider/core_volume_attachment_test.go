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

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_attachment.test_volume_attachment"
	datasourceName := "data.oci_core_volume_attachments.test_volume_attachments"

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
				),
			},

			// verify datasource
			{
				Config: config + `
variable "volume_attachment_attachment_type" { default = "attachmentType" }
variable "volume_attachment_availability_domain" { default = "availabilityDomain" }
variable "volume_attachment_display_name" { default = "displayName" }
variable "volume_attachment_is_read_only" { default = false }

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
                ` + compartmentIdVariableStr + VolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_id"),

					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.attachment_type", "attachmentType"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.display_name", "displayName"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_attachments.0.is_read_only", "false"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_attachments.0.volume_id"),
				),
			},
		},
	})
}
