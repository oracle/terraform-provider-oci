// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"
)

type ResourceIdentitySwiftPasswordTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
	Res          *baremetal.SwiftPassword
}

func (s *ResourceIdentitySwiftPasswordTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

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
		resource "baremetal_identity_swift_password" "t" {
			user_id = "user_id"
			description = "` + description + `"
		}
	`
	s.Config += testProviderConfig

	s.TimeCreated = time.Now()
	s.ResourceName = "baremetal_identity_swift_password.t"

	s.Res = &baremetal.SwiftPassword{
		ID:          "id",
		Password:    "password",
		TimeCreated: s.TimeCreated,
		State:       baremetal.ResourceActive,
		UserID:      "user_id",
		Description: description,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	s.Client.On("CreateSwiftPassword", "user_id", description, (*baremetal.RetryTokenOptions)(nil)).
		Return(s.Res, nil).Once()
	s.Client.On("ListSwiftPasswords", "user_id").
		Return(
			&baremetal.ListSwiftPasswords{
				SwiftPasswords: []baremetal.SwiftPassword{
					*s.Res,
				},
			}, nil).Twice()
	s.Client.On("DeleteSwiftPassword", "id", "user_id", (*baremetal.IfMatchOptions)(nil)).Return(nil).Once()
}

func (s *ResourceIdentitySwiftPasswordTestSuite) TestCreateSwiftPassword() {
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

func (s ResourceIdentitySwiftPasswordTestSuite) TestUpdateDescriptionUpdatesSwiftPassword() {
	config := `
		resource "baremetal_identity_swift_password" "t" {
			user_id = "user_id"
			description = "nah nah nah"
		}
  `
	config += testProviderConfig

	res := &baremetal.SwiftPassword{}
	*res = *s.Res
	res.Description = "nah nah nah"
	s.Client.On("UpdateSwiftPassword", "id", "user_id", &baremetal.UpdateIdentityOptions{Description: res.Description}).
		Return(res, nil)
	s.Client.On("ListSwiftPasswords", "user_id").
		Return(
			&baremetal.ListSwiftPasswords{
				SwiftPasswords: []baremetal.SwiftPassword{
					*res,
				},
			}, nil).Twice()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "password", "password"),
				),
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", res.Description),
				),
			},
		},
	})
}

func TestResourceIdentitySwiftPasswordTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentitySwiftPasswordTestSuite))
}
