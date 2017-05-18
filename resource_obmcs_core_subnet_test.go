// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

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

type ResourceCoreSubnetTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Subnet
	DeletedRes   *baremetal.Subnet
}

func (s *ResourceCoreSubnetTestSuite) SetupTest() {
	s.Client = GetTestProvider()

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
	cidr_block = "10.0.0.0/16"
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
}

resource "baremetal_core_internet_gateway" "CompleteIG" {
    compartment_id = "${var.compartment_id}"
    display_name = "CompleteIG"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
}

resource "baremetal_core_route_table" "RouteForComplete" {
    compartment_id = "${var.compartment_id}"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
    display_name = "RouteTableForComplete"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${baremetal_core_internet_gateway.CompleteIG.id}"
    }
}

resource "baremetal_core_security_list" "WebSubnet" {
    compartment_id = "${var.compartment_id}"
    display_name = "Public"
    vcn_id = "${baremetal_core_virtual_network.t.id}"
    egress_security_rules = [{
        destination = "0.0.0.0/0"
        protocol = "6"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 80
            "min" = 80
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
	{
	protocol = "6"
	source = "10.0.0.0/16"
    }]
}


resource "baremetal_core_subnet" "WebSubnetAD1" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block = "10.0.0.0/16"
  display_name = "WebSubnetAD1"
  compartment_id = "${var.compartment_id}"
  vcn_id = "${baremetal_core_virtual_network.t.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.WebSubnet.id}"]
}

	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_subnet.t"
	s.Res = &baremetal.Subnet{
		AvailabilityDomain: "availabilitydomainid",
		CIDRBlock:          "10.10.10.0/24",
		CompartmentID:      "compartmentid",
		DisplayName:        "display_name",
		ID:                 "id",
		RouteTableID:       "routetableid",
		SecurityListIDs: []string{
			// Note: sorted by schema.HashString
			"slid2",
			"slid1",
		},
		State: baremetal.ResourceAvailable,
		TimeCreated: baremetal.Time{
			Time: time.Now(),
		},
		VcnID:            "vcnid",
		VirtualRouterIP:  "virtualrouterip",
		VirtualRouterMac: "virtualroutermac",
	}

	s.DeletedRes = &baremetal.Subnet{}
	*s.DeletedRes = *s.Res
	s.DeletedRes.State = baremetal.ResourceTerminated

	opts := &baremetal.CreateSubnetOptions{}
	opts.DisplayName = "display_name"
	opts.RouteTableID = s.Res.RouteTableID
	opts.SecurityListIDs = s.Res.SecurityListIDs
	s.Client.On(
		"CreateSubnet",
		s.Res.AvailabilityDomain,
		s.Res.CIDRBlock,
		s.Res.CompartmentID,
		s.Res.VcnID,
		opts).Return(s.Res, nil)
	s.Client.On("DeleteSubnet", s.Res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreSubnetTestSuite) TestCreateResourceCoreSubnet() {
	s.Client.On("GetSubnet", s.Res.ID).Return(s.Res, nil).Times(2)
	s.Client.On("GetSubnet", s.Res.ID).Return(s.DeletedRes, nil).Times(2)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", s.Res.AvailabilityDomain),

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "id", s.Res.ID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
				),
			},
		},
	})
}

func (s *ResourceCoreSubnetTestSuite) TestCreateResourceCoreSubnetWithoutDisplayName() {
	if IsAccTest() {
		s.T().Skip()
	}
	s.Client.On("GetSubnet", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetSubnet", "id").Return(s.DeletedRes, nil).Times(2)

	s.Config = `
resource "baremetal_core_subnet" "t" {
  availability_domain = "availabilitydomainid"
  compartment_id = "${var.compartment_id}"
  cidr_block          = "10.10.10.0/24"
  route_table_id      = "routetableid"
  vcn_id              = "vcnid"
  security_list_ids   = ["slid1", "slid2"]
}
	`

	s.Config += testProviderConfig()

	s.Res.DisplayName = ""

	opts := &baremetal.CreateSubnetOptions{}
	opts.RouteTableID = s.Res.RouteTableID
	opts.SecurityListIDs = s.Res.SecurityListIDs
	s.Client.On(
		"CreateSubnet",
		s.Res.AvailabilityDomain,
		s.Res.CIDRBlock,
		s.Res.CompartmentID,
		s.Res.VcnID,
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

func (s ResourceCoreSubnetTestSuite) TestUpdateCompartmentIDForcesNewSubnet() {
	if IsAccTest() {
		s.T().Skip()
	}
	config := `
resource "baremetal_core_subnet" "t" {
  availability_domain = "availabilitydomainid"
  compartment_id      = "new_compartment_id"
  display_name        = "display_name"
  cidr_block          = "10.10.10.0/24"
  route_table_id      = "routetableid"
  vcn_id              = "vcnid"
  security_list_ids   = ["slid1", "slid2"]
}
	`

	config += testProviderConfig()

	res := &baremetal.Subnet{
		AvailabilityDomain: "availabilitydomainid",
		CIDRBlock:          "10.10.10.0/24",
		CompartmentID:      "new_compartment_id",
		DisplayName:        "display_name",
		ID:                 "new_id",
		RouteTableID:       "routetableid",
		SecurityListIDs: []string{
			// Note: sorted by schema.HashString
			"slid2",
			"slid1",
		},
		State: baremetal.ResourceAvailable,
		TimeCreated: baremetal.Time{
			Time: time.Now(),
		},
		VcnID:            "vcnid",
		VirtualRouterIP:  "virtualrouterip",
		VirtualRouterMac: "virtualroutermac",
	}
	delRes := &baremetal.Subnet{}
	*delRes = *res
	delRes.State = baremetal.ResourceTerminated
	opts := &baremetal.CreateSubnetOptions{}
	opts.DisplayName = "display_name"
	opts.RouteTableID = res.RouteTableID
	opts.SecurityListIDs = res.SecurityListIDs

	s.Client.On("GetSubnet", s.Res.ID).Return(s.Res, nil).Times(2)
	s.Client.On("DeleteSubnet", s.Res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)
	s.Client.On("GetSubnet", s.Res.ID).Return(s.DeletedRes, nil).Times(2)

	s.Client.On(
		"CreateSubnet",
		res.AvailabilityDomain,
		res.CIDRBlock,
		res.CompartmentID,
		res.VcnID,
		opts).Return(res, nil).Once()

	s.Client.On("GetSubnet", res.ID).Return(res, nil).Times(2)
	s.Client.On("DeleteSubnet", res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)
	s.Client.On("GetSubnet", res.ID).Return(delRes, nil).Times(2)

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

				),
			},
		},
	})
}

func (s *ResourceCoreSubnetTestSuite) TestTerminateSubnet() {
	if IsAccTest() {
		s.T().Skip()
	}
	s.Client.On("GetSubnet", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetSubnet", "id").Return(s.DeletedRes, nil)

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

	s.Client.On("DeleteSubnet", s.Res.ID, (*baremetal.IfMatchOptions)(nil)).Return(nil)

}

func TestResourceCoreSubnetTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreSubnetTestSuite))
}
