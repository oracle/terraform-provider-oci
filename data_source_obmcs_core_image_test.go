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
	Client       mockableClient
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
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_images" "t" {
      compartment_id = "compartment_id"
      limit = 1
      page = "page"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_images.t"

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
	opts := &baremetal.ListImagesOptions{}
	opts.Limit = 1
	opts.Page = "page"

	s.Client.On("ListImages", "compartment_id", opts).Return(s.List, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListImages", "compartment_id", opts)
}

func (s *ResourceCoreImagesTestSuite) TestReadImagesWithPagination() {
	opts := &baremetal.ListImagesOptions{}
	opts.Limit = 1
	opts.Page = "page"

	listVal := *s.List
	list := &listVal
	list.NextPage = "nextpage"
	s.Client.On("ListImages", "compartment_id", opts).Return(list, nil)

	opts2 := &baremetal.ListImagesOptions{}
	opts2.Limit = 1
	opts2.Page = "nextpage"

	list2Val := *s.List
	list2 := &list2Val
	b3 := s.List.Images[0]
	b3.ID = "id3"
	b4 := s.List.Images[1]
	b4.ID = "id4"
	list2.Images = []baremetal.Image{b3, b4}
	s.Client.On("ListImages", "compartment_id", opts2).Return(list2, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "images.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "images.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListImages", "compartment_id", opts2)
}

func TestResourceCoreImagesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreImagesTestSuite))
}
