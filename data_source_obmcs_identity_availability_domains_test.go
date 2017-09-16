// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityAvailabilityDomainsTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListAvailabilityDomains
}

func (s *DatasourceIdentityAvailabilityDomainsTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	data "oci_identity_availability_domains" "t" {
		compartment_id = "${var.compartment_id}"
	}`

	s.ResourceName = "data.oci_identity_availability_domains.t"
}

func (s *DatasourceIdentityAvailabilityDomainsTestSuite) TestAccIdentityAvailabilityDomains_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domains.0.name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domains.1.name"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityAvailabilityDomainsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityAvailabilityDomainsTestSuite))
}
