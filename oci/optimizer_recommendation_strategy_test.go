// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	recommendationStrategySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{repType: Required, create: `true`},
		"name":                      Representation{repType: Required, create: `name`},
		"recommendation_name":       Representation{repType: Required, create: `${oci_optimizer_recommendation.test_recommendation.name}`},
	}

	recommendationStrategyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{repType: Required, create: `true`},
		"name":                      Representation{repType: Required, create: `name`},
		"recommendation_name":       Representation{repType: Required, create: `${oci_optimizer_recommendation.test_recommendation.name}`},
	}

	RecommendationStrategyResourceConfig = RecommendationResourceDependencies + generateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationRepresentation)
)

// issue-routing-tag: optimizer/default
func TestOptimizerRecommendationStrategyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerRecommendationStrategyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_optimizer_recommendation_strategies.test_recommendation_strategies"
	singularDatasourceName := "data.oci_optimizer_recommendation_strategy.test_recommendation_strategy"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_optimizer_recommendation_strategies", "test_recommendation_strategies", Required, Create, recommendationStrategyDataSourceRepresentation) +
					compartmentIdVariableStr + RecommendationStrategyResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					generateDataSourceFromRepresentationMap("oci_optimizer_recommendation_strategy", "test_recommendation_strategy", Required, Create, recommendationStrategySingularDataSourceRepresentation) +
					compartmentIdVariableStr + RecommendationStrategyResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id_in_subtree", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "recommendation_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
				),
			},
		},
	})
}
