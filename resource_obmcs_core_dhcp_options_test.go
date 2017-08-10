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

	resource "baremetal_core_dhcp_options" "opt1" {
	  compartment_id = "${var.compartment_id}"
	  vcn_id = "${baremetal_core_virtual_network.t.id}"
	  display_name = "display_name"
	  options {
	    type = "DomainNameServer"
	    server_type = "VcnLocalPlusInternet"
	  }
	}

	resource "baremetal_core_dhcp_options" "opt2" {
	  compartment_id = "${var.compartment_id}"
	  vcn_id = "${baremetal_core_virtual_network.t.id}"
	  display_name = "display_name"
	  options {
	    type = "DomainNameServer"
	    server_type = "VcnLocalPlusInternet"
	  }
	  options {
	    type = "SearchDomain"
	    search_domain_names = [ "test.com" ]
	  }
	}

	resource "baremetal_core_dhcp_options" "opt3" {
	  compartment_id = "${var.compartment_id}"
	  vcn_id = "${baremetal_core_virtual_network.t.id}"
	  display_name = "display_name"
	  options {
	    type = "DomainNameServer"
	    server_type = "CustomDnsServer"
	    custom_dns_servers = [  "8.8.4.4", "8.8.8.8" ]
	  }
	}

	resource "baremetal_core_dhcp_options" "opt4" {
	  compartment_id = "${var.compartment_id}"
	  vcn_id = "${baremetal_core_virtual_network.t.id}"
	  display_name = "display_name"
	  options {
	    type = "DomainNameServer"
	    server_type = "CustomDnsServer"
	    custom_dns_servers = [  "8.8.4.4", "8.8.8.8" ]
	  }
	  options {
	    type = "SearchDomain"
	    search_domain_names = [ "test.com" ]
	  }
	}
	`
	s.Config += testProviderConfig()
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

					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt1", "display_name", "display_name"),

					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt1", "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt1", "options.0.server_type", "VcnLocalPlusInternet"),

					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt2", "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt2", "options.0.server_type", "VcnLocalPlusInternet"),
					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt2", "options.1.type", "SearchDomain"),

					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt3", "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt3", "options.0.server_type", "CustomDnsServer"),

					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt4", "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt4", "options.0.server_type", "CustomDnsServer"),
					resource.TestCheckResourceAttr("baremetal_core_dhcp_options.opt4", "options.1.type", "SearchDomain"),
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
