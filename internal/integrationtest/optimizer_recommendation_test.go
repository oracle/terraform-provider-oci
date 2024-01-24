// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OptimizerRecommendationRequiredOnlyResource = OptimizerRecommendationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Required, acctest.Create, OptimizerRecommendationRepresentation)

	OptimizerRecommendationResourceConfig = OptimizerRecommendationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Optional, acctest.Update, OptimizerRecommendationRepresentation)

	OptimizerOptimizerRecommendationSingularDataSourceRepresentation = map[string]interface{}{
		"recommendation_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_optimizer_recommendations.test_recommendations.recommendation_collection.0.items.0.id}`},
	}

	OptimizerOptimizerRecommendationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"category_id":               acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_optimizer_categories.test_categories.category_collection.0.items[0], "id")}`},
		"include_organization":      acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: `PENDING`, Update: `DISMISSED`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: OptimizerRecommendationDataSourceFilterRepresentation}}

	// to filter the list of recommendation and get one recommendation containing the supportlevels
	// we will use the supportlevels to Create the optimizer profile
	OptimizerRecommendationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`cost-management-compute-host-underutilized-name`}},
	}

	OptimizerRecommendationRepresentation = map[string]interface{}{
		"recommendation_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_optimizer_recommendations.test_recommendations.recommendation_collection.0.items.0.id}`},
		"status":            acctest.Representation{RepType: acctest.Required, Create: `PENDING`, Update: `DISMISSED`},
	}

	OptimizerRecommendationResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", acctest.Required, acctest.Create, OptimizerOptimizerCategoryDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_recommendations", "test_recommendations", acctest.Required, acctest.Create, OptimizerOptimizerRecommendationDataSourceRepresentation)
)

// issue-routing-tag: optimizer/default
func TestOptimizerRecommendationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerRecommendationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_optimizer_recommendation.test_recommendation"
	datasourceName := "data.oci_optimizer_recommendations.test_recommendations"
	singularDatasourceName := "data.oci_optimizer_recommendation.test_recommendation"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OptimizerRecommendationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Optional, acctest.Create, OptimizerRecommendationRepresentation), "optimizer", "recommendation", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OptimizerRecommendationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Required, acctest.Create, OptimizerRecommendationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "recommendation_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "PENDING"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OptimizerRecommendationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OptimizerRecommendationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Optional, acctest.Create, OptimizerRecommendationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + OptimizerRecommendationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Optional, acctest.Update, OptimizerRecommendationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", acctest.Required, acctest.Create, OptimizerOptimizerCategoryDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_recommendations", "test_recommendations", acctest.Required, acctest.Create, OptimizerOptimizerRecommendationDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", acctest.Required, acctest.Create, OptimizerOptimizerRecommendationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OptimizerRecommendationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
		// verify resource import
		{
			Config:            config + OptimizerRecommendationRequiredOnlyResource,
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
