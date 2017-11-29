// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityGroupTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityGroupTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig()
	s.ResourceName = "oci_identity_group.t"
}

func (s *ResourceIdentityGroupTestSuite) TestAccResourceIdentityGroup_basic() {
	token, tokenFn := tokenize()
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + tokenFn(`
				resource "oci_identity_group" "t" {
					name = "{{.token}}"
					description = "tf test group"
				}`, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", token),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test group"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceActive),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
			// verify update
			{
				Config: s.Config + tokenFn(`
				resource "oci_identity_group" "t" {
					name = "{{.token}}"
					description = "tf test group (updated)"
				}`, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test group (updated)"),
				),
			},
		},
	})
}

func TestResourceIdentityGroupTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityGroupTestSuite))
}
