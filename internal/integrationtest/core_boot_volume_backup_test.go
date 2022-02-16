// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	BootVolumeBackupRequiredOnlyResource = BootVolumeBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", acctest.Required, acctest.Create, bootVolumeBackupRepresentation)

	BootVolumeBackupResourceConfig = BootVolumeBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", acctest.Optional, acctest.Update, bootVolumeBackupRepresentation)

	bootVolumeBackupSingularDataSourceRepresentation = map[string]interface{}{
		"boot_volume_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_boot_volume_backup.test_boot_volume_backup.id}`},
	}

	bootVolumeBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"boot_volume_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: bootVolumeBackupDataSourceFilterRepresentation}}
	bootVolumeBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_boot_volume_backup.test_boot_volume_backup.id}`}},
	}

	bootVolumeBackupRepresentation = map[string]interface{}{
		"boot_volume_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_boot_volume.test_boot_volume.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `INCREMENTAL`},
	}
	bootVolumeBackupId, bootVolumeId, instanceId string
	BootVolumeBackupResourceDependencies         = BootVolumeOptionalResource
)

// issue-routing-tag: core/blockStorage
func TestCoreBootVolumeBackupResource_basic(t *testing.T) {
	if httpreplay.ShouldRetryImmediately() {
		t.Skip("TestCoreBootVolumeBackupResource_basic test is flaky in httpreplay mode, skip this test for checkin test.")
	}

	httpreplay.SetScenario("TestCoreBootVolumeBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_boot_volume_backup.test_boot_volume_backup"
	datasourceName := "data.oci_core_boot_volume_backups.test_boot_volume_backups"
	singularDatasourceName := "data.oci_core_boot_volume_backup.test_boot_volume_backup"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+BootVolumeBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(bootVolumeBackupRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "core", "bootVolumeBackup", t)

	acctest.ResourceTest(t, testAccCheckCoreBootVolumeBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BootVolumeBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", acctest.Required, acctest.Create, bootVolumeBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "boot_volume_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BootVolumeBackupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BootVolumeBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(bootVolumeBackupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "boot_volume_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentIdU, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BootVolumeBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(bootVolumeBackupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "boot_volume_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_boot_volume_backups", "test_boot_volume_backups", acctest.Optional, acctest.Update, bootVolumeBackupDataSourceRepresentation) +
				compartmentIdVariableStr + BootVolumeBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", acctest.Optional, acctest.Update, bootVolumeBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.boot_volume_id"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.source_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.time_request_received"),
				resource.TestCheckResourceAttr(datasourceName, "boot_volume_backups.0.type", "INCREMENTAL"),
				resource.TestCheckResourceAttrSet(datasourceName, "boot_volume_backups.0.unique_size_in_gbs"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_boot_volume_backup", "test_boot_volume_backup", acctest.Required, acctest.Create, bootVolumeBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BootVolumeBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "boot_volume_backup_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_request_received"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "INCREMENTAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "unique_size_in_gbs"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + BootVolumeBackupResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreBootVolumeBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlockstorageClient()

	if bootVolumeBackupId != "" || bootVolumeId != "" {
		deleteSourceBootVolumeBackupToCopy()
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_boot_volume_backup" {
			noResourceFound = false
			request := oci_core.GetBootVolumeBackupRequest{}

			tmp := rs.Primary.ID
			request.BootVolumeBackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetBootVolumeBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.BootVolumeBackupLifecycleStateTerminated): true,
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreBootVolumeBackup") {
		resource.AddTestSweepers("CoreBootVolumeBackup", &resource.Sweeper{
			Name:         "CoreBootVolumeBackup",
			Dependencies: acctest.DependencyGraph["bootVolumeBackup"],
			F:            sweepCoreBootVolumeBackupResource,
		})
	}
}

func sweepCoreBootVolumeBackupResource(compartment string) error {
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()
	bootVolumeBackupIds, err := getBootVolumeBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, bootVolumeBackupId := range bootVolumeBackupIds {
		if ok := acctest.SweeperDefaultResourceId[bootVolumeBackupId]; !ok {
			deleteBootVolumeBackupRequest := oci_core.DeleteBootVolumeBackupRequest{}

			deleteBootVolumeBackupRequest.BootVolumeBackupId = &bootVolumeBackupId

			deleteBootVolumeBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteBootVolumeBackup(context.Background(), deleteBootVolumeBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting BootVolumeBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", bootVolumeBackupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bootVolumeBackupId, bootVolumeBackupSweepWaitCondition, time.Duration(3*time.Minute),
				bootVolumeBackupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getBootVolumeBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BootVolumeBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()

	listBootVolumeBackupsRequest := oci_core.ListBootVolumeBackupsRequest{}
	listBootVolumeBackupsRequest.CompartmentId = &compartmentId
	listBootVolumeBackupsRequest.LifecycleState = oci_core.BootVolumeBackupLifecycleStateAvailable
	listBootVolumeBackupsResponse, err := blockstorageClient.ListBootVolumeBackups(context.Background(), listBootVolumeBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BootVolumeBackup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, bootVolumeBackup := range listBootVolumeBackupsResponse.Items {
		id := *bootVolumeBackup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BootVolumeBackupId", id)
	}
	return resourceIds, nil
}

func bootVolumeBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bootVolumeBackupResponse, ok := response.Response.(oci_core.GetBootVolumeBackupResponse); ok {
		return bootVolumeBackupResponse.LifecycleState != oci_core.BootVolumeBackupLifecycleStateTerminated
	}
	return false
}

func bootVolumeBackupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BlockstorageClient().GetBootVolumeBackup(context.Background(), oci_core.GetBootVolumeBackupRequest{
		BootVolumeBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
