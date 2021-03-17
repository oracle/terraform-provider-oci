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
		generateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationRepresentation)

	RecommendationResourceConfig = RecommendationResourceDependencies +
		generateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Optional, Update, recommendationRepresentation)

	recommendationSingularDataSourceRepresentation = map[string]interface{}{
		"recommendation_id": Representation{repType: Required, create: `${data.oci_optimizer_recommendations.test_recommendations.recommendation_collection.0.items.0.id}`},
	}

	recommendationDataSourceRepresentation = map[string]interface{}{
		"category_id":               Representation{repType: Required, create: `${lookup(data.oci_optimizer_categories.test_categories.category_collection.0.items[0], "id")}`},
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{repType: Required, create: `true`},
		"name":                      Representation{repType: Optional, create: `name`},
		"state":                     Representation{repType: Optional, create: `ACTIVE`},
		"status":                    Representation{repType: Optional, create: `PENDING`, update: `DISMISSED`},
		"filter":                    RepresentationGroup{Required, recommendationDataSourceFilterRepresentation}}

	// to filter the list of recommendation and get one recommendation containing the supportlevels
	// we will use the supportlevels to create the optimizer profile
	recommendationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`cost-management-compute-host-underutilized-name`}},
	}

	recommendationRepresentation = map[string]interface{}{
		"recommendation_id": Representation{repType: Required, create: `${data.oci_optimizer_recommendations.test_recommendations.recommendation_collection.0.items.0.id}`},
		"status":            Representation{repType: Required, create: `PENDING`, update: `DISMISSED`},
	}

	RecommendationResourceDependencies = generateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", Required, Create, optimizerCategoryDataSourceRepresentation) +
		generateDataSourceFromRepresentationMap("oci_optimizer_recommendations", "test_recommendations", Required, Create, recommendationDataSourceRepresentation)
)

func TestOptimizerRecommendationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerRecommendationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_optimizer_recommendation.test_recommendation"
	datasourceName := "data.oci_optimizer_recommendations.test_recommendations"
	singularDatasourceName := "data.oci_optimizer_recommendation.test_recommendation"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+RecommendationResourceDependencies+
		generateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Optional, Create, recommendationRepresentation), "optimizer", "recommendation", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + RecommendationResourceDependencies +
					generateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "recommendation_id"),
					resource.TestCheckResourceAttr(resourceName, "status", "PENDING"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RecommendationResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + RecommendationResourceDependencies +
					generateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Optional, Create, recommendationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Optional, Update, recommendationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config: config + compartmentIdVariableStr +
					generateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", Required, Create, optimizerCategoryDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_optimizer_recommendations", "test_recommendations", Required, Create, recommendationDataSourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + RecommendationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}
