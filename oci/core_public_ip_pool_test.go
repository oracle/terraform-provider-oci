// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v29/common"
	oci_core "github.com/oracle/oci-go-sdk/v29/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PublicIpPoolRequiredOnlyResource = PublicIpPoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Required, Create, publicIpPoolRepresentation)

	PublicIpPoolResourceConfig = PublicIpPoolResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Update, publicIpPoolRepresentation)

	publicIpPoolSingularDataSourceRepresentation = map[string]interface{}{
		"public_ip_pool_id": Representation{repType: Required, create: `${oci_core_public_ip_pool.test_public_ip_pool.id}`},
	}

	publicIpPoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"byoip_range_id": Representation{repType: Optional, create: `${var.byoip_range_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"filter":         RepresentationGroup{Required, publicIpPoolDataSourceFilterRepresentation}}
	publicIpPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_public_ip_pool.test_public_ip_pool.id}`}},
	}

	publicIpPoolRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	PublicIpPoolResourceDependencies = DefinedTagsDependencies + byoipRangeIdVariableStr + publicIpPoolCidrBlockVariableStr
)

func TestCorePublicIpPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePublicIpPoolResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_public_ip_pool.test_public_ip_pool"
	datasourceName := "data.oci_core_public_ip_pools.test_public_ip_pools"
	singularDatasourceName := "data.oci_core_public_ip_pool.test_public_ip_pool"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCorePublicIpPoolDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + PublicIpPoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Required, Create, publicIpPoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + PublicIpPoolResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + PublicIpPoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Create, publicIpPoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PublicIpPoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Create,
						representationCopyWithNewProperties(publicIpPoolRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				Config: config + compartmentIdVariableStr + PublicIpPoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Update, publicIpPoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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
					generateDataSourceFromRepresentationMap("oci_core_public_ip_pools", "test_public_ip_pools", Optional, Update, publicIpPoolDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
					compartmentIdVariableStr + PublicIpPoolResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Optional, Update, publicIpPoolRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", Required, Create, publicIpPoolSingularDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", Required, Create, publicIpPoolCapacityRepresentation) +
					compartmentIdVariableStr + PublicIpPoolResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if !inSweeperExcludeList("CorePublicIpPool") {
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

			deletePublicIpPoolRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeletePublicIpPool(context.Background(), deletePublicIpPoolRequest)
			if error != nil {
				fmt.Printf("Error deleting PublicIpPool %s %s, It is possible that the resource is already deleted. Please verify manually \n", publicIpPoolId, error)
				continue
			}
			waitTillCondition(testAccProvider, &publicIpPoolId, publicIpPoolSweepWaitCondition, time.Duration(3*time.Minute),
				publicIpPoolSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getPublicIpPoolIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "PublicIpPoolId")
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
		addResourceIdToSweeperResourceIdMap(compartmentId, "PublicIpPoolId", id)
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
