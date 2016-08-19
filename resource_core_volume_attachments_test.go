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

type ResourceCoreVolumeAttachmentsTestSuite struct {
	suite.Suite
	Client       *MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVolumeAttachmentsTestSuite) SetupTest() {
	s.Client = &MockClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_volume_attachments" "t" {
      availability_domain = "availability_domain"
      compartment_id = "compartment_id"
      limit = 1
      page = "page"
      instance_id = "instance_id"
      volume_id = "volume_id"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_volume_attachments.t"
}

func (s *ResourceCoreVolumeAttachmentsTestSuite) TestReadVolumeAttachments() {
	opts := []baremetal.Options{
		baremetal.Options{
			AvailabilityDomain: "availability_domain",
			Limit:              1,
			Page:               "page",
			InstanceID:         "instance_id",
			VolumeID:           "volume_id",
		},
	}

	s.Client.On(
		"ListVolumeAttachments",
		"compartment_id",
		opts,
	).Return(
		&baremetal.VolumeAttachmentList{
			VolumeAttachments: []baremetal.VolumeAttachment{
				baremetal.VolumeAttachment{
					AttachmentType:     "attachment_type",
					AvailabilityDomain: "availability_domain",
					CompartmentID:      "compartment_id",
					DisplayName:        "display_name",
					ID:                 "id1",
					InstanceID:         "instance_id",
					State:              baremetal.ResourceAttached,
					TimeCreated:        baremetal.Time{Time: time.Now()},
					VolumeID:           "volume_id",
				},
				baremetal.VolumeAttachment{
					AttachmentType:     "attachment_type",
					AvailabilityDomain: "availability_domain",
					CompartmentID:      "compartment_id",
					DisplayName:        "display_name",
					ID:                 "id2",
					InstanceID:         "instance_id",
					State:              baremetal.ResourceAttached,
					TimeCreated:        baremetal.Time{Time: time.Now()},
					VolumeID:           "volume_id",
				},
			},
		},
		nil,
	)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.1.id", "id2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVolumeAttachments", "compartment_id", opts)
}

func TestResourceCoreVolumeAttachmentsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeAttachmentsTestSuite))
}
