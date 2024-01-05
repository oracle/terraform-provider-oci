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
	OnesubscriptionOnesubscriptionRatecardDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"part_number":     acctest.Representation{RepType: acctest.Optional, Create: `partNumber`},
		"time_from":       acctest.Representation{RepType: acctest.Optional, Create: `timeFrom`},
		"time_to":         acctest.Representation{RepType: acctest.Optional, Create: `timeTo`},
	}
)

// issue-routing-tag: onesubscription/default
func TestOnesubscriptionRatecardResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOnesubscriptionRatecardResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_onesubscription_ratecards.test_ratecards"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_ratecards", "test_ratecards", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionRatecardDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.#"),
				resource.TestCheckResourceAttr(datasourceName, "rate_cards.0.currency.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.discretionary_discount_percentage"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.is_tier"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.net_unit_price"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.overage_price"),
				resource.TestCheckResourceAttr(datasourceName, "rate_cards.0.product.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.subscribed_service_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "rate_cards.0.time_start"),
			),
		},
	})
}
