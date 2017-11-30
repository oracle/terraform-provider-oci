// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type ResourceDBHomeTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.DBHome
	DeletedRes   *baremetal.DBHome
}

func (s *ResourceDBHomeTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
resource "oci_database_db_home" "testDBHome" {
	#Required
	database {
		#Required
		admin_password = "${var.database_admin_password}"
		db_name = "${var.database_db_name}"
	}
	db_system_id = "${var.db_system_id}"
	db_version = "${var.db_version}"
}

	`

	s.Config += testProviderConfig()

	s.ResourceName = "oci_database_db_home.t"

}

func (s *ResourceDBHomeTestSuite) TestCreateResourceDBHome() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
				),
			},
		},
	})
}

func (s ResourceDBHomeTestSuite) TestUpdateDBHomeDBVersion() {
	config := `
resource "oci_database_db_home" "testDBHome" {
	#Required
	database {
		#Required
		admin_password = "${var.database_admin_password}"
		db_name = "${var.database_db_name}"
	}
	db_system_id = "${var.db_system_id}"
	db_version = "${var.db_version}"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "db_version", "newValue"),
				),
			},
		},
	})
}

func (s ResourceDBHomeTestSuite) TestUpdateDatabaseForcesNewDBHome() {

	config := `
resource "oci_database_db_home" "testDBHome" {
	#Required
	database {
		#Required
		admin_password = "${var.database_admin_password}"
		db_name = "${var.database_db_name}"
	}
	db_system_id = "${var.db_system_id}"
	db_version = "${var.db_version}"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "database", "newValue"),
				),
			},
		},
	})
}

func (s ResourceDBHomeTestSuite) TestUpdateDBSystemIDForcesNewDBHome() {

	config := `
resource "oci_database_db_home" "testDBHome" {
	#Required
	database {
		#Required
		admin_password = "${var.database_admin_password}"
		db_name = "${var.database_db_name}"
	}
	db_system_id = "${var.db_system_id}"
	db_version = "${var.db_version}"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "db_system_id", "newValue"),
				),
			},
		},
	})
}

func (s ResourceDBHomeTestSuite) TestUpdateDisplayNameForcesNewDBHome() {

	config := `
resource "oci_database_db_home" "testDBHome" {
	#Required
	database {
		#Required
		admin_password = "${var.database_admin_password}"
		db_name = "${var.database_db_name}"
	}
	db_system_id = "${var.db_system_id}"
	db_version = "${var.db_version}"

	#Optional
	display_name = "${var.display_name}"
}
`

	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "newValue"),
				),
			},
		},
	})
}

func TestResourceDBHomeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceDBHomeTestSuite))
}
