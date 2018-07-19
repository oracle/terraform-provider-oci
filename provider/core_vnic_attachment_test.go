// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	VnicAttachmentRequiredOnlyResource = VnicAttachmentResourceDependencies + `
resource "oci_core_vnic_attachment" "test_vnic_attachment" {
	#Required
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.t.id}"
	}
	instance_id = "${oci_core_instance.t.id}"
}
`

	VnicAttachmentResourceConfig = VnicAttachmentResourceDependencies + `
resource "oci_core_vnic_attachment" "test_vnic_attachment" {
	#Required
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.t.id}"

		#Optional
		assign_public_ip = "${var.vnic_attachment_create_vnic_details_assign_public_ip}"
		defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.vnic_attachment_create_vnic_details_defined_tags_value}")}"
		display_name = "${var.vnic_attachment_create_vnic_details_display_name}"
		freeform_tags = "${var.vnic_attachment_create_vnic_details_freeform_tags}"
		hostname_label = "${var.vnic_attachment_create_vnic_details_hostname_label}"
		private_ip = "${var.vnic_attachment_create_vnic_details_private_ip}"
		skip_source_dest_check = "${var.vnic_attachment_create_vnic_details_skip_source_dest_check}"
	}
	instance_id = "${oci_core_instance.t.id}"

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
variable "vnic_attachment_create_vnic_details_freeform_tags" { default = {"Department"= "Accounting"} }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "attachvnictestinstance" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "10.0.1.5" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = "0" }

`
	VnicAttachmentResourceDependencies = instanceDnsConfig + VnicPropertyVariables + VnicResourceConfig
)

func TestCoreVnicAttachmentResource_basic(t *testing.T) {
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
		CheckDestroy: testAccCheckCoreVnicAttachmentDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + VnicAttachmentPropertyVariables + compartmentIdVariableStr + VnicAttachmentRequiredOnlyResource,
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
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", "attachvnictestinstance"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.1.5"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.skip_source_dest_check", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "nic_index", "0"),
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
variable "vnic_attachment_create_vnic_details_freeform_tags" { default = {"Department"= "Accounting"} }
variable "vnic_attachment_create_vnic_details_hostname_label" { default = "attachvnictestinstance" }
variable "vnic_attachment_create_vnic_details_private_ip" { default = "10.0.1.5" }
variable "vnic_attachment_create_vnic_details_skip_source_dest_check" { default = false }
variable "vnic_attachment_display_name" { default = "displayName" }
variable "vnic_attachment_nic_index" { default = "0" }

data "oci_core_vnic_attachments" "test_vnic_attachments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	instance_id = "${oci_core_instance.t.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_vnic_attachment.test_vnic_attachment.id}"]
    }
}
                ` + compartmentIdVariableStr + VnicAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),

					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.0.display_name", "displayName"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "vnic_attachments.0.nic_index", "0"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_attachments.0.vnic_id"),
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

func testAccCheckCoreVnicAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_vnic_attachment" {
			noResourceFound = false
			request := oci_core.GetVnicAttachmentRequest{}

			tmp := rs.Primary.ID
			request.VnicAttachmentId = &tmp

			response, error := client.GetVnicAttachment(context.Background(), request)

			if error == nil {
				return fmt.Errorf("resource still exists")
			}
			//Verify that exception is for 'not found'.
			if response.RawResponse.StatusCode != 404 {
				return error
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
