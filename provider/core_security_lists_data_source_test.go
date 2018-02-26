// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreSecurityListTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreSecurityListTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-vcn"
	}`
	s.ResourceName = "data.oci_core_security_lists.t"
}

func (s *DatasourceCoreSecurityListTestSuite) TestAccDatasourceCoreSecurityLists_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_core_security_lists" "t" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					filter {
						name = "display_name"
						values = ["Default Security List.*"]
						regex = true
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.display_name", "Default Security List for -tf-vcn"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.ingress_security_rules.0.tcp_options.0.max", "22"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.state", string(core.SecurityListLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "security_lists.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "security_lists.0.vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "security_lists.0.time_created"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreSecurityListTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreSecurityListTestSuite))
}
