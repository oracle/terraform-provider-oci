// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/oracle/oci-go-sdk/v46/common"
	oci_core "github.com/oracle/oci-go-sdk/v46/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	byoipRangeSingularDataSourceRepresentation = map[string]interface{}{
		"byoip_range_id": Representation{repType: Required, create: `${var.byoip_range_id}`},
	}

	byoipRangeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
	}

	ByoipRangeResourceConfig = byoipRangeIdVariableStr
)

// issue-routing-tag: core/vcnip
func TestCoreByoipRangeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreByoipRangeResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_byoip_ranges.test_byoip_ranges"
	singularDatasourceName := "data.oci_core_byoip_range.test_byoip_range"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_core_byoip_ranges", "test_byoip_ranges", Required, Create, byoipRangeDataSourceRepresentation) +
				compartmentIdVariableStr + ByoipRangeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "byoip_range_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_core_byoip_range", "test_byoip_range", Required, Create, byoipRangeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ByoipRangeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "byoip_range_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cidr_block", publicIpPoolCidrBlock),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "test_BYOIP_range_do_not_delete"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "validation_token"),
			),
		},
	})
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreByoipRange") {
		resource.AddTestSweepers("CoreByoipRange", &resource.Sweeper{
			Name:         "CoreByoipRange",
			Dependencies: DependencyGraph["byoipRange"],
			F:            sweepCoreByoipRangeResource,
		})
	}
}

func sweepCoreByoipRangeResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	byoipRangeIds, err := getByoipRangeIds(compartment)
	if err != nil {
		return err
	}
	for _, byoipRangeId := range byoipRangeIds {
		if ok := SweeperDefaultResourceId[byoipRangeId]; !ok {
			deleteByoipRangeRequest := oci_core.DeleteByoipRangeRequest{}

			deleteByoipRangeRequest.ByoipRangeId = &byoipRangeId

			deleteByoipRangeRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteByoipRange(context.Background(), deleteByoipRangeRequest)
			if error != nil {
				fmt.Printf("Error deleting ByoipRange %s %s, It is possible that the resource is already deleted. Please verify manually \n", byoipRangeId, error)
				continue
			}
			waitTillCondition(testAccProvider, &byoipRangeId, byoipRangeSweepWaitCondition, time.Duration(3*time.Minute),
				byoipRangeSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getByoipRangeIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ByoipRangeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listByoipRangesRequest := oci_core.ListByoipRangesRequest{}
	listByoipRangesRequest.CompartmentId = &compartmentId
	tmp := "ACTIVE"
	listByoipRangesRequest.LifecycleState = &tmp
	listByoipRangesResponse, err := virtualNetworkClient.ListByoipRanges(context.Background(), listByoipRangesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ByoipRange list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, byoipRange := range listByoipRangesResponse.Items {
		id := *byoipRange.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ByoipRangeId", id)
	}
	return resourceIds, nil
}

func byoipRangeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if byoipRangeResponse, ok := response.Response.(oci_core.GetByoipRangeResponse); ok {
		return byoipRangeResponse.LifecycleState != oci_core.ByoipRangeLifecycleStateDeleted
	}
	return false
}

func byoipRangeSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetByoipRange(context.Background(), oci_core.GetByoipRangeRequest{
		ByoipRangeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
