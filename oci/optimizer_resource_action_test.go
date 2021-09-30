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
	recommendationDataSourceRepresentationForResourceAction = map[string]interface{}{
		"category_id":               Representation{RepType: Required, Create: `${lookup(data.oci_optimizer_categories.test_categories.category_collection.0.items[0], "id")}`},
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{RepType: Required, Create: `true`},
		"name":                      Representation{RepType: Optional, Create: `name`},
		"state":                     Representation{RepType: Optional, Create: `ACTIVE`},
		"status":                    Representation{RepType: Optional, Create: `PENDING`, Update: `DISMISSED`},
		"filter":                    RepresentationGroup{Required, recommendationDataSourceFilterForResourceAction}}

	recommendationDataSourceFilterForResourceAction = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `name`},
		"values": Representation{RepType: Required, Create: []string{`cost-management-object-storage-enable-olm-name`}},
	}

	ResourceActionRequiredOnlyResource = ResourceActionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Required, Create, resourceActionRepresentation)

	ResourceActionResourceConfig = ResourceActionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Optional, Update, resourceActionRepresentation)

	resourceActionSingularDataSourceRepresentation = map[string]interface{}{
		"resource_action_id": Representation{RepType: Required, Create: `${data.oci_optimizer_resource_actions.test_resource_actions.resource_action_collection.0.items.0.id}`},
	}

	resourceActionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{RepType: Required, Create: `true`},
		"recommendation_id":         Representation{RepType: Required, Create: `${data.oci_optimizer_recommendations.test_recommendations.recommendation_collection.0.items.0.id}`},
		"name":                      Representation{RepType: Optional, Create: `name`},
		"resource_type":             Representation{RepType: Optional, Create: `resourceType`},
		"state":                     Representation{RepType: Optional, Create: `ACTIVE`},
		"status":                    Representation{RepType: Optional, Create: `PENDING`, Update: `DISMISSED`},
		"filter":                    RepresentationGroup{Required, resourceActionDataSourceFilterRepresentation}}

	resourceActionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `status`},
		"values": Representation{RepType: Required, Create: []string{`PENDING`, `DISMISSED`, `POSTPONED`}},
	}

	resourceActionRepresentation = map[string]interface{}{
		"resource_action_id": Representation{RepType: Required, Create: `${data.oci_optimizer_resource_actions.test_resource_actions.resource_action_collection.0.items.0.id}`},
		"status":             Representation{RepType: Required, Create: `PENDING`, Update: `DISMISSED`},
	}

	ResourceActionResourceDependencies = GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", Required, Create, optimizerCategoryDataSourceRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_optimizer_recommendations", "test_recommendations", Required, Create, recommendationDataSourceRepresentationForResourceAction) +
		GenerateDataSourceFromRepresentationMap("oci_optimizer_resource_actions", "test_resource_actions", Required, Create, resourceActionDataSourceRepresentation)
)

// issue-routing-tag: optimizer/default
func TestOptimizerResourceActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerResourceActionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_optimizer_resource_action.test_resource_action"
	datasourceName := "data.oci_optimizer_resource_actions.test_resource_actions"
	singularDatasourceName := "data.oci_optimizer_resource_action.test_resource_action"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ResourceActionResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Optional, Create, resourceActionRepresentation), "optimizer", "resourceAction", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ResourceActionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Required, Create, resourceActionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "resource_action_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "PENDING"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ResourceActionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ResourceActionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Optional, Create, resourceActionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + ResourceActionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Optional, Update, resourceActionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + ResourceActionResourceDependencies,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_action_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_action_collection.0.items.0.recommendation_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_action_collection.0.items.0.resource_type"),
				resource.TestCheckResourceAttr(datasourceName, "resource_action_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "resource_action_collection.0.items.0.status", "DISMISSED"),

				resource.TestCheckResourceAttrSet(datasourceName, "resource_action_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_action_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + ResourceActionResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Required, Create, resourceActionSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISMISSED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_status_begin"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ResourceActionResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
