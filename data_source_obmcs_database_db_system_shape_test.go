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

type DatabaseDBSystemShapeTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDBSystemShapeTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_database_db_system_shapes" "t" {
      availability_domain = "availability"
      compartment_id = "${var.compartment_id}"
      limit = 1
      page = "page"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_database_db_system_shapes.t"
}

func (s *DatabaseDBSystemShapeTestSuite) TestReadDBSystemShapes() {
	opts := &baremetal.PageListOptions{}
	opts.Page = "page"

	res := &baremetal.ListDBSystemShapes{}
	res.NextPage = "nextpage"
	res.DBSystemShapes = []baremetal.DBSystemShape{
		{
			AvailableCoreCount: 1,
			Name:               "name1",
			Shape:              "shape1",
		},
		{
			AvailableCoreCount: 2,
			Name:               "name2",
			Shape:              "shape2",
		},
	}

	s.Client.On(
		"ListDBSystemShapes", "availability", "compartmentid", uint64(1), opts,
	).Return(res, nil)

	opts2 := &baremetal.PageListOptions{}
	opts2.Page = "nextpage"
	s.Client.On(
		"ListDBSystemShapes", "availability", "compartmentid", uint64(1), opts2,
	).Return(
		&baremetal.ListDBSystemShapes{
			DBSystemShapes: []baremetal.DBSystemShape{
				{
					AvailableCoreCount: 1,
					Name:               "name3",
					Shape:              "shape3",
				},
				{
					AvailableCoreCount: 2,
					Name:               "name4",
					Shape:              "shape4",
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
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_system_shapes.0.name", "name1"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_system_shapes.3.name", "name4"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_system_shapes.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(
		s.T(), "ListDBSystemShapes", "availability", "compartmentid", uint64(1), opts2,
	)
}

func TestDatabaseDBSystemShapeTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDBSystemShapeTestSuite))
}
