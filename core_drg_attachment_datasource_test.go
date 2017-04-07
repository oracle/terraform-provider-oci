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

type CoreDrgAttachmentDatasourceTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *CoreDrgAttachmentDatasourceTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_drg_attachments" "t" {
      compartment_id = "compartment_id"
			drg_id = "drg_id"
      limit = 1
      page = "page"
			vcn_id = "vcn_id"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_core_drg_attachments.t"
}

func (s *CoreDrgAttachmentDatasourceTestSuite) TestReadDrgAttachments() {
	opts := &baremetal.ListDrgAttachmentsOptions{}
	opts.DrgID = "drg_id"
	opts.Limit = 1
	opts.Page = "page"
	opts.VcnID = "vcn_id"

	s.Client.On(
		"ListDrgAttachments",
		"compartment_id",
		opts,
	).Return(
		&baremetal.ListDrgAttachments{
			DrgAttachments: []baremetal.DrgAttachment{
				{
					CompartmentID: "compartment_id",
					DrgID:         "drg_id",
					DisplayName:   "display_name",
					ID:            "id1",
					State:         baremetal.ResourceAttached,
					TimeCreated:   baremetal.Time{Time: time.Now()},
					VcnID:         "vcn_id",
				},
				{
					CompartmentID: "compartment_id",
					DrgID:         "drg_id",
					DisplayName:   "display_name",
					ID:            "id2",
					State:         baremetal.ResourceAttached,
					TimeCreated:   baremetal.Time{Time: time.Now()},
					VcnID:         "vcn_id",
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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.0.compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.#", "2"),
				),
			},
		},
	},
	)
}

func (s *CoreDrgAttachmentDatasourceTestSuite) TestReadPagedDrgAttachments() {
	opts := &baremetal.ListDrgAttachmentsOptions{}
	opts.DrgID = "drg_id"
	opts.Limit = 1
	opts.Page = "page"
	opts.VcnID = "vcn_id"

	res := &baremetal.ListDrgAttachments{}
	res.NextPage = "nextpage"
	res.DrgAttachments = []baremetal.DrgAttachment{
		{
			CompartmentID: "compartment_id",
			DrgID:         "drg_id",
			DisplayName:   "display_name",
			ID:            "id1",
			State:         baremetal.ResourceAttached,
			TimeCreated:   baremetal.Time{Time: time.Now()},
			VcnID:         "vcn_id",
		},
		{
			CompartmentID: "compartment_id",
			DrgID:         "drg_id",
			DisplayName:   "display_name",
			ID:            "id2",
			State:         baremetal.ResourceAttached,
			TimeCreated:   baremetal.Time{Time: time.Now()},
			VcnID:         "vcn_id",
		},
	}

	s.Client.On("ListDrgAttachments", "compartment_id", opts).Return(res, nil)

	opts2 := &baremetal.ListDrgAttachmentsOptions{}
	opts2.DrgID = "drg_id"
	opts2.Limit = 1
	opts2.Page = "nextpage"
	opts2.VcnID = "vcn_id"

	s.Client.On(
		"ListDrgAttachments",
		"compartment_id",
		opts2,
	).Return(
		&baremetal.ListDrgAttachments{
			DrgAttachments: []baremetal.DrgAttachment{
				{
					CompartmentID: "compartment_id",
					DrgID:         "drg_id",
					DisplayName:   "display_name",
					ID:            "id3",
					State:         baremetal.ResourceAttached,
					TimeCreated:   baremetal.Time{Time: time.Now()},
					VcnID:         "vcn_id",
				},
				{
					CompartmentID: "compartment_id",
					DrgID:         "drg_id",
					DisplayName:   "display_name",
					ID:            "id4",
					State:         baremetal.ResourceAttached,
					TimeCreated:   baremetal.Time{Time: time.Now()},
					VcnID:         "vcn_id",
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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "limit", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "page", "page"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.0.compartment_id", "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "drg_attachments.#", "4"),
				),
			},
		},
	},
	)

	s.Client.AssertCalled(s.T(), "ListDrgAttachments", "compartment_id", opts2)
}

func TestCoreDrgAttachmentDatasourceTestSuite(t *testing.T) {
	suite.Run(t, new(CoreDrgAttachmentDatasourceTestSuite))
}
