// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreImagesTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListImages
}

func (s *ResourceCoreImagesTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = `
    data "oci_core_images" "t" {
      compartment_id = "${var.compartment_id}"
      limit = 1
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.oci_core_images.t"

	b1 := baremetal.Image{
		BaseImageID:            "base_image_id",
		CompartmentID:          "compartment_id",
		CreateImageAllowed:     true,
		DisplayName:            "display_name",
		ID:                     "id1",
		State:                  baremetal.ResourceAvailable,
		OperatingSystem:        "operating_system",
		OperatingSystemVersion: "operating_system_version",
		TimeCreated:            baremetal.Time{Time: time.Now()},
	}

	b2 := b1
	b2.ID = "id2"

	s.List = &baremetal.ListImages{
		Images: []baremetal.Image{b1, b2},
	}
}

func (s *ResourceCoreImagesTestSuite) TestReadImages() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.1.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "images.#"),
				),
			},
		},
	},
	)

}

func TestResourceCoreImagesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreImagesTestSuite))
}
