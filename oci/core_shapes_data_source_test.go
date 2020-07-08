// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreShapeTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreShapeTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "t" {
		compartment_id = "${var.compartment_id}"
	}
	data "oci_core_shape" "t" {
		compartment_id = "${var.compartment_id}"
		availability_domain = "${data.oci_identity_availability_domains.t.availability_domains.0.name}"
		filter {
			name = "name"
			values = ["VM.Standard2.1"]
		}
	}`
	s.ResourceName = "data.oci_core_shape.t"
}

func (s *DatasourceCoreShapeTestSuite) TestAccDatasourceCoreShape_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.0.name", "VM.Standard2.1"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreShapeTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreShapeTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreShapeTestSuite))
}
