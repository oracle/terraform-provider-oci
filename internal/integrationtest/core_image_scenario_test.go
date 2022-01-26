// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreImageTestSuite struct {
	suite.Suite
	Providers              map[string]terraform.ResourceProvider
	Config                 string
	OperatingSystem        string
	OperatingSystemVersion string
}

func (s *ResourceCoreImageTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + instanceConfig + DefinedTagsDependencies
	s.OperatingSystem = "Oracle Linux"
}

func (s *ResourceCoreImageTestSuite) TestAccResourceCoreImage_objectStorageImageSources() {
	/*
	 * This test requires an image to have been exported relative to bucket_name and object_name below.
	 * bucket_name: test-fixtures-bucket
	 * object_name: test-fixtures-image-export
	 * as well as an image PAR url supplied via env var image_par
	 */
	imagePar := utils.GetEnvSettingWithBlankDefault("image_par")
	if imagePar == "" {
		s.T().Skip("Dependency image_par not defined for test")
	}

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: acctest.LegacyTestProviderConfig() + `
					data "oci_objectstorage_namespace" "t" {}
					resource "oci_core_image" "i1" {
						compartment_id = "${var.tenancy_ocid}"
						image_source_details {
							source_type = "objectStorageTuple"
							namespace_name = "${data.oci_objectstorage_namespace.t.namespace}"
							bucket_name = "test-fixtures-bucket"
							object_name = "test-fixtures-image-export"
							operating_system = "Oracle Linux"
							operating_system_version = "7.2"
						}
						timeouts {
							create = "30m"
						}
					}
					resource "oci_core_image" "i2" {
						compartment_id = "${var.tenancy_ocid}"
						image_source_details {
							source_type = "objectStorageUri"
							source_uri = "` + imagePar + `"
							source_image_type = "QCOW2"
						}
						timeouts {
							create = "30m"
						}
					}
				`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckNoResourceAttr("oci_core_image.i1", "instance_id"),
					resource.TestCheckNoResourceAttr("oci_core_image.i1", "base_image_id"),
					resource.TestCheckResourceAttr("oci_core_image.i1", "operating_system", "Oracle Linux"),
					resource.TestCheckResourceAttr("oci_core_image.i1", "operating_system_version", "7.2"),

					resource.TestCheckNoResourceAttr("oci_core_image.i2", "instance_id"),
					resource.TestCheckNoResourceAttr("oci_core_image.i2", "base_image_id"),
					resource.TestCheckResourceAttr("oci_core_image.i2", "operating_system", "Custom"),
					resource.TestCheckResourceAttr("oci_core_image.i2", "operating_system_version", "Custom"),
				),
			},
			// data source
			{
				Config: s.Config + fmt.Sprintf(`
				data "oci_core_images" "allOracleImages" {
					compartment_id = "${var.tenancy_ocid}"
					operating_system = "%s"
					shape = "VM.Standard2.1"
				}

				data "oci_core_images" "t" {
					compartment_id = "${var.tenancy_ocid}"
					operating_system = "${lookup(data.oci_core_images.allOracleImages.images[0], "operating_system")}"
					operating_system_version = "${lookup(data.oci_core_images.allOracleImages.images[0], "operating_system_version")}"

					filter {
						name = "launch_options.is_pv_encryption_in_transit_enabled"
						values = ["true"]
					}

                    sort_by = "TIMECREATED"
                    sort_order = "DESC"
				}`, s.OperatingSystem),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("data.oci_core_images.t", "images.0.create_image_allowed", "true"),
					resource.TestCheckResourceAttr("data.oci_core_images.t", "images.0.state", string(core.ImageLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet("data.oci_core_images.t", "images.0.launch_mode"),
					resource.TestCheckResourceAttr("data.oci_core_images.t", "images.0.launch_options.#", "1"),
					resource.TestCheckResourceAttrSet("data.oci_core_images.t", "images.0.launch_options.0.boot_volume_type"),
					resource.TestCheckResourceAttrSet("data.oci_core_images.t", "images.0.launch_options.0.firmware"),
					resource.TestCheckResourceAttrSet("data.oci_core_images.t", "images.0.launch_options.0.network_type"),
					resource.TestCheckResourceAttrSet("data.oci_core_images.t", "images.0.launch_options.0.remote_data_volume_type"),
					resource.TestCheckResourceAttr("data.oci_core_images.t", "images.0.launch_options.0.is_pv_encryption_in_transit_enabled", "true"),
					resource.TestCheckResourceAttr("data.oci_core_images.t", "images.0.operating_system", s.OperatingSystem),
					acctest.TestCheckResourceAttributesEqual("data.oci_core_images.t", "images.0.operating_system_version", "data.oci_core_images.allOracleImages", "images.0.operating_system_version"),
					resource.TestCheckResourceAttrSet("data.oci_core_images.t", "images.0.time_created"),
					// This test filters to official images, which do not derive from another so the below properties are expected to be null
					resource.TestCheckResourceAttr("data.oci_core_images.t", "images.0.base_image_id", ""),
					resource.TestCheckResourceAttr("data.oci_core_images.t", "images.0.instance_id", ""),
					resource.TestCheckResourceAttr("data.oci_core_images.t", "images.0.compartment_id", ""),
				),
			},
		},
	})
}

// issue-routing-tag: core/computeImaging
func TestResourceCoreImageTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreImageTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreImageTestSuite))
}
