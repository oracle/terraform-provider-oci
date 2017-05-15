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

type ResourceIdentityCompartmentsTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListCompartments
}

func (s *ResourceIdentityCompartmentsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_identity_compartments" "t" {
      compartment_id = "${var.compartment_id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_identity_compartments.t"

	b1 := baremetal.Compartment{
		ID:            "id",
		Name:          "compartmentname",
		CompartmentID: "compartment",
		Description:   "blah",
		State:         baremetal.ResourceActive,
		TimeCreated:   time.Now(),
	}

	b2 := b1
	b2.ID = "id2"

	s.List = &baremetal.ListCompartments{
		Compartments: []baremetal.Compartment{b1, b2},
	}
}

func (s *ResourceIdentityCompartmentsTestSuite) TestReadCompartments() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartments.0.id", "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartments.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "compartments.#", "2"),
				),
			},
		},
	},
	)
}

func TestResourceIdentityCompartmentsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityCompartmentsTestSuite))
}
