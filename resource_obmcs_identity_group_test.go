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

	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = testProviderConfig() + `
		resource "baremetal_identity_group" "t" {
			name = "-tf-group"
			description = "automated test group"
		}
	`

	s.ResourceName = "baremetal_identity_group.t"
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
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-group"),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test group"),

					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceActive),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
			{
				Config: testProviderConfig() + `
					resource "baremetal_identity_group" "t" {
						name = "-tf-group"
						description = "automated test group (updated)"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test group (updated)"),
				),
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
