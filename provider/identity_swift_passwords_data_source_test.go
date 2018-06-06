// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentitySwiftPasswordsTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *DatasourceIdentitySwiftPasswordsTestSuite) SetupTest() {
	_, tokenFn := tokenize()

	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.token}}"
		description = "tf test user"
	}
	resource "oci_identity_swift_password" "t" {
		user_id = "${oci_identity_user.t.id}"
		description = "tf test user swift password"
	}`, nil)
	s.ResourceName = "data.oci_identity_swift_passwords.p"
}

func (s *DatasourceIdentitySwiftPasswordsTestSuite) TestAccDatasourceIdentitySwiftPasswords_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				data "oci_identity_swift_passwords" "p" {
					user_id = "${oci_identity_user.t.id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "passwords.#"),
				),
			},
			{
				Config: s.Config + `
				data "oci_identity_swift_passwords" "p" {
					user_id = "${oci_identity_user.t.id}"
					filter {
						name   = "description"
						values = ["${oci_identity_swift_password.t.description}"]
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "passwords.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "passwords.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "passwords.0.user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "passwords.0.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "passwords.0.description", "tf test user swift password"),
					resource.TestCheckResourceAttr(s.ResourceName, "passwords.0.state", string(identity.SwiftPasswordLifecycleStateActive)),
					// TODO: These fields are not being returned by the service call but are still showing up in the datasource
					// resource.TestCheckNoResourceAttr(s.ResourceName, "passwords.0.expires_on",
					// resource.TestCheckNoResourceAttr(s.ResourceName, "passwords.0.inactive_state"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentitySwiftPasswordsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentitySwiftPasswordsTestSuite))
}
