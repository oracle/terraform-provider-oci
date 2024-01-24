// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// This test requires long-lived replication Target resource to do the LIST and get operation on replicationTarget resource

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/common"

	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FileStorageReplicationTargetSingularDataSourceRepresentation = map[string]interface{}{
		"replication_target_id": acctest.Representation{RepType: acctest.Required, Create: `${var.replication_target_id}`},
	}

	FileStorageReplicationTargetDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_static_resource}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `replication-terraform-test`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${var.replication_target_id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	FileStorageReplicationTargetResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: file_storage/default
func TestFileStorageReplicationTargetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageReplicationTargetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	replicationTargetId := utils.GetEnvSettingWithBlankDefault("replication_target_id")
	replicationTargetIdVariableStr := fmt.Sprintf("variable \"replication_target_id\" { default = \"%s\" }\n", replicationTargetId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id_for_static_resource")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id_for_static_resource\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_file_storage_replication_targets.test_replication_targets"
	singularDatasourceName := "data.oci_file_storage_replication_target.test_replication_target"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_replication_targets", "test_replication_targets", acctest.Required, acctest.Create, FileStorageReplicationTargetDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageReplicationTargetResourceConfig + replicationTargetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttrSet(datasourceName, "replication_targets.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_targets.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_targets.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "replication_targets.0.display_name", "replication-terraform-test"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_targets.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_targets.0.recovery_point_time"),
				resource.TestCheckResourceAttr(datasourceName, "replication_targets.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_targets.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_replication_target", "test_replication_target", acctest.Required, acctest.Create, FileStorageReplicationTargetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageReplicationTargetResourceConfig + replicationTargetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replication_target_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "delta_progress"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "delta_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_snapshot_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recovery_point_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replication_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("FileStorageReplicationTarget") {
		resource.AddTestSweepers("FileStorageReplicationTarget", &resource.Sweeper{
			Name:         "FileStorageReplicationTarget",
			Dependencies: acctest.DependencyGraph["replicationTarget"],
			F:            sweepFileStorageReplicationTargetResource,
		})
	}
}

func sweepFileStorageReplicationTargetResource(compartment string) error {
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()
	replicationTargetIds, err := getFileStorageReplicationTargetIds(compartment)
	if err != nil {
		return err
	}
	for _, replicationTargetId := range replicationTargetIds {
		if ok := acctest.SweeperDefaultResourceId[replicationTargetId]; !ok {
			deleteReplicationTargetRequest := oci_file_storage.DeleteReplicationTargetRequest{}

			deleteReplicationTargetRequest.ReplicationTargetId = &replicationTargetId

			deleteReplicationTargetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")
			_, error := fileStorageClient.DeleteReplicationTarget(context.Background(), deleteReplicationTargetRequest)
			if error != nil {
				fmt.Printf("Error deleting ReplicationTarget %s %s, It is possible that the resource is already deleted. Please verify manually \n", replicationTargetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &replicationTargetId, FileStorageReplicationTargetSweepWaitCondition, time.Duration(3*time.Minute),
				FileStorageReplicationTargetSweepResponseFetchOperation, "file_storage", true)
		}
	}
	return nil
}

func getFileStorageReplicationTargetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ReplicationTargetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()

	listReplicationTargetsRequest := oci_file_storage.ListReplicationTargetsRequest{}
	listReplicationTargetsRequest.CompartmentId = &compartmentId

	availabilityDomains, err := acctest.GetAvalabilityDomains(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting availabilityDomains required for ReplicationTarget list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, availabilityDomainName := range availabilityDomains {
		listReplicationTargetsRequest.AvailabilityDomain = &availabilityDomainName

		listReplicationTargetsRequest.LifecycleState = oci_file_storage.ListReplicationTargetsLifecycleStateActive
		listReplicationTargetsResponse, err := fileStorageClient.ListReplicationTargets(context.Background(), listReplicationTargetsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting ReplicationTarget list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, replicationTarget := range listReplicationTargetsResponse.Items {
			id := *replicationTarget.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ReplicationTargetId", id)
		}

	}
	return resourceIds, nil
}

func FileStorageReplicationTargetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if replicationTargetResponse, ok := response.Response.(oci_file_storage.GetReplicationTargetResponse); ok {
		return replicationTargetResponse.LifecycleState != oci_file_storage.ReplicationTargetLifecycleStateDeleted
	}
	return false
}

func FileStorageReplicationTargetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FileStorageClient().GetReplicationTarget(context.Background(), oci_file_storage.GetReplicationTargetRequest{
		ReplicationTargetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
