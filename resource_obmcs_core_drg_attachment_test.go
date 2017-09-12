// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreDrgAttachmentTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Res          *baremetal.DrgAttachment
	DetachedRes  *baremetal.DrgAttachment
}

func (s *ResourceCoreDrgAttachmentTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
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

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify a drg attachment can be created
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config + `
				resource "oci_core_drg_attachment" "t" {
					compartment_id = "${var.compartment_id}"
					drg_id = "${oci_core_drg.t.id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAttached),
				),
			},
			// verify drg attachment update
			{
				Config: s.Config + `
				resource "oci_core_drg_attachment" "t" {
					compartment_id = "${var.compartment_id}"
					drg_id = "${oci_core_drg.t.id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					display_name = "-tf-drg-attachment"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-drg-attachment"),
				),
			},
		},
	})
}

func TestResourceCoreDrgAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDrgAttachmentTestSuite))
}
