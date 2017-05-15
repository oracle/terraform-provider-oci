// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreInstancesTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreInstancesTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
    data "baremetal_core_instances" "s" {
      compartment_id = "${var.compartment_id}"
      availability_domain = "availabilityid"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_instances.s"
}

func (s *ResourceCoreInstancesTestSuite) TestResourceListInstances() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.1.id", "id2"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.#", "2"),
				),
			},
		},
	},
	)

}

func (s *ResourceCoreInstancesTestSuite) TestResourceListInstancesPaged() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(s.ResourceName, "availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.availability_domain", "availabilityid"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.0.id", "id1"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.3.id", "id4"),
					resource.TestCheckResourceAttr(s.ResourceName, "instances.#", "4"),
				),
			},
		},
	},
	)

}

func TestResourceCoreInstancesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInstancesTestSuite))
}
