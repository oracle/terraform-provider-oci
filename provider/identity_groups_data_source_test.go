// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityGroupsTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	Token        string
	TokenFn      TokenFn
}

func (s *DatasourceIdentityGroupsTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + s.TokenFn(`
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
					compartment_id = "${var.tenancy_ocid}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.#"),
				),
			},
			// Test cascading filters
			{
				Config: s.Config + s.TokenFn(`
				data "oci_identity_groups" "t" {
					compartment_id = "${var.tenancy_ocid}"
					filter {
						name   = "name"
						values = ["{{.token}}", "Administrators"]
					}
					filter {
						name   = "description"
						values = ["automated test group"]
					}
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "groups.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.description", "automated test group"),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.state", string(identity.GroupLifecycleStateActive)),
					// TODO: This field is not being returned by the service call but is still showing up in the datasource
					// resource.TestCheckNoResourceAttr(s.ResourceName, "groups.0.inactive_state"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityGroupsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityGroupsTestSuite))
}
