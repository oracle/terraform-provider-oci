// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/stretchr/testify/suite"
	//"github.com/oracle/oci-go-sdk/v65/core"
)

type DatasourceCoreIPSecStatusTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
}

func (s *DatasourceCoreIPSecStatusTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	resource "oci_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-drg"
	}
	resource "oci_core_cpe" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-cpe"
		ip_address = "123.123.123.123"
	}
	resource "oci_core_ipsec" "t" {
		compartment_id = "${var.compartment_id}"
		cpe_id = "${oci_core_cpe.t.id}"
		drg_id = "${oci_core_drg.t.id}"
		display_name = "-tf-ipsec"
		static_routes = ["10.0.0.0/16"]
	}`
	s.ResourceName = "data.oci_core_ipsec_status.s"
}

func (s *DatasourceCoreIPSecStatusTestSuite) TestAccDatasourceCoreIPSecStatus_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + `
				data "oci_core_ipsec_status" "s" {
					ipsec_id = "${oci_core_ipsec.t.id}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "tunnels.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "tunnels.0.ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "tunnels.0.time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "tunnels.0.time_state_modified"),
					// TODO: During testing, the service returned "DOWN" for the state which is not an expected value (not defined in the spec)
					// TODO: Also will need to investigate why "DOWN" was returned when Available was expected
					//resource.TestCheckResourceAttr(s.ResourceName, "tunnels.0.state", string(core.IpSecConnectionLifecycleStateAvailable)),
				),
			},
		},
	},
	)
}

// issue-routing-tag: core/default
func TestDatasourceCoreIPSecStatusTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreIPSecStatusTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreIPSecStatusTestSuite))
}
