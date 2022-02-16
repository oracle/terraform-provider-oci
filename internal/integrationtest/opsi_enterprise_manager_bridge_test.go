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
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v58/opsi"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	EnterpriseManagerBridgeRequiredOnlyResource = EnterpriseManagerBridgeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", acctest.Required, acctest.Create, enterpriseManagerBridgeRepresentation)

	EnterpriseManagerBridgeResourceConfig = EnterpriseManagerBridgeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", acctest.Optional, acctest.Update, enterpriseManagerBridgeRepresentation)

	enterpriseManagerBridgeSingularDataSourceRepresentation = map[string]interface{}{
		"enterprise_manager_bridge_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id}`},
	}

	enterpriseManagerBridgeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: enterpriseManagerBridgeDataSourceFilterRepresentation}}
	enterpriseManagerBridgeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id}`}},
	}

	enterpriseManagerBridgeRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"object_storage_bucket_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesEnterpriseManagerBridgeRepresentation},
	}

	ignoreChangesEnterpriseManagerBridgeRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	EnterpriseManagerBridgeResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiEnterpriseManagerBridgeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiEnterpriseManagerBridgeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge"
	datasourceName := "data.oci_opsi_enterprise_manager_bridges.test_enterprise_manager_bridges"
	singularDatasourceName := "data.oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+EnterpriseManagerBridgeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", acctest.Optional, acctest.Create, enterpriseManagerBridgeRepresentation), "opsi", "enterpriseManagerBridge", t)

	acctest.ResourceTest(t, testAccCheckOpsiEnterpriseManagerBridgeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", acctest.Required, acctest.Create, enterpriseManagerBridgeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", acctest.Optional, acctest.Create, enterpriseManagerBridgeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace_name"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + EnterpriseManagerBridgeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(enterpriseManagerBridgeRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace_name"),
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
			Config: config + compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", acctest.Optional, acctest.Update, enterpriseManagerBridgeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace_name"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_enterprise_manager_bridges", "test_enterprise_manager_bridges", acctest.Optional, acctest.Update, enterpriseManagerBridgeDataSourceRepresentation) +
				compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", acctest.Optional, acctest.Update, enterpriseManagerBridgeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"), // This is hash generated value. so we cannot match exact value
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "enterprise_manager_bridge_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "enterprise_manager_bridge_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", acctest.Required, acctest.Create, enterpriseManagerBridgeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + EnterpriseManagerBridgeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "enterprise_manager_bridge_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_storage_bucket_status_details"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_storage_namespace_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + EnterpriseManagerBridgeResourceConfig,
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

func testAccCheckOpsiEnterpriseManagerBridgeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_enterprise_manager_bridge" {
			noResourceFound = false
			request := oci_opsi.GetEnterpriseManagerBridgeRequest{}

			tmp := rs.Primary.ID
			request.EnterpriseManagerBridgeId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetEnterpriseManagerBridge(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpsiEnterpriseManagerBridge") {
		resource.AddTestSweepers("OpsiEnterpriseManagerBridge", &resource.Sweeper{
			Name:         "OpsiEnterpriseManagerBridge",
			Dependencies: acctest.DependencyGraph["enterpriseManagerBridge"],
			F:            sweepOpsiEnterpriseManagerBridgeResource,
		})
	}
}

func sweepOpsiEnterpriseManagerBridgeResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	enterpriseManagerBridgeIds, err := getEnterpriseManagerBridgeIds(compartment)
	if err != nil {
		return err
	}
	for _, enterpriseManagerBridgeId := range enterpriseManagerBridgeIds {
		if ok := acctest.SweeperDefaultResourceId[enterpriseManagerBridgeId]; !ok {
			deleteEnterpriseManagerBridgeRequest := oci_opsi.DeleteEnterpriseManagerBridgeRequest{}

			deleteEnterpriseManagerBridgeRequest.EnterpriseManagerBridgeId = &enterpriseManagerBridgeId

			deleteEnterpriseManagerBridgeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteEnterpriseManagerBridge(context.Background(), deleteEnterpriseManagerBridgeRequest)
			if error != nil {
				fmt.Printf("Error deleting EnterpriseManagerBridge %s %s, It is possible that the resource is already deleted. Please verify manually \n", enterpriseManagerBridgeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &enterpriseManagerBridgeId, enterpriseManagerBridgeSweepWaitCondition, time.Duration(3*time.Minute),
				enterpriseManagerBridgeSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getEnterpriseManagerBridgeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EnterpriseManagerBridgeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listEnterpriseManagerBridgesRequest := oci_opsi.ListEnterpriseManagerBridgesRequest{}
	listEnterpriseManagerBridgesRequest.CompartmentId = &compartmentId
	listEnterpriseManagerBridgesRequest.LifecycleState = []oci_opsi.LifecycleStateEnum{oci_opsi.LifecycleStateActive, oci_opsi.LifecycleStateNeedsAttention}
	listEnterpriseManagerBridgesResponse, err := operationsInsightsClient.ListEnterpriseManagerBridges(context.Background(), listEnterpriseManagerBridgesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting EnterpriseManagerBridge list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, enterpriseManagerBridge := range listEnterpriseManagerBridgesResponse.Items {
		id := *enterpriseManagerBridge.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "EnterpriseManagerBridgeId", id)
	}
	return resourceIds, nil
}

func enterpriseManagerBridgeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if enterpriseManagerBridgeResponse, ok := response.Response.(oci_opsi.GetEnterpriseManagerBridgeResponse); ok {
		return enterpriseManagerBridgeResponse.LifecycleState != oci_opsi.LifecycleStateDeleted
	}
	return false
}

func enterpriseManagerBridgeSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetEnterpriseManagerBridge(context.Background(), oci_opsi.GetEnterpriseManagerBridgeRequest{
		EnterpriseManagerBridgeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
