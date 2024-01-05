// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"regexp"
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/suite"
)

type DatabaseDBSystemShapeTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
}

func (s *DatabaseDBSystemShapeTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
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
				Config: s.Config + `
					data "oci_database_db_system_shapes" "t" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestMatchResourceAttr(s.ResourceName, "db_system_shapes.#", regexp.MustCompile("[1-9][0-9]*")), // At least one image returned.
					resource.TestMatchResourceAttr(s.ResourceName, "db_system_shapes.0.name", regexp.MustCompile(`VM\.Standard.+`)),
					resource.TestMatchResourceAttr(s.ResourceName, "db_system_shapes.0.shape", regexp.MustCompile(`VM\.Standard.+`)),
				),
			},
		},
	},
	)
}

// issue-routing-tag: database/default
func TestDatasourceDatabaseDBSystemShapeTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceDatabaseDBSystemShapeTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatabaseDBSystemShapeTestSuite))
}
