// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeAttachmentsTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVolumeAttachmentsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = instanceConfig + `
		resource "baremetal_core_volume" "t" {
			availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
			size_in_mbs = 262144
		}
		resource "baremetal_core_volume_attachment" "t" {
			attachment_type = "iscsi"
			compartment_id = "${var.compartment_id}"
			instance_id = "${baremetal_core_instance.t.id}"
			volume_id = "${baremetal_core_volume.t.id}"
		}
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_volume_attachments.t"
}

func (s *ResourceCoreVolumeAttachmentsTestSuite) TestReadVolumeAttachments() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				    data "baremetal_core_volume_attachments" "t" {
				      availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
				      compartment_id = "${var.compartment_id}"
				      limit = 1
				      instance_id = "${baremetal_core_instance.t.id}"
				      volume_id = "${baremetal_core_volume.t.id}"
				    }`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_attachments.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.#", "1"),
				),
			},
		},
	},
	)
}

func TestResourceCoreVolumeAttachmentsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeAttachmentsTestSuite))
}
