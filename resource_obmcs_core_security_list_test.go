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

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type ResourceCoreSecurityListTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.SecurityList
	DeletingRes  *baremetal.SecurityList
	DeletedRes   *baremetal.SecurityList
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
}

func (s *ResourceCoreSecurityListTestSuite) TestCreateResourceCoreSecurityList() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "Public"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.icmp_options.0.code", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.stateless", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.tcp_options.0.max", "2"),
				),
			},
		},
	})
}

func (s ResourceCoreSecurityListTestSuite) TestUpdateSecurityList() {

	config := `
		resource "baremetal_core_security_list" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
      egress_security_rules {
				destination = "destination"
				icmp_options {
					"code" = 1
					"type" = 2
				}
				protocol = "protocol"
				stateless = true
			}
      ingress_security_rules {
				tcp_options {
					"max" = 3
					"min" = 1
				}
				protocol = "protocol"
				source = "source"
			}
			vcn_id = "vcn_id"
		}
	`
	config += testProviderConfig()

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
					resource.TestCheckResourceAttr(s.ResourceName, "egress_security_rules.0.icmp_options.0.code", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "ingress_security_rules.0.tcp_options.0.max", "3"),
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
