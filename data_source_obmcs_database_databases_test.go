// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DatabaseDatabasesTestSuite struct {
	suite.Suite
	Client    mockableClient
	Config    string
	Provider  terraform.ResourceProvider
	Providers map[string]terraform.ResourceProvider
}

func (s *DatabaseDatabasesTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = databaseConfig

	s.Config += testProviderConfig()
}

func (s *DatabaseDatabasesTestSuite) TestReadDatabases() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "baremetal_database_db_systems" "t" {
				  compartment_id = "${var.compartment_id}"
				}
				data "baremetal_database_db_homes" "t" {
				  compartment_id = "${var.compartment_id}"
				  db_system_id = "${baremetal_database_db_system.t.id}"
				}
			        data "baremetal_database_databases" "t" {
				      compartment_id = "${var.compartment_id}"
				      db_home_id = "${data.baremetal_database_db_homes.t.id}"
				}
			        data "baremetal_database_db_nodes" "t" {
				      compartment_id = "${var.compartment_id}"
				      db_system_id = "${baremetal_database_db_system.t.id}"
			        }
				`,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet("data.baremetal_database_db_systems.t", "db_systems.#"),
					resource.TestCheckResourceAttrSet("data.baremetal_database_db_homes.t", "db_homes.#"),
					resource.TestCheckResourceAttrSet("data.baremetal_database_databases", "databases.#"),
					resource.TestCheckResourceAttrSet("data.baremetal_database_db_nodes", "db_nodes.#"),
				),
			},
		},
	},
	)
}

func TestDatabaseDatabasesTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDatabasesTestSuite))
}
