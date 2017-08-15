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

type ResourceIdentityCompartmentTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
	Res          *baremetal.Compartment
}

func (s *ResourceIdentityCompartmentTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	configfn := func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	}

	s.Provider = Provider(configfn)
	p := s.Provider.(*schema.Provider)
	res := p.ResourcesMap["baremetal_identity_compartment"]
	res.Delete = func(d *schema.ResourceData, m interface{}) (e error) {
		return nil
	}
	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = `
		resource "baremetal_identity_compartment" "t" {
			name = "test-compartment"
			description = "automated test compartment"
		}
	`
	s.Config += testProviderConfig()
	s.ResourceName = "baremetal_identity_compartment.t"
	s.Res = &baremetal.Compartment{
		ID:            "id!",
		Name:          "test-compartment",
		Description:   "automated test compartment",
		CompartmentID: "cid!",
		State:         baremetal.ResourceActive,
		TimeCreated:   s.TimeCreated,
	}

}

func (s *ResourceIdentityCompartmentTestSuite) TestCreateResourceIdentityCompartment() {

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
				),
			},
		},
	})
}

func TestResourceIdentityCompartmentTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityCompartmentTestSuite))
}
