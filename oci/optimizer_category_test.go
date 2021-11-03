// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	categorySingularDataSourceRepresentation = map[string]interface{}{
		"category_id": Representation{RepType: Required, Create: `${lookup(data.oci_optimizer_categories.test_categories.category_collection.0.items[0], "id")}`},
	}

	optimizerCategoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{RepType: Required, Create: `true`},
		"name":                      Representation{RepType: Optional, Create: `name`},
		"state":                     Representation{RepType: Optional, Create: `CREATED`},
		"filter":                    RepresentationGroup{Required, categoryDataSourceFilterRepresentation},
	}
	categoryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `name`},
		"values": Representation{RepType: Required, Create: []string{`cost-management-name`}},
	}

	OptimizerCategoryResourceConfig = ""
)

// issue-routing-tag: optimizer/default
func TestOptimizerCategoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerCategoryResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_optimizer_categories.test_categories"
	singularDatasourceName := "data.oci_optimizer_category.test_category"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", Required, Create, optimizerCategoryDataSourceRepresentation) +
				compartmentIdVariableStr + OptimizerCategoryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", Required, Create, optimizerCategoryDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_optimizer_category", "test_category", Required, Create, categorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + OptimizerCategoryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
