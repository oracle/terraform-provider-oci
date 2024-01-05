// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CorePublicIpPoolRequiredOnlyResource = CorePublicIpPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Required, acctest.Create, CorePublicPoolRepresentation)

	CorePublicIpPoolResourceConfig = CorePublicIpPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Optional, acctest.Update, CorePublicPoolRepresentation)

	CoreCorePublicIpPoolSingularDataSourceRepresentation = map[string]interface{}{
		"public_ip_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_public_ip_pool.test_public_ip_pool.id}`},
	}

	CoreCorePublicIpPoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"byoip_range_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.byoip_range_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CorePublicIpPoolDataSourceFilterRepresentation}}
	CorePublicIpPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_public_ip_pool.test_public_ip_pool.id}`}},
	}

	CorePublicPoolRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CorePublicIpPoolResourceDependencies = DefinedTagsDependencies + byoipRangeIdVariableStr + publicIpPoolCidrBlockVariableStr
)

// issue-routing-tag: core/vcnip
func TestCorePublicIpPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePublicIpPoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_public_ip_pool.test_public_ip_pool"
	datasourceName := "data.oci_core_public_ip_pools.test_public_ip_pools"
	singularDatasourceName := "data.oci_core_public_ip_pool.test_public_ip_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CorePublicIpPoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Optional, acctest.Create, CorePublicPoolRepresentation), "core", "publicIpPool", t)

	acctest.ResourceTest(t, testAccCheckCorePublicIpPoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CorePublicIpPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", acctest.Required, acctest.Create, publicIpPoolCapacityRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Required, acctest.Create, CorePublicPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CorePublicIpPoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CorePublicIpPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", acctest.Required, acctest.Create, publicIpPoolCapacityRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Optional, acctest.Create, CorePublicPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CorePublicIpPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", acctest.Required, acctest.Create, publicIpPoolCapacityRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CorePublicPoolRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
			Config: config + compartmentIdVariableStr + CorePublicIpPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", acctest.Required, acctest.Create, publicIpPoolCapacityRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Optional, acctest.Update, CorePublicPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_public_ip_pools", "test_public_ip_pools", acctest.Optional, acctest.Update, CoreCorePublicIpPoolDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", acctest.Required, acctest.Create, publicIpPoolCapacityRepresentation) +
				compartmentIdVariableStr + CorePublicIpPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Optional, acctest.Update, CorePublicPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Required, acctest.Create, CoreCorePublicIpPoolSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", acctest.Required, acctest.Create, publicIpPoolCapacityRepresentation) +
				compartmentIdVariableStr + CorePublicIpPoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip_pool_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cidr_blocks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CorePublicIpPoolRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCorePublicIpPoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_public_ip_pool" {
			noResourceFound = false
			request := oci_core.GetPublicIpPoolRequest{}

			tmp := rs.Primary.ID
			request.PublicIpPoolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CorePublicIpPool") {
		resource.AddTestSweepers("CorePublicIpPool", &resource.Sweeper{
			Name:         "CorePublicIpPool",
			Dependencies: acctest.DependencyGraph["publicIpPool"],
			F:            sweepCorePublicIpPoolResource,
		})
	}
}

func sweepCorePublicIpPoolResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	publicIpPoolIds, err := getCorePublicIpPoolIds(compartment)
	if err != nil {
		return err
	}
	for _, publicIpPoolId := range publicIpPoolIds {
		if ok := acctest.SweeperDefaultResourceId[publicIpPoolId]; !ok {
			deletePublicIpPoolRequest := oci_core.DeletePublicIpPoolRequest{}

			deletePublicIpPoolRequest.PublicIpPoolId = &publicIpPoolId

			deletePublicIpPoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeletePublicIpPool(context.Background(), deletePublicIpPoolRequest)
			if error != nil {
				fmt.Printf("Error deleting PublicIpPool %s %s, It is possible that the resource is already deleted. Please verify manually \n", publicIpPoolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &publicIpPoolId, CorePublicIpPoolSweepWaitCondition, time.Duration(3*time.Minute),
				CorePublicIpPoolSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCorePublicIpPoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PublicIpPoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listPublicIpPoolsRequest := oci_core.ListPublicIpPoolsRequest{}
	listPublicIpPoolsRequest.CompartmentId = &compartmentId
	listPublicIpPoolsResponse, err := virtualNetworkClient.ListPublicIpPools(context.Background(), listPublicIpPoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PublicIpPool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, publicIpPool := range listPublicIpPoolsResponse.Items {
		id := *publicIpPool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PublicIpPoolId", id)
	}
	return resourceIds, nil
}

func CorePublicIpPoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if publicIpPoolResponse, ok := response.Response.(oci_core.GetPublicIpPoolResponse); ok {
		return publicIpPoolResponse.LifecycleState != oci_core.PublicIpPoolLifecycleStateDeleted
	}
	return false
}

func CorePublicIpPoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetPublicIpPool(context.Background(), oci_core.GetPublicIpPoolRequest{
		PublicIpPoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
