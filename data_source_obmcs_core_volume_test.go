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

type ResourceCoreVolumesTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVolumesTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_volumes" "t" {
      availability_domain = "availability_domain"
      compartment_id = "compartment_id"
      limit = 1
      page = "page"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_volumes.t"
}

func (s *ResourceCoreVolumesTestSuite) TestReadVolumes() {
	opts := &baremetal.ListVolumesOptions{}
	opts.AvailabilityDomain = "availability_domain"
	opts.Limit = 1
	opts.Page = "page"

	s.Client.On(
		"ListVolumes",
		"compartment_id",
		opts,
	).Return(
		&baremetal.ListVolumes{
			Volumes: []baremetal.Volume{
				{
					AvailabilityDomain: "availability_domain",
					CompartmentID:      "compartment_id",
					DisplayName:        "display_name",
					ID:                 "id1",
					SizeInMBs:          123,
					State:              baremetal.ResourceAvailable,
					TimeCreated:        baremetal.Time{Time: time.Now()},
				},
				{
					AvailabilityDomain: "availability_domain",
					CompartmentID:      "compartment_id",
					DisplayName:        "display_name",
					ID:                 "id2",
					SizeInMBs:          123,
					State:              baremetal.ResourceAvailable,
					TimeCreated:        baremetal.Time{Time: time.Now()},
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVolumes", "compartment_id", opts)
}

func (s *ResourceCoreVolumesTestSuite) TestReadVolumesWithPagination() {
	opts := &baremetal.ListVolumesOptions{}
	opts.AvailabilityDomain = "availability_domain"
	opts.Limit = 1
	opts.Page = "page"

	res := &baremetal.ListVolumes{}
	res.NextPage = "nextpage"
	res.Volumes = []baremetal.Volume{
		{
			AvailabilityDomain: "availability_domain",
			CompartmentID:      "compartment_id",
			DisplayName:        "display_name",
			ID:                 "id1",
			SizeInMBs:          123,
			State:              baremetal.ResourceAvailable,
			TimeCreated:        baremetal.Time{Time: time.Now()},
		},
		{
			AvailabilityDomain: "availability_domain",
			CompartmentID:      "compartment_id",
			DisplayName:        "display_name",
			ID:                 "id2",
			SizeInMBs:          123,
			State:              baremetal.ResourceAvailable,
			TimeCreated:        baremetal.Time{Time: time.Now()},
		},
	}

	s.Client.On(
		"ListVolumes",
		"compartment_id",
		opts,
	).Return(res, nil)

	opts2 := &baremetal.ListVolumesOptions{}
	opts2.AvailabilityDomain = "availability_domain"
	opts2.Limit = 1
	opts2.Page = "nextpage"

	s.Client.On(
		"ListVolumes",
		"compartment_id",
		opts2,
	).Return(
		&baremetal.ListVolumes{
			Volumes: []baremetal.Volume{
				{
					AvailabilityDomain: "availability_domain",
					CompartmentID:      "compartment_id",
					DisplayName:        "display_name",
					ID:                 "id3",
					SizeInMBs:          123,
					State:              baremetal.ResourceAvailable,
					TimeCreated:        baremetal.Time{Time: time.Now()},
				},
				{
					AvailabilityDomain: "availability_domain",
					CompartmentID:      "compartment_id",
					DisplayName:        "display_name",
					ID:                 "id4",
					SizeInMBs:          123,
					State:              baremetal.ResourceAvailable,
					TimeCreated:        baremetal.Time{Time: time.Now()},
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "volumes.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVolumes", "compartment_id", opts2)
}

func TestResourceCoreVolumesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumesTestSuite))
}
