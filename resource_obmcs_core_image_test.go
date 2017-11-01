// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreImageTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreImageTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + instanceConfig
	s.ResourceName = "oci_core_image.t"
}

func (s *ResourceCoreImageTestSuite) TestAccResourceCoreImage_basic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// create image
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_image" "t" {
						compartment_id = "${var.compartment_id}"
						instance_id = "${oci_core_instance.t.id}"
						timeouts {
							create = "30m"
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "base_image_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
				),
			},
			// update image display name
			{
				Config: s.Config + `
					resource "oci_core_image" "t" {
						compartment_id = "${var.compartment_id}"
						instance_id = "${oci_core_instance.t.id}"
						display_name = "-tf-image"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-image"),
				),
			},
		},
	})
}

func TestResourceCoreImageTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreImageTestSuite))
}
