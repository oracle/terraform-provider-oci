// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"fmt"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreImageTestSuite struct {
	suite.Suite
	Config                 string
	Providers              map[string]terraform.ResourceProvider
	ResourceName           string
	FilterExpression       string
	OperatingSystem        string
	OperatingSystemVersion string
}

func (s *DatasourceCoreImageTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = testProviderConfig()
	s.ResourceName = "data.oci_core_images.t"
	s.OperatingSystem = "Oracle Linux"
}

func (s *DatasourceCoreImageTestSuite) TestAccImage_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + fmt.Sprintf(`
				data "oci_core_images" "allOracleImages" {
					compartment_id = "${var.tenancy_ocid}"
					operating_system = "%s"
					shape = "VM.Standard1.8"
				}

				data "oci_core_images" "t" {
					compartment_id = "${var.tenancy_ocid}"
					operating_system = "${lookup(data.oci_core_images.allOracleImages.images[0], "operating_system")}"
					operating_system_version = "${lookup(data.oci_core_images.allOracleImages.images[0], "operating_system_version")}"
				
					filter {
						name = "display_name"
						values = ["${lookup(data.oci_core_images.allOracleImages.images[0], "display_name")}"]
						regex = true
					}
				}`, s.OperatingSystem),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "images.#", "1"),
					TestCheckResourceAttributesEqual(s.ResourceName, "images.0.id", "data.oci_core_images.allOracleImages", "images.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.create_image_allowed", "true"),
					TestCheckResourceAttributesEqual(s.ResourceName, "images.0.display_name", "data.oci_core_images.allOracleImages", "images.0.display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.state", string(core.ImageLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.launch_mode"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.launch_options.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.launch_options.0.boot_volume_type"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.launch_options.0.firmware"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.launch_options.0.network_type"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.launch_options.0.remote_data_volume_type"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.operating_system", s.OperatingSystem),
					TestCheckResourceAttributesEqual(s.ResourceName, "images.0.operating_system_version", "data.oci_core_images.allOracleImages", "images.0.operating_system_version"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.time_created"),
					// This test filters to official images, which do not derive from another so the below properties are expected to be null
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.base_image_id", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.instance_id", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.compartment_id", ""),
				),
			},
			{
				Config: s.Config + fmt.Sprintf(`
				data "oci_core_images" "allOracleImages" {
					compartment_id = "${var.tenancy_ocid}"
					operating_system = "%s"
					shape = "VM.Standard1.8"
				}

				data "oci_core_images" "t" {
					compartment_id = "${var.tenancy_ocid}"
					operating_system = "${lookup(data.oci_core_images.allOracleImages.images[0], "operating_system")}"
					operating_system_version = "${lookup(data.oci_core_images.allOracleImages.images[0], "operating_system_version")}"

					filter {
						name = "launch_options.is_pv_encryption_in_transit_enabled"
						values = ["true"]
					}
				}`, s.OperatingSystem),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.create_image_allowed", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.state", string(core.ImageLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.launch_mode"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.launch_options.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.launch_options.0.boot_volume_type"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.launch_options.0.firmware"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.launch_options.0.network_type"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.launch_options.0.remote_data_volume_type"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.launch_options.0.is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.operating_system", s.OperatingSystem),
					TestCheckResourceAttributesEqual(s.ResourceName, "images.0.operating_system_version", "data.oci_core_images.allOracleImages", "images.0.operating_system_version"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.time_created"),
					// This test filters to official images, which do not derive from another so the below properties are expected to be null
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.base_image_id", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.instance_id", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.compartment_id", ""),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreImageTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreImageTestSuite))
}
