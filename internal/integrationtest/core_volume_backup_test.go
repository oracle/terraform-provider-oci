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
	"github.com/oracle/oci-go-sdk/v56/common"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	VolumeBackupRequiredOnlyResource = VolumeBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", acctest.Required, acctest.Create, volumeBackupRepresentation)

	volumeBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"volume_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_volume.test_volume.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: volumeBackupDataSourceFilterRepresentation}}
	volumeBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_volume_backup.test_volume_backup.id}`}},
	}

	volumeBackupFromSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"source_volume_backup_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_volume_backup.test_volume_backup_copy.source_volume_backup_id}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: volumeBackupFromSourceDataSourceFilterRepresentation}}
	volumeBackupFromSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_volume_backup.test_volume_backup_copy.id}`}},
	}

	volumeBackupRepresentation = map[string]interface{}{
		"volume_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_volume.test_volume.id}`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"type":          acctest.Representation{RepType: acctest.Optional, Create: `FULL`},
	}
	volumeBackupWithSourceDetailsRepresentation = map[string]interface{}{
		"source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: volumeBackupSourceDetailsRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}

	volumeBackupId, volumeId                string
	volumeBackupSourceDetailsRepresentation = map[string]interface{}{}

	VolumeBackupResourceDependencies = VolumeResourceConfig
)

// issue-routing-tag: core/blockStorage
func TestCoreVolumeBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_volume_backup.test_volume_backup"
	datasourceName := "data.oci_core_volume_backups.test_volume_backups"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+VolumeBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(volumeBackupRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "core", "volumeBackup", t)

	acctest.ResourceTest(t, testAccCheckCoreVolumeBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", acctest.Required, acctest.Create, volumeBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VolumeBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(volumeBackupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "FULL"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

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
			Config: config + compartmentIdVariableStr + VolumeBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(volumeBackupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "FULL"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_volume_backups", "test_volume_backups", acctest.Optional, acctest.Update, volumeBackupDataSourceRepresentation) +
				compartmentIdVariableStr + VolumeBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_backup", "test_volume_backup", acctest.Optional, acctest.Update, volumeBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_id"),

				resource.TestCheckResourceAttr(datasourceName, "volume_backups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "volume_backups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_backups.0.kms_key_id"),
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
	})
}

func testAccCheckCoreVolumeBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlockstorageClient()

	if volumeBackupId != "" || volumeId != "" {
		deleteSourceVolumeBackupToCopy()
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_backup" {
			noResourceFound = false
			request := oci_core.GetVolumeBackupRequest{}

			tmp := rs.Primary.ID
			request.VolumeBackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreVolumeBackup") {
		resource.AddTestSweepers("CoreVolumeBackup", &resource.Sweeper{
			Name:         "CoreVolumeBackup",
			Dependencies: acctest.DependencyGraph["volumeBackup"],
			F:            sweepCoreVolumeBackupResource,
		})
	}
}

func sweepCoreVolumeBackupResource(compartment string) error {
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()
	volumeBackupIds, err := getVolumeBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeBackupId := range volumeBackupIds {
		if ok := acctest.SweeperDefaultResourceId[volumeBackupId]; !ok {
			deleteVolumeBackupRequest := oci_core.DeleteVolumeBackupRequest{}

			deleteVolumeBackupRequest.VolumeBackupId = &volumeBackupId

			deleteVolumeBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteVolumeBackup(context.Background(), deleteVolumeBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting VolumeBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeBackupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &volumeBackupId, volumeBackupSweepWaitCondition, time.Duration(3*time.Minute),
				volumeBackupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVolumeBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VolumeBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VolumeBackupId", id)
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

func volumeBackupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BlockstorageClient().GetVolumeBackup(context.Background(), oci_core.GetVolumeBackupRequest{
		VolumeBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
