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
	OptimizerOptimizerCategorySingularDataSourceRepresentation = map[string]interface{}{
		"category_id": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_optimizer_categories.test_categories.category_collection.0.items[0], "id")}`},
	}

	OptimizerOptimizerCategoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"include_organization":      acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: OptimizerOptimizerCategoryDataSourceFilterRepresentation},
	}
	OptimizerOptimizerCategoryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`cost-management-name`}},
	}

	OptimizerCategoryResourceConfig = ""
)

// issue-routing-tag: optimizer/default
func TestOptimizerCategoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerCategoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	childTenancyId := utils.GetEnvSettingWithBlankDefault("child_tenancy_id")
	childTenancyIdVarStr := fmt.Sprintf("variable \"child_tenancy_id\" { default = \"%s\" }\n", childTenancyId)

	datasourceName := "data.oci_optimizer_categories.test_categories"
	singularDatasourceName := "data.oci_optimizer_category.test_category"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", acctest.Required, acctest.Create, OptimizerOptimizerCategoryDataSourceRepresentation) +
				compartmentIdVariableStr + OptimizerCategoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "category_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "category_collection.0.items.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "category_collection.#"),
			),
		},
		{
			Config: config + compartmentIdVariableStr + OptimizerCategoryResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", acctest.Optional, acctest.Create, OptimizerOptimizerCategoryDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "category_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "category_collection.0.items.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "category_collection.#"),
			),
		},
		{
			Config: config + compartmentIdVariableStr + OptimizerCategoryResourceConfig + childTenancyIdVarStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OptimizerOptimizerCategoryDataSourceRepresentation, map[string]interface{}{
						"child_tenancy_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.child_tenancy_id}`}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "category_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "category_collection.0.items.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "category_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", acctest.Required, acctest.Create, OptimizerOptimizerCategoryDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_optimizer_category", "test_category", acctest.Required, acctest.Create, OptimizerOptimizerCategorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + OptimizerCategoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "category_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_cost_saving"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recommendation_counts.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_counts.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
