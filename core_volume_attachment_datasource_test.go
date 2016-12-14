package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeAttachmentsTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVolumeAttachmentsTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
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
	opts := &baremetal.ListVolumeAttachmentsOptions{}
	opts.AvailabilityDomain = "availability_domain"
	opts.Limit = 1
	opts.Page = "page"
	opts.InstanceID = "instance_id"
	opts.VolumeID = "volume_id"

	s.Client.On(
		"ListVolumeAttachments",
		"compartment_id",
		opts,
	).Return(
		&baremetal.ListVolumeAttachments{
			VolumeAttachments: []baremetal.VolumeAttachment{
				{
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
				{
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
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.#", "2"),
				),
			},
		},
	},
	)
	s.Client.AssertCalled(s.T(), "ListVolumeAttachments", "compartment_id", opts)
}

func (s *ResourceCoreVolumeAttachmentsTestSuite) TestReadVolumeAttachmentsWithPaging() {
	opts := &baremetal.ListVolumeAttachmentsOptions{}
	opts.AvailabilityDomain = "availability_domain"
	opts.Limit = 1
	opts.Page = "page"
	opts.InstanceID = "instance_id"
	opts.VolumeID = "volume_id"

	s.Client.On(
		"ListVolumeAttachments",
		"compartment_id",
		opts,
	).Return(
		&baremetal.ListVolumeAttachments{
			ResourceContainer: baremetal.ResourceContainer{
				NextPage: "nextpage",
			},
			VolumeAttachments: []baremetal.VolumeAttachment{
				{
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
				{
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

	opts2 := &baremetal.ListVolumeAttachmentsOptions{}
	opts2.AvailabilityDomain = "availability_domain"
	opts2.Limit = 1
	opts2.Page = "nextpage"
	opts2.InstanceID = "instance_id"
	opts2.VolumeID = "volume_id"

	s.Client.On(
		"ListVolumeAttachments",
		"compartment_id",
		opts2,
	).Return(
		&baremetal.ListVolumeAttachments{
			VolumeAttachments: []baremetal.VolumeAttachment{
				{
					AttachmentType:     "attachment_type",
					AvailabilityDomain: "availability_domain",
					CompartmentID:      "compartment_id",
					DisplayName:        "display_name",
					ID:                 "id3",
					InstanceID:         "instance_id",
					State:              baremetal.ResourceAttached,
					TimeCreated:        baremetal.Time{Time: time.Now()},
					VolumeID:           "volume_id",
				},
				{
					AttachmentType:     "attachment_type",
					AvailabilityDomain: "availability_domain",
					CompartmentID:      "compartment_id",
					DisplayName:        "display_name",
					ID:                 "id4",
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
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.availability_domain", "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_attachments.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVolumeAttachments", "compartment_id", opts2)
}

func TestResourceCoreVolumeAttachmentsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeAttachmentsTestSuite))
}
