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
	VolumeGroupBackupCopyResourceDependencies = SourceVolumeListDependency +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	volumeGroupBackupFromSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"filter":         RepresentationGroup{Required, volumeGroupBackupFromSourceDataSourceFilterRepresentation}}

	volumeGroupBackupFromSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume_group_backup.test_volume_group_backup_copy.id}`}},
	}

	volumeGroupBackupWithSourceDetailsRepresentation = map[string]interface{}{
		"source_details": RepresentationGroup{Required, volumeGroupBackupSourceDetailsRepresentation},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
	}

	volumeGroupBackupId, volumeGroupId           string
	volumeGroupBackupSourceDetailsRepresentation = map[string]interface{}{}
)

func TestResourceCoreVolumeGroupBackup_copy(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreVolumeGroupBackup_copy")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameCopy := "oci_core_volume_group_backup.test_volume_group_backup_copy"
	datasourceName := "data.oci_core_volume_group_backups.test_volume_group_backups"

	if getEnvSettingWithBlankDefault("source_region") == "" {
		t.Skip("Skipping TestResourceCoreVolumeGroupBackup_copy test because there is no source region specified")
	}

	err := createSourceVolumeGroupBackupToCopy()
	if err != nil {
		t.Fatalf("Unable to create source Volume group and VolumeGroupBackup to copy. Error: %v", err)
	}

	volumeGroupBackupSourceDetailsRepresentation = map[string]interface{}{
		"volume_group_backup_id": Representation{repType: Required, create: volumeGroupBackupId},
		"region":                 Representation{repType: Required, create: getEnvSettingWithBlankDefault("source_region")},
		"kms_key_id":             Representation{repType: Optional, create: getEnvSettingWithBlankDefault("kms_key_ocid")},
	}

	volumeGroupBackupWithSourceDetailsRepresentation = getUpdatedRepresentationCopy("source_details", RepresentationGroup{Required, volumeGroupBackupSourceDetailsRepresentation}, volumeGroupBackupWithSourceDetailsRepresentation)

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeGroupBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config +
					compartmentIdVariableStr + VolumeGroupBackupCopyResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup_copy", Required, Create, volumeGroupBackupWithSourceDetailsRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_group_id"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceNameCopy, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeGroupBackupCopyResourceDependencies,
			},
			// verify create from the backup with optionals
			{
				Config: config +
					compartmentIdVariableStr + VolumeGroupBackupCopyResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup_copy", Optional, Create, volumeGroupBackupWithSourceDetailsRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
					resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_group_id"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "source_volume_group_backup_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceNameCopy, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config +
					compartmentIdVariableStr + VolumeGroupBackupCopyResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup_copy", Optional, Update, volumeGroupBackupWithSourceDetailsRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
					resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_group_id"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "source_volume_group_backup_id"),

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
					generateDataSourceFromRepresentationMap("oci_core_volume_group_backups", "test_volume_group_backups", Optional, Update, volumeGroupBackupFromSourceDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeGroupBackupCopyResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Required, Create, volumeBackupRepresentation) +
					generateResourceFromRepresentationMap("oci_core_volume", "test_volume", Required, Create, volumeRepresentation) +
					generateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup_copy", Optional, Update, volumeGroupBackupWithSourceDetailsRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.type"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.volume_group_id"),
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

func createSourceVolumeGroupBackupToCopy() error {
	sourceRegion := getEnvSettingWithBlankDefault("source_region")

	var err error
	volumeId, err = createVolumeInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion)
	if err != nil {
		log.Printf("[WARN] failed to createVolumeInRegion with the error %v", err)
		return err
	}

	volumeGroupId, err = createVolumeGroupInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion, &volumeId)
	if err != nil {
		log.Printf("[WARN] failed to createVolumeGroupInRegion with the error %v", err)
		return err
	}

	volumeGroupBackupId, err = createVolumeGroupBackupInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion, &volumeGroupId)
	if err != nil {
		log.Printf("[WARN] failed to createVolumeGroupBackupInRegion with the error %v", err)
		return err
	}

	return nil
}

func deleteSourceVolumeGroupBackupToCopy() {
	sourceRegion := getEnvSettingWithBlankDefault("source_region")

	var err error
	err = deleteVolumeGroupBackupInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion, volumeGroupBackupId)
	if err != nil {
		log.Printf("[WARN] failed to deleteVolumeGroupBackupInRegion with error %v", err)
	}

	err = deleteVolumeGroupInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion, volumeGroupId)
	if err != nil {
		log.Printf("[WARN] failed to deleteVolumeGroupInRegion with error %v", err)
	}
}
