// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreImageTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreImageTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + instanceConfig
	s.ResourceName = "oci_core_image.t"
}

func (s *ResourceCoreImageTestSuite) TestAccResourceCoreImage_basic() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// create image
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_image" "t" {
						compartment_id = "${var.tenancy_ocid}"
						instance_id = "${oci_core_instance.t.id}"
						timeouts {
							create = "30m"
						}
					}
					data "oci_core_images" "t" {
						compartment_id = "${var.tenancy_ocid}"
					
						filter {
							name = "id"
							values = ["${oci_core_image.t.id}"]
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "base_image_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_image_allowed", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_mode", "NATIVE"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.boot_volume_type", "ISCSI"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.firmware", "UEFI_64"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.network_type", "VFIO"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.remote_data_volume_type", "PARAVIRTUALIZED"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "size_in_mbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.ImageLifecycleStateAvailable)),
					func(ts *terraform.State) (err error) {
						resId, err = fromInstanceState(ts, s.ResourceName, "id")
						return err
					},

					resource.TestCheckResourceAttrSet("data.oci_core_images.t", "images.0.compartment_id"),
				),
			},
			// update image display name
			{
				Config: s.Config + `
					resource "oci_core_image" "t" {
						compartment_id = "${var.tenancy_ocid}"
						instance_id = "${oci_core_instance.t.id}"
						display_name = "-tf-image"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-image"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "base_image_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_image_allowed", "true"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.ImageLifecycleStateAvailable)),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// ForceNew image by changing the launch mode to EMULATED
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_image" "t" {
						compartment_id = "${var.tenancy_ocid}"
						instance_id = "${oci_core_instance.t.id}"
						launch_mode = "EMULATED"
						display_name = "-tf-image"
						timeouts {
							create = "30m"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "launch_mode", "EMULATED"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.boot_volume_type", "IDE"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.firmware", "BIOS"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.network_type", "E1000"),
					resource.TestCheckResourceAttr(s.ResourceName, "launch_options.0.remote_data_volume_type", "SCSI"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "size_in_mbs"),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("resource updated when it was supposed to be recreated.")
						}
						return err
					},
				),
			},
			// Update compartment_id to ForceNew
			{
				Config: s.Config + `
					variable "update_compartment_id" {
						default = "` + getRequiredEnvSetting("compartment_id_for_update") + `"
					}
					resource "oci_core_image" "t" {
						compartment_id = "${var.update_compartment_id}"
						instance_id = "${oci_core_instance.t.id}"
						launch_mode = "EMULATED"
						display_name = "-tf-image"
						timeouts {
							create = "30m"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-image"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", getRequiredEnvSetting("compartment_id_for_update")),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "base_image_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "create_image_allowed", "true"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "operating_system"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "operating_system_version"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.ImageLifecycleStateAvailable)),
					func(ts *terraform.State) (err error) {
						resId2, err = fromInstanceState(ts, s.ResourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("resource updated when it was supposed to be recreated.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceCoreImageTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreImageTestSuite))
}
