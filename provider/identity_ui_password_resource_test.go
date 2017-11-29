// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUIPasswordTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityUIPasswordTestSuite) SetupTest() {
	_, tokenFn := tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "-tf-user"
		description = "tf test user"
	}`, nil)

	s.ResourceName = "oci_identity_ui_password.t"
}

func (s *ResourceIdentityUIPasswordTestSuite) TestAccIdentityUIPassword_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_identity_ui_password" "t" {
					user_id = "${oci_identity_user.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
				),
			},
		},
	})
}

func TestResourceIdentityUIPasswordTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUIPasswordTestSuite))
}
