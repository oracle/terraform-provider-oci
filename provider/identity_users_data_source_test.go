// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityUsersTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	Token        string
	TokenFn      TokenFn
}

func (s *DatasourceIdentityUsersTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + s.TokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.token}}"
		description = "automated test user"
	}`, nil)
	s.ResourceName = "data.oci_identity_users.t"
}

func (s *DatasourceIdentityUsersTestSuite) TestAccDatasourceIdentityUsers_basic() {
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
				data "oci_identity_users" "t" {
					compartment_id = "${var.compartment_id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "users.#"),
				),
			},
			{
				Config: s.Config + s.TokenFn(`
				data "oci_identity_users" "t" {
					compartment_id = "${var.compartment_id}"
					filter {
						name = "name"
						values = ["{{.token}}"]
					}
				}`, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "users.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "users.0.name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "users.0.description", "automated test user"),
					resource.TestCheckResourceAttr(s.ResourceName, "users.0.state", "ACTIVE"),
					resource.TestCheckResourceAttr(s.ResourceName, "users.0.inactive_state", "0"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "users.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "users.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "users.0.time_created"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityUsersTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityUsersTestSuite))
}
