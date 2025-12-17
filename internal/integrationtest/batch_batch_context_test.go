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
	BatchBatchContextDisplayName       = "batchcontext" + time.Now().UTC().Format("20060102150405")
	BatchBatchContextDisplayNameUpdate = BatchBatchContextDisplayName + "2"

	BatchBatchContextRequiredOnlyResource = BatchBatchContextResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Required, acctest.Create, BatchBatchContextRepresentation)

	BatchBatchContextResourceConfig = BatchBatchContextResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Optional, acctest.Update, BatchBatchContextRepresentation)

	BatchBatchContextSingularDataSourceRepresentation = map[string]interface{}{
		"batch_context_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_batch_batch_context.test_batch_context.id}`},
	}

	BatchBatchContextDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: BatchBatchContextDisplayName, Update: BatchBatchContextDisplayNameUpdate},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_batch_batch_context.test_batch_context.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextDataSourceFilterRepresentation}}
	BatchBatchContextDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_batch_batch_context.test_batch_context.id}`}},
	}

	BatchBatchContextRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fleets":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextFleetsRepresentation},
		"network":        acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextNetworkRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: BatchBatchContextDisplayName, Update: BatchBatchContextDisplayNameUpdate},
	}
	BatchBatchContextFleetsRepresentation = map[string]interface{}{
		"max_concurrent_tasks": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"name":                 acctest.Representation{RepType: acctest.Required, Create: `name`},
		"shape":                acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextFleetsShapeRepresentation},
		"type":                 acctest.Representation{RepType: acctest.Required, Create: `SERVICE_MANAGED_FLEET`},
	}
	BatchBatchContextNetworkRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	BatchBatchContextFleetsShapeRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `16`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1`},
		"shape_name":    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`},
	}

	BatchBatchContextResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: batch/default
func TestBatchBatchContextResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBatchBatchContextResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_batch_batch_context.test_batch_context"
	datasourceName := "data.oci_batch_batch_contexts.test_batch_contexts"
	singularDatasourceName := "data.oci_batch_batch_context.test_batch_context"
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BatchBatchContextResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Optional, acctest.Create, BatchBatchContextRepresentation), "batch", "batchContext", t)

	acctest.ResourceTest(t, testAccCheckBatchBatchContextDestroy, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + BatchBatchContextResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Optional, acctest.Create, BatchBatchContextRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", BatchBatchContextDisplayName),
				resource.TestCheckResourceAttr(resourceName, "fleets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.max_concurrent_tasks", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.memory_in_gbs", "16"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.shape_name", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.type", "SERVICE_MANAGED_FLEET"),
				resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "network.0.vnics.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},
		// verify datasource (same Batch Context, no extra creates)
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_contexts", "test_batch_contexts", acctest.Optional, acctest.Update, BatchBatchContextDataSourceRepresentation) +
				compartmentIdVariableStr + BatchBatchContextResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Optional, acctest.Update, BatchBatchContextRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", BatchBatchContextDisplayNameUpdate),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "batch_context_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "batch_context_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Required, acctest.Create, BatchBatchContextSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BatchBatchContextResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "batch_context_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", BatchBatchContextDisplayNameUpdate),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.max_concurrent_tasks", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.0.memory_in_gbs", "16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.0.ocpus", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.type", "SERVICE_MANAGED_FLEET"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network.0.vnics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}

func testAccCheckBatchBatchContextDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BatchComputingClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_batch_batch_context" {
			noResourceFound = false
			request := oci_batch.GetBatchContextRequest{}

			tmp := rs.Primary.ID
			request.BatchContextId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "batch")

			response, err := client.GetBatchContext(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_batch.BatchContextLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BatchBatchContext") {
		resource.AddTestSweepers("BatchBatchContext", &resource.Sweeper{
			Name:         "BatchBatchContext",
			Dependencies: acctest.DependencyGraph["batchContext"],
			F:            sweepBatchBatchContextResource,
		})
	}
}

func sweepBatchBatchContextResource(compartment string) error {
	batchComputingClient := acctest.GetTestClients(&schema.ResourceData{}).BatchComputingClient()
	batchContextIds, err := getBatchBatchContextIds(compartment)
	if err != nil {
		return err
	}
	for _, batchContextId := range batchContextIds {
		if ok := acctest.SweeperDefaultResourceId[batchContextId]; !ok {
			deleteBatchContextRequest := oci_batch.DeleteBatchContextRequest{}

			deleteBatchContextRequest.BatchContextId = &batchContextId

			deleteBatchContextRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "batch")
			_, error := batchComputingClient.DeleteBatchContext(context.Background(), deleteBatchContextRequest)
			if error != nil {
				fmt.Printf("Error deleting BatchContext %s %s, It is possible that the resource is already deleted. Please verify manually \n", batchContextId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &batchContextId, BatchBatchContextSweepWaitCondition, time.Duration(3*time.Minute),
				BatchBatchContextSweepResponseFetchOperation, "batch", true)
		}
	}
	return nil
}

func getBatchBatchContextIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BatchContextId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	batchComputingClient := acctest.GetTestClients(&schema.ResourceData{}).BatchComputingClient()

	listBatchContextsRequest := oci_batch.ListBatchContextsRequest{}
	listBatchContextsRequest.CompartmentId = &compartmentId
	listBatchContextsRequest.LifecycleState = oci_batch.BatchContextLifecycleStateNeedsAttention
	listBatchContextsResponse, err := batchComputingClient.ListBatchContexts(context.Background(), listBatchContextsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BatchContext list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, batchContext := range listBatchContextsResponse.Items {
		id := *batchContext.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BatchContextId", id)
	}
	return resourceIds, nil
}

func BatchBatchContextSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if batchContextResponse, ok := response.Response.(oci_batch.GetBatchContextResponse); ok {
		return batchContextResponse.LifecycleState != oci_batch.BatchContextLifecycleStateDeleted
	}
	return false
}

func BatchBatchContextSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BatchComputingClient().GetBatchContext(context.Background(), oci_batch.GetBatchContextRequest{
		BatchContextId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
