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

type ResourceIdentityGroupTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
	Res          *baremetal.Group
}

func (s *ResourceIdentityGroupTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	configfn := func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	}

	s.Provider = Provider(configfn)
	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = `
		resource "baremetal_identity_group" "t" {
			name = "name!"
			description = "desc!"
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_identity_group.t"
	s.Res = &baremetal.Group{
		ID:            "id!",
		Name:          "name!",
		Description:   "desc!",
		CompartmentID: "cid!",
		State:         baremetal.ResourceActive,
		TimeCreated:   s.TimeCreated,
	}
	s.Client.On("CreateGroup", "name!", "desc!", (*baremetal.RetryTokenOptions)(nil)).Return(s.Res, nil)
	s.Client.On("DeleteGroup", "id!", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceIdentityGroupTestSuite) TestCreateResourceIdentityGroup() {
	s.Client.On("GetGroup", "id!").Return(s.Res, nil)

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
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
				),
			},
		},
	})
}

func (s *ResourceIdentityGroupTestSuite) TestCreateResourceIdentityGroupPolling() {
	s.Res.State = baremetal.ResourceCreating
	s.Client.On("GetGroup", "id!").Return(s.Res, nil).Once()

	u := *s.Res
	u.State = baremetal.ResourceActive
	s.Client.On("GetGroup", "id!").Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check:             resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceActive),
			},
		},
	})
}

func (s *ResourceIdentityGroupTestSuite) TestUpdateResourceIdentityGroupDescription() {
	s.Client.On("GetGroup", "id!").Return(s.Res, nil).Twice()

	c := `
		resource "baremetal_identity_group" "t" {
			name = "name!"
			description = "newdesc!"
		}
	`

	c += testProviderConfig

	u := *s.Res
	u.Description = "newdesc!"
	opts := &baremetal.UpdateIdentityOptions{}
	opts.Description = "newdesc!"
	s.Client.On("UpdateGroup", "id!", "newdesc!", opts).
		Return(&u, nil)
	s.Client.On("GetGroup", "id!").Return(&u, nil)

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

func (s *ResourceIdentityGroupTestSuite) TestFailedUpdateResourceIdentityGroupDescription() {
	s.Client.On("GetGroup", "id!").Return(s.Res, nil).Times(3)

	c := `
		resource "baremetal_identity_group" "t" {
			name = "name!"
			description = "newdesc!"
		}
	`
	c += testProviderConfig

	opts := &baremetal.UpdateIdentityOptions{}
	opts.Description = "newdesc!"
	s.Client.On("UpdateGroup", "id!", opts).
		Return(nil, errors.New("FAILED!")).Once()

	u := *s.Res
	u.Description = "newdesc!"
	s.Client.On("UpdateGroup", "id!", opts).
		Return(&u, nil)
	s.Client.On("GetGroup", "id!").Return(&u, nil)

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
				Check:       resource.TestCheckResourceAttr(s.ResourceName, "description", "desc!"),
			},
			{
				Config: c,
				Check:  resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
			},
		},
	})
}

func (s *ResourceIdentityGroupTestSuite) TestUpdateResourceIdentityGroupNameShouldCreateNew() {
	s.Client.On("GetGroup", "id!").Return(s.Res, nil)

	c := `
		resource "baremetal_identity_group" "t" {
			name = "newname!"
			description = "desc!"
		}
	`

	c += testProviderConfig

	u := *s.Res
	u.ID = "newid!"
	u.Name = "newname!"
	s.Client.On("CreateGroup", "newname!", "desc!", (*baremetal.RetryTokenOptions)(nil)).Return(&u, nil)
	s.Client.On("GetGroup", "newid!").Return(&u, nil)
	s.Client.On("DeleteGroup", "newid!", (*baremetal.IfMatchOptions)(nil)).Return(nil)

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
				Check:  resource.TestCheckResourceAttr(s.ResourceName, "name", "newname!"),
			},
		},
	})
}

func (s *ResourceIdentityGroupTestSuite) TestDeleteResourceIdentityGroup() {
	s.Client.On("GetGroup", "id!").Return(s.Res, nil)

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

	s.Client.AssertCalled(s.T(), "DeleteGroup", "id!", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceIdentityGroupTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityGroupTestSuite))
}
