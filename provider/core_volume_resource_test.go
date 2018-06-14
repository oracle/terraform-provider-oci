// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceCoreVolumeTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}`

	s.ResourceName = "oci_core_volume.t"
}

func (s *ResourceCoreVolumeTestSuite) TestCreateResourceCoreVolume_basic() {
	var resId, resId2 string
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_hydrated", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "source_details"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "volume_backup_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "source_details.#", "0"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_core_volume.t", "id")
						return err
					},
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-volume"),
					resource.TestCheckResourceAttr(s.ResourceName, "source_details.#", "0"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_hydrated", "true"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "source_details"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "volume_backup_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeLifecycleStateAvailable)),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_core_volume.t", "id")
						return err
					},
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
			// create a clone off the existing volume
			{
				Config: s.Config + `
				resource "oci_core_volume" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					display_name = "-tf-volume"
					size_in_gbs = 50
				}
				resource "oci_core_volume" "u" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					display_name = "-tf-volume-clone"
					size_in_gbs = 50
					source_details {
						type = "volumE" 	# case-insensitive
						id = "${oci_core_volume.t.id}"
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("oci_core_volume.u", "availability_domain"),
					resource.TestCheckResourceAttrSet("oci_core_volume.u", "id"),
					resource.TestCheckResourceAttrSet("oci_core_volume.u", "source_details.0.id"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "display_name", "-tf-volume-clone"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "source_details.#", "1"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "source_details.0.type", "volume"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "state", string(core.VolumeLifecycleStateAvailable)),
					resource.TestCheckResourceAttr("oci_core_volume.u", "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet("oci_core_volume.u", "is_hydrated"), // u (clone) isn't necessarily hydrated yet.
					resource.TestCheckNoResourceAttr("oci_core_volume.u", "volume_backup_id"),
					resource.TestCheckResourceAttrSet("oci_core_volume.u", "time_created"),
					resource.TestCheckResourceAttr("oci_core_volume.t", "is_hydrated", "true"), // t (source) should be hydrated now that u (clone) is available.
				),
			},
			// ensure that changing the case for source_details.?.type (polymorphic discriminator) is a no-op.
			{
				Config: s.Config + `
				resource "oci_core_volume" "t" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					display_name = "-tf-volume"
					size_in_gbs = 50
				}
				resource "oci_core_volume" "u" {
					availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
					compartment_id = "${var.compartment_id}"
					display_name = "-tf-volume-clone"
					size_in_gbs = 50
					source_details {
						type = "VoLume" 	# case-insensitive
						id = "${oci_core_volume.t.id}"
					}
				}`,
				PlanOnly: true,
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_hydrated", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeLifecycleStateAvailable)),
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
					size_in_gbs = 1024
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_gbs", "1024"),
					resource.TestMatchResourceAttr(s.ResourceName, "display_name", regexp.MustCompile(`[^\-tf\-volume]`)),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_core_volume.t", "id")
						if resId == resId2 {
							return fmt.Errorf("expected different ocid, got the same")
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
