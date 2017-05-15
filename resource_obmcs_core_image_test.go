// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreImageTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Image
	DeletedRes   *baremetal.Image
}

func (s *ResourceCoreImageTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)
	s.Providers = map[string]terraform.ResourceProvider{"baremetal": s.Provider}

	s.ResourceName = "baremetal_core_image.t"
	s.Config = `
		resource "baremetal_core_image" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
			instance_id = "instance_id"
		}
	`
	s.Config += testProviderConfig()

}

func (s *ResourceCoreImageTestSuite) TestCreateImage() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "base_image_id"),

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func (s *ResourceCoreImageTestSuite) TestCreateImageWithoutDisplayName() {

	s.Config = `
		resource "baremetal_core_image" "t" {
			compartment_id = "${var.compartment_id}"
			instance_id = "instance_id"
		}
	`
	s.Config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
				),
			},
		},
	})
}

func (s ResourceCoreImageTestSuite) TestUpdateImageDisplayName() {

	config := `
		resource "baremetal_core_image" "t" {
			compartment_id = "${var.compartment_id}"
			instance_id = "instance_id"
			display_name = "new_display_name"
		}
	`
	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "new_display_name"),
				),
			},
		},
	})
}

func (s *ResourceCoreImageTestSuite) TestDeleteImage() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

}

func TestResourceCoreImageTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreImageTestSuite))
}
