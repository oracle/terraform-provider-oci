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

type DatabaseDBSystemShapeTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDBSystemShapeTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = `
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.compartment_id}"
}
    data "oci_database_db_system_shapes" "t" {
      availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
      compartment_id = "${var.compartment_id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.oci_database_db_system_shapes.t"
}

func (s *DatabaseDBSystemShapeTestSuite) TestReadDBSystemShapes() {
	resource.UnitTest(s.T(), resource.TestCase{
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

func TestDatabaseDBSystemShapeTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDBSystemShapeTestSuite))
}
