// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUserTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityUserTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig()

	s.ResourceName = "oci_identity_user.t"
}

func (s *ResourceIdentityUserTestSuite) TestAccResourceIdentityUser_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_identity_user" "t" {
					name = "-tf-user"
					description = "automated test user"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-user"),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test user"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceActive),
				),
			},
			// verify update
			{
				Config: s.Config + `
				resource "oci_identity_user" "t" {
					name = "-tf-user"
					description = "automated test user (updated)"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test user (updated)"),
				),
			},
		},
	})
}

func TestResourceIdentityUserTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserTestSuite))
}
