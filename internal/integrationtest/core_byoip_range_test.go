// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreCoreByoipRangeSingularDataSourceRepresentation = map[string]interface{}{
		"byoip_range_id": acctest.Representation{RepType: acctest.Required, Create: `${var.byoip_range_id}`},
	}

	CoreCoreByoipRangeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}
	CoreCoreByoipv6RangeSingularDataSouorceRepresentation = map[string]interface{}{
		"byoip_range_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("byoipv6_range_ocid")},
	}

	byoipv6RangeId            = utils.GetEnvSettingWithBlankDefault("byoipv6_range_ocid")
	byoipv6RangeIdVariableStr = fmt.Sprintf("variable \"byoipv6_range_id\" { default = \"%s\" }\n", byoipv6RangeId)

	ByoipRangeResourceConfig   = byoipRangeIdVariableStr
	Byoipv6RangeResourceConfig = byoipv6RangeIdVariableStr
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_byoip_ranges", "test_byoip_ranges", acctest.Required, acctest.Create, CoreCoreByoipRangeDataSourceRepresentation) +
				compartmentIdVariableStr + ByoipRangeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "byoip_range_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_byoip_range", "test_byoip_range", acctest.Required, acctest.Create, CoreCoreByoipRangeSingularDataSourceRepresentation) +
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

		// verify byoipv6 singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_byoip_range", "test_byoip_range", acctest.Required, acctest.Create, CoreCoreByoipv6RangeSingularDataSouorceRepresentation) +
				compartmentIdVariableStr + Byoipv6RangeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "byoip_range_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttr(singularDatasourceName, "cidr_block", publicIpPoolCidrBlock),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "test_BYOIP_range_do_not_delete"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ipv6cidr_block", "2000:1000::/48"),
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
	byoipRangeIds, err := getCoreByoipRangeIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &byoipRangeId, CoreByoipRangeSweepWaitCondition, time.Duration(3*time.Minute),
				CoreByoipRangeSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreByoipRangeIds(compartment string) ([]string, error) {
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

func CoreByoipRangeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if byoipRangeResponse, ok := response.Response.(oci_core.GetByoipRangeResponse); ok {
		return byoipRangeResponse.LifecycleState != oci_core.ByoipRangeLifecycleStateDeleted
	}
	return false
}

func CoreByoipRangeSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetByoipRange(context.Background(), oci_core.GetByoipRangeRequest{
		ByoipRangeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
