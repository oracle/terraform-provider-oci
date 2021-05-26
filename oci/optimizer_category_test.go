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
	categorySingularDataSourceRepresentation = map[string]interface{}{
		"category_id": Representation{repType: Required, create: `${lookup(data.oci_optimizer_categories.test_categories.category_collection.0.items[0], "id")}`},
	}

	optimizerCategoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{repType: Required, create: `true`},
		"name":                      Representation{repType: Optional, create: `name`},
		"state":                     Representation{repType: Optional, create: `CREATED`},
		"filter":                    RepresentationGroup{Required, categoryDataSourceFilterRepresentation},
	}
	categoryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`cost-management-name`}},
	}

	OptimizerCategoryResourceConfig = ""
)

func TestOptimizerCategoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerCategoryResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_optimizer_categories.test_categories"
	singularDatasourceName := "data.oci_optimizer_category.test_category"

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
					generateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", Required, Create, optimizerCategoryDataSourceRepresentation) +
					compartmentIdVariableStr + OptimizerCategoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_optimizer_categories", "test_categories", Required, Create, optimizerCategoryDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_optimizer_category", "test_category", Required, Create, categorySingularDataSourceRepresentation) +
					compartmentIdVariableStr + OptimizerCategoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}
