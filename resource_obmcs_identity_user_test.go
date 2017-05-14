// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"




	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUserTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
	Res          *baremetal.User
}

func (s *ResourceIdentityUserTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = `
		resource "baremetal_identity_user" "t" {
			name = "name1"
			description = "desc!"
		}
	`
	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_identity_user.t"
	s.Res = &baremetal.User{
		ID:            "id!",
		Name:          "name1",
		Description:   "desc!",
		CompartmentID: "cid!",
		State:         baremetal.ResourceActive,
		TimeCreated:   s.TimeCreated,
	}
	s.Client.On("CreateUser", "name1", "desc!", (*baremetal.RetryTokenOptions)(nil)).
		Return(s.Res, nil)
	s.Client.On("DeleteUser", "id!", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceIdentityUserTestSuite) TestCreateResourceIdentityUser() {
	s.Client.On("GetUser", "id!").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", s.Res.Name),
					resource.TestCheckResourceAttr(s.ResourceName, "description", s.Res.Description),

					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestCreateResourceIdentityUserPolling() {
	if IsAccTest() {
		s.T().Skip()
	}
	s.Res.State = baremetal.ResourceCreating
	s.Client.On("GetUser", "id!").Return(s.Res, nil).Once()

	u := *s.Res
	u.State = baremetal.ResourceActive
	s.Client.On("GetUser", "id!").Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceActive),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestUpdateResourceIdentityUserDescription() {
	s.Client.On("GetUser", "id!").Return(s.Res, nil).Twice()

	c := `

		resource "baremetal_identity_user" "t" {
			name = "name1"
			description = "newdesc!"
		}
	`
	c += testProviderConfig()

	u := *s.Res
	u.Description = "newdesc!"
	opts := &baremetal.UpdateIdentityOptions{}
	opts.Description = "newdesc!"
	s.Client.On("UpdateUser", "id!", opts).Return(&u, nil)
	s.Client.On("GetUser", "id!").Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: c,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestFailedUpdateResourceIdentityUserDescription() {
	if IsAccTest() {
		s.T().Skip()
	}
	s.Client.On("GetUser", "id!").Return(s.Res, nil).Times(3)

	c := `

		resource "baremetal_identity_user" "t" {
			name = "name1"
			description = "newdesc!"
		}

	`

	c += testProviderConfig()

	opts := &baremetal.UpdateIdentityOptions{}
	opts.Description = "newdesc!"
	s.Client.On("UpdateUser", "id!", opts).
		Return(nil, errors.New("FAILED!")).Once()

	u := *s.Res
	u.Description = "newdesc!"
	s.Client.On("UpdateUser", "id!", opts).Return(&u, nil)
	s.Client.On("GetUser", "id!").Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:      c,
				ExpectError: regexp.MustCompile(`FAILED`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "desc!"),
				),
			},
			{
				Config: c,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestUpdateResourceIdentityUserNameShouldCreateNew() {

	s.Client.On("GetUser", "id!").Return(s.Res, nil)

	c := `
		resource "baremetal_identity_user" "t" {
			name = "newname1"
			description = "desc!"
		}
	`

	c += testProviderConfig()

	u := *s.Res
	u.ID = "newid!"
	u.Name = "newname1"
	s.Client.On("CreateUser", "newnam1!", "desc!", (*baremetal.RetryTokenOptions)(nil)).
		Return(&u, nil)
	s.Client.On("GetUser", "newid!").Return(&u, nil)
	s.Client.On("DeleteUser", "newid!", (*baremetal.IfMatchOptions)(nil)).Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: c,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "newname1"),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestDeleteResourceIdentityUser() {
	s.Client.On("GetUser", "id!").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeleteUser", "id!", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceIdentityUserTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserTestSuite))
}
