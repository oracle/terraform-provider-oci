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
	"github.com/oracle/oci-go-sdk/v43/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v43/opsi"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	EnterpriseManagerBridgeRequiredOnlyResource = EnterpriseManagerBridgeResourceDependencies +
		generateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Required, Create, enterpriseManagerBridgeRepresentation)

	EnterpriseManagerBridgeResourceConfig = EnterpriseManagerBridgeResourceDependencies +
		generateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Update, enterpriseManagerBridgeRepresentation)

	enterpriseManagerBridgeSingularDataSourceRepresentation = map[string]interface{}{
		"enterprise_manager_bridge_id": Representation{repType: Required, create: `${oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id}`},
	}

	enterpriseManagerBridgeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`},
		"id":             Representation{repType: Optional, create: `${oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id}`},
		"state":          Representation{repType: Optional, create: []string{`Active`}},
		"filter":         RepresentationGroup{Required, enterpriseManagerBridgeDataSourceFilterRepresentation}}
	enterpriseManagerBridgeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id}`}},
	}

	enterpriseManagerBridgeRepresentation = map[string]interface{}{
		"compartment_id":             Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":               Representation{repType: Required, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":               Representation{repType: Required, create: `displayName`},
		"freeform_tags":              Representation{repType: Required, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"object_storage_bucket_name": Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"description":                Representation{repType: Required, create: `description`, update: `description2`},
		"lifecycle":                  RepresentationGroup{Required, ignoreChangesEnterpriseManagerBridgeRepresentation},
	}

	// DBX-5754 - Defined_tags should not be required field.
	ignoreChangesEnterpriseManagerBridgeRepresentation = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}

	EnterpriseManagerBridgeResourceDependencies = DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation) +
		generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

func TestOpsiEnterpriseManagerBridgeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiEnterpriseManagerBridgeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge"
	datasourceName := "data.oci_opsi_enterprise_manager_bridges.test_enterprise_manager_bridges"
	singularDatasourceName := "data.oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+EnterpriseManagerBridgeResourceDependencies+
		generateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Create, enterpriseManagerBridgeRepresentation), "opsi", "enterpriseManagerBridge", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOpsiEnterpriseManagerBridgeDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies +
					generateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Required, Create, enterpriseManagerBridgeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies +
					generateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Create, enterpriseManagerBridgeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + EnterpriseManagerBridgeResourceDependencies +
					generateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Create,
						representationCopyWithNewProperties(enterpriseManagerBridgeRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Update, enterpriseManagerBridgeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "object_storage_bucket_name"),
					resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace_name"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_opsi_enterprise_manager_bridges", "test_enterprise_manager_bridges", Optional, Update, enterpriseManagerBridgeDataSourceRepresentation) +
					compartmentIdVariableStr + EnterpriseManagerBridgeResourceDependencies +
					generateResourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Optional, Update, enterpriseManagerBridgeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_opsi_enterprise_manager_bridge", "test_enterprise_manager_bridge", Required, Create, enterpriseManagerBridgeSingularDataSourceRepresentation) +
					compartmentIdVariableStr + EnterpriseManagerBridgeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "enterprise_manager_bridge_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "opsi")

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
	if !inSweeperExcludeList("OpsiEnterpriseManagerBridge") {
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

			deleteEnterpriseManagerBridgeRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteEnterpriseManagerBridge(context.Background(), deleteEnterpriseManagerBridgeRequest)
			if error != nil {
				fmt.Printf("Error deleting EnterpriseManagerBridge %s %s, It is possible that the resource is already deleted. Please verify manually \n", enterpriseManagerBridgeId, error)
				continue
			}
			waitTillCondition(testAccProvider, &enterpriseManagerBridgeId, enterpriseManagerBridgeSweepWaitCondition, time.Duration(3*time.Minute),
				enterpriseManagerBridgeSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getEnterpriseManagerBridgeIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "EnterpriseManagerBridgeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := GetTestClients(&schema.ResourceData{}).operationsInsightsClient()

	listEnterpriseManagerBridgesRequest := oci_opsi.ListEnterpriseManagerBridgesRequest{}
	listEnterpriseManagerBridgesRequest.CompartmentId = &compartmentId
	listEnterpriseManagerBridgesRequest.LifecycleState = []oci_opsi.LifecycleStateEnum{oci_opsi.LifecycleStateActive}
	listEnterpriseManagerBridgesResponse, err := operationsInsightsClient.ListEnterpriseManagerBridges(context.Background(), listEnterpriseManagerBridgesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting EnterpriseManagerBridge list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, enterpriseManagerBridge := range listEnterpriseManagerBridgesResponse.Items {
		id := *enterpriseManagerBridge.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "EnterpriseManagerBridgeId", id)
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
