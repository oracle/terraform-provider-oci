// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeBackupTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Res          *baremetal.VolumeBackup
	DeletedRes   *baremetal.VolumeBackup
}

func (s *ResourceCoreVolumeBackupTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
		data "oci_identity_availability_domains" "ADs" {
  			compartment_id = "${var.compartment_id}"
		}
		resource "oci_core_volume" "t" {
			availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
			compartment_id = "${var.compartment_id}"
			display_name = "-tf-volume"
			size_in_gbs = 50
		}`
	s.ResourceName = "oci_core_volume_backup.t"
}

func (s *ResourceCoreVolumeBackupTestSuite) TestAccResourceCoreVolumeBackup_basic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceAvailable),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "unique_size_in_mbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "unique_size_in_gbs"),
				),
			},
			// verify update
			{
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
						display_name = "-tf-volume-backup"
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-volume-backup"),
				),
			},
			// verify conventional restore
			{
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
						display_name = "-tf-volume-backup"
					}
					resource "oci_core_volume" "t2" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-volume-restored"
						size_in_gbs = 50
						volume_backup_id = "${oci_core_volume_backup.t.id}"
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("oci_core_volume.t2", "display_name", "-tf-volume-restored"),
				),
			},
			// verify clone from backup
			{
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
						display_name = "-tf-volume-backup"
					}
					resource "oci_core_volume" "u" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-volume-clone"
						size_in_gbs = 50
						source_details {
							type = "volumeBackup"
							id = "${oci_core_volume_backup.t.id}"
						}
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("oci_core_volume.u", "source_details.0.id"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "display_name", "-tf-volume-clone"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "source_details.0.type", "volumeBackup"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "state", baremetal.ResourceAvailable),
				),
			},
		},
	})
}

func TestResourceCoreVolumeBackupTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVolumeBackupTestSuite))
}
