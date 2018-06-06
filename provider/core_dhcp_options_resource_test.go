// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	core "github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreDHCPOptionsTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

var defaultDhcpOpts = `
resource "oci_core_default_dhcp_options" "default" {
	manage_default_resource_id = "${oci_core_virtual_network.t.default_dhcp_options_id}"
	options {
		type = "DomainNameServer"
		server_type = "CustomDnsServer"
		custom_dns_servers = [ "8.8.4.4", "8.8.8.8" ]
	}
	options {
		type = "SearchDomain"
		search_domain_names = [ "test.com" ]
	}
}
`

var additionalDhcpOption4 = `
	resource "oci_core_dhcp_options" "opt4" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		display_name = "display_name4"
		options {
			type = "DomainNameServer"
			server_type = "CustomDnsServer"
			custom_dns_servers = [ "8.8.4.4", "8.8.8.8" ]
		}
		options {
			type = "SearchDomain"
			search_domain_names = [ "test.com" ]
		}
	}`

func (s *ResourceCoreDHCPOptionsTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "network_name"
	}

	resource "oci_core_virtual_network" "t2" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "network_name"
	}

	resource "oci_core_dhcp_options" "opt1" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		display_name = "display_name1"
		options {
			type = "DomainNameServer"
			server_type = "VcnLocalPlusInternet"
		}
	}

	resource "oci_core_dhcp_options" "opt2" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		display_name = "display_name2"
		options {
			type = "domainNAMEserver"	# case-insensitive
			server_type = "VcnLocalPlusInternet"
		}
		options {
			type = "seaRCHdomAIN"		# case-insensitive
			search_domain_names = [ "test.com" ]
		}
	}

	resource "oci_core_dhcp_options" "opt3" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		display_name = "display_name3"
		options {
			type = "DomainNameServer"
			server_type = "CustomDnsServer"
			custom_dns_servers = [ "8.8.4.4", "8.8.8.8" ]
		}
	}`
}

func (s *ResourceCoreDHCPOptionsTestSuite) TestAccResourceCoreDHCPOptions_basic() {
	var resDefaultId, resOpt4Id, resId2 string

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config + additionalDhcpOption4 + defaultDhcpOpts,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt1", "display_name", "display_name1"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt1", "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt1", "options.0.server_type", "VcnLocalPlusInternet"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt1", "options.0.custom_dns_servers.#", "0"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt1", "options.0.search_domain_names.#", "0"),

					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "display_name", "display_name2"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "options.#", "2"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "options.0.server_type", "VcnLocalPlusInternet"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "options.1.type", "SearchDomain"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "options.0.custom_dns_servers.#", "0"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt2", "options.1.search_domain_names.#"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "options.1.search_domain_names.0", "test.com"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "options.0.custom_dns_servers.#", "0"),

					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt3", "display_name", "display_name3"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt3", "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt3", "options.0.server_type", "CustomDnsServer"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt3", "options.0.search_domain_names.#", "0"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt3", "options.0.custom_dns_servers.#"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt3", "options.0.custom_dns_servers.0", "8.8.4.4"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt3", "options.0.custom_dns_servers.1", "8.8.8.8"),

					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt4", "display_name", "display_name4"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt4", "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt4", "options.0.server_type", "CustomDnsServer"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt4", "options.0.custom_dns_servers.#"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt4", "options.0.custom_dns_servers.0", "8.8.4.4"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt4", "options.0.custom_dns_servers.1", "8.8.8.8"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt4", "options.1.type", "SearchDomain"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt4", "options.1.search_domain_names.0", "test.com"),

					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "options.0.server_type", "CustomDnsServer"),
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "options.1.type", "SearchDomain"),

					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt1", "vcn_id"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt1", "id"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt1", "time_created"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt1", "state", string(core.DhcpOptionsLifecycleStateAvailable)),

					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt2", "vcn_id"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt2", "id"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt2", "time_created"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "state", string(core.DhcpOptionsLifecycleStateAvailable)),

					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt3", "vcn_id"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt3", "id"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt3", "time_created"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt3", "state", string(core.DhcpOptionsLifecycleStateAvailable)),

					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt4", "vcn_id"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt4", "id"),
					resource.TestCheckResourceAttrSet("oci_core_dhcp_options.opt4", "time_created"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt4", "state", string(core.DhcpOptionsLifecycleStateAvailable)),

					resource.TestCheckResourceAttrSet("oci_core_default_dhcp_options.default", "manage_default_resource_id"),
					resource.TestCheckResourceAttrSet("oci_core_default_dhcp_options.default", "id"),
					resource.TestCheckResourceAttrSet("oci_core_default_dhcp_options.default", "time_created"),
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "state", string(core.DhcpOptionsLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr("oci_core_default_dhcp_options.default", "vcn_id"),
					resource.TestCheckNoResourceAttr("oci_core_default_dhcp_options.default", "compartment_id"),

					func(s *terraform.State) (err error) {
						if resDefaultId, err = fromInstanceState(s, "oci_core_default_dhcp_options.default", "id"); err != nil {
							return err
						}
						resOpt4Id, err = fromInstanceState(s, "oci_core_dhcp_options.opt4", "id")
						return err
					},
				),
			},
			// Verify removing default DHCP options
			{
				Config: s.Config + additionalDhcpOption4,
				Check:  nil,
			},
			// Verify adding default DHCP options again
			{
				Config: s.Config + additionalDhcpOption4 + defaultDhcpOpts,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "options.0.type", "DomainNameServer"),
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "options.0.server_type", "CustomDnsServer"),
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "options.0.custom_dns_servers.0", "8.8.4.4"),
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "options.0.custom_dns_servers.1", "8.8.8.8"),
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "options.1.type", "SearchDomain"),
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "options.1.search_domain_names.0", "test.com"),
					resource.TestCheckResourceAttrSet("oci_core_default_dhcp_options.default", "manage_default_resource_id"),
					resource.TestCheckResourceAttrSet("oci_core_default_dhcp_options.default", "id"),
					resource.TestCheckResourceAttrSet("oci_core_default_dhcp_options.default", "time_created"),
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "state", string(core.DhcpOptionsLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr("oci_core_default_dhcp_options.default", "vcn_id"),
					resource.TestCheckNoResourceAttr("oci_core_default_dhcp_options.default", "compartment_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_core_default_dhcp_options.default", "id")
						if resDefaultId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// ensure that changing the case for options.?.type (polymorphic discriminator) is a no-op.
			{
				Config: s.Config + `
					resource "oci_core_dhcp_options" "opt4" { # Same as additionalDhcpOption4 but with diff casing for 'type'.
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						display_name = "display_name4"
						options {
							type = "domAINnameSERVER"	# case-insensitive
							server_type = "CustomDnsServer"
							custom_dns_servers = [ "8.8.4.4", "8.8.8.8" ]
						}
						options {
							type = "seaRCHdomAIN"		# case-insensitive
							search_domain_names = [ "test.com" ]
						}
					}` + defaultDhcpOpts,
				PlanOnly: true,
			},
			// Changing to a different vcn should be a ForceNew
			{
				Config: s.Config + `
					resource "oci_core_dhcp_options" "opt4" { # Same as alternate additionalDhcpOption4 above, but with diff vcn.
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t2.id}"
						display_name = "display_name4"
						options {
							type = "domAINnameSERVER"	# case-insensitive
							server_type = "CustomDnsServer"
							custom_dns_servers = [ "8.8.4.4", "8.8.8.8" ]
						}
						options {
							type = "seaRCHdomAIN"		# case-insensitive
							search_domain_names = [ "test.com" ]
						}
					}` + defaultDhcpOpts,
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_core_dhcp_options.opt4", "id")
						if resOpt4Id == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// Test invalid options type
			{
				Config: legacyTestProviderConfig() + defaultDhcpOpts + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name = "network_name"
					}

					resource "oci_core_dhcp_options" "opt6" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						display_name = "display_name6"
						options {
							type = "DomainNameServer_ShouldNotWork"
							server_type = "CustomDnsServer"
							custom_dns_servers = [ "8.8.4.4", "8.8.8.8" ]
						}
					}`,
				ExpectError: regexp.MustCompile("expected options.0.type to be one of"),
			},
			// Test invalid options server_type
			{
				Config: legacyTestProviderConfig() + defaultDhcpOpts + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name = "network_name"
					}

					resource "oci_core_dhcp_options" "opt6" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						display_name = "display_name6"
						options {
							type = "DomainNameServer"
							server_type = "CustomDnsServer_ShouldNotWork"
							custom_dns_servers = [ "8.8.4.4", "8.8.8.8" ]
						}
					}`,
				ExpectError: regexp.MustCompile("expected options.0.server_type to be one of"),
			},
			// Verify result of strange polymorphic options
			{
				Config: s.Config,
				Check:  nil,
			},
			{
				Config: legacyTestProviderConfig() + defaultDhcpOpts + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name = "network_name"
					}

					resource "oci_core_dhcp_options" "opt5" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						display_name = "display_name5"
						options {
							type = "DomainNameServer"
						}
						options {
							type = "SearchDomain"
						}
					}`,
				ExpectError: regexp.MustCompile("InvalidParameter.*JSON input"),
			},
			{
				Config: legacyTestProviderConfig() + defaultDhcpOpts + `
					resource "oci_core_virtual_network" "t" {
						cidr_block = "10.0.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name = "network_name"
					}

					resource "oci_core_dhcp_options" "opt6" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						display_name = "display_name6"
						options {
							type = "DomainNameServer"
							server_type = "CustomDnsServer"
							custom_dns_servers = [ "8.8.4.4", "8.8.8.8" ]
							search_domain_names = [ "test.com" ]
						}
						options {
							type = "SearchDomain"
							server_type = "CustomDnsServer"
							custom_dns_servers = [ "8.8.4.4", "8.8.8.8" ]
							search_domain_names = [ "test.com" ]
						}
					}`,
				ExpectError: regexp.MustCompile("should not be specified for type"),
			},
		},
	})
}

func TestResourceCoreDHCPOptionsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDHCPOptionsTestSuite))
}
