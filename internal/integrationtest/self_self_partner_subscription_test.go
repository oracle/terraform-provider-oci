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
	SelfSelfPartnerSubscriptionDataSourceRepresentation = map[string]interface{}{
		"listing_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.product_id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	SelfSelfPartnerSubscriptionResourceConfig = ""
)

// issue-routing-tag: self/default
func TestSelfSelfPartnerSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSelfSelfPartnerSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	productId := utils.GetEnvSettingWithBlankDefault("product_id")
	productIdVariableStr := fmt.Sprintf("variable \"product_id\" { default = \"%s\" }\n", productId)

	datasourceName := "data.oci_self_self_partner_subscriptions.test_self_partner_subscriptions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_self_self_partner_subscriptions", "test_self_partner_subscriptions", acctest.Required, acctest.Create, SelfSelfPartnerSubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + productIdVariableStr + SelfSelfPartnerSubscriptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "listing_id", productId),

				resource.TestCheckResourceAttrSet(datasourceName, "listing_subscriptions_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "listing_subscriptions_collection.0.items.0.product_id", productId),
			),
		},
	})
}
