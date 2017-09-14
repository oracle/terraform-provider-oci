// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentitySwiftPasswordTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
	Res          *baremetal.SwiftPassword
}

func (s *ResourceIdentitySwiftPasswordTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}

	description := "blah blah blah"
	s.Config = `
		resource "oci_identity_user" "t" {
			name = "name1"
			description = "desc!"
		}
		resource "oci_identity_swift_password" "t" {
			user_id = "${oci_identity_user.t.id}"
			description = "` + description + `"
		}
	`
	s.Config += testProviderConfig()

	s.TimeCreated = time.Now()
	s.ResourceName = "oci_identity_swift_password.t"

}

func (s *ResourceIdentitySwiftPasswordTestSuite) TestCreateSwiftPassword() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
				),
			},
		},
	})
}

func (s ResourceIdentitySwiftPasswordTestSuite) TestUpdateDescriptionUpdatesSwiftPassword() {
	config := `
		resource "oci_identity_user" "t" {
			name = "-tf-user"
			description = "automated test user"
		}
		resource "oci_identity_swift_password" "t" {
			user_id = "${oci_identity_user.t.id}"
			description = "automated test swift password"
		}
  `
	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
				),
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test swift password"),
				),
			},
		},
	})
}

func TestResourceIdentitySwiftPasswordTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentitySwiftPasswordTestSuite))
}
