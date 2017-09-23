// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreShapeTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreShapeTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	data "oci_identity_availability_domains" "t" {
		compartment_id = "${var.compartment_id}"
	}
	data "oci_core_shape" "s" {
		compartment_id = "${var.compartment_id}"
		availability_domain = "${data.oci_identity_availability_domains.t.availability_domains.0.name}"
	}`
	s.ResourceName = "data.oci_core_shape.s"
}

func (s *DatasourceCoreShapeTestSuite) TestAccDatasourceCoreShape_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shapes.0.name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shapes.1.name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shapes.#"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreShapeTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreShapeTestSuite))
}
