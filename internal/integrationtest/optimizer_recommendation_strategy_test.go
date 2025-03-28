// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OptimizerOptimizerRecommendationStrategySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"name":                      acctest.Representation{RepType: acctest.Required, Create: `name`},
		"recommendation_name":       acctest.Representation{RepType: acctest.Required, Create: `${oci_optimizer_recommendation.test_recommendation.name}`},
	}

	OptimizerOptimizerRecommendationStrategyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"name":                      acctest.Representation{RepType: acctest.Required, Create: `name`},
		"recommendation_name":       acctest.Representation{RepType: acctest.Required, Create: `${oci_optimizer_recommendation.test_recommendation.name}`},
	}

	OptimizerRecommendationStrategyResourceConfig = OptimizerRecommendationResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Required, acctest.Create, OptimizerRecommendationRepresentation)
)

// issue-routing-tag: optimizer/default
func TestOptimizerRecommendationStrategyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerRecommendationStrategyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_optimizer_recommendation_strategies.test_recommendation_strategies"
	singularDatasourceName := "data.oci_optimizer_recommendation_strategy.test_recommendation_strategy"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_recommendation_strategies", "test_recommendation_strategies", acctest.Required, acctest.Create, OptimizerOptimizerRecommendationStrategyDataSourceRepresentation) +
				compartmentIdVariableStr + OptimizerRecommendationStrategyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "name"),
				resource.TestCheckResourceAttrSet(datasourceName, "recommendation_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "recommendation_strategy_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "recommendation_strategy_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_recommendation_strategy", "test_recommendation_strategy", acctest.Required, acctest.Create, OptimizerOptimizerRecommendationStrategySingularDataSourceRepresentation) +
				compartmentIdVariableStr + OptimizerRecommendationStrategyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recommendation_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
