// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	BootVolumeAttachmentResourceConfig = BootVolumeAttachmentResourceDependencies + `

`

	BootVolumeAttachmentResourceDependencies = BootVolumeResourceDependencies
)

func TestCoreBootVolumeAttachmentResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_boot_volume_attachments.test_boot_volume_attachments"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource can retrieve a specific attachment using server-side filtering
			{
				Config: config + `
data "oci_core_boot_volume_attachments" "test_boot_volume_attachments" {
	#Required
	availability_domain = "${oci_core_instance.test_instance.availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	boot_volume_id = "${oci_core_instance.test_instance.boot_volume_id}"
	instance_id = "${oci_core_instance.test_instance.id}"
}
                ` + compartmentIdVariableStr + BootVolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),

					resource.TestCheckResourceAttr(datasourceName, "boot_volume_attachments.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_attachments.0.availability_domain"),
					TestCheckResourceAttributesEqual(datasourceName, "boot_volume_attachments.0.boot_volume_id", "oci_core_instance.test_instance", "boot_volume_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_attachments.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_attachments.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_attachments.0.id"),
					TestCheckResourceAttributesEqual(datasourceName, "boot_volume_attachments.0.instance_id", "oci_core_instance.test_instance", "id"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_attachments.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_attachments.0.time_created"),
				),
			},
			// verify datasource can retrieve all boot volume attachments in a compartment by specifying no filtering options
			{
				Config: config + `
data "oci_core_boot_volume_attachments" "test_boot_volume_attachments" {
	#Required
	availability_domain = "${oci_core_instance.test_instance.availability_domain}"
	compartment_id = "${var.compartment_id}"
}
                ` + compartmentIdVariableStr + BootVolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestMatchResourceAttr(datasourceName, "boot_volume_attachments.#", regexp.MustCompile("[1-9][0-9]*")),
				),
			},
		},
	})
}
