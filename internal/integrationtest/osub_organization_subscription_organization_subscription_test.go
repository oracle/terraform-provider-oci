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
	OsubOrganizationSubscriptionOsubOrganizationSubscriptionOrganizationSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_ids":    acctest.Representation{RepType: acctest.Required, Create: `subscriptionIds`},
		"x_one_origin_region": acctest.Representation{RepType: acctest.Required, Create: `${var.region}`},
	}

	OsubOrganizationSubscriptionOrganizationSubscriptionResourceConfig = ""
)

// issue-routing-tag: osub_organization_subscription/default
func TestOsubOrganizationSubscriptionOrganizationSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsubOrganizationSubscriptionOrganizationSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	oneRegion := utils.GetEnvSettingWithBlankDefault("region")
	oneRegionVariableStr := fmt.Sprintf("variable \"x_one_origin_region\" { default = \"%s\" }\n", oneRegion)

	datasourceName := "data.oci_osub_organization_subscription_organization_subscriptions.test_organization_subscriptions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osub_organization_subscription_organization_subscriptions", "test_organization_subscriptions", acctest.Required, acctest.Create, OsubOrganizationSubscriptionOsubOrganizationSubscriptionOrganizationSubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + oneRegionVariableStr + OsubOrganizationSubscriptionOrganizationSubscriptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_ids", "subscriptionIds"),
				resource.TestCheckResourceAttrSet(datasourceName, "x_one_origin_region"),

				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.#"),
				resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.currency.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.time_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.total_value"),
			),
		},
	})
}
