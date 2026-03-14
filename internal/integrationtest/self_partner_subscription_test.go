// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	SelfPartnerSubscriptionDataSourceRepresentation = map[string]interface{}{
		"listing_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.listing_id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}
	SelfPartnerSubscriptionResourceConfig = ""
)

// issue-routing-tag: self/default
func TestSelfPartnerSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSelfPartnerSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	listingId := utils.GetEnvSettingWithBlankDefault("listing_id")
	listingIdVariableStr := fmt.Sprintf("variable \"listing_id\" { default = \"%s\" }\n", listingId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_self_partner_subscriptions.test_partner_subscriptions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_self_partner_subscriptions", "test_partner_subscriptions", acctest.Required, acctest.Create, SelfPartnerSubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + SelfPartnerSubscriptionResourceConfig + listingIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "listing_subscriptions_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "listing_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "listing_subscriptions_collection.#"),
				resource.TestCheckResourceAttrWith(datasourceName, "listing_subscriptions_collection.0.items.#", func(value string) error {
					count, err := strconv.Atoi(value)
					if err != nil {
						return fmt.Errorf("expected a number, got %s", value)
					}
					if count < 1 {
						return fmt.Errorf("expected at least 1 item, got %d", count)
					}
					return nil
				}),
			),
		},
	})
}
