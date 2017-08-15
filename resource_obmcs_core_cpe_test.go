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

type ResourceCoreCpeTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Cpe
}

func (s *ResourceCoreCpeTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated = baremetal.Time{Time: time.Now()}
	s.Config = `

		resource "baremetal_core_cpe" "t" {
			compartment_id = "${var.compartment_id}"
			display_name = "displayname"
      			ip_address = "123.123.123.123"
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_cpe.t"

}

func (s *ResourceCoreCpeTestSuite) TestCreateResourceCoreCpe() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "displayname"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "ip_address", "123.123.123.123"),
				),
			},
		},
	})
}

func (s ResourceCoreCpeTestSuite) TestUpdateForcesNewCoreCpe() {

	updateForcingChangeConfig := `

  resource "baremetal_core_cpe" "t" {
    compartment_id = "${var.compartment_id}"
    display_name = "displayname"
    ip_address = "111.222.111.222"
  }

  `
	updateForcingChangeConfig += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: updateForcingChangeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "ip_address", "111.222.111.222"),
				),
			},
		},
	})

}

func (s *ResourceCoreCpeTestSuite) TestDeleteResourceCoreCpe() {

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

func TestResourceCoreCpeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreCpeTestSuite))
}
