// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/oci-go-sdk/v65/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVolumeBackupTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
}

func (s *DatasourceCoreVolumeBackupTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	resource "oci_core_volume" "t" {
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		compartment_id = "${var.compartment_id}"
	}
	resource "oci_core_volume_backup" "t" {
		volume_id = "${oci_core_volume.t.id}"
		display_name = "-tf-volume-backup"
	}`
	s.ResourceName = "data.oci_core_volume_backups.t"
}

func (s *DatasourceCoreVolumeBackupTestSuite) TestAccDatasourceCoreVolumeBackup_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + `
				data "oci_core_volume_backups" "t" {
					compartment_id = "${var.compartment_id}"
					volume_id = "${oci_core_volume.t.id}"
					filter {
						name = "id"
						values = ["${oci_core_volume_backup.t.id}"]
					}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "volume_backups.0.id", "oci_core_volume_backup.t", "id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "volume_backups.0.volume_id", "oci_core_volume_backup.t", "volume_id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "volume_backups.0.time_created", "oci_core_volume_backup.t", "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.display_name", "-tf-volume-backup"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.expiration_time", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.source_type", string(core.VolumeBackupSourceTypeManual)),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.type", string(core.VolumeBackupTypeIncremental)),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.state", string(core.VolumeBackupLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.0.size_in_gbs", "50"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "volume_backups.0.unique_size_in_mbs", "oci_core_volume_backup.t", "unique_size_in_mbs"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "volume_backups.0.unique_size_in_gbs", "oci_core_volume_backup.t", "unique_size_in_gbs"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "volume_backups.0.time_request_received", "oci_core_volume_backup.t", "time_request_received"),
				),
			},
			// Server-side filtering tests.
			{
				Config: s.Config + `
				data "oci_core_volume_backups" "t" {
					compartment_id = "${var.compartment_id}"
					volume_id = "${oci_core_volume.t.id}"
					display_name = "${oci_core_volume_backup.t.display_name}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "volume_backups.0.id", "oci_core_volume_backup.t", "id"),
				),
			},
			{
				Config: s.Config + `
				data "oci_core_volume_backups" "t" {
					compartment_id = "${var.compartment_id}"
					volume_id = "${oci_core_volume.t.id}"
					state = "` + string(core.VolumeBackupLifecycleStateAvailable) + `"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "volume_backups.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "volume_backups.0.id", "oci_core_volume_backup.t", "id"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: core/blockStorage
func TestDatasourceCoreVolumeBackupTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreVolumeBackupTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreVolumeBackupTestSuite))
}
