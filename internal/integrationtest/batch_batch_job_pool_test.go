// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BatchBatchJobPoolRequiredOnlyResource = BatchBatchJobPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_job_pool", "test_batch_job_pool", acctest.Required, acctest.Create, BatchBatchJobPoolRepresentation)

	BatchBatchJobPoolResourceConfig = BatchBatchJobPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_job_pool", "test_batch_job_pool", acctest.Optional, acctest.Update, BatchBatchJobPoolRepresentation)

	BatchBatchJobPoolSingularDataSourceRepresentation = map[string]interface{}{
		"batch_job_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_batch_batch_job_pool.test_batch_job_pool.id}`},
	}

	BatchBatchJobPoolDataSourceRepresentation = map[string]interface{}{
		"batch_context_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.existing_batch_context_id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_batch_batch_job_pool.test_batch_job_pool.id}`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchJobPoolDataSourceFilterRepresentation}}
	BatchBatchJobPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_batch_batch_job_pool.test_batch_job_pool.id}`}},
	}

	BatchBatchJobPoolRepresentation = map[string]interface{}{
		"batch_context_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_batch_context_id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}

	// No Batch Context creation here; we use an existing context passed in via variable.
	BatchBatchJobPoolResourceDependencies = ""
)

// issue-routing-tag: batch/default
func TestBatchBatchJobPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBatchBatchJobPoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	existingBatchContextId := utils.GetEnvSettingWithBlankDefault("existing_batch_context_id")
	if existingBatchContextId == "" {
		t.Skip("existing_batch_context_id must be set to an existing Batch Context OCID")
	}
	existingBatchContextVariableStr := fmt.Sprintf("variable \"existing_batch_context_id\" { default = \"%s\" }\n", existingBatchContextId)

	resourceName := "oci_batch_batch_job_pool.test_batch_job_pool"
	datasourceName := "data.oci_batch_batch_job_pools.test_batch_job_pools"
	singularDatasourceName := "data.oci_batch_batch_job_pool.test_batch_job_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+existingBatchContextVariableStr+BatchBatchJobPoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_job_pool", "test_batch_job_pool", acctest.Optional, acctest.Create, BatchBatchJobPoolRepresentation), "batch", "batchJobPool", t)

	acctest.ResourceTest(t, testAccCheckBatchBatchJobPoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + existingBatchContextVariableStr + BatchBatchJobPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_job_pool", "test_batch_job_pool", acctest.Required, acctest.Create, BatchBatchJobPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "batch_context_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + existingBatchContextVariableStr + BatchBatchJobPoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + existingBatchContextVariableStr + BatchBatchJobPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_job_pool", "test_batch_job_pool", acctest.Optional, acctest.Create, BatchBatchJobPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "batch_context_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + existingBatchContextVariableStr + BatchBatchJobPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_job_pool", "test_batch_job_pool", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(BatchBatchJobPoolRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "batch_context_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + existingBatchContextVariableStr + BatchBatchJobPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_job_pool", "test_batch_job_pool", acctest.Optional, acctest.Update, BatchBatchJobPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "batch_context_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_job_pools", "test_batch_job_pools", acctest.Optional, acctest.Update, BatchBatchJobPoolDataSourceRepresentation) +
				compartmentIdVariableStr + existingBatchContextVariableStr + BatchBatchJobPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_job_pool", "test_batch_job_pool", acctest.Optional, acctest.Update, BatchBatchJobPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "batch_context_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "batch_job_pool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "batch_job_pool_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_job_pool", "test_batch_job_pool", acctest.Required, acctest.Create, BatchBatchJobPoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + existingBatchContextVariableStr + BatchBatchJobPoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "batch_job_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + BatchBatchJobPoolRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckBatchBatchJobPoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BatchComputingClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_batch_batch_job_pool" {
			noResourceFound = false
			request := oci_batch.GetBatchJobPoolRequest{}

			tmp := rs.Primary.ID
			request.BatchJobPoolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "batch")

			response, err := client.GetBatchJobPool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_batch.BatchJobPoolLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BatchBatchJobPool") {
		resource.AddTestSweepers("BatchBatchJobPool", &resource.Sweeper{
			Name:         "BatchBatchJobPool",
			Dependencies: acctest.DependencyGraph["batchJobPool"],
			F:            sweepBatchBatchJobPoolResource,
		})
	}
}

func sweepBatchBatchJobPoolResource(compartment string) error {
	batchComputingClient := acctest.GetTestClients(&schema.ResourceData{}).BatchComputingClient()
	batchJobPoolIds, err := getBatchBatchJobPoolIds(compartment)
	if err != nil {
		return err
	}
	for _, batchJobPoolId := range batchJobPoolIds {
		if ok := acctest.SweeperDefaultResourceId[batchJobPoolId]; !ok {
			deleteBatchJobPoolRequest := oci_batch.DeleteBatchJobPoolRequest{}

			deleteBatchJobPoolRequest.BatchJobPoolId = &batchJobPoolId

			deleteBatchJobPoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "batch")
			_, error := batchComputingClient.DeleteBatchJobPool(context.Background(), deleteBatchJobPoolRequest)
			if error != nil {
				fmt.Printf("Error deleting BatchJobPool %s %s, It is possible that the resource is already deleted. Please verify manually \n", batchJobPoolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &batchJobPoolId, BatchBatchJobPoolSweepWaitCondition, time.Duration(3*time.Minute),
				BatchBatchJobPoolSweepResponseFetchOperation, "batch", true)
		}
	}
	return nil
}

func getBatchBatchJobPoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BatchJobPoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	batchComputingClient := acctest.GetTestClients(&schema.ResourceData{}).BatchComputingClient()

	listBatchJobPoolsRequest := oci_batch.ListBatchJobPoolsRequest{}
	listBatchJobPoolsRequest.CompartmentId = &compartmentId
	listBatchJobPoolsRequest.LifecycleState = oci_batch.BatchJobPoolLifecycleStateNeedsAttention
	listBatchJobPoolsResponse, err := batchComputingClient.ListBatchJobPools(context.Background(), listBatchJobPoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BatchJobPool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, batchJobPool := range listBatchJobPoolsResponse.Items {
		id := *batchJobPool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BatchJobPoolId", id)
	}
	return resourceIds, nil
}

func BatchBatchJobPoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if batchJobPoolResponse, ok := response.Response.(oci_batch.GetBatchJobPoolResponse); ok {
		return batchJobPoolResponse.LifecycleState != oci_batch.BatchJobPoolLifecycleStateDeleted
	}
	return false
}

func BatchBatchJobPoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BatchComputingClient().GetBatchJobPool(context.Background(), oci_batch.GetBatchJobPoolRequest{
		BatchJobPoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
