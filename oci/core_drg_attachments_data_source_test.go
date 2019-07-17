// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreDrgAttachmentTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreDrgAttachmentTestSuite) SetupTest() {
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-vcn"
	}
	resource "oci_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-drg"
	}
	resource "oci_core_drg_attachment" "t" {
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
				Config: s.Config + `
				data "oci_core_drg_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					drg_id = "${oci_core_drg.t.id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
					
					filter {
						name = "id"
						values = ["${oci_core_drg_attachment.t.id}"]
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.0.display_name", "-tf-drg-attachment"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_attachments.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_attachments.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_attachments.0.drg_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drg_attachments.0.vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.0.state", string(core.DrgAttachmentLifecycleStateAttached)),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreDrgAttachmentTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreDrgAttachmentTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreDrgAttachmentTestSuite))
}
