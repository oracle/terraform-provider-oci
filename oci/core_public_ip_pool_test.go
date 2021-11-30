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
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_core "github.com/oracle/oci-go-sdk/v53/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PublicIpPoolRequiredOnlyResource = PublicIpPoolResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Required, Create, publicIpPoolRepresentation)

	PublicIpPoolResourceConfig = PublicIpPoolResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Update, publicIpPoolRepresentation)

	publicIpPoolSingularDataSourceRepresentation = map[string]interface{}{
		"public_ip_pool_id": Representation{RepType: Required, Create: `${oci_core_public_ip_pool.test_public_ip_pool.id}`},
	}

	publicIpPoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"byoip_range_id": Representation{RepType: Optional, Create: `${var.byoip_range_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         RepresentationGroup{Required, publicIpPoolDataSourceFilterRepresentation}}
	publicIpPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_public_ip_pool.test_public_ip_pool.id}`}},
	}

	publicIpPoolRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	PublicIpPoolResourceDependencies = DefinedTagsDependencies + byoipRangeIdVariableStr + publicIpPoolCidrBlockVariableStr
)

// issue-routing-tag: core/vcnip
func TestCorePublicIpPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePublicIpPoolResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_public_ip_pool.test_public_ip_pool"
	datasourceName := "data.oci_core_public_ip_pools.test_public_ip_pools"
	singularDatasourceName := "data.oci_core_public_ip_pool.test_public_ip_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+PublicIpPoolResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Create, publicIpPoolRepresentation), "core", "publicIpPool", t)

	ResourceTest(t, testAccCheckCorePublicIpPoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PublicIpPoolResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Required, Create, publicIpPoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PublicIpPoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PublicIpPoolResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Create, publicIpPoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PublicIpPoolResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Create,
					RepresentationCopyWithNewProperties(publicIpPoolRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + PublicIpPoolResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Update, publicIpPoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				GenerateDataSourceFromRepresentationMap("oci_core_public_ip_pools", "test_public_ip_pools", Optional, Update, publicIpPoolDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
				compartmentIdVariableStr + PublicIpPoolResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Update, publicIpPoolRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "byoip_range_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "public_ip_pool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "public_ip_pool_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Required, Create, publicIpPoolSingularDataSourceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
				compartmentIdVariableStr + PublicIpPoolResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cidr_blocks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + PublicIpPoolResourceConfig,
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

func testAccCheckCorePublicIpPoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_public_ip_pool" {
			noResourceFound = false
			request := oci_core.GetPublicIpPoolRequest{}

			tmp := rs.Primary.ID
			request.PublicIpPoolId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")

			response, err := client.GetPublicIpPool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.PublicIpPoolLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("CorePublicIpPool") {
		resource.AddTestSweepers("CorePublicIpPool", &resource.Sweeper{
			Name:         "CorePublicIpPool",
			Dependencies: DependencyGraph["publicIpPool"],
			F:            sweepCorePublicIpPoolResource,
		})
	}
}

func sweepCorePublicIpPoolResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	publicIpPoolIds, err := getPublicIpPoolIds(compartment)
	if err != nil {
		return err
	}
	for _, publicIpPoolId := range publicIpPoolIds {
		if ok := SweeperDefaultResourceId[publicIpPoolId]; !ok {
			deletePublicIpPoolRequest := oci_core.DeletePublicIpPoolRequest{}

			deletePublicIpPoolRequest.PublicIpPoolId = &publicIpPoolId

			deletePublicIpPoolRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeletePublicIpPool(context.Background(), deletePublicIpPoolRequest)
			if error != nil {
				fmt.Printf("Error deleting PublicIpPool %s %s, It is possible that the resource is already deleted. Please verify manually \n", publicIpPoolId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &publicIpPoolId, publicIpPoolSweepWaitCondition, time.Duration(3*time.Minute),
				publicIpPoolSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getPublicIpPoolIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "PublicIpPoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listPublicIpPoolsRequest := oci_core.ListPublicIpPoolsRequest{}
	listPublicIpPoolsRequest.CompartmentId = &compartmentId
	listPublicIpPoolsResponse, err := virtualNetworkClient.ListPublicIpPools(context.Background(), listPublicIpPoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PublicIpPool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, publicIpPool := range listPublicIpPoolsResponse.Items {
		id := *publicIpPool.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "PublicIpPoolId", id)
	}
	return resourceIds, nil
}

func publicIpPoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if publicIpPoolResponse, ok := response.Response.(oci_core.GetPublicIpPoolResponse); ok {
		return publicIpPoolResponse.LifecycleState != oci_core.PublicIpPoolLifecycleStateDeleted
	}
	return false
}

func publicIpPoolSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetPublicIpPool(context.Background(), oci_core.GetPublicIpPoolRequest{
		PublicIpPoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
