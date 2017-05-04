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

type ResourceIdentityGroupsTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListGroups
}

func (s *ResourceIdentityGroupsTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_identity_groups" "t" {
      compartment_id = "compartment"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "data.baremetal_identity_groups.t"

	b1 := baremetal.Group{
		ID:            "id",
		Name:          "groupname",
		CompartmentID: "compartment",
		Description:   "blah",
		State:         baremetal.ResourceActive,
		TimeCreated:   time.Now(),
	}

	b2 := b1
	b2.ID = "id2"

	s.List = &baremetal.ListGroups{
		Groups: []baremetal.Group{b1, b2},
	}
}

func (s *ResourceIdentityGroupsTestSuite) TestReadGroups() {
	s.Client.On("ListGroups", (*baremetal.ListOptions)(nil)).Return(s.List, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "groups.0.id", "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "groups.#", "2"),
				),
			},
		},
	},
	)
}

func TestResourceIdentityGroupsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityGroupsTestSuite))
}
