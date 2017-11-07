// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	baremetal "github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVolumeAttachmentTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVolumeAttachmentTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + instanceConfig + `
	resource "oci_core_volume" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-volume"
	}
	resource "oci_core_volume_attachment" "t" {
		compartment_id = "${var.compartment_id}"
		instance_id = "${oci_core_instance.t.id}"
		volume_id = "${oci_core_volume.t.id}"
		attachment_type = "iscsi"
	}`
	s.ResourceName = "data.oci_core_volume_attachments.t"
}

func (s *DatasourceCoreVolumeAttachmentTestSuite) TestAccDatasourceCoreVolumeAttachment_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_core_volume_attachments" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
					volume_id = "${oci_core_volume.t.id}"
					filter {
						name = "id"
						values = ["${oci_core_volume_attachment.t.id}"]
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.attachment_type", "iscsi"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.state", "ATTACHED"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.instance_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.volume_id"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreVolumeAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVolumeAttachmentTestSuite))
}
