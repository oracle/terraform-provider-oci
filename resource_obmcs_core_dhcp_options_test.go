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

type ResourceCoreDHCPOptionsTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.DHCPOptions
	DeletedRes   *baremetal.DHCPOptions
}

func (s *ResourceCoreDHCPOptionsTestSuite) SetupTest() {
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
		display_name = "network_name"
	}
	resource "baremetal_core_dhcp_options" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
     		options {
			type = "DomainNameServer"
			custom_dns_servers = [ "8.8.8.8" ]
			server_type = "CustomDnsServer"
		}
     		vcn_id = "${baremetal_core_virtual_network.t.id}"
	}
	`
	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_dhcp_options.t"
}

func (s *ResourceCoreDHCPOptionsTestSuite) TestCreateResourceCoreDHCPOptions() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr(s.ResourceName, "options.0.server_type", "CustomDnsServer"),
				),
			},
		},
	})
}

func (s *ResourceCoreDHCPOptionsTestSuite) TestDeleteDHCPOptions() {

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

}

func TestResourceCoreDHCPOptionsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDHCPOptionsTestSuite))
}
