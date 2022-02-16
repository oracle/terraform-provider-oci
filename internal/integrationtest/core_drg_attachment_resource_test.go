// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"fmt"

	"github.com/oracle/oci-go-sdk/v58/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreDrgAttachmentTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreDrgAttachmentTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
		resource "oci_core_virtual_network" "t" {
			cidr_block = "10.0.0.0/16"
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-vcn"
		}
		resource "oci_core_drg" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-drg"
		}`

	s.ResourceName = "oci_core_drg_attachment.t"
}

func (s *ResourceCoreDrgAttachmentTestSuite) TestAccResourceCoreDrgAttachment_basic() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify a drg attachment can be created
			{
				Config: s.Config + `
				resource "oci_core_drg_attachment" "t" {
					drg_id = "${oci_core_drg.t.id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.DrgAttachmentLifecycleStateAttached)),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_core_drg.t", "id")
						return err
					},
				),
			},
			// verify drg attachment Update
			{
				Config: s.Config + `
				resource "oci_core_drg_attachment" "t" {
					drg_id = "${oci_core_drg.t.id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					display_name = "-tf-drg-attachment"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-drg-attachment"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.DrgAttachmentLifecycleStateAttached)),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, "oci_core_drg.t", "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: core/virtualNetwork
func TestResourceCoreDrgAttachmentTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreDrgAttachmentTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreDrgAttachmentTestSuite))
}
