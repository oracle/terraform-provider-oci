// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	subscriptionProductSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"tenancy_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"usage_period_key": acctest.Representation{RepType: acctest.Required, Create: `${var.usage_period_key}`},
		"producttype":      acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
	}

	subscriptionProductDataSourceRepresentation = map[string]interface{}{
		"subscription_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"tenancy_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"usage_period_key": acctest.Representation{RepType: acctest.Required, Create: `${var.usage_period_key}`},
		"producttype":      acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
	}

	SubscriptionProductResourceConfig = ""
)

// issue-routing-tag: usage_proxy/default
func TestUsageProxySubscriptionProductResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestUsageProxySubscriptionProductResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	usagePeriodKey := utils.GetEnvSettingWithBlankDefault("usage_period_key")
	usagePeriodVariableStr := fmt.Sprintf("variable \"usage_period_key\" { default = \"%s\" }\n", usagePeriodKey)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_usage_proxy_subscription_products.test_subscription_products"
	singularDatasourceName := "data.oci_usage_proxy_subscription_product.test_subscription_product"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_products", "test_subscription_products", acctest.Required, acctest.Create, subscriptionProductDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + usagePeriodVariableStr + SubscriptionProductResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tenancy_id"),
				resource.TestCheckResourceAttr(datasourceName, "usage_period_key", "1890"),

				resource.TestCheckResourceAttrSet(datasourceName, "product_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "product_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_product", "test_subscription_product", acctest.Required, acctest.Create, subscriptionProductSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + usagePeriodVariableStr + SubscriptionProductResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "usage_period_key", "1890"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "5"),
			),
		},
	})
}
