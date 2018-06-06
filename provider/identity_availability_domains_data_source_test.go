// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityAvailabilityDomainsTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         identity.ListAvailabilityDomainsResponse
}

func (s *DatasourceIdentityAvailabilityDomainsTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = testProviderConfig()
	s.ResourceName = "data.oci_identity_availability_domains.t"
}

func (s *DatasourceIdentityAvailabilityDomainsTestSuite) TestAccIdentityAvailabilityDomains_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			// Verify expected number of ADs in expected order
			{
				Config: s.Config + `
				data "oci_identity_availability_domains" "t" {
					compartment_id = "${var.tenancy_ocid}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domains.#", "3"),
					resource.TestMatchResourceAttr(s.ResourceName, "availability_domains.0.name", regexp.MustCompile(`\w*-AD-1`)),
					resource.TestMatchResourceAttr(s.ResourceName, "availability_domains.1.name", regexp.MustCompile(`\w*-AD-2`)),
					resource.TestMatchResourceAttr(s.ResourceName, "availability_domains.2.name", regexp.MustCompile(`\w*-AD-3`)),
				),
			},
			// Verify regex filtering
			{
				Config: s.Config + `
				data "oci_identity_availability_domains" "t" {
					compartment_id = "${var.tenancy_ocid}"
					filter {
						name = "name"
						values = ["\\w*-AD-2"]
						regex = true
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
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
