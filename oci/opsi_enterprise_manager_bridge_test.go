// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v48/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v48/opsi"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	EnterpriseManagerBridgeRequiredOnlyResource = EnterpriseManagerBridgeResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Required, Create, enterpriseManagerBridgeRepresentation)

	EnterpriseManagerBridgeResourceConfig = EnterpriseManagerBridgeResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Update, enterpriseManagerBridgeRepresentation)

	enterpriseManagerBridgeSingularDataSourceRepresentation = map[string]interface{}{
		"enterprise_manager_bridge_id": Representation{RepType: Required, Create: `${oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id}`},
	}

	enterpriseManagerBridgeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`},
		"id":             Representation{RepType: Optional, Create: `${oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id}`},
		"state":          Representation{RepType: Optional, Create: []string{`ACTIVE`}},
		"filter":         RepresentationGroup{Required, enterpriseManagerBridgeDataSourceFilterRepresentation}}
	enterpriseManagerBridgeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id}`}},
	}

	enterpriseManagerBridgeRepresentation = map[string]interface{}{
		"compartment_id":             Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":               Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"object_storage_bucket_name": Representation{RepType: Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"defined_tags":               Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"freeform_tags":              Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                  RepresentationGroup{Required, ignoreChangesEnterpriseManagerBridgeRepresentation},
	}

	ignoreChangesEnterpriseManagerBridgeRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`}},
	}

	EnterpriseManagerBridgeResourceDependencies = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiEnterpriseManagerBridgeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiEnterpriseManagerBridgeResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge"
	datasourceName := "data.oci_opsi_enterprise_manager_bridges.test_enterprise_manager_bridges"
	singularDatasourceName := "data.oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+EnterpriseManagerBridgeResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Create, enterpriseManagerBridgeRepresentation), "opsi", "enterpriseManagerBridge", t)

	ResourceTest(t, testAccCheckOpsiEnterpriseManagerBridgeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Required, Create, enterpriseManagerBridgeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Create, enterpriseManagerBridgeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Create,
					RepresentationCopyWithNewProperties(enterpriseManagerBridgeRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Update, enterpriseManagerBridgeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_opsi_enterprise_manager_bridges", "test_enterprise_manager_bridges", Optional, Update, enterpriseManagerBridgeDataSourceRepresentation) +
				compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Update, enterpriseManagerBridgeRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
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
				GenerateDataSourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Required, Create, enterpriseManagerBridgeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + EnterpriseManagerBridgeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "enterprise_manager_bridge_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
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
	client := testAccProvider.Meta().(*OracleClients).operationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_enterprise_manager_bridge" {
			noResourceFound = false
			request := oci_opsi.GetEnterpriseManagerBridgeRequest{}

			tmp := rs.Primary.ID
			request.EnterpriseManagerBridgeId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "opsi")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("OpsiEnterpriseManagerBridge") {
		resource.AddTestSweepers("OpsiEnterpriseManagerBridge", &resource.Sweeper{
			Name:         "OpsiEnterpriseManagerBridge",
			Dependencies: DependencyGraph["enterpriseManagerBridge"],
			F:            sweepOpsiEnterpriseManagerBridgeResource,
		})
	}
}

func sweepOpsiEnterpriseManagerBridgeResource(compartment string) error {
	operationsInsightsClient := GetTestClients(&schema.ResourceData{}).operationsInsightsClient()
	enterpriseManagerBridgeIds, err := getEnterpriseManagerBridgeIds(compartment)
	if err != nil {
		return err
	}
	for _, enterpriseManagerBridgeId := range enterpriseManagerBridgeIds {
		if ok := SweeperDefaultResourceId[enterpriseManagerBridgeId]; !ok {
			deleteEnterpriseManagerBridgeRequest := oci_opsi.DeleteEnterpriseManagerBridgeRequest{}

			deleteEnterpriseManagerBridgeRequest.EnterpriseManagerBridgeId = &enterpriseManagerBridgeId

			deleteEnterpriseManagerBridgeRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteEnterpriseManagerBridge(context.Background(), deleteEnterpriseManagerBridgeRequest)
			if error != nil {
				fmt.Printf("Error deleting EnterpriseManagerBridge %s %s, It is possible that the resource is already deleted. Please verify manually \n", enterpriseManagerBridgeId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &enterpriseManagerBridgeId, enterpriseManagerBridgeSweepWaitCondition, time.Duration(3*time.Minute),
				enterpriseManagerBridgeSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getEnterpriseManagerBridgeIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "EnterpriseManagerBridgeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := GetTestClients(&schema.ResourceData{}).operationsInsightsClient()

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
		AddResourceIdToSweeperResourceIdMap(compartmentId, "EnterpriseManagerBridgeId", id)
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

func enterpriseManagerBridgeSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.operationsInsightsClient().GetEnterpriseManagerBridge(context.Background(), oci_opsi.GetEnterpriseManagerBridgeRequest{
		EnterpriseManagerBridgeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
