// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type ResourceDatabaseDBSystemTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Res          *baremetal.DBSystem
}

func (s *ResourceDatabaseDBSystemTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders

	s.Config = databaseConfig

	s.Config += testProviderConfig()

	s.ResourceName = "oci_database_db_system.t"
}

func (s *ResourceDatabaseDBSystemTestSuite) TestCreateResourceDatabaseDBSystem() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "MyTFDatabaseNode0"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func TestResourceDatabaseDBSystemTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceDatabaseDBSystemTestSuite))
}
