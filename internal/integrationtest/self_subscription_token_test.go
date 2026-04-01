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
	SelfSubscriptionTokenSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
	}

	//SelfSubscriptionTokenResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_self_products", "test_products", acctest.Required, acctest.Create, SelfProductSingularDataSourceRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_self_subscription", "test_subscription", acctest.Required, acctest.Create, SelfSubscriptionRepresentation)
	SelfSubscriptionTokenResourceConfig = ""
)

// issue-routing-tag: self/default
func TestSelfSubscriptionTokenResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSelfSubscriptionTokenResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_self_subscription_token.test_subscription_token"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_self_subscription_token", "test_subscription_token", acctest.Required, acctest.Create, SelfSubscriptionTokenSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SelfSubscriptionTokenResourceConfig + subscriptionIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "token"),
			),
		},
	})
}
