// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatabaseDBVersionTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDBVersionTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_database_db_versions" "t" {
      compartment_id = "compartmentid"
      limit = 1
      page = "page"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_database_db_versions.t"
}

func (s *DatabaseDBVersionTestSuite) TestReadDBVersions() {
	opts := &baremetal.PageListOptions{}
	opts.Page = "page"

	res := &baremetal.ListDBVersions{}
	res.NextPage = "nextpage"
	res.DBVersions = []baremetal.DBVersion{
		{Version: "version1"}, {Version: "version2"},
	}

	s.Client.On(
		"ListDBVersions", "compartmentid", uint64(1), opts,
	).Return(res, nil)

	opts2 := &baremetal.PageListOptions{}
	opts2.Page = "nextpage"
	s.Client.On(
		"ListDBVersions", "compartmentid", uint64(1), opts2,
	).Return(
		&baremetal.ListDBVersions{
			DBVersions: []baremetal.DBVersion{
				{Version: "version3"}, {Version: "version4"},
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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_versions.0.version", "version1"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_versions.3.version", "version4"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_versions.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(
		s.T(), "ListDBVersions", "compartmentid", uint64(1), opts2,
	)
}

func TestDatabaseDBVersionTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDBVersionTestSuite))
}
