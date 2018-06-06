// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DatabaseDBSystemShapeTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDBSystemShapeTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
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
				Config: s.Config + `
					data "oci_database_db_system_shapes" "t" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_system_shapes.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_system_shapes.0.name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_system_shapes.0.shape"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_system_shapes.0.available_core_count"),
				),
			},
			// Client-side filtering.
			{
				Config: s.Config + `
					data "oci_database_db_system_shapes" "t" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
						filter {
							name = "shape"
							values = ["non-existent-db-shape"]
						}
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "db_system_shapes.#", "0"),
				),
			},
			{
				Config: s.Config + `
					data "oci_database_db_system_shapes" "t" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
						filter {
							name = "shape"
							values = ["VM\\.Standard.+"]
							regex = true
						}
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(s.ResourceName, "db_system_shapes.#", regexp.MustCompile("[1-9][0-9]*")), // At least one image returned.
					resource.TestMatchResourceAttr(s.ResourceName, "db_system_shapes.0.name", regexp.MustCompile(`VM\.Standard.+`)),
					resource.TestMatchResourceAttr(s.ResourceName, "db_system_shapes.0.shape", regexp.MustCompile(`VM\.Standard.+`)),
				),
			},
		},
	},
	)
}

func TestDatasourceDatabaseDBSystemShapeTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDBSystemShapeTestSuite))
}
