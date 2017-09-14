// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVolumeBackupTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListVolumeBackups
}

func (s *DatasourceCoreVolumeBackupTestSuite) SetupTest() {
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
		display_name = "display_name"
		size_in_mbs = 51200
	}
	resource "oci_core_volume_backup" "t" {
		volume_id = "${oci_core_volume.t.id}"
		display_name = "display_name"
	}`
	s.ResourceName = "data.oci_core_volume_backups.t"
}

func (s *DatasourceCoreVolumeBackupTestSuite) TestAccDatasourceCoreVolumeBackup_basic() {
	resource.Test(s.T(), resource.TestCase{
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
				data "oci_core_volume_backups" "t" {
					compartment_id = "${var.compartment_id}"
					volume_id = "${oci_core_volume.t.id}"
					limit = 1
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_backups.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.#", "1"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreVolumeBackupTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVolumeBackupTestSuite))
}
