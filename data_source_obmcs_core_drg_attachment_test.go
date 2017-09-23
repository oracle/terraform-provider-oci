// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreDrgAttachmentTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreDrgAttachmentTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "network_name"
	}
	resource "oci_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}
	resource "oci_core_drg_attachment" "t" {
		compartment_id = "${var.compartment_id}"
		drg_id = "${oci_core_drg.t.id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		display_name = "-tf-drg-attachment"
	}`
	s.ResourceName = "data.oci_core_drg_attachments.t"
}

func (s *DatasourceCoreDrgAttachmentTestSuite) TestAccDatasourceCoreDrgAttachment_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config:            s.Config,
			},
			// todo: investigate, related issue with TestAccDatasourceCoreIPConnections_basic
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config + `
				data "oci_core_drg_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					drg_id = "${oci_core_drg.t.id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					limit = 1
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.0.display_name", "-tf-drg-attachment"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_attachments.0.id"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreDrgAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreDrgAttachmentTestSuite))
}
