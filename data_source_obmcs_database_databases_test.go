// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DatabaseDatabasesTestSuite struct {
	suite.Suite
	Client    *baremetal.Client
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
		"oci": s.Provider,
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
				data "oci_database_db_systems" "t" {
				  compartment_id = "${var.compartment_id}"
				}
				data "oci_database_db_homes" "t" {
				  compartment_id = "${var.compartment_id}"
				  db_system_id = "${oci_database_db_system.t.id}"
				}
			        data "oci_database_databases" "t" {
				      compartment_id = "${var.compartment_id}"
				      db_home_id = "${data.oci_database_db_homes.t.id}"
				}
				data "oci_database_database" "t" {
				      database_id = "${data.oci_database_databases.t.databases.0.id}"
				}
			        data "oci_database_db_nodes" "t" {
				      compartment_id = "${var.compartment_id}"
				      db_system_id = "${oci_database_db_system.t.id}"
			        }
				`,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet("data.oci_database_db_systems.t", "db_systems.#"),
					resource.TestCheckResourceAttrSet("data.oci_database_db_homes.t", "db_homes.#"),
					resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.#"),
					resource.TestCheckResourceAttrSet("data.oci_database_db_nodes.t", "db_nodes.#"),
					resource.TestCheckResourceAttrSet("data.oci_database_database.t", "id"),
				),
			},
		},
	},
	)
}

func TestDatabaseDatabasesTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDatabasesTestSuite))
}
