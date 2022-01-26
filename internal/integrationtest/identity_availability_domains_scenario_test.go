// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"regexp"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/identity"
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
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.ProviderTestConfig()
	s.ResourceName = "data.oci_identity_availability_domains.t"
}

func (s *DatasourceIdentityAvailabilityDomainsTestSuite) TestAccIdentityAvailabilityDomains_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			// Verify expected number of ADs in expected order. Expect this to fail in single AD regions
			{
				Config: s.Config + `
				data "oci_identity_availability_domains" "t" {
					compartment_id = "${var.tenancy_ocid}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "availability_domains.#", "1"),
					resource.TestMatchResourceAttr(s.ResourceName, "availability_domains.0.name", regexp.MustCompile(".*AD-2")),
				),
			},
		},
	},
	)
}

// issue-routing-tag: identity/default
func TestDatasourceIdentityAvailabilityDomainsTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceIdentityAvailabilityDomainsTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceIdentityAvailabilityDomainsTestSuite))
}
