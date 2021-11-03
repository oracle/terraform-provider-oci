// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	historyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{RepType: Required, Create: `true`},
		"name":                      Representation{RepType: Optional, Create: `name`},
		"recommendation_id":         Representation{RepType: Optional, Create: `${oci_optimizer_recommendation.test_recommendation.recommendation_id}`},
		"recommendation_name":       Representation{RepType: Optional, Create: `${oci_optimizer_recommendation.test_recommendation.name}`},
		"resource_type":             Representation{RepType: Optional, Create: `resourceType`},
		"state":                     Representation{RepType: Optional, Create: `ACTIVE`},
		"status":                    Representation{RepType: Optional, Create: `PENDING`},
	}

	HistoryResourceConfig = RecommendationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationRepresentation)
)

// issue-routing-tag: optimizer/default
func TestOptimizerHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_optimizer_histories.test_histories"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_optimizer_histories", "test_histories", Required, Create, historyDataSourceRepresentation) +
				compartmentIdVariableStr + HistoryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.category_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.compartment_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.estimated_cost_saving"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.recommendation_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.recommendation_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.resource_action_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.resource_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.action.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.time_created"),
			),
		},
	})
}
