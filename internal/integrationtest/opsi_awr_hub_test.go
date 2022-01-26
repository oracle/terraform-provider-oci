// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v56/opsi"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AwrHubRequiredOnlyResource = AwrHubResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Required, acctest.Create, awrHubRepresentation)

	AwrHubResourceConfig = AwrHubResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Optional, acctest.Update, awrHubRepresentation)

	awrHubSingularDataSourceRepresentation = map[string]interface{}{
		"awr_hub_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
	}

	awrHubDataSourceRepresentation = map[string]interface{}{
		"operations_insights_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                               acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"state":                            acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"filter":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: awrHubDataSourceFilterRepresentation}}
	awrHubDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_awr_hub.test_awr_hub.id}`}},
	}

	awrHubRepresentation = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"object_storage_bucket_name":       acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"operations_insights_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesawrHubRepresentation},
	}

	ignoreChangesawrHubRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	AwrHubResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Required, acctest.Create, operationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiAwrHubResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiAwrHubResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_opsi_awr_hub.test_awr_hub"
	datasourceName := "data.oci_opsi_awr_hubs.test_awr_hubs"
	singularDatasourceName := "data.oci_opsi_awr_hub.test_awr_hub"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AwrHubResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Optional, acctest.Create, awrHubRepresentation), "operationsinsights", "awrHub", t)

	acctest.ResourceTest(t, testAccCheckOpsiAwrHubDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + AwrHubResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Required, acctest.Create, awrHubRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + AwrHubResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + AwrHubResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Optional, acctest.Create, awrHubRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AwrHubResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Optional, acctest.Update, awrHubRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hubs", "test_awr_hubs", acctest.Optional, acctest.Update, awrHubDataSourceRepresentation) +
				compartmentIdVariableStr + AwrHubResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Optional, acctest.Update, awrHubRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "operations_insights_warehouse_id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "awr_hub_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "awr_hub_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Required, acctest.Create, awrHubSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AwrHubResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "awr_hub_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "awr_mailbox_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AwrHubResourceConfig,
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

func testAccCheckOpsiAwrHubDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_awr_hub" {
			noResourceFound = false
			request := oci_opsi.GetAwrHubRequest{}

			tmp := rs.Primary.ID
			request.AwrHubId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetAwrHub(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.AwrHubLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpsiAwrHub") {
		resource.AddTestSweepers("OpsiAwrHub", &resource.Sweeper{
			Name:         "OpsiAwrHub",
			Dependencies: acctest.DependencyGraph["awrHub"],
			F:            sweepOpsiAwrHubResource,
		})
	}
}

func sweepOpsiAwrHubResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	awrHubIds, err := getAwrHubIds(compartment)
	if err != nil {
		return err
	}
	for _, awrHubId := range awrHubIds {
		if ok := acctest.SweeperDefaultResourceId[awrHubId]; !ok {
			deleteAwrHubRequest := oci_opsi.DeleteAwrHubRequest{}

			deleteAwrHubRequest.AwrHubId = &awrHubId

			deleteAwrHubRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteAwrHub(context.Background(), deleteAwrHubRequest)
			if error != nil {
				fmt.Printf("Error deleting AwrHub %s %s, It is possible that the resource is already deleted. Please verify manually \n", awrHubId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &awrHubId, awrHubSweepWaitCondition, time.Duration(3*time.Minute),
				awrHubSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getAwrHubIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AwrHubId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listAwrHubsRequest := oci_opsi.ListAwrHubsRequest{}
	listAwrHubsRequest.CompartmentId = &compartmentId

	operationsInsightsWarehouseIds, error := getOperationsInsightsWarehouseIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting operationsInsightsWarehouseId required for AwrHub resource requests \n")
	}
	for _, operationsInsightsWarehouseId := range operationsInsightsWarehouseIds {
		listAwrHubsRequest.OperationsInsightsWarehouseId = &operationsInsightsWarehouseId

		listAwrHubsRequest.LifecycleState = []oci_opsi.AwrHubLifecycleStateEnum{oci_opsi.AwrHubLifecycleStateActive}
		listAwrHubsResponse, err := operationsInsightsClient.ListAwrHubs(context.Background(), listAwrHubsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting AwrHub list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, awrHub := range listAwrHubsResponse.Items {
			id := *awrHub.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AwrHubId", id)
		}

	}
	return resourceIds, nil
}

func awrHubSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if awrHubResponse, ok := response.Response.(oci_opsi.GetAwrHubResponse); ok {
		return awrHubResponse.LifecycleState != oci_opsi.AwrHubLifecycleStateDeleted
	}
	return false
}

func awrHubSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetAwrHub(context.Background(), oci_opsi.GetAwrHubRequest{
		AwrHubId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
