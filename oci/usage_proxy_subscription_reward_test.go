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
	subscriptionRewardSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": Representation{RepType: Required, Create: `${var.subscription_id}`},
		"tenancy_id":      Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
	}

	subscriptionRewardDataSourceRepresentation = map[string]interface{}{
		"subscription_id": Representation{RepType: Required, Create: `${var.subscription_id}`},
		"tenancy_id":      Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
	}

	SubscriptionRewardResourceConfig = ""
)

// issue-routing-tag: usage_proxy/default
func TestUsageProxySubscriptionRewardResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestUsageProxySubscriptionRewardResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := getEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_usage_proxy_subscription_rewards.test_subscription_rewards"
	singularDatasourceName := "data.oci_usage_proxy_subscription_reward.test_subscription_reward"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_rewards", "test_subscription_rewards", Required, Create, subscriptionRewardDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + SubscriptionRewardResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_reward", "test_subscription_reward", Required, Create, subscriptionRewardSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + SubscriptionRewardResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "summary.#", "1"),
			),
		},
	})
}
