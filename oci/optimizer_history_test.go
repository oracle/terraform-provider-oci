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
	historyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"compartment_id_in_subtree": Representation{repType: Required, create: `true`},
		"name":                      Representation{repType: Optional, create: `name`},
		"recommendation_id":         Representation{repType: Optional, create: `${oci_optimizer_recommendation.test_recommendation.recommendation_id}`},
		"recommendation_name":       Representation{repType: Optional, create: `${oci_optimizer_recommendation.test_recommendation.name}`},
		"resource_type":             Representation{repType: Optional, create: `resourceType`},
		"state":                     Representation{repType: Optional, create: `ACTIVE`},
		"status":                    Representation{repType: Optional, create: `PENDING`},
	}

	HistoryResourceConfig = RecommendationResourceDependencies +
		generateResourceFromRepresentationMap("oci_optimizer_recommendation", "test_recommendation", Required, Create, recommendationRepresentation)
)

func TestOptimizerHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOptimizerHistoryResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_optimizer_histories.test_histories"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_optimizer_histories", "test_histories", Required, Create, historyDataSourceRepresentation) +
					compartmentIdVariableStr + HistoryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.category_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.compartment_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.estimated_cost_saving"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.recommendation_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.recommendation_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.resource_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.resource_action_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.resource_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.action.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.status"),
					resource.TestCheckResourceAttrSet(datasourceName, "history_collection.0.items.0.time_created"),
				),
			},
		},
	})
}
