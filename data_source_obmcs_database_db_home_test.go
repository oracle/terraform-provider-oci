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

type DatabaseDBHomeTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDBHomeTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_database_db_home" "t" {
      db_home_id = "id"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_database_db_home.t"
}

func (s *DatabaseDBHomeTestSuite) TestReadDBHome() {
	dbHome := &baremetal.DBHome{}
	dbHome.CompartmentID = "compartment_id"
	dbHome.DBSystemID = "db_system_id"
	dbHome.ID = "id"

	s.Client.On("GetDBHome", "id").Return(dbHome, nil)

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
					resource.TestCheckResourceAttr(s.ResourceName, "db_system_id", "db_system_id"),
				),
			},
		},
	},
	)
}

func TestDatabaseDBHomeTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDBHomeTestSuite))
}
