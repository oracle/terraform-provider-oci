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

type ResourceCoreVolumesTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVolumesTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.Config = `
	data "baremetal_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	resource "baremetal_core_volume" "t" {
		availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
		size_in_mbs = 262144
	}
    data "baremetal_core_volumes" "t" {
      availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
      compartment_id = "${var.compartment_id}"
      limit = 1
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_volumes.t"
}

func (s *ResourceCoreVolumesTestSuite) TestReadVolumes() {

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
					resource.TestCheckResourceAttrSet(s.ResourceName, "volumes.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volumes.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volumes.#"),
				),
			},
		},
	},
	)

}

func (s *ResourceCoreVolumesTestSuite) TestReadVolumesWithPagination() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "volumes.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volumes.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volumes.#"),
				),
			},
		},
	},
	)

}

func TestResourceCoreVolumesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumesTestSuite))
}
