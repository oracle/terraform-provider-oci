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

type ResourceIdentityUIPasswordTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.UIPassword
}

func (s *ResourceIdentityUIPasswordTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}

	s.Config = `
		resource "oci_identity_user" "t" {
			name = "-tf-user"
			description = "automated test user"
		}
		resource "oci_identity_ui_password" "t" {
			user_id = "${oci_identity_user.t.id}"
		}
	`
	s.Config += testProviderConfig()

	s.TimeCreated = baremetal.Time{Time: time.Now()}
	s.ResourceName = "oci_identity_ui_password.t"
}

func (s *ResourceIdentityUIPasswordTestSuite) TestCreateUIPassword() {
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

func (s ResourceIdentityUIPasswordTestSuite) TestUpdateVersionForcesNewUIPassword() {
	config := `
		resource "oci_identity_user" "t" {
			name = "-tf-user"
			description = "automated test user"
		}
		resource "oci_identity_ui_password" "t" {
			user_id = "${oci_identity_user.t.id}"
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
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
				),
			},
		},
	})
}

func TestResourceIdentityUIPasswordTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUIPasswordTestSuite))
}
