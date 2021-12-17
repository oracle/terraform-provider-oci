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
	VolumeGroupBackupCopyResourceDependencies = SourceVolumeListDependency +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	volumeGroupBackupFromSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: volumeGroupBackupFromSourceDataSourceFilterRepresentation}}

	volumeGroupBackupFromSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_volume_group_backup.test_volume_group_backup_copy.id}`}},
	}

	volumeGroupBackupWithSourceDetailsRepresentation = map[string]interface{}{
		"source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: volumeGroupBackupSourceDetailsRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}

	volumeGroupBackupId, volumeGroupId           string
	volumeGroupBackupSourceDetailsRepresentation = map[string]interface{}{}
)

func TestResourceCoreVolumeGroupBackup_copy(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreVolumeGroupBackup_copy")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameCopy := "oci_core_volume_group_backup.test_volume_group_backup_copy"
	datasourceName := "data.oci_core_volume_group_backups.test_volume_group_backups"

	if utils.GetEnvSettingWithBlankDefault("source_region") == "" {
		t.Skip("Skipping TestResourceCoreVolumeGroupBackup_copy test because there is no source region specified")
	}

	err := createSourceVolumeGroupBackupToCopy()
	if err != nil {
		t.Fatalf("Unable to Create source Volume Group and VolumeGroupBackup to copy. Error: %v", err)
	}

	volumeGroupBackupSourceDetailsRepresentation = map[string]interface{}{
		"volume_group_backup_id": acctest.Representation{RepType: acctest.Required, Create: volumeGroupBackupId},
		"region":                 acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("source_region")},
		"kms_key_id":             acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("kms_key_ocid")},
	}

	volumeGroupBackupWithSourceDetailsRepresentation = acctest.GetUpdatedRepresentationCopy("source_details", acctest.RepresentationGroup{RepType: acctest.Required, Group: volumeGroupBackupSourceDetailsRepresentation}, volumeGroupBackupWithSourceDetailsRepresentation)

	var resId string

	acctest.ResourceTest(t, testAccCheckCoreVolumeGroupBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config +
				compartmentIdVariableStr + VolumeGroupBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup_copy", acctest.Required, acctest.Create, volumeGroupBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_group_id"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameCopy, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VolumeGroupBackupCopyResourceDependencies,
		},
		// verify Create from the backup with optionals
		{
			Config: config +
				compartmentIdVariableStr + VolumeGroupBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup_copy", acctest.Optional, acctest.Create, volumeGroupBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_group_id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "source_volume_group_backup_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameCopy, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config +
				compartmentIdVariableStr + VolumeGroupBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup_copy", acctest.Optional, acctest.Update, volumeGroupBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_group_id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "source_volume_group_backup_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_volume_group_backups", "test_volume_group_backups", acctest.Optional, acctest.Update, volumeGroupBackupFromSourceDataSourceRepresentation) +
				compartmentIdVariableStr + VolumeGroupBackupCopyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", acctest.Required, acctest.Create, volumeBackupRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume", "test_volume", acctest.Required, acctest.Create, volumeRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup_copy", acctest.Optional, acctest.Update, volumeGroupBackupWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	})

}

func createSourceVolumeGroupBackupToCopy() error {
	sourceRegion := utils.GetEnvSettingWithBlankDefault("source_region")

	var err error
	volumeId, err = createVolumeInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion)
	if err != nil {
		log.Printf("[WARN] failed to createVolumeInRegion with the error %v", err)
		return err
	}

	volumeGroupId, err = createVolumeGroupInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, &volumeId)
	if err != nil {
		log.Printf("[WARN] failed to createVolumeGroupInRegion with the error %v", err)
		return err
	}

	volumeGroupBackupId, err = createVolumeGroupBackupInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, &volumeGroupId)
	if err != nil {
		log.Printf("[WARN] failed to createVolumeGroupBackupInRegion with the error %v", err)
		return err
	}

	return nil
}

func deleteSourceVolumeGroupBackupToCopy() {
	sourceRegion := utils.GetEnvSettingWithBlankDefault("source_region")

	var err error
	err = deleteVolumeGroupBackupInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, volumeGroupBackupId)
	if err != nil {
		log.Printf("[WARN] failed to deleteVolumeGroupBackupInRegion with error %v", err)
	}

	err = deleteVolumeGroupInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, volumeGroupId)
	if err != nil {
		log.Printf("[WARN] failed to deleteVolumeGroupInRegion with error %v", err)
	}
}
