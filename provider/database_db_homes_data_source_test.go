// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceDBHomesTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListDBHomes
}

func (s *DatasourceDBHomesTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = `
resource "oci_database_db_home" "testDBHome" {
	#Required
	database {
		#Required
		admin_password = "${var.database_admin_password}"
		db_name = "${var.database_db_name}"

		#Optional
		character_set = "${var.database_character_set}"
		db_workload = "${var.database_db_workload}"
		ncharacter_set = "${var.database_ncharacter_set}"
		pdb_name = "${var.database_pdb_name}"
	}
	db_system_id = "${var.db_system_id}"
	db_version = "${var.db_version}"

	#Optional
	display_name = "${var.display_name}"
}

	`
	s.Config += testProviderConfig()
	s.ResourceName = "data.oci_database_db_homes.t"
}

func (s *DatasourceDBHomesTestSuite) TestReadDBHomes() {

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
data "oci_database_db_homes" "testDBHomes" {
	#Required
	compartment_id = "${var.compartment_id}"
	db_system_id = "${var.db_system_id}"
}

				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_homes.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_homes.#"),
				),
			},
		},
	},
	)
}

func TestDatasourceDBHomesTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceDBHomesTestSuite))
}
