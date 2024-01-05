// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OsubUsageOsubUsageComputedUsageAggregatedDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"time_from":           acctest.Representation{RepType: acctest.Required, Create: `2019-11-20T08:00:00Z`},
		"time_to":             acctest.Representation{RepType: acctest.Required, Create: `2019-11-20T23:59:59Z`},
		"grouping":            acctest.Representation{RepType: acctest.Required, Create: `MONTHLY`},
		"parent_product":      acctest.Representation{RepType: acctest.Required, Create: `${var.parent_product}`},
		"x_one_origin_region": acctest.Representation{RepType: acctest.Required, Create: `${var.region}`},
	}
	OsubUsageComputedUsageAggregatedParentProductDataSourceRepresentation = map[string]interface{}{}

	OsubUsageComputedUsageAggregatedResourceConfig = ""
)

// issue-routing-tag: osub_usage/default
func TestOsubUsageComputedUsageAggregatedResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsubUsageComputedUsageAggregatedResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	parentProduct := utils.GetEnvSettingWithBlankDefault("parent_product")
	parentProductVariableStr := fmt.Sprintf("variable \"parent_product\" { default = \"%s\" }\n", parentProduct)

	oneRegion := utils.GetEnvSettingWithBlankDefault("region")
	oneRegionVariableStr := fmt.Sprintf("variable \"x_one_origin_region\" { default = \"%s\" }\n", oneRegion)

	datasourceName := "data.oci_osub_usage_computed_usage_aggregateds.test_computed_usage_aggregateds"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osub_usage_computed_usage_aggregateds", "test_computed_usage_aggregateds", acctest.Required, acctest.Create, OsubUsageOsubUsageComputedUsageAggregatedDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + parentProductVariableStr + oneRegionVariableStr + OsubUsageComputedUsageAggregatedResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "grouping"),
				resource.TestCheckResourceAttr(datasourceName, "parent_product", parentProduct),
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttrSet(datasourceName, "time_from"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_to"),

				resource.TestCheckResourceAttrSet(datasourceName, "computed_usage_aggregateds.#"),
				resource.TestCheckResourceAttr(datasourceName, "computed_usage_aggregateds.0.aggregated_computed_usages.#", "3"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usage_aggregateds.0.currency_code"),
				resource.TestCheckResourceAttr(datasourceName, "computed_usage_aggregateds.0.parent_product.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usage_aggregateds.0.parent_subscribed_service_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usage_aggregateds.0.plan_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usage_aggregateds.0.pricing_model"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usage_aggregateds.0.rate_card_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usage_aggregateds.0.subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usage_aggregateds.0.time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usage_aggregateds.0.time_start"),
			),
		},
	})
}
