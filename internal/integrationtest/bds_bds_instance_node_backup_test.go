// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsBdsInstanceNodeBackupSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"node_backup_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance_node_backup.test_bds_instance_node_backup.id}`},
	}

	BdsBdsInstanceNodeBackupDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}
	BdsBdsInstanceNodeBackupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)

	BdsBdsInstanceNodeBackupRepresentation = map[string]interface{}{
		"bds_instance_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"level_type_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsInstanceNodeBackupLevelTypeDetailsRepresentation},
		"backup_type":        acctest.Representation{RepType: acctest.Optional, Create: `FULL`},
	}

	BdsBdsInstanceNodeBackupLevelTypeDetailsRepresentation = map[string]interface{}{
		"level_type": acctest.Representation{RepType: acctest.Required, Create: `NODE_TYPE_LEVEL`},
		"node_type":  acctest.Representation{RepType: acctest.Optional, Create: `MASTER`},
	}
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceNodeBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceNodeBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	//bdsInstanceId := utils.GetEnvSettingWithBlankDefault("bds_instance_ocid")
	//bdsInstanceIdVariableStr := fmt.Sprintf("variable \"bds_instance_id\" { default = \"%s\" }\n", bdsInstanceId)

	resourceName := "oci_bds_bds_instance_node_backup.test_bds_instance_node_backup"
	datasourceName := "data.oci_bds_bds_instance_node_backups.test_bds_instance_node_backups"
	singularDatasourceName := "data.oci_bds_bds_instance_node_backup.test_bds_instance_node_backup"

	var resId string
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceNodeBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_backup", "test_bds_instance_node_backup", acctest.Optional, acctest.Create, BdsBdsInstanceNodeBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_type", "FULL"),
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "level_type_details.0.level_type", "NODE_TYPE_LEVEL"),
				resource.TestCheckResourceAttr(resourceName, "level_type_details.0.node_type", "MASTER"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
		// verify datasource
		{
			Config: config + acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_backup", "test_bds_instance_node_backup", acctest.Optional, acctest.Create, BdsBdsInstanceNodeBackupRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_node_backups", "test_bds_instance_node_backups", acctest.Required, acctest.Create, BdsBdsInstanceNodeBackupDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceNodeBackupResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backups.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backups.0.backup_trigger_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backups.0.backup_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backups.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backups.0.node_host_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backups.0.node_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backups.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config + acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_backup", "test_bds_instance_node_backup", acctest.Optional, acctest.Create, BdsBdsInstanceNodeBackupRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_node_backup", "test_bds_instance_node_backup", acctest.Required, acctest.Create, BdsBdsInstanceNodeBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceNodeBackupResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_trigger_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("BdsBdsInstanceNodeBackup") {
		resource.AddTestSweepers("BdsBdsInstanceNodeBackup", &resource.Sweeper{
			Name:         "BdsBdsInstanceNodeBackup",
			Dependencies: acctest.DependencyGraph["bdsInstanceNodeBackup"],
			F:            sweepBdsBdsInstanceNodeBackupResource,
		})
	}
}

func sweepBdsBdsInstanceNodeBackupResource(compartment string) error {
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()
	bdsInstanceNodeBackupIds, err := getBdsBdsInstanceNodeBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsInstanceNodeBackupId := range bdsInstanceNodeBackupIds {
		if ok := acctest.SweeperDefaultResourceId[bdsInstanceNodeBackupId]; !ok {
			deleteNodeBackupRequest := oci_bds.DeleteNodeBackupRequest{}

			deleteNodeBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")
			_, error := bdsClient.DeleteNodeBackup(context.Background(), deleteNodeBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsInstanceNodeBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsInstanceNodeBackupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bdsInstanceNodeBackupId, BdsBdsInstanceNodeBackupSweepWaitCondition, time.Duration(3*time.Minute),
				BdsBdsInstanceNodeBackupSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsBdsInstanceNodeBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BdsInstanceNodeBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()

	listNodeBackupsRequest := oci_bds.ListNodeBackupsRequest{}
	//listNodeBackupsRequest.CompartmentId = &compartmentId

	bdsInstanceIds, error := getBdsInstanceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bdsInstanceId required for BdsInstanceNodeBackup resource requests \n")
	}
	for _, bdsInstanceId := range bdsInstanceIds {
		listNodeBackupsRequest.BdsInstanceId = &bdsInstanceId

		listNodeBackupsRequest.LifecycleState = oci_bds.NodeBackupLifecycleStateActive
		listNodeBackupsResponse, err := bdsClient.ListNodeBackups(context.Background(), listNodeBackupsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting BdsInstanceNodeBackup list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, bdsInstanceNodeBackup := range listNodeBackupsResponse.Items {
			id := *bdsInstanceNodeBackup.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BdsInstanceNodeBackupId", id)
		}

	}
	return resourceIds, nil
}

func BdsBdsInstanceNodeBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bdsInstanceNodeBackupResponse, ok := response.Response.(oci_bds.GetNodeBackupResponse); ok {
		return bdsInstanceNodeBackupResponse.LifecycleState != oci_bds.NodeBackupLifecycleStateDeleted
	}
	return false
}

func BdsBdsInstanceNodeBackupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetNodeBackup(context.Background(), oci_bds.GetNodeBackupRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
