// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUIPasswordTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.UIPassword
}

func (s *ResourceIdentityUIPasswordTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.Config = `
		resource "baremetal_identity_ui_password" "t" {
			user_id = "user_id"
			version = "1"
		}
	`
	s.Config += testProviderConfig

	s.TimeCreated = baremetal.Time{Time: time.Now()}
	s.ResourceName = "baremetal_identity_ui_password.t"

	s.Res = &baremetal.UIPassword{
		Password:    "password",
		TimeCreated: s.TimeCreated,
		UserID:      "user_id",
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.Client.On("CreateOrResetUIPassword", "user_id", (*baremetal.RetryTokenOptions)(nil)).
		Return(s.Res, nil).Once()
}

func (s *ResourceIdentityUIPasswordTestSuite) TestCreateUIPassword() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "user_id", "user_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "password", "password"),
				),
			},
		},
	})
}

func (s ResourceIdentityUIPasswordTestSuite) TestUpdateVersionForcesNewUIPassword() {
	config := `
		resource "baremetal_identity_ui_password" "t" {
			user_id = "user_id"
			version = "2"
		}
  `
	config += testProviderConfig

	res := &baremetal.UIPassword{
		Password:    "new_password",
		TimeCreated: s.TimeCreated,
		UserID:      "user_id",
	}

	s.Client.On("CreateOrResetUIPassword", "user_id", (*baremetal.RetryTokenOptions)(nil)).
		Return(res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "password", "new_password"),
				),
			},
		},
	})
}

func TestResourceIdentityUIPasswordTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUIPasswordTestSuite))
}
