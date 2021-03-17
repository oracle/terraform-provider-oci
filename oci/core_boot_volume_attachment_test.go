// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	bootVolumeAttachmentDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"boot_volume_id":      Representation{repType: Optional, create: `${oci_core_instance.test_instance.boot_volume_id}`},
		"instance_id":         Representation{repType: Optional, create: `${oci_core_instance.test_instance.id}`},
	}

	BootVolumeAttachmentResourceConfig = generateResourceFromRepresentationMap("oci_core_boot_volume", "test_boot_volume", Required, Create, bootVolumeRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		AvailabilityDomainConfig
)

func TestCoreBootVolumeAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreBootVolumeAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_boot_volume_attachments.test_boot_volume_attachments"

	saveConfigContent("", "", "", t)

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
