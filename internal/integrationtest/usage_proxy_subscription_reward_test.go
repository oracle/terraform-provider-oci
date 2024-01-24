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
	UsageProxyUsageProxySubscriptionRewardSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	UsageProxyUsageProxySubscriptionRewardDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	UsageProxySubscriptionRewardResourceConfig = ""
)

// issue-routing-tag: usage_proxy/default
func TestUsageProxySubscriptionRewardResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestUsageProxySubscriptionRewardResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_usage_proxy_subscription_rewards.test_subscription_rewards"
	singularDatasourceName := "data.oci_usage_proxy_subscription_reward.test_subscription_reward"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_rewards", "test_subscription_rewards", acctest.Required, acctest.Create, UsageProxyUsageProxySubscriptionRewardDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + UsageProxySubscriptionRewardResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tenancy_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "reward_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "reward_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "reward_collection.0.items.0.summary.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_reward", "test_subscription_reward", acctest.Required, acctest.Create, UsageProxyUsageProxySubscriptionRewardSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + UsageProxySubscriptionRewardResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "summary.#", "1"),
			),
		},
	})
}
