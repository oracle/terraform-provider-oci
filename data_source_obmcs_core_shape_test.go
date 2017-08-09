// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreShapeTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreShapeTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_identity_availability_domains" "t" {
      compartment_id = "${var.compartment_id}"
    }
    data "baremetal_core_shape" "s" {
      compartment_id = "${var.compartment_id}"
      availability_domain = "${data.baremetal_identity_availability_domains.t.availability_domains.0.name}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_shape.s"

}

func (s *ResourceCoreShapeTestSuite) TestResourceReadCoreShape() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shapes.0.name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shapes.1.name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shapes.#"),
				),
			},
		},
	},
	)
}

func TestResourceCoreShapeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreShapeTestSuite))
}
