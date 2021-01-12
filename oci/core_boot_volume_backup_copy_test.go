// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	bootVolumeBackupFromSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                 Representation{repType: Optional, create: `displayName`},
		"source_boot_volume_backup_id": Representation{repType: Optional, create: `${oci_core_boot_volume_backup.test_boot_volume_backup_copy.source_boot_volume_backup_id}`},
		"state":                        Representation{repType: Optional, create: `AVAILABLE`},
		"filter":                       RepresentationGroup{Required, bootVolumeBackupFromSourceDataSourceFilterRepresentation}}
	bootVolumeBackupFromSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_boot_volume_backup.test_boot_volume_backup_copy.id}`}},
	}
	bootVolumeBackupWithSourceDetailsRepresentation = map[string]interface{}{
		"source_details": RepresentationGroup{Required, bootVolumeBackupSourceDetailsRepresentation},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"display_name":   Representation{repType: Optional, create: `displayName`},
	}
	bootVolumeBackupSourceDetailsRepresentation = map[string]interface{}{}
	BootVolumeBackupCopyResourceDependencies    = BootVolumeBackupResourceDependencies + generateResourceFromRepresentationMap("oci_kms_key", "test_key", Required, Create, keyRepresentation)
)

func TestResourceCoreBootVolumeBackup_copy(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreBootVolumeBackup_copy")
	defer httpreplay.SaveScenario()

	if getEnvSettingWithBlankDefault("source_region") == "" {
		t.Skip("Skipping TestCoreBootVolumeBackupResource_copy test because there is no source region specified")
	}
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameCopy := "oci_core_boot_volume_backup.test_boot_volume_backup_copy"
	datasourceName := "data.oci_core_boot_volume_backups.test_boot_volume_backups"

	err := createSourceBootVolumeBackupToCopy()
	if err != nil {
		t.Fatalf("Unable to create source BootVolume and BootVolumeBackup to copy. Error: %v", err)
	}

	bootVolumeBackupSourceDetailsRepresentation = map[string]interface{}{
		"boot_volume_backup_id": Representation{repType: Required, create: bootVolumeBackupId},
		"region":                Representation{repType: Required, create: getEnvSettingWithBlankDefault("source_region")},
		"kms_key_id":            Representation{repType: Optional, create: `${oci_kms_key.test_key.id}`},
	}

	bootVolumeBackupWithSourceDetailsRepresentation = getUpdatedRepresentationCopy("source_details", RepresentationGroup{Required, bootVolumeBackupSourceDetailsRepresentation}, bootVolumeBackupWithSourceDetailsRepresentation)

	var resId string
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreBootVolumeBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config +
					compartmentIdVariableStr + BootVolumeBackupCopyResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup_copy", Required, Create, bootVolumeBackupWithSourceDetailsRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "boot_volume_id"),
					resource.TestCheckResourceAttr(resourceNameCopy, "source_boot_volume_backup_id", bootVolumeBackupId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceNameCopy, "id")

						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BootVolumeBackupCopyResourceDependencies,
			},
			// verify create from the backup with optionals
			{
				Config: config +
					compartmentIdVariableStr + BootVolumeBackupCopyResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup_copy", Optional, Create, bootVolumeBackupWithSourceDetailsRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
					resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "boot_volume_id"),
					resource.TestCheckResourceAttr(resourceNameCopy, "source_boot_volume_backup_id", bootVolumeBackupId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceNameCopy, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config +
					compartmentIdVariableStr + BootVolumeBackupCopyResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup_copy", Optional, Update, bootVolumeBackupWithSourceDetailsRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
					resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "boot_volume_id"),
					resource.TestCheckResourceAttr(resourceNameCopy, "source_boot_volume_backup_id", bootVolumeBackupId),

					func(s *terraform.State) (err error) {
						resId2, err := fromInstanceState(s, resourceNameCopy, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_boot_volume_backups", "test_boot_volume_backups", Optional, Update, bootVolumeBackupFromSourceDataSourceRepresentation) +
					compartmentIdVariableStr + BootVolumeBackupCopyResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup_copy", Optional, Update, bootVolumeBackupWithSourceDetailsRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}

func createSourceBootVolumeBackupToCopy() error {
	sourceRegion := getEnvSettingWithBlankDefault("source_region")

	var err error
	instanceId, bootVolumeId, err = createBootVolumeInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion)
	if err != nil {
		log.Printf("[WARN] failed to createBootVolumeInRegion with the error %v", err)
		return err
	}

	bootVolumeBackupId, err = createBootVolumeBackupInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion, &bootVolumeId)
	if err != nil {
		log.Printf("[WARN] failed to createBootVolumeBackupInRegion with the error %v", err)
		return err
	}

	return nil
}

func deleteSourceBootVolumeBackupToCopy() {
	sourceRegion := getEnvSettingWithBlankDefault("source_region")

	var err error
	err = deleteBootVolumeBackupInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion, bootVolumeBackupId)
	if err != nil {
		log.Printf("[WARN] failed to deleteBootVolumeBackupInRegion with error %v", err)
	}

	err = terminateInstanceInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion, instanceId)
	if err != nil {
		log.Printf("[WARN] failed to terminateInstanceInRegion with error %v", err)
	}
}
