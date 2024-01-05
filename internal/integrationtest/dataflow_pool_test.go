// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataflowPoolRequiredOnlyResource = DataflowPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Required, acctest.Create, DataflowPoolRepresentation)

	DataflowPoolResourceConfig = DataflowPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Optional, acctest.Update, DataflowPoolRepresentation)

	DataflowDataflowPoolSingularDataSourceRepresentation = map[string]interface{}{
		"pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_pool.test_pool.id}`},
	}

	DataflowDataflowPoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `myPool`, Update: `displayName2`},
		"display_name_starts_with": acctest.Representation{RepType: acctest.Optional, Create: `displayNameStartsWith`},
		//"owner_principal_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_dataflow_owner_principal.test_owner_principal.id}`},
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `SCHEDULED`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowPoolDataSourceFilterRepresentation}}
	DataflowPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataflow_pool.test_pool.id}`}},
	}

	DataflowPoolRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configurations": acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowPoolConfigurationsRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `myPool`, Update: `displayName2`},
		//"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		//"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"idle_timeout_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"schedules":               acctest.RepresentationGroup{RepType: acctest.Required, Group: DataflowPoolSchedulesRepresentation},
		//"state":                   acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
	}
	DataflowPoolConfigurationsRepresentation = map[string]interface{}{
		"max":   acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `3`},
		"min":   acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `0`},
		"shape": acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`, Update: `VM.Standard2.2`},
		//"shape_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataflowPoolConfigurationsShapeConfigRepresentation},
	}
	DataflowPoolSchedulesRepresentation = map[string]interface{}{
		"day_of_week": acctest.Representation{RepType: acctest.Required, Create: `SUNDAY`, Update: `MONDAY`},
		"start_time":  acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		//"stop_time":   acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	//DataflowPoolConfigurationsShapeConfigRepresentation = map[string]interface{}{
	//	"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.1`},
	//	"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.1`},
	//}

	DataflowPoolResourceDependencies = "" //DefinedTagsDependencies
)

// issue-routing-tag: dataflow/default
func TestDataflowPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowPoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dataflow_pool.test_pool"
	datasourceName := "data.oci_dataflow_pools.test_pools"
	singularDatasourceName := "data.oci_dataflow_pool.test_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataflowPoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Optional, acctest.Create, DataflowPoolRepresentation), "dataflow", "pool", t)

	acctest.ResourceTest(t, testAccCheckDataflowPoolDestroy, []resource.TestStep{
		// verify Create
		{
			PreConfig: func() {
				fmt.Println("step 1")
			},
			Config: config + compartmentIdVariableStr + DataflowPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Required, acctest.Create, DataflowPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myPool"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			PreConfig: func() {
				fmt.Println("step 2")
			},
			Config: config + compartmentIdVariableStr + DataflowPoolResourceDependencies,
		},
		// verify Create with optionals
		{
			PreConfig: func() {
				fmt.Println("step 3")
			},
			Config: config + compartmentIdVariableStr + DataflowPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Optional, acctest.Create, DataflowPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.max", "3"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.min", "0"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.shape", "VM.Standard2.1"),
				//resource.TestCheckResourceAttr(resourceName, "configurations.0.shape_config.0.memory_in_gbs", "16"),
				//resource.TestCheckResourceAttr(resourceName, "configurations.0.shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myPool"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idle_timeout_in_minutes", "10"),
				//resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.day_of_week", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.start_time", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//			return errExport
					//}
					//}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			PreConfig: func() {
				fmt.Println("step 4")
			},
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataflowPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataflowPoolRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.max", "3"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.min", "0"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.shape", "VM.Standard2.1"),
				//resource.TestCheckResourceAttr(resourceName, "configurations.0.shape_config.0.memory_in_gbs", "16"),
				//resource.TestCheckResourceAttr(resourceName, "configurations.0.shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myPool"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idle_timeout_in_minutes", "10"),
				//resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.day_of_week", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.start_time", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			PreConfig: func() {
				fmt.Println("step 5")
			},
			Config: config + compartmentIdVariableStr + DataflowPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Optional, acctest.Update, DataflowPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.max", "3"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.min", "0"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.shape", "VM.Standard2.2"),
				//resource.TestCheckResourceAttr(resourceName, "configurations.0.shape_config.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "configurations.0.shape_config.0.memory_in_gbs", 16),
				//resource.TestCheckResourceAttr(resourceName, "configurations.0.shape_config.0.ocpus", 1),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idle_timeout_in_minutes", "11"),
				//resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.day_of_week", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "schedules.0.start_time", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			PreConfig: func() {
				fmt.Println("step 6")
			},
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_pools", "test_pools", acctest.Optional, acctest.Update, DataflowDataflowPoolDataSourceRepresentation) +
				compartmentIdVariableStr + DataflowPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Optional, acctest.Update, DataflowPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_starts_with", "displayNameStartsWith"),
				//resource.TestCheckResourceAttrSet(datasourceName, "owner_principal_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "SCHEDULED"),

				resource.TestCheckResourceAttr(datasourceName, "pool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pool_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			PreConfig: func() {
				fmt.Println("step 7")
			},
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_pool", "test_pool", acctest.Required, acctest.Create, DataflowDataflowPoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataflowPoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configurations.0.max", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configurations.0.min", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configurations.0.shape", "VM.Standard2.2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "configurations.0.shape_config.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "configurations.0.shape_config.0.memory_in_gbs", 16),
				//resource.TestCheckResourceAttr(singularDatasourceName, "configurations.0.shape_config.0.ocpus", 1),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idle_timeout_in_minutes", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner_user_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pool_metrics.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.day_of_week", "MONDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.start_time", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "SCHEDULED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataflowPoolRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataflowPoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataFlowClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataflow_pool" {
			noResourceFound = false
			request := oci_dataflow.GetPoolRequest{}

			tmp := rs.Primary.ID
			request.PoolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")

			response, err := client.GetPool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dataflow.PoolLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataflowPool") {
		resource.AddTestSweepers("DataflowPool", &resource.Sweeper{
			Name:         "DataflowPool",
			Dependencies: acctest.DependencyGraph["pool"],
			F:            sweepDataflowPoolResource,
		})
	}
}

func sweepDataflowPoolResource(compartment string) error {
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()
	poolIds, err := getDataflowPoolIds(compartment)
	if err != nil {
		return err
	}
	for _, poolId := range poolIds {
		if ok := acctest.SweeperDefaultResourceId[poolId]; !ok {

			stopPoolRequest := oci_dataflow.StopPoolRequest{}
			stopPoolRequest.PoolId = &poolId
			stopPoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")
			_, errorStopping := dataFlowClient.StopPool(context.Background(), stopPoolRequest)
			if errorStopping != nil {
				fmt.Printf("Error stopping Pool %s %s, It is possible that the resource is already stopped. Please verify manually \n", poolId, errorStopping)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &poolId, DataflowPoolStopWaitCondition, time.Duration(3*time.Minute),
				DataflowPoolSweepResponseFetchOperation, "dataflow", true)

			deletePoolRequest := oci_dataflow.DeletePoolRequest{}
			deletePoolRequest.PoolId = &poolId
			deletePoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")
			_, error := dataFlowClient.DeletePool(context.Background(), deletePoolRequest)
			if error != nil {
				fmt.Printf("Error deleting Pool %s %s, It is possible that the resource is already deleted. Please verify manually \n", poolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &poolId, DataflowPoolSweepWaitCondition, time.Duration(3*time.Minute),
				DataflowPoolSweepResponseFetchOperation, "dataflow", true)
		}
	}
	return nil
}

func getDataflowPoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()

	listPoolsRequest := oci_dataflow.ListPoolsRequest{}
	listPoolsRequest.CompartmentId = &compartmentId
	listPoolsRequest.LifecycleState = oci_dataflow.ListPoolsLifecycleStateActive
	listPoolsResponse, err := dataFlowClient.ListPools(context.Background(), listPoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Pool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, pool := range listPoolsResponse.Items {
		id := *pool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PoolId", id)
	}
	return resourceIds, nil
}

func DataflowPoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if poolResponse, ok := response.Response.(oci_dataflow.GetPoolResponse); ok {
		return poolResponse.LifecycleState != oci_dataflow.PoolLifecycleStateDeleted
	}
	return false
}

func DataflowPoolStopWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	fmt.Println("waiting until pool is stopped")
	if poolResponse, ok := response.Response.(oci_dataflow.GetPoolResponse); ok {
		return poolResponse.LifecycleState != oci_dataflow.PoolLifecycleStateStopped
	}
	return false
}

func DataflowPoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataFlowClient().GetPool(context.Background(), oci_dataflow.GetPoolRequest{
		PoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
