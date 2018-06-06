// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type CoreConsoleHistoryDataDatasourceTestSuite struct {
	suite.Suite
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *CoreConsoleHistoryDataDatasourceTestSuite) SetupTest() {
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + instanceConfig + `
	resource "oci_core_console_history" "t" {
		instance_id = "${oci_core_instance.t.id}"
	}
	data "oci_core_console_history_data" "s" {
		console_history_id = "${oci_core_console_history.t.id}"
		length = 10240
	}`
	s.ResourceName = "data.oci_core_console_history_data.s"
}

func (s *CoreConsoleHistoryDataDatasourceTestSuite) TestAccDatasourceCoreConsoleHistory_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "console_history_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "data"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "length"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreConsoleHistoryTestSuite(t *testing.T) {
	suite.Run(t, new(CoreConsoleHistoryDataDatasourceTestSuite))
}
