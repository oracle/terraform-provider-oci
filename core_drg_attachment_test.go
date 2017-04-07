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
)

type ResourceCoreDrgAttachmentTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.DrgAttachment
	DetachedRes  *baremetal.DrgAttachment
}

func (s *ResourceCoreDrgAttachmentTestSuite) SetupTest() {
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
		resource "baremetal_core_drg_attachment" "t" {
			compartment_id = "compartment_id"
			display_name = "display_name"
			drg_id = "drg_id"
			vcn_id = "vcn_id"
		}
	`
	s.Config += testProviderConfig

	s.ResourceName = "baremetal_core_drg_attachment.t"
	s.Res = &baremetal.DrgAttachment{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		DrgID:         "drg_id",
		ID:            "id",
		State:         baremetal.ResourceAttached,
		TimeCreated:   s.TimeCreated,
		VcnID:         "vcn_id",
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.DetachedRes = &baremetal.DrgAttachment{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		DrgID:         "drg_id",
		ID:            "id",
		State:         baremetal.ResourceDetached,
		TimeCreated:   s.TimeCreated,
		VcnID:         "vcn_id",
	}
	s.DetachedRes.ETag = "etag"
	s.DetachedRes.RequestID = "opcrequestid"

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On(
		"CreateDrgAttachment",
		"drg_id",
		"vcn_id",
		opts).Return(s.Res, nil)
	s.Client.On("DeleteDrgAttachment", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreDrgAttachmentTestSuite) TestCreateResourceCoreDrgAttachment() {
	s.Client.On("GetDrgAttachment", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetDrgAttachment", "id").Return(s.DetachedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_id", s.Res.DrgID),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
					resource.TestCheckResourceAttr(s.ResourceName, "vcn_id", s.Res.VcnID),
				),
			},
		},
	})
}

func (s *ResourceCoreDrgAttachmentTestSuite) TestDetachVolume() {
	s.Client.On("GetDrgAttachment", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetDrgAttachment", "id").Return(s.DetachedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeleteDrgAttachment", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreDrgAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDrgAttachmentTestSuite))
}
