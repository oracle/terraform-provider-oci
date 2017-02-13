// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreDrgsTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreDrgsTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_drgs" "t" {
      compartment_id = "compartment_id"
      limit = 1
      page = "page"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_drgs.t"
}

func (s *ResourceCoreDrgsTestSuite) TestReadDrgs() {
	opts := &baremetal.ListOptions{}
	opts.Limit = 1
	opts.Page = "page"

	s.Client.On(
		"ListDrgs",
		"compartment_id",
		opts,
	).Return(
		&baremetal.ListDrgs{
			Drgs: []baremetal.Drg{
				{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id1",
					State:         baremetal.ResourceAttached,
					TimeCreated:   baremetal.Time{Time: time.Now()},
				},
				{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id2",
					State:         baremetal.ResourceAttached,
					TimeCreated:   baremetal.Time{Time: time.Now()},
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
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.0.compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListDrgs", "compartment_id", opts)
}

func (s *ResourceCoreDrgsTestSuite) TestReadDrgsPaged() {
	opts := &baremetal.ListOptions{}
	opts.Limit = 1
	opts.Page = "page"

	res := &baremetal.ListDrgs{}
	res.NextPage = "nextpage"
	res.Drgs = []baremetal.Drg{
		{
			CompartmentID: "compartment_id",
			DisplayName:   "display_name",
			ID:            "id1",
			State:         baremetal.ResourceAttached,
			TimeCreated:   baremetal.Time{Time: time.Now()},
		},
		{
			CompartmentID: "compartment_id",
			DisplayName:   "display_name",
			ID:            "id2",
			State:         baremetal.ResourceAttached,
			TimeCreated:   baremetal.Time{Time: time.Now()},
		},
	}

	s.Client.On(
		"ListDrgs",
		"compartment_id",
		opts,
	).Return(res, nil)

	opts2 := &baremetal.ListOptions{}
	opts2.Page = "nextpage"
	opts2.Limit = 1

	s.Client.On(
		"ListDrgs",
		"compartment_id",
		opts2,
	).Return(
		&baremetal.ListDrgs{
			Drgs: []baremetal.Drg{
				{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id3",
					State:         baremetal.ResourceAttached,
					TimeCreated:   baremetal.Time{Time: time.Now()},
				},
				{
					CompartmentID: "compartment_id",
					DisplayName:   "display_name",
					ID:            "id4",
					State:         baremetal.ResourceAttached,
					TimeCreated:   baremetal.Time{Time: time.Now()},
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
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.0.compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.#", "4"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.2.id", "id3"),
					resource.TestCheckResourceAttr(s.ResourceName, "drgs.3.id", "id4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListDrgs", "compartment_id", opts2)

}

func TestResourceCoreDrgsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDrgsTestSuite))
}
