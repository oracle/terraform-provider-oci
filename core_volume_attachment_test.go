package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeAttachmentTestSuite struct {
	suite.Suite
	Client       *client.MockClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.VolumeAttachment
	DetachedRes  *baremetal.VolumeAttachment
	Opts         []baremetal.Options
}

func (s *ResourceCoreVolumeAttachmentTestSuite) SetupTest() {
	s.Client = &client.MockClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
		resource "baremetal_core_volume_attachment" "t" {
			attachment_type = "attachment_type"
			compartment_id = "compartment_id"
			instance_id = "instance_id"
			volume_id = "volume_id"
		}
	`
	s.Config += testProviderConfig

	s.ResourceName = "baremetal_core_volume_attachment.t"
	s.Res = &baremetal.VolumeAttachment{
		AttachmentType:     "attachment_type",
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "display_name",
		ID:                 "id",
		InstanceID:         "instance_id",
		State:              baremetal.ResourceAttached,
		TimeCreated:        s.TimeCreated,
		VolumeID:           "volume_id",
		ETag:               "etag",
		OPCRequestID:       "opc_request_id",
	}

	s.DetachedRes = &baremetal.VolumeAttachment{
		AttachmentType:     "attachment_type",
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "display_name",
		ID:                 "id",
		InstanceID:         "instance_id",
		State:              baremetal.ResourceDetached,
		TimeCreated:        s.TimeCreated,
		VolumeID:           "volume_id",
		ETag:               "etag",
		OPCRequestID:       "opc_request_id",
	}

	// opts := baremetal.Options{}
	s.Opts = []baremetal.Options(nil)
	s.Client.On(
		"AttachVolume",
		"compartment_id",
		"instance_id",
		"attachment_type",
		"volume_id",
		s.Opts).Return(s.Res, nil)
	s.Client.On("DetachVolume", "id").Return(nil)
}

func (s *ResourceCoreVolumeAttachmentTestSuite) TestCreateResourceCoreVolumeAttachment() {
	s.Client.On("GetVolumeAttachment", "id", []baremetal.Options(nil)).Return(s.Res, nil).Times(2)
	s.Client.On("GetVolumeAttachment", "id", []baremetal.Options(nil)).Return(s.DetachedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "attachment_type", s.Res.AttachmentType),
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", s.Res.AvailabilityDomain),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "instance_id", s.Res.InstanceID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_id", s.Res.VolumeID),
				),
			},
		},
	})
}

func (s *ResourceCoreVolumeAttachmentTestSuite) TestDetachVolume() {
	s.Client.On("GetVolumeAttachment", "id", []baremetal.Options(nil)).Return(s.Res, nil).Times(2)
	s.Client.On("GetVolumeAttachment", "id", []baremetal.Options(nil)).Return(s.DetachedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DetachVolume", "id")
}

func TestResourceCoreVolumeAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeAttachmentTestSuite))
}
