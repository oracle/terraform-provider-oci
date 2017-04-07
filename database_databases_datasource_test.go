// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"

	"github.com/stretchr/testify/suite"
)

type DatabaseDatabasesTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDatabasesTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_database_databases" "t" {
      compartment_id = "compartment_id"
      db_home_id = "db_home_id"
      limit = 1
      page = "page"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_database_databases.t"
}

func (s *DatabaseDatabasesTestSuite) TestReadDatabases() {
	opts := &baremetal.PageListOptions{}
	opts.Page = "page"

	res := &baremetal.ListDatabases{}
	res.NextPage = "nextpage"
	res.Databases = []baremetal.Database{{ID: "1"}, {ID: "2"}}

	s.Client.On(
		"ListDatabases", "compartment_id", "db_home_id", uint64(1), opts,
	).Return(res, nil)

	opts2 := &baremetal.PageListOptions{}
	opts2.Page = "nextpage"
	s.Client.On(
		"ListDatabases", "compartment_id", "db_home_id", uint64(1), opts2,
	).Return(
		&baremetal.ListDatabases{
			Databases: []baremetal.Database{{ID: "3"}, {ID: "4"}},
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
					resource.TestCheckResourceAttr(s.ResourceName, "db_home_id", "db_home_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "databases.0.id", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "databases.3.id", "4"),
					resource.TestCheckResourceAttr(s.ResourceName, "databases.#", "4"),
				),
			},
		},
	},
	)
}

func TestDatabaseDatabasesTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDatabasesTestSuite))
}
