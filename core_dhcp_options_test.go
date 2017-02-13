// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

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

type ResourceCoreDHCPOptionsTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.DHCPOptions
	DeletedRes   *baremetal.DHCPOptions
}

func (s *ResourceCoreDHCPOptionsTestSuite) SetupTest() {
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
		resource "baremetal_core_dhcp_options" "t" {
			compartment_id = "compartment_id"
			display_name = "display_name"
      options {
				type = "type"
				custom_dns_servers = [ "custom_dns_servers" ]
				server_type = "server_type"
			}
      options {
				type = "type"
				custom_dns_servers = [ "custom_dns_servers" ]
				server_type = "server_type"
			}
			vcn_id = "vcn_id"
		}
	`
	s.Config += testProviderConfig

	s.ResourceName = "baremetal_core_dhcp_options.t"

	entities := []baremetal.DHCPDNSOption{
		{
			Type:             "type",
			CustomDNSServers: []string{"custom_dns_servers"},
			ServerType:       "server_type",
		},
		{
			Type:             "type",
			CustomDNSServers: []string{"custom_dns_servers"},
			ServerType:       "server_type",
		},
	}

	s.Res = &baremetal.DHCPOptions{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		ID:            "id",
		Options:       entities,
		State:         baremetal.ResourceAvailable,
		TimeCreated:   s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	deletedRes := *s.Res
	s.DeletedRes = &deletedRes
	s.DeletedRes.State = baremetal.ResourceTerminated

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = "display_name"
	s.Client.On("CreateDHCPOptions", "compartment_id", "vcn_id", entities, opts).Return(s.Res, nil)
	s.Client.On("DeleteDHCPOptions", "id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceCoreDHCPOptionsTestSuite) TestCreateResourceCoreDHCPOptions() {
	s.Client.On("GetDHCPOptions", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetDHCPOptions", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", s.Res.DisplayName),
					resource.TestCheckResourceAttr(s.ResourceName, "options.0.type", "type"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.1.server_type", "server_type"),
				),
			},
		},
	})
}

func (s ResourceCoreDHCPOptionsTestSuite) TestUpdateDHCPOptions() {
	s.Client.On("GetDHCPOptions", "id").Return(s.Res, nil).Times(3)

	config := `
		resource "baremetal_core_dhcp_options" "t" {
			compartment_id = "compartment_id"
			display_name = "display_name"
      options {
				type = "new_type"
				custom_dns_servers = [ "new_custom_dns_servers" ]
				server_type = "new_server_type"
			}
			vcn_id = "vcn_id"
		}
	`
	config += testProviderConfig

	entities := []baremetal.DHCPDNSOption{
		{
			Type:             "new_type",
			CustomDNSServers: []string{"new_custom_dns_servers"},
			ServerType:       "new_server_type",
		},
	}

	res := &baremetal.DHCPOptions{
		CompartmentID: "compartment_id",
		DisplayName:   "display_name",
		ID:            "id",
		Options:       entities,
		State:         baremetal.ResourceAvailable,
		TimeCreated:   s.TimeCreated,
	}
	res.ETag = "etag"
	res.RequestID = "opcrequestid"

	opts := &baremetal.UpdateDHCPDNSOptions{}
	opts.Options = []baremetal.DHCPDNSOption{
		{
			Type:             "new_type",
			CustomDNSServers: []string{"new_custom_dns_servers"},
			ServerType:       "new_server_type",
		},
	}

	s.Client.On("UpdateDHCPOptions", "id", opts).Return(res, nil)
	s.Client.On("GetDHCPOptions", "id").Return(res, nil).Times(2)
	s.Client.On("GetDHCPOptions", "id").Return(s.DeletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config: config,
				Check:  resource.TestCheckResourceAttr(s.ResourceName, "options.0.type", "new_type"),
			},
		},
	})
}

func (s *ResourceCoreDHCPOptionsTestSuite) TestDeleteDHCPOptions() {
	s.Client.On("GetDHCPOptions", "id").Return(s.Res, nil).Times(2)
	s.Client.On("GetDHCPOptions", "id").Return(s.DeletedRes, nil)

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

	s.Client.AssertCalled(s.T(), "DeleteDHCPOptions", "id", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceCoreDHCPOptionsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDHCPOptionsTestSuite))
}
