// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
	"testing"
	"regexp"
)

type ResourceCoreVolumeTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Res          *baremetal.Volume
	DeletedRes   *baremetal.Volume
}

func (s *ResourceCoreVolumeTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
		data "oci_identity_availability_domains" "ADs" {
			compartment_id = "${var.compartment_id}"
		}`

	s.ResourceName = "oci_core_volume.t"
}

func (s *ResourceCoreVolumeTestSuite) TestCreateResourceCoreVolume() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// create volume
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_volume" "t" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
				),
			},
			// update volume
			{
				Config: s.Config + `
					resource "oci_core_volume" "t" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-volume"
						size_in_mbs = 51200
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-volume"),
				),
			},
			// verify changing volume size is destructive
			{
				Config: s.Config + `
					resource "oci_core_volume" "t" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
						size_in_mbs = 102400
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", "102400"),
					resource.TestMatchResourceAttr(s.ResourceName, "display_name", regexp.MustCompile(`[^\-tf\-volume]`)),
				),
			},
		},
	})
}

func TestResourceCoreVolumeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeTestSuite))
}
