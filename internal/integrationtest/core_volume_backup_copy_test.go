// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var (
	VolumeBackupCopyRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup_copy", acctest.Required, acctest.Create, CoreVolumeBackupWithSourceDetailsRepresentation)

	VolumeBackupCopyResourceDependencies = CoreVolumeBackupResourceDependencies
)

// issue-routing-tag: core/blockStorage
func TestResourceCoreVolumeBackup_copy(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreVolumeBackup_copy")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameCopy := "oci_core_volume_backup.test_volume_backup_copy"
	datasourceName := "data.oci_core_volume_backups.test_volume_backups"

	if utils.GetEnvSettingWithBlankDefault("source_region") == "" {
		t.Skip("Skipping TestCoreVolumeBackupResource_copy test because there is no source region specified")
	}

	err := createSourceVolumeBackupToCopy()
	if err != nil {
		t.Fatalf("Unable to Create source Volume and VolumeBackup to copy. Error: %v", err)
	}

	volumeBackupSourceDetailsRepresentation = map[string]interface{}{
		"volume_backup_id": acctest.Representation{RepType: acctest.Required, Create: volumeBackupId},
		"region":           acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("source_region")},
		"kms_key_id":       acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
	}
	CoreVolumeBackupWithSourceDetailsRepresentation = acctest.GetUpdatedRepresentationCopy("source_details", acctest.RepresentationGroup{RepType: acctest.Required, Group: volumeBackupSourceDetailsRepresentation}, CoreVolumeBackupWithSourceDetailsRepresentation)

	var resId string

	acctest.ResourceTest(t, testAccCheckCoreVolumeBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config +
				compartmentIdVariableStr + VolumeBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup_copy", acctest.Required, acctest.Create, CoreVolumeBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameCopy, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VolumeBackupCopyResourceDependencies,
		},
		// verify Create from the backup with optionals
		{
			Config: config +
				compartmentIdVariableStr + VolumeBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup_copy", acctest.Optional, acctest.Create, CoreVolumeBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "source_volume_backup_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameCopy, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config +
				compartmentIdVariableStr + VolumeBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup_copy", acctest.Optional, acctest.Update, CoreVolumeBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "source_volume_backup_id"),

				func(s *terraform.State) (err error) {
					resId2, err := acctest.FromInstanceState(s, resourceNameCopy, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_volume_backups", "test_volume_backups", acctest.Optional, acctest.Update, CoreVolumeBackupFromSourceDataSourceRepresentation) +
				compartmentIdVariableStr + VolumeBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup_copy", acctest.Optional, acctest.Update, CoreVolumeBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "source_volume_backup_id"),

				resource.TestCheckResourceAttr(datasourceName, "volume_backups.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.volume_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.source_volume_backup_id"),
			),
		},
		// verify resource import
		{
			Config:            config + VolumeBackupCopyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"source_details",
			},
			ResourceName: resourceNameCopy,
		},
	})
}

func createSourceVolumeBackupToCopy() error {
	sourceRegion := utils.GetEnvSettingWithBlankDefault("source_region")

	var err error
	volumeId, err = createVolumeInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion)
	if err != nil {
		log.Printf("[WARN] failed to createVolumeInRegion with the error %v", err)
		return err
	}

	volumeBackupId, err = createVolumeBackupInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, &volumeId)
	if err != nil {
		log.Printf("[WARN] failed to createVolumeBackupInRegion with the error %v", err)
		return err
	}

	return nil
}

func deleteSourceVolumeBackupToCopy() {
	sourceRegion := utils.GetEnvSettingWithBlankDefault("source_region")

	var err error
	err = deleteVolumeBackupInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, volumeBackupId)
	if err != nil {
		log.Printf("[WARN] failed to deleteVolumeBackupInRegion with error %v", err)
	}

	err = deleteVolumeInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, volumeId)
	if err != nil {
		log.Printf("[WARN] failed to deleteVolumeInRegion with error %v", err)
	}
}
