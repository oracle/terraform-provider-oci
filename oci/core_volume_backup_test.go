// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"

	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	VolumeBackupRequiredOnlyResource = VolumeBackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Required, Create, volumeBackupRepresentation)

	volumeBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"volume_id":      Representation{repType: Optional, create: `${oci_core_volume.test_volume.id}`},
		"filter":         RepresentationGroup{Required, volumeBackupDataSourceFilterRepresentation}}
	volumeBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume_backup.test_volume_backup.id}`}},
	}

	volumeBackupFromSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":            Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"source_volume_backup_id": Representation{repType: Optional, create: `${oci_core_volume_backup.test_volume_backup_copy.source_volume_backup_id}`},
		"state":                   Representation{repType: Optional, create: `AVAILABLE`},
		"filter":                  RepresentationGroup{Required, volumeBackupFromSourceDataSourceFilterRepresentation}}
	volumeBackupFromSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_volume_backup.test_volume_backup_copy.id}`}},
	}

	volumeBackupRepresentation = map[string]interface{}{
		"volume_id":     Representation{repType: Required, create: `${oci_core_volume.test_volume.id}`},
		"defined_tags":  Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags": Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"type":          Representation{repType: Optional, create: `FULL`},
	}
	volumeBackupWithSourceDetailsRepresentation = map[string]interface{}{
		"source_details": RepresentationGroup{Required, volumeBackupSourceDetailsRepresentation},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
	}

	volumeBackupId, volumeId                string
	volumeBackupSourceDetailsRepresentation = map[string]interface{}{}

	VolumeBackupResourceDependencies = VolumeResourceConfig
)

func TestCoreVolumeBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeBackupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_volume_backup.test_volume_backup"
	datasourceName := "data.oci_core_volume_backups.test_volume_backups"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Required, Create, volumeBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Optional, Create, volumeBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "type", "FULL"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Optional, Update, volumeBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "type", "FULL"),
					resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_core_volume_backups", "test_volume_backups", Optional, Update, volumeBackupDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", Optional, Update, volumeBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_id"),

					resource.TestCheckResourceAttr(datasourceName, "volume_backups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.size_in_mbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.source_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.time_request_received"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.type", "FULL"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.unique_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.unique_size_in_mbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.volume_id"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func TestCoreVolumeBackupResource_copy(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameCopy := "oci_core_volume_backup.test_volume_backup_copy"
	datasourceName := "data.oci_core_volume_backups.test_volume_backups"

	if getEnvSettingWithBlankDefault("source_region") == "" {
		t.Skip("Skipping TestCoreVolumeBackupResource_copy test because there is no source region specified")
	}

	err := createSourceVolumeBackupToCopy()
	if err != nil {
		t.Fatalf("Unable to create source Volume and VolumeBackup to copy. Error: %v", err)
	}

	volumeBackupSourceDetailsRepresentation = map[string]interface{}{
		"volume_backup_id": Representation{repType: Required, create: volumeBackupId},
		"region":           Representation{repType: Required, create: getEnvSettingWithBlankDefault("source_region")},
	}
	volumeBackupWithSourceDetailsRepresentation = getUpdatedRepresentationCopy("source_details", RepresentationGroup{Required, volumeBackupSourceDetailsRepresentation}, volumeBackupWithSourceDetailsRepresentation)

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreVolumeBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config +
					compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup_copy", Required, Create, volumeBackupWithSourceDetailsRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceNameCopy, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies,
			},
			// verify create from the backup with optionals
			{
				Config: config +
					compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup_copy", Optional, Create, volumeBackupWithSourceDetailsRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
					resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_id"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "source_volume_backup_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceNameCopy, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config +
					compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup_copy", Optional, Update, volumeBackupWithSourceDetailsRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "compartment_id"),
					resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "type"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "volume_id"),
					resource.TestCheckResourceAttrSet(resourceNameCopy, "source_volume_backup_id"),

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
					generateDataSourceFromRepresentationMap("oci_core_volume_backups", "test_volume_backups", Optional, Update, volumeBackupFromSourceDataSourceRepresentation) +
					compartmentIdVariableStr + VolumeBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup_copy", Optional, Update, volumeBackupWithSourceDetailsRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "source_volume_backup_id"),

					resource.TestCheckResourceAttr(datasourceName, "volume_backups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.type"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.volume_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.source_volume_backup_id"),
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

func createSourceVolumeBackupToCopy() error {
	sourceRegion := getEnvSettingWithBlankDefault("source_region")

	var err error
	volumeId, err = createVolumeInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion)
	if err != nil {
		log.Printf("[WARN] failed to createVolumeInRegion with the error %v", err)
		return err
	}

	volumeBackupId, err = createVolumeBackupInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion, &volumeId)
	if err != nil {
		log.Printf("[WARN] failed to createVolumeBackupInRegion with the error %v", err)
		return err
	}

	return nil
}

func deleteSourceVolumeBackupToCopy() {
	sourceRegion := getEnvSettingWithBlankDefault("source_region")

	var err error
	err = deleteVolumeBackupInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion, volumeBackupId)
	if err != nil {
		log.Printf("[WARN] failed to deleteVolumeBackupInRegion with error %v", err)
	}

	err = deleteVolumeInRegion(GetTestClients(&schema.ResourceData{}), sourceRegion, volumeId)
	if err != nil {
		log.Printf("[WARN] failed to deleteVolumeInRegion with error %v", err)
	}
}

func testAccCheckCoreVolumeBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockstorageClient

	if volumeBackupId != "" || volumeId != "" {
		deleteSourceVolumeBackupToCopy()
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_backup" {
			noResourceFound = false
			request := oci_core.GetVolumeBackupRequest{}

			tmp := rs.Primary.ID
			request.VolumeBackupId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			response, err := client.GetVolumeBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VolumeBackupLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreVolumeBackup") {
		resource.AddTestSweepers("CoreVolumeBackup", &resource.Sweeper{
			Name:         "CoreVolumeBackup",
			Dependencies: DependencyGraph["volumeBackup"],
			F:            sweepCoreVolumeBackupResource,
		})
	}
}

func sweepCoreVolumeBackupResource(compartment string) error {
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient
	volumeBackupIds, err := getVolumeBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeBackupId := range volumeBackupIds {
		if ok := SweeperDefaultResourceId[volumeBackupId]; !ok {
			deleteVolumeBackupRequest := oci_core.DeleteVolumeBackupRequest{}

			deleteVolumeBackupRequest.VolumeBackupId = &volumeBackupId

			deleteVolumeBackupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteVolumeBackup(context.Background(), deleteVolumeBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting VolumeBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeBackupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &volumeBackupId, volumeBackupSweepWaitCondition, time.Duration(3*time.Minute),
				volumeBackupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVolumeBackupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "VolumeBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := GetTestClients(&schema.ResourceData{}).blockstorageClient

	listVolumeBackupsRequest := oci_core.ListVolumeBackupsRequest{}
	listVolumeBackupsRequest.CompartmentId = &compartmentId
	listVolumeBackupsRequest.LifecycleState = oci_core.VolumeBackupLifecycleStateAvailable
	listVolumeBackupsResponse, err := blockstorageClient.ListVolumeBackups(context.Background(), listVolumeBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VolumeBackup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, volumeBackup := range listVolumeBackupsResponse.Items {
		id := *volumeBackup.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "VolumeBackupId", id)
	}
	return resourceIds, nil
}

func volumeBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if volumeBackupResponse, ok := response.Response.(oci_core.GetVolumeBackupResponse); ok {
		return volumeBackupResponse.LifecycleState != oci_core.VolumeBackupLifecycleStateTerminated
	}
	return false
}

func volumeBackupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.blockstorageClient.GetVolumeBackup(context.Background(), oci_core.GetVolumeBackupRequest{
		VolumeBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
