// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ResourceCoreSecurityListTestSuite struct {
	suite.Suite
	Client       mockableClient
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
`
	s.SLConfig = `
resource "baremetal_core_security_list" "t" {
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
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "Public"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.stateless", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.tcp_options.0.max", "80"),
				),
			},
		},
	})
}

func (s *ResourceCoreSecurityListTestSuite) TestCreateResourceCoreSecurityListRemoveRules() {

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
