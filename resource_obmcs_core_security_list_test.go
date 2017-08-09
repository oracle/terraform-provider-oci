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

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

type ResourceCoreSecurityListTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	SLConfig     string
	ResourceName string
}

func extraWait(ew crud.ExtraWaitPostCreateDelete) {
	return
}

func (s *ResourceCoreSecurityListTestSuite) SetupTest() {
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
`
	s.SLConfig = `
resource "baremetal_core_security_list" "t" {
	compartment_id = "${var.compartment_id}"
	display_name = "security_list0"
	vcn_id = "${baremetal_core_virtual_network.t.id}"
	egress_security_rules = [{
		destination = "0.0.0.0/0"
		protocol = "6"
	}]
	ingress_security_rules = [{
		protocol = "1"
		source = "0.0.0.0/0"
		icmp_options {
			"type" = 3
			"code" = 4
		}
	},
	{
		protocol = "6"
		source = "0.0.0.0/0"
		tcp_options {
			"min" = 80
			"max" = 80
		}
	},
	{
		protocol = "17"
		source = "10.0.0.0/16"
		udp_options {
			"min" = 319
			"max" = 320
		}
	}]
}
	`
	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_security_list.t"
}

func (s *ResourceCoreSecurityListTestSuite) TestCreateResourceCoreSecurityList() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config + s.SLConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "security_list0"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.icmp_options.0.type", "3"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.1.tcp_options.0.max", "80"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.2.udp_options.0.max", "320"),
				),
			},
		},
	})
}

func (s *ResourceCoreSecurityListTestSuite) TestCreateResourceCoreSecurityListUpdateRules() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config + s.SLConfig,
			},
			{
				Config: s.Config + `
					resource "baremetal_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "security_list1"
						vcn_id = "${baremetal_core_virtual_network.t.id}"
						egress_security_rules = [{
							destination = "0.0.0.0/0"
							protocol = "17"
							stateless = true
						}]
						ingress_security_rules = [{
							protocol = "1"
							source = "0.0.0.0/0"
							stateless = true
							icmp_options {
								"type" = 5
								"code" = 0
							}
						},
						{
							protocol = "6"
							source = "0.0.0.0/0"
							stateless = true
							tcp_options {
								"min" = 80
								"max" = 82
							}
						},
						{
							protocol = "17"
							source = "10.0.0.0/16"
							stateless = true
						}]
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "security_list1"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.protocol", "17"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.icmp_options.0.type", "5"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.1.tcp_options.0.max", "82"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.1.stateless", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.2.stateless", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "ingress_security_rules.2.udp_options"),
				),
			},
			// todo: consistent 500 error from server without this step
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
					resource "baremetal_core_security_list" "t" {
						compartment_id = "${var.compartment_id}"
						display_name = "Public"
						vcn_id = "${baremetal_core_virtual_network.t.id}"
						egress_security_rules = []
						ingress_security_rules = []
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "Public"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "0"),
				),
			},
		},
	})
}

func (s *ResourceCoreSecurityListTestSuite) TestDeleteSecurityList() {

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

func TestResourceCoreSecurityListTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreSecurityListTestSuite))
}
