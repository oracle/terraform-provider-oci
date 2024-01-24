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
	OsubSubscriptionOsubSubscriptionRatecardDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"part_number":         acctest.Representation{RepType: acctest.Optional, Create: `partNumber`},
		"time_from":           acctest.Representation{RepType: acctest.Optional, Create: `timeFrom`},
		"time_to":             acctest.Representation{RepType: acctest.Optional, Create: `timeTo`},
		"x_one_origin_region": acctest.Representation{RepType: acctest.Required, Create: `${var.x_one_origin_region}`},
	}
)

// issue-routing-tag: osub_subscription/default
func TestOsubSubscriptionRatecardResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsubSubscriptionRatecardResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	oneRegion := utils.GetEnvSettingWithBlankDefault("x_one_origin_region")
	oneRegionVariableStr := fmt.Sprintf("variable \"x_one_origin_region\"  { default = \"%s\" }\n", oneRegion)

	datasourceName := "data.oci_osub_subscription_ratecards.test_ratecards"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osub_subscription_ratecards", "test_ratecards", acctest.Required, acctest.Create, OsubSubscriptionOsubSubscriptionRatecardDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + oneRegionVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "x_one_origin_region"),

				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.#"),
				resource.TestCheckResourceAttr(datasourceName, "rate_cards.0.currency.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.discretionary_discount_percentage"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.is_tier"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.net_unit_price"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.overage_price"),
				resource.TestCheckResourceAttr(datasourceName, "rate_cards.0.product.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.time_start"),
			),
		},
	})
}
