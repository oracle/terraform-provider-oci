// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
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

func (s *ResourceCoreVolumeTestSuite) TestCreateResourceCoreVolume_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// specify size in MBs and GBs, expect error
			{
				Config: s.Config + `
				resource "oci_core_volume" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					size_in_mbs = 51200	
					size_in_gbs = 50
				}`,
				ExpectError: regexp.MustCompile("Megabytes and Gigabytes"),
			},
			// create volume, use default size
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
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-volume"),
				),
			},
			// explicit volume size in MBs, noop
			{
				Config: s.Config + `
				resource "oci_core_volume" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					display_name = "-tf-volume"
					size_in_mbs = 51200 //specify same size as default value, does nothing
				}`,
				ExpectNonEmptyPlan: false,
			},
			// migrate size_in_mbs to size_in_gbs
			{
				Config: s.Config + `
				resource "oci_core_volume" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					display_name = "-tf-volume"
					size_in_gbs = 50 //specify same size in GB, does nothing
				}`,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func (s *ResourceCoreVolumeTestSuite) TestCreateResourceCoreVolume_destructive() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// create volume, use default size
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
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_volume.t", "id")
						return err
					},
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
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_core_volume.t", "id")
						if resId == resId2 {
							return fmt.Errorf("Expected different ocid, got the same.")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceCoreVolumeTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeTestSuite))
}
