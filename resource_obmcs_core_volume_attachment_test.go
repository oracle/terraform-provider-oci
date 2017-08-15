// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeAttachmentTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.VolumeAttachment
	DetachedRes  *baremetal.VolumeAttachment
}

func (s *ResourceCoreVolumeAttachmentTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

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

	s.ResourceName = "baremetal_core_volume_attachment.t"

}

func (s *ResourceCoreVolumeAttachmentTestSuite) TestCreateResourceCoreVolumeAttachment() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "attachment_type", "iscsi"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAttached),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "chap_secret"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "chap_username"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ipv4"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "iqn"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "port"),
				),
			},
		},
	})
}

func (s *ResourceCoreVolumeAttachmentTestSuite) TestDetachVolume() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

}

func TestResourceCoreVolumeAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeAttachmentTestSuite))
}
