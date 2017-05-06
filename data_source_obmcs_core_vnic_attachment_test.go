// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"



	"github.com/oracle/terraform-provider-baremetal/crud"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVnicAttachmentsTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVnicAttachmentsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
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
	opts := &baremetal.ListVnicAttachmentsOptions{}
	opts.AvailabilityDomain = "availabilityid"
	opts.VnicID = "vnicid"
	opts.InstanceID = "instanceid"

	s.Client.On(
		"ListVnicAttachments",
		"compartmentid",
		opts,
	).Return(
		&baremetal.ListVnicAttachments{
			Attachments: []baremetal.VnicAttachment{
				{
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
				{
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
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
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
	opts := &baremetal.ListVnicAttachmentsOptions{}
	opts.AvailabilityDomain = "availabilityid"
	opts.VnicID = "vnicid"
	opts.InstanceID = "instanceid"

	res := &baremetal.ListVnicAttachments{}
	res.NextPage = "nextpage"
	res.Attachments = []baremetal.VnicAttachment{
		{
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
		{
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
	}

	s.Client.On(
		"ListVnicAttachments",
		"compartmentid",
		opts,
	).Return(res, nil)

	opts2 := &baremetal.ListVnicAttachmentsOptions{}
	opts2.AvailabilityDomain = "availabilityid"
	opts2.VnicID = "vnicid"
	opts2.InstanceID = "instanceid"
	opts2.Page = "nextpage"

	s.Client.On(
		"ListVnicAttachments",
		"compartmentid",
		opts2,
	).Return(
		&baremetal.ListVnicAttachments{
			Attachments: []baremetal.VnicAttachment{
				{
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
				{
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
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
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
