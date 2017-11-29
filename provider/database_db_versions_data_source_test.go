// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatabaseDBVersionTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDBVersionTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	data "oci_database_db_versions" "t" {
		compartment_id = "${var.compartment_id}"
	}`
	s.ResourceName = "data.oci_database_db_versions.t"
}

func (s *DatabaseDBVersionTestSuite) TestAccDatasourceDatabaseDBVersion_basic() {
	resource.Test(s.T(), resource.TestCase{
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

func TestDatasourceDatabaseDBVersionTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDBVersionTestSuite))
}
