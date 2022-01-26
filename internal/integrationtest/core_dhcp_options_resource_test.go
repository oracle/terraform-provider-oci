// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var defaultDhcpOpts = `
resource "oci_core_default_dhcp_options" "default" {
	manage_default_resource_id = "${oci_core_virtual_network.t.default_dhcp_options_id}"
	compartment_id = "${var.compartment_id}"
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

// issue-routing-tag: core/virtualNetwork
func TestResourceCoreDHCPOptions_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreDHCPOptions_basic")
	defer httpreplay.SaveScenario()

	var resDefaultId, resOpt4Id, resId2 string

	provider := acctest.TestAccProvider

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	config := acctest.LegacyTestProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "network_name"
	}

	resource "oci_core_virtual_network" "t2" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "network_name2"
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

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// Test invalid options type
			{
				Config: acctest.LegacyTestProviderConfig() + defaultDhcpOpts + `
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
				Config: acctest.LegacyTestProviderConfig() + defaultDhcpOpts + `
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
				Config: config,
				Check:  nil,
			},
			{
				Config: acctest.LegacyTestProviderConfig() + defaultDhcpOpts + `
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
				ExpectError: regexp.MustCompile(".*JSON input"),
			},
			{
				Config: config + additionalDhcpOption4 + defaultDhcpOpts + compartmentIdUVariableStr,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(

					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt1", "display_name", "display_name1"),
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt1", "options", map[string]string{
						"type":        "DomainNameServer",
						"server_type": "VcnLocalPlusInternet",
					}, []string{}),

					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "display_name", "display_name2"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt2", "options.#", "2"),
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt2", "options", map[string]string{
						"type":        "DomainNameServer",
						"server_type": "VcnLocalPlusInternet",
					}, []string{}),
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt2", "options", map[string]string{
						"type":                  "SearchDomain",
						"search_domain_names.0": "test.com",
					}, []string{}),

					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt3", "display_name", "display_name3"),
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt3", "options", map[string]string{
						"type":                 "DomainNameServer",
						"server_type":          "CustomDnsServer",
						"custom_dns_servers.0": "8.8.4.4",
						"custom_dns_servers.1": "8.8.8.8",
					}, []string{}),

					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt4", "display_name", "display_name4"),
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt4", "options", map[string]string{
						"type":                 "DomainNameServer",
						"server_type":          "CustomDnsServer",
						"custom_dns_servers.0": "8.8.4.4",
						"custom_dns_servers.1": "8.8.8.8",
					}, []string{}),
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt4", "options", map[string]string{
						"type":                  "SearchDomain",
						"search_domain_names.0": "test.com",
					}, []string{}),

					acctest.CheckResourceSetContainsElementWithProperties("oci_core_default_dhcp_options.default", "options", map[string]string{
						"type":        "DomainNameServer",
						"server_type": "CustomDnsServer",
					}, []string{}),
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_default_dhcp_options.default", "options", map[string]string{
						"type": "SearchDomain",
					}, []string{}),

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
					resource.TestCheckResourceAttrSet("oci_core_default_dhcp_options.default", "compartment_id"),
					func(s *terraform.State) (err error) {
						if resDefaultId, err = acctest.FromInstanceState(s, "oci_core_default_dhcp_options.default", "id"); err != nil {
							return err
						}
						resOpt4Id, err = acctest.FromInstanceState(s, "oci_core_dhcp_options.opt4", "id")
						return err
					},
				),
			},
			// Verify removing default DHCP options
			{
				Config: config + additionalDhcpOption4,
				Check:  nil,
			},
			// Verify adding default DHCP options again
			{
				Config: config + additionalDhcpOption4 + defaultDhcpOpts,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_default_dhcp_options.default", "options", map[string]string{
						"type":                 "DomainNameServer",
						"server_type":          "CustomDnsServer",
						"custom_dns_servers.0": "8.8.4.4",
						"custom_dns_servers.1": "8.8.8.8",
					}, []string{}),
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_default_dhcp_options.default", "options", map[string]string{
						"type":                  "SearchDomain",
						"search_domain_names.0": "test.com",
					}, []string{}),
					resource.TestCheckResourceAttrSet("oci_core_default_dhcp_options.default", "manage_default_resource_id"),
					resource.TestCheckResourceAttrSet("oci_core_default_dhcp_options.default", "id"),
					resource.TestCheckResourceAttrSet("oci_core_default_dhcp_options.default", "time_created"),
					resource.TestCheckResourceAttr("oci_core_default_dhcp_options.default", "state", string(core.DhcpOptionsLifecycleStateAvailable)),
					resource.TestCheckNoResourceAttr("oci_core_default_dhcp_options.default", "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, "oci_core_default_dhcp_options.default", "id")
						if resDefaultId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// ensure that changing the case for options.?.type (polymorphic discriminator) is a no-op.
			{
				Config: config + `
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
				Config: config + `
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, "oci_core_dhcp_options.opt4", "id")
						if resOpt4Id == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
		},
	})
}

//If you set DhcpDnsOption to `VcnLocalPlusInternet`, and you assign a DNS label to the VCN during creation, the search domain name in the VCN's default set of DHCP options is automatically set to the VCN domain
//To avoid multiple applies we perform an apply after the Create in order have the options match what the user has in the config
//This test makes sure we handle that case correctly and that there is a non empty plan after the apply
// issue-routing-tag: core/virtualNetwork
func TestResourceCoreDHCPOptions_avoidServiceDefault(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreDHCPOptions_avoidServiceDefault")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: acctest.LegacyTestProviderConfig() + defaultDhcpOpts + `
					resource "oci_core_virtual_network" "t" {
						cidr_block     = "10.1.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name   = "testVcn"
						dns_label      = "tftestvcn"
					}

					resource "oci_core_dhcp_options" "opt" {
						compartment_id = "${var.compartment_id}"
 						vcn_id         = "${oci_core_virtual_network.t.id}"
 						display_name   = "testDhcpOptions"

 						// required
 						options {
					    	type        = "DomainNameServer"
						    server_type = "VcnLocalPlusInternet"
					    }
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt", "display_name", "testDhcpOptions"),
					resource.TestCheckResourceAttr("oci_core_dhcp_options.opt", "options.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt", "options", map[string]string{
						"type":        "DomainNameServer",
						"server_type": "VcnLocalPlusInternet",
					}, []string{}),
				),
			},
		},
	})
}

// issue-routing-tag: core/virtualNetwork
func TestResourceCoreDHCPOptions_changeOptionsServerType(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreDHCPOptions_changeOptionsServerType")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: acctest.LegacyTestProviderConfig() + defaultDhcpOpts + `
					resource "oci_core_virtual_network" "t" {
						cidr_block     = "10.1.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name   = "testVcn"
						dns_label      = "tftestvcn"
					}

					resource "oci_core_dhcp_options" "opt" {
						compartment_id = "${var.compartment_id}"
  						vcn_id         = "${oci_core_virtual_network.t.id}"
  						display_name   = "testDhcpOptions"

  						// required
  						options {
 					    	type               = "DomainNameServer"
						    server_type        = "CustomDnsServer"
    						custom_dns_servers = ["8.8.4.4", "8.8.8.8"]
					    }
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt", "options", map[string]string{
						"server_type":          "CustomDnsServer",
						"custom_dns_servers.#": "2",
					}, []string{}),
				),
			},
			{
				Config: acctest.LegacyTestProviderConfig() + defaultDhcpOpts + `
					resource "oci_core_virtual_network" "t" {
						cidr_block     = "10.1.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name   = "testVcn"
						dns_label      = "tftestvcn"
					}

					resource "oci_core_dhcp_options" "opt" {
						compartment_id = "${var.compartment_id}"
  						vcn_id         = "${oci_core_virtual_network.t.id}"
  						display_name   = "testDhcpOptions"

  						// required
  						options {
 					    	type        = "DomainNameServer"
						    server_type = "VcnLocalPlusInternet"
					    }
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt", "options", map[string]string{
						"server_type": "VcnLocalPlusInternet",
					}, []string{}),
				),
			},
		},
	})
}

// issue-routing-tag: core/virtualNetwork
func TestResourceCoreDHCPOptions_changeOptionsOrder(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreDHCPOptions_changeOptionsOrder")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: acctest.LegacyTestProviderConfig() + `
					resource "oci_core_virtual_network" "t" {
						cidr_block     = "10.1.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name   = "testVcn"
						dns_label      = "tftestvcn"
					}

					resource "oci_core_dhcp_options" "opt" {
						compartment_id = "${var.compartment_id}"
						vcn_id         = "${oci_core_virtual_network.t.id}"
  						display_name   = "testDhcpOptions"

  						// required
  						options {
    						type = "DomainNameServer"
    						server_type = "VcnLocalPlusInternet"
  						}

  						options {
    						type = "SearchDomain"
    						search_domain_names = ["test.com"]
						}
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt", "options", map[string]string{
						"type":        "DomainNameServer",
						"server_type": "VcnLocalPlusInternet",
					}, []string{}),
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt", "options", map[string]string{
						"type":                  "SearchDomain",
						"search_domain_names.0": "test.com",
					}, []string{}),
				),
			},
			{
				Config: acctest.LegacyTestProviderConfig() + `
					resource "oci_core_virtual_network" "t" {
						cidr_block     = "10.1.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name   = "testVcn"
						dns_label      = "tftestvcn"
					}

					resource "oci_core_dhcp_options" "opt" {
						compartment_id = "${var.compartment_id}"
						vcn_id         = "${oci_core_virtual_network.t.id}"
  						display_name   = "testDhcpOptions"

  						// required
						options {
    						type = "SearchDomain"
    						search_domain_names = ["test.com"]
						}

  						options {
    						type = "DomainNameServer"
    						server_type = "VcnLocalPlusInternet"
  						}
					}`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				Config: acctest.LegacyTestProviderConfig() + `
					resource "oci_core_virtual_network" "t" {
						cidr_block     = "10.1.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name   = "testVcn"
						dns_label      = "tftestvcn"
					}

					resource "oci_core_dhcp_options" "opt" {
						compartment_id = "${var.compartment_id}"
						vcn_id         = "${oci_core_virtual_network.t.id}"
  						display_name   = "testDhcpOptions"

  						// required
  						options {
    						type = "DomainNameServer"
							server_type = "CustomDnsServer"
    						custom_dns_servers = ["8.8.4.4", "8.8.8.8"]
						}
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt", "options", map[string]string{
						"type":                 "DomainNameServer",
						"server_type":          "CustomDnsServer",
						"custom_dns_servers.0": "8.8.4.4",
						"custom_dns_servers.1": "8.8.8.8",
					}, []string{}),
				),
			},
			{
				Config: acctest.LegacyTestProviderConfig() + `
					resource "oci_core_virtual_network" "t" {
						cidr_block     = "10.1.0.0/16"
						compartment_id = "${var.compartment_id}"
						display_name   = "testVcn"
						dns_label      = "tftestvcn"
					}

					resource "oci_core_dhcp_options" "opt" {
						compartment_id = "${var.compartment_id}"
						vcn_id         = "${oci_core_virtual_network.t.id}"
  						display_name   = "testDhcpOptions"

  						// required
  						options {
    						type = "DomainNameServer"
    						server_type = "VcnLocalPlusInternet"
  						}
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					acctest.CheckResourceSetContainsElementWithProperties("oci_core_dhcp_options.opt", "options", map[string]string{
						"type":        "DomainNameServer",
						"server_type": "VcnLocalPlusInternet",
					}, []string{}),
				),
			},
		},
	})
}
