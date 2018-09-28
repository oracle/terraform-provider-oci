// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	bootVolumeAttachmentDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${oci_core_instance.test_instance.availability_domain}`, update: `${oci_core_instance.test_instance.availability_domain}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"boot_volume_id":      Representation{repType: Optional, create: `${oci_core_instance.test_instance.boot_volume_id}`},
		"instance_id":         Representation{repType: Optional, create: `${oci_core_instance.test_instance.id}`},
	}

	BootVolumeAttachmentResourceConfig = BootVolumeResourceConfig
)

func TestCoreBootVolumeAttachmentResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_boot_volume_attachments.test_boot_volume_attachments"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource can retrieve a specific attachment using server-side filtering
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_boot_volume_attachments", "test_boot_volume_attachments", Optional, Create, bootVolumeAttachmentDataSourceRepresentation) +
					compartmentIdVariableStr + BootVolumeAttachmentResourceConfig,
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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_boot_volume_attachments", "test_boot_volume_attachments", Required, Update, bootVolumeAttachmentDataSourceRepresentation) +
					compartmentIdVariableStr + BootVolumeAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestMatchResourceAttr(datasourceName, "boot_volume_attachments.#", regexp.MustCompile("[1-9][0-9]*")),
				),
			},
		},
	})
}
