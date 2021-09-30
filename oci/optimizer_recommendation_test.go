// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	RecommendationRequiredOnlyResource = RecommendationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationRepresentation)

	RecommendationResourceConfig = RecommendationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Optional, Update, recommendationRepresentation)

	recommendationSingularDataSourceRepresentation = map[string]interface{}{
		"recommendation_id": Representation{RepType: Required, Create: `${data.oci_optimizer_recommendations.test_recommendations.recommendation_collection.0.items.0.id}`},
	}

	recommendationDataSourceRepresentation = map[string]interface{}{
		"category_id":               Representation{RepType: Required, Create: `${lookup(data.oci_optimizer_categories.test_categories.category_collection.0.items[0], "id")}`},
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{RepType: Required, Create: `true`},
		"name":                      Representation{RepType: Optional, Create: `name`},
		"state":                     Representation{RepType: Optional, Create: `ACTIVE`},
		"status":                    Representation{RepType: Optional, Create: `PENDING`, Update: `DISMISSED`},
		"filter":                    RepresentationGroup{Required, recommendationDataSourceFilterRepresentation}}

	// to filter the list of recommendation and get one recommendation containing the supportlevels
	// we will use the supportlevels to Create the optimizer profile
	recommendationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `name`},
		"values": Representation{RepType: Required, Create: []string{`cost-management-compute-host-underutilized-name`}},
	}

	recommendationRepresentation = map[string]interface{}{
		"recommendation_id": Representation{RepType: Required, Create: `${data.oci_optimizer_recommendations.test_recommendations.recommendation_collection.0.items.0.id}`},
		"status":            Representation{RepType: Required, Create: `PENDING`, Update: `DISMISSED`},
	}

	RecommendationResourceDependencies = GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", Required, Create, optimizerCategoryDataSourceRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_optimizer_recommendations", "test_recommendations", Required, Create, recommendationDataSourceRepresentation)
)

// issue-routing-tag: optimizer/default
func TestOptimizerRecommendationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerRecommendationResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_optimizer_recommendation.test_recommendation"
	datasourceName := "data.oci_optimizer_recommendations.test_recommendations"
	singularDatasourceName := "data.oci_optimizer_recommendation.test_recommendation"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+RecommendationResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Optional, Create, recommendationRepresentation), "optimizer", "recommendation", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RecommendationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "recommendation_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "PENDING"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RecommendationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RecommendationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Optional, Create, recommendationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "category_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "description"),
				resource.TestCheckResourceAttrSet(resourceName, "estimated_cost_saving"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "importance"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttrSet(resourceName, "recommendation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_counts.#"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "PENDING"),
				resource.TestCheckResourceAttrSet(resourceName, "time_status_begin"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + RecommendationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Optional, Update, recommendationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "category_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "description"),
				resource.TestCheckResourceAttrSet(resourceName, "estimated_cost_saving"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "importance"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttrSet(resourceName, "recommendation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_counts.#"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISMISSED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_status_begin"),

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
			Config: config + compartmentIdVariableStr +
				GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", Required, Create, optimizerCategoryDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_optimizer_recommendations", "test_recommendations", Required, Create, recommendationDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "category_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "recommendation_collection.0.items.0.name"),
				resource.TestCheckResourceAttr(datasourceName, "recommendation_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "recommendation_collection.0.items.0.status", "DISMISSED"),

				resource.TestCheckResourceAttrSet(datasourceName, "recommendation_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "recommendation_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RecommendationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recommendation_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "category_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_cost_saving"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "importance"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_counts.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISMISSED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "supported_levels.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_status_begin"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + RecommendationResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				// Resource counts is dynamically computed during each GET call and is returned as a list so the order might not always be the same
				"resource_counts",
			},
			ResourceName: resourceName,
		},
	})
}
