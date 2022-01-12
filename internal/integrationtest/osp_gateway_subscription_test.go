// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	OspSubscriptionResourceConfig = OspSubscriptionResourceDependencies

	ospSubscriptionSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"osp_home_region": acctest.Representation{RepType: acctest.Required, Create: `${var.home_region}`},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_osp_gateway_subscriptions.test_subscriptions.subscription_collection.0.items[0], "id")}`},
	}

	ospSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"osp_home_region": acctest.Representation{RepType: acctest.Required, Create: `${var.home_region}`},
	}

	OspSubscriptionResourceDependencies = ""
)

// issue-routing-tag: osp_gateway/default
func TestOspGatewaySubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOspGatewaySubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	homeRegion := utils.GetEnvSettingWithBlankDefault("region")
	regionVariableStr := fmt.Sprintf("variable \"home_region\" { default = \"%s\" }\n", homeRegion)

	datasourceName := "data.oci_osp_gateway_subscriptions.test_subscriptions"
	singularDatasourceName := "data.oci_osp_gateway_subscription.test_subscription"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osp_gateway_subscriptions", "test_subscriptions", acctest.Required, acctest.Create, ospSubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + regionVariableStr + OspSubscriptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "osp_home_region", homeRegion),

				resource.TestCheckResourceAttr(datasourceName, "subscription_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscription_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osp_gateway_subscriptions", "test_subscriptions", acctest.Required, acctest.Create, ospSubscriptionDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osp_gateway_subscription", "test_subscription", acctest.Required, acctest.Create, ospSubscriptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + regionVariableStr + OspSubscriptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "osp_home_region", homeRegion),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "bill_to_cust_account_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "billing_address.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "currency_code"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gsi_org_code"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_intent_to_pay"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "language_code"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "organization_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "payment_gateway.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "payment_options.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ship_to_cust_acct_role_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ship_to_cust_acct_site_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_plan_number"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tax_info.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_start"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "upgrade_state"),
			),
		},
	})
}
