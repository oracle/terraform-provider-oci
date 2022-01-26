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
	VolumeGroupBackupRequiredOnlyResource = VolumeGroupBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", acctest.Required, acctest.Create, volumeGroupBackupRepresentation)

	volumeGroupBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"volume_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_volume_group.test_volume_group.id}`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: volumeGroupBackupDataSourceFilterRepresentation}}
	volumeGroupBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_volume_group_backup.test_volume_group_backup.id}`}},
	}

	volumeGroupBackupRepresentation = map[string]interface{}{
		"volume_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_volume_group.test_volume_group.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"type":            acctest.Representation{RepType: acctest.Optional, Create: `INCREMENTAL`},
	}

	VolumeGroupBackupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group", "test_volume_group", acctest.Required, acctest.Create, volumeGroupRepresentation) +
		SourceVolumeListDependency +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/blockStorage
func TestCoreVolumeGroupBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVolumeGroupBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_volume_group_backup.test_volume_group_backup"
	datasourceName := "data.oci_core_volume_group_backups.test_volume_group_backups"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+VolumeGroupBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", acctest.Optional, acctest.Create, volumeGroupBackupRepresentation), "core", "volumeGroupBackup", t)

	acctest.ResourceTest(t, testAccCheckCoreVolumeGroupBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + VolumeGroupBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", acctest.Required, acctest.Create, volumeGroupBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "volume_group_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VolumeGroupBackupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VolumeGroupBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", acctest.Optional, acctest.Create, volumeGroupBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_backup_ids.#"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_group_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VolumeGroupBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(volumeGroupBackupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_backup_ids.#"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_group_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + VolumeGroupBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", acctest.Optional, acctest.Update, volumeGroupBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "INCREMENTAL"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_backup_ids.#"),
				resource.TestCheckResourceAttrSet(resourceName, "volume_group_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_volume_group_backups", "test_volume_group_backups", acctest.Optional, acctest.Update, volumeGroupBackupDataSourceRepresentation) +
				compartmentIdVariableStr + VolumeGroupBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_volume_group_backup", "test_volume_group_backup", acctest.Optional, acctest.Update, volumeGroupBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_id"),

				resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.size_in_mbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.source_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.time_request_received"),
				resource.TestCheckResourceAttr(datasourceName, "volume_group_backups.0.type", "INCREMENTAL"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.unique_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.unique_size_in_mbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.volume_backup_ids.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "volume_group_backups.0.volume_group_id"),
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

func testAccCheckCoreVolumeGroupBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlockstorageClient()

	if volumeGroupBackupId != "" || volumeGroupId != "" {
		deleteSourceVolumeGroupBackupToCopy()
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_volume_group_backup" {
			noResourceFound = false
			request := oci_core.GetVolumeGroupBackupRequest{}

			tmp := rs.Primary.ID
			request.VolumeGroupBackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetVolumeGroupBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VolumeGroupBackupLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreVolumeGroupBackup") {
		resource.AddTestSweepers("CoreVolumeGroupBackup", &resource.Sweeper{
			Name:         "CoreVolumeGroupBackup",
			Dependencies: acctest.DependencyGraph["volumeGroupBackup"],
			F:            sweepCoreVolumeGroupBackupResource,
		})
	}
}

func sweepCoreVolumeGroupBackupResource(compartment string) error {
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()
	volumeGroupBackupIds, err := getVolumeGroupBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, volumeGroupBackupId := range volumeGroupBackupIds {
		if ok := acctest.SweeperDefaultResourceId[volumeGroupBackupId]; !ok {
			deleteVolumeGroupBackupRequest := oci_core.DeleteVolumeGroupBackupRequest{}

			deleteVolumeGroupBackupRequest.VolumeGroupBackupId = &volumeGroupBackupId

			deleteVolumeGroupBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := blockstorageClient.DeleteVolumeGroupBackup(context.Background(), deleteVolumeGroupBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting VolumeGroupBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", volumeGroupBackupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &volumeGroupBackupId, volumeGroupBackupSweepWaitCondition, time.Duration(3*time.Minute),
				volumeGroupBackupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getVolumeGroupBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VolumeGroupBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockstorageClient := acctest.GetTestClients(&schema.ResourceData{}).BlockstorageClient()

	listVolumeGroupBackupsRequest := oci_core.ListVolumeGroupBackupsRequest{}
	listVolumeGroupBackupsRequest.CompartmentId = &compartmentId
	listVolumeGroupBackupsResponse, err := blockstorageClient.ListVolumeGroupBackups(context.Background(), listVolumeGroupBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VolumeGroupBackup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, volumeGroupBackup := range listVolumeGroupBackupsResponse.Items {
		id := *volumeGroupBackup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VolumeGroupBackupId", id)
	}
	return resourceIds, nil
}

func volumeGroupBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if volumeGroupBackupResponse, ok := response.Response.(oci_core.GetVolumeGroupBackupResponse); ok {
		return volumeGroupBackupResponse.LifecycleState != oci_core.VolumeGroupBackupLifecycleStateTerminated
	}
	return false
}

func volumeGroupBackupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BlockstorageClient().GetVolumeGroupBackup(context.Background(), oci_core.GetVolumeGroupBackupRequest{
		VolumeGroupBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
