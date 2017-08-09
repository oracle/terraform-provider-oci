// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-baremetal/client"
)

type DatabaseDBVersionTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDBVersionTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_database_db_versions" "t" {
      compartment_id = "${var.compartment_id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_database_db_versions.t"
}

func (s *DatabaseDBVersionTestSuite) TestReadDBVersions() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(s.ResourceName, "db_versions.0.version"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_versions.1.version"),
				),
			},
		},
	},
	)
}

func TestDatabaseDBVersionTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDBVersionTestSuite))
}
