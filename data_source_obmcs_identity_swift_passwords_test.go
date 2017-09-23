// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentitySwiftPasswordsTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *DatasourceIdentitySwiftPasswordsTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_identity_user" "t" {
		name = "-tf-user"
		description = "tf test user"
	}
	resource "oci_identity_swift_password" "t" {
		user_id = "${oci_identity_user.t.id}"
		description = "tf test user swift password"
	}`
	s.ResourceName = "data.oci_identity_swift_passwords.p"
}

func (s *DatasourceIdentitySwiftPasswordsTestSuite) TestAccDatasourceIdentitySwiftPasswords_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "oci_identity_swift_passwords" "p" {
					user_id = "${oci_identity_user.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "passwords.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "passwords.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "passwords.0.description", "tf test user swift password"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentitySwiftPasswordsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentitySwiftPasswordsTestSuite))
}
