// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

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
}

func (s *DatasourceIdentityGroupsTestSuite) SetupTest() {
	_, tokenFn := tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + tokenFn(`
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
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.id"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityGroupsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityGroupsTestSuite))
}
