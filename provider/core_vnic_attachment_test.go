// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	VnicAttachmentRequiredOnlyResource = VnicAttachmentResourceDependencies + `
resource "oci_core_vnic_attachment" "test_vnic_attachment" {
	#Required
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.test_subnet.id}"
	}
	instance_id = "${oci_core_instance.test_instance.id}"
}
`

	VnicAttachmentResourceConfig = VnicAttachmentResourceDependencies + `
resource "oci_core_vnic_attachment" "test_vnic_attachment" {
	#Required
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.test_subnet.id}"

		#Optional
		assign_public_ip = "${var.vnic_attachment_create_vnic_details_assign_public_ip}"
		defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.vnic_attachment_create_vnic_details_defined_tags_value}")}"
		display_name = "${var.vnic_attachment_create_vnic_details_display_name}"
		freeform_tags = "${var.vnic_attachment_create_vnic_details_freeform_tags}"
		hostname_label = "${var.vnic_attachment_create_vnic_details_hostname_label}"
		private_ip = "${var.vnic_attachment_create_vnic_details_private_ip}"
		skip_source_dest_check = "${var.vnic_attachment_create_vnic_details_skip_source_dest_check}"
	}
	instance_id = "${oci_core_instance.test_instance.id}"

	#Optional
	display_name = "${var.vnic_attachment_display_name}"
	nic_index = "${var.vnic_attachment_nic_index}"
}
`
	VnicAttachmentPropertyVariables = `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = false }
variable "vnic_attachment_create_vnic_details_defined_tags_value" { default = "definedTags" }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName" }
variable "vnic_attachment_create_vnic_details_freeform_tags" { default = "freeformTags" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = 10 }

`
	VnicAttachmentResourceDependencies = VnicPropertyVariables + VnicResourceConfig
)

func TestCoreVnicAttachmentResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_vnic_attachment.test_vnic_attachment"
	datasourceName := "data.oci_core_vnic_attachments.test_vnic_attachments"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + VnicAttachmentPropertyVariables + compartmentIdVariableStr + VnicAttachmentRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VnicAttachmentResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + VnicAttachmentPropertyVariables + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "privateIp"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "vnic_attachment_availability_domain" { default = "availabilityDomain" }
variable "vnic_attachment_create_vnic_details_assign_public_ip" { default = false }
variable "vnic_attachment_create_vnic_details_defined_tags_value" { default = "definedTags" }
variable "vnic_attachment_create_vnic_details_display_name" { default = "displayName" }
variable "vnic_attachment_create_vnic_details_freeform_tags" { default = "freeformTags" }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "hostnameLabel" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "privateIp" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = 10 }

data "oci_core_vnic_attachments" "test_vnic_attachments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.vnic_attachment_availability_domain}"
	instance_id = "${oci_core_instance.test_instance.id}"
	vnic_id = "${oci_core_vnic.test_vnic.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_vnic_attachment.test_vnic_attachment.id}"]
    }
}
                ` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "availability_domain", "availabilityDomain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_id"),

					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.0.display_name", "displayName"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.0.nic_index", "10"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.vnic_id"),
				),
			},
		},
	})
}
