// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreIPSecConnectionConfigTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreIPSecConnectionConfigTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	resource "oci_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}
	resource "oci_core_cpe" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "displayname"
		ip_address = "123.123.123.123"
		depends_on = ["oci_core_drg.t"]
	}
	resource "oci_core_ipsec" "t" {
		compartment_id = "${var.compartment_id}"
		cpe_id = "${oci_core_cpe.t.id}"
		drg_id = "${oci_core_drg.t.id}"
		display_name = "display_name"
		static_routes = ["10.0.0.0/16"]
	}
	data "oci_core_ipsec_config" "s" {
		ipsec_id = "${oci_core_ipsec.t.id}"
	}`
	s.ResourceName = "data.oci_core_ipsec_config.s"
}

func (s *DatasourceCoreIPSecConnectionConfigTestSuite) TestAccDatasourceCoreIPSecConnectionConfig_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "tunnels.0.ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "tunnels.0.shared_secret"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "tunnels.0.time_created"),
				),
			},
		},
	},
	)

}

// issue-routing-tag: core/default
func TestDatasourceCoreIPSecConnectionConfigTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreIPSecConnectionConfigTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreIPSecConnectionConfigTestSuite))
}
