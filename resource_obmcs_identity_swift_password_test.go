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

	"github.com/oracle/terraform-provider-baremetal/client"
)

type ResourceIdentitySwiftPasswordTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
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
		"baremetal": s.Provider,
	}

	description := "blah blah blah"
	s.Config = `
		resource "baremetal_identity_user" "t" {
			name = "name1"
			description = "desc!"
		}
		resource "baremetal_identity_swift_password" "t" {
			user_id = "${baremetal_identity_user.t.id}"
			description = "` + description + `"
		}
	`
	s.Config += testProviderConfig()

	s.TimeCreated = time.Now()
	s.ResourceName = "baremetal_identity_swift_password.t"

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
		resource "baremetal_identity_user" "t" {
			name = "name1"
			description = "desc!"
		}
		resource "baremetal_identity_swift_password" "t" {
			user_id = "${baremetal_identity_user.t.id}"
			description = "nah nah nah"
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
					resource.TestCheckResourceAttr(s.ResourceName, "description", "nah nah nah"),
				),
			},
		},
	})
}

func TestResourceIdentitySwiftPasswordTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentitySwiftPasswordTestSuite))
}
