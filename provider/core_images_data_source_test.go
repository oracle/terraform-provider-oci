// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"fmt"
	"regexp"

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
	s.Config = testProviderConfig()
	s.ResourceName = "data.oci_core_images.t"
	// This test will need to be updated when this image is removed from ListImages.
	s.FilterExpression = ".*2017.12.18-0"
	s.OperatingSystem = "Oracle Linux"
	s.OperatingSystemVersion = "7.4"
}

func (s *DatasourceCoreImageTestSuite) TestAccImage_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + fmt.Sprintf(`
				data "oci_core_images" "t" {
					compartment_id = "${var.tenancy_ocid}"
					operating_system = "%s"
					operating_system_version = "%s"
				
					filter {
						name = "display_name"
						values = ["%s"]
						regex = true
					}
				}`, s.OperatingSystem, s.OperatingSystemVersion, s.FilterExpression),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(s.ResourceName, "images.#", regexp.MustCompile("[1-9][0-9]*")),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.create_image_allowed", "true"),
					resource.TestMatchResourceAttr(s.ResourceName, "images.0.display_name", regexp.MustCompile(s.FilterExpression)),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.state", string(core.ImageLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.operating_system", s.OperatingSystem),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.operating_system_version", s.OperatingSystemVersion),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.base_image_id", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.instance_id", ""),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.compartment_id"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreImageTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreImageTestSuite))
}
