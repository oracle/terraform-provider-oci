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

type ResourceCoreVirtualNetworkTestSuite struct {
	suite.Suite
	Client       *client.MockClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.VirtualNetwork
	DeletedRes   *baremetal.VirtualNetwork
	Opts         []baremetal.Options
}

func (s *ResourceCoreVirtualNetworkTestSuite) SetupTest() {
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
		resource "baremetal_core_virtual_network" "t" {
			cidr_block = "cidr_block"
			compartment_id = "compartment_id"
			display_name = "display_name"
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_core_virtual_network.t"
	s.Res = &baremetal.VirtualNetwork{
		CidrBlock: "cidr_block",
		CompartmentID:      "compartment_id",
		DefaultRoutingTableID:      "default_routing_table_id",
		DefaultSecurityListID:      "default_security_list_id",
		DisplayName:        "display_name",
		ID:                 "id",
		State:              baremetal.ResourceAvailable,
		TimeCreated:        s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.DeletedRes = &baremetal.VirtualNetwork{
		CidrBlock: "cidr_block",
		CompartmentID:      "compartment_id",
		DefaultRoutingTableID:      "default_routing_table_id",
		DefaultSecurityListID:      "default_security_list_id",
		DisplayName:        "display_name",
		ID:                 "id",
		State:              baremetal.ResourceTerminated,
		TimeCreated:        s.TimeCreated,
	}
	s.DeletedRes.ETag = "etag"
	s.DeletedRes.RequestID = "opcrequestid"

	opts := baremetal.Options{DisplayName: "display_name"}
	s.Opts = []baremetal.Options{opts}
	s.Client.On(
		"CreateVirtualNetwork",
		"cidr_block",
		"compartment_id",
		s.Opts).Return(s.Res, nil)
	s.Client.On("DeleteVirtualNetwork", "id").Return(nil)
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestCreateResourceCoreVirtualNetwork() {
	s.Client.On("GetVirtualNetwork", "id", []baremetal.Options(nil)).Return(s.Res, nil).Times(2)
	s.Client.On("GetVirtualNetwork", "id", []baremetal.Options(nil)).Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", s.Res.CidrBlock),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "default_routing_table_id", s.Res.DefaultRoutingTableID),
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

func (s *ResourceCoreVirtualNetworkTestSuite) TestCreateResourceCoreVirtualNetworkWithoutDisplayName() {
	s.Client.On("GetVirtualNetwork", "id", []baremetal.Options(nil)).Return(s.Res, nil)

	s.Config = `
		resource "baremetal_core_virtual_network" "t" {
			cidr_block = "cidr_block"
			compartment_id = "compartment_id"
		}
	`
	s.Config += testProviderConfig

	opts := baremetal.Options{}
	s.Client.On(
		"CreateVirtualNetwork",
		"cidr_block",
		"compartment_id", []baremetal.Options{opts}).Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
				),
			},
		},
	})
}

func (s ResourceCoreVirtualNetworkTestSuite) TestUpdateCidrBlockForcesNewVirtualNetwork() {
	s.Client.On("GetVirtualNetwork", "id", []baremetal.Options(nil)).Return(s.Res, nil)

	config := `
		resource "baremetal_core_virtual_network" "t" {
			cidr_block = "new_cidr_block"
			compartment_id = "compartment_id"
		}
  `
	config += testProviderConfig

	res := &baremetal.VirtualNetwork{
		CidrBlock: "new_cidr_block",
		CompartmentID:      "compartment_id",
		DefaultRoutingTableID:      "default_routing_table_id",
		DefaultSecurityListID:      "default_security_list_id",
		DisplayName:        "display_name",
		ID:                 "new_id",
		State:              baremetal.ResourceAvailable,
		TimeCreated:        s.TimeCreated,
	}
	res.ETag = "etag"
	res.RequestID = "opcrequestid"

	opts := baremetal.Options{}
	s.Client.On(
		"CreateVirtualNetwork",
		res.CidrBlock,
		res.CompartmentID, []baremetal.Options{opts}).Return(res, nil)

	s.Client.On("GetVirtualNetwork", res.ID, []baremetal.Options(nil)).Return(res, nil)
	s.Client.On("DeleteVirtualNetwork", res.ID).Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "cidr_block", res.CidrBlock),
				),
			},
		},
	})
}

func (s *ResourceCoreVirtualNetworkTestSuite) TestDeleteVirtualNetwork() {
	s.Client.On("GetVirtualNetwork", "id", []baremetal.Options(nil)).Return(s.Res, nil).Times(2)
	s.Client.On("GetVirtualNetwork", "id", []baremetal.Options(nil)).Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "DeleteVirtualNetwork", "id")
}

func TestResourceCoreVirtualNetworkTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVirtualNetworkTestSuite))
}
