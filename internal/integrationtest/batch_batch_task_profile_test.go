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
	BatchBatchTaskProfileRequiredOnlyResource = BatchBatchTaskProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_profile", "test_batch_task_profile", acctest.Required, acctest.Create, BatchBatchTaskProfileRepresentation)

	BatchBatchTaskProfileResourceConfig = BatchBatchTaskProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_profile", "test_batch_task_profile", acctest.Optional, acctest.Update, BatchBatchTaskProfileRepresentation)

	BatchBatchTaskProfileSingularDataSourceRepresentation = map[string]interface{}{
		"batch_task_profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_batch_batch_task_profile.test_batch_task_profile.id}`},
	}

	BatchBatchTaskProfileDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_batch_batch_task_profile.test_batch_task_profile.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchTaskProfileDataSourceFilterRepresentation}}
	BatchBatchTaskProfileDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_batch_batch_task_profile.test_batch_task_profile.id}`}},
	}

	BatchBatchTaskProfileRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"min_memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"min_ocpus":         acctest.Representation{RepType: acctest.Required, Create: `10`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	BatchBatchTaskProfileResourceDependencies = ""
)

// issue-routing-tag: batch/default
func TestBatchBatchTaskProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBatchBatchTaskProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_batch_batch_task_profile.test_batch_task_profile"
	datasourceName := "data.oci_batch_batch_task_profiles.test_batch_task_profiles"
	singularDatasourceName := "data.oci_batch_batch_task_profile.test_batch_task_profile"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BatchBatchTaskProfileResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_profile", "test_batch_task_profile", acctest.Optional, acctest.Create, BatchBatchTaskProfileRepresentation), "batch", "batchTaskProfile", t)

	acctest.ResourceTest(t, testAccCheckBatchBatchTaskProfileDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BatchBatchTaskProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_profile", "test_batch_task_profile", acctest.Required, acctest.Create, BatchBatchTaskProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "min_memory_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "min_ocpus", "10"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BatchBatchTaskProfileResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BatchBatchTaskProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_profile", "test_batch_task_profile", acctest.Optional, acctest.Create, BatchBatchTaskProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "min_memory_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "min_ocpus", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BatchBatchTaskProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_profile", "test_batch_task_profile", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(BatchBatchTaskProfileRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "min_memory_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "min_ocpus", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + BatchBatchTaskProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_profile", "test_batch_task_profile", acctest.Optional, acctest.Update, BatchBatchTaskProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "min_memory_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "min_ocpus", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_task_profiles", "test_batch_task_profiles", acctest.Optional, acctest.Update, BatchBatchTaskProfileDataSourceRepresentation) +
				compartmentIdVariableStr + BatchBatchTaskProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_profile", "test_batch_task_profile", acctest.Optional, acctest.Update, BatchBatchTaskProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "batch_task_profile_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "batch_task_profile_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_task_profile", "test_batch_task_profile", acctest.Required, acctest.Create, BatchBatchTaskProfileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BatchBatchTaskProfileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "batch_task_profile_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_memory_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_ocpus", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + BatchBatchTaskProfileRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckBatchBatchTaskProfileDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BatchComputingClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_batch_batch_task_profile" {
			noResourceFound = false
			request := oci_batch.GetBatchTaskProfileRequest{}

			tmp := rs.Primary.ID
			request.BatchTaskProfileId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "batch")

			response, err := client.GetBatchTaskProfile(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_batch.BatchTaskProfileLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BatchBatchTaskProfile") {
		resource.AddTestSweepers("BatchBatchTaskProfile", &resource.Sweeper{
			Name:         "BatchBatchTaskProfile",
			Dependencies: acctest.DependencyGraph["batchTaskProfile"],
			F:            sweepBatchBatchTaskProfileResource,
		})
	}
}

func sweepBatchBatchTaskProfileResource(compartment string) error {
	batchComputingClient := acctest.GetTestClients(&schema.ResourceData{}).BatchComputingClient()
	batchTaskProfileIds, err := getBatchBatchTaskProfileIds(compartment)
	if err != nil {
		return err
	}
	for _, batchTaskProfileId := range batchTaskProfileIds {
		if ok := acctest.SweeperDefaultResourceId[batchTaskProfileId]; !ok {
			deleteBatchTaskProfileRequest := oci_batch.DeleteBatchTaskProfileRequest{}

			deleteBatchTaskProfileRequest.BatchTaskProfileId = &batchTaskProfileId

			deleteBatchTaskProfileRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "batch")
			_, error := batchComputingClient.DeleteBatchTaskProfile(context.Background(), deleteBatchTaskProfileRequest)
			if error != nil {
				fmt.Printf("Error deleting BatchTaskProfile %s %s, It is possible that the resource is already deleted. Please verify manually \n", batchTaskProfileId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &batchTaskProfileId, BatchBatchTaskProfileSweepWaitCondition, time.Duration(3*time.Minute),
				BatchBatchTaskProfileSweepResponseFetchOperation, "batch", true)
		}
	}
	return nil
}

func getBatchBatchTaskProfileIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BatchTaskProfileId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	batchComputingClient := acctest.GetTestClients(&schema.ResourceData{}).BatchComputingClient()

	listBatchTaskProfilesRequest := oci_batch.ListBatchTaskProfilesRequest{}
	listBatchTaskProfilesRequest.CompartmentId = &compartmentId
	listBatchTaskProfilesRequest.LifecycleState = oci_batch.BatchTaskProfileLifecycleStateActive
	listBatchTaskProfilesResponse, err := batchComputingClient.ListBatchTaskProfiles(context.Background(), listBatchTaskProfilesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BatchTaskProfile list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, batchTaskProfile := range listBatchTaskProfilesResponse.Items {
		id := *batchTaskProfile.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BatchTaskProfileId", id)
	}
	return resourceIds, nil
}

func BatchBatchTaskProfileSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if batchTaskProfileResponse, ok := response.Response.(oci_batch.GetBatchTaskProfileResponse); ok {
		return batchTaskProfileResponse.LifecycleState != oci_batch.BatchTaskProfileLifecycleStateDeleted
	}
	return false
}

func BatchBatchTaskProfileSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BatchComputingClient().GetBatchTaskProfile(context.Background(), oci_batch.GetBatchTaskProfileRequest{
		BatchTaskProfileId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
