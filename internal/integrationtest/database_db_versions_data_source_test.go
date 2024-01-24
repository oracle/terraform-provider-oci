// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"regexp"

	"github.com/stretchr/testify/suite"
)

type DatabaseDBVersionTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
}

func (s *DatabaseDBVersionTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig()
	s.ResourceName = "data.oci_database_db_versions.t"
}

func (s *DatabaseDBVersionTestSuite) TestAccDatasourceDatabaseDBVersion_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + `
					data "oci_database_db_versions" "t" {
						compartment_id = "${var.compartment_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_versions.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_versions.0.supports_pdb"),
					resource.TestMatchResourceAttr(s.ResourceName, "db_versions.0.version", regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)),
				),
			},
			{
				Config: s.Config + `
					data "oci_database_db_versions" "t" {
						compartment_id = "${var.compartment_id}"
						db_system_shape = "BM.DenseIO2.52"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "db_versions.#"),
					resource.TestCheckResourceAttr(s.ResourceName, "db_system_shape", "BM.DenseIO2.52"),
				),
			},
			// Client-side filtering.
			{
				Config: s.Config + `
					data "oci_database_db_versions" "t" {
						compartment_id = "${var.compartment_id}"
						db_system_shape = "BM.DenseIO2.52"
						filter {
							name = "version"
							values = ["12\\.\\d+\\.\\d+\\.\\d+"]
							regex = true
						}
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestMatchResourceAttr(s.ResourceName, "db_versions.#", regexp.MustCompile("[1-9][0-9]*")), // At least one version returned.
					resource.TestMatchResourceAttr(s.ResourceName, "db_versions.0.version", regexp.MustCompile(`12\.\d+\.\d+\.\d+`)),
				),
			},
			{
				Config: s.Config + `
					data "oci_database_db_versions" "t" {
						compartment_id = "${var.compartment_id}"
						db_system_shape = "BM.DenseIO2.52"
						filter {
							name = "version"
							values = ["non-existent-version"]
						}
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "db_versions.#", "0"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: database/default
func TestDatasourceDatabaseDBVersionTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceDatabaseDBVersionTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatabaseDBVersionTestSuite))
}
