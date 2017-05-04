// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUsersTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListUsers
}

func (s *ResourceIdentityUsersTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_identity_users" "t" {
      compartment_id = "compartment"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_identity_users.t"

	b1 := baremetal.User{
		ID:            "id",
		Name:          "username",
		CompartmentID: "compartment",
		Description:   "blah",
		State:         baremetal.ResourceActive,
		TimeCreated:   time.Now(),
	}

	b2 := b1
	b2.ID = "id2"

	s.List = &baremetal.ListUsers{
		Users: []baremetal.User{b1, b2},
	}
}

func (s *ResourceIdentityUsersTestSuite) TestReadUsers() {
	s.Client.On("ListUsers", (*baremetal.ListOptions)(nil)).Return(s.List, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "users.0.id", "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "users.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "users.#", "2"),
				),
			},
		},
	},
	)
}

func TestResourceIdentityUsersTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUsersTestSuite))
}
