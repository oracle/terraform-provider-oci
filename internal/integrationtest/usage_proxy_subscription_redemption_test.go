// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	subscriptionRedemptionSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	subscriptionRedemptionDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	SubscriptionRedemptionResourceConfig = ""
)

// issue-routing-tag: usage_proxy/default
func TestUsageProxySubscriptionRedemptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestUsageProxySubscriptionRedemptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_usage_proxy_subscription_redemptions.test_subscription_redemptions"
	singularDatasourceName := "data.oci_usage_proxy_subscription_redemption.test_subscription_redemption"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_redemptions", "test_subscription_redemptions", acctest.Required, acctest.Create, subscriptionRedemptionDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + SubscriptionRedemptionResourceConfig,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tenancy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "redemption_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_redemption", "test_subscription_redemption", acctest.Required, acctest.Create, subscriptionRedemptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + SubscriptionRedemptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
	})
}
