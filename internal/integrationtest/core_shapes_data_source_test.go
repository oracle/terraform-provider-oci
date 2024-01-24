// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreShapeTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
}

type DatasourceCoreFlexShapeTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
}

func (s *DatasourceCoreShapeTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
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

func (s *DatasourceCoreFlexShapeTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "flex_shape_test" {
		compartment_id = "${var.compartment_id}"
	}
	data "oci_core_shape" "flex_shape_test" {
		compartment_id = "${var.compartment_id}"
		availability_domain = "${data.oci_identity_availability_domains.flex_shape_test.availability_domains.0.name}"
		filter {
			name = "name"
			values = ["VM.Standard.E3.Flex"]
		}
	}`
	s.ResourceName = "data.oci_core_shape.flex_shape_test"
}

func (s *DatasourceCoreShapeTestSuite) TestAccDatasourceCoreShape_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.0.name", "VM.Standard2.1"),
				),
			},
		},
	},
	)
}

func (s *DatasourceCoreFlexShapeTestSuite) TestAccDatasourceFlexCoreShape_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.0.name", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.0.memory_options.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shapes.0.memory_options.0.max_per_numa_node_in_gbs"),
					resource.TestCheckResourceAttr(s.ResourceName, "shapes.0.ocpu_options.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shapes.0.ocpu_options.0.max_per_numa_node"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestDatasourceCoreShapeTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreShapeTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreShapeTestSuite))
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestDatasourceCoreFlexShapeTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreFlexShapeTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreFlexShapeTestSuite))
}
