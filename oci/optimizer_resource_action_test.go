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
		"category_id":               Representation{repType: Required, create: `${lookup(data.oci_optimizer_categories.test_categories.category_collection.0.items[0], "id")}`},
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{repType: Required, create: `true`},
		"name":                      Representation{repType: Optional, create: `name`},
		"state":                     Representation{repType: Optional, create: `ACTIVE`},
		"status":                    Representation{repType: Optional, create: `PENDING`, update: `DISMISSED`},
		"filter":                    RepresentationGroup{Required, recommendationDataSourceFilterForResourceAction}}

	recommendationDataSourceFilterForResourceAction = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`cost-management-object-storage-enable-olm-name`}},
	}

	ResourceActionRequiredOnlyResource = ResourceActionResourceDependencies +
		generateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Required, Create, resourceActionRepresentation)

	ResourceActionResourceConfig = ResourceActionResourceDependencies +
		generateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Optional, Update, resourceActionRepresentation)

	resourceActionSingularDataSourceRepresentation = map[string]interface{}{
		"resource_action_id": Representation{repType: Required, create: `${data.oci_optimizer_resource_actions.test_resource_actions.resource_action_collection.0.items.0.id}`},
	}

	resourceActionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{repType: Required, create: `true`},
		"recommendation_id":         Representation{repType: Required, create: `${data.oci_optimizer_recommendations.test_recommendations.recommendation_collection.0.items.0.id}`},
		"name":                      Representation{repType: Optional, create: `name`},
		"resource_type":             Representation{repType: Optional, create: `resourceType`},
		"state":                     Representation{repType: Optional, create: `ACTIVE`},
		"status":                    Representation{repType: Optional, create: `PENDING`, update: `DISMISSED`},
		"filter":                    RepresentationGroup{Required, resourceActionDataSourceFilterRepresentation}}

	resourceActionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `status`},
		"values": Representation{repType: Required, create: []string{`PENDING`, `DISMISSED`, `POSTPONED`}},
	}

	resourceActionRepresentation = map[string]interface{}{
		"resource_action_id": Representation{repType: Required, create: `${data.oci_optimizer_resource_actions.test_resource_actions.resource_action_collection.0.items.0.id}`},
		"status":             Representation{repType: Required, create: `PENDING`, update: `DISMISSED`},
	}

	ResourceActionResourceDependencies = generateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", Required, Create, optimizerCategoryDataSourceRepresentation) +
		generateDataSourceFromRepresentationMap("oci_optimizer_recommendations", "test_recommendations", Required, Create, recommendationDataSourceRepresentationForResourceAction) +
		generateDataSourceFromRepresentationMap("oci_optimizer_resource_actions", "test_resource_actions", Required, Create, resourceActionDataSourceRepresentation)
)

func TestOptimizerResourceActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerResourceActionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_optimizer_resource_action.test_resource_action"
	datasourceName := "data.oci_optimizer_resource_actions.test_resource_actions"
	singularDatasourceName := "data.oci_optimizer_resource_action.test_resource_action"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ResourceActionResourceDependencies+
		generateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Optional, Create, resourceActionRepresentation), "optimizer", "resourceAction", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ResourceActionResourceDependencies +
					generateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Required, Create, resourceActionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "resource_action_id"),
					resource.TestCheckResourceAttr(resourceName, "status", "PENDING"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ResourceActionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ResourceActionResourceDependencies +
					generateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Optional, Create, resourceActionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config: config + compartmentIdVariableStr + ResourceActionResourceDependencies +
					generateResourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Optional, Update, resourceActionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
				Config: config + compartmentIdVariableStr + ResourceActionResourceDependencies,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_optimizer_resource_action", "test_resource_action", Required, Create, resourceActionSingularDataSourceRepresentation) +
					compartmentIdVariableStr,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}
