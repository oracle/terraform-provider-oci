// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"fmt"
	"github.com/stretchr/testify/suite"
	"regexp"
)

type DatasourceCoreImageTestSuite struct {
	suite.Suite
	Client                 *baremetal.Client
	Config                 string
	Provider               terraform.ResourceProvider
	Providers              map[string]terraform.ResourceProvider
	ResourceName           string
	FilterExpression       string
	OperatingSystem        string
	OperatingSystemVersion string
}

func (s *DatasourceCoreImageTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
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
					compartment_id = "${var.compartment_id}"
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
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.state", "AVAILABLE"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.operating_system", s.OperatingSystem),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.operating_system_version", s.OperatingSystemVersion),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.time_created"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreImageTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreImageTestSuite))
}
