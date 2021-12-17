// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	bootVolumeBackupFromSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"source_boot_volume_backup_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_boot_volume_backup.test_boot_volume_backup_copy.source_boot_volume_backup_id}`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: bootVolumeBackupFromSourceDataSourceFilterRepresentation}}
	bootVolumeBackupFromSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_boot_volume_backup.test_boot_volume_backup_copy.id}`}},
	}
	bootVolumeBackupWithSourceDetailsRepresentation = map[string]interface{}{
		"source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: bootVolumeBackupSourceDetailsRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}
	bootVolumeBackupSourceDetailsRepresentation = map[string]interface{}{}
	BootVolumeBackupCopyResourceDependencies    = BootVolumeBackupResourceDependencies
)

// issue-routing-tag: core/blockStorage
func TestResourceCoreBootVolumeBackup_copy(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreBootVolumeBackup_copy")
	defer httpreplay.SaveScenario()

	if utils.GetEnvSettingWithBlankDefault("source_region") == "" {
		t.Skip("Skipping TestCoreBootVolumeBackupResource_copy test because there is no source region specified")
	}
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameCopy := "oci_core_boot_volume_backup.test_boot_volume_backup_copy"
	datasourceName := "data.oci_core_boot_volume_backups.test_boot_volume_backups"

	err := createSourceBootVolumeBackupToCopy()
	if err != nil {
		t.Fatalf("Unable to Create source BootVolume and BootVolumeBackup to copy. Error: %v", err)
	}

	bootVolumeBackupSourceDetailsRepresentation = map[string]interface{}{
		"boot_volume_backup_id": acctest.Representation{RepType: acctest.Required, Create: bootVolumeBackupId},
		"region":                acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("source_region")},
		"kms_key_id":            acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
	}

	bootVolumeBackupWithSourceDetailsRepresentation = acctest.GetUpdatedRepresentationCopy("source_details", acctest.RepresentationGroup{RepType: acctest.Required, Group: bootVolumeBackupSourceDetailsRepresentation}, bootVolumeBackupWithSourceDetailsRepresentation)

	var resId string
	acctest.ResourceTest(t, testAccCheckCoreBootVolumeBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config +
				compartmentIdVariableStr + BootVolumeBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup_copy", acctest.Required, acctest.Create, bootVolumeBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "boot_volume_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "source_boot_volume_backup_id", bootVolumeBackupId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameCopy, "id")

					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BootVolumeBackupCopyResourceDependencies,
		},
		// verify Create from the backup with optionals
		{
			Config: config +
				compartmentIdVariableStr + BootVolumeBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup_copy", acctest.Optional, acctest.Create, bootVolumeBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "boot_volume_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "source_boot_volume_backup_id", bootVolumeBackupId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameCopy, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config +
				compartmentIdVariableStr + BootVolumeBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup_copy", acctest.Optional, acctest.Update, bootVolumeBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "boot_volume_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "source_boot_volume_backup_id", bootVolumeBackupId),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_boot_volume_backups", "test_boot_volume_backups", acctest.Optional, acctest.Update, bootVolumeBackupFromSourceDataSourceRepresentation) +
				compartmentIdVariableStr + BootVolumeBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup_copy", acctest.Optional, acctest.Update, bootVolumeBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(resourceNameCopy, "source_boot_volume_backup_id", bootVolumeBackupId),

				resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.0.display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.boot_volume_id"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.0.source_boot_volume_backup_id", bootVolumeBackupId),
			),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"source_details",
			},
			ResourceName: resourceNameCopy,
		},
	})
}

func createSourceBootVolumeBackupToCopy() error {
	sourceRegion := utils.GetEnvSettingWithBlankDefault("source_region")

	var err error
	instanceId, bootVolumeId, err = createBootVolumeInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion)
	if err != nil {
		log.Printf("[WARN] failed to createBootVolumeInRegion with the error %v", err)
		return err
	}

	bootVolumeBackupId, err = createBootVolumeBackupInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, &bootVolumeId)
	if err != nil {
		log.Printf("[WARN] failed to createBootVolumeBackupInRegion with the error %v", err)
		return err
	}

	return nil
}

func deleteSourceBootVolumeBackupToCopy() {
	sourceRegion := utils.GetEnvSettingWithBlankDefault("source_region")

	var err error
	err = deleteBootVolumeBackupInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, bootVolumeBackupId)
	if err != nil {
		log.Printf("[WARN] failed to deleteBootVolumeBackupInRegion with error %v", err)
	}

	err = terminateInstanceInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, instanceId)
	if err != nil {
		log.Printf("[WARN] failed to terminateInstanceInRegion with error %v", err)
	}
}
