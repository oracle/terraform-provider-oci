// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityGroupsTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	Token        string
	TokenFn      TokenFn
}

func (s *DatasourceIdentityGroupsTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + s.TokenFn(`
	resource "oci_identity_group" "t" {
		name = "{{.token}}"
		description = "automated test group"
	}`, nil)
	s.ResourceName = "data.oci_identity_groups.t"
}

func (s *DatasourceIdentityGroupsTestSuite) TestAccDatasourceIdentityGroups_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "oci_identity_groups" "t" {
					compartment_id = "${var.compartment_id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.#"),
				),
			},
			// Test cascading filters
			{
				Config: s.Config + s.TokenFn(`
				data "oci_identity_groups" "t" {
					compartment_id = "${var.compartment_id}"
					filter {
						name   = "name"
						values = ["{{.token}}", "Administrators"]
					}
					filter {
						name   = "description"
						values = ["automated test group"]
					}
				}`, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "groups.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.description", "automated test group"),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.state", "ACTIVE"),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.inactive_state", "0"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityGroupsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityGroupsTestSuite))
}
