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
	Client       mockableClient
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
	s.Res = &baremetal.VolumeAttachment{
		AttachmentType:     "iscsi",
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "display_name",
		ID:                 "id",
		InstanceID:         "instance_id",
		State:              baremetal.ResourceAttached,
		TimeCreated:        s.TimeCreated,
		VolumeID:           "volume_id",
		CHAPSecret:         "chap_secret",
		CHAPUsername:       "chap_username",
		IPv4:               "ipv4",
		IQN:                "iqn",
		Port:               12345,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.DetachedRes = &baremetal.VolumeAttachment{
		AttachmentType:     "iscsi",
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "display_name",
		ID:                 "id",
		InstanceID:         "instance_id",
		State:              baremetal.ResourceDetached,
		TimeCreated:        s.TimeCreated,
		VolumeID:           "volume_id",
	}
	s.DetachedRes.ETag = "etag"
	s.DetachedRes.RequestID = "opcrequestid"

	s.Client.On(
		"AttachVolume",
		"iscsi",
		"instance_id",
		"volume_id",
		(*baremetal.CreateOptions)(nil)).Return(s.Res, nil)
	s.Client.On("DetachVolume", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreVolumeAttachmentTestSuite) TestCreateResourceCoreVolumeAttachment() {
	s.Client.On("GetVolumeAttachment", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVolumeAttachment", "id").Return(s.DetachedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "attachment_type", s.Res.AttachmentType),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "instance_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
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
	s.Client.On("GetVolumeAttachment", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVolumeAttachment", "id").Return(s.DetachedRes, nil)

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

	s.Client.AssertCalled(s.T(), "DetachVolume", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreVolumeAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeAttachmentTestSuite))
}
