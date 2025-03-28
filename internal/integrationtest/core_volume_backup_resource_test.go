// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreVolumeBackupTestSuite struct {
	suite.Suite
	Providers    map[string]*schema.Provider
	Config       string
	ResourceName string
}

func (s *ResourceCoreVolumeBackupTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
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
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "expiration_time"),
					resource.TestCheckResourceAttr(s.ResourceName, "source_type", string(core.VolumeBackupSourceTypeManual)),
					resource.TestCheckResourceAttr(s.ResourceName, "type", string(core.VolumeBackupTypeIncremental)),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeBackupLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "unique_size_in_mbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "unique_size_in_gbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_request_received"),
					func(ts *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// verify Update
			{
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
						display_name = "-tf-volume-backup"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-volume-backup"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "expiration_time"),
					resource.TestCheckResourceAttr(s.ResourceName, "source_type", string(core.VolumeBackupSourceTypeManual)),
					resource.TestCheckResourceAttr(s.ResourceName, "type", string(core.VolumeBackupTypeIncremental)),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeBackupLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "unique_size_in_mbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "unique_size_in_gbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_request_received"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId2 != resId {
							return fmt.Errorf("expected same volume bakcup ocid, got different")
						}
						return err
					},
				),
			},
			// verify ForceNew when changing the backup type to FULL
			{
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
						display_name = "-tf-volume-backup"
						type = "FULL"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-volume-backup"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "expiration_time"),
					resource.TestCheckResourceAttr(s.ResourceName, "source_type", string(core.VolumeBackupSourceTypeManual)),
					resource.TestCheckResourceAttr(s.ResourceName, "type", string(core.VolumeBackupTypeFull)),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeBackupLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "unique_size_in_mbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "unique_size_in_gbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_request_received"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId2 == resId {
							return fmt.Errorf("expected different volume backup ocid, got same")
						}

						resId = resId2
						return err
					},
				),
			},
			// verify conventional restore
			{
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
						display_name = "-tf-volume-backup"
						type = "FULL"
					}
					resource "oci_core_volume" "t2" {
						availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
						compartment_id = "${var.compartment_id}"
						display_name = "-tf-volume-restored"
						size_in_gbs = 50
						volume_backup_id = "${oci_core_volume_backup.t.id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "volume_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VolumeBackupLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr(s.ResourceName, "size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "unique_size_in_mbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "unique_size_in_gbs"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_request_received"),
					resource.TestCheckResourceAttr("oci_core_volume.t2", "display_name", "-tf-volume-restored"),
					resource.TestCheckResourceAttrSet("oci_core_volume.t2", "source_details.0.id"),
					resource.TestCheckResourceAttr("oci_core_volume.t2", "source_details.0.type", "volumeBackup"),
					resource.TestCheckResourceAttr("oci_core_volume.t2", "state", string(core.VolumeLifecycleStateAvailable)),
					resource.TestCheckResourceAttr("oci_core_volume.t2", "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr("oci_core_volume.t2", "size_in_gbs", "50"),
					// Only set during "Create" scenarios
					resource.TestCheckNoResourceAttr("oci_core_volume.t2", "time_request_received"),
					func(ts *terraform.State) (err error) {
						var backupId, volBackupId string
						if backupId, err = acctest.FromInstanceState(ts, s.ResourceName, "id"); err == nil {
							if volBackupId, err = acctest.FromInstanceState(ts, "oci_core_volume.t2", "volume_backup_id"); err == nil {
								if volBackupId != backupId {
									return fmt.Errorf("volume created from different backup than expected")
								}
							}
							return err
						}
						return err
					},
				),
			},
			// verify clone from backup
			{
				Config: s.Config + `
					resource "oci_core_volume_backup" "t" {
						volume_id = "${oci_core_volume.t.id}"
						display_name = "-tf-volume-backup"
						type = "FULL"
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet("oci_core_volume.u", "source_details.0.id"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "display_name", "-tf-volume-clone"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "source_details.0.type", "volumeBackup"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "state", string(core.VolumeLifecycleStateAvailable)),
					resource.TestCheckResourceAttr("oci_core_volume.u", "size_in_mbs", "51200"),
					resource.TestCheckResourceAttr("oci_core_volume.u", "size_in_gbs", "50"),
					// Only set during "Create" scenarios
					resource.TestCheckNoResourceAttr("oci_core_volume.u", "time_request_received"),
					// Only present if specific in configuration
					resource.TestCheckNoResourceAttr("oci_core_volume.u", "volume_backup_id"),
				),
			},
		},
	})
}

// issue-routing-tag: core/blockStorage
func TestResourceCoreVolumeBackupTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreVolumeBackupTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceCoreVolumeBackupTestSuite))
}
