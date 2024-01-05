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
	OptimizerOptimizerResourceRecommendationDataSourceRepresentationForResourceAction = map[string]interface{}{
		"category_id":               acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_optimizer_categories.test_categories.category_collection.0.items[0], "id")}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"include_organization":      acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"include_resource_metadata": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: `PENDING`, Update: `DISMISSED`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: OptimizerOptimizerResourceRecommendationDataSourceFilterForResourceAction}}

	OptimizerOptimizerResourceRecommendationDataSourceFilterForResourceAction = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`cost-management-object-storage-enable-olm-name`}},
	}

	OptimizerOptimizerResourceActionRequiredOnlyResource = OptimizerResourceActionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", acctest.Required, acctest.Create, OptimizerResourceActionRepresentation)

	ResourceActionResourceConfig = OptimizerResourceActionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", acctest.Optional, acctest.Update, OptimizerResourceActionRepresentation)

	OptimizerOptimizerResourceActionSingularDataSourceRepresentation = map[string]interface{}{
		"resource_action_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_optimizer_resource_actions.test_resource_actions.resource_action_collection.0.items.0.id}`},
	}

	OptimizerOptimizerResourceActionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"recommendation_id":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_optimizer_recommendations.test_recommendations.recommendation_collection.0.items.0.id}`},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"resource_type":             acctest.Representation{RepType: acctest.Optional, Create: `resourceType`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: `PENDING`, Update: `DISMISSED`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: OptimizerResourceActionDataSourceFilterRepresentation}}

	OptimizerResourceActionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `status`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`PENDING`, `DISMISSED`, `POSTPONED`}},
	}

	OptimizerResourceActionRepresentation = map[string]interface{}{
		"resource_action_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_optimizer_resource_actions.test_resource_actions.resource_action_collection.0.items.0.id}`},
		"status":             acctest.Representation{RepType: acctest.Required, Create: `PENDING`, Update: `DISMISSED`},
	}

	OptimizerResourceActionResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", acctest.Required, acctest.Create, OptimizerOptimizerCategoryDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_recommendations", "test_recommendations", acctest.Required, acctest.Create, OptimizerOptimizerResourceRecommendationDataSourceRepresentationForResourceAction) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_resource_actions", "test_resource_actions", acctest.Required, acctest.Create, OptimizerOptimizerResourceActionDataSourceRepresentation)
)

// issue-routing-tag: optimizer/default
func TestOptimizerResourceActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerResourceActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_optimizer_resource_action.test_resource_action"
	datasourceName := "data.oci_optimizer_resource_actions.test_resource_actions"
	singularDatasourceName := "data.oci_optimizer_resource_action.test_resource_action"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OptimizerResourceActionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", acctest.Optional, acctest.Create, OptimizerResourceActionRepresentation), "optimizer", "resourceAction", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OptimizerResourceActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", acctest.Required, acctest.Create, OptimizerResourceActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "resource_action_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "PENDING"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OptimizerResourceActionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OptimizerResourceActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", acctest.Optional, acctest.Create, OptimizerResourceActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "action.#"),
				resource.TestCheckResourceAttrSet(resourceName, "category_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "estimated_cost_saving"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttrSet(resourceName, "recommendation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_action_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_type"),
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
			Config: config + compartmentIdVariableStr + OptimizerResourceActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", acctest.Optional, acctest.Update, OptimizerResourceActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_name"),
				resource.TestCheckResourceAttrSet(resourceName, "estimated_cost_saving"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttrSet(resourceName, "recommendation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_action_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_type"),
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
			Config: config + compartmentIdVariableStr + OptimizerResourceActionResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "recommendation_id"),

				resource.TestCheckResourceAttr(datasourceName, "resource_action_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_action_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_action_collection.0.items.0.recommendation_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_action_collection.0.items.0.resource_type"),
				resource.TestCheckResourceAttr(datasourceName, "resource_action_collection.0.items.0.state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "resource_action_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_action_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + ResourceActionResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", acctest.Required, acctest.Create, OptimizerOptimizerResourceActionSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_action_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "action.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "category_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_cost_saving"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_status_begin"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OptimizerOptimizerResourceActionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
