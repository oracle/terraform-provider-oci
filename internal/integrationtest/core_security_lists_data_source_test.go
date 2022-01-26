// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/oracle/oci-go-sdk/v56/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreSecurityListTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreSecurityListTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-vcn"
	}
	resource "oci_core_security_list" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-security-list"
		vcn_id = "${oci_core_virtual_network.t.id}"
	}`
	s.ResourceName = "data.oci_core_security_lists.t"
}

func (s *DatasourceCoreSecurityListTestSuite) TestAccDatasourceCoreSecurityLists_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "security_lists.0.id", "oci_core_virtual_network.t", "default_security_list_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.display_name", "Default Security List for -tf-vcn"),
					acctest.CheckResourceSetContainsElementWithProperties(s.ResourceName, "security_lists.0.ingress_security_rules", map[string]string{
						"tcp_options.0.max": "22",
					},
						nil),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.state", string(core.SecurityListLifecycleStateAvailable)),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "security_lists.0.vcn_id", "oci_core_virtual_network.t", "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "security_lists.0.time_created"),
				),
			},
			// Test that enum fields such as 'state' can be filtered with multiple values
			{
				Config: s.Config + fmt.Sprintf(`
				data "oci_core_security_lists" "t" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					filter {
						name = "state"
						values = ["%s", "%s"]
						regex = true
					}
				}`, string(core.SecurityListLifecycleStateTerminated), string(core.SecurityListLifecycleStateAvailable)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.state", string(core.SecurityListLifecycleStateAvailable)),
				),
			},
			// Test that items can be filtered out
			{
				Config: s.Config + fmt.Sprintf(`
				data "oci_core_security_lists" "t" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					filter {
						name = "state"
						values = ["%s"]
						regex = true
					}
				}`, string(core.SecurityListLifecycleStateTerminated)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.#", "0"),
				),
			},
			// Server-side filtering tests.
			{
				Config: s.Config + `
				data "oci_core_security_lists" "t" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					display_name = "Default Security List for -tf-vcn"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "security_lists.0.id", "oci_core_virtual_network.t", "default_security_list_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.display_name", "Default Security List for -tf-vcn"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.state", string(core.SecurityListLifecycleStateAvailable)),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "security_lists.0.vcn_id", "oci_core_virtual_network.t", "id"),
				),
			},
			{
				Config: s.Config + `
				data "oci_core_security_lists" "t" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					display_name = "-tf-security-list"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "security_lists.0.id", "oci_core_security_list.t", "id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "security_lists.0.display_name", "oci_core_security_list.t", "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.0.state", string(core.SecurityListLifecycleStateAvailable)),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "security_lists.0.vcn_id", "oci_core_virtual_network.t", "id"),
				),
			},
			{
				Config: s.Config + `
				data "oci_core_security_lists" "t" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					state = "` + string(core.SecurityListLifecycleStateAvailable) + `"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "security_lists.#", "2"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: core/virtualNetwork
func TestDatasourceCoreSecurityListTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreSecurityListTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreSecurityListTestSuite))
}
