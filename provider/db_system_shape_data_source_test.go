// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type DatabaseDBSystemShapeTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDBSystemShapeTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	data "oci_database_db_system_shapes" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
	}`
	s.ResourceName = "data.oci_database_db_system_shapes.t"
}

func (s *DatabaseDBSystemShapeTestSuite) TestAccDatasourceDatabaseDBSystemShape_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_system_shapes.0.name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_system_shapes.3.name"),
				),
			},
		},
	},
	)
}

func TestDatasourceDatabaseDBSystemShapeTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDBSystemShapeTestSuite))
}
