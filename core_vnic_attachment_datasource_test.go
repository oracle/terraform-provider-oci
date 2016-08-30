package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVnicAttachmentsTestSuite struct {
	suite.Suite
	Client       *client.MockClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVnicAttachmentsTestSuite) SetupTest() {
	s.Client = &client.MockClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_vnic_attachments" "s" {
      compartment_id = "compartmentid"
      availability_domain = "availabilityid"
      vnic_id = "vnicid"
      instance_id = "instanceid"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_vnic_attachments.s"

}

func (s *ResourceCoreVnicAttachmentsTestSuite) TestResourceReadCoreVnicAttachments() {
	opts := []baremetal.Options{
		baremetal.Options{
			AvailabilityDomain: "availabilityid",
			VnicID:             "vnicid",
			InstanceID:         "instanceid",
		},
	}

	s.Client.On(
		"ListVnicAttachments",
		"compartmentid",
		opts,
	).Return(
		&baremetal.ListVnicAttachments{
			Attachments: []baremetal.VnicAttachment{
				baremetal.VnicAttachment{
					ID:                 "id1",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "att1",
					InstanceID:         "instanceid",
					State:              baremetal.ResourceAttached,
					SubnetID:           "subnetid",
					VnicID:             "vnicid",
					TimeCreated:        time.Now(),
				},
				baremetal.VnicAttachment{
					ID:                 "id2",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "att2",
					InstanceID:         "instanceid",
					State:              baremetal.ResourceAttached,
					SubnetID:           "subnetid",
					VnicID:             "vnicid",
					TimeCreated:        time.Now().Add(crud.FiveMinutes),
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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_id", "vnicid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instance_id", "instanceid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.0.availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.#", "2"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVnicAttachments", "compartmentid", opts)

}

func (s *ResourceCoreVnicAttachmentsTestSuite) TestResourceReadCoreVnicAttachmentsWithPaging() {
	opts := []baremetal.Options{
		baremetal.Options{
			AvailabilityDomain: "availabilityid",
			VnicID:             "vnicid",
			InstanceID:         "instanceid",
		},
	}

	s.Client.On(
		"ListVnicAttachments",
		"compartmentid",
		opts,
	).Return(
		&baremetal.ListVnicAttachments{
			ResourceContainer: baremetal.ResourceContainer{
				NextPage: "nextpage",
			},
			Attachments: []baremetal.VnicAttachment{
				baremetal.VnicAttachment{
					ID:                 "id1",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "att1",
					InstanceID:         "instanceid",
					State:              baremetal.ResourceAttached,
					SubnetID:           "subnetid",
					VnicID:             "vnicid",
					TimeCreated:        time.Now(),
				},
				baremetal.VnicAttachment{
					ID:                 "id2",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "att2",
					InstanceID:         "instanceid",
					State:              baremetal.ResourceAttached,
					SubnetID:           "subnetid",
					VnicID:             "vnicid",
					TimeCreated:        time.Now().Add(crud.FiveMinutes),
				},
			},
		},
		nil,
	)

	opts2 := []baremetal.Options{
		baremetal.Options{
			AvailabilityDomain: "availabilityid",
			VnicID:             "vnicid",
			InstanceID:         "instanceid",
			Page:               "nextpage",
		},
	}

	s.Client.On(
		"ListVnicAttachments",
		"compartmentid",
		opts2,
	).Return(
		&baremetal.ListVnicAttachments{
			Attachments: []baremetal.VnicAttachment{
				baremetal.VnicAttachment{
					ID:                 "id3",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "att1",
					InstanceID:         "instanceid",
					State:              baremetal.ResourceAttached,
					SubnetID:           "subnetid",
					VnicID:             "vnicid",
					TimeCreated:        time.Now(),
				},
				baremetal.VnicAttachment{
					ID:                 "id4",
					AvailabilityDomain: "availabilityid",
					CompartmentID:      "compartmentid",
					DisplayName:        "att2",
					InstanceID:         "instanceid",
					State:              baremetal.ResourceAttached,
					SubnetID:           "subnetid",
					VnicID:             "vnicid",
					TimeCreated:        time.Now().Add(crud.FiveMinutes),
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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartmentid"),
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_id", "vnicid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instance_id", "instanceid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.0.availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.0.id", "id1"),
					//	resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListVnicAttachments", "compartmentid", opts2)

}

func TestResourceCoreVnicAttachmentsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVnicAttachmentsTestSuite))
}
