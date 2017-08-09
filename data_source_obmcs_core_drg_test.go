// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-baremetal/client"
)

type ResourceCoreDrgsTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreDrgsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
	resource "baremetal_core_drg" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_drgs.t"
}

func (s *ResourceCoreDrgsTestSuite) TestReadDrgs() {

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
				data "baremetal_core_drgs" "t" {
					compartment_id = "${var.compartment_id}"
					limit = 1
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "drgs.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "drgs.#"),
				),
			},
		},
	},
	)

}

func TestResourceCoreDrgsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreDrgsTestSuite))
}
