// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
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
			name = "groupname"
			description = "group desc!"
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_identity_group.t"
	s.Res = &baremetal.Group{
		ID:            "id!",
		Name:          "groupname",
		Description:   "group desc!",
		CompartmentID: "cid!",
		State:         baremetal.ResourceActive,
		TimeCreated:   s.TimeCreated,
	}

}

func (s *ResourceIdentityGroupTestSuite) TestCreateResourceIdentityGroup() {

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

func (s *ResourceIdentityGroupTestSuite) TestCreateResourceIdentityGroupPolling() {
	s.Res.State = baremetal.ResourceCreating

	u := *s.Res
	u.State = baremetal.ResourceActive

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

	c := `
		resource "baremetal_identity_group" "t" {
			name = "groupname"
			description = "newdesc!"
		}
	`

	c += testProviderConfig()

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

	c := `
		resource "baremetal_identity_group" "t" {
			name = "groupname"
			description = "newdesc!"
		}
	`
	c += testProviderConfig()

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
				Check:       resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
			},
			{
				Config: c,
				Check:  resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
			},
		},
	})
}

func (s *ResourceIdentityGroupTestSuite) TestUpdateResourceIdentityGroupNameShouldCreateNew() {

	c := `
		resource "baremetal_identity_group" "t" {
			name = "groupname2"
			description = "desc!"
		}
	`

	c += testProviderConfig()

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
				Check:  resource.TestCheckResourceAttr(s.ResourceName, "name", "groupname2"),
			},
		},
	})
}

func (s *ResourceIdentityGroupTestSuite) TestDeleteResourceIdentityGroup() {

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

}

func TestResourceIdentityGroupTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityGroupTestSuite))
}
