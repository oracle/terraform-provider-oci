// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type CoreDrgAttachmentDatasourceTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *CoreDrgAttachmentDatasourceTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
	resource "baremetal_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "network_name"
	}
	resource "baremetal_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}
	resource "baremetal_core_drg_attachment" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
		drg_id = "${baremetal_core_drg.t.id}"
		vcn_id = "${baremetal_core_virtual_network.t.id}"
	}
    data "baremetal_core_drg_attachments" "t" {
        compartment_id = "${var.compartment_id}"
	drg_id = "${baremetal_core_drg.t.id}"
        limit = 1
	vcn_id = "${baremetal_core_virtual_network.t.id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_drg_attachments.t"
}

func (s *CoreDrgAttachmentDatasourceTestSuite) TestReadDrgAttachments() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.0.display_name", "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_attachments.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.#", "1"),
				),
			},
		},
	},
	)
}

func TestCoreDrgAttachmentDatasourceTestSuite(t *testing.T) {
	suite.Run(t, new(CoreDrgAttachmentDatasourceTestSuite))
}
