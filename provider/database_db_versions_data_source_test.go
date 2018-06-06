// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"regexp"

	"github.com/stretchr/testify/suite"
)

type DatabaseDBVersionTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatabaseDBVersionTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig()
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
				Config: s.Config + `
					data "oci_database_db_versions" "t" {
						compartment_id = "${var.compartment_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_versions.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_versions.0.supports_pdb"),
					resource.TestMatchResourceAttr(s.ResourceName, "db_versions.0.version", regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)),
				),
			},
			{
				Config: s.Config + `
					data "oci_database_db_versions" "t" {
						compartment_id = "${var.compartment_id}"
						db_system_shape = "BM.DenseIO1.36"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_versions.#"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_system_shape", "BM.DenseIO1.36"),
				),
			},
			// Client-side filtering.
			{
				Config: s.Config + `
					data "oci_database_db_versions" "t" {
						compartment_id = "${var.compartment_id}"
						db_system_shape = "BM.DenseIO1.36"
						filter {
							name = "version"
							values = ["12\\.\\d+\\.\\d+\\.\\d+"]
							regex = true
						}
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(s.ResourceName, "db_versions.#", regexp.MustCompile("[1-9][0-9]*")), // At least one version returned.
					resource.TestMatchResourceAttr(s.ResourceName, "db_versions.0.version", regexp.MustCompile(`12\.\d+\.\d+\.\d+`)),
				),
			},
			{
				Config: s.Config + `
					data "oci_database_db_versions" "t" {
						compartment_id = "${var.compartment_id}"
						db_system_shape = "BM.DenseIO1.36"
						filter {
							name = "version"
							values = ["non-existent-version"]
						}
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "db_versions.#", "0"),
				),
			},
		},
	},
	)
}

func TestDatasourceDatabaseDBVersionTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseDBVersionTestSuite))
}
