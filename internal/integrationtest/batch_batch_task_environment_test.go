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
	batchTaskEnvironmentImageURL                  = utils.GetEnvSettingWithDefault("batch_task_environment_image_url", "mad.ocir.io/idokjuh8gabq/arnold:latest")
	BatchBatchTaskEnvironmentRequiredOnlyResource = BatchBatchTaskEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_environment", "test_batch_task_environment", acctest.Required, acctest.Create, BatchBatchTaskEnvironmentRepresentation)

	BatchBatchTaskEnvironmentResourceConfig = BatchBatchTaskEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_environment", "test_batch_task_environment", acctest.Optional, acctest.Update, BatchBatchTaskEnvironmentRepresentation)

	BatchBatchTaskEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"batch_task_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_batch_batch_task_environment.test_batch_task_environment.id}`},
	}

	BatchBatchTaskEnvironmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_batch_batch_task_environment.test_batch_task_environment.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchTaskEnvironmentDataSourceFilterRepresentation}}
	BatchBatchTaskEnvironmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_batch_batch_task_environment.test_batch_task_environment.id}`}},
	}

	BatchBatchTaskEnvironmentRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"image_url":         acctest.Representation{RepType: acctest.Required, Create: batchTaskEnvironmentImageURL},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"security_context":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: BatchBatchTaskEnvironmentSecurityContextRepresentation},
		"volumes":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: BatchBatchTaskEnvironmentVolumesRepresentation},
		"working_directory": acctest.Representation{RepType: acctest.Optional, Create: `workingDirectory`},
	}
	BatchBatchTaskEnvironmentSecurityContextRepresentation = map[string]interface{}{
		"fs_group":     acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"run_as_group": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"run_as_user":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	BatchBatchTaskEnvironmentVolumesRepresentation = map[string]interface{}{
		"local_mount_directory_path": acctest.Representation{RepType: acctest.Required, Create: `localMountDirectoryPath`},
		"mount_target_export_path":   acctest.Representation{RepType: acctest.Required, Create: `mountTargetExportPath`},
		"mount_target_fqdn":          acctest.Representation{RepType: acctest.Required, Create: `mountTargetFqdn`},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `name`},
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `NFS`},
	}

	BatchBatchTaskEnvironmentResourceDependencies = ""
)

// issue-routing-tag: batch/default
func TestBatchBatchTaskEnvironmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBatchBatchTaskEnvironmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_batch_batch_task_environment.test_batch_task_environment"
	datasourceName := "data.oci_batch_batch_task_environments.test_batch_task_environments"
	singularDatasourceName := "data.oci_batch_batch_task_environment.test_batch_task_environment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BatchBatchTaskEnvironmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_environment", "test_batch_task_environment", acctest.Optional, acctest.Create, BatchBatchTaskEnvironmentRepresentation), "batch", "batchTaskEnvironment", t)

	acctest.ResourceTest(t, testAccCheckBatchBatchTaskEnvironmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BatchBatchTaskEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_environment", "test_batch_task_environment", acctest.Required, acctest.Create, BatchBatchTaskEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "image_url", batchTaskEnvironmentImageURL),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BatchBatchTaskEnvironmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BatchBatchTaskEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_environment", "test_batch_task_environment", acctest.Optional, acctest.Create, BatchBatchTaskEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_url", batchTaskEnvironmentImageURL),
				resource.TestCheckResourceAttr(resourceName, "security_context.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_context.0.fs_group", "10"),
				resource.TestCheckResourceAttr(resourceName, "security_context.0.run_as_group", "10"),
				resource.TestCheckResourceAttr(resourceName, "security_context.0.run_as_user", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "volumes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.local_mount_directory_path", "localMountDirectoryPath"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.mount_target_export_path", "mountTargetExportPath"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.mount_target_fqdn", "mountTargetFqdn"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "working_directory", "workingDirectory"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BatchBatchTaskEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_environment", "test_batch_task_environment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(BatchBatchTaskEnvironmentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_url", batchTaskEnvironmentImageURL),
				resource.TestCheckResourceAttr(resourceName, "security_context.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_context.0.fs_group", "10"),
				resource.TestCheckResourceAttr(resourceName, "security_context.0.run_as_group", "10"),
				resource.TestCheckResourceAttr(resourceName, "security_context.0.run_as_user", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "volumes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.local_mount_directory_path", "localMountDirectoryPath"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.mount_target_export_path", "mountTargetExportPath"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.mount_target_fqdn", "mountTargetFqdn"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "working_directory", "workingDirectory"),

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
			Config: config + compartmentIdVariableStr + BatchBatchTaskEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_environment", "test_batch_task_environment", acctest.Optional, acctest.Update, BatchBatchTaskEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "image_url", batchTaskEnvironmentImageURL),
				resource.TestCheckResourceAttr(resourceName, "security_context.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_context.0.fs_group", "10"),
				resource.TestCheckResourceAttr(resourceName, "security_context.0.run_as_group", "10"),
				resource.TestCheckResourceAttr(resourceName, "security_context.0.run_as_user", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "volumes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.local_mount_directory_path", "localMountDirectoryPath"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.mount_target_export_path", "mountTargetExportPath"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.mount_target_fqdn", "mountTargetFqdn"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "volumes.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "working_directory", "workingDirectory"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_task_environments", "test_batch_task_environments", acctest.Optional, acctest.Update, BatchBatchTaskEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr + BatchBatchTaskEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_task_environment", "test_batch_task_environment", acctest.Optional, acctest.Update, BatchBatchTaskEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "batch_task_environment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "batch_task_environment_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_task_environment", "test_batch_task_environment", acctest.Required, acctest.Create, BatchBatchTaskEnvironmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BatchBatchTaskEnvironmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "batch_task_environment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "image_url", batchTaskEnvironmentImageURL),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_context.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_context.0.fs_group", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_context.0.run_as_group", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_context.0.run_as_user", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "volumes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "volumes.0.local_mount_directory_path", "localMountDirectoryPath"),
				resource.TestCheckResourceAttr(singularDatasourceName, "volumes.0.mount_target_export_path", "mountTargetExportPath"),
				resource.TestCheckResourceAttr(singularDatasourceName, "volumes.0.mount_target_fqdn", "mountTargetFqdn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "volumes.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "volumes.0.type", "NFS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "working_directory", "workingDirectory"),
			),
		},
		// verify resource import
		{
			Config:                  config + BatchBatchTaskEnvironmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckBatchBatchTaskEnvironmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BatchComputingClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_batch_batch_task_environment" {
			noResourceFound = false
			request := oci_batch.GetBatchTaskEnvironmentRequest{}

			tmp := rs.Primary.ID
			request.BatchTaskEnvironmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "batch")

			response, err := client.GetBatchTaskEnvironment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_batch.BatchTaskEnvironmentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BatchBatchTaskEnvironment") {
		resource.AddTestSweepers("BatchBatchTaskEnvironment", &resource.Sweeper{
			Name:         "BatchBatchTaskEnvironment",
			Dependencies: acctest.DependencyGraph["batchTaskEnvironment"],
			F:            sweepBatchBatchTaskEnvironmentResource,
		})
	}
}

func sweepBatchBatchTaskEnvironmentResource(compartment string) error {
	batchComputingClient := acctest.GetTestClients(&schema.ResourceData{}).BatchComputingClient()
	batchTaskEnvironmentIds, err := getBatchBatchTaskEnvironmentIds(compartment)
	if err != nil {
		return err
	}
	for _, batchTaskEnvironmentId := range batchTaskEnvironmentIds {
		if ok := acctest.SweeperDefaultResourceId[batchTaskEnvironmentId]; !ok {
			deleteBatchTaskEnvironmentRequest := oci_batch.DeleteBatchTaskEnvironmentRequest{}

			deleteBatchTaskEnvironmentRequest.BatchTaskEnvironmentId = &batchTaskEnvironmentId

			deleteBatchTaskEnvironmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "batch")
			_, error := batchComputingClient.DeleteBatchTaskEnvironment(context.Background(), deleteBatchTaskEnvironmentRequest)
			if error != nil {
				fmt.Printf("Error deleting BatchTaskEnvironment %s %s, It is possible that the resource is already deleted. Please verify manually \n", batchTaskEnvironmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &batchTaskEnvironmentId, BatchBatchTaskEnvironmentSweepWaitCondition, time.Duration(3*time.Minute),
				BatchBatchTaskEnvironmentSweepResponseFetchOperation, "batch", true)
		}
	}
	return nil
}

func getBatchBatchTaskEnvironmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BatchTaskEnvironmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	batchComputingClient := acctest.GetTestClients(&schema.ResourceData{}).BatchComputingClient()

	listBatchTaskEnvironmentsRequest := oci_batch.ListBatchTaskEnvironmentsRequest{}
	listBatchTaskEnvironmentsRequest.CompartmentId = &compartmentId
	listBatchTaskEnvironmentsRequest.LifecycleState = oci_batch.BatchTaskEnvironmentLifecycleStateActive
	listBatchTaskEnvironmentsResponse, err := batchComputingClient.ListBatchTaskEnvironments(context.Background(), listBatchTaskEnvironmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BatchTaskEnvironment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, batchTaskEnvironment := range listBatchTaskEnvironmentsResponse.Items {
		id := *batchTaskEnvironment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BatchTaskEnvironmentId", id)
	}
	return resourceIds, nil
}

func BatchBatchTaskEnvironmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if batchTaskEnvironmentResponse, ok := response.Response.(oci_batch.GetBatchTaskEnvironmentResponse); ok {
		return batchTaskEnvironmentResponse.LifecycleState != oci_batch.BatchTaskEnvironmentLifecycleStateDeleted
	}
	return false
}

func BatchBatchTaskEnvironmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BatchComputingClient().GetBatchTaskEnvironment(context.Background(), oci_batch.GetBatchTaskEnvironmentRequest{
		BatchTaskEnvironmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
