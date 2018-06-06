// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreCpeTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreCpeTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
		resource "oci_core_cpe" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-cpe"
			ip_address = "123.123.123.123"
		}
	`

	s.ResourceName = "oci_core_cpe.t"
}

func (s *ResourceCoreCpeTestSuite) TestAccResourceCoreCpe_basic() {
	compartmentId := getCompartmentIDForLegacyTests()
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-cpe"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "ip_address", "123.123.123.123"),
				),
			},
		},
	})
}

func TestResourceCoreCpeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreCpeTestSuite))
}
