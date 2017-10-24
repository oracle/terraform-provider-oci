// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"regexp"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
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
	s.Config = testProviderConfig()
	s.ResourceName = "data.oci_identity_availability_domains.t"
}

func (s *DatasourceIdentityAvailabilityDomainsTestSuite) TestAccIdentityAvailabilityDomains_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			// Verify expected number of ADs
			{
				Config: s.Config + `
				data "oci_identity_availability_domains" "t" {
					compartment_id = "${var.compartment_id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domains.#", "3"),
				),
			},
			// Verify regex filtering
			{
				Config: s.Config + `
				data "oci_identity_availability_domains" "t" {
					compartment_id = "${var.compartment_id}"
					filter {
						name = "name"
						values = ["\\w*:\\w{3}-AD-2"]
						regex = true
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domains.#", "1"),
					resource.TestMatchResourceAttr(s.ResourceName, "availability_domains.0.name", regexp.MustCompile(".*AD-2")),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityAvailabilityDomainsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityAvailabilityDomainsTestSuite))
}
