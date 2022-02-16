// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/oracle/oci-go-sdk/v58/common"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	byoipRangeSingularDataSourceRepresentation = map[string]interface{}{
		"byoip_range_id": acctest.Representation{RepType: acctest.Required, Create: `${var.byoip_range_id}`},
	}

	byoipRangeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	ByoipRangeResourceConfig = byoipRangeIdVariableStr
)

// issue-routing-tag: core/vcnip
func TestCoreByoipRangeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreByoipRangeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_byoip_ranges.test_byoip_ranges"
	singularDatasourceName := "data.oci_core_byoip_range.test_byoip_range"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_byoip_ranges", "test_byoip_ranges", acctest.Required, acctest.Create, byoipRangeDataSourceRepresentation) +
				compartmentIdVariableStr + ByoipRangeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "byoip_range_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_byoip_range", "test_byoip_range", acctest.Required, acctest.Create, byoipRangeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ByoipRangeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreByoipRange") {
		resource.AddTestSweepers("CoreByoipRange", &resource.Sweeper{
			Name:         "CoreByoipRange",
			Dependencies: acctest.DependencyGraph["byoipRange"],
			F:            sweepCoreByoipRangeResource,
		})
	}
}

func sweepCoreByoipRangeResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	byoipRangeIds, err := getByoipRangeIds(compartment)
	if err != nil {
		return err
	}
	for _, byoipRangeId := range byoipRangeIds {
		if ok := acctest.SweeperDefaultResourceId[byoipRangeId]; !ok {
			deleteByoipRangeRequest := oci_core.DeleteByoipRangeRequest{}

			deleteByoipRangeRequest.ByoipRangeId = &byoipRangeId

			deleteByoipRangeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteByoipRange(context.Background(), deleteByoipRangeRequest)
			if error != nil {
				fmt.Printf("Error deleting ByoipRange %s %s, It is possible that the resource is already deleted. Please verify manually \n", byoipRangeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &byoipRangeId, byoipRangeSweepWaitCondition, time.Duration(3*time.Minute),
				byoipRangeSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getByoipRangeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ByoipRangeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ByoipRangeId", id)
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

func byoipRangeSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetByoipRange(context.Background(), oci_core.GetByoipRangeRequest{
		ByoipRangeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
