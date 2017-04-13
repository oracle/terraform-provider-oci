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

	"errors"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVirtualNetworkTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.VirtualNetwork
	DeletedRes   *baremetal.VirtualNetwork
	DeletingRes  *baremetal.VirtualNetwork
}

func (s *ResourceCoreVirtualNetworkTestSuite) SetupTest() {
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
		resource "baremetal_core_virtual_network" "t" {
			cidr_block = "cidr_block"
			compartment_id = "compartment_id"
			display_name = "display_name"
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_core_virtual_network.t"
	s.Res = &baremetal.VirtualNetwork{
		CidrBlock:             "cidr_block",
		CompartmentID:         "compartment_id",
		DefaultRouteTableID:   "default_route_table_id",
		DefaultSecurityListID: "default_security_list_id",
		DisplayName:           "display_name",
		ID:                    "id",
		State:                 baremetal.ResourceAvailable,
		TimeCreated:           s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.DeletingRes = &baremetal.VirtualNetwork{
		CidrBlock:             "cidr_block",
		CompartmentID:         "compartment_id",
		DefaultRouteTableID:   "default_route_table_id",
		DefaultSecurityListID: "default_security_list_id",
		DisplayName:           "display_name",
		ID:                    "id",
		State:                 baremetal.ResourceTerminating,
		TimeCreated:           s.TimeCreated,
	}

	s.DeletedRes = &baremetal.VirtualNetwork{
		CidrBlock:             "cidr_block",
		CompartmentID:         "compartment_id",
		DefaultRouteTableID:   "default_route_table_id",
		DefaultSecurityListID: "default_security_list_id",
		DisplayName:           "display_name",
		ID:                    "id",
		State:                 baremetal.ResourceTerminated,
		TimeCreated:           s.TimeCreated,
	}
	s.DeletedRes.ETag = "etag"
	s.DeletedRes.RequestID = "opcrequestid"

	opts := &baremetal.CreateVcnOptions{}
	opts.DisplayName = "display_name"
	s.Client.On(
		"CreateVirtualNetwork",
		"cidr_block",
		"compartment_id",
		opts).Return(s.Res, nil)
	s.Client.On("DeleteVirtualNetwork", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestCreateResourceCoreVirtualNetwork() {
	s.Client.On("GetVirtualNetwork", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVirtualNetwork", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", s.Res.CidrBlock),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "default_route_table_id", s.Res.DefaultRouteTableID),
					resource.TestCheckResourceAttr(s.ResourceName, "default_security_list_id", s.Res.DefaultSecurityListID),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
				),
			},
		},
	})
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestDeleteResourceCoreVirtualNetwork() {
	s.Client.On("GetVirtualNetwork", s.Res.ID).Return(s.Res, nil).Times(2)
	s.Client.On("GetVirtualNetwork", s.Res.ID).Return(s.DeletingRes, nil).Times(2)
	s.Client.On("GetVirtualNetwork", s.Res.ID).Return(nil, errors.New("blah blah does not exist"))

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
				),
			},
			{
				Config: testProviderConfig,
				Check: resource.ComposeTestCheckFunc(
					testNoInstanceState("baremetal_core_virtual_network"),
				),
			},
		},
	})
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestCreateResourceCoreVirtualNetworkWithoutDisplayName() {
	s.Client.On("GetVirtualNetwork", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVirtualNetwork", "id").Return(s.DeletedRes, nil)

	s.Config = `
		resource "baremetal_core_virtual_network" "t" {
			cidr_block = "cidr_block"
			compartment_id = "compartment_id"
		}
	`
	s.Config += testProviderConfig

	opts := &baremetal.CreateVcnOptions{}
	s.Client.On(
		"CreateVirtualNetwork",
		"cidr_block",
		"compartment_id", opts).Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
				),
			},
		},
	})
}

func (s ResourceCoreVirtualNetworkTestSuite) TestUpdateCidrBlockForcesNewVirtualNetwork() {
	// Step 1 uses the mocking in Setup plus the following two Get mocks to create
	// and then destroy the original resource.
	s.Client.On("GetVirtualNetwork", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVirtualNetwork", "id").Return(s.DeletedRes, nil)

	config := `
		resource "baremetal_core_virtual_network" "t" {
			cidr_block = "new_cidr_block"
			compartment_id = "compartment_id"
		}
  `
	config += testProviderConfig

	res := &baremetal.VirtualNetwork{
		CidrBlock:             "new_cidr_block",
		CompartmentID:         "compartment_id",
		DefaultRouteTableID:   "default_route_table_id",
		DefaultSecurityListID: "default_security_list_id",
		DisplayName:           "display_name",
		ID:                    "new_id",
		State:                 baremetal.ResourceAvailable,
		TimeCreated:           s.TimeCreated,
	}
	res.ETag = "etag"
	res.RequestID = "opcrequestid"

	// Step 2 creates a new resource and then uses the same Get pattern from
	// above in order to delete, this time with the newly created resource.
	opts := &baremetal.CreateVcnOptions{}
	s.Client.On(
		"CreateVirtualNetwork",
		res.CidrBlock,
		res.CompartmentID, opts).Return(res, nil)

	s.Client.On("DeleteVirtualNetwork", res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)
	deletedRes := &baremetal.VirtualNetwork{
		ID:    res.ID,
		State: baremetal.ResourceTerminated,
	}

	s.Client.On("GetVirtualNetwork", res.ID).Return(res, nil).Times(2)
	s.Client.On("GetVirtualNetwork", res.ID).Return(deletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", res.CidrBlock),
				),
			},
		},
	})
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestDeleteVirtualNetwork() {
	s.Client.On("GetVirtualNetwork", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetVirtualNetwork", "id").Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "DeleteVirtualNetwork", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreVirtualNetworkTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVirtualNetworkTestSuite))
}
