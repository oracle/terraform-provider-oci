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

type ResourceCoreVolumeTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Volume
	DeletedRes   *baremetal.Volume
}

func (s *ResourceCoreVolumeTestSuite) SetupTest() {
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
		data "baremetal_identity_availability_domains" "ADs" {
  			compartment_id = "${var.compartment_id}"
		}
		resource "baremetal_core_volume" "t" {
			availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
			compartment_id = "${var.compartment_id}"
			display_name = "display_name"
			size_in_mbs = 262144
		}
	`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_volume.t"

}

func (s *ResourceCoreVolumeTestSuite) TestCreateResourceCoreVolume() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),

					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					//resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", strconv.Itoa(s.Res.SizeInMBs)),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
				),
			},
		},
	})
}

func (s ResourceCoreVolumeTestSuite) TestUpdateVolumeDisplayName() {

	config := `
		data "baremetal_identity_availability_domains" "ADs" {
  			compartment_id = "${var.compartment_id}"
		}
		resource "baremetal_core_volume" "t" {
			availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
			compartment_id = "${var.compartment_id}"
			display_name = "new_display_name"
			size_in_mbs = 262144
		}
	`
	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "new_display_name"),
				),
			},
		},
	})
}

func (s ResourceCoreVolumeTestSuite) TestUpdateAvailabilityDomainForcesNewVolume() {

	config := `
		data "baremetal_identity_availability_domains" "ADs" {
  			compartment_id = "${var.compartment_id}"
		}
		resource "baremetal_core_volume" "t" {
			availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.1.name}"
			compartment_id = "${var.compartment_id}"
			size_in_mbs = 262144
		}
  `
	config += testProviderConfig()

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
				),
			},
		},
	})
}

func (s *ResourceCoreVolumeTestSuite) TestDeleteVolume() {

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

func TestResourceCoreVolumeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeTestSuite))
}
