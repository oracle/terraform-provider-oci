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

type ResourceIdentityGroupsTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListGroups
}

func (s *ResourceIdentityGroupsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
	resource "baremetal_identity_group" "t" {
		name = "groupname"
		description = "group desc!"
	}
  `
	s.Config += testProviderConfig()
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

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				    data "baremetal_identity_groups" "t" {
				      compartment_id = "${var.compartment_id}"
				    }`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.#"),
				),
			},
		},
	},
	)
}

func TestResourceIdentityGroupsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityGroupsTestSuite))
}
