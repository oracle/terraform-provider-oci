// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/client/mocks"

	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type ResourceCoreInstanceTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Instance
	DeletedRes   *baremetal.Instance
}

func (s *ResourceCoreInstanceTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

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
		resource "baremetal_core_instance" "t" {
			availability_domain = "availability_domain"
			compartment_id = "compartment_id"
			display_name = "display_name"
      image = "imageid"
      shape = "shapeid"
      subnet_id = "subnetid"
      metadata {
        ssh_authorized_keys = "mypublickey"
      }
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_core_instance.t"
	s.Res = &baremetal.Instance{
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "display_name",
		ID:                 "id",
		ImageID:            "imageid",
		Metadata: map[string]string{
			"ssh_authorized_keys": "mypublickey",
		},
		Region:      "region",
		Shape:       "shapeid",
		State:       baremetal.ResourceRunning,
		TimeCreated: s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.DeletedRes = &baremetal.Instance{}
	*s.DeletedRes = *s.Res
	s.DeletedRes.State = baremetal.ResourceTerminated

	opts := &baremetal.LaunchInstanceOptions{}
	opts.DisplayName = "display_name"
	opts.Metadata = s.Res.Metadata
	s.Client.On(
		"LaunchInstance",
		s.Res.AvailabilityDomain,
		s.Res.CompartmentID,
		s.Res.ImageID,
		s.Res.Shape,
		"subnetid",
		opts).Return(s.Res, nil)
	s.Client.On("TerminateInstance", s.Res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)

	listVnicOpts := &baremetal.ListVnicAttachmentsOptions{}
	listVnicOpts.AvailabilityDomain = s.Res.AvailabilityDomain
	listVnicOpts.InstanceID = s.Res.ID

	listVnicOpts2 := &baremetal.ListVnicAttachmentsOptions{}
	listVnicOpts2.AvailabilityDomain = "new_availability_domain"
	listVnicOpts2.InstanceID = "new_id"

	vnic := &baremetal.Vnic{}
	vnic.PublicIPAddress = "0.0.0.0"
	vnic.PrivateIPAddress = "0.0.0.0"
	vnicAttachment := &baremetal.VnicAttachment{
		ID:                 "id1",
		AvailabilityDomain: "availabilityid",
		CompartmentID:      "compartmentid",
		DisplayName:        "att1",
		InstanceID:         "instanceid",
		State:              baremetal.ResourceAttached,
		SubnetID:           "subnetid",
		VnicID:             "vnicid",
		TimeCreated:        time.Now(),
	}
	vnicList := &baremetal.ListVnicAttachments{
		Attachments: []baremetal.VnicAttachment{
			*vnicAttachment,
		},
	}
	s.Client.On("ListVnicAttachments", s.Res.CompartmentID, listVnicOpts).Return(vnicList, nil)
	s.Client.On("ListVnicAttachments", s.Res.CompartmentID, listVnicOpts2).Return(vnicList, nil)
	s.Client.On("GetVnic", "vnicid").Return(vnic, nil)
}

func (s *ResourceCoreInstanceTestSuite) TestCreateResourceCoreInstance() {
	s.Client.On("GetInstance", "id").Return(s.Res, nil).Twice()
	s.Client.On("GetInstance", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", s.Res.AvailabilityDomain),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "image", s.Res.ImageID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
					resource.TestCheckResourceAttr(s.ResourceName, "public_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ip", "0.0.0.0"),
				),
			},
		},
	})
}

func (s *ResourceCoreInstanceTestSuite) TestCreateResourceCoreInstanceWithoutDisplayName() {
	s.Client.On("GetInstance", "id").Return(s.Res, nil).Twice()
	s.Client.On("GetInstance", "id").Return(s.DeletedRes, nil)

	s.Config = `
		resource "baremetal_core_instance" "t" {
			availability_domain = "availability_domain"
			compartment_id = "compartment_id"
      image = "imageid"
      shape = "shapeid"
      subnet_id = "subnetid"
      metadata {
        ssh_authorized_keys = "mypublickey"
      }
		}
	`
	s.Config += testProviderConfig

	opts := &baremetal.LaunchInstanceOptions{}
	opts.Metadata = s.Res.Metadata

	s.Client.On(
		"LaunchInstance",
		s.Res.AvailabilityDomain,
		s.Res.CompartmentID,
		s.Res.ImageID,
		s.Res.Shape,
		"subnetid",
		opts).Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
				),
			},
		},
	})
}

func (s ResourceCoreInstanceTestSuite) TestUpdateInstanceDisplayName() {
	s.Client.On("GetInstance", "id").Return(s.Res, nil).Times(2)

	config := `
		resource "baremetal_core_instance" "t" {
			availability_domain = "availability_domain"
			compartment_id = "compartment_id"
      image = "imageid"
      shape = "shapeid"
      subnet_id = "subnetid"
      display_name = "new_display_name"
      metadata {
        ssh_authorized_keys = "mypublickey"
      }
		}
	`
	config += testProviderConfig

	res := &baremetal.Instance{
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "new_display_name",
		ID:                 "id",
		ImageID:            "imageid",
		Metadata: map[string]string{
			"ssh_authorized_keys": "mypublickey",
		},
		Region:      "region",
		Shape:       "shapeid",
		State:       baremetal.ResourceRunning,
		TimeCreated: s.TimeCreated,
	}
	res.ETag = "etag"
	res.RequestID = "opcrequestid"

	opts := &baremetal.UpdateOptions{}
	opts.DisplayName = "new_display_name"
	s.Client.On("UpdateInstance", "id", opts).Return(res, nil)
	s.Client.On("GetInstance", "id").Return(res, nil).Times(2)
	s.Client.On("GetInstance", "id").Return(s.DeletedRes, nil).Times(2)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", res.DisplayName),
				),
			},
		},
	})
}

func (s ResourceCoreInstanceTestSuite) TestUpdateAvailabilityDomainForcesNewInstance() {
	s.Client.On("GetInstance", "id").Return(s.Res, nil).Times(2)

	config := `
		resource "baremetal_core_instance" "t" {
			availability_domain = "new_availability_domain"
			compartment_id = "compartment_id"
			display_name = "display_name"
      image = "imageid"
      shape = "shapeid"
      subnet_id = "subnetid"
      metadata {
        ssh_authorized_keys = "mypublickey"
      }
		}
	`

	config += testProviderConfig

	res := &baremetal.Instance{
		AvailabilityDomain: "new_availability_domain",
		CompartmentID:      "compartment_id",
		DisplayName:        "display_name",
		ID:                 "new_id",
		ImageID:            "imageid",
		Metadata: map[string]string{
			"ssh_authorized_keys": "mypublickey",
		},
		Region:      "region",
		Shape:       "shapeid",
		State:       baremetal.ResourceRunning,
		TimeCreated: s.TimeCreated,
	}
	res.ETag = "etag"
	res.RequestID = "opcrequestid"

	opts := &baremetal.LaunchInstanceOptions{}
	opts.DisplayName = "display_name"
	opts.Metadata = s.Res.Metadata
	s.Client.On(
		"LaunchInstance",
		res.AvailabilityDomain,
		res.CompartmentID,
		res.ImageID,
		res.Shape,
		"subnetid",
		opts).Return(res, nil)

	s.Client.On("GetInstance", s.Res.ID).Return(s.DeletedRes, nil)
	s.Client.On("GetInstance", res.ID).Return(res, nil).Times(2)
	s.Client.On("TerminateInstance", res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)
	s.Client.On("GetInstance", "new_id").Return(s.DeletedRes, nil).Twice()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", res.AvailabilityDomain),
				),
			},
		},
	})
}

func (s *ResourceCoreInstanceTestSuite) TestTerminateInstance() {
	s.Client.On("GetInstance", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetInstance", "id").Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "TerminateInstance", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestIsStatefulResource(t *testing.T) {
	var sr crud.StatefulResource
	sr = &InstanceResourceCrud{}
	if sr == nil {
		t.Fail()
	}
}

func TestResourceCoreInstanceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInstanceTestSuite))
}
