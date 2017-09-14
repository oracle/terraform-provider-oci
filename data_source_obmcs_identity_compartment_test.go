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

type DatasourceIdentityCompartmentsTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListCompartments
}

func (s *DatasourceIdentityCompartmentsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = `
    data "oci_identity_compartments" "t" {
      compartment_id = "${var.compartment_id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.oci_identity_compartments.t"

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

func (s *DatasourceIdentityCompartmentsTestSuite) TestReadCompartments() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartments.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartments.#"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityCompartmentsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityCompartmentsTestSuite))
}
