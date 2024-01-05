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
	computedUsageSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"computed_usage_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.computed_usage_id}`},
		"x_one_origin_region": acctest.Representation{RepType: acctest.Required, Create: `${var.region}`},
	}

	computedUsageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"time_from":           acctest.Representation{RepType: acctest.Required, Create: `2019-11-22T08:00:00Z`},
		"time_to":             acctest.Representation{RepType: acctest.Required, Create: `2019-11-22T23:59:59Z`},
		"computed_product":    acctest.Representation{RepType: acctest.Required, Create: `${var.computed_product}`},
		"parent_product":      acctest.Representation{RepType: acctest.Required, Create: `${var.parent_product}`},
		"x_one_origin_region": acctest.Representation{RepType: acctest.Required, Create: `${var.region}`},
	}
	computedUsageParentProductDataSourceRepresentation = map[string]interface{}{}

	ComputedUsageResourceConfig = ""
)

// issue-routing-tag: osub_usage/default
func TestOsubUsageComputedUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsubUsageComputedUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	computedUsageId := utils.GetEnvSettingWithBlankDefault("computed_usage_id")
	computedUsageIdVariableId := fmt.Sprintf("variable \"computed_usage_id\" { default = \"%s\" }\n", computedUsageId)

	oneRegion := utils.GetEnvSettingWithBlankDefault("region")
	oneRegionVariableStr := fmt.Sprintf("variable \"x_one_origin_region\" { default = \"%s\" }\n", oneRegion)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	computedProduct := utils.GetEnvSettingWithBlankDefault("computed_product")
	computedProductVariableStr := fmt.Sprintf("variable \"computed_product\" { default = \"%s\" }\n", computedProduct)

	parentProduct := utils.GetEnvSettingWithBlankDefault("parent_product")
	parentProductVariableStr := fmt.Sprintf("variable \"parent_product\" { default = \"%s\" }\n", parentProduct)

	datasourceName := "data.oci_osub_usage_computed_usages.test_computed_usages"
	singularDatasourceName := "data.oci_osub_usage_computed_usage.test_computed_usage"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osub_usage_computed_usages", "test_computed_usages", acctest.Required, acctest.Create, computedUsageDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + computedProductVariableStr + parentProductVariableStr + oneRegionVariableStr + ComputedUsageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "computed_product", computedProduct),
				resource.TestCheckResourceAttr(datasourceName, "parent_product", parentProduct),
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionId),
				resource.TestCheckResourceAttrSet(datasourceName, "time_from"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_to"),
				resource.TestCheckResourceAttr(datasourceName, "x_one_origin_region", oneRegion),

				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.computed_usage_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.cost"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.cost_rounded"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.currency_code"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.data_center"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.is_invoiced"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.mqs_message_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.net_unit_price"),
				resource.TestCheckResourceAttr(datasourceName, "computed_usages.0.parent_product.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.parent_subscribed_service_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.plan_number"),
				resource.TestCheckResourceAttr(datasourceName, "computed_usages.0.product.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.quantity"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.rate_card_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.time_metered_on"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.time_of_arrival"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "computed_usages.0.unit_of_measure"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osub_usage_computed_usage", "test_computed_usage", acctest.Required, acctest.Create, computedUsageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + computedUsageIdVariableId + oneRegionVariableStr + ComputedUsageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "computed_usage_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "x_one_origin_region"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "commitment_service_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost_rounded"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "currency_code"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_center"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_invoiced"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mqs_message_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "net_unit_price"),
				resource.TestCheckResourceAttr(singularDatasourceName, "parent_product.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_subscribed_service_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan_number"),
				resource.TestCheckResourceAttr(singularDatasourceName, "product.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "quantity"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "rate_card_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_metered_on"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_arrival"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "unit_of_measure"),
			),
		},
	})
}
