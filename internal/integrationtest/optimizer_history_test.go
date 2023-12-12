// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OptimizerOptimizerHistoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"include_resource_metadata": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"recommendation_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_optimizer_recommendation.test_recommendation.recommendation_id}`},
		"recommendation_name":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_optimizer_recommendation.test_recommendation.name}`},
		"resource_type":             acctest.Representation{RepType: acctest.Optional, Create: `resourceType`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: `PENDING`},
	}

	OptimizerHistoryResourceConfig = OptimizerRecommendationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Required, acctest.Create, OptimizerRecommendationRepresentation)
)

// issue-routing-tag: optimizer/default
func TestOptimizerHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_optimizer_histories.test_histories"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_histories", "test_histories", acctest.Required, acctest.Create, OptimizerOptimizerHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + OptimizerHistoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
